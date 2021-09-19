package handler

import (
	"github.com/pastequo/services.fizzbuzz/internal/entity"
	"github.com/pastequo/services.fizzbuzz/models"
)

func mapperFizzBuzzParamsTransportToEntity(params models.FizzbuzzParams) entity.FizzBuzzParams {
	max := int(*params.Limit)

	mult1 := int(params.Word1.Multiple)
	mult2 := int(params.Word2.Multiple)

	return entity.FizzBuzzParams{
		Max: max,
		FirstWord: entity.FizzBuzzWord{
			Word:     params.Word1.Word,
			Multiple: mult1,
		},
		SecondWord: entity.FizzBuzzWord{
			Word:     params.Word2.Word,
			Multiple: mult2,
		},
	}
}

func mapperFizzBuzzParamsEntityToTransport(params *entity.FizzBuzzParams) *models.FizzbuzzParams {
	if params == nil {
		return nil
	}

	limit := int32(params.Max)

	return &models.FizzbuzzParams{
		Limit: &limit,
		Word1: &models.FizzbuzzWord{
			Multiple: int32(params.FirstWord.Multiple),
			Word:     params.FirstWord.Word,
		},
		Word2: &models.FizzbuzzWord{
			Multiple: int32(params.SecondWord.Multiple),
			Word:     params.SecondWord.Word,
		},
	}
}
