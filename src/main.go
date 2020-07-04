package main

import (
	"github.com/filecoffee/proxy/modules"
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
	"log"
	"net/http"
	"os"
)

func main() {
	/* sets up the cache */
	modules.CacheSetup()

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
		data, err := modules.GetFromCache(c.Params("upload"))
		if err == nil {
			c.Send(data)
			return
		}

		request := modules.GetUpload(c.Params("upload"))
		if request.StatusCode != http.StatusOK {
			_ = c.Status(http.StatusNotFound).Render("index", fiber.Map{
				"Content": "Upload not found",
			})
			return
		}
		c.Send(request.Body)
		request.Body.Close()

		if os.Getenv("CACHE") == "true" {
			if modules.CheckIfCached(c.Params("upload")) {
				return
			}
			err = modules.Cache(c.Params("upload"), modules.GetUpload(c.Params("upload")))
			if err != nil {
				log.Print(err)
			}
		}

		modules.Setup()
		return
	})

	app.Get("*", func(c *fiber.Ctx) {
		_ = c.Render("index", fiber.Map{
			"Content": "Content not found",
		})
	})

	_ = app.Listen("0.0.0.0:8080")
}
