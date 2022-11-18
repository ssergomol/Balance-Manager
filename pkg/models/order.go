package models

type Order struct {
	UserID      uint   `json:"user_id"`
	ServiceID   uint   `json:"service_id"`
	Price       string `json:"price"`
	IsPositive  bool   `json:"is_positive"`
	Description string `json:"description"`
}
