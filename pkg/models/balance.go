package models

type Balance struct {
	ID  uint        `json:"id"`
	Sum interface{} `json:"sum"`
}
