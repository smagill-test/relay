// Code generated by go-swagger; DO NOT EDIT.

package workflow_runs

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

// GetWorkflowRunsReader is a Reader for the GetWorkflowRuns structure.
type GetWorkflowRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWorkflowRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWorkflowRunsOK creates a GetWorkflowRunsOK with default headers values
func NewGetWorkflowRunsOK() *GetWorkflowRunsOK {
	return &GetWorkflowRunsOK{}
}

/*GetWorkflowRunsOK handles this case with default header values.

The list of workflow runs
*/
type GetWorkflowRunsOK struct {
	Payload *GetWorkflowRunsOKBody
}

func (o *GetWorkflowRunsOK) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflowName}/runs][%d] getWorkflowRunsOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowRunsOK) GetPayload() *GetWorkflowRunsOKBody {
	return o.Payload
}

func (o *GetWorkflowRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowRunsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowRunsDefault creates a GetWorkflowRunsDefault with default headers values
func NewGetWorkflowRunsDefault(code int) *GetWorkflowRunsDefault {
	return &GetWorkflowRunsDefault{
		_statusCode: code,
	}
}

/*GetWorkflowRunsDefault handles this case with default header values.

An error occurred
*/
type GetWorkflowRunsDefault struct {
	_statusCode int

	Payload *GetWorkflowRunsDefaultBody
}

// Code gets the status code for the get workflow runs default response
func (o *GetWorkflowRunsDefault) Code() int {
	return o._statusCode
}

func (o *GetWorkflowRunsDefault) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflowName}/runs][%d] getWorkflowRuns default  %+v", o._statusCode, o.Payload)
}

func (o *GetWorkflowRunsDefault) GetPayload() *GetWorkflowRunsDefaultBody {
	return o.Payload
}

func (o *GetWorkflowRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowRunsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWorkflowRunsDefaultBody Error response
swagger:model GetWorkflowRunsDefaultBody
*/
type GetWorkflowRunsDefaultBody struct {

	// error
	Error *models.Error `json:"error,omitempty"`
}

// Validate validates this get workflow runs default body
func (o *GetWorkflowRunsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowRunsDefaultBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWorkflowRuns default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowRunsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowRunsDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowRunsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWorkflowRunsOKBody The response type for retrieving the list of workflow runs
swagger:model GetWorkflowRunsOKBody
*/
type GetWorkflowRunsOKBody struct {

	// A list of workflow run summaries
	Runs []*models.WorkflowRunSummary `json:"runs"`
}

// Validate validates this get workflow runs o k body
func (o *GetWorkflowRunsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRuns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowRunsOKBody) validateRuns(formats strfmt.Registry) error {

	if swag.IsZero(o.Runs) { // not required
		return nil
	}

	for i := 0; i < len(o.Runs); i++ {
		if swag.IsZero(o.Runs[i]) { // not required
			continue
		}

		if o.Runs[i] != nil {
			if err := o.Runs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getWorkflowRunsOK" + "." + "runs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowRunsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowRunsOKBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowRunsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
