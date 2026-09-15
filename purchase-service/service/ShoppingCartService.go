package service

import (
	"fmt"
	"purchase-service/models"
	"purchase-service/repository"
)

type ShoppingCartService struct {
	Repo *repository.ShoppingCartRepo
}


func (ss *ShoppingCartService) Update(sc *models.ShoppingCart) (*models.ShoppingCart, error) {
	cart, err := ss.Repo.Update(sc)

	return cart, err
}

func (ss *ShoppingCartService) Save(sc *models.ShoppingCart) (*models.ShoppingCart, error) {
	cart, err := ss.Repo.Save(sc)

	return cart, err
}

func (ss *ShoppingCartService) Delete(id string) (*models.ShoppingCart, error) {
	cart, err := ss.Repo.GetById(id)

	if err != nil {
		fmt.Println("Error getting cart to delete")

		return nil, err
	}

	_, err = ss.Repo.Delete(&cart)

	return &cart, err
}

func (ss *ShoppingCartService) GetAll() ([]models.ShoppingCart, error) {
	carts, err := ss.Repo.GetAll()

	return carts, err
}

func (ss *ShoppingCartService) GetById(id string) (models.ShoppingCart, error) {
	cart, err := ss.Repo.GetById(id)

	return cart, err
}

func (ss *ShoppingCartService) GetByUserId(id string) (models.ShoppingCart, error) {
	cart, err := ss.Repo.GetByUserId(id)

	return cart, err
}
