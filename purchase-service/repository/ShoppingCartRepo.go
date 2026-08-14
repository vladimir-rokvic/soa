package repository

import (
	"purchase-service/models"

	"gorm.io/gorm"
)

type ShoppingCartRepo struct {
	Db *gorm.DB
}


func (sr *ShoppingCartRepo) Update(sc *models.ShoppingCart) (*models.ShoppingCart, error) {
	result := sr.Db.Save(sc)

	return sc, result.Error
}

//Ovo je zapravo create i guess
func (sr *ShoppingCartRepo) Save(sc *models.ShoppingCart) (*models.ShoppingCart, error) {
	result := sr.Db.Create(sc)

	return sc, result.Error
}

func (sr *ShoppingCartRepo) Delete(sc *models.ShoppingCart) (*models.ShoppingCart, error) {
	result := sr.Db.Delete(sc)

	return sc, result.Error
}

func (sr *ShoppingCartRepo) GetAll() ([]models.ShoppingCart, error) {
	var shoppingCarts []models.ShoppingCart
	result := sr.Db.Preload("Items").Find(&shoppingCarts)

	return shoppingCarts, result.Error
}

func (sr *ShoppingCartRepo) GetById(id string) (models.ShoppingCart, error) {
	var cart models.ShoppingCart
	result := sr.Db.Preload("Items").First(&cart, "id = ?", id)

	return cart, result.Error
}

func (sr *ShoppingCartRepo) GetByUserId(id string) (models.ShoppingCart, error) {
	var cart models.ShoppingCart
	result := sr.Db.Preload("Items").First(&cart, "user_id = ?", id)

	return cart, result.Error
}
