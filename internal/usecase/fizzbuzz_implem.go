package usecase

import (
	"context"
	"fmt"
	"sync"

	"github.com/pastequo/libs.golang.utils/logutil"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/pastequo/services.fizzbuzz/internal/algo"
	"github.com/pastequo/services.fizzbuzz/internal/common"
	"github.com/pastequo/services.fizzbuzz/internal/entity"
	"github.com/pastequo/services.fizzbuzz/internal/repo"
)

var (
	initOnce     sync.Once
	queryCounter *prometheus.CounterVec
)

type FizzBuzzComp struct {
	rep repo.FizzBuzzParam
}

func NewFizzBuzzComp(rep repo.FizzBuzzParam) FizzBuzzComp {
	initOnce.Do(func() {
		queryCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace:   "algo",
			Subsystem:   "fizzbuzz",
			Name:        "query_total",
			Help:        "count how many times each request has been sent.",
			ConstLabels: nil,
		}, []string{"limit", "word1", "multiple1", "word2", "multiple2"})

		err := prometheus.Register(queryCounter)
		if err != nil {
			logutil.GetDefaultLogger().Warn("Failed to register query counter metric")
		}
	})

	return FizzBuzzComp{rep: rep}
}

func (f FizzBuzzComp) ComputeFizzBuzz(ctx context.Context, params entity.FizzBuzzParams) (string, error) {
	logger := logutil.GetLogger(ctx)

	err := params.Validate()
	if err != nil {
		logger.WithError(err).Errorf("Invalid parameters: %+v", params)

		return "", err
	}

	err = f.rep.AddFizzBuzzParam(ctx, params)
	if err != nil {
		logger.WithError(err).Errorf("Failed to count params: %+v", params)

		return "", err
	}

	queryCounter.WithLabelValues(getMetricValues(params)...).Inc()

	return algo.ComputeFizzBuzzString(params), nil
}

func (f FizzBuzzComp) GetFizzBuzzStats(ctx context.Context) (uint, *entity.FizzBuzzParams, error) {
	logger := logutil.GetLogger(ctx)

	max, params, err := f.rep.GetMaxFizzBuzzParam(ctx)
	if err != nil {
		logger.WithError(err).Errorf("Failed to get most used params")

		return 0, nil, err
	}

	if params == nil && max > 0 {
		logger.Warnf("Inconsistent result from repo: params is nil & max %d", max)

		return 0, nil, common.NewErrInternalErrorMsg("Inconsistent result")
	}

	if params != nil {
		err = params.Validate()
		if err != nil {
			logger.WithError(err).Errorf("got invalid entity: %+v", *params)

			return 0, nil, err
		}
	}

	return max, params, nil
}

func getMetricValues(params entity.FizzBuzzParams) []string {
	return []string{
		fmt.Sprintf("%d", params.Max),
		params.FirstWord.Word,
		fmt.Sprintf("%d", params.FirstWord.Multiple),
		params.SecondWord.Word,
		fmt.Sprintf("%d", params.SecondWord.Multiple),
	}
}
