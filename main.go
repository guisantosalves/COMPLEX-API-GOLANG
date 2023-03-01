package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/guisantosalves/go-api-fiber/src/database"
	Router "github.com/guisantosalves/go-api-fiber/src/routes"
	"github.com/joho/godotenv"
)

// b != nil -> true
// x != nil -> true
// c != nil -> true
// nil != nil -> false

func main() {
	app := fiber.New()

	// initialize and open migrations
	database.ConnectDB()

	// default route
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Go api with fiber")
	})

	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// verify env
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3005"
	}

	// set routes and up the server
	Router.SetRoutes(app)
	app.Listen(PORT)

}
