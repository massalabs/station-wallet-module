// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/massalabs/thyra-plugin-wallet/api/server/models"
)

// RestAccountExportFileOKCode is the HTTP code returned for type RestAccountExportFileOK
const RestAccountExportFileOKCode int = 200

/*
RestAccountExportFileOK Download the account file

swagger:response restAccountExportFileOK
*/
type RestAccountExportFileOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewRestAccountExportFileOK creates RestAccountExportFileOK with default headers values
func NewRestAccountExportFileOK() *RestAccountExportFileOK {

	return &RestAccountExportFileOK{}
}

// WithPayload adds the payload to the rest account export file o k response
func (o *RestAccountExportFileOK) WithPayload(payload io.ReadCloser) *RestAccountExportFileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest account export file o k response
func (o *RestAccountExportFileOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestAccountExportFileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// RestAccountExportFileBadRequestCode is the HTTP code returned for type RestAccountExportFileBadRequest
const RestAccountExportFileBadRequestCode int = 400

/*
RestAccountExportFileBadRequest Bad request

swagger:response restAccountExportFileBadRequest
*/
type RestAccountExportFileBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRestAccountExportFileBadRequest creates RestAccountExportFileBadRequest with default headers values
func NewRestAccountExportFileBadRequest() *RestAccountExportFileBadRequest {

	return &RestAccountExportFileBadRequest{}
}

// WithPayload adds the payload to the rest account export file bad request response
func (o *RestAccountExportFileBadRequest) WithPayload(payload *models.Error) *RestAccountExportFileBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest account export file bad request response
func (o *RestAccountExportFileBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestAccountExportFileBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RestAccountExportFileNotFoundCode is the HTTP code returned for type RestAccountExportFileNotFound
const RestAccountExportFileNotFoundCode int = 404

/*
RestAccountExportFileNotFound Not found.

swagger:response restAccountExportFileNotFound
*/
type RestAccountExportFileNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRestAccountExportFileNotFound creates RestAccountExportFileNotFound with default headers values
func NewRestAccountExportFileNotFound() *RestAccountExportFileNotFound {

	return &RestAccountExportFileNotFound{}
}

// WithPayload adds the payload to the rest account export file not found response
func (o *RestAccountExportFileNotFound) WithPayload(payload *models.Error) *RestAccountExportFileNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest account export file not found response
func (o *RestAccountExportFileNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestAccountExportFileNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RestAccountExportFileInternalServerErrorCode is the HTTP code returned for type RestAccountExportFileInternalServerError
const RestAccountExportFileInternalServerErrorCode int = 500

/*
RestAccountExportFileInternalServerError Internal Server Error - The server has encountered a situation it does not know how to handle.

swagger:response restAccountExportFileInternalServerError
*/
type RestAccountExportFileInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRestAccountExportFileInternalServerError creates RestAccountExportFileInternalServerError with default headers values
func NewRestAccountExportFileInternalServerError() *RestAccountExportFileInternalServerError {

	return &RestAccountExportFileInternalServerError{}
}

// WithPayload adds the payload to the rest account export file internal server error response
func (o *RestAccountExportFileInternalServerError) WithPayload(payload *models.Error) *RestAccountExportFileInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the rest account export file internal server error response
func (o *RestAccountExportFileInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RestAccountExportFileInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
