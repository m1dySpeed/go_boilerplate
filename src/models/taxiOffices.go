package models

import (
	"github.com/google/uuid"
	"github.com/m1dySpeed/province_taxi_service/src/database"
	"gorm.io/gorm"
	"strings"
	"time"
)

const TaxiOfficesTableName = "taxi_offices"

type TaxiOffice struct {
	Id           string    `gorm:"primaryKey" json:"id,omitempty"`
	CityId       string    `json:"cityId,omitempty"`
	Name         string    `json:"name,omitempty"`
	Latitude     float64   `json:"latitude,omitempty"`
	Longitude    float64   `json:"longitude,omitempty"`
	WorkingHours string    `json:"workingHours,omitempty"`
	OfficerName  string    `json:"officerName,omitempty"`
	Phone        string    `json:"phone,omitempty"`
	CreatedAt    time.Time `gorm:"autoCreateTime:true" json:"createdAt,omitempty"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime:true" json:"updatedAt,omitempty"`
}

type TaxiOfficeModel struct{}

var archive = new(TaxiOfficesArchiveModel)

func (to *TaxiOfficeModel) Create(o *TaxiOffice) (office *TaxiOffice, err error) {
	var city *City

	o.Id = uuid.New().String()

	db := database.GetDb()

	if err := db.Where("id = ?", o.CityId).Take(&city).Error; err != nil {
		return nil, err
	}

	if err := db.Create(&o).Error; err != nil {
		return nil, err
	}

	return o, nil
}

func (to *TaxiOfficeModel) Update(id string, o map[string]any) (office *TaxiOffice, err error) {
	var updated *TaxiOffice

	db := database.GetDb()

	o["updated_at"] = time.Now()

	if err := db.Table(TaxiOfficesTableName).Where("id = ?", id).Updates(o).Error; err != nil {
		return nil, err
	}

	if err := db.Where("id = ?", id).Take(&updated).Error; err != nil {
		return nil, err
	}

	return updated, nil
}

func (to *TaxiOfficeModel) Get(id string) (office *TaxiOffice, err error) {
	var o *TaxiOffice
	db := database.GetDb()

	if err := db.Where("id = ?", id).Take(&o).Error; err != nil {
		return nil, err
	}

	return o, nil
}

func (to *TaxiOfficeModel) GetList() (ol []*TaxiOffice, err error) {
	var officeList []*TaxiOffice
	db := database.GetDb()

	if err := db.Find(&officeList).Error; err != nil {
		return nil, err
	}

	return officeList, nil
}

func (to *TaxiOfficeModel) Delete(id string) (msg map[string]string, err error) {
	var office *TaxiOffice

	db := database.GetDb()

	if err := db.Where("id = ?", id).Take(&office).Error; err != nil {
		return nil, err
	}

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := archive.Create(office, tx); err != nil {
			return err
		}

		if err := tx.Where("id = ?", id).Delete(&office).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return map[string]string{"msg": "OK"}, nil
}

func (to *TaxiOfficeModel) Recover(id string) (o *TaxiOffice, err error) {
	old, err := archive.Get(id)

	if err != nil {
		return nil, err
	}

	db := database.GetDb()

	rec := TaxiOffice{
		Id:           old.Id,
		CityId:       old.CityId,
		Name:         old.Name,
		Latitude:     old.Latitude,
		Longitude:    old.Longitude,
		WorkingHours: strings.Trim(old.WorkingHours, " "),
		OfficerName:  strings.Trim(old.OfficerName, " "),
		Phone:        old.Phone,
		CreatedAt:    old.CreatedAt,
		UpdatedAt:    time.Now(),
	}

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&rec).Error; err != nil {
			return err
		}

		if err := archive.Delete(old, tx); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &rec, nil
}
