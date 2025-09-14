package dto

// Request DTOs
type CreateRecipeRequest struct {
	Title       string `json:"title"`
	MakingTime  string `json:"making_time"`
	Serves      string `json:"serves"`
	Ingredients string `json:"ingredients"`
	Cost        int    `json:"cost"`
}

type UpdateRecipeRequest struct {
	Title       string `json:"title"`
	MakingTime  string `json:"making_time"`
	Serves      string `json:"serves"`
	Ingredients string `json:"ingredients"`
	Cost        int    `json:"cost"`
}

// Response DTOs
type CreateRecipeResponse struct {
	Message string                    `json:"message"`
	Recipe  []RecipeInfoWithTimestamp `json:"recipe"`
}

type GetRecipeResponse struct {
	Message string       `json:"message,omitempty"`
	Recipe  []RecipeInfo `json:"recipe"`
}

type GetAllRecipesResponse struct {
	Recipes []RecipeInfo `json:"recipes"`
}

type UpdateRecipeResponse struct {
	Message string       `json:"message"`
	Recipe  []RecipeInfo `json:"recipe"`
}

type DeleteRecipeResponse struct {
	Message string `json:"message"`
}

// RecipeInfo used inside response
type RecipeInfo struct {
	ID          int64  `json:"id,omitempty"`
	Title       string `json:"title"`
	MakingTime  string `json:"making_time"`
	Serves      string `json:"serves"`
	Ingredients string `json:"ingredients"`
	Cost        int    `json:"cost"`
}

type RecipeInfoWithTimestamp struct {
	ID          int64  `json:"id,omitempty"`
	Title       string `json:"title"`
	MakingTime  string `json:"making_time"`
	Serves      string `json:"serves"`
	Ingredients string `json:"ingredients"`
	Cost        int    `json:"cost"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type ValidationErrorResponse struct {
	Message  string  `json:"message"`
	Required *string `json:"required"`
}
