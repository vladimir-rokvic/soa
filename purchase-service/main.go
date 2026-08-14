package main

import (
	"fmt"
	"log"
	"net/http"
	"purchase-service/controller"
	"purchase-service/models"
	"purchase-service/repository"
	"purchase-service/service"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init_db() *gorm.DB {
	dsn := "host=database user=postgres password=mypassword dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to the database")
		fmt.Println(err)
		return nil
	} else {
		fmt.Println("Succsessfully connected to the database!")
	}

	db.AutoMigrate(models.OrderItem{})
	db.AutoMigrate(models.ShoppingCart{})

	return db;
}

func main() {

	db := init_db()
	if db == nil {
		fmt.Println("GG")
		return
	}

	sc_repo := repository.ShoppingCartRepo{Db: db}
	sc_service := service.ShoppingCartService{Repo: &sc_repo}

	oi_repo := repository.OrderItemRepo{Db: db}
	oi_service := service.OrderItemService{Repo: &oi_repo}
	oi_controller := controller.OrderItemController{Service: &oi_service}

	sc_controller := controller.ShoppingCartController{
		Service: &sc_service,
		ItemService: &oi_service,
	}

	router := mux.NewRouter()
	sc_router := router.PathPrefix("/purchase/sc").Subrouter()
	oi_router := router.PathPrefix("/purchase/oi").Subrouter()

	sc_router.HandleFunc("/", sc_controller.GetAll).Methods("GET")
	sc_router.HandleFunc("/", sc_controller.Save).Methods("POST")
	sc_router.HandleFunc("/{id}", sc_controller.Delete).Methods("DELETE")
	sc_router.HandleFunc("/", sc_controller.Update).Methods("PUT")
	sc_router.HandleFunc("/{id}", sc_controller.GetById).Methods("GET")
	sc_router.HandleFunc("/user/{id}", sc_controller.GetByUserId).Methods("GET")
	sc_router.HandleFunc("/user/{id}", sc_controller.BuyItems).Methods("POST")
	sc_router.HandleFunc("/addItem", sc_controller.AddItem).Methods("POST")

	oi_router.HandleFunc("/", oi_controller.GetAll).Methods("GET")
	oi_router.HandleFunc("/", oi_controller.Save).Methods("POST")
	oi_router.HandleFunc("/{id}", oi_controller.Delete).Methods("DELETE")
	oi_router.HandleFunc("/", oi_controller.Update).Methods("PUT")
	oi_router.HandleFunc("/{id}", oi_controller.GetById).Methods("GET")
	oi_router.HandleFunc("/sc/{id}", oi_controller.GetByShoppingCartId).Methods("GET")


	log.Fatal(http.ListenAndServe(":8080", router))
}
