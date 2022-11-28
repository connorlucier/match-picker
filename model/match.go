package model

import "time"

type Match struct {
	ID         *string   `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	HomeTeamID *string   `json:"homeTeamId" gorm:"type:uuid"`
	AwayTeamID *string   `json:"awayTeamId" gorm:"type:uuid"`
	SeasonID   *string   `json:"seasonId" gorm:"type:uuid"`
	StartsAt   time.Time `json:"startsAt" gorm:"not null"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime:nano"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime:nano"`
}
