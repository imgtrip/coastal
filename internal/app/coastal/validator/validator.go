package validator

import (
	"errors"
	"gopkg.in/go-playground/validator.v8"
)

var validate *validator.Validate

func IsOwner(reqUserId uint64, tokenUserId uint64) bool {
	return reqUserId == tokenUserId && tokenUserId != 0
}

func MustOwner(reqUserId uint64, tokenUserId uint64) error {
	if IsOwner(reqUserId, tokenUserId) {
		return nil
	}
	return errors.New("permission denied")
}

func New() *validator.Validate {
	if validate == nil {
		config := &validator.Config{TagName: "validate"}
		validate = validator.New(config)
	}
	return validate
}

func Field(str string, tag string) error {
	return New().Field(str, tag)
}
func Structure(structure interface{}) error {
	return New().Struct(structure)
}
