package main

import (
	"fmt"
	"purchase-service/models"

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

}
