package Transformers

import (
	"github.com/Jeffail/gabs/v2"
	"gopkg.in/go-playground/validator.v9"
)

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func ParsingValidationError(err error) (interface{}){
	validationErrors := err.(validator.ValidationErrors)
	validationErrFriendlyJson := gabs.New()
	for _, v := range validationErrors {
		error := ValidationError{v.Field(), v.Tag()}
		validationErrFriendlyJson.ArrayAppend(error,"errors")
	}
	return validationErrFriendlyJson.Path("errors").Data()
}
