package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pastequo/services.fizzbuzz/internal/entity"
	"github.com/pastequo/services.fizzbuzz/models"
)

func TestEntityToTransport(t *testing.T) {
	params := entity.FizzBuzzParams{
		Max:        7,
		FirstWord:  entity.FizzBuzzWord{Word: "fizz", Multiple: 2},
		SecondWord: entity.FizzBuzzWord{Word: "buzz", Multiple: 3},
	}

	limit := int32(7)

	expected := models.FizzbuzzParams{
		Limit: &limit,
		Word1: &models.FizzbuzzWord{Word: "fizz", Multiple: 2},
		Word2: &models.FizzbuzzWord{Word: "buzz", Multiple: 3},
	}

	ret := mapperFizzBuzzParamsEntityToTransport(&params)

	assert.Equal(t, expected, *ret)
}

func TestTransportToEntity(t *testing.T) {
	expected := entity.FizzBuzzParams{
		Max:        7,
		FirstWord:  entity.FizzBuzzWord{Word: "fizz", Multiple: 2},
		SecondWord: entity.FizzBuzzWord{Word: "buzz", Multiple: 3},
	}

	limit := int32(7)

	params := models.FizzbuzzParams{
		Limit: &limit,
		Word1: &models.FizzbuzzWord{Word: "fizz", Multiple: 2},
		Word2: &models.FizzbuzzWord{Word: "buzz", Multiple: 3},
	}

	ret := mapperFizzBuzzParamsTransportToEntity(params)

	assert.Equal(t, expected, ret)
}
