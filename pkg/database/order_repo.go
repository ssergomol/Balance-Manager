package database

import "github.com/ssergomol/Balance-Manager/pkg/models"

type OrderRepo struct {
	store *Storage
}

func (r *OrderRepo) CreateOrder(order models.Order) {
	r.store.DB.Query("INSERT INTO (service_id, price, description) VALUES ($1, $2, $3);",
		order.ServiceID, order.Price, order.Description,
	)
}
