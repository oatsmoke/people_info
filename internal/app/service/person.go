package service

import (
	"encoding/json"
	"fmt"
	"github.com/oatsmoke/people_info/internal/app/model"
	"github.com/oatsmoke/people_info/internal/app/repository"
	"io"
	"net/http"
	"sync"
)

type PersonService struct {
	repositoryPeople repository.People
}

type chData struct {
	resource string
	result   interface{}
	err      error
}

func NewPersonService(repositoryPeople repository.People) *PersonService {
	return &PersonService{
		repositoryPeople: repositoryPeople,
	}
}

func (s *PersonService) Create(person model.Person) error {
	ch := make(chan chData)
	wg := sync.WaitGroup{}
	wg.Add(3)
	go age(&wg, ch, person.Name)
	go gender(&wg, ch, person.Name)
	go nationality(&wg, ch, person.Name)
	go func(wg *sync.WaitGroup, ch chan chData) {
		wg.Wait()
		close(ch)
	}(&wg, ch)
	for item := range ch {
		if item.err != nil {
			_ = fmt.Errorf("PersonService Create: %s", item.err.Error())
			continue
		}
		switch item.resource {
		case "agify":
			person.Age = item.result.(int)
		case "genderize":
			person.Gender = item.result.(string)
		case "nationalize":
			person.Nationality = item.result.(string)
		}
	}
	if err := s.repositoryPeople.Create(person); err != nil {
		return fmt.Errorf("PersonService Create: %s", err.Error())
	}
	return nil
}

func (s *PersonService) Change(person model.Person) error {
	return s.repositoryPeople.Change(person)
}

func (s *PersonService) Delete(personId int) error {
	return s.repositoryPeople.Delete(personId)
}

func (s *PersonService) GetPersonByFilters(filters model.Filters) ([]model.Person, error) {
	str := ""
	if filters.Name != "" {
		str = fmt.Sprintf("name = '%s'", filters.Name)
	}
	if filters.Surname != "" {
		if str != "" {
			str = fmt.Sprintf("%s AND", str)
		}
		str = fmt.Sprintf("%s surname = '%s'", str, filters.Surname)
	}
	if filters.Patronymic != "" {
		if str != "" {
			str = fmt.Sprintf("%s AND", str)
		}
		str = fmt.Sprintf("%s patronymic = '%s'", str, filters.Patronymic)
	}
	if filters.Age != 0 {
		if str != "" {
			str = fmt.Sprintf("%s AND", str)
		}
		str = fmt.Sprintf("%s age = %d", str, filters.Age)
	}
	if filters.Gender != "" {
		if str != "" {
			str = fmt.Sprintf("%s AND", str)
		}
		str = fmt.Sprintf("%s gender = '%s'", str, filters.Gender)
	}
	if filters.Nationality != "" {
		if str != "" {
			str = fmt.Sprintf("%s AND", str)
		}
		str = fmt.Sprintf("%s nationality = '%s'", str, filters.Nationality)
	}
	if str != "" {
		str = fmt.Sprintf("WHERE %s ", str)
	}
	return s.repositoryPeople.GetPersonByFilters(str, filters.Limit, filters.Page)
}

func age(wg *sync.WaitGroup, ch chan chData, name string) {
	defer wg.Done()
	age, err := getData("agify", name)
	if err != nil {
		ch <- chData{err: err}
		return
	}
	var ageData model.AgeData
	if err := json.Unmarshal(age, &ageData); err != nil {
		ch <- chData{err: err}
		return
	}
	ch <- chData{resource: "agify", result: ageData.Age, err: nil}
}

func gender(wg *sync.WaitGroup, ch chan chData, name string) {
	defer wg.Done()
	gender, err := getData("genderize", name)
	if err != nil {
		ch <- chData{err: err}
		return
	}
	var genderData model.GenderData
	if err := json.Unmarshal(gender, &genderData); err != nil {
		ch <- chData{err: err}
		return
	}
	ch <- chData{resource: "genderize", result: genderData.Gender, err: nil}
}

func nationality(wg *sync.WaitGroup, ch chan chData, name string) {
	defer wg.Done()
	nationality, err := getData("nationalize", name)
	if err != nil {
		ch <- chData{err: err}
		return
	}
	var nationalityData model.NationalityDate
	if err := json.Unmarshal(nationality, &nationalityData); err != nil {
		ch <- chData{err: err}
		return
	}
	ch <- chData{resource: "nationalize", result: nationalityData.Country[0].CountryId, err: nil}
}

func getData(resource, name string) ([]byte, error) {
	url := fmt.Sprintf("https://api.%s.io/?name=%s", resource, name)
	response, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("PersonService getData(https://api.%s.io/?name=%s): %s",
			resource, name, err.Error())
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("PersonService getData(https://api.%s.io/?name=%s): %s",
			resource, name, err.Error())
	}
	return body, nil
}
