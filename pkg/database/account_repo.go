package database

import (
	"database/sql"

	"github.com/shopspring/decimal"
	"github.com/ssergomol/Balance-Manager/pkg/models"
)

type AccountRepo struct {
	store *Storage
}

func (r *AccountRepo) CreateAccount(account models.Account) {
	r.store.db.Query(
		"INSERT INTO (sum) VALUES ($1)",
		account.Sum,
	)
}

func (r *AccountRepo) ReserveFunds(account models.Account) error {

	switch account.ServiceID {

	// Debit funds from the balance
	case 1:
		var oldBalanceSum string
		err := r.store.db.QueryRow("SELECT sum from balances WHERE user_id = $1", account.UserID).Scan(&oldBalanceSum)
		if err != nil {
			return err
		}

		balanceSum, err := decimal.NewFromString(oldBalanceSum)
		if err != nil {
			return err
		}

		deltaSum, err := decimal.NewFromString(account.Sum)
		if err != nil {
			return err
		}

		r.store.db.QueryRow("UPDATE balances SET sum = $1 WHERE user_id = $2", balanceSum.Sub(deltaSum).String(), account.UserID)

		var oldAccountSum string
		if err := r.store.db.QueryRow("SELECT sum from accounts WHERE id = $1", account.ID).Scan(&oldAccountSum); err != nil {
			if err == sql.ErrNoRows {
				r.store.db.QueryRow(
					`INSERT INTO accounts (id, user_id, sum) VALUES ($1, $2, $3)`,
					account.ID, account.UserID, account.Sum,
				)

				return nil
			}
			return err
		}

		accountSum, err := decimal.NewFromString(oldAccountSum)
		if err != nil {
			return err
		}
		r.store.db.QueryRow("UPDATE accounts SET sum = $1 WHERE id = $2", accountSum.Add(deltaSum).String(), account.ID)

	// Debit funds from the account
	case 2:
		var oldAccountSum string
		if err := r.store.db.QueryRow("SELECT sum from accounts WHERE id = $1", account.ID).Scan(&oldAccountSum); err != nil {
			return err
		}

		accountSum, err := decimal.NewFromString(oldAccountSum)
		if err != nil {
			return err
		}

		deltaSum, err := decimal.NewFromString(account.Sum)
		if err != nil {
			return err
		}

		r.store.db.QueryRow("UPDATE accounts SET sum = $1 WHERE id = $2", accountSum.Sub(deltaSum).String(), account.ID)
	}
	return nil
}

func (r *AccountRepo) GetAccountSum(id uint, userID uint) (models.Account, error) {
	var sum string
	if err := r.store.db.QueryRow("SELECT sum from accounts WHERE id = $1", id).Scan(&sum); err != nil {
		if err == sql.ErrNoRows {
			r.store.db.QueryRow(
				`INSERT INTO accounts (id, user_id, sum) VALUES ($1, $2, $3)`,
				id, userID, "0.00",
			)
			account := models.Account{
				ID:     id,
				UserID: userID,
				Sum:    "0.00",
			}
			return account, nil
		}
		return models.Account{}, err
	}

	account := models.Account{
		ID:     id,
		UserID: userID,
		Sum:    sum,
	}
	return account, nil
}
