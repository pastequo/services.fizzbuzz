package handler

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"

	"github.com/pastequo/services.fizzbuzz/internal/common"
	"github.com/pastequo/services.fizzbuzz/internal/conf"
	"github.com/pastequo/services.fizzbuzz/internal/usecase"
	"github.com/pastequo/services.fizzbuzz/models"
	"github.com/pastequo/services.fizzbuzz/restapi/operations"
)

func GetFizzBuzzHandler(fb usecase.FizzBuzzComputer) func(operations.FizzbuzzParams) middleware.Responder {
	return operations.FizzbuzzHandlerFunc(func(params operations.FizzbuzzParams) middleware.Responder {
		// Ensure that the context has a deadline.
		ctx, cancel := context.WithTimeout(params.HTTPRequest.Context(), conf.GetGlobalTimeout())
		defer cancel()

		// Params validation is made by go-swagger.

		// Convert params into entity.
		fbParams := mapperFizzBuzzParamsTransportToEntity(*params.FizzbuzzParams)

		// Call usecase.
		ret, err := fb.ComputeFizzBuzz(ctx, fbParams)
		if err != nil {
			switch {
			case errors.As(err, &common.ErrInvalidEntity{}):
				tmp := models.ErrorMsgTypeErrInvalidObject

				return operations.NewFizzbuzzBadRequest().WithPayload(&models.ErrorMsg{
					Type:    &tmp,
					Message: err.Error(),
				})
			case errors.As(err, &common.ErrInternalError{}):
				return operations.NewFizzbuzzInternalServerError()
			default:
				return operations.NewFizzbuzzInternalServerError()
			}
		}

		if len(ret) == 0 {
			return operations.NewFizzbuzzNoContent()
		}

		return operations.NewFizzbuzzOK().WithPayload(ret)
	})
}

func GetStatsHandler(fb usecase.FizzBuzzComputer) func(params operations.StatsParams) middleware.Responder {
	return operations.StatsHandlerFunc(func(params operations.StatsParams) middleware.Responder {
		// Ensure that the context has a deadline.
		ctx, cancel := context.WithTimeout(params.HTTPRequest.Context(), conf.GetGlobalTimeout())
		defer cancel()

		max, fbParams, err := fb.GetFizzBuzzStats(ctx)
		if err != nil {
			return operations.NewFizzbuzzInternalServerError()
		}

		count := int32(max)

		return operations.NewStatsOK().WithPayload(&operations.StatsOKBody{
			Count:     &count,
			Parameter: mapperFizzBuzzParamsEntityToTransport(fbParams),
		})
	})
}
