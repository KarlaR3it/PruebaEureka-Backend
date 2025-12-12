package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/KarlaR3it/PruebaEureka-Backend/internal/models"
	"github.com/KarlaR3it/PruebaEureka-Backend/internal/service"
)

type mockAreaRepo struct{ mock.Mock }

func (m *mockAreaRepo) Create(ctx context.Context, area *models.Area) error {
	args := m.Called(ctx, area)
	return args.Error(0)
}

func (m *mockAreaRepo) ExistsByName(ctx context.Context, name string) (bool, error) {
	args := m.Called(ctx, name)
	return args.Bool(0), args.Error(1)
}

func TestAreaService_CreateArea(t *testing.T) {
	mockRepo := new(mockAreaRepo)
	svc := service.NewAreaService(mockRepo)

	area := &models.Area{Name: "Test"}
	mockRepo.On("ExistsByName", mock.Anything, "Test").Return(false, nil)
	mockRepo.On("Create", mock.Anything, area).Return(nil)

	err := svc.CreateArea(context.Background(), area)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func (m *mockAreaRepo) GetAll(ctx context.Context) ([]models.Area, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Area), args.Error(1)
}

func TestAreaService_GetAllAreas(t *testing.T) {
	mockRepo := new(mockAreaRepo)
	svc := service.NewAreaService(mockRepo)

	mockAreas := []models.Area{{Name: "Area1"}, {Name: "Area2"}}
	mockRepo.On("GetAll", mock.Anything).Return(mockAreas, nil)

	areas, err := svc.GetAllAreas(context.Background())
	assert.NoError(t, err)
	assert.Len(t, areas, 2)
	mockRepo.AssertExpectations(t)
}

func (m *mockAreaRepo) GetAreaCounts(ctx context.Context) ([]models.AreaResponse, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.AreaResponse), args.Error(1)
}

func TestAreaService_GetAreaCounts(t *testing.T) {
	mockRepo := new(mockAreaRepo)
	svc := service.NewAreaService(mockRepo)

	mockCounts := []models.AreaResponse{
		{Name: "Recursos Humanos", PersonsQuantity: 5},
		{Name: "Ventas", PersonsQuantity: 3},
	}
	mockRepo.On("GetAreaCounts", mock.Anything).Return(mockCounts, nil)

	counts, err := svc.GetAreaCounts(context.Background())
	assert.NoError(t, err)
	assert.Len(t, counts, 2)
	assert.Equal(t, "Recursos Humanos", counts[0].Name)
	assert.EqualValues(t, 5, counts[0].PersonsQuantity)
	mockRepo.AssertExpectations(t)
}
