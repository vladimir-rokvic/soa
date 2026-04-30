package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole int
const (
	Admin UserRole = iota
	Vodic
	Turista
)

var userRoleNames = map[UserRole] string {
	Admin: "admin",
	Vodic: "vodic",
	Turista: "turista",
}

func (ur UserRole) String() string {
	return userRoleNames[ur]
}

var namesUserRole = map[string] UserRole {
	"admin": Admin,
	"vodic": Vodic,
	"turista": Turista,
}

func StringToUserRole(role string) UserRole {
	return namesUserRole[role]
}

type User struct {
	ID uuid.UUID `json:"id"`
	Username string `json:"username" gorm:"unique"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	UserRole UserRole `json:"role"`
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.ID = uuid.New()
	return nil
}
