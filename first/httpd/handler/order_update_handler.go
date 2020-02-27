package handler

import (
	"first/db"
	"first/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

func UpdateOrder(cartOrder models.CartOrder) int {

	var order models.DbOrder
	database := db.OpenDbConnection()
	fmt.Println("CartOrder", cartOrder)
	order.UserId = 123
	order.Discount = cartOrder.DiscountPercent
	if err := database.Table("orders").Where("id = ?", cartOrder.OrderId).First(&order).Error; err != nil {
		// error handling...
		if gorm.IsRecordNotFoundError(err) {
			var temporder models.DbOrder
			database.Table("orders").Last(&temporder)
			order.OrderId = temporder.OrderId + 1
			order.IsCompleted = "N"
			cartOrder.OrderId = order.OrderId
			fmt.Println("Inserting")
			fmt.Println("Order", order)
			database.Table("orders").Create(&order)
			UpdateOrderDetails(cartOrder)
		}
	} else {
		fmt.Println("Updating")
		//database.Table("orders").Model(&order).Where("id = ?", cartOrder.OrderId).Update("is_payment_recieved", "Y")
		UpdateOrderDetails(cartOrder)
	}
	//database.Find(&products)
	defer database.Close()
	return order.OrderId
}

func UpdateOrderDetails(cartOrder models.CartOrder) {

	var orderDetails models.DbOrderDetails
	//var tempOrderDetails models.DbOrderDetails
	database := db.OpenDbConnection().Table("ordersdetails")
	database.Last(&orderDetails)
	lastorderdetailsID := orderDetails.Id
	fmt.Println("CartOrder", cartOrder)
	cartItems := cartOrder.CartItems

	if err := database.Where("orders_id = ?", cartOrder.OrderId).First(&orderDetails).Error; err != nil {
		// error handling...
		if gorm.IsRecordNotFoundError(err) {
			for _, cartItem := range cartItems {
				lastorderdetailsID++
				tempOrderDetails := models.DbOrderDetails{Id: lastorderdetailsID, OrderID: cartOrder.OrderId, Quantity: cartItem.Quantity, ItemID: cartItem.Id}
				database.Create(&tempOrderDetails)
			}
		} else {
			database.Where("orders_id = ?", cartOrder.OrderId).Delete(&models.CartOrder{})
			database.Last(&orderDetails)
			lastorderdetailsID := orderDetails.Id
			for _, cartItem := range cartItems {
				lastorderdetailsID++
				orderDetails := models.DbOrderDetails{Quantity: cartItem.Quantity, ItemID: cartItem.Id}
				database.Create(&orderDetails)
			}
		}
	}
	//database.Find(&products)
	defer database.Close()

}

func CompleteOrder(checkoutorder models.CheckoutOrder) (bool, string) {

	var order models.DbOrder
	database := db.OpenDbConnection().Table("orders")
	fmt.Println("checkoutorder", checkoutorder)

	if err := database.Where("id = ?", checkoutorder.OrderID).First(&order).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("Invalid Order ID Recieved")
			return false, "Invalid Order ID Recieved"
		}
	} else {
		upderr := database.Where("id = ? AND is_payment_recieved = ?", checkoutorder.OrderID, "N").First(&order).Error
		if upderr != nil {
			fmt.Println("Order Already Completed")
			return false, "Order Already Completed"
		} else {
			return false, "Order couldn't be Completed"
		}
	}
	//database.Find(&products)
	defer database.Close()
	return false, "Order couldn't be Completed"
}
