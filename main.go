package main

import (
	"go_gorms_lat/config"
	"go_gorms_lat/model"
	"go_gorms_lat/repo"
	"log"
)

func main() {

	config := config.NewConfigDB()

	db := config.DbConn()

	enigmaDb, _ := db.DB()

	defer config.DBTutup(enigmaDb)

	custRepo := repo.NewCustomerRepository(db)

	// Insert Data Baru

	// id := uuid.New()

	// customer01 := model.Customer{
	// 	Id:      id.String(),
	// 	Name:    "Utasmi",
	// 	Phone:   "3535456",
	// 	Balance: 20000,
	// }

	// err := custRepo.Create(&customer01)

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	//Update Data yang udh ada

	// customerUpdate := model.Customer{
	// 	Id: "36a3ebd0-9422-4973-9692-37e0114481bd",
	// }

	// err := custRepo.Update(&customerUpdate, map[string]interface{}{
	// 	"Name":    "Pong-Pong",
	// 	"Balance": "11000",
	// })

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	//Delete

	customerDelete := model.Customer{
		Id: "36a3ebd0-9422-4973-9692-37e0114481bd",
	}

	err := custRepo.Delete(&customerDelete)

	if err != nil {
		log.Println(err.Error())
	}

}
