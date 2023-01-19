package request

import (
	"github.com/dicapisar/job_scraper/api/share"
	"github.com/go-playground/validator/v10"
)

type Search struct {
	Title        string `json:"title" validate:"required,min=3,max=32"`
	CountToFind  int    `json:"countToFind" validate:"required"`
	Location     string `json:"location" validate:"required,min=3,max=32"`
	Interval     uint8  `json:"interval" validate:"required,gte=1,lte=60"`
	TypeInterval string `json:"type_interval" validate:"required"`
}

func (s *Search) Validate() []*share.ErrorResponse {
	var validate = validator.New()

	var errors []*share.ErrorResponse

	errValidate := validate.Struct(s)

	if errValidate != nil {
		for _, err := range errValidate.(validator.ValidationErrors) {
			var element share.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			element.TypeField = err.Kind().String()
			errors = append(errors, &element)
		}
	} else {
		if s.TypeInterval != "hour" && s.TypeInterval != "minute" && s.TypeInterval != "day" {
			var element share.ErrorResponse
			element.FailedField = "type_interval"
			element.Tag = "value entered not allowed, try with 'hour' or 'minute' or 'day'"
			element.Value = s.TypeInterval
			errors = append(errors, &element)
		}
	}
	return errors
}
