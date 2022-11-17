package database

import "github.com/ssergomol/Balance-Manager/pkg/models"

type AccountRepo struct {
	store *Storage
}

func (r *AccountRepo) CreateAccount(account models.Account) {
	r.store.db.Query(
		"INSERT INTO (sum) VALUES ($1)",
		account.Sum,
	)
}
