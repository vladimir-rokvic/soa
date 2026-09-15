package repository

import (
	"purchase-service/models"

	"gorm.io/gorm"
)


type OrderItemRepo struct {
	Db *gorm.DB
}


func (or *OrderItemRepo) Update(oi *models.OrderItem) (*models.OrderItem, error) {
	result := or.Db.Save(oi)

	return oi, result.Error
}

//I ovo je zapravo create i guess(iskreno ne znam kako da se osecam oko ovoga)
func (or *OrderItemRepo) Save(oi *models.OrderItem) (*models.OrderItem, error) {
	result := or.Db.Create(oi)

	return oi, result.Error
}

func (or *OrderItemRepo) Delete(oi *models.OrderItem) (*models.OrderItem, error) {
	result := or.Db.Delete(oi)

	return oi, result.Error
}

func (or *OrderItemRepo) GetAll() ([]models.OrderItem, error) {
	var items []models.OrderItem
	result := or.Db.Find(&items)

	return items, result.Error
}

func (or *OrderItemRepo) GetById(id string) (models.OrderItem, error) {
	var item models.OrderItem
	result := or.Db.First(&item, "id = ?", id)

	return item, result.Error
}

func (or *OrderItemRepo) GetByShoppingCartId(id string) ([]models.OrderItem, error) {
	var items []models.OrderItem
	result := or.Db.Find(&items, "shopping_cart_id = ?", id)

	return items, result.Error
}
