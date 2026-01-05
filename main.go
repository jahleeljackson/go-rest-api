package main

import (
	"github.com/jahleeljackson/go-rest-api/controllers"
	"github.com/jahleeljackson/go-rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.POST("/logs", controllers.CreateLog)
	router.GET("/logs", controllers.FindLogs)
	router.GET("/logs/:id", controllers.FindLog)
	router.PATCH("/posts/:id", controllers.UpdateLog)
	router.DELETE("/posts/:id", controllers.DeletePost)

	router.Run("localhost:8080")
}
