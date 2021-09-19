package memory

import (
	"context"
	"sync"

	"github.com/pastequo/libs.golang.utils/logutil"

	"github.com/pastequo/services.fizzbuzz/internal/entity"
)

type FizzBuzzRepo struct {
	lock     sync.RWMutex
	requests map[entity.FizzBuzzParams]uint
}

func NewFizzBuzzRepo() *FizzBuzzRepo {
	return &FizzBuzzRepo{
		lock:     sync.RWMutex{},
		requests: make(map[entity.FizzBuzzParams]uint),
	}
}

func (f *FizzBuzzRepo) AddFizzBuzzParam(ctx context.Context, params entity.FizzBuzzParams) error {
	f.lock.Lock()
	defer f.lock.Unlock()

	f.requests[params]++

	logutil.GetLogger(ctx).Debugf("requests: %+v", f.requests)

	return nil
}

func (f *FizzBuzzRepo) GetMaxFizzBuzzParam(ctx context.Context) (uint, *entity.FizzBuzzParams, error) {
	f.lock.RLock()
	defer f.lock.RUnlock()

	max := uint(0)

	var ret *entity.FizzBuzzParams

	logutil.GetLogger(ctx).Debugf("requests: %+v", f.requests)

	for params, count := range f.requests {
		if count > max {
			max = count
			tmp := params
			ret = &tmp
		}
	}

	logutil.GetLogger(ctx).Debugf("return: %d", max)

	return max, ret, nil
}
