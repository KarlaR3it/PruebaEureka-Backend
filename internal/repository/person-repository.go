package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/KarlaR3it/PruebaEureka-Backend/internal/models"
)

type PersonRepository interface {
	Create(ctx context.Context, person *models.Person) error
	GetAll(ctx context.Context) ([]models.Person, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

type personRepository struct {
	db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) PersonRepository {
	return &personRepository{db: db}
}

func (r *personRepository) Create(ctx context.Context, person *models.Person) error {
	return r.db.WithContext(ctx).Create(person).Error
}

func (r *personRepository) GetAll(ctx context.Context) ([]models.Person, error) {
	var persons []models.Person
	err := r.db.WithContext(ctx).Preload("Area").Find(&persons).Error
	return persons, err
}

func (r *personRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Person{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}
