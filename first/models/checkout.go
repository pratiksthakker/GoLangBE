package models

type CheckoutOrder struct {
	UserName string `json:"user_name"`
	OrderID  int    `json:"order_id"`
	Status   string `json:"order_status"`
}
