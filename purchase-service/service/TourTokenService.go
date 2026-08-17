package service

import (
	"encoding/json"
	"net/http"
	"purchase-service/dto"
	"purchase-service/models"
	"purchase-service/repository"

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
