// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codeguruprofiler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The structure representing the putPermissionRequest.
type PutPermissionInput struct {
	_ struct{} `type:"structure"`

	// The list of actions that the users and roles can perform on the profiling
	// group.
	//
	// ActionGroup is a required field
	ActionGroup ActionGroup `location:"uri" locationName:"actionGroup" type:"string" required:"true" enum:"true"`

	// The list of role and user ARNs or the accountId that needs access (wildcards
	// are not allowed).
	//
	// Principals is a required field
	Principals []string `locationName:"principals" min:"1" type:"list" required:"true"`

	// The name of the profiling group.
	//
	// ProfilingGroupName is a required field
	ProfilingGroupName *string `location:"uri" locationName:"profilingGroupName" min:"1" type:"string" required:"true"`

	// A unique identifier for the current revision of the policy. This is required,
	// if a policy exists for the profiling group. This is not required when creating
	// the policy for the first time.
	RevisionId *string `locationName:"revisionId" type:"string"`
}

// String returns the string representation
func (s PutPermissionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutPermissionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutPermissionInput"}
	if len(s.ActionGroup) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ActionGroup"))
	}

	if s.Principals == nil {
		invalidParams.Add(aws.NewErrParamRequired("Principals"))
	}
	if s.Principals != nil && len(s.Principals) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Principals", 1))
	}

	if s.ProfilingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProfilingGroupName"))
	}
	if s.ProfilingGroupName != nil && len(*s.ProfilingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProfilingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutPermissionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Principals != nil {
		v := s.Principals

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "principals", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "revisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ActionGroup) > 0 {
		v := s.ActionGroup

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "actionGroup", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ProfilingGroupName != nil {
		v := *s.ProfilingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "profilingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The structure representing the putPermissionResponse.
type PutPermissionOutput struct {
	_ struct{} `type:"structure"`

	// The resource-based policy.
	//
	// Policy is a required field
	Policy *string `locationName:"policy" type:"string" required:"true"`

	// A unique identifier for the current revision of the policy.
	//
	// RevisionId is a required field
	RevisionId *string `locationName:"revisionId" type:"string" required:"true"`
}

// String returns the string representation
func (s PutPermissionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutPermissionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Policy != nil {
		v := *s.Policy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "policy", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "revisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opPutPermission = "PutPermission"

// PutPermissionRequest returns a request value for making API operation for
// Amazon CodeGuru Profiler.
//
// Provides permission to the principals. This overwrites the existing permissions,
// and is not additive.
//
//    // Example sending a request using PutPermissionRequest.
//    req := client.PutPermissionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeguruprofiler-2019-07-18/PutPermission
func (c *Client) PutPermissionRequest(input *PutPermissionInput) PutPermissionRequest {
	op := &aws.Operation{
		Name:       opPutPermission,
		HTTPMethod: "PUT",
		HTTPPath:   "/profilingGroups/{profilingGroupName}/policy/{actionGroup}",
	}

	if input == nil {
		input = &PutPermissionInput{}
	}

	req := c.newRequest(op, input, &PutPermissionOutput{})
	return PutPermissionRequest{Request: req, Input: input, Copy: c.PutPermissionRequest}
}

// PutPermissionRequest is the request type for the
// PutPermission API operation.
type PutPermissionRequest struct {
	*aws.Request
	Input *PutPermissionInput
	Copy  func(*PutPermissionInput) PutPermissionRequest
}

// Send marshals and sends the PutPermission API request.
func (r PutPermissionRequest) Send(ctx context.Context) (*PutPermissionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutPermissionResponse{
		PutPermissionOutput: r.Request.Data.(*PutPermissionOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutPermissionResponse is the response type for the
// PutPermission API operation.
type PutPermissionResponse struct {
	*PutPermissionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutPermission request.
func (r *PutPermissionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
