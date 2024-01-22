package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"people_info/model"
	"strconv"
)

func (h *Handler) createPerson(c *gin.Context) {
	person := model.Person{
		Name:       c.PostForm("name"),
		Surname:    c.PostForm("surname"),
		Patronymic: c.PostForm("patronymic"),
	}
	err := h.service.Create(person)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, "")
}

func (h *Handler) changePerson(c *gin.Context) {
	personId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: err.Error()})
		return
	}
	person := model.Person{
		Id:         personId,
		Name:       c.PostForm("name"),
		Surname:    c.PostForm("surname"),
		Patronymic: c.PostForm("patronymic"),
	}
	if err := h.service.Change(person); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, "")
}

func (h *Handler) deletePerson(c *gin.Context) {
	personId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: err.Error()})
		return
	}
	if err := h.service.Delete(personId); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, "")
}

func (h *Handler) getPersonByFilters(c *gin.Context) {
	age, err := strconv.Atoi(c.DefaultQuery("age", "0"))
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: err.Error()})
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: err.Error()})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: err.Error()})
		return
	}
	filters := model.Filters{
		Name:        c.Query("name"),
		Surname:     c.Query("surname"),
		Patronymic:  c.Query("patronymic"),
		Age:         age,
		Gender:      c.Query("gender"),
		Nationality: c.Query("nationality"),
		Page:        page,
		Limit:       limit,
	}
	rows, err := h.service.GetPersonByFilters(filters)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, rows)
}
