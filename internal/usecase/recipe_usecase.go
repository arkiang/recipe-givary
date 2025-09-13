package usecase

import (
	"givery-recip/internal/dto"
	"givery-recip/internal/models"
)

type RecipeUsecase interface {
	Create(recipe *dto.CreateRecipeRequest) (*models.Recipe, error)
	GetByID(id int64) (*models.Recipe, error)
	GetList() ([]models.Recipe, error)
	Update(id int64, recipe *dto.UpdateRecipeRequest) (*models.Recipe, error)
	Delete(id int64) error
}
