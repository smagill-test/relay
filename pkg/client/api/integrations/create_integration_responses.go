// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/puppetlabs/nebula-cli/pkg/client/api/models"
)

// CreateIntegrationReader is a Reader for the CreateIntegration structure.
type CreateIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateIntegrationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateIntegrationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateIntegrationCreated creates a CreateIntegrationCreated with default headers values
func NewCreateIntegrationCreated() *CreateIntegrationCreated {
	return &CreateIntegrationCreated{}
}

/*CreateIntegrationCreated handles this case with default header values.

The newly created integration
*/
type CreateIntegrationCreated struct {
	Payload *CreateIntegrationCreatedBody
}

func (o *CreateIntegrationCreated) Error() string {
	return fmt.Sprintf("[POST /api/integrations][%d] createIntegrationCreated  %+v", 201, o.Payload)
}

func (o *CreateIntegrationCreated) GetPayload() *CreateIntegrationCreatedBody {
	return o.Payload
}

func (o *CreateIntegrationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateIntegrationCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIntegrationDefault creates a CreateIntegrationDefault with default headers values
func NewCreateIntegrationDefault(code int) *CreateIntegrationDefault {
	return &CreateIntegrationDefault{
		_statusCode: code,
	}
}

/*CreateIntegrationDefault handles this case with default header values.

An error occurred
*/
type CreateIntegrationDefault struct {
	_statusCode int

	Payload *CreateIntegrationDefaultBody
}

// Code gets the status code for the create integration default response
func (o *CreateIntegrationDefault) Code() int {
	return o._statusCode
}

func (o *CreateIntegrationDefault) Error() string {
	return fmt.Sprintf("[POST /api/integrations][%d] createIntegration default  %+v", o._statusCode, o.Payload)
}

func (o *CreateIntegrationDefault) GetPayload() *CreateIntegrationDefaultBody {
	return o.Payload
}

func (o *CreateIntegrationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateIntegrationDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateIntegrationBody Required fields to create a nebula external integration
swagger:model CreateIntegrationBody
*/
type CreateIntegrationBody struct {

	// auth
	Auth interface{} `json:"auth,omitempty"`

	// A descriptive integration name
	Name string `json:"name,omitempty"`
}

// Validate validates this create integration body
func (o *CreateIntegrationBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateIntegrationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateIntegrationBody) UnmarshalBinary(b []byte) error {
	var res CreateIntegrationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateIntegrationCreatedBody create integration created body
swagger:model CreateIntegrationCreatedBody
*/
type CreateIntegrationCreatedBody struct {
	models.IntegrationEntity

	// integration
	Integration *models.Integration `json:"integration,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *CreateIntegrationCreatedBody) UnmarshalJSON(raw []byte) error {
	// CreateIntegrationCreatedBodyAO0
	var createIntegrationCreatedBodyAO0 models.IntegrationEntity
	if err := swag.ReadJSON(raw, &createIntegrationCreatedBodyAO0); err != nil {
		return err
	}
	o.IntegrationEntity = createIntegrationCreatedBodyAO0

	// CreateIntegrationCreatedBodyAO1
	var dataCreateIntegrationCreatedBodyAO1 struct {
		Integration *models.Integration `json:"integration,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataCreateIntegrationCreatedBodyAO1); err != nil {
		return err
	}

	o.Integration = dataCreateIntegrationCreatedBodyAO1.Integration

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o CreateIntegrationCreatedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	createIntegrationCreatedBodyAO0, err := swag.WriteJSON(o.IntegrationEntity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, createIntegrationCreatedBodyAO0)

	var dataCreateIntegrationCreatedBodyAO1 struct {
		Integration *models.Integration `json:"integration,omitempty"`
	}

	dataCreateIntegrationCreatedBodyAO1.Integration = o.Integration

	jsonDataCreateIntegrationCreatedBodyAO1, errCreateIntegrationCreatedBodyAO1 := swag.WriteJSON(dataCreateIntegrationCreatedBodyAO1)
	if errCreateIntegrationCreatedBodyAO1 != nil {
		return nil, errCreateIntegrationCreatedBodyAO1
	}
	_parts = append(_parts, jsonDataCreateIntegrationCreatedBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this create integration created body
func (o *CreateIntegrationCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.IntegrationEntity
	if err := o.IntegrationEntity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIntegration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateIntegrationCreatedBody) validateIntegration(formats strfmt.Registry) error {

	if swag.IsZero(o.Integration) { // not required
		return nil
	}

	if o.Integration != nil {
		if err := o.Integration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createIntegrationCreated" + "." + "integration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateIntegrationCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateIntegrationCreatedBody) UnmarshalBinary(b []byte) error {
	var res CreateIntegrationCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateIntegrationDefaultBody Error response
swagger:model CreateIntegrationDefaultBody
*/
type CreateIntegrationDefaultBody struct {

	// error
	Error *models.Error `json:"error,omitempty"`
}

// Validate validates this create integration default body
func (o *CreateIntegrationDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateIntegrationDefaultBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createIntegration default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateIntegrationDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateIntegrationDefaultBody) UnmarshalBinary(b []byte) error {
	var res CreateIntegrationDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
