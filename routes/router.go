package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/henriettelehtonen/GO_API/controllers"
)

func Router(app *fiber.App) {
	// ENDPOINTS
	app.Get("/birds", controllers.GetBirds)
	app.Get("/birds/:id", controllers.GetSingleBirdByID)
	app.Post("/birds", controllers.CreateBird)
	app.Delete("/birds/:id", controllers.DeleteBirdByID)
	// testataan patchia, koska ei haluta päivittää koko lintua
	app.Patch("/birds/:id", controllers.UpdateBirdByID)
	app.Get("/uhanalaisuus", controllers.GetUhanalaisuus)


}