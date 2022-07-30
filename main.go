package main

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/recover"
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
	app.Use(recover.New())

	//    app.Static("/", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{}, "layouts/main")
	})

	app.Post("/", func(c *fiber.Ctx) error {
		chains, err := GetEvolutionChain(strings.ToLower(c.FormValue("pokemon")))

		if err != nil {
			return c.Render("index", fiber.Map{
				"Error": err.Error(),
			}, "layouts/main")
		}

		if len(chains) == 0 {
			return c.Status(fiber.StatusNoContent).Render("index", fiber.Map{
				"Info": "This pokemon doesn't have any evolution",
			}, "layouts/main")
		}

		return c.Render("index", fiber.Map{
			"Pokemon": c.FormValue("pokemon"),
			"Chains":  chains,
		}, "layouts/main")
	})

	app.Listen(":3000")
}
