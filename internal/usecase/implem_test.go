package usecase_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/pastequo/services.fizzbuzz/internal/common"
	"github.com/pastequo/services.fizzbuzz/internal/entity"
	mock_repo "github.com/pastequo/services.fizzbuzz/internal/repo/mock"
	"github.com/pastequo/services.fizzbuzz/internal/usecase"
)

func TestNominal(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_repo.NewMockFizzBuzzParam(ctrl)
	uc := usecase.NewFizzBuzzComp(m)

	m.EXPECT().AddFizzBuzzParam(gomock.Any(), gomock.Any()).Return(nil)

	ctx := context.Background()

	ret, err := uc.ComputeFizzBuzz(ctx, entity.FizzBuzzParams{
		Max:        7,
		FirstWord:  entity.FizzBuzzWord{Word: "fizz", Multiple: 2},
		SecondWord: entity.FizzBuzzWord{Word: "buzz", Multiple: 3},
	})

	assert.Nil(t, err)
	assert.Equal(t, "1,fizz,buzz,fizz,5,fizzbuzz,7", ret)
}

func TestRepoFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_repo.NewMockFizzBuzzParam(ctrl)
	uc := usecase.NewFizzBuzzComp(m)

	m.EXPECT().AddFizzBuzzParam(gomock.Any(), gomock.Any()).Return(common.NewErrInternalErrorMsg("fake failed"))

	ctx := context.Background()

	_, err := uc.ComputeFizzBuzz(ctx, entity.FizzBuzzParams{
		Max:        7,
		FirstWord:  entity.FizzBuzzWord{Word: "fizz", Multiple: 2},
		SecondWord: entity.FizzBuzzWord{Word: "buzz", Multiple: 3},
	})

	assert.NotNil(t, err)
}

func TestInvalidEntity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_repo.NewMockFizzBuzzParam(ctrl)
	uc := usecase.NewFizzBuzzComp(m)

	ctx := context.Background()

	_, err := uc.ComputeFizzBuzz(ctx, entity.FizzBuzzParams{
		Max:        -7,
		FirstWord:  entity.FizzBuzzWord{Word: "fizz", Multiple: 2},
		SecondWord: entity.FizzBuzzWord{Word: "buzz", Multiple: 3},
	})

	assert.NotNil(t, err)
}
