package repository

import (
	"purchase-service/models"

	"gorm.io/gorm"
)

type TourTokenRepo struct {
	Db *gorm.DB
}

func (ttr *TourTokenRepo) Save(token *models.TourPurchaseToken) error {
	result := ttr.Db.Create(token)

	return result.Error
}

func (ttr *TourTokenRepo) Delete(token *models.TourPurchaseToken) error {
	result := ttr.Db.Delete(token)

	return result.Error
}

func (ttr *TourTokenRepo) Update(token *models.TourPurchaseToken) error {
	result := ttr.Db.Save(token)

	return result.Error
}

func (ttr *TourTokenRepo) GetByUserId(id string) ([]models.TourPurchaseToken, error) {
	var tokens []models.TourPurchaseToken
	result := ttr.Db.Find(&tokens, "user_id = ?", id)

	return tokens, result.Error
}
