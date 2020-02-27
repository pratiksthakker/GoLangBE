package models

type DbOrder struct {
	OrderId     int     `gorm:"column:id"`
	Discount    float32 `gorm:"column:discount"`
	Status      string  `gorm:"column:status"`
	IsCompleted string  `gorm:"column:is_payment_recieved"`
	UserId      int     `gorm:"column:user_id"`
}
