package entity

import (
	"github.com/go-playground/validator/v10"

	"github.com/pastequo/services.fizzbuzz/internal/common"
)

// Structs

// FizzBuzzWord is a word for the fizzbuzz algorithm.
type FizzBuzzWord struct {
	Word     string
	Multiple int `validate:"required,gte=1"`
}

// FizzBuzzParams gathers all fizzbuzz algorithm parameters.
type FizzBuzzParams struct {
	FirstWord  FizzBuzzWord
	SecondWord FizzBuzzWord
	Max        int `validate:"required,gte=1"`
}

// Validator

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Validate implements the entity valdiation.
func (f FizzBuzzWord) Validate() error {
	if err := validate.Struct(f); err != nil {
		return common.NewErrInvalidEntity("FizzBuzzWord", err.Error())
	}

	return nil
}

// Validate implements the entity valdiation.
func (f FizzBuzzParams) Validate() error {
	if err := validate.Struct(f); err != nil {
		return common.NewErrInvalidEntity("FizzBuzzParams", err.Error())
	}

	return nil
}
