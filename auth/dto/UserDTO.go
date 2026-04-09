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
}

func UserToDTO(user *model.User) UserDTO {
	var dto UserDTO
	dto.ID = user.ID
	dto.Username = user.Username
	dto.Email = user.Email
	dto.UserRole = user.UserRole.String()
	return dto
}
