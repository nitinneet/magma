// Code generated by go-swagger; DO NOT EDIT.

package a_p_ns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// PutLTENetworkIDAPNSAPNNameReader is a Reader for the PutLTENetworkIDAPNSAPNName structure.
type PutLTENetworkIDAPNSAPNNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDAPNSAPNNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDAPNSAPNNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDAPNSAPNNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDAPNSAPNNameNoContent creates a PutLTENetworkIDAPNSAPNNameNoContent with default headers values
func NewPutLTENetworkIDAPNSAPNNameNoContent() *PutLTENetworkIDAPNSAPNNameNoContent {
	return &PutLTENetworkIDAPNSAPNNameNoContent{}
}

/*PutLTENetworkIDAPNSAPNNameNoContent handles this case with default header values.

Success
*/
type PutLTENetworkIDAPNSAPNNameNoContent struct {
}

func (o *PutLTENetworkIDAPNSAPNNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/apns/{apn_name}][%d] putLteNetworkIdApnsApnNameNoContent ", 204)
}

func (o *PutLTENetworkIDAPNSAPNNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDAPNSAPNNameDefault creates a PutLTENetworkIDAPNSAPNNameDefault with default headers values
func NewPutLTENetworkIDAPNSAPNNameDefault(code int) *PutLTENetworkIDAPNSAPNNameDefault {
	return &PutLTENetworkIDAPNSAPNNameDefault{
		_statusCode: code,
	}
}

/*PutLTENetworkIDAPNSAPNNameDefault handles this case with default header values.

Unexpected Error
*/
type PutLTENetworkIDAPNSAPNNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID APNS APN name default response
func (o *PutLTENetworkIDAPNSAPNNameDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDAPNSAPNNameDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/apns/{apn_name}][%d] PutLTENetworkIDAPNSAPNName default  %+v", o._statusCode, o.Payload)
}

func (o *PutLTENetworkIDAPNSAPNNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDAPNSAPNNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}