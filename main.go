package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"givery-recip/infrastructure/persistent"
	"givery-recip/internal/handler"
	"givery-recip/internal/usecase"

	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sqlx.Connect("sqlite3", "./data.db")
	if err != nil {
		log.Fatalln("Failed to connect to DB:", err)
	}

	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatalln(err)
	}
	if err := goose.Up(db.DB, "./migrations"); err != nil {
		log.Fatalln("Failed to run migrations:", err)
	}

	recipeRepo := persistent.NewRecipeRepository(db)
	recipeUC := usecase.NewRecipeUsecase(recipeRepo)

	app := fiber.New()
	handler.NewRecipeHandler(app, recipeUC)

	log.Println("Server running at http://localhost:8081")
	if err := app.Listen(":8081"); err != nil {
		log.Fatalln(err)
	}
}
