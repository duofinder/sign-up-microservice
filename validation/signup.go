package validation

import (
	"encoding/json"

	validator "github.com/go-playground/validator/v10"
)

func Validate[T any](body string, obj *T) (*T, error) {
	err := json.Unmarshal([]byte(body), obj)
	if err != nil {
		return nil, err
	}

	err = validator.New().Struct(obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
