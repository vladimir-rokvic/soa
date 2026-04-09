package service

import (
	"auth_service/model"
	"auth_service/repo"
	"fmt"
)

type UserService struct {
	UserRepo *repo.UserRepo
}

func (user_service *UserService) GetAll() ([]model.User){
	users, err := user_service.UserRepo.GetAll()

	if err != nil {
		fmt.Println("Error fetching all users")
		fmt.Println(err)
		return nil
	}

	return users
}

func (user_service *UserService) GetById(id string) *model.User {
	user, err := user_service.UserRepo.GetById(id)

	if err != nil {
		fmt.Printf("Error fetching user with id: %d ", id)
		fmt.Println(err)

		return nil
	}

	return &user
}
