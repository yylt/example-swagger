//go:generate swagger generate model
//go:generate swagger generate operation
//go:generate swagger generate server --exclude-main
package api

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/yylt/example-swagger/api/restapi"
	"github.com/yylt/example-swagger/api/restapi/operations"
	"net/http"
)

type Api struct {
	log     log.Logger
	Handler http.Handler
}

func New(log log.Logger, baseurl string) (*Api, error) {
	api := &Api{log: log}
	// Load embedded swagger file.
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return nil, fmt.Errorf("failed to load embedded swagger file: %v", err.Error())
	}

	swaggerSpec.Spec().BasePath = baseurl

	// Create new service API.
	openAPI := operations.NewExampleAPI(swaggerSpec)

	// Skip the  redoc middleware, only serving the OpenAPI specification and
	// the API itself via RoutesHandler. See:
	// https://github.com/go-swagger/go-swagger/issues/1779
	openAPI.Middleware = func(b middleware.Builder) http.Handler {
		return middleware.Spec(baseurl, swaggerSpec.Raw(), openAPI.Context().RoutesHandler(b))
	}

	openAPI.GetEchoHandler = operations.GetEchoHandlerFunc(api.getEchoHandler)
	openAPI.GetEchoNameHandler = operations.GetEchoNameHandlerFunc(api.getEchoNameHandler)
	openAPI.PostEchoHandler = operations.PostEchoHandlerFunc(api.postEchoHandler)
	openAPI.PostEchoNameHandler = operations.PostEchoNameHandlerFunc(api.postEchoNameHandler)

	api.Handler = openAPI.Serve(nil)

	return api, nil
}
