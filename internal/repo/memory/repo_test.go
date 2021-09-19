package memory_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pastequo/services.fizzbuzz/internal/entity"
	"github.com/pastequo/services.fizzbuzz/internal/repo/memory"
)

func TestMultipleAppend(t *testing.T) {
	repo := memory.NewFizzBuzzRepo()
	ctx := context.Background()

	count, _, err := repo.GetMaxFizzBuzzParam(ctx)
	assert.Equal(t, uint(0), count)
	assert.Nil(t, err)

	err = repo.AddFizzBuzzParam(ctx, makeFizzBuzzParams(1, "a", 2, "b", 3))
	assert.Nil(t, err)

	count, params, err := repo.GetMaxFizzBuzzParam(ctx)
	assert.Equal(t, uint(1), count)
	assert.NotNil(t, params)
	assert.Equal(t, makeFizzBuzzParams(1, "a", 2, "b", 3), *params)
	assert.Nil(t, err)

	err = repo.AddFizzBuzzParam(ctx, makeFizzBuzzParams(42, "aa", 22, "bb", 33))
	assert.Nil(t, err)

	err = repo.AddFizzBuzzParam(ctx, makeFizzBuzzParams(42, "aa", 22, "bb", 33))
	assert.Nil(t, err)

	count, params, err = repo.GetMaxFizzBuzzParam(ctx)
	assert.Equal(t, uint(2), count)
	assert.NotNil(t, params)
	assert.Equal(t, makeFizzBuzzParams(42, "aa", 22, "bb", 33), *params)
	assert.Nil(t, err)
}

func TestConcurrentAppend(t *testing.T) {
	repo := memory.NewFizzBuzzRepo()
	ctx := context.Background()

	toAppend := make(chan entity.FizzBuzzParams, 10)
	errors := make(chan error, 10)

	go func() {
		for params := range toAppend {
			errors <- repo.AddFizzBuzzParam(ctx, params)
		}
	}()

	go func() {
		for params := range toAppend {
			errors <- repo.AddFizzBuzzParam(ctx, params)
		}
	}()

	go func() {
		for params := range toAppend {
			errors <- repo.AddFizzBuzzParam(ctx, params)
		}
	}()

	for i := 0; i < 10; i++ {
		toAppend <- makeFizzBuzzParams(42, "aa", 22, "bb", 33)
	}

	close(toAppend)

	for i := 0; i < 10; i++ {
		err := <-errors
		assert.Nil(t, err)
	}

	close(errors)

	count, params, err := repo.GetMaxFizzBuzzParam(ctx)
	assert.Equal(t, uint(10), count)
	assert.NotNil(t, params)
	assert.Equal(t, makeFizzBuzzParams(42, "aa", 22, "bb", 33), *params)
	assert.Nil(t, err)
}

func makeFizzBuzzParams(max int, w1 string, n1 int, w2 string, n2 int) entity.FizzBuzzParams {
	return entity.FizzBuzzParams{
		Max:        max,
		FirstWord:  entity.FizzBuzzWord{Word: w1, Multiple: n1},
		SecondWord: entity.FizzBuzzWord{Word: w2, Multiple: n2},
	}
}
