package usecase

import (
	"errors"
	"givery-recip/internal/dto"
	"givery-recip/internal/models"
	"givery-recip/internal/repository"
	"givery-recip/internal/util"
	"time"
)

type recipeUsecase struct {
	repo repository.RecipeRepository
}

func (r recipeUsecase) Create(recipe *dto.CreateRecipeRequest) (*models.Recipe, error) {
	if validate := util.ValidateCreateRecipe(recipe); validate != nil {
		return nil, validate
	}

	recipeCreated := models.Recipe{
		Title:       recipe.Title,
		MakingTime:  recipe.MakingTime,
		Serves:      recipe.Serves,
		Ingredients: recipe.Ingredients,
		Cost:        recipe.Cost,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return r.repo.Create(&recipeCreated)
}

func (r recipeUsecase) GetByID(id int64) (*models.Recipe, error) {
	if id <= 0 {
		return nil, errors.New("invalid recipe ID")
	}
	return r.repo.GetByID(id)
}

func (r recipeUsecase) GetList() ([]models.Recipe, error) {
	return r.repo.GetList()
}

func (r recipeUsecase) Update(id int64, recipe *dto.UpdateRecipeRequest) (*models.Recipe, error) {
	if id <= 0 {
		return nil, errors.New("invalid recipe ID")
	}

	if validate := util.ValidateUpdateRecipe(recipe); validate != nil {
		return nil, validate
	}

	recipeUpdated := models.Recipe{
		ID:          id,
		Title:       recipe.Title,
		MakingTime:  recipe.MakingTime,
		Serves:      recipe.Serves,
		Ingredients: recipe.Ingredients,
		Cost:        recipe.Cost,
		UpdatedAt:   time.Now(),
	}

	return r.repo.Update(&recipeUpdated)
}

func (r recipeUsecase) Delete(id int64) error {
	if id <= 0 {
		return errors.New("invalid recipe ID")
	}
	return r.repo.Delete(id)
}

func NewRecipeUsecase(repo repository.RecipeRepository) RecipeUsecase {
	return &recipeUsecase{repo: repo}
}
