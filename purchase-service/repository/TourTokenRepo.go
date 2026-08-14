package repository

import (
	"purchase-service/models"

	"gorm.io/gorm"
)

type TourTokenRepo struct {
	Db *gorm.DB
}

func (ttr *TourTokenRepo) Save(token *models.TourPackageToken) error {
	result := ttr.Db.Create(token)

	return result.Error
}

func (ttr *TourTokenRepo) Delete(token *models.TourPackageToken) error {
	result := ttr.Db.Delete(token)

	return result.Error
}

func (ttr *TourTokenRepo) Update(token *models.TourPackageToken) error {
	result := ttr.Db.Save(token)

	return result.Error
}
