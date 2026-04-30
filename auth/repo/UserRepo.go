package repo

import (
	"auth_service/model"

	"golang.org/x/crypto/bcrypt"
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
	result := user_repo.Db.First(&user, "id = ?", id)

	return user, result.Error
}

func (user_repo *UserRepo) Save(user *model.User) error {
	//hashovanje sifre radi bezbednosti
	//10 je default vrednost
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashed)
	if err != nil {
		return err
	}

	result := user_repo.Db.Save(user)
	return result.Error
}

func (user_repo *UserRepo) GetByUsername(username string) (model.User, error) {
	var user model.User
	result := user_repo.Db.First(&user, "username = ?", username)

	return user, result.Error
}
