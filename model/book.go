package model

import (
	"time"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	MediumPrice float32        `json:"medium_price"`
	Author      string         `json:"author"`
	ImageURL    string         `json:"img_url"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"`
}
