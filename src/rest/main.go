package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/m1dySpeed/province_taxi_service/src/rest/routers"
)

func Init(e *gin.Engine) {
	citiesRg := e.Group("/cities")
	{
		cities := new(routers.CityController)

		citiesRg.POST("/create", cities.Create)
		citiesRg.DELETE(":id", cities.Delete)
		citiesRg.GET(":id", cities.Get)
	}

	taxiOfficesRg := e.Group("taxi-offices")
	{
		to := new(routers.TaxiOfficeController)

		taxiOfficesRg.POST("/create", to.Create)
		taxiOfficesRg.POST("/recover/:id", to.Recover)
		taxiOfficesRg.PUT(":id", to.Update)
		taxiOfficesRg.GET(":id", to.Get)
		taxiOfficesRg.GET("/list", to.GetList)
		taxiOfficesRg.DELETE(":id", to.Delete)
	}

	carsRG := e.Group("cars")
	{
		c := new(routers.CarsController)

		carsRG.POST("/create", c.Create)
		carsRG.PUT(":id", c.Update)
		carsRG.GET(":id", c.Get)
		carsRG.GET("/list", c.GetList)
		carsRG.DELETE(":id", c.Delete)
	}
}
