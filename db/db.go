package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

var DB *gorm.DB
var err error

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func DbURL() string {

	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbName := viper.GetString("DB_NAME")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")

	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbHost, dbPort)
}

func Open() {

	DB, err = gorm.Open(postgres.Open(DbURL()), &gorm.Config{})

	if err != nil {
		log.Fatal("Problem opening database: " + err.Error())
	}
}
