package service

import (
	"auth_service/dto"
	"auth_service/model"
	"auth_service/repo"
	"fmt"

	"github.com/google/uuid"
)

type UserService struct {
	UserRepo *repo.UserRepo
}

func (user_service *UserService) GetAll() ([]dto.UserDTO){
	users, err := user_service.UserRepo.GetAll()

	if err != nil {
		fmt.Println("Error fetching all users")
		fmt.Println(err)
		return nil
	}

	var dtos []dto.UserDTO
	for _, u := range users {
		dtos = append(dtos, dto.UserToDTO(&u))
	}

	return dtos
}

func (user_service *UserService) GetById(id string) *dto.UserDTO {
	user, err := user_service.UserRepo.GetById(id)

	if err != nil {
		fmt.Printf("Error fetching user with id: %s ", id)
		fmt.Println(err)

		return nil
	}

	dto := dto.UserToDTO(&user)
	return &dto
}

func (user_service *UserService) GetByUUID(id uuid.UUID) *model.User {
	user, err := user_service.UserRepo.GetById(id.String())

	if err != nil {
		fmt.Printf("Error fetching user with id: %s ", id)
		fmt.Println(err)

		return nil
	}

	return &user
}

func (user_service *UserService) GetByUsername(username string) *model.User {
	user, err := user_service.UserRepo.GetByUsername(username)
	if err != nil {
		fmt.Printf("Error getting user with username: %s ", username)
		fmt.Println(err)
		return nil
	}

	return &user
}

//TODO:Check if a user with the given username already exists
func (user_service *UserService) Save(user *model.User) error {
	err := user_service.UserRepo.Save(user)
	return err
}

func (user_service *UserService) UpdateUser(user *model.User, dto *dto.UserUpdateDTO) error {
	user.Username = dto.Username
	user.Email = dto.Email
	user.Password = dto.Password
	user.UserRole = model.StringToUserRole(dto.UserRole)

	err := user_service.Save(user)

	return err
}
