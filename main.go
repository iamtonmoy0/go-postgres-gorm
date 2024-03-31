package main

import (
	"log"
	"net/http"

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

func(r *Repository)CreateBook(c *fiber.Ctx) error{
	book:Book{}
  //   response:=fiber.Map{"msg":""}
	err:=c.BodyParser(&body)
	if err != nil {
		log.Fatalf("failed to  parse body %v", err)
	}
	err:=r.DB.Create(&book)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(err)
	}
	c.Status(http.StatusOK).JSON(&fiber.Map{"msg":"new book created"})
	return nil
}


func(r *Repository)GetBook(c *fiber.Ctx) error{


}
func(r *Repository)GetBooks(c *fiber.Ctx) error{
	books:=&[]models.Books{}

err:=r.DB.Find(books)
if err != nil {
	c.Status(400).SendString("failed to find books")
	return err
}
c.Status(201).JSON(&fiber.Map{"msg":"Successfully fetched data ","data":books})
return nil
}


func(r *Repository)UpdateBook(c *fiber.Ctx) error{}
func(r *Repository)DeleteBook(c *fiber.Ctx) error{}


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
