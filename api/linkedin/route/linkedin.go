package route

import (
	"github.com/dicapisar/job_scraper/api/linkedin/handle"
	"github.com/gofiber/fiber/v2"
)

func GetRouter(app *fiber.App) fiber.Router {

	linkedinRouteGroup := app.Group("/linkedin").Name("linkedinHandleRoute")
	linkedinRouteGroup.Post("", handle.GenerateSearchHandle)

	return linkedinRouteGroup
}
