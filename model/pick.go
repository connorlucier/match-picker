package model

import "time"

// TODO add unique constraint on (userId, groupId, matchId)
type Pick struct {
	ID            *string   `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID        *string   `json:"userId" gorm:"type:uuid;not null"`
	GroupID       *string   `json:"groupId" gorm:"type:uuid;not null"`
	MatchID       *string   `json:"matchId" gorm:"type:uuid;not null"`
	WinningTeamID *string   `json:"winningTeamId" gorm:"type:uuid"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime:nano"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"autoUpdateTime:nano"`
}
