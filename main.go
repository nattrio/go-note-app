package main

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nattrio/go-note-app/config"
	"github.com/nattrio/go-note-app/controller"
	"github.com/nattrio/go-note-app/model"
	"github.com/nattrio/go-note-app/repository"
	"github.com/nattrio/go-note-app/router"
	"github.com/nattrio/go-note-app/service"
)

func main() {
	fmt.Println("Running server...")

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("error loading config", err)
	}

	// Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("notes").AutoMigrate(&model.Note{})

	// Init Repository
	noteRepository := repository.NewNoteRepositoryImpl(db)

	// Init Service
	noteService := service.NewNoteServiceImpl(noteRepository, validate)

	// Init Controller
	noteController := controller.NewNoteController(noteService)

	// Routes
	routes := router.NewRouter(noteController)

	// Run App
	app := fiber.New()
	app.Mount("/api", routes)

	log.Fatal(app.Listen(":3000"))
}
