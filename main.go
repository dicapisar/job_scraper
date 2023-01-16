package main

import (
	"fmt"
	"github.com/dicapisar/job_scraper/domain"
	"github.com/dicapisar/job_scraper/scraper/linkedin"
)

func main() {
	linkedinScraper := linkedin.Scraper{}

	search := domain.JobSearch{
		Title:       "Software Developer",
		CountToFind: 20,
		Location:    domain.Bogota,
	}

	jobResults := linkedinScraper.GenerateJobResults(&search)

	fmt.Println(len(*jobResults))

}
