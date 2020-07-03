package main

import (
	"github.com/filecoffee/proxy/modules"
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
	"net/http"
)

func main() {
	engine := html.New("./views", ".html")
	engine.Reload(true)
	app := fiber.New(&fiber.Settings{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) {
		_ = c.Render("index", fiber.Map{
			"Content": "This is a file.coffee proxy instance, for all features head over to the official site.",
		})
	})

	app.Get("/u/:upload", func(c *fiber.Ctx) {
		request := modules.GetUpload(c.Params("upload"))
		if request.StatusCode != http.StatusOK {
			_ = c.Status(http.StatusNotFound).Render("index", fiber.Map{
				"Content": "Upload not found",
			})
			return
		}
		c.Send(request.Body)
		request.Body.Close()
		return
	})

	app.Get("*", func(c *fiber.Ctx) {
		_ = c.Render("index", fiber.Map{
			"Content": "Content not found",
		})
	})

	_ = app.Listen("0.0.0.0:8080")
}
