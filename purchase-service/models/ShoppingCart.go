package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type ShoppingCart struct {
	ID uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"user_id" gorm:"unique"`
	Items []OrderItem `json:"items"`
	Price float32 `json:"price"`
}

func (sc *ShoppingCart) BeforeCreate(scope *gorm.DB) error {
	sc.ID = uuid.New()
	sc.Price = 0

	return nil
}
