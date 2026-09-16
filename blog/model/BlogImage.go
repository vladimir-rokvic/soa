package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogImage struct {
	ID uuid.UUID `json:"id"`
	BlogID uuid.UUID `json:"blog_id"`
	Path string `json:"path"`
}

func (image *BlogImage) BeforeCreate(scope *gorm.DB) error {
	image.ID = uuid.New()
	return nil
}
