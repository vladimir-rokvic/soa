package dto

type UserUpdateDTO struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	UserRole string `json:"role"`
}
