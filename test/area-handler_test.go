package test

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "context"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"

    "github.com/KarlaR3it/PruebaEureka-Backend/internal/handler"
    "github.com/KarlaR3it/PruebaEureka-Backend/internal/models"
)

type mockAreaService struct{ mock.Mock }

func (m *mockAreaService) CreateArea(ctx context.Context, area *models.Area) error {
    args := m.Called(ctx, area)
    return args.Error(0)
}

func TestAreaHandler_Create(t *testing.T) {
    gin.SetMode(gin.TestMode)
    mockSvc := new(mockAreaService)
    h := handler.NewAreaHandler(mockSvc)

    router := gin.Default()
    router.POST("/areas", h.Create)

    body := []byte(`{"name":"Test"}`)
    req, _ := http.NewRequest("POST", "/areas", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()

    mockSvc.On("CreateArea", mock.Anything, mock.AnythingOfType("*models.Area")).Return(nil)

    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusCreated, w.Code)
    mockSvc.AssertExpectations(t)
}

func (m *mockAreaService) GetAllAreas(ctx context.Context) ([]models.Area, error) {
    args := m.Called(ctx)
    return args.Get(0).([]models.Area), args.Error(1)
}

func TestAreaHandler_GetAll(t *testing.T) {
    gin.SetMode(gin.TestMode)
    mockSvc := new(mockAreaService)
    h := handler.NewAreaHandler(mockSvc)

    router := gin.Default()
    router.GET("/areas", h.GetAll)

    req, _ := http.NewRequest("GET", "/areas", nil)
    w := httptest.NewRecorder()

    mockAreas := []models.Area{{Name: "Area1"}, {Name: "Area2"}}
    mockSvc.On("GetAllAreas", mock.Anything).Return(mockAreas, nil)

    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
    mockSvc.AssertExpectations(t)
}

func (m *mockAreaService) GetAreaCounts(ctx context.Context) ([]models.AreaResponse, error) {
    args := m.Called(ctx)
    return args.Get(0).([]models.AreaResponse), args.Error(1)
}

func TestAreaHandler_GetAreaCounts(t *testing.T) {
    gin.SetMode(gin.TestMode)
    mockSvc := new(mockAreaService)
    h := handler.NewAreaHandler(mockSvc)

    router := gin.Default()
    router.GET("/areas/counts", h.GetAreaCounts)

    req, _ := http.NewRequest("GET", "/areas/counts", nil)
    w := httptest.NewRecorder()

    mockCounts := []models.AreaResponse{{Name: "Area1", PersonsQuantity: 5}, {Name: "Area2", PersonsQuantity: 3}}
    mockSvc.On("GetAreaCounts", mock.Anything).Return(mockCounts, nil)

    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
    mockSvc.AssertExpectations(t)
}
