// Code generated by go-swagger; DO NOT EDIT.

package access_control

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

// ResendInviteReader is a Reader for the ResendInvite structure.
type ResendInviteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResendInviteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewResendInviteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewResendInviteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResendInviteAccepted creates a ResendInviteAccepted with default headers values
func NewResendInviteAccepted() *ResendInviteAccepted {
	return &ResendInviteAccepted{}
}

/*ResendInviteAccepted handles this case with default header values.

The operation completed successfully
*/
type ResendInviteAccepted struct {
	Payload *ResendInviteAcceptedBody
}

func (o *ResendInviteAccepted) Error() string {
	return fmt.Sprintf("[POST /api/invites/{inviteId}/resend][%d] resendInviteAccepted  %+v", 202, o.Payload)
}

func (o *ResendInviteAccepted) GetPayload() *ResendInviteAcceptedBody {
	return o.Payload
}

func (o *ResendInviteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ResendInviteAcceptedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResendInviteDefault creates a ResendInviteDefault with default headers values
func NewResendInviteDefault(code int) *ResendInviteDefault {
	return &ResendInviteDefault{
		_statusCode: code,
	}
}

/*ResendInviteDefault handles this case with default header values.

An error occurred
*/
type ResendInviteDefault struct {
	_statusCode int

	Payload *ResendInviteDefaultBody
}

// Code gets the status code for the resend invite default response
func (o *ResendInviteDefault) Code() int {
	return o._statusCode
}

func (o *ResendInviteDefault) Error() string {
	return fmt.Sprintf("[POST /api/invites/{inviteId}/resend][%d] resendInvite default  %+v", o._statusCode, o.Payload)
}

func (o *ResendInviteDefault) GetPayload() *ResendInviteDefaultBody {
	return o.Payload
}

func (o *ResendInviteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ResendInviteDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ResendInviteAcceptedBody Success response
swagger:model ResendInviteAcceptedBody
*/
type ResendInviteAcceptedBody struct {

	// Did this succeed?
	Success bool `json:"success,omitempty"`
}

// Validate validates this resend invite accepted body
func (o *ResendInviteAcceptedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ResendInviteAcceptedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ResendInviteAcceptedBody) UnmarshalBinary(b []byte) error {
	var res ResendInviteAcceptedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ResendInviteDefaultBody Error response
swagger:model ResendInviteDefaultBody
*/
type ResendInviteDefaultBody struct {

	// error
	Error *models.Error `json:"error,omitempty"`
}

// Validate validates this resend invite default body
func (o *ResendInviteDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ResendInviteDefaultBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resendInvite default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ResendInviteDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ResendInviteDefaultBody) UnmarshalBinary(b []byte) error {
	var res ResendInviteDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
