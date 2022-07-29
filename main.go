package main

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use("/static", filesystem.New(filesystem.Config{
		Root:   http.Dir("./static"),
		Browse: true,
	}))

	//    app.Static("/", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{}, "layouts/main")
	})

	app.Post("/evolution", func(c *fiber.Ctx) error {
		chains := GetEvolutionChain(strings.ToLower(c.FormValue("pokemon")))

		if len(chains) == 0 {
			return c.SendString("This pokemon doesn't have any evolution")
		}

		return c.Render("chain", fiber.Map{
			"Pokemon": c.FormValue("pokemon"),
			"Chains":  chains,
		}, "layouts/main")
	})

	app.Listen(":3000")
}
