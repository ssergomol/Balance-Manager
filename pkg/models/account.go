package models

type Account struct {
	ID  uint        `json:"id"`
	Sum interface{} `json:"sum"`
}
