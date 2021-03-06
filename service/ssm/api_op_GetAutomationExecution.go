// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetAutomationExecutionInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for an existing automation execution to examine. The
	// execution ID is returned by StartAutomationExecution when the execution of
	// an Automation document is initiated.
	//
	// AutomationExecutionId is a required field
	AutomationExecutionId *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s GetAutomationExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAutomationExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAutomationExecutionInput"}

	if s.AutomationExecutionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutomationExecutionId"))
	}
	if s.AutomationExecutionId != nil && len(*s.AutomationExecutionId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("AutomationExecutionId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetAutomationExecutionOutput struct {
	_ struct{} `type:"structure"`

	// Detailed information about the current state of an automation execution.
	AutomationExecution *AutomationExecution `type:"structure"`
}

// String returns the string representation
func (s GetAutomationExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetAutomationExecution = "GetAutomationExecution"

// GetAutomationExecutionRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Get detailed information about a particular Automation execution.
//
//    // Example sending a request using GetAutomationExecutionRequest.
//    req := client.GetAutomationExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/GetAutomationExecution
func (c *Client) GetAutomationExecutionRequest(input *GetAutomationExecutionInput) GetAutomationExecutionRequest {
	op := &aws.Operation{
		Name:       opGetAutomationExecution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAutomationExecutionInput{}
	}

	req := c.newRequest(op, input, &GetAutomationExecutionOutput{})
	return GetAutomationExecutionRequest{Request: req, Input: input, Copy: c.GetAutomationExecutionRequest}
}

// GetAutomationExecutionRequest is the request type for the
// GetAutomationExecution API operation.
type GetAutomationExecutionRequest struct {
	*aws.Request
	Input *GetAutomationExecutionInput
	Copy  func(*GetAutomationExecutionInput) GetAutomationExecutionRequest
}

// Send marshals and sends the GetAutomationExecution API request.
func (r GetAutomationExecutionRequest) Send(ctx context.Context) (*GetAutomationExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAutomationExecutionResponse{
		GetAutomationExecutionOutput: r.Request.Data.(*GetAutomationExecutionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAutomationExecutionResponse is the response type for the
// GetAutomationExecution API operation.
type GetAutomationExecutionResponse struct {
	*GetAutomationExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAutomationExecution request.
func (r *GetAutomationExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
