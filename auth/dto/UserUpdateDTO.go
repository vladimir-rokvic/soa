package dto

type UserUpdateDTO struct {
	Username string `json:"username"`
	Email string `json:"email"`
	UserRole string `json:"role"`
	ImagePath string `json:"image_path"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Bio string `json:"bio"`
	Motto string `json:"motto"`
}
