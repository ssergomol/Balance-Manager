package database

type ConfigDB struct {
	databaseURL string
}

func NewConfig() *ConfigDB {
	return &ConfigDB{
		databaseURL: "host=localhost dbname=balance_manager sslmode=disable",
	}
}
