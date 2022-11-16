package models

type Order struct {
	ID          uint        `json:"id"`
	ServiceID   uint        `json:"service_id"`
	Price       interface{} `json:"price"`
	Description string      `json:"description"`
	Date        string      `json:"execution_date"`
}
