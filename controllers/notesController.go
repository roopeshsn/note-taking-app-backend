package controllers

import (
	"net/http"
	"note-taking-app-backend/utils"
	"note-taking-app-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type NoteResponse struct {
	Id uint `json:"id"`
	Note string `json:"note"`
}

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

	var noteResponses []NoteResponse

	for _, note := range notes {
		noteResponses = append(noteResponses, NoteResponse{
			Id:   note.ID,
			Note: note.Note,
		})
	}

    c.JSON(http.StatusOK, gin.H{
		"notes": noteResponses,
	})
}

func DeleteNote(c *gin.Context) {
	userIdRaw, _ := c.Get("id")

	userId, ok := userIdRaw.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user ID",
		})
		return
	}

	var req struct {
		Sid string `json:"sid" binding:"required"`
		Id uint `json:"id" binding:"required"`
	}

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil  {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var count int64

	if err := utils.DB.Model(&models.User{}).
		Where("id = ? AND id IN (SELECT uid FROM notes WHERE id = ?)", userId, req.Id).
		Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to check note ownership",
		})
		return
	}

	if count == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Note does not exist or does not belong to the user",
		})
		return
	}

	// Delete the note using the association method
	if err := utils.DB.Where("id = ?", req.Id).Delete(&models.Note{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete the note",
		})
		return
	}

	c.Status(http.StatusOK)
}