package models

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Imageurl    uint8  `json:"imageurl"`
	Quantity    int    `json:"quantity"`
}
