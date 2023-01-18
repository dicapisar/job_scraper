package linkedin

import "github.com/dicapisar/job_scraper/domain"

type JobDetailCollectorResult struct {
	Title          string
	Id             string
	DateAgo        string
	Url            string
	Description    string
	Company        string
	SeniorityLevel string
	EmploymentType string
	JobFunction    string
	Industries     string
	Location       string
}

func (j *JobDetailCollectorResult) ParseToLinkedinJobModel() *domain.LinkedinJob {
	model := domain.LinkedinJob{
		JobId:          &j.Id,
		Title:          &j.Title,
		DateAgo:        &j.DateAgo,
		Url:            &j.Url,
		Description:    &j.Description,
		Company:        &j.Company,
		SeniorityLevel: &j.SeniorityLevel,
		EmploymentType: &j.EmploymentType,
		JobFunction:    &j.JobFunction,
		Industries:     &j.Industries,
		Location:       &j.Location,
	}

	return &model
}
