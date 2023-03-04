package env

import (
	"log"
	"os"
	"strconv"
)

type EnvList struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
	DBPort     int
}

func SetEnv() EnvList {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("SETENV error")
	}
	return EnvList{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     port,
	}
}
