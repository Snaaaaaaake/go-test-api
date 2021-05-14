package service

import (
	test_api "go-test-modern-api/pkg"
	"go-test-modern-api/pkg/repository"
)

type CityService struct {
	repo repository.Cities
}

func NewCityService(repo *repository.Cities) *CityService {
	return &CityService{*repo}
}

func (s *CityService) GetAllCities() ([]test_api.City, error) {
	cities, err := s.repo.GetAllCities()
	return cities, err
}

func (s *CityService) GetOneCity(nameEn string) (test_api.City, error) {
	city, err := s.repo.GetOneCity(nameEn)
	return city, err
}
