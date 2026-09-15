package service

import (
	"fmt"
	"purchase-service/models"
	"purchase-service/repository"
)

type OrderItemService struct {
	Repo *repository.OrderItemRepo
}


func (os *OrderItemService) Update(oi *models.OrderItem) (*models.OrderItem, error) {

	oi_old, err := os.Repo.GetById(oi.ID.String())
	if err != nil {
		fmt.Println("Error getting old item")

		return nil, err
	}

	//Ne znam da li je ovo najbolji nacin da se ovo uradi
	if oi.Price != oi_old.Price {
		oi_old.Price = oi.Price
	}
	if oi.ShoppingCartID != oi_old.ShoppingCartID {
		oi_old.ShoppingCartID = oi.ShoppingCartID
	}
	if oi.TourName != oi_old.TourName {
		oi_old.TourName = oi.TourName
	}
	if oi.TourId != oi_old.TourId {
		oi_old.TourId = oi.TourId
	}

	item, err := os.Repo.Update(&oi_old)

	return item, err
}

func (os *OrderItemService) Save(oi *models.OrderItem) (*models.OrderItem, error) {
	item, err := os.Repo.Save(oi)

	return item, err
}

func (os *OrderItemService) Delete(id string) (*models.OrderItem, error) {
	item, err := os.Repo.GetById(id)

	if err != nil { 
		fmt.Println("Error getting item for deletion")

		return nil, err
	}

	_, err = os.Repo.Delete(&item)

	return &item, err
}

func (os *OrderItemService) GetAll() ([]models.OrderItem, error) {
	items, err := os.Repo.GetAll()

	return items, err
}

func (os *OrderItemService) GetById(id string) (models.OrderItem, error) {
	item, err := os.Repo.GetById(id)

	return item, err
}

func (os *OrderItemService) GetByShoppingCartId(id string) ([]models.OrderItem, error) {
	items, err := os.Repo.GetByShoppingCartId(id)

	return items, err
}
