package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	db          *sql.DB
	config      *ConfigDB
	userRepo    *UserRepo
	orderRepo   *OrderRepo
	balanceRepo *BalanceRepo
	accountRepo *AccountRepo
}

func NewDB(config *ConfigDB) *Storage {
	return &Storage{
		config: config,
	}
}

func (s *Storage) Connect() error {
	db, err := sql.Open("postgres", s.config.dbServer)
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE DATABASE " + s.config.dbName)
	if err != nil {
		return err
	}

	db.Close()

	db, err = sql.Open("postgres", s.config.dbURL)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	s.db = db

	return LoadSQLFile(s.db)
}

func (s *Storage) Disconnect() error {
	return s.db.Close()
}

func (s *Storage) Order() *OrderRepo {
	if s.orderRepo != nil {
		return s.orderRepo
	}

	s.orderRepo = &OrderRepo{
		store: s,
	}
	return s.orderRepo
}

func (s *Storage) User() *UserRepo {
	if s.orderRepo != nil {
		return s.userRepo
	}

	s.userRepo = &UserRepo{
		store: s,
	}
	return s.userRepo
}

func (s *Storage) Balance() *BalanceRepo {
	if s.balanceRepo != nil {
		return s.balanceRepo
	}

	s.balanceRepo = &BalanceRepo{
		store: s,
	}
	return s.balanceRepo
}

func (s *Storage) Account() *AccountRepo {
	if s.accountRepo != nil {
		return s.accountRepo
	}

	s.accountRepo = &AccountRepo{
		store: s,
	}
	return s.accountRepo
}
