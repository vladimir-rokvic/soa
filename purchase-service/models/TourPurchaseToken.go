package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TokenState int

const (
	INACTIVE TokenState = iota
	ACTIVE
	COMPLETED
	ABANDONED
)

type TourPurchaseToken struct {
	ID uuid.UUID
	UserId uuid.UUID
	TourId string
	Price float32
	Status TokenState
}

var stateName = map[TokenState]string {
	INACTIVE: "Inactive",
	ACTIVE: "Active",
	COMPLETED: "Completed",
	ABANDONED: "Abandoned",
}

func (ts TokenState) String() string {
	return stateName[ts]
}

func GenerateToken(userId uuid.UUID, item OrderItem) TourPurchaseToken {
	var ret TourPurchaseToken

	ret.UserId = userId
	ret.TourId = item.TourId
	ret.Price = item.Price
	ret.Status = INACTIVE

	return ret
}

func (tpt *TourPurchaseToken) BeforeCreate(scope *gorm.DB) error {
	tpt.ID = uuid.New()
	return nil
}
