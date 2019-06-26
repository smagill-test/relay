// Code generated by go-swagger; DO NOT EDIT.

package workflow_runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new workflow runs v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for workflow runs v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateWorkflowRun creates new workflow run associated with given workflow
*/
func (a *Client) CreateWorkflowRun(params *CreateWorkflowRunParams, authInfo runtime.ClientAuthInfoWriter) (*CreateWorkflowRunCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateWorkflowRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createWorkflowRun",
		Method:             "POST",
		PathPattern:        "/api/workflows/{workflow_name}/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateWorkflowRunReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateWorkflowRunCreated), nil

}

/*
GetWorkflowRun gets a workflow run accessed with a workflow name and run number
*/
func (a *Client) GetWorkflowRun(params *GetWorkflowRunParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkflowRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowRunParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkflowRun",
		Method:             "GET",
		PathPattern:        "/api/workflows/{workflow_name}/runs/{run_number}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWorkflowRunReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkflowRunOK), nil

}

/*
GetWorkflowRunLogs returns the logs from an entire workflow execution accessed by workflow name and run number
*/
func (a *Client) GetWorkflowRunLogs(params *GetWorkflowRunLogsParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkflowRunLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowRunLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkflowRunLogs",
		Method:             "GET",
		PathPattern:        "/api/workflows/{workflow_name}/runs/{run_number}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWorkflowRunLogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkflowRunLogsOK), nil

}

/*
GetWorkflowRunStepLogs returns the logs from a single workflow run step acessed by workflow name run number and step name
*/
func (a *Client) GetWorkflowRunStepLogs(params *GetWorkflowRunStepLogsParams, authInfo runtime.ClientAuthInfoWriter) (*GetWorkflowRunStepLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowRunStepLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkflowRunStepLogs",
		Method:             "GET",
		PathPattern:        "/api/workflows/{workflow_name}/runs/{run_number}/steps/{step_name}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWorkflowRunStepLogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkflowRunStepLogsOK), nil

}

/*
ListWorkflowRuns gets all workflow runs associated with given workflow id
*/
func (a *Client) ListWorkflowRuns(params *ListWorkflowRunsParams, authInfo runtime.ClientAuthInfoWriter) (*ListWorkflowRunsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListWorkflowRunsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listWorkflowRuns",
		Method:             "GET",
		PathPattern:        "/api/workflows/{workflow_name}/runs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListWorkflowRunsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListWorkflowRunsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
