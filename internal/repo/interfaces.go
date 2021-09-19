package repo

import (
	"context"

	"github.com/pastequo/services.fizzbuzz/internal/entity"
)

type FizzBuzzParam interface {
	AddFizzBuzzParam(ctx context.Context, param entity.FizzBuzzParams) error
	GetMaxFizzBuzzParam(ctx context.Context) (uint, *entity.FizzBuzzParams, error)
}
