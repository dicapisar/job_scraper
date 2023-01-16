package scraper

import "github.com/dicapisar/job_scraper/domain"

type Scraper interface {
	GenerateJobResults(search *domain.JobSearch) *[]domain.JobResult
}
