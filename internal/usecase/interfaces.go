package usecase

import (
	"context"

	"github.com/pastequo/services.fizzbuzz/internal/entity"
)

type FizzBuzzComputer interface {
	ComputeFizzBuzz(context.Context, entity.FizzBuzzParams) (string, error)
	GetFizzBuzzStats(context.Context) (uint, *entity.FizzBuzzParams, error)
}
