package domain

type Job interface {
	GetTypeScraper() *string
	Save() error
}
