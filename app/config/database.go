package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/freekup/product-scrapper/app"
	"github.com/freekup/product-scrapper/app/repository"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	Username string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASS"`
	DBName   string `mapstructure:"DB_NAME"`
}

func initDatabaseConfig() {
	dbConfig := DBConfig{}
	viper.Unmarshal(&dbConfig)

	if dbConfig.Host == "" {
		log.Fatal("DB config is empty")
	}

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.DBName)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	app.App.Repository = repository.NewRepository(db)
}
