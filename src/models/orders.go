package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	ID                    uuid.UUID `gorm:"type:uuid;not null"`
	CarID                 uuid.UUID `gorm:"type:uuid;not null"`
	PickupAddressStr      string    `gorm:"type:varchar(255);not null"`
	DestinationAddressStr string    `gorm:"type:varchar(255);not null"`
	PickupLatitude        float64   `gorm:"type:double precision;not null"`
	PickupLongitude       float64   `gorm:"type:double precision;not null"`
	DestinationLatitude   float64   `gorm:"type:double precision;not null"`
	DestinationLongitude  float64   `gorm:"type:double precision;not null"`
	CreatedAt             time.Time `gorm:"type:timestamp(0) without time zone;not null"`
	UpdatedAt             time.Time `gorm:"type:timestamp(0) without time zone;not null"`
}
