package builder

import (
	"quiz_master/dto"
)

type OptionRequestGetOrDelete func(*dto.RequestGetOrDeleteQuestion)

func NewRequestGetOrDelete(options ...OptionRequestGetOrDelete) *dto.RequestGetOrDeleteQuestion {
	r := &dto.RequestGetOrDeleteQuestion{}
	for _, o := range options {
		o(r)
	}
	return r
}

func GetOrDeleteWithNumber(number string) OptionRequestGetOrDelete {
	return func(r *dto.RequestGetOrDeleteQuestion) {
		r.Number = number
	}
}
