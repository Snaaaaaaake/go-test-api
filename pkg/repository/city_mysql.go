package repository

import (
	"fmt"
	test_api "go-test-modern-api/pkg"

	"github.com/jmoiron/sqlx"
)

type CityMysql struct {
	db *sqlx.DB
}

func NewCityMysql(db *sqlx.DB) *CityMysql {
	return &CityMysql{db}
}

func (r *CityMysql) GetAllCities() ([]test_api.City, error) {
	var response []test_api.City
	query := "SELECT name_en, name_ru FROM city"
	err := r.db.Select(&response, query)
	return response, err
}

func (r *CityMysql) GetOneCity(nameEn string) (test_api.City, error) {
	var response test_api.City
	query := fmt.Sprintf("SELECT name_en, name_ru FROM city WHERE name_en='%s'", nameEn)
	err := r.db.Get(&response, query)
	return response, err
}
