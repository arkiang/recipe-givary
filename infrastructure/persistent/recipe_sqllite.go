package persistent

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"givery-recip/internal/models"
	"givery-recip/internal/repository"
)

type recipeRepository struct {
	db *sqlx.DB
}

func (r recipeRepository) Create(recipe *models.Recipe) (*models.Recipe, error) {
	query := `
		INSERT INTO recipes (title, making_time, serves, ingredients, cost, created_at, updated_at)
		VALUES (:title, :making_time, :serves, :ingredients, :cost, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`
	result, err := r.db.NamedExec(query, recipe)
	if err != nil {
		return nil, err
	}

	// Get the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Fetch the full created record including created_at / updated_at
	var created models.Recipe
	err = r.db.Get(&created, `SELECT * FROM recipes WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (r recipeRepository) GetByID(id int64) (*models.Recipe, error) {
	var recipe models.Recipe
	err := r.db.Get(&recipe, `SELECT * FROM recipes WHERE id = ?`, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &recipe, nil
}

func (r recipeRepository) GetList() ([]models.Recipe, error) {
	var recipes []models.Recipe
	err := r.db.Select(&recipes, `SELECT * FROM recipes`)
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

func (r recipeRepository) Update(recipe *models.Recipe) (*models.Recipe, error) {
	query := `
		UPDATE recipes
		SET title = :title,
		    making_time = :making_time,
		    serves = :serves,
		    ingredients = :ingredients,
		    cost = :cost,
		    updated_at = CURRENT_TIMESTAMP
		WHERE id = :id`
	result, err := r.db.NamedExec(query, recipe)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, nil
	}

	// Fetch the updated record
	var updated models.Recipe
	err = r.db.Get(&updated, `SELECT * FROM recipes WHERE id = ?`, recipe.ID)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func (r recipeRepository) Delete(id int64) error {
	result, err := r.db.Exec(`DELETE FROM recipes WHERE id = ?`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func NewRecipeRepository(db *sqlx.DB) repository.RecipeRepository {
	return &recipeRepository{db: db}
}
