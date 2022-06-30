package main

import (
	"go_gorms_lat/config"
	"go_gorms_lat/model"
	"go_gorms_lat/repo"
	"log"

	"github.com/google/uuid"
)

func main() {

	config := config.NewConfigDB()

	db := config.DbConn()

	enigmaDb, _ := db.DB()

	defer config.DBTutup(enigmaDb)

	custRepo := repo.NewCustomerRepository(db)

	//Insert Data Baru

	id := uuid.New()

	customer01 := model.Customer{
		Id:      id.String(),
		Name:    "Poki Poi",
		Phone:   "07689999",
		Balance: 20000,
	}

	err := custRepo.Create(&customer01)

	if err != nil {
		log.Println(err.Error())
	}

	//Update Data yang udh ada

	// customerUpdate := model.Customer{
	// 	Id: "e4c95648-6fbe-49b1-b88b-c91a048c4ceb",
	// }

	// err := custRepo.Update(&customerUpdate, map[string]interface{}{
	// 	"Name":    "Pong-Pong",
	// 	"Balance": "11000",
	// })

	// if err != nil {
	// 	log.Println(err.Error())
	// }

}
