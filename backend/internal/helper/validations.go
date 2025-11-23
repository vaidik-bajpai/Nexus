package helper

import "github.com/go-playground/validator/v10"

func RegisterCustomValidations(v *validator.Validate) error {
	validations := []struct {
		tag      string
		validate validator.Func
	}{
		{"color_or_url", colorOrURL},
	}

	for _, val := range validations {
		if err := v.RegisterValidation(val.tag, val.validate); err != nil {
			return err
		}
	}

	return nil
}

func colorOrURL(fl validator.FieldLevel) bool {
	val := fl.Field().String()

	v := validator.New()

	errColor := v.Var(val, "iscolor")
	if errColor == nil {
		return true
	}

	errURL := v.Var(val, "url")
	return errURL == nil
}
