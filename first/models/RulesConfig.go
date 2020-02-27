package models

type RulesConfig struct {
	ID            int     `json:"id"`
	Qty           int     `json:"qty"`
	DiscountValue float32 `json:"discount_value"`
	DiscountType  string  `json:"discount_type"`
	RulesId       int     `json:"rules_id"`
	ItemID        int     `json:"item_id"`
}
