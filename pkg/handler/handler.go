package handler

import (
	"go-test-modern-api/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	mainRouter := gin.New()
	userRouter := mainRouter.Group("/user")
	{
		userRouter.POST("/singup", h.SignUp)
		userRouter.POST("/singin", h.SignIn)
	}
	cityRouter := mainRouter.Group("/city", h.userIdentity)
	{
		cityRouter.GET("/", h.GetAllCities)
		cityRouter.GET("/:name", h.GetCityByName)
	}
	return mainRouter
}
