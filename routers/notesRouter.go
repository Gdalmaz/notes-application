package routers

import (
	"notes-application/controllers"

	"github.com/gofiber/fiber/v2"
)

func NotesRouter(app *fiber.App){
	api := app.Group("/api")
	v1 := api.Group("/v1")
	notes := v1.Group("/notes")


	notes.Post("/",controllers.CreateNotes)
	notes.Post("/update-notes",controllers.UpdateNotes)
	notes.Delete("/",controllers.DeleteNotes)
	notes.Get("/",controllers.GetAllNotes)
}