package config

import (
	"database/sql"
	"fmt"
	"os"

	"go_gorms_lat/model"
	"go_gorms_lat/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Db *gorm.DB
}

func (c *Config) initDb() {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	env := os.Getenv("ENV")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	utils.IsError(err)

	c.Db = db

	if env == "dev" {
		c.Db = db.Debug()
	} else if env == "migration" {
		c.Db = db.Debug()
		err := c.Db.AutoMigrate(&model.Customer{})

		if err != nil {
			return
		}
	} else {
		c.Db = db
	}

}

func (c *Config) DbConn() *gorm.DB {
	return c.Db
}

func (c *Config) DBTutup(*sql.DB) {
	db, err := c.Db.DB()
	utils.IsError(err)
	err = db.Close()
	utils.IsError(err)

}

func NewConfigDB() Config {
	cfg := Config{}
	cfg.initDb()
	return cfg
}
