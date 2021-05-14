package service

import (
	test_api "go-test-modern-api/pkg"
	"go-test-modern-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user test_api.User) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, error)
}

type Cities interface {
	GetAllCities() ([]test_api.City, error)
	GetOneCity(nameEn string) (test_api.City, error)
}

type Service struct {
	Authorization
	Cities
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(&repos.Authorization),
		Cities:        NewCityService(&repos.Cities),
	}
}
