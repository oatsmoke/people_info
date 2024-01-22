package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oatsmoke/people_info/internal/app/model"
)

type Repository struct {
	People
}

type People interface {
	Create(person model.Person) error
	Change(person model.Person) error
	Delete(personId int) error
	GetPersonByFilters(srt string, limit, page int) ([]model.Person, error)
}

func NewRepository(connectionDB *pgxpool.Pool) *Repository {
	return &Repository{
		People: NewPersonRepository(connectionDB),
	}
}
