package middleware

import (
	"os"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v5"
	"note-taking-app-backend/utils"
	"note-taking-app-backend/models"
)

func RequireAuth(c *gin.Context) {
	var req struct {
		Sid string `json:"sid" binding:"required"`
	}

	if c.ShouldBindBodyWith(&req, binding.JSON) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request!",
		})
		return
	}

	// Parse takes the token string and a function for looking up the key.
	token, err := jwt.Parse(req.Sid, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check expiration of sid
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Session expired!",
			})
			return
		}

		// Find user
		var user models.User
		if err := utils.DB.First(&user, claims["sub"]).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "User not found!",
			})
			return
		}

		// Attach user ID to the request
		c.Set("id", user.ID)
		c.Next()
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
}