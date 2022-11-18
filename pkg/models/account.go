package models

type Account struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	ServiceID uint   `json:"service_id"`
	Sum       string `json:"sum"`
}
