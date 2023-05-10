package routers

import "github.com/gin-gonic/gin"

const CarsTableName = "cars"

type CarsController struct{}

func (cc *CarsController) Create(ctx *gin.Context) {}

func (cc *CarsController) Get(ctx *gin.Context) {}

func (cc *CarsController) GetList(ctx *gin.Context) {}

func (cc *CarsController) Update(ctx *gin.Context) {}

func (cc *CarsController) Delete(ctx *gin.Context) {}
