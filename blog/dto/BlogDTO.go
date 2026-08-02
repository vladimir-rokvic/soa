package dto

import (
	"time"

	"github.com/google/uuid"
)

type BlogDTO struct {
	ID uuid.UUID `json:"id"`
	AuthorID uuid.UUID `json:"author_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"date_created"`
}
