package repository

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Type     string
	Username string
	Password string
	Protocol string
	Adress   string
	Port     string
	Database string
}

const (
	usersTable  = "user"
	citiesTable = "city"
)

func NewDBInstance(cfg Config) (*sqlx.DB, error) {
	params := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Protocol, cfg.Adress, cfg.Port, cfg.Database)
	db, err := sqlx.Open(cfg.Type, params)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 4)
	return db, nil
}
