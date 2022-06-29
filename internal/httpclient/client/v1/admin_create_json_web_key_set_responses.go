// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/models"
)

// AdminCreateJSONWebKeySetReader is a Reader for the AdminCreateJSONWebKeySet structure.
type AdminCreateJSONWebKeySetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminCreateJSONWebKeySetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAdminCreateJSONWebKeySetCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAdminCreateJSONWebKeySetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAdminCreateJSONWebKeySetCreated creates a AdminCreateJSONWebKeySetCreated with default headers values
func NewAdminCreateJSONWebKeySetCreated() *AdminCreateJSONWebKeySetCreated {
	return &AdminCreateJSONWebKeySetCreated{}
}

/* AdminCreateJSONWebKeySetCreated describes a response with status code 201, with default header values.

jsonWebKeySet
*/
type AdminCreateJSONWebKeySetCreated struct {
	Payload *models.JSONWebKeySet
}

func (o *AdminCreateJSONWebKeySetCreated) Error() string {
	return fmt.Sprintf("[POST /admin/keys/{set}][%d] adminCreateJsonWebKeySetCreated  %+v", 201, o.Payload)
}
func (o *AdminCreateJSONWebKeySetCreated) GetPayload() *models.JSONWebKeySet {
	return o.Payload
}

func (o *AdminCreateJSONWebKeySetCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONWebKeySet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCreateJSONWebKeySetDefault creates a AdminCreateJSONWebKeySetDefault with default headers values
func NewAdminCreateJSONWebKeySetDefault(code int) *AdminCreateJSONWebKeySetDefault {
	return &AdminCreateJSONWebKeySetDefault{
		_statusCode: code,
	}
}

/* AdminCreateJSONWebKeySetDefault describes a response with status code -1, with default header values.

oAuth2ApiError
*/
type AdminCreateJSONWebKeySetDefault struct {
	_statusCode int

	Payload *models.OAuth2APIError
}

// Code gets the status code for the admin create Json web key set default response
func (o *AdminCreateJSONWebKeySetDefault) Code() int {
	return o._statusCode
}

func (o *AdminCreateJSONWebKeySetDefault) Error() string {
	return fmt.Sprintf("[POST /admin/keys/{set}][%d] adminCreateJsonWebKeySet default  %+v", o._statusCode, o.Payload)
}
func (o *AdminCreateJSONWebKeySetDefault) GetPayload() *models.OAuth2APIError {
	return o.Payload
}

func (o *AdminCreateJSONWebKeySetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuth2APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
