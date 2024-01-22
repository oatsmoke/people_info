package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oatsmoke/people_info/internal/app/model"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) createPerson(c *gin.Context) {
	name := c.PostForm("name")
	if err := validEmptyOrLen(name); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: "name " + err.Error()})
		return
	}
	surname := c.PostForm("surname")
	if err := validEmptyOrLen(surname); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: "surname " + err.Error()})
		return
	}
	patronymic := c.PostForm("patronymic")
	if err := validLen(patronymic); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: "patronymic " + err.Error()})
		return
	}
	person := model.Person{
		Name:       name,
		Surname:    surname,
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
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: "id " + err.Error()})
		return
	}
	name := c.PostForm("name")
	if err := validEmptyOrLen(name); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: "name " + err.Error()})
		return
	}
	surname := c.PostForm("surname")
	if err := validEmptyOrLen(surname); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: "surname " + err.Error()})
		return
	}
	patronymic := c.PostForm("patronymic")
	if err := validLen(patronymic); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: "patronymic " + err.Error()})
		return
	}
	age, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: "age " + err.Error()})
		return
	}
	if err := validInt(age); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: "age " + err.Error()})
		return
	}
	gender := c.PostForm("gender")
	if err := validEmptyOrLen(gender); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: "gender " + err.Error()})
		return
	}
	nationality := c.PostForm("nationality")
	if err := validEmptyOrLen(nationality); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: "nationality " + err.Error()})
		return
	}
	person := model.Person{
		Id:          personId,
		Name:        name,
		Surname:     surname,
		Patronymic:  patronymic,
		Age:         age,
		Gender:      gender,
		Nationality: nationality,
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
	name := c.Query("name")
	if err := validLen(name); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: "name " + err.Error()})
		return
	}
	surname := c.Query("surname")
	if err := validLen(surname); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: "surname " + err.Error()})
		return
	}
	patronymic := c.Query("patronymic")
	if err := validLen(patronymic); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: "patronymic " + err.Error()})
		return
	}
	age, err := strconv.Atoi(c.DefaultQuery("age", "0"))
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: err.Error()})
		return
	}
	gender := c.Query("gender")
	if err := validLen(gender); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: "gender " + err.Error()})
		return
	}
	nationality := c.Query("nationality")
	if err := validLen(nationality); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse{Message: "nationality " + err.Error()})
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
		Name:        name,
		Surname:     surname,
		Patronymic:  patronymic,
		Age:         age,
		Gender:      gender,
		Nationality: nationality,
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

func validEmptyOrLen(data string) error {
	if data == "" || len(data) > 20 {
		return fmt.Errorf("incorrect data")
	}
	return nil
}

func validLen(data string) error {
	if len(data) > 20 {
		return fmt.Errorf("incorrect data")
	}
	return nil
}

func validInt(data int) error {
	if 0 > data && data > 150 {
		return fmt.Errorf("incorrect data")
	}
	return nil
}
