package domain

import "gorm.io/gorm"

type LinkedinJob struct {
	Id             uint    `gorm:"primary key;autoIncrement;type:serial" json:"id"`
	JobId          *string `json:"jobId"`
	Title          *string `json:"title"`
	DateAgo        *string `json:"date_ago"`
	Url            *string `json:"url"`
	Description    *string `json:"description"`
	Company        *string `json:"company"`
	SeniorityLevel *string `json:"seniority_level"`
	EmploymentType *string `json:"employment_type"`
	JobFunction    *string `json:"job_function"`
	Industries     *string `json:"industries"`
	Location       *string `json:"location"`
}

func (l *LinkedinJob) GetTypeScraper() *string {
	typeScraper := "linkedin"
	return &typeScraper
}

func MigrateLinkedinJob(db *gorm.DB) error {
	err := db.AutoMigrate(&LinkedinJob{})
	return err
}

func (l *LinkedinJob) Save() error {
	return nil
}
