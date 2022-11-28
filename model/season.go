package model

import "time"

type Season struct {
	ID        *string   `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	StartYear int       `json:"startYear" gorm:"not null"`
	EndYear   int       `json:"endYear" gorm:"not null"`
	LeagueID  *string   `json:"leagueId" gorm:"type:uuid"`
	Matches   []Match   `json:"matches"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime:nano"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime:nano"`
}
