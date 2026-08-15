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
	db.AutoMigrate(models.TourPurchaseToken{})
	db.AutoMigrate(models.InterestPoint{})
	db.AutoMigrate(models.TourExecution{})

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

	token_repo := repository.TourTokenRepo{Db: db}
	token_service := service.TourTokenService{Repo: &token_repo}

	sc_controller := controller.ShoppingCartController{
		Service: &sc_service,
		ItemService: &oi_service,
		TokenService: &token_service,
	}

	token_controller := controller.TokenController{
		Service: &token_service,
	}

	router := mux.NewRouter()
	sc_router := router.PathPrefix("/purchase/sc").Subrouter()
	oi_router := router.PathPrefix("/purchase/oi").Subrouter()
	token_router := router.PathPrefix("/purchase/tokens").Subrouter()

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

	token_router.HandleFunc("/user/{id}", token_controller.GetByUserId).Methods("GET")
	token_router.HandleFunc("/active/user/{id}", token_controller.GetActiveByUserId).Methods("GET")
	token_router.HandleFunc("/start/{id}", token_controller.StartTour).Methods("POST")


	log.Fatal(http.ListenAndServe(":8080", router))
}
