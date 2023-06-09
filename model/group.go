package model

import "time"

type Group struct {
	ID        *string   `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      *string   `json:"name" gorm:"not null"`
	UserID    *string   `json:"userID" gorm:"not null"`
	Members   []User    `json:"members" gorm:"many2many:group_members"`
	Leagues   []League  `json:"leagues" gorm:"many2many:group_leagues"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime:nano"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime:nano"`
}
