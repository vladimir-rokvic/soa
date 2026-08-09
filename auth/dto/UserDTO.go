package dto

import (
	"auth_service/model"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID uuid.UUID `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	UserRole string `json:"role"`
	ImagePath string `json:"image_path"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Bio string `json:"bio"`
	Motto string `json:"motto"`
}

func UserToDTO(user *model.User) UserDTO {
	var dto UserDTO
	dto.ID = user.ID
	dto.Username = user.Username
	dto.Email = user.Email
	dto.UserRole = user.UserRole.String()
	dto.ImagePath = user.ImagePath 
	dto.FirstName = user.FirstName 
	dto.LastName = user.LastName 
	dto.Bio = user.Bio 
	dto.Motto = user.Motto 
	return dto
}
