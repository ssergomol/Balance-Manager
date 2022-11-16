package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	DB          *sql.DB
	Config      *ConfigDB
	userRepo    *UserRepo
	orderRepo   *OrderRepo
	balanceRepo *BalanceRepo
	accountRepo *AccountRepo
}

func NewDB(config *ConfigDB) *Storage {
	return &Storage{
		Config: config,
	}
}

func (s *Storage) Connect() error {
	db, err := sql.Open("postgres", s.Config.databaseURL)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	s.DB = db
	return nil
}

func (s *Storage) Disconnect() error {
	return s.DB.Close()
}
