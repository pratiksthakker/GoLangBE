package handler

import (
	"first/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckOut(c *gin.Context) {

	var checkoutorder models.CheckoutOrder
	c.BindJSON(&checkoutorder)
	iscompleted, message := CompleteOrder(checkoutorder)
	if iscompleted {
		checkoutorder.Status = "Order Completed Successfully"
	} else {
		checkoutorder.Status = message + "   Please contact HelpDesk"
	}
	c.JSON(http.StatusOK, checkoutorder)

}
