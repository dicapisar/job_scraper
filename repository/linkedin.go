package repository

import (
	"fmt"
	"github.com/dicapisar/job_scraper/domain"
)

func (r *Repository) CreateLinkedinJob(job *domain.LinkedinJob) error {
	err := r.DB.Create(&job).Error

	if err == nil {
		fmt.Printf("Job Linkedin saved success: %v \n", job)
	}

	return err
}
