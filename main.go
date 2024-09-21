package main

import (
	"enakazoo-backend/infra"
	"enakazoo-backend/internal/handlers"

	_ "enakazoo-backend/docs"

	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
)

// @title Anekazoo API
// @version 1.0
// @description API documentation for Anekazoo backend.
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	db := infra.InitDB()

	handlers.NewAnimalHandler(app, db)

	app.Listen(":3000")
}
