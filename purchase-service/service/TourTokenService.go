package service

import (
	"purchase-service/models"
	"purchase-service/repository"
)

type TourTokenService struct {
	Repo repository.TourTokenRepo
}

func (tts *TourTokenService) Save(token *models.TourPackageToken) error {
	err := tts.Repo.Save(token)

	return err
}

func (tts *TourTokenService) Delete(token *models.TourPackageToken) error { 
	err := tts.Repo.Delete(token)

	return err
}

func (tts *TourTokenService) Update(token *models.TourPackageToken) error {
	err := tts.Repo.Update(token)
	
	return err
}
