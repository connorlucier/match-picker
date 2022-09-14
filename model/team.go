package model

import (
	"time"
)

type Team struct {
	ID         *string   `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name       *string   `json:"name" gorm:"not null"`
	Abbr       *string   `json:"abbr" gorm:"not null"`
	Conference *string   `json:"conference"`
	Division   *string   `json:"division"`
	LeagueID   *string   `json:"leagueId" gorm:"type:uuid;foreignKey"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime:nano"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime:nano"`
}
