package main

import (
	"github.com/gin-gonic/gin"
	productcontroller "github.com/hadi25nov/go-rest-api-gin/controllers"
	"github.com/hadi25nov/go-rest-api-gin/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products/:id", productcontroller.Show)
	r.POST("/api/products", productcontroller.Create)
	r.PUT("/api/products/:id", productcontroller.Update)
	r.DELETE("/api/product/", productcontroller.Delete)

	r.Run(":8082")
}
