package model

import "time"

type League struct {
	ID        *string   `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      *string   `json:"name" gorm:"not null"`
	Abbr      *string   `json:"abbr" gorm:"not null"`
	Sport     *string   `json:"sport" gorm:"not null"`
	Level     *string   `json:"level" gorm:"not null"`
	Teams     []Team    `json:"teams"`
	Seasons   []Season  `json:"seasons"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime:nano"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime:nano"`
}
