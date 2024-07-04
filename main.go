package main

import (
	"notes-application/database"
	"notes-application/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	routers.UserRouter(app)
	routers.NotesRouter(app)
	app.Listen(":8000")
}
