package routers

import (
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m1dySpeed/province_taxi_service/src/errobjects"
	"github.com/m1dySpeed/province_taxi_service/src/models"
)

const CitiesTableName = "cities"

type CityController struct{}

var cityModel = new(models.CityModel)

func (c *CityController) Create(ctx *gin.Context) {
	var city *models.City

	if checkErr := ctx.ShouldBindJSON(&city); checkErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errobjects.CityCreateFail)
		return
	}

	newCity, err := cityModel.Create(city)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errobjects.CityCreateFail)
	}

	ctx.JSON(http.StatusOK, &newCity)
	return
}

func (c *CityController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errobjects.CityWrongIdentifier)
		return
	}

	res, err := cityModel.Delete(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errobjects.CityDeleteFail)
		return
	}

	ctx.JSON(http.StatusOK, res)
	return
}

func (c *CityController) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errobjects.CityWrongIdentifier)
		return
	}

	city, err := cityModel.Get(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, errobjects.CityNotFound)
		return
	}

	ctx.JSON(http.StatusOK, city)
	return
}
