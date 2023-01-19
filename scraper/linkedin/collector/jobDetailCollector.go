package collector

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/dicapisar/job_scraper/scraper/linkedin/result"
	"github.com/gocolly/colly/v2"
	"golang.org/x/net/html"
	"strings"
)

const urlDetailJob = "https://www.linkedin.com/jobs-guest/jobs/api/jobPosting/%s"

type JobDetailCollector struct {
	collector *colly.Collector
}

func (j *JobDetailCollector) GetDetailJob(infoCollectorResult *result.JobInfoCollectorResult) *result.JobDetailCollectorResult {

	jobDetail := result.JobDetailCollectorResult{}
	initialFillJobDetailFromJobInfoCollectorResult(&jobDetail, infoCollectorResult)
	j.initializeNewJobDetailScraper(&jobDetail)

	url := generateUrlJobDetail(infoCollectorResult)

	err := j.collector.Visit(url)

	if err != nil {
		fmt.Printf("Error Visiting Url: %s \n", url)
	}

	return &jobDetail
}

func (j *JobDetailCollector) initializeNewJobDetailScraper(jobDetail *result.JobDetailCollectorResult) {
	c := colly.NewCollector(colly.AllowedDomains("www.linkedin.com", "linkedin.com"))
	j.collector = c

	j.collector.OnHTML("body", func(h *colly.HTMLElement) {
		jobDetail.Description = getDescriptionFromHTMLElement(h)
		jobDetail.Company = getCompanyFromHTMLElement(h)
		jobDetail.SeniorityLevel = getSeniorityLevel(h)
		jobDetail.EmploymentType = getEmploymentType(h)
		jobDetail.JobFunction = getJobFunction(h)
		jobDetail.Industries = getIndustries(h)
		jobDetail.Location = getLocation(h)
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

func initialFillJobDetailFromJobInfoCollectorResult(jobDetail *result.JobDetailCollectorResult, result *result.JobInfoCollectorResult) {
	jobDetail.Title = result.Title
	jobDetail.Id = result.Id
	jobDetail.DateAgo = result.DateAgo
	jobDetail.Url = result.Url
}

func generateUrlJobDetail(result *result.JobInfoCollectorResult) string {
	return fmt.Sprintf(urlDetailJob, result.Id)
}

func getDescriptionFromHTMLElement(h *colly.HTMLElement) string {
	selection := h.DOM
	descriptionSelection := selection.Find("div.show-more-less-html__markup")
	descriptionNodes := descriptionSelection.Nodes[0].FirstChild
	descriptionText := generateTextFromNode(descriptionNodes, h, descriptionSelection)

	return descriptionText
}

func generateTextFromNode(n *html.Node, h *colly.HTMLElement, s *goquery.Selection) string {

	if n == nil {
		return ""
	}

	if n.Type == html.TextNode {
		return fmt.Sprintf("%s \n %v", strings.TrimSpace(n.Data), generateTextFromNode(n.NextSibling, h, s))
	}

	htmlElement := getTextFromNode(n, h, s, 0) //colly.NewHTMLElementFromSelectionNode(h.Response, s, n, 0).Text

	return fmt.Sprintf("%s \n %v", htmlElement, generateTextFromNode(n.NextSibling, h, s))

}

func getTextFromNode(node *html.Node, h *colly.HTMLElement, s *goquery.Selection, i int) string {

	if node != nil {
		switch node.Data {
		case "p":
			htmlElement := colly.NewHTMLElementFromSelectionNode(h.Response, s, node, i)
			return strings.TrimSpace(htmlElement.Text) + "\n"
		case "ul":
			liText := getTextFromNode(node.FirstChild, h, s, i)
			return liText + "\n"
		case "li":
			htmlElement := colly.NewHTMLElementFromSelectionNode(h.Response, s, node, i)
			return strings.TrimSpace(htmlElement.Text) + "\n" + getTextFromNode(node.NextSibling, h, s, i)
		case "strong":
			htmlElement := colly.NewHTMLElementFromSelectionNode(h.Response, s, node, i)
			return strings.TrimSpace(htmlElement.Text) + "\n"
		case "br":
			return ""
		case "em":
			htmlElement := colly.NewHTMLElementFromSelectionNode(h.Response, s, node, i)
			return strings.TrimSpace(htmlElement.Text) + "\n"
		default:
			fmt.Printf("out-of-schema node %v \n", node)
		}
	}

	return ""
}

func getCompanyFromHTMLElement(h *colly.HTMLElement) string {
	selection := h.DOM
	nameCompany := selection.Find("a.topcard__org-name-link").Text()
	nameCompany = strings.TrimSpace(nameCompany)
	return nameCompany
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

func getLocation(h *colly.HTMLElement) string {
	selection := h.DOM
	location := selection.Find("span.topcard__flavor.topcard__flavor--bullet").Text()
	location = strings.TrimSpace(location)
	return location
}
