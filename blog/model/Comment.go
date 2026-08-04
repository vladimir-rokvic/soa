package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID uuid.UUID `json:"id"`
	Text string `json:"text"`
	CreatedAt time.Time `json:"date_posted"`
	AuthorID uuid.UUID `json:"author_id"`
	BlogID uuid.UUID `json:"blog_id"`
}

func (c *Comment) BeforeCreate (tx *gorm.DB) (error) {
	c.ID = uuid.New()

	return nil
}
