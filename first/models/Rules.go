package models

type Rules struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CouponCode  string `json:"coupon_code"`
}
