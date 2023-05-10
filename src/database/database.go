package database

import (
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var envNames = []string{
	"DB_HOST",
	"DB_USER",
	"DB_PASSWORD",
	"DB_DBNAME",
	"DB_PORT",
	"DB_SSLMODE",
	"DB_TIMEZONE",
}

var db *gorm.DB

func Init() (err error) {
	dsn := ""

	for n := range envNames {
		// Cut DB_ prefix
		str := strings.ToLower(strings.Split(envNames[n], "_")[1])

		// Like "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
		dsn += str + "=" + os.Getenv(envNames[n]) + " "
	}

	var error error

	db, error = gorm.Open(postgres.Open(strings.Trim(dsn, "")), &gorm.Config{})

	if error != nil {
		return error
	}
	return nil
}

func GetDb() *gorm.DB {
	return db
}
