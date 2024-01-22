package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oatsmoke/people_info/internal/app/model"
)

type PersonRepository struct {
	connectionDB *pgxpool.Pool
}

func NewPersonRepository(connectionDB *pgxpool.Pool) *PersonRepository {
	return &PersonRepository{
		connectionDB: connectionDB,
	}
}

func (r *PersonRepository) Create(person model.Person) error {
	query := `INSERT INTO people (name, surname, patronymic, age, gender, nationality)
				VALUES ($1, $2, $3, $4, $5, $6);`
	_, err := r.connectionDB.Exec(context.Background(), query,
		person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality)
	if err != nil {
		return fmt.Errorf("PersonRepository Create: %s", err.Error())
	}
	return nil
}

func (r *PersonRepository) Change(person model.Person) error {
	query := `UPDATE people 
				SET name=$2, surname=$3, patronymic=$4, age=$5, gender=$6, nationality=$7
				WHERE id = $1;`
	_, err := r.connectionDB.Exec(context.Background(), query,
		person.Id, person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality)
	if err != nil {
		return fmt.Errorf("PersonRepository Change: %s", err.Error())
	}
	return nil
}

func (r *PersonRepository) Delete(personId int) error {
	query := `DELETE FROM people 
       			WHERE id = $1;`
	_, err := r.connectionDB.Exec(context.Background(), query, personId)
	if err != nil {
		return fmt.Errorf("PersonRepository Delete: %s", err.Error())
	}
	return nil
}

func (r *PersonRepository) GetPersonByFilters(str string, limit, page int) ([]model.Person, error) {
	var people []model.Person
	var person model.Person
	query := "SELECT * FROM people " + str + "LIMIT $1 OFFSET $2;"
	rows, err := r.connectionDB.Query(context.Background(), query, limit, (page-1)*limit)
	if err != nil {
		return nil, fmt.Errorf("PersonRepository GetPersonByFilters: %s", err.Error())
	}
	for rows.Next() {
		err = rows.Scan(
			&person.Id,
			&person.Name,
			&person.Surname,
			&person.Patronymic,
			&person.Age,
			&person.Gender,
			&person.Nationality)
		if err != nil {
			return nil, fmt.Errorf("PersonRepository GetPersonByFilters: %s", err.Error())
		}
		people = append(people, person)
	}
	return people, nil
}
