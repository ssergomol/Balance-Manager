package models

type Transfer struct {
	FromUserID uint   `json:"from_user_id"`
	FromID     uint   `json:"from_id"`
	ToUserID   uint   `json:"to_user_id"`
	ToID       uint   `json:"to_id"`
	Sum        string `json:"sum"`
}
