package repository

import (
	"fmt"
	test_api "go-test-modern-api/pkg"

	"github.com/jmoiron/sqlx"
)

type AuthMysql struct {
	db *sqlx.DB
}

func NewAuthMysql(db *sqlx.DB) *AuthMysql {
	return &AuthMysql{db}
}

func (r *AuthMysql) CreateUser(user test_api.User) (string, error) {
	var response test_api.User
	query := fmt.Sprintf(`INSERT INTO %s (entity_id, name, photo_url, retailer) VALUES("%s", "%s", "%s", "%s")`,
		usersTable, user.Id, user.Name, user.Password, user.Username)
	r.db.QueryRow(query)
	query = fmt.Sprintf(`SELECT entity_id, name, photo_url, retailer FROM %s WHERE entity_id="%s"`, usersTable, user.Id)
	err := r.db.Get(&response, query)
	if err != nil {
		return "", err
	}
	return response.Id, nil
}

func (r *AuthMysql) GetUser(username string, password string) (test_api.User, error) {
	var user test_api.User
	query := fmt.Sprintf(`SELECT name, photo_url, retailer, entity_id FROM %s WHERE retailer="%s" AND photo_url="%s"`, usersTable, username, password)
	err := r.db.Get(&user, query)
	return user, err
}
