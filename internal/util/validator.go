package util

import (
	"givery-recip/internal/common"
	"givery-recip/internal/dto"
	"strings"
)

func ValidateUpdateRecipe(recipe *dto.UpdateRecipeRequest) *common.ValidationError {
	var missing []string

	if recipe.Title == "" {
		missing = append(missing, "title")
	}
	if recipe.MakingTime == "" {
		missing = append(missing, "making_time")
	}
	if recipe.Serves == "" {
		missing = append(missing, "serves")
	}
	if recipe.Ingredients == "" {
		missing = append(missing, "ingredients")
	}
	if recipe.Cost < 0 {
		missing = append(missing, "cost")
	}

	if len(missing) > 0 {
		return &common.ValidationError{Missing: strings.Join(missing, ", ")}
	}

	return nil
}

func ValidateCreateRecipe(recipe *dto.CreateRecipeRequest) *common.ValidationError {
	var missing []string

	if recipe.Title == "" {
		missing = append(missing, "title")
	}
	if recipe.MakingTime == "" {
		missing = append(missing, "making_time")
	}
	if recipe.Serves == "" {
		missing = append(missing, "serves")
	}
	if recipe.Ingredients == "" {
		missing = append(missing, "ingredients")
	}
	if recipe.Cost < 0 {
		missing = append(missing, "cost")
	}

	if len(missing) > 0 {
		return &common.ValidationError{Missing: strings.Join(missing, ", ")}
	}

	return nil
}
