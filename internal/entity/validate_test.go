package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pastequo/services.fizzbuzz/internal/entity"
)

func TestNominal(t *testing.T) {
	params := entity.FizzBuzzParams{
		Max:        7,
		FirstWord:  entity.FizzBuzzWord{Word: "fizz", Multiple: 2},
		SecondWord: entity.FizzBuzzWord{Word: "buzz", Multiple: 3},
	}

	err := params.Validate()
	assert.Nil(t, err)
}

func TestMissingWord(t *testing.T) {
	word := entity.FizzBuzzWord{Word: "plouf", Multiple: 4}

	params := entity.FizzBuzzParams{Max: 42} //nolint: exhaustivestruct
	assert.NotNil(t, params.Validate())

	paramsOnly1 := entity.FizzBuzzParams{Max: 42, FirstWord: word} //nolint: exhaustivestruct
	assert.NotNil(t, paramsOnly1.Validate())

	paramsOnly2 := entity.FizzBuzzParams{Max: 42, SecondWord: word} //nolint: exhaustivestruct
	assert.NotNil(t, paramsOnly2.Validate())
}

func TestNegativeMultiple(t *testing.T) {
	params := entity.FizzBuzzParams{
		Max:        7,
		FirstWord:  entity.FizzBuzzWord{Word: "neg", Multiple: -2},
		SecondWord: entity.FizzBuzzWord{Word: "pos", Multiple: 3},
	}

	err := params.Validate()
	assert.NotNil(t, err)

	params = entity.FizzBuzzParams{
		Max:        7,
		SecondWord: entity.FizzBuzzWord{Word: "neg", Multiple: -2},
		FirstWord:  entity.FizzBuzzWord{Word: "pos", Multiple: 3},
	}

	err = params.Validate()
	assert.NotNil(t, err)
}

func TestZeroMultiple(t *testing.T) {
	params := entity.FizzBuzzParams{
		Max:        7,
		FirstWord:  entity.FizzBuzzWord{Word: "zero", Multiple: 0},
		SecondWord: entity.FizzBuzzWord{Word: "pos", Multiple: 3},
	}

	err := params.Validate()
	assert.NotNil(t, err)

	params = entity.FizzBuzzParams{
		Max:        7,
		SecondWord: entity.FizzBuzzWord{Word: "zero", Multiple: 0},
		FirstWord:  entity.FizzBuzzWord{Word: "pos", Multiple: 3},
	}

	err = params.Validate()
	assert.NotNil(t, err)
}

func TestZeroLimit(t *testing.T) {
	params := entity.FizzBuzzParams{
		Max:        0,
		FirstWord:  entity.FizzBuzzWord{Word: "fizz", Multiple: 2},
		SecondWord: entity.FizzBuzzWord{Word: "buzz", Multiple: 3},
	}

	err := params.Validate()
	assert.NotNil(t, err)
}

func TestNegativeLimit(t *testing.T) {
	params := entity.FizzBuzzParams{
		Max:        -7,
		FirstWord:  entity.FizzBuzzWord{Word: "fizz", Multiple: 2},
		SecondWord: entity.FizzBuzzWord{Word: "buzz", Multiple: 3},
	}

	err := params.Validate()
	assert.NotNil(t, err)
}
