package api

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/yylt/example-swagger/api/restapi/operations"
)

func (a *Api) getEchoHandler(params operations.GetEchoParams) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

func (a *Api) getEchoNameHandler(params operations.GetEchoNameParams) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

func (a *Api) postEchoHandler(params operations.PostEchoParams) middleware.Responder {
	return middleware.NotImplemented("not impl")
}

func (a *Api) postEchoNameHandler(params operations.PostEchoNameParams) middleware.Responder {
	return middleware.NotImplemented("not impl")
}
