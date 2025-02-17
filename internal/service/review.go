package service

import (
	"errors"
	"pizzaria/internal/models"
)

func ValidaReviewRating(review *models.Review) error {
	if review.Rating < 1 || review.Rating > 5 {
		return errors.New("Rating must be between 1 e 5")
	}
	return nil
}
