package main

import (
	"blog_service/controller"
	"blog_service/model"
	"blog_service/repo"
	"blog_service/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init_db() *gorm.DB {
	connection_url := 
	"host=database user=postgres password=mypassword dbname=postgres port=5432"

	db, err := gorm.Open(postgres.Open(connection_url), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database")
		fmt.Println(err)
		return nil
	} else {
		fmt.Println("Succsessfully connected to the database!")
	}

	db.AutoMigrate(model.Blog{})

	return db
}

func main() {
	db := init_db()
	if db == nil {
		return
	}

	router := mux.NewRouter().PathPrefix("/blog").Subrouter()

	repo := repo.BlogRepo{Db: db}
	service := service.BlogService{Repository: &repo}
	controller := controller.BlogController{Service: &service}

	router.HandleFunc("/", controller.GetAll).Methods("GET")
	router.HandleFunc("/", controller.Save).Methods("POST")
	router.HandleFunc("/{id}", controller.GetById).Methods("GET")
	router.HandleFunc("/", controller.UpdateBlog).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}
