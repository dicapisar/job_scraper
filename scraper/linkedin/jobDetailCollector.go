package linkedin

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"golang.org/x/net/html"
	"strings"
)

const urlDetailJob = "https://www.linkedin.com/jobs-guest/jobs/api/jobPosting/%s"

type jobDetailCollector struct {
	collector *colly.Collector
}

func (j *jobDetailCollector) GetDetailJob(result *JobInfoCollectorResult) *JobDetailCollectorResult {

	jobDetail := JobDetailCollectorResult{}
	initialFillJobDetailFromJobInfoCollectorResult(&jobDetail, result)
	j.initializeNewJobDetailScraper(&jobDetail)

	url := generateUrlJobDetail(result)

	err := j.collector.Visit(url)

	if err != nil {
		fmt.Printf("Error Visiting Url: %s \n", url)
	}

	return &jobDetail
}

func (j *jobDetailCollector) initializeNewJobDetailScraper(jobDetail *JobDetailCollectorResult) {
	c := colly.NewCollector(colly.AllowedDomains("www.linkedin.com", "linkedin.com"))
	j.collector = c

	j.collector.OnHTML("body", func(h *colly.HTMLElement) {
		jobDetail.Description = getDescriptionFromHTMLElement(h)
		jobDetail.Company = getCompanyFromHTMLElement(h)
		jobDetail.SeniorityLevel = getSeniorityLevel(h)
		jobDetail.EmploymentType = getEmploymentType(h)
		jobDetail.JobFunction = getJobFunction(h)
		jobDetail.Industries = getIndustries(h)
	})

	j.collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
		r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
		fmt.Printf("Visiting %s\n", r.URL)
	})

	j.collector.OnError(func(r *colly.Response, e error) {
		fmt.Printf("Error while scraping: %s \n", e.Error())
	})
}

func initialFillJobDetailFromJobInfoCollectorResult(jobDetail *JobDetailCollectorResult, result *JobInfoCollectorResult) {
	jobDetail.Title = result.Title
	jobDetail.Id = result.Id
	jobDetail.DateAgo = result.DateAgo
	jobDetail.Url = result.Url
}

func generateUrlJobDetail(result *JobInfoCollectorResult) string {
	return fmt.Sprintf(urlDetailJob, result.Id)
}

func getDescriptionFromHTMLElement(h *colly.HTMLElement) string {
	selection := h.DOM
	descriptionSelection := selection.Find("div.show-more-less-html__markup")
	descriptionText := strings.TrimSpace(descriptionSelection.Nodes[0].FirstChild.Data) + fmt.Sprintln()
	description := descriptionSelection.Children()
	i := 0
	description.Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			descriptionText = descriptionText + getTextFromNode(n, h, s, i)
			i = i + 1
		}
	})

	return descriptionText
}

func getTextFromNode(node *html.Node, h *colly.HTMLElement, s *goquery.Selection, i int) string {

	switch node.Data {
	case "p":
		htmlElement := colly.NewHTMLElementFromSelectionNode(h.Response, s, node, i)
		return htmlElement.Text + fmt.Sprintln()
	case "ul":
		stringUl := ""
		htmlElement := colly.NewHTMLElementFromSelectionNode(h.Response, s, node, i)
		list := htmlElement.DOM.Children()
		i := 0
		list.Each(func(_ int, s *goquery.Selection) {
			for _, n := range s.Nodes {
				text := getTextFromNode(n, h, s, i)
				stringUl = stringUl + fmt.Sprintln() + text
			}
		})
		return stringUl + fmt.Sprintln()
	case "li":
		htmlElement := colly.NewHTMLElementFromSelectionNode(h.Response, s, node, i)
		return htmlElement.Text
	case "strong":
		htmlElement := colly.NewHTMLElementFromSelectionNode(h.Response, s, node, i)
		return htmlElement.Text + fmt.Sprintln()
	case "br":
		return ""
	case "em":
		htmlElement := colly.NewHTMLElementFromSelectionNode(h.Response, s, node, i)
		return htmlElement.Text + fmt.Sprintln()
	default:
		fmt.Printf("out-of-schema node %v \n", node)
	}

	return ""
}

func getCompanyFromHTMLElement(h *colly.HTMLElement) string {
	selection := h.DOM
	nameCompany := selection.Find("a.topcard__org-name-link").Text()
	nameCompany = strings.TrimSpace(nameCompany)
	return nameCompany
}
func getCountApplicantsFromHTMLElement(h *colly.HTMLElement) string {
	return ""
}
func getSeniorityLevel(h *colly.HTMLElement) string {
	return *getStringFromDescriptionJobCriteriaListByIndex(h, 0)
}
func getEmploymentType(h *colly.HTMLElement) string {
	return *getStringFromDescriptionJobCriteriaListByIndex(h, 1)
}
func getJobFunction(h *colly.HTMLElement) string {
	return *getStringFromDescriptionJobCriteriaListByIndex(h, 2)
}
func getIndustries(h *colly.HTMLElement) string {
	return *getStringFromDescriptionJobCriteriaListByIndex(h, 3)
}

func getStringFromDescriptionJobCriteriaListByIndex(h *colly.HTMLElement, index int) *string {
	selection := h.DOM
	jobCriteriaNodeList := selection.Find("ul.description__job-criteria-list").Children().Nodes
	textHtmlElement := ""
	if len(jobCriteriaNodeList)-1 >= index {
		node := jobCriteriaNodeList[index]
		htmlElement := colly.NewHTMLElementFromSelectionNode(h.Response, selection, node, index)
		resultSelection := htmlElement.ChildTexts("span.description__job-criteria-text")
		textHtmlElement = strings.TrimSpace(resultSelection[index])
	}

	return &textHtmlElement
}
