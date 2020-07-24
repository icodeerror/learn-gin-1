package main

import (
	"github.com/gin-gonic/gin"
	"github.com/icodeerror/learn-gin-1/controllers"
	"github.com/icodeerror/learn-gin-1/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.Run()
}
