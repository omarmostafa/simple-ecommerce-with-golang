package infrastructures

import "gopkg.in/go-playground/validator.v9"

func Validate(req interface{}) (error) {
	v := validator.New()
	err := v.Struct(req)
	if err!=nil {
		return err
	}
	return  nil
}

