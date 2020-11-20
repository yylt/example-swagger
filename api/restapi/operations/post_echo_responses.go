// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostEchoOKCode is the HTTP code returned for type PostEchoOK
const PostEchoOKCode int = 200

/*PostEchoOK post echo o k

swagger:response postEchoOK
*/
type PostEchoOK struct {
}

// NewPostEchoOK creates PostEchoOK with default headers values
func NewPostEchoOK() *PostEchoOK {

	return &PostEchoOK{}
}

// WriteResponse to the client
func (o *PostEchoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostEchoBadRequestCode is the HTTP code returned for type PostEchoBadRequest
const PostEchoBadRequestCode int = 400

/*PostEchoBadRequest post echo bad request

swagger:response postEchoBadRequest
*/
type PostEchoBadRequest struct {
}

// NewPostEchoBadRequest creates PostEchoBadRequest with default headers values
func NewPostEchoBadRequest() *PostEchoBadRequest {

	return &PostEchoBadRequest{}
}

// WriteResponse to the client
func (o *PostEchoBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}