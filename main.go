package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/iamtonmoy0/go-postgres-gorm/models"
	"github.com/iamtonmoy0/go-postgres-gorm/storage"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create-book", r.CreateBook)
	api.Get("/get-book/:id", r.GetBook)
	api.Get("/get-books", r.GetBooks)
	api.Put("/update-book/:id", r.UpdateBook)
	api.Delete("/delete-book/:id", r.DeleteBook)

}

func (r *Repository) CreateBook(c *fiber.Ctx) error {
	book := Book{}
	//   response:=fiber.Map{"msg":""}
	err := c.BodyParser(&book)
	if err != nil {
		log.Fatalf("failed to  parse body %v", err)
	}
	err = r.DB.Create(&book).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(err)
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"msg": "new book created"})
	return nil
}

func (r *Repository) GetBook(c *fiber.Ctx) error {
	return nil
}
func (r *Repository) GetBooks(c *fiber.Ctx) error {
	books := &[]models.Books{}

	err := r.DB.Find(books)
	if err != nil {
		c.Status(400).SendString("failed to find books")

	}
	c.Status(201).JSON(&fiber.Map{"msg": "Successfully fetched data ", "data": books})
	return nil
}

func (r *Repository) UpdateBook(c *fiber.Ctx) error {
	book := models.Books{}
	id := c.Params("id")
	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{"msg": "id should not be empty"})
		return nil
	}
	err := r.DB.Delete(book, id)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{"msg": "failed to delete the item"})
	}
	c.Status(http.StatusOK).JSON(fiber.Map{"msg": "Data deleted successfully"})
	return nil

}
func (r *Repository) DeleteBook(c *fiber.Ctx) error {
	return nil

}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// config
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USERNAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database")
	}
	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal(err)
	}
	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)

	app.Listen(":3000")
}
