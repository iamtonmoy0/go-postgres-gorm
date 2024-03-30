package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct{
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func(r *Repository) SetupRoutes(app *fiber.App) {
api:=app.Group("/api")
	api.Post("/create-book",r.CreateBook)
	api.Get("/get-book/:id",r.GetBook)
	api.Get("/get-books",r.GetBooks)
	api.Put("/update-book/:id",r.UpdateBook)
	api.Delete("/delete-book/:id", r.DeleteBook)

}
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	r := Repository{
		DB:db
	}

	app := fiber.New()
	r.SetupRoutes(app)

	app.Listen(":3000")
}
