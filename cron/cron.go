package cron

import (
	"fmt"
	"github.com/dicapisar/job_scraper/infra"
	"github.com/robfig/cron/v3"
)

type Cron struct {
	Job *cron.Cron
}

func (c *Cron) AddFunc(spec string, function *func()) error {
	addFunc, err := c.Job.AddFunc(spec, *function)

	if err != nil {
		return err
	}

	fmt.Printf("New Job adding, id job: %v \n", addFunc)

	return nil
}

func (c *Cron) DeleteJob(idJOb uint) {
	c.Job.Remove(cron.EntryID(idJOb))
	fmt.Printf("Job deleted, id job: %v \n", idJOb)
}

func GetCron() *Cron {
	if infra.Cron == nil {
		c := cron.New()
		localCron := Cron{
			Job: c,
		}
		infra.Cron = &localCron
	}
	return infra.Cron
}
