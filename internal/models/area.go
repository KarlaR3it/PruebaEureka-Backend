package models

import (
	"time"
)

type Area struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null;unique" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Persons   []Person  `gorm:"foreignKey:AreaID" json:"persons,omitempty"`
}

type AreaResponse struct {
	Name            string `json:"name"`
	PersonsQuantity int    `json:"persons_quantity"`
}

func (Area) TableName() string {
	return "areas"
}
