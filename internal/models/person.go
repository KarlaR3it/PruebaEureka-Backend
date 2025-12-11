package models

import (
	"time"
)

type Person struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"type:varchar(100);not null;unique" json:"email"`
	AreaID    uint      `gorm:"not null" json:"area_id"`
	Area      Area      `gorm:"foreignKey:AreaID" json:"area"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type PersonResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Area  string `json:"area"`
}

func (Person) TableName() string {
	return "persons"
}
