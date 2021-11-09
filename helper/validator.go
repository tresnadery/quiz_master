package helper

import (
	"errors"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni           *ut.UniversalTranslator
	errorMessages = ""
)

func Validate(i interface{}) error {
	newValidator := validator.New()
	err := newValidator.Struct(i)
	if err != nil {
		en := en.New()
		uni = ut.New(en, en)
		trans, _ := uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(newValidator, trans)
		customMessage(trans, newValidator)
		errs := err.(validator.ValidationErrors)
		for i, e := range errs {
			errorMessages += e.Translate(trans)
			if i != len(errs)-1 {
				errorMessages += "\n"
			}
		}
		return errors.New(errorMessages)
	}
	return nil
}

func customMessage(trans ut.Translator, newValidator *validator.Validate) {
	newValidator.RegisterTranslation("required_if", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})
}
