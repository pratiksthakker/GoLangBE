package models

type CartOrder struct {
	UserName string `json:"user_name"`
	OrderId  int    `json:"order_id"`

	ZipCode    string `json:"zip_code"`
	CouponCode string `json:"coupon_code"`
	CartItems  []struct {
		Id       int     `json:"id"`
		Quantity int     `json:"quantity"`
		Price    float32 `json:"price"`
	} `json:"cart_items"`
	TotalPrice          float32 `json:"total_price"`
	DiscountedPrice     float32 `json:"discounted_price"`
	DiscountPercent     float32 `json:"discount_percent"`
	DiscountDescription string  `json:"discount_desc"`
}
