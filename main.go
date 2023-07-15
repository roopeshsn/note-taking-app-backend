package main

import (
	"os"

	"github.com/gin-gonic/gin"
    "note-taking-app-backend/utils"
	"note-taking-app-backend/migration"
	"note-taking-app-backend/controllers"
	"note-taking-app-backend/middleware"
)

func init() {
	utils.LoadEnv()
	utils.ConnectDB()
}

func main() {
	arg1 := os.Args[1]

	if arg1 == "m" {
        migration.Migrate()
    } else {
		r := gin.Default()
		r.POST("/signup", controllers.Signup)
		r.POST("/login", controllers.Login)
		r.POST("/notes", middleware.RequireAuth, controllers.CreateNote)
		r.GET("/notes",  middleware.RequireAuth, controllers.GetNotes)
		r.DELETE("/notes",  middleware.RequireAuth, controllers.DeleteNote)
		r.Run()
	}
}




