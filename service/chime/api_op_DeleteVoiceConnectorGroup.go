// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteVoiceConnectorGroupInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime Voice Connector group ID.
	//
	// VoiceConnectorGroupId is a required field
	VoiceConnectorGroupId *string `location:"uri" locationName:"voiceConnectorGroupId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVoiceConnectorGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVoiceConnectorGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVoiceConnectorGroupInput"}

	if s.VoiceConnectorGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VoiceConnectorGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVoiceConnectorGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.VoiceConnectorGroupId != nil {
		v := *s.VoiceConnectorGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "voiceConnectorGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteVoiceConnectorGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteVoiceConnectorGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteVoiceConnectorGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteVoiceConnectorGroup = "DeleteVoiceConnectorGroup"

// DeleteVoiceConnectorGroupRequest returns a request value for making API operation for
// Amazon Chime.
//
// Deletes the specified Amazon Chime Voice Connector group. Any VoiceConnectorItems
// and phone numbers associated with the group must be removed before it can
// be deleted.
//
//    // Example sending a request using DeleteVoiceConnectorGroupRequest.
//    req := client.DeleteVoiceConnectorGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DeleteVoiceConnectorGroup
func (c *Client) DeleteVoiceConnectorGroupRequest(input *DeleteVoiceConnectorGroupInput) DeleteVoiceConnectorGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteVoiceConnectorGroup,
		HTTPMethod: "DELETE",
		HTTPPath:   "/voice-connector-groups/{voiceConnectorGroupId}",
	}

	if input == nil {
		input = &DeleteVoiceConnectorGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteVoiceConnectorGroupOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteVoiceConnectorGroupRequest{Request: req, Input: input, Copy: c.DeleteVoiceConnectorGroupRequest}
}

// DeleteVoiceConnectorGroupRequest is the request type for the
// DeleteVoiceConnectorGroup API operation.
type DeleteVoiceConnectorGroupRequest struct {
	*aws.Request
	Input *DeleteVoiceConnectorGroupInput
	Copy  func(*DeleteVoiceConnectorGroupInput) DeleteVoiceConnectorGroupRequest
}

// Send marshals and sends the DeleteVoiceConnectorGroup API request.
func (r DeleteVoiceConnectorGroupRequest) Send(ctx context.Context) (*DeleteVoiceConnectorGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVoiceConnectorGroupResponse{
		DeleteVoiceConnectorGroupOutput: r.Request.Data.(*DeleteVoiceConnectorGroupOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVoiceConnectorGroupResponse is the response type for the
// DeleteVoiceConnectorGroup API operation.
type DeleteVoiceConnectorGroupResponse struct {
	*DeleteVoiceConnectorGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVoiceConnectorGroup request.
func (r *DeleteVoiceConnectorGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
