package handler

import (
	test_api "go-test-modern-api/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllCities(c *gin.Context) {
	cities, err := h.services.GetAllCities()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cities)
}

func (h *Handler) GetCityByName(c *gin.Context) {
	name := c.Param("name")
	city, err := h.services.Cities.GetOneCity(name)
	if city == (test_api.City{}) {
		NewErrorResponse(c, http.StatusNotFound, "Нет такого города")
		return
	}
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, city)
}
