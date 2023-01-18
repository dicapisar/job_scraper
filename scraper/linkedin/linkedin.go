package linkedin

import (
	"github.com/dicapisar/job_scraper/domain"
)

type Scraper struct {
}

func (s *Scraper) GenerateJobResults(search *domain.JobSearch) *[]domain.Job {

	collectorListJob := listJobCollector{}
	listJobCollectorResult := collectorListJob.GetJobList(search)
	removeExcessResults(listJobCollectorResult, search)

	collectorDetailJob := jobDetailCollector{}
	//JobDetailCollectorResultList := make([]JobDetailCollectorResult, 0, 1)

	jobs := make([]domain.Job, 0, 1)

	for _, jobResult := range *listJobCollectorResult {
		jobDetailResult := collectorDetailJob.GetDetailJob(&jobResult)
		jobs = append(jobs, jobDetailResult.ParseToLinkedinJobModel())
		//JobDetailCollectorResultList = append(JobDetailCollectorResultList, *jobDetailResult)
	}

	return &jobs
}

func removeExcessResults(listJobCollectorResult *[]JobInfoCollectorResult, search *domain.JobSearch) {
	countExcess := len(*listJobCollectorResult) - search.CountToFind

	if countExcess <= 0 {
		return
	}

	lastSlice := len(*listJobCollectorResult) - countExcess

	*listJobCollectorResult = (*listJobCollectorResult)[0:lastSlice]
}
