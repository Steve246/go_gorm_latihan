package main

import (
	"go_gorms_lat/config"
	"go_gorms_lat/model"
)

func main() {

	config := config.NewConfigDB()

	db := config.DbConn()

	enigmaDb, _ := db.DB()

	defer config.DBTutup(enigmaDb)

	err := db.AutoMigrate(model.Customer{})

	if err != nil {
		panic(err)
	}

}
