package service

import (
	"context"
	"fmt"

	"github.com/KarlaR3it/PruebaEureka-Backend/internal/models"
	"github.com/KarlaR3it/PruebaEureka-Backend/internal/repository"
)

type PersonService interface {
	CreatePerson(ctx context.Context, person *models.Person) error
	GetAllPersons(ctx context.Context) ([]models.Person, error)
}

type personService struct {
	repo repository.PersonRepository
}

func NewPersonService(repo repository.PersonRepository) PersonService {
	return &personService{repo: repo}
}

func (s *personService) CreatePerson(ctx context.Context, person *models.Person) error {
	if person.Name == "" || person.Email == "" || person.AreaID == 0 {
		return fmt.Errorf("nombre, email y área son obligatorios")
	}

	exists, err := s.repo.ExistsByEmail(ctx, person.Email)
	if err != nil {
		return fmt.Errorf("error al validar email: %v", err)
	}
	if exists {
		return fmt.Errorf("el email %s ya está registrado", person.Email)
	}

	return s.repo.Create(ctx, person)
}

func (s *personService) GetAllPersons(ctx context.Context) ([]models.Person, error) {
	return s.repo.GetAll(ctx)
}
