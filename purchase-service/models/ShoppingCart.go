package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type ShoppingCart struct {
	ID uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id" gorm:"unique"`
	Items []OrderItem `json:"items"`
}

func (sc *ShoppingCart) BeforeCreate(scope *gorm.DB) error {
	sc.ID = uuid.New()

	return nil
}
