package api

import (
	"fmt"
	"github.com/dicapisar/job_scraper/api/linkedin/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"os"
	"time"
)

func InitializeApi() {
	app := fiber.New(fiber.Config{
		AppName:           os.Getenv("NAME_APP"),
		EnablePrintRoutes: true,
	})

	// Middleware
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${blue} ${time} ${yellow} ${pid} ${blue} ${status} ${green} ${locals:requestid} " +
			"${yellow} - ${method} ${latency} ${path}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
	}))

	route.GetRouter(app)

	err := app.Listen(":3000")

	if err != nil {
		fmt.Println("Error initializing server...")
	}

}
