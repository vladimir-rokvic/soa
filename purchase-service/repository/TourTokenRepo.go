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

func (ttr *TourTokenRepo) GetActiveByUserId(id string) (models.TourExecution, error) {
	var te models.TourExecution
	result := ttr.Db.Where("status = 1").Preload("Points").First(&te, "user_id = ?", id)

	return te, result.Error
}

func (ttr *TourTokenRepo) GetById(id string) (models.TourPurchaseToken, error) {
	var token models.TourPurchaseToken
	result := ttr.Db.First(&token, "id = ?", id)

	return token, result.Error
}

//Zapravo je create
func (ttr *TourTokenRepo) SaveTE(te *models.TourExecution) error {
	result := ttr.Db.Create(te)

	return result.Error
}

func (ttr *TourTokenRepo) SavePoint(point *models.InterestPoint) error {
	result := ttr.Db.Create(point)

	return result.Error
}

func (ttr *TourTokenRepo) GetAllTE() ([]models.TourExecution, error) {
	var tes []models.TourExecution
	result := ttr.Db.Preload("Points").Find(&tes)

	return tes, result.Error
}

func (ttr *TourTokenRepo) GetTEById(id string) (models.TourExecution, error) {
	var te models.TourExecution
	result := ttr.Db.Preload("Points").First(&te, "id = ?", id)

	return te, result.Error
}
