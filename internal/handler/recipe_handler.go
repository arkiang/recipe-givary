package handler

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"givery-recip/internal/common"
	"givery-recip/internal/dto"
	"givery-recip/internal/usecase"
	"strconv"
)

type RecipeHandler struct {
	uc usecase.RecipeUsecase
}

func NewRecipeHandler(app *fiber.App, uc usecase.RecipeUsecase) {
	h := &RecipeHandler{uc}

	app.Post("/recipes", h.Create)
	app.Get("/recipes/:id", h.GetByID)
	app.Get("/recipes", h.GetList)
	app.Patch("/recipes/:id", h.Update)
	app.Delete("/recipes/:id", h.Delete)
}

func (h *RecipeHandler) Create(c *fiber.Ctx) error {
	var recipe dto.CreateRecipeRequest
	if err := c.BodyParser(&recipe); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{
			Message: err.Error(),
		})
	}

	created, err := h.uc.Create(&recipe)
	if err != nil {
		var e *common.ValidationError
		switch {
		case errors.As(err, &e):
			return c.Status(fiber.StatusOK).JSON(dto.ValidationErrorResponse{
				Message:  "Recipe creation failed!",
				Required: &e.Missing,
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ValidationErrorResponse{
				Message: err.Error(),
			})
		}
	}

	resp := dto.CreateRecipeResponse{
		Message: "Recipe successfully created!",
		Recipe: []dto.RecipeInfoWithTimestamp{{
			ID:          created.ID,
			Title:       created.Title,
			MakingTime:  created.MakingTime,
			Serves:      created.Serves,
			Ingredients: created.Ingredients,
			Cost:        created.Cost,
			CreatedAt:   created.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   created.UpdatedAt.Format("2006-01-02 15:04:05"),
		}},
	}

	return c.JSON(resp)
}

func (h *RecipeHandler) GetByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{
			Message: "invalid recipe ID",
		})
	}

	recipe, err := h.uc.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ValidationErrorResponse{
			Message: err.Error(),
		})
	}

	if recipe == nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ValidationErrorResponse{
			Message: "recipe not found",
		})
	}

	resp := dto.GetRecipeResponse{
		Message: "Recipe details by id",
		Recipe: []dto.RecipeInfo{{
			ID:          recipe.ID,
			Title:       recipe.Title,
			MakingTime:  recipe.MakingTime,
			Serves:      recipe.Serves,
			Ingredients: recipe.Ingredients,
			Cost:        recipe.Cost,
		}},
	}

	return c.JSON(resp)
}

func (h *RecipeHandler) GetList(c *fiber.Ctx) error {
	recipes, err := h.uc.GetList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ValidationErrorResponse{
			Message: err.Error(),
		})
	}

	resp := dto.GetAllRecipesResponse{
		Recipes: make([]dto.RecipeInfo, len(recipes)),
	}

	for i, r := range recipes {
		resp.Recipes[i] = dto.RecipeInfo{
			ID:          r.ID,
			Title:       r.Title,
			MakingTime:  r.MakingTime,
			Serves:      r.Serves,
			Ingredients: r.Ingredients,
			Cost:        r.Cost,
		}
	}

	return c.JSON(resp)
}

func (h *RecipeHandler) Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{
			Message: "invalid recipe ID",
		})
	}

	var recipe dto.UpdateRecipeRequest
	if err := c.BodyParser(&recipe); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{
			Message: err.Error(),
		})
	}

	update, err := h.uc.Update(id, &recipe)
	if err != nil {
		var e *common.ValidationError
		switch {
		case errors.As(err, &e):
			return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{
				Message:  "Recipe update failed!",
				Required: &e.Missing,
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ValidationErrorResponse{
				Message: err.Error(),
			})
		}
	}

	if update == nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ValidationErrorResponse{
			Message: "recipe not found",
		})
	}

	resp := dto.UpdateRecipeResponse{
		Message: "Recipe successfully updated!",
		Recipe: []dto.RecipeInfo{{
			ID:          update.ID,
			Title:       update.Title,
			MakingTime:  update.MakingTime,
			Serves:      update.Serves,
			Ingredients: update.Ingredients,
			Cost:        update.Cost,
		}},
	}

	return c.JSON(resp)
}

func (h *RecipeHandler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{
			Message: "invalid recipe ID",
		})
	}

	err = h.uc.Delete(id)

	if errors.Is(err, sql.ErrNoRows) {
		return c.Status(fiber.StatusNotFound).JSON(dto.ValidationErrorResponse{
			Message: "No recipe found",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{"message": "Recipe successfully removed!"})
}
