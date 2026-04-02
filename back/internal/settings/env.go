package settings

import (
	"fmt"
	"os"
)

var DbHost string
var DbPort string
var DbUser string
var DbPassword string
var DbName string
var AppPort string

func LoadEnv() {
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
	AppPort = os.Getenv("APP_PORT")

	fmt.Println("Environment variables louded")
}
