package service

import (
	"encoding/json"
	"math"
	"net/http"
	"purchase-service/dto"
	"purchase-service/models"
	"purchase-service/repository"
	"time"

	"github.com/google/uuid"
)

type TourTokenService struct {
	Repo *repository.TourTokenRepo
}

func (tts *TourTokenService) Save(token *models.TourPurchaseToken) error {
	err := tts.Repo.Save(token)

	return err
}

func (tts *TourTokenService) Delete(token *models.TourPurchaseToken) error { 
	err := tts.Repo.Delete(token)

	return err
}

func (tts *TourTokenService) Update(token *models.TourPurchaseToken) error {
	err := tts.Repo.Update(token)
	
	return err
}

func (tts *TourTokenService) GetByUserId(id string) ([]models.TourPurchaseToken, error) {
	tokens, err := tts.Repo.GetByUserId(id)

	return tokens, err
}

func (tts *TourTokenService) GetActiveByUserId(id string) (models.TourExecution, error) {
	te, err := tts.Repo.GetActiveByUserId(id)

	return te, err
}

func (tts *TourTokenService) GetById(id string) (models.TourPurchaseToken, error) {
	token, err := tts.Repo.GetById(id)

	return token, err
}

func getPointFromDTO(pointDTO dto.InterestPointDTO, teId uuid.UUID) models.InterestPoint {
	point := models.InterestPoint {
		Lat: pointDTO.Lat,
		Lng: pointDTO.Lng,
		TourExecutionID: teId,
	}

	return point
}

func (tts *TourTokenService) CreateTourExecution(currentPostion dto.CurrentPosDTO, token models.TourPurchaseToken) (models.TourExecution, error) {
	te := models.TourExecution{
		TokenId: token.ID,
		UserId: token.UserId,
		Status: models.ACTIVE,
		CurrentLat: currentPostion.CurrentLat,
		CurrentLng: currentPostion.CurrentLng,
	}

	res, err := http.Get("http://tours-service:8080/tours/" + token.TourId)

	if err != nil {
		return te, err
	}

	var tour dto.TourDTO
	err = json.NewDecoder(res.Body).Decode(&tour)

	if err != nil {
		return te, err
	}

	err = tts.Repo.SaveTE(&te)

	if err != nil {
		return te, err
	}

	start_point := getPointFromDTO(tour.StartPoint, te.ID)
	err = tts.Repo.SavePoint(&start_point)

	if err != nil {
		return te, err
	}
	
	end_point := getPointFromDTO(tour.EndPoint, te.ID)
	err = tts.Repo.SavePoint(&end_point)

	if err != nil {
		return te, err
	}

	return te, err
}

func (tts *TourTokenService) GetTEById(id string) (models.TourExecution, error) {
	te, err := tts.Repo.GetTEById(id)

	return te, err
}

func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func distanceBetweenPoints(lat1, lng1, lat2, lng2 float64) float64 {
	//u metrima
	const earthRadius = 6371000.0

	dLat := toRadians(lat2 - lat1)
	dLng := toRadians(lng2 - lng1)

	lat1Rad := toRadians(lat1)
	lat2Rad := toRadians(lat2)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*
			math.Cos(lat2Rad)*
			math.Sin(dLng/2)*
			math.Sin(dLng/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}

func (tts *TourTokenService) UpdateTour(te *models.TourExecution, currentPos dto.CurrentPosDTO) (models.TourExecution, error){
	te.CurrentLat = currentPos.CurrentLat
	te.CurrentLng = currentPos.CurrentLng

	for i := range te.Points {
		p := &te.Points[i]

		d := distanceBetweenPoints(
			float64(currentPos.CurrentLat),
			float64(currentPos.CurrentLng),
			float64(p.Lat),
			float64(p.Lng),
			)

		if d <= 100 && p.TimeCompleted.IsZero() {
			p.TimeCompleted = time.Now()

			err := tts.Repo.UpdatePoint(p)
			if err != nil {
				return models.TourExecution{}, err
			}

			te.TimeEnded = p.TimeCompleted
		}
	}

	ret, err := tts.Repo.UpdateTE(te)

	if err != nil {
		return models.TourExecution{}, err
	}

	return ret, err
}
