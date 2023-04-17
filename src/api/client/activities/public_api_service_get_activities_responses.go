// Code generated by go-swagger; DO NOT EDIT.

package activities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tkhq/tkcli/src/api/models"
)

// PublicAPIServiceGetActivitiesReader is a Reader for the PublicAPIServiceGetActivities structure.
type PublicAPIServiceGetActivitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicAPIServiceGetActivitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicAPIServiceGetActivitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPublicAPIServiceGetActivitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPublicAPIServiceGetActivitiesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPublicAPIServiceGetActivitiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPublicAPIServiceGetActivitiesOK creates a PublicAPIServiceGetActivitiesOK with default headers values
func NewPublicAPIServiceGetActivitiesOK() *PublicAPIServiceGetActivitiesOK {
	return &PublicAPIServiceGetActivitiesOK{}
}

/*
PublicAPIServiceGetActivitiesOK describes a response with status code 200, with default header values.

A successful response.
*/
type PublicAPIServiceGetActivitiesOK struct {
	Payload *models.V1GetActivitiesResponse
}

// IsSuccess returns true when this public Api service get activities o k response has a 2xx status code
func (o *PublicAPIServiceGetActivitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this public Api service get activities o k response has a 3xx status code
func (o *PublicAPIServiceGetActivitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this public Api service get activities o k response has a 4xx status code
func (o *PublicAPIServiceGetActivitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this public Api service get activities o k response has a 5xx status code
func (o *PublicAPIServiceGetActivitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this public Api service get activities o k response a status code equal to that given
func (o *PublicAPIServiceGetActivitiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the public Api service get activities o k response
func (o *PublicAPIServiceGetActivitiesOK) Code() int {
	return 200
}

func (o *PublicAPIServiceGetActivitiesOK) Error() string {
	return fmt.Sprintf("[POST /public/v1/query/list_activities][%d] publicApiServiceGetActivitiesOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceGetActivitiesOK) String() string {
	return fmt.Sprintf("[POST /public/v1/query/list_activities][%d] publicApiServiceGetActivitiesOK  %+v", 200, o.Payload)
}

func (o *PublicAPIServiceGetActivitiesOK) GetPayload() *models.V1GetActivitiesResponse {
	return o.Payload
}

func (o *PublicAPIServiceGetActivitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetActivitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicAPIServiceGetActivitiesForbidden creates a PublicAPIServiceGetActivitiesForbidden with default headers values
func NewPublicAPIServiceGetActivitiesForbidden() *PublicAPIServiceGetActivitiesForbidden {
	return &PublicAPIServiceGetActivitiesForbidden{}
}

/*
PublicAPIServiceGetActivitiesForbidden describes a response with status code 403, with default header values.

Returned when the user does not have permission to access the resource.
*/
type PublicAPIServiceGetActivitiesForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this public Api service get activities forbidden response has a 2xx status code
func (o *PublicAPIServiceGetActivitiesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this public Api service get activities forbidden response has a 3xx status code
func (o *PublicAPIServiceGetActivitiesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this public Api service get activities forbidden response has a 4xx status code
func (o *PublicAPIServiceGetActivitiesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this public Api service get activities forbidden response has a 5xx status code
func (o *PublicAPIServiceGetActivitiesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this public Api service get activities forbidden response a status code equal to that given
func (o *PublicAPIServiceGetActivitiesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the public Api service get activities forbidden response
func (o *PublicAPIServiceGetActivitiesForbidden) Code() int {
	return 403
}

func (o *PublicAPIServiceGetActivitiesForbidden) Error() string {
	return fmt.Sprintf("[POST /public/v1/query/list_activities][%d] publicApiServiceGetActivitiesForbidden  %+v", 403, o.Payload)
}

func (o *PublicAPIServiceGetActivitiesForbidden) String() string {
	return fmt.Sprintf("[POST /public/v1/query/list_activities][%d] publicApiServiceGetActivitiesForbidden  %+v", 403, o.Payload)
}

func (o *PublicAPIServiceGetActivitiesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PublicAPIServiceGetActivitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicAPIServiceGetActivitiesNotFound creates a PublicAPIServiceGetActivitiesNotFound with default headers values
func NewPublicAPIServiceGetActivitiesNotFound() *PublicAPIServiceGetActivitiesNotFound {
	return &PublicAPIServiceGetActivitiesNotFound{}
}

/*
PublicAPIServiceGetActivitiesNotFound describes a response with status code 404, with default header values.

Returned when the resource does not exist.
*/
type PublicAPIServiceGetActivitiesNotFound struct {
	Payload string
}

// IsSuccess returns true when this public Api service get activities not found response has a 2xx status code
func (o *PublicAPIServiceGetActivitiesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this public Api service get activities not found response has a 3xx status code
func (o *PublicAPIServiceGetActivitiesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this public Api service get activities not found response has a 4xx status code
func (o *PublicAPIServiceGetActivitiesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this public Api service get activities not found response has a 5xx status code
func (o *PublicAPIServiceGetActivitiesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this public Api service get activities not found response a status code equal to that given
func (o *PublicAPIServiceGetActivitiesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the public Api service get activities not found response
func (o *PublicAPIServiceGetActivitiesNotFound) Code() int {
	return 404
}

func (o *PublicAPIServiceGetActivitiesNotFound) Error() string {
	return fmt.Sprintf("[POST /public/v1/query/list_activities][%d] publicApiServiceGetActivitiesNotFound  %+v", 404, o.Payload)
}

func (o *PublicAPIServiceGetActivitiesNotFound) String() string {
	return fmt.Sprintf("[POST /public/v1/query/list_activities][%d] publicApiServiceGetActivitiesNotFound  %+v", 404, o.Payload)
}

func (o *PublicAPIServiceGetActivitiesNotFound) GetPayload() string {
	return o.Payload
}

func (o *PublicAPIServiceGetActivitiesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicAPIServiceGetActivitiesDefault creates a PublicAPIServiceGetActivitiesDefault with default headers values
func NewPublicAPIServiceGetActivitiesDefault(code int) *PublicAPIServiceGetActivitiesDefault {
	return &PublicAPIServiceGetActivitiesDefault{
		_statusCode: code,
	}
}

/*
PublicAPIServiceGetActivitiesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PublicAPIServiceGetActivitiesDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this public Api service get activities default response has a 2xx status code
func (o *PublicAPIServiceGetActivitiesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this public Api service get activities default response has a 3xx status code
func (o *PublicAPIServiceGetActivitiesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this public Api service get activities default response has a 4xx status code
func (o *PublicAPIServiceGetActivitiesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this public Api service get activities default response has a 5xx status code
func (o *PublicAPIServiceGetActivitiesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this public Api service get activities default response a status code equal to that given
func (o *PublicAPIServiceGetActivitiesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the public Api service get activities default response
func (o *PublicAPIServiceGetActivitiesDefault) Code() int {
	return o._statusCode
}

func (o *PublicAPIServiceGetActivitiesDefault) Error() string {
	return fmt.Sprintf("[POST /public/v1/query/list_activities][%d] PublicApiService_GetActivities default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceGetActivitiesDefault) String() string {
	return fmt.Sprintf("[POST /public/v1/query/list_activities][%d] PublicApiService_GetActivities default  %+v", o._statusCode, o.Payload)
}

func (o *PublicAPIServiceGetActivitiesDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *PublicAPIServiceGetActivitiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
