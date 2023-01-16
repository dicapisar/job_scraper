package linkedin

import (
	"fmt"
	"github.com/dicapisar/job_scraper/domain"
	"github.com/dicapisar/job_scraper/util"
	"github.com/gocolly/colly/v2"
	"math"
	"strings"
)

const (
	urlListJobs    = "https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=%s&position=1&pageNum=0&start=%d"
	countPageIndex = 25
)

type listJobCollector struct {
	collector *colly.Collector
}

func (l *listJobCollector) GetJobList(search *domain.JobSearch) *[]JobInfoCollectorResult {

	jobInfoCollector := make([]JobInfoCollectorResult, 0, 1)
	l.initializeNewListJobScraper(&jobInfoCollector)

	countSearch := math.Trunc(float64(search.CountToFind / countPageIndex))

	for i := 0; i <= int(countSearch); i++ {
		pageStartIndex := 25

		url := generateUrlListJob(search, uint8(pageStartIndex))

		err := l.collector.Visit(url)

		if err != nil {
			fmt.Printf("Error Visiting Url: %s", url)
		}

		pageStartIndex = pageStartIndex + 25

	}

	return &jobInfoCollector
}

func (l *listJobCollector) initializeNewListJobScraper(jobInfoCollectorList *[]JobInfoCollectorResult) {

	jobInfoCollectorFound := JobInfoCollectorResult{}

	c := colly.NewCollector(colly.AllowedDomains("www.linkedin.com", "linkedin.com"))

	l.collector = c

	l.collector.OnHTML("div.base-card", func(h *colly.HTMLElement) {
		jobInfoCollectorFound.Url = getUrlFromHTMLElement(h)
		jobInfoCollectorFound.Title = getJobTitleFromHTMLElement(h)
		jobInfoCollectorFound.Id = *util.GetInfoJobId(&jobInfoCollectorFound.Url)
		jobInfoCollectorFound.DateAgo = getDateAgoFromHTMLElement(h)

		*jobInfoCollectorList = append(*jobInfoCollectorList, jobInfoCollectorFound)
		jobInfoCollectorFound = JobInfoCollectorResult{}
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
		r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
		fmt.Printf("Visiting %s\n", r.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while scraping: %s\n", e.Error())
	})

}

func getUrlFromHTMLElement(h *colly.HTMLElement) string {
	selection := h.DOM
	jobUrl, _ := selection.Find("a.base-card__full-link").Attr("href")
	return jobUrl
}

func getJobTitleFromHTMLElement(h *colly.HTMLElement) string {
	selection := h.DOM
	jobTitle := selection.Find("h3.base-search-card__title").Text()
	jobTitle = strings.TrimSpace(jobTitle)
	return jobTitle
}

func getDateAgoFromHTMLElement(h *colly.HTMLElement) string {
	selection := h.DOM
	dateAgo, _ := selection.Find("time.job-search-card__listdate").Attr("datetime")
	dateAgo = strings.TrimSpace(dateAgo)
	return dateAgo
}

func generateUrlListJob(search *domain.JobSearch, pageStartIndex uint8) string {
	keyword := search.GetKeyword()
	return fmt.Sprintf(urlListJobs, keyword, pageStartIndex)
}
