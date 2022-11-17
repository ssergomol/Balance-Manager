package database

import "github.com/ssergomol/Balance-Manager/pkg/models"

type BalanceRepo struct {
	store *Storage
}

func (r *BalanceRepo) CreateBalance(balance models.Balance) {
	r.store.db.QueryRow(
		`INSERT INTO balances (sum) VALUES ($1)`,
		balance.Sum,
	)
}
