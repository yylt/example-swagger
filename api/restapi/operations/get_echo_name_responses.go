// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetEchoNameOKCode is the HTTP code returned for type GetEchoNameOK
const GetEchoNameOKCode int = 200

/*GetEchoNameOK get echo name o k

swagger:response getEchoNameOK
*/
type GetEchoNameOK struct {
}

// NewGetEchoNameOK creates GetEchoNameOK with default headers values
func NewGetEchoNameOK() *GetEchoNameOK {

	return &GetEchoNameOK{}
}

// WriteResponse to the client
func (o *GetEchoNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetEchoNameBadRequestCode is the HTTP code returned for type GetEchoNameBadRequest
const GetEchoNameBadRequestCode int = 400

/*GetEchoNameBadRequest get echo name bad request

swagger:response getEchoNameBadRequest
*/
type GetEchoNameBadRequest struct {
}

// NewGetEchoNameBadRequest creates GetEchoNameBadRequest with default headers values
func NewGetEchoNameBadRequest() *GetEchoNameBadRequest {

	return &GetEchoNameBadRequest{}
}

// WriteResponse to the client
func (o *GetEchoNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
