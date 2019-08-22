// Code generated by go-swagger; DO NOT EDIT.

package workflow_secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/puppetlabs/nebula-cli/pkg/client/api/models"
)

// ListWorkflowSecretsReader is a Reader for the ListWorkflowSecrets structure.
type ListWorkflowSecretsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListWorkflowSecretsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListWorkflowSecretsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListWorkflowSecretsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListWorkflowSecretsOK creates a ListWorkflowSecretsOK with default headers values
func NewListWorkflowSecretsOK() *ListWorkflowSecretsOK {
	return &ListWorkflowSecretsOK{}
}

/*ListWorkflowSecretsOK handles this case with default header values.

A list of secrets
*/
type ListWorkflowSecretsOK struct {
	Payload *ListWorkflowSecretsOKBody
}

func (o *ListWorkflowSecretsOK) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflowName}/secrets][%d] listWorkflowSecretsOK  %+v", 200, o.Payload)
}

func (o *ListWorkflowSecretsOK) GetPayload() *ListWorkflowSecretsOKBody {
	return o.Payload
}

func (o *ListWorkflowSecretsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListWorkflowSecretsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListWorkflowSecretsDefault creates a ListWorkflowSecretsDefault with default headers values
func NewListWorkflowSecretsDefault(code int) *ListWorkflowSecretsDefault {
	return &ListWorkflowSecretsDefault{
		_statusCode: code,
	}
}

/*ListWorkflowSecretsDefault handles this case with default header values.

An error occurred
*/
type ListWorkflowSecretsDefault struct {
	_statusCode int

	Payload *ListWorkflowSecretsDefaultBody
}

// Code gets the status code for the list workflow secrets default response
func (o *ListWorkflowSecretsDefault) Code() int {
	return o._statusCode
}

func (o *ListWorkflowSecretsDefault) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflowName}/secrets][%d] listWorkflowSecrets default  %+v", o._statusCode, o.Payload)
}

func (o *ListWorkflowSecretsDefault) GetPayload() *ListWorkflowSecretsDefaultBody {
	return o.Payload
}

func (o *ListWorkflowSecretsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListWorkflowSecretsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListWorkflowSecretsDefaultBody Error response
swagger:model ListWorkflowSecretsDefaultBody
*/
type ListWorkflowSecretsDefaultBody struct {

	// error
	Error *models.Error `json:"error,omitempty"`
}

// Validate validates this list workflow secrets default body
func (o *ListWorkflowSecretsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListWorkflowSecretsDefaultBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("listWorkflowSecrets default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListWorkflowSecretsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListWorkflowSecretsDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListWorkflowSecretsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListWorkflowSecretsOKBody The response type for retrieving the secrets in a workflow
swagger:model ListWorkflowSecretsOKBody
*/
type ListWorkflowSecretsOKBody struct {

	// A list of secrets
	Secrets []*models.WorkflowSecretSummary `json:"secrets"`
}

// Validate validates this list workflow secrets o k body
func (o *ListWorkflowSecretsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSecrets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListWorkflowSecretsOKBody) validateSecrets(formats strfmt.Registry) error {

	if swag.IsZero(o.Secrets) { // not required
		return nil
	}

	for i := 0; i < len(o.Secrets); i++ {
		if swag.IsZero(o.Secrets[i]) { // not required
			continue
		}

		if o.Secrets[i] != nil {
			if err := o.Secrets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listWorkflowSecretsOK" + "." + "secrets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListWorkflowSecretsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListWorkflowSecretsOKBody) UnmarshalBinary(b []byte) error {
	var res ListWorkflowSecretsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
