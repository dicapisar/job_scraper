package handle

import (
	"fmt"
	"github.com/dicapisar/job_scraper/api/linkedin/dto/request"
	"github.com/dicapisar/job_scraper/domain"
	"github.com/dicapisar/job_scraper/infra"
	"github.com/dicapisar/job_scraper/scraper/linkedin"
	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron/v3"
)

func GenerateSearchHandle(c *fiber.Ctx) error {
	search := request.Search{}
	err := c.BodyParser(&search)

	if err != nil {
		return err
	}

	validateError := search.Validate()

	if validateError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validateError)
	}

	scraper := linkedin.Scraper{}
	jobSearch := domain.JobSearch{}

	jobSearch.ParseFromLinkedinSearch(&search)

	result := scraper.GenerateJobResults(&jobSearch)

	err = generateCron()

	if err != nil {
		fmt.Println("error generating cron")
	}

	for _, job := range *result {

		err = infra.DBRepository.CreateLinkedinJob(job.(*domain.LinkedinJob))

		if err != nil {
			fmt.Printf("error on saved job on database: %v", err.Error())
		}
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func generateCron() error {

	c := cron.New(cron.WithSeconds())

	addFunc, err := c.AddFunc("@every 1m", func() {
		fmt.Println("every minute")
	})

	fmt.Println(addFunc) // guardar info

	if err != nil {
		return err
	}

	c.Start()

	return nil
}
