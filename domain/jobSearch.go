package domain

import "strings"

type JobSearch struct {
	Title       string
	CountToFind int
	Location    Location
}

func (j *JobSearch) GetKeyword() string {
	return strings.ReplaceAll(j.Title, " ", "+")
}
