package dto

type OrderItemDTO struct {
	UserId string `json:"user_id"`
	TourTitle string `json:"tour_title"`
	TourId string `json:"tour_id"`
	TourPrice float32 `json:"tour_price"`
}
