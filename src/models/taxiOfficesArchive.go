package models

import (
	"github.com/m1dySpeed/province_taxi_service/src/database"
	"gorm.io/gorm"
	"strings"
	"time"
)

type TaxiOfficesArchive struct {
	Id           string    `gorm:"primaryKey" json:"id"`
	CityId       string    `json:"cityId"`
	Name         string    `json:"name"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	WorkingHours string    `json:"workingHours"`
	OfficerName  string    `json:"officerName"`
	Phone        string    `json:"phone"`
	CreatedAt    time.Time `gorm:"autoCreateTime:true" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime:true" json:"updated_at"`
	DeletedAt    time.Time `gorm:"autoUpdateTime:true" json:"deleted_at"`
}

type TaxiOfficesArchiveModel struct{}

func (toa *TaxiOfficesArchiveModel) Create(office *TaxiOffice, tx *gorm.DB) (err error) {
	r := TaxiOfficesArchive{
		Id:           office.Id,
		CityId:       office.CityId,
		Name:         office.Name,
		Latitude:     office.Latitude,
		Longitude:    office.Longitude,
		WorkingHours: strings.Trim(office.WorkingHours, " "),
		OfficerName:  strings.Trim(office.OfficerName, " "),
		Phone:        office.Phone,
		CreatedAt:    office.CreatedAt,
		UpdatedAt:    office.UpdatedAt,
		DeletedAt:    time.Now(),
	}

	if err := tx.Create(&r).Error; err != nil {
		return err
	}

	return nil
}

func (toa *TaxiOfficesArchiveModel) Get(id string) (oldOffice *TaxiOfficesArchive, err error) {
	db := database.GetDb()
	var office *TaxiOfficesArchive

	if err := db.Where("id = ?", id).Take(&office).Error; err != nil {
		return nil, err
	}

	return office, nil
}

func (toa *TaxiOfficesArchiveModel) Delete(old *TaxiOfficesArchive, tx *gorm.DB) error {
	if err := tx.Where("id = ?", old.Id).Delete(&old).Error; err != nil {
		return err
	}

	return nil
}
