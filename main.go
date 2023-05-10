package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/m1dySpeed/province_taxi_service/src/database"
	"github.com/m1dySpeed/province_taxi_service/src/rest"
)

func main() {

	s := gin.Default()

	var port string

	if err := godotenv.Load(); err != nil {
		gin.SetMode(gin.ReleaseMode)

		port = ":3000"
	}

	port = os.Getenv("HTTP_PORT")

	if err := database.Init(); err != nil {
		panic(err)
	}

	rest.Init(s)

	s.Run(port)
}
