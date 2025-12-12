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

	exists, err := s.repo.ExistsByName(ctx, area.Name)
	if err != nil {
		return fmt.Errorf("error al validar nombre del área: %v", err)
	}
	if exists {
		return fmt.Errorf("el área %s ya existe", area.Name)
	}

	return s.repo.Create(ctx, area)
}

func (s *areaService) GetAllAreas(ctx context.Context) ([]models.Area, error) {
	return s.repo.GetAll(ctx)
}

func (s *areaService) GetAreaCounts(ctx context.Context) ([]models.AreaResponse, error) {
	return s.repo.GetAreaCounts(ctx)
}
