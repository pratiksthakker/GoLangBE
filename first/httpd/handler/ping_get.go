package handler

import (
	"first/db"
	"first/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingGet() gin.HandlerFunc {

	return func(c *gin.Context) {
		var product models.Item
		products := []models.Item{}
		database := db.OpenDbConnection()
		database.AutoMigrate(&product)
		database.Raw("Select * from item").Scan(&products)
		//database.Find(&products)
		c.JSON(http.StatusOK, products)
		defer database.Close()
	}
}
