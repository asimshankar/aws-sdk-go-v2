// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package synthetics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type StopCanaryInput struct {
	_ struct{} `type:"structure"`

	// The name of the canary that you want to stop. To find the names of your canaries,
	// use DescribeCanaries.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopCanaryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopCanaryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopCanaryInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StopCanaryInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type StopCanaryOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopCanaryOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StopCanaryOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opStopCanary = "StopCanary"

// StopCanaryRequest returns a request value for making API operation for
// Synthetics.
//
// Stops the canary to prevent all future runs. If the canary is currently running,
// Synthetics stops waiting for the current run of the specified canary to complete.
// The run that is in progress completes on its own, publishes metrics, and
// uploads artifacts, but it is not recorded in Synthetics as a completed run.
//
// You can use StartCanary to start it running again with the canary’s current
// schedule at any point in the future.
//
//    // Example sending a request using StopCanaryRequest.
//    req := client.StopCanaryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/synthetics-2017-10-11/StopCanary
func (c *Client) StopCanaryRequest(input *StopCanaryInput) StopCanaryRequest {
	op := &aws.Operation{
		Name:       opStopCanary,
		HTTPMethod: "POST",
		HTTPPath:   "/canary/{name}/stop",
	}

	if input == nil {
		input = &StopCanaryInput{}
	}

	req := c.newRequest(op, input, &StopCanaryOutput{})
	return StopCanaryRequest{Request: req, Input: input, Copy: c.StopCanaryRequest}
}

// StopCanaryRequest is the request type for the
// StopCanary API operation.
type StopCanaryRequest struct {
	*aws.Request
	Input *StopCanaryInput
	Copy  func(*StopCanaryInput) StopCanaryRequest
}

// Send marshals and sends the StopCanary API request.
func (r StopCanaryRequest) Send(ctx context.Context) (*StopCanaryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopCanaryResponse{
		StopCanaryOutput: r.Request.Data.(*StopCanaryOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopCanaryResponse is the response type for the
// StopCanary API operation.
type StopCanaryResponse struct {
	*StopCanaryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopCanary request.
func (r *StopCanaryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
