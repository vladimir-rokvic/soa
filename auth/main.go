package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"auth_service/controller"
	"auth_service/model"
	"auth_service/repo"
	"auth_service/service"
)

func init_db() *gorm.DB {
	connection_url := 
	"host=database user=postgres password=mypassword dbname=postgres port=5432"

	db, err := gorm.Open(postgres.Open(connection_url), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database: ")
		fmt.Println(err)
		return nil
	}

	db.AutoMigrate(model.User{})

	return db
}

func main() {
	db := init_db()
	if db == nil {
		return
	}

	router := mux.NewRouter()

	repo := &repo.UserRepo{Db: db}
	ser := &service.UserService{UserRepo: repo}
	controller := controller.UserController{UserService: ser}

	router.HandleFunc("/users/all", controller.GetAll)

	log.Fatal(http.ListenAndServe(":8080", router))
}
