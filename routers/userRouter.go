package routers

import (
	"notes-application/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App){
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user:= v1.Group("/user")


	user.Post("/",controllers.SignUp)
	user.Post("/login",controllers.LogIn)
	user.Put("/",controllers.UpdatePassword)
	user.Delete("/",controllers.DeleteAccount)
	user.Get("/",controllers.LogOut)
}