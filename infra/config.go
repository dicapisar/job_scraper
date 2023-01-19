package infra

import (
	"github.com/dicapisar/job_scraper/cron"
	"github.com/dicapisar/job_scraper/repository"
)

var (
	DBRepository *repository.Repository
	Cron         *cron.Cron
)
