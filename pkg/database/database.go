package database

import (
	"database/sql"
)

type Storage struct {
	DB     *sql.DB
	Config *ConfigDB
	// TODO: add data repos
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
