package model

import "time"

type User struct {
	ID           *string   `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	FirstName    *string   `json:"firstName" gorm:"not null"`
	LastName     *string   `json:"lastName" gorm:"not null"`
	EmailAddress *string   `json:"emailAddress" gorm:"unique; not null"`
	PhoneNumber  *string   `json:"phoneNumber" gorm:"unique; not null"`
	Groups       []Group   `json:"groups"`
	CreatedAt    time.Time `json:"createdAt" gorm:"autoCreateTime:nano"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"autoUpdateTime:nano"`
}
