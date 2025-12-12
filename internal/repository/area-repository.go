package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/KarlaR3it/PruebaEureka-Backend/internal/models"
)

type AreaRepository interface {
	Create(ctx context.Context, area *models.Area) error
	GetAll(ctx context.Context) ([]models.Area, error)
	GetAreaCounts(ctx context.Context) ([]models.AreaResponse, error)
	ExistsByName(ctx context.Context, name string) (bool, error)
}

type areaRepository struct {
	db *gorm.DB
}

func NewAreaRepository(db *gorm.DB) AreaRepository {
	return &areaRepository{db: db}
}

func (r *areaRepository) Create(ctx context.Context, area *models.Area) error {
	return r.db.WithContext(ctx).Create(area).Error
}

func (r *areaRepository) GetAll(ctx context.Context) ([]models.Area, error) {
	var areas []models.Area
	err := r.db.WithContext(ctx).Find(&areas).Error
	return areas, err
}

func (r *areaRepository) GetAreaCounts(ctx context.Context) ([]models.AreaResponse, error) {
	var results []models.AreaResponse
	err := r.db.WithContext(ctx).
		Model(&models.Area{}).
		Select("areas.name, COUNT(persons.id) as persons_quantity").
		Joins("LEFT JOIN persons ON persons.area_id = areas.id").
		Group("areas.id, areas.name").
		Scan(&results).Error
	return results, err
}

func (r *areaRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Area{}).Where("name = ?", name).Count(&count).Error
	return count > 0, err
}
