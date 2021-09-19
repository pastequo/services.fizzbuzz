// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/pastequo/libs.golang.utils/gitutil"
	"github.com/pastequo/libs.golang.utils/logutil"
	"github.com/pastequo/libs.golang.utils/panicrecover"
	"github.com/pastequo/libs.golang.utils/prometrics"
	"github.com/pastequo/libs.golang.utils/queryid"

	"github.com/pastequo/services.fizzbuzz/internal/conf"
	"github.com/pastequo/services.fizzbuzz/internal/handler"
	"github.com/pastequo/services.fizzbuzz/internal/repo/memory"
	"github.com/pastequo/services.fizzbuzz/internal/usecase"
	"github.com/pastequo/services.fizzbuzz/restapi/operations"
)

//go:generate swagger generate server --target ../../services.fizzbuzz --name Fizzbuzz --spec ../target/swagger.yaml --principal interface{}

type configurationFlags struct {
	ConfFile string `short:"c" long:"conf" description:"Path to configuration file" value-name:"FILE"`
}

var confFlags configurationFlags

func configureFlags(api *operations.FizzbuzzAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "Configuration Options",
			Options:          &confFlags,
		},
	}
}

func configureAPI(api *operations.FizzbuzzAPI) http.Handler {
	configureGlobal()

	api.Logger = logutil.GetMethodLogger("").Printf
	api.ServeError = errors.ServeError
	api.UseSwaggerUI()

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	repo := memory.NewFizzBuzzRepo()
	uc := usecase.NewFizzBuzzComp(repo)

	api.FizzbuzzHandler = operations.FizzbuzzHandlerFunc(handler.GetFizzBuzzHandler(uc))
	api.StatsHandler = operations.StatsHandlerFunc(handler.GetStatsHandler(uc))

	api.HealthcheckHandler = operations.HealthcheckHandlerFunc(func(params operations.HealthcheckParams) middleware.Responder {
		return operations.NewHealthcheckOK()
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {
		logger := logutil.GetDefaultLogger()

		logger.Debug("Calling shutdown")

		ctx, cancel := context.WithTimeout(context.Background(), conf.GetGracefulShutdown())
		defer cancel()

		err := prometrics.Shutdown(ctx)
		if err != nil {
			logger.WithError(err).Warn("Failed to shutdown server")
		}
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialised but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	logger := logutil.GetDefaultLogger()

	metricsMiddleware, err := prometrics.GetMiddleware("fizzbuzz", os.Getenv("hostname"), handler)
	if err != nil {
		logger.WithError(err).Warn("failed to create prometheus middleware")

		return queryid.GetMiddleware(panicrecover.GetMiddleware(handler))
	}

	return queryid.GetMiddleware(panicrecover.GetMiddleware(metricsMiddleware))
}

func configureGlobal() {
	// Set log output
	logutil.SetOutput(os.Stdout)
	logutil.SetFormatter(logutil.Text)

	// Dump version information
	logger := logutil.GetDefaultLogger()
	logger.Infof("git version: %v", gitutil.CommitID)

	// Parse configuration
	conf.ParseConfiguration(confFlags.ConfFile)

	// Set log level according to configuration
	level := conf.GetLogsLevel()
	logutil.SetLevel(level)

	// nolint:gomnd
	// Start metrics endpoint
	prometrics.StartServer(7777)
}
