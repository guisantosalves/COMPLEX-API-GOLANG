package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/guisantosalves/go-api-fiber/src/database"
	"github.com/guisantosalves/go-api-fiber/src/routes"
	"github.com/joho/godotenv"
)

// b != nil -> true
// x != nil -> true
// c != nil -> true
// nil != nil -> false

func main() {
	app := fiber.New()

	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// initialize and open migrations
	// open only one time and use that connection
	database.ConnectDB()

	// default route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Go api with fiber")
	})

	// verify env
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3005"
	}

	// set routes and up the server
	routes.SetRoutes(app)
	app.Listen(":" + PORT)

}
