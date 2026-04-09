package repo

import (
	"auth_service/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func (user_repo *UserRepo) GetAll() ([]model.User, error) {
	var users []model.User
	result := user_repo.Db.Find(&users)
	return users, result.Error
}

func (user_repo *UserRepo) GetById(id string) (model.User, error) {
	var user model.User
	result := user_repo.Db.First(&user, id)

	return user, result.Error
}
