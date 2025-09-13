package repository

import (
	"givery-recip/internal/models"
)

type RecipeRepository interface {
	Create(recipe *models.Recipe) (*models.Recipe, error)
	GetByID(id int64) (*models.Recipe, error)
	GetList() ([]models.Recipe, error)
	Update(recipe *models.Recipe) (*models.Recipe, error)
	Delete(id int64) error
}
