package models

type User struct {
	ID        uint `json:"id"`
	BalanceID uint `json:"balance_id"`
	AccountID uint `json:"account_id"`
	OrderID   uint `json:"order_id"`
}
