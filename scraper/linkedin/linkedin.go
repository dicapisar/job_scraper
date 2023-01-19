package linkedin

import (
	"github.com/dicapisar/job_scraper/domain"
	"github.com/dicapisar/job_scraper/scraper/linkedin/collector"
	"github.com/dicapisar/job_scraper/scraper/linkedin/result"
)

type Scraper struct {
}

func (s *Scraper) GenerateJobResults(search *domain.JobSearch) *[]domain.Job {

	collectorListJob := collector.ListJobCollector{}
	listJobCollectorResult := collectorListJob.GetJobList(search)
	removeExcessResults(listJobCollectorResult, search)

	collectorDetailJob := collector.JobDetailCollector{}

	jobs := make([]domain.Job, 0, 1)

	for _, jobResult := range *listJobCollectorResult {
		jobDetailResult := collectorDetailJob.GetDetailJob(&jobResult)
		jobs = append(jobs, jobDetailResult.ParseToLinkedinJobModel())
	}

	return &jobs
}

func removeExcessResults(listJobCollectorResult *[]result.JobInfoCollectorResult, search *domain.JobSearch) {
	countExcess := len(*listJobCollectorResult) - search.CountToFind

	if countExcess <= 0 {
		return
	}

	lastSlice := len(*listJobCollectorResult) - countExcess

	*listJobCollectorResult = (*listJobCollectorResult)[0:lastSlice]
}
