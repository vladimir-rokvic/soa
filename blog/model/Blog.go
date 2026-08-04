package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Blog struct {
	ID uuid.UUID `json:"id"`
	AuthorID uuid.UUID `json:"author_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"date_created"`
	Comments []Comment `json:"comments"`
}

func (blog *Blog) BeforeCreate(scope *gorm.DB) error {
	blog.ID = uuid.New()
	return nil
}
