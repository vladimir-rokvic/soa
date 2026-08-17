package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InterestPoint struct {
	ID uuid.UUID
	Lat float32
	Lng float32
	TimeCompleted time.Time
	TourExecutionID uuid.UUID `json:"tour_execution_id"`
}

type TourExecution struct {
	ID uuid.UUID `json:"id"`
	TokenId uuid.UUID `json:"token_id"`
	UserId uuid.UUID `json:"user_id"`
	Status TokenState `json:"status"`
	CurrentLat float32 `json:"current_lat"`
	CurrentLng float32 `json:"current_lng"`
	//Ovo ce biti januar 1, godina 1 mogu da koristim IsZero()
	TimeEnded time.Time `json:"time_ended"`
	CreatedAt time.Time `json:"time_started"`
	Points []InterestPoint `json:"points"`
}

func (te *TourExecution) BeforeCreate(scope *gorm.DB) error {
	te.ID = uuid.New()

	return nil
}

func (ip *InterestPoint) BeforeCreate(scope *gorm.DB) error {
	ip.ID = uuid.New()

	return nil
}
