package models

import (
	"github.com/google/uuid"
)

type Car struct {
	Id               string    `gorm:"primaryKey" json:"id,omitempty"`
	OfficeID         uuid.UUID `gorm:"type:uuid;not null" json:"officeId"`
	GovNumber        string    `gorm:"type:varchar(255);not null" json:"govNumber"`
	DriverSurname    string    `gorm:"type:varchar(255);not null" json:"driverSurname"`
	DriverName       string    `gorm:"type:varchar(255);not null" json:"driverName"`
	DriverLicense    string    `gorm:"type:varchar(255);not null" json:"driverLicense"`
	AddDriverName    string    `gorm:"type:varchar(255);not null" json:"addDriverName"`
	AddDriverSurname string    `gorm:"type:varchar(255);not null" json:"addDriverSurname"`
	AddDriverLicense string    `gorm:"type:varchar(255);not null" json:"addDriverLicense"`
	Notice           string    `gorm:"type:varchar(255);not null" json:"notice"`
}
