package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nattrio/go-note-app/controller"
)

func NewRouter(noteController *controller.NoteController) *fiber.App {
	router := fiber.New()

	router.Get("/heathchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "welcome to server",
		})
	})

	router.Route("/notes", func(route fiber.Router) {
		route.Post("/", noteController.Create)
		route.Get("/", noteController.FindAll)
	})

	router.Route("/notes/:noteId", func(route fiber.Router) {
		route.Patch("/", noteController.Update)
		route.Get("/", noteController.FindById)
		route.Delete("/", noteController.Delete)
	})

	return router
}
