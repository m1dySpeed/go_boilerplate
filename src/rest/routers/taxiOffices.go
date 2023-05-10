package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/m1dySpeed/province_taxi_service/src/errobjects"
	"github.com/m1dySpeed/province_taxi_service/src/models"
	"net/http"
)

type TaxiOfficeController struct{}

var toModel = new(models.TaxiOfficeModel)

func (to *TaxiOfficeController) Create(ctx *gin.Context) {
	var office *models.TaxiOffice

	if err := ctx.ShouldBindJSON(&office); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errobjects.TOCreateFail)
		return
	}

	o, err := toModel.Create(office)

	if err != nil {
		switch err.Error() {
		case "record not found":
			ctx.AbortWithStatusJSON(http.StatusNotFound, errobjects.CityNotFound)
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errobjects.TOCreateFail)
			return
		}
	}

	ctx.JSON(http.StatusOK, o)
	return
}

func (to *TaxiOfficeController) Update(ctx *gin.Context) {
	var office map[string]any
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errobjects.TOWrongIdentifier)
		return
	}

	if err := ctx.ShouldBindJSON(&office); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errobjects.TOUpdateFail)
		return
	}

	o, err := toModel.Update(id, office)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errobjects.TOUpdateFail)
		return
	}

	ctx.JSON(http.StatusOK, o)
	return
}

func (to *TaxiOfficeController) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errobjects.TOWrongIdentifier)
	}

	o, err := toModel.Get(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, errobjects.TONotFound)
	}

	ctx.JSON(http.StatusOK, o)
	return
}

func (to *TaxiOfficeController) GetList(ctx *gin.Context) {
	officeList, err := toModel.GetList()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errobjects.TORequestError)
	}

	ctx.JSON(http.StatusOK, officeList)
	return
}

func (to *TaxiOfficeController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errobjects.TOWrongIdentifier)
		return
	}

	res, err := toModel.Delete(id)

	if err != nil {
		switch err.Error() {
		case "record not found":
			ctx.AbortWithStatusJSON(http.StatusNotFound, errobjects.TONotFound)
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errobjects.TODeleteFail)
			return
		}
	}

	ctx.JSON(http.StatusOK, res)
	return
}

func (to *TaxiOfficeController) Recover(ctx *gin.Context) {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errobjects.TOWrongIdentifier)
		return
	}

	res, err := toModel.Recover(id)

	if err != nil {
		switch err.Error() {
		case "record not found":
			ctx.AbortWithStatusJSON(http.StatusNotFound, errobjects.TONotFound)
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errobjects.TODeleteFail)
			return
		}
	}

	ctx.JSON(http.StatusOK, res)
	return
}
