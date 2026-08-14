package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type OrderItem struct {
	ID uuid.UUID `json:"id"`
	TourName string `json:"tour_name"`
	Price float32 `json:"price"`
	TourId string `json:"tour_id"`
	ShoppingCartID uuid.UUID `json:"shopping_cart_id"`
}

func (oi *OrderItem) BeforeCreate(scope *gorm.DB) error {
	oi.ID = uuid.New()

	return nil
}
