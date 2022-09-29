package main

import (
	"github.com/gin-gonic/gin"
	"github.com/polatyener-dev/golang_restapi/models"
	"github.com/polatyener-dev/golang_restapi/controllers/product"
)

func main() {
	r := gin.Default()
	models.ConnectDB()

	r.GET("api/products", product.Index)
	r.GET("api/products/:id", product.Show)
	r.POST("api/products", product.Create)
	r.PUT("api/products/:id", product.Update)
	r.DELETE("api/products", product.Delete)

	r.Run()
}