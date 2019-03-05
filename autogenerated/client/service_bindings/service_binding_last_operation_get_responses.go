// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/openservicebrokerapi/osb-checker/autogenerated/models"
)

// ServiceBindingLastOperationGetReader is a Reader for the ServiceBindingLastOperationGet structure.
type ServiceBindingLastOperationGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBindingLastOperationGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServiceBindingLastOperationGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewServiceBindingLastOperationGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 410:
		result := NewServiceBindingLastOperationGetGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServiceBindingLastOperationGetOK creates a ServiceBindingLastOperationGetOK with default headers values
func NewServiceBindingLastOperationGetOK() *ServiceBindingLastOperationGetOK {
	return &ServiceBindingLastOperationGetOK{}
}

/*ServiceBindingLastOperationGetOK handles this case with default header values.

OK
*/
type ServiceBindingLastOperationGetOK struct {
	/*Indicates when to retry the request
	 */
	RetryAfter string

	Payload *models.LastOperationResource
}

func (o *ServiceBindingLastOperationGetOK) Error() string {
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetOK  %+v", 200, o.Payload)
}

func (o *ServiceBindingLastOperationGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header RetryAfter
	o.RetryAfter = response.GetHeader("RetryAfter")

	o.Payload = new(models.LastOperationResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingLastOperationGetBadRequest creates a ServiceBindingLastOperationGetBadRequest with default headers values
func NewServiceBindingLastOperationGetBadRequest() *ServiceBindingLastOperationGetBadRequest {
	return &ServiceBindingLastOperationGetBadRequest{}
}

/*ServiceBindingLastOperationGetBadRequest handles this case with default header values.

Bad Request
*/
type ServiceBindingLastOperationGetBadRequest struct {
	Payload *models.Error
}

func (o *ServiceBindingLastOperationGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBindingLastOperationGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingLastOperationGetGone creates a ServiceBindingLastOperationGetGone with default headers values
func NewServiceBindingLastOperationGetGone() *ServiceBindingLastOperationGetGone {
	return &ServiceBindingLastOperationGetGone{}
}

/*ServiceBindingLastOperationGetGone handles this case with default header values.

Gone
*/
type ServiceBindingLastOperationGetGone struct {
	Payload *models.Error
}

func (o *ServiceBindingLastOperationGetGone) Error() string {
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetGone  %+v", 410, o.Payload)
}

func (o *ServiceBindingLastOperationGetGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
