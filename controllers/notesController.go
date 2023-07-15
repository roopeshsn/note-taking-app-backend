package controllers

import (
	"net/http"
	"note-taking-app-backend/utils"
	"note-taking-app-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateNote(c *gin.Context) {
	var req struct {
		Sid string `json:"sid" binding:"required"`
		Note string `json:"note" binding:"required"`
	}

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil  {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId, _ := c.Get("id")

	// We do not require Sid to be stored in the DB. 
	note := models.Note{Uid: userId.(uint), Note: req.Note}
	result := utils.DB.Create(&note)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create a note!",
		})
		return
	}
	c.Status(http.StatusOK)
}

func GetNotes(c *gin.Context) {
	userId, _ := c.Get("id")

	var notes []models.Note

	if err := utils.DB.Where("uid = ?", userId).Find(&notes).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes"})
        return
    }

    c.JSON(http.StatusOK, notes)
}

func DeleteNote(c *gin.Context) {
	userId, _ := c.Get("id")

	var req struct {
		Sid string `json:"sid" binding:"required"`
		Id uint `json:"id" binding:"required"`
	}

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil  {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request!",
		})
		return
	}

	if err := utils.DB.Where("id = ? AND uid = ?", req.Id, userId).Delete(&models.Note{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found or does not belong to the user!"})
		return
	}

	// if err := utils.DB.Delete(&models.Note{}, req.Id).Error; err != nil {
    //     c.JSON(http.StatusInternalServerError, err)
    //     return
    // }
	c.Status(http.StatusOK)
}