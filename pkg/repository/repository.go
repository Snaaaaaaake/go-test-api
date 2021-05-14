package repository

import (
	test_api "go-test-modern-api/pkg"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user test_api.User) (string, error)
	GetUser(username string, password string) (test_api.User, error)
}
type Cities interface {
	GetAllCities() ([]test_api.City, error)
	GetOneCity(nameEn string) (test_api.City, error)
}

type Repository struct {
	Authorization
	Cities
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMysql(db),
		Cities:        NewCityMysql(db),
	}
}
