package service

import (
	"purchase-service/models"
	"purchase-service/repository"
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

func (tts *TourTokenService) CreateTourExecution(token TourTokenService) (models.TourExecution, error) {
}
