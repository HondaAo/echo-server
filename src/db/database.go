package db

import (
	"fmt"
	"log"

	"github.com/HondaAo/snippet/src/env"
	"github.com/HondaAo/snippet/src/pkg/driver/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB *gorm.DB

func InitDB(env env.EnvList) DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", env.DBUser, env.DBPassword, env.DBHost, env.DBPort, env.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}

	err = db.AutoMigrate(&models.Video{}, &models.Script{}, &models.Idioms{}, &models.ScriptIdioms{}, &models.Words{})
	if err != nil {
		log.Fatal("DB migration error")
	}
	return db
}
