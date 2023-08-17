package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Startpage",
		})
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("about", fiber.Map{
			"Title": "About",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
