package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
	"github.com/qinains/fastergoding"
)

func main() {

    fastergoding.Run()

    engine := handlebars.New("./views", ".hbs")

    app := fiber.New(fiber.Config{
        Views: engine,
    })

    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("index", fiber.Map{})
    })

    app.Post("/evolution", func(c *fiber.Ctx) error {
        chains := GetEvolutionChain(strings.ToLower(c.FormValue("pokemon")))

        if len(chains) == 0 {
            return c.SendString("This pokemon doesn't have any evolution")
        }

        return c.Render("chain", fiber.Map{
            "Chains": chains,
        })
    })

    app.Listen(":3000")
}
