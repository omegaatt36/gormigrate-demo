package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/omegaatt36/gormigrate-demo/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Rdb *gorm.DB

func InitDb() {
	username := config.GetString("DB_USERNAME")
	password := config.GetString("DB_PASSWORD")
	host := config.GetString("DB_HOST")
	port := config.GetString("DB_PORT")
	options := config.GetString("DB_OPTIONS")
	databaseName := config.GetString("DB_DATABASE")

	if username == "" || password == "" || host == "" ||
		databaseName == "" {
		panic("[WARNING] GORM ENV is empty")
	}

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s %s", host, username, databaseName, password, options)
	if port != "" {
		dsn = fmt.Sprintf("%s port=%s", dsn, port)
	}

	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: 180 * time.Millisecond,
			LogLevel:      logger.Info,
			Colorful:      true,
		})})
	if err != nil {
		panic(err)
	}

	Rdb = db
}
