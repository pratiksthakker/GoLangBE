package handler

import (
	"first/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Addtocart(c *gin.Context) {

	//var rule models.Rules
	var cart []models.Item
	c.BindJSON(&cart)
	fmt.Println("got cart:", cart)
	//rules := GetRules
	//newCart := RunRulesOnCart
	//fmt.Println(newCart)
}
