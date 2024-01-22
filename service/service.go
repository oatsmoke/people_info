package service

import (
	"people_info/model"
	"people_info/repository"
)

type Service struct {
	Person
}

type Person interface {
	Create(person model.Person) error
	Change(person model.Person) error
	Delete(personId int) error
	GetPersonByFilters(filters model.Filters) ([]model.Person, error)
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Person: NewPersonService(repository.People),
	}
}
