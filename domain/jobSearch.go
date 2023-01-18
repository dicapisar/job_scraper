package domain

import (
	"fmt"
	"github.com/dicapisar/job_scraper/api/linkedin/dto/request"
	"strings"
)

type JobSearch struct {
	Title       string
	CountToFind int
	Location    string
}

func (j *JobSearch) GetKeyword() string {
	return strings.ReplaceAll(j.Title, " ", "+")
}

func (j *JobSearch) GetLocation() string {
	if j.Location == "" {
		return ""
	}
	return fmt.Sprintf("&location=%s", j.Location)
}

func (j *JobSearch) ParseFromLinkedinSearch(search *request.Search) {
	j.Title = search.Title
	j.CountToFind = search.CountToFind
	j.Location = search.Location
}
