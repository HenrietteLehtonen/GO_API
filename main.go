package main

import (
	"fmt"

	"github.com/henriettelehtonen/GO_API/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/henriettelehtonen/GO_API/routes"
)

func main() {

	fmt.Println("Hello World from GO API!")

	// FIBER TEST 
	app := fiber.New()
    // app.Get("/", func(c *fiber.Ctx) error {
    //     return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
    // })

	// Database connection
	 configs.ConnectDB()

	 // Routes
	routes.Router(app)

    app.Listen(":3000")
}