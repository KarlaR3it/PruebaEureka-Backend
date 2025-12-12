package test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/KarlaR3it/PruebaEureka-Backend/internal/handler"
	"github.com/KarlaR3it/PruebaEureka-Backend/internal/models"
)

type mockPersonService struct{ mock.Mock }

func (m *mockPersonService) CreatePerson(ctx context.Context, person *models.Person) error {
	args := m.Called(ctx, person)
	return args.Error(0)
}

func (m *mockPersonService) GetAllPersons(ctx context.Context) ([]models.Person, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Person), args.Error(1)
}

func TestPersonHandler_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockPersonService)
	h := handler.NewPersonHandler(mockSvc)

	router := gin.Default()
	router.POST("/persons", h.Create)

	body := []byte(`{"name":"Juan Perez","email":"juan.perez@email.com","area_id":1}`)
	req, _ := http.NewRequest("POST", "/persons", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mockSvc.On("CreatePerson", mock.Anything, mock.AnythingOfType("*models.Person")).Return(nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestPersonHandler_GetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockPersonService)
	h := handler.NewPersonHandler(mockSvc)

	router := gin.Default()
	router.GET("/persons", h.GetAll)

	req, _ := http.NewRequest("GET", "/persons", nil)
	w := httptest.NewRecorder()

	mockPersons := []models.Person{
		{Name: "Juan Perez", Email: "juan.perez@email.com", AreaID: 1},
		{Name: "Maria Lopez", Email: "maria.lopez@email.com", AreaID: 2},
	}
	mockSvc.On("GetAllPersons", mock.Anything).Return(mockPersons, nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}
