package main

import (
	"first/httpd/api/services"
	"first/httpd/handler"
	"fmt"

	"github.com/gin-gonic/gin"
)

var server = services.Server{}

func main() {
	fmt.Println("First Go APP")

	r := gin.Default()

	r.Use(Cors())
	r.POST("/addToCart", handler.RunRulesOnCart)
	r.POST("/checkOut", handler.CheckOut)
	r.GET("/items", handler.GetAllItems)
	r.OPTIONS("/items")

	r.Run()
}

func Cors() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Access-Control-Allow-Methods", "DELECT,POST,PUT")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "Contect-Type")
		c.Next()
	}
}
