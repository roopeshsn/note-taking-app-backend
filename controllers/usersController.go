package controllers

import (
	"os"
	"net/http"
	"time"
	"golang.org/x/crypto/bcrypt"

	"note-taking-app-backend/utils"
	"note-taking-app-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Signup(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Create user row
	user := models.User{Name: req.Name, Email: req.Email, Password: string(hashedPassword)}
	result := utils.DB.Create(&user)

	if result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	c.Status(http.StatusOK)
}

func Login(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User
	if err := utils.DB.First(&user, "email = ?", req.Email).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Compare password to hashPassword
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sid": tokenString,
	})
}