package models

import (
	"github.com/google/uuid"
	"github.com/m1dySpeed/province_taxi_service/src/database"
)

type City struct {
	Id        string  `gorm:"primaryKey" json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CityModel struct{}

func (c *CityModel) Create(city *City) (r *City, err error) {
	city.Id = uuid.New().String()

	if err := database.GetDb().Create(&city).Error; err != nil {
		return nil, err
	}

	return city, nil
}

func (c *CityModel) Delete(id string) (msg map[string]string, err error) {
	if err := database.GetDb().Where("id = ?", id).Delete(&City{}).Error; err != nil {
		return nil, err
	} else {
		result := map[string]string{"msg": "OK"}
		return result, nil
	}
}

func (c *CityModel) Get(id string) (city *City, err error) {
	var res *City

	if err := database.GetDb().Where("id = ?", id).Take(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
