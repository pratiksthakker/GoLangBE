package handler

import (
	"first/db"
	"first/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllItems(c *gin.Context) {

	products := []models.Item{}
	database := db.OpenDbConnection().Table("item")
	//database.AutoMigrate(&product)
	database.Find(&products)
	//database.Find(&products)
	c.JSON(http.StatusOK, products)
	defer database.Close()

}
