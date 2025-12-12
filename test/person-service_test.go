package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/KarlaR3it/PruebaEureka-Backend/internal/models"
	"github.com/KarlaR3it/PruebaEureka-Backend/internal/service"
)

type mockPersonRepo struct{ mock.Mock }

func (m *mockPersonRepo) Create(ctx context.Context, person *models.Person) error {
	args := m.Called(ctx, person)
	return args.Error(0)
}

func (m *mockPersonRepo) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	args := m.Called(ctx, email)
	return args.Bool(0), args.Error(1)
}

func (m *mockPersonRepo) GetAll(ctx context.Context) ([]models.Person, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Person), args.Error(1)
}

func TestPersonService_CreatePerson(t *testing.T) {
	mockRepo := new(mockPersonRepo)
	svc := service.NewPersonService(mockRepo)

	person := &models.Person{
		Name:   "Juan Perez",
		Email:  "juan.perez@email.com",
		AreaID: 1,
	}
	mockRepo.On("ExistsByEmail", mock.Anything, "juan.perez@email.com").Return(false, nil)
	mockRepo.On("Create", mock.Anything, person).Return(nil)

	err := svc.CreatePerson(context.Background(), person)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPersonService_GetAllPersons(t *testing.T) {
	mockRepo := new(mockPersonRepo)
	svc := service.NewPersonService(mockRepo)

	mockPersons := []models.Person{
		{Name: "Juan Perez", Email: "juan.perez@email.com", AreaID: 1},
		{Name: "Maria Lopez", Email: "maria.lopez@email.com", AreaID: 2},
	}
	mockRepo.On("GetAll", mock.Anything).Return(mockPersons, nil)

	persons, err := svc.GetAllPersons(context.Background())
	assert.NoError(t, err)
	assert.Len(t, persons, 2)
	mockRepo.AssertExpectations(t)
}
