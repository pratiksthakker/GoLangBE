package models

type DbOrderDetails struct {
	Id       int `gorm:"column:id"`
	Quantity int `gorm:"column:qty"`
	OrderID  int `gorm:"column:orders_id"`
	ItemID   int `gorm:"column:item_id"`
}
