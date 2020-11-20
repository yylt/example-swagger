// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostEchoNameOKCode is the HTTP code returned for type PostEchoNameOK
const PostEchoNameOKCode int = 200

/*PostEchoNameOK post echo name o k

swagger:response postEchoNameOK
*/
type PostEchoNameOK struct {
}

// NewPostEchoNameOK creates PostEchoNameOK with default headers values
func NewPostEchoNameOK() *PostEchoNameOK {

	return &PostEchoNameOK{}
}

// WriteResponse to the client
func (o *PostEchoNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostEchoNameBadRequestCode is the HTTP code returned for type PostEchoNameBadRequest
const PostEchoNameBadRequestCode int = 400

/*PostEchoNameBadRequest post echo name bad request

swagger:response postEchoNameBadRequest
*/
type PostEchoNameBadRequest struct {
}

// NewPostEchoNameBadRequest creates PostEchoNameBadRequest with default headers values
func NewPostEchoNameBadRequest() *PostEchoNameBadRequest {

	return &PostEchoNameBadRequest{}
}

// WriteResponse to the client
func (o *PostEchoNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
