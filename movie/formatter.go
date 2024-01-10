package movie

import (
	"theater/entities"
	"time"
)

type MovieFormatter struct {
	ID          int       `json:"id"`
	Title       string    `json:"name"`
	Description string    `json:"description"`
	Rating      float32   `json:"rating"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FormatMovie(movie entities.Movie) MovieFormatter {
	formatter := MovieFormatter{
		ID:          int(movie.ID),
		Title:       movie.Title,
		Description: movie.Description,
		Rating:      movie.Rating,
		Image:       movie.Image,
		CreatedAt:   movie.CreatedAt,
		UpdatedAt:   movie.UpdatedAt,
	}

	return formatter
}
