package database

import "github.com/ssergomol/Balance-Manager/pkg/models"

type UserRepo struct {
	store *Storage
}

func (r *UserRepo) CreateUser(user models.User) {
	r.store.DB.Query("INSERT INTO (balance_id, account_id, order_id) VALUES ($1, $2, $3)",
		user.BalanceID, user.AccountID, user.OrderID,
	)
}
