package algo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pastequo/services.fizzbuzz/internal/algo"
	"github.com/pastequo/services.fizzbuzz/internal/entity"
)

func TestNominal(t *testing.T) {
	expected := "1,fizz,buzz,fizz,5,fizzbuzz,7"

	ret := algo.ComputeFizzBuzzString(entity.FizzBuzzParams{
		Max:        7,
		FirstWord:  entity.FizzBuzzWord{Word: "fizz", Multiple: 2},
		SecondWord: entity.FizzBuzzWord{Word: "buzz", Multiple: 3},
	})

	assert.Equal(t, expected, ret)
}

func TestSameValue(t *testing.T) {
	expected := "1,ploufplouf,3,ploufplouf,5,ploufplouf,7,ploufplouf"

	word := entity.FizzBuzzWord{Word: "plouf", Multiple: 2}

	ret := algo.ComputeFizzBuzzString(entity.FizzBuzzParams{
		Max:        8,
		FirstWord:  word,
		SecondWord: word,
	})

	assert.Equal(t, expected, ret)
}

func TestEmptyWord(t *testing.T) {
	expected := "1,2,,4,fives,,7"

	ret := algo.ComputeFizzBuzzString(entity.FizzBuzzParams{
		Max:        7,
		FirstWord:  entity.FizzBuzzWord{Word: "", Multiple: 3},
		SecondWord: entity.FizzBuzzWord{Word: "fives", Multiple: 5},
	})

	assert.Equal(t, expected, ret)

	expected = ",,,"

	ret = algo.ComputeFizzBuzzString(entity.FizzBuzzParams{
		Max:        4,
		FirstWord:  entity.FizzBuzzWord{Word: "", Multiple: 1},
		SecondWord: entity.FizzBuzzWord{Word: "", Multiple: 5},
	})

	assert.Equal(t, expected, ret)
}

func TestEmptyResult(t *testing.T) {
	expected := ""

	word := entity.FizzBuzzWord{Word: "", Multiple: 1}

	ret := algo.ComputeFizzBuzzString(entity.FizzBuzzParams{
		Max:        1,
		FirstWord:  word,
		SecondWord: word,
	})

	assert.Equal(t, expected, ret)
}
