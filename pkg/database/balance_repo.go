package database

import (
	"database/sql"

	"github.com/shopspring/decimal"
	"github.com/ssergomol/Balance-Manager/pkg/models"
)

type BalanceRepo struct {
	store *Storage
}

func (r *BalanceRepo) CreateBalance(balance models.Balance) {
	r.store.db.QueryRow(
		`INSERT INTO balances (sum) VALUES ($1)`,
		balance.Sum,
	)
}

func (r *BalanceRepo) ReplenishBalance(balance models.Balance) error {
	var oldSum string
	if err := r.store.db.QueryRow("SELECT sum from balances WHERE user_id = $1", balance.ID).Scan(&oldSum); err != nil {
		if err == sql.ErrNoRows {
			r.store.db.QueryRow(
				`INSERT INTO balances (user_id, sum) VALUES ($1, $2)`,
				balance.ID, balance.Sum,
			)

			orderFrom := models.Order{
				UserID:      balance.ID,
				ServiceID:   4,
				IsPositive:  true,
				Price:       balance.Sum,
				Description: "Replenish balance",
			}

			err = r.store.Order().CreateOrder(orderFrom)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	err := r.store.db.QueryRow("SELECT sum from balances WHERE user_id = $1", balance.ID).Scan(&oldSum)
	if err != nil {
		return err
	}

	sum, err := decimal.NewFromString(oldSum)
	if err != nil {
		return err
	}

	deltaSum, err := decimal.NewFromString(balance.Sum)
	if err != nil {
		return err
	}
	newSum := sum.Add(deltaSum)

	r.store.db.QueryRow("UPDATE balances SET sum = $1 WHERE user_id = $2", newSum.String(), balance.ID)

	orderFrom := models.Order{
		UserID:      balance.ID,
		ServiceID:   4,
		IsPositive:  true,
		Price:       balance.Sum,
		Description: "Replenish balance",
	}

	err = r.store.Order().CreateOrder(orderFrom)
	if err != nil {
		return err
	}
	return nil
}

func (r *BalanceRepo) GetBalance(userID uint) (models.Balance, error) {
	var sum string
	if err := r.store.db.QueryRow("SELECT sum from balances WHERE user_id = $1", userID).Scan(&sum); err != nil {
		if err == sql.ErrNoRows {
			r.store.db.QueryRow(
				`INSERT INTO balances (user_id, sum) VALUES ($1, $2)`,
				userID, "0.00",
			)
			balance := models.Balance{
				ID:  userID,
				Sum: "0.00",
			}
			return balance, nil
		}
		return models.Balance{}, err
	}

	balance := models.Balance{
		ID:  userID,
		Sum: sum,
	}
	return balance, nil
}
