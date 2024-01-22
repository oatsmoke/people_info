package handler

import (
	"github.com/gin-gonic/gin"
	"people_info/service"
)

type errorResponse struct {
	Message string `json:"message"`
}

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	person := router.Group("/person")
	{
		person.POST("", h.createPerson)
		person.PUT("/:id", h.changePerson)
		person.DELETE("/:id", h.deletePerson)
		person.GET("", h.getPersonByFilters)
	}

	return router
}
