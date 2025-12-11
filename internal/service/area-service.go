package service

import (
    "context"
    "fmt"
	
    "github.com/KarlaR3it/PruebaEureka-Backend/internal/models"
    "github.com/KarlaR3it/PruebaEureka-Backend/internal/repository"
)

type AreaService interface {
    CreateArea(ctx context.Context, area *models.Area) error
    GetAllAreas(ctx context.Context) ([]models.Area, error)
    GetAreaCounts(ctx context.Context) ([]models.AreaResponse, error)
}

type areaService struct {
    repo repository.AreaRepository
}

func NewAreaService(repo repository.AreaRepository) AreaService {
    return &areaService{repo: repo}
}

func (s *areaService) CreateArea(ctx context.Context, area *models.Area) error {
    if area.Name == "" {
        return fmt.Errorf("el nombre del área es requerido")
    }
    // Aquí podrías validar unicidad si lo deseas (consultando el repo)
    return s.repo.Create(ctx, area)
}

func (s *areaService) GetAllAreas(ctx context.Context) ([]models.Area, error) {
    return s.repo.GetAll(ctx)
}

func (s *areaService) GetAreaCounts(ctx context.Context) ([]models.AreaResponse, error) {
    return s.repo.GetAreaCounts(ctx)
}
