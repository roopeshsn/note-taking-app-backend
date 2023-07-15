package migration


import (
    "note-taking-app-backend/utils"
	"note-taking-app-backend/models"
)

func Migrate() {
	utils.LoadEnv()
	utils.ConnectDB()
	utils.DB.AutoMigrate(&models.Note{})
	utils.DB.AutoMigrate(&models.User{})
}