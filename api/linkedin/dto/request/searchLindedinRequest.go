package request

import (
	"github.com/dicapisar/job_scraper/api/share"
	"github.com/go-playground/validator/v10"
)

type Search struct {
	Title       string `json:"title" validate:"require, min=3, max=32"`
	CountToFind int    `json:"countToFind" validate:"require"`
	Location    string `json:"location" validate:"require, min=3, max=32"`
}

func (s *Search) Validate() []*share.ErrorResponse {
	var validate = validator.New()

	var errors []*share.ErrorResponse

	err := validate.Struct(&s)

	if errors != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element share.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
