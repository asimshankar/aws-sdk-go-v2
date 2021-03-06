// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Removes the specified directory as a publisher to the specified SNS topic.
type DeregisterEventTopicInput struct {
	_ struct{} `type:"structure"`

	// The Directory ID to remove as a publisher. This directory will no longer
	// send messages to the specified SNS topic.
	//
	// DirectoryId is a required field
	DirectoryId *string `type:"string" required:"true"`

	// The name of the SNS topic from which to remove the directory as a publisher.
	//
	// TopicName is a required field
	TopicName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeregisterEventTopicInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterEventTopicInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterEventTopicInput"}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}

	if s.TopicName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TopicName"))
	}
	if s.TopicName != nil && len(*s.TopicName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TopicName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a DeregisterEventTopic request.
type DeregisterEventTopicOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeregisterEventTopicOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeregisterEventTopic = "DeregisterEventTopic"

// DeregisterEventTopicRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Removes the specified directory as a publisher to the specified SNS topic.
//
//    // Example sending a request using DeregisterEventTopicRequest.
//    req := client.DeregisterEventTopicRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/DeregisterEventTopic
func (c *Client) DeregisterEventTopicRequest(input *DeregisterEventTopicInput) DeregisterEventTopicRequest {
	op := &aws.Operation{
		Name:       opDeregisterEventTopic,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeregisterEventTopicInput{}
	}

	req := c.newRequest(op, input, &DeregisterEventTopicOutput{})
	return DeregisterEventTopicRequest{Request: req, Input: input, Copy: c.DeregisterEventTopicRequest}
}

// DeregisterEventTopicRequest is the request type for the
// DeregisterEventTopic API operation.
type DeregisterEventTopicRequest struct {
	*aws.Request
	Input *DeregisterEventTopicInput
	Copy  func(*DeregisterEventTopicInput) DeregisterEventTopicRequest
}

// Send marshals and sends the DeregisterEventTopic API request.
func (r DeregisterEventTopicRequest) Send(ctx context.Context) (*DeregisterEventTopicResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterEventTopicResponse{
		DeregisterEventTopicOutput: r.Request.Data.(*DeregisterEventTopicOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterEventTopicResponse is the response type for the
// DeregisterEventTopic API operation.
type DeregisterEventTopicResponse struct {
	*DeregisterEventTopicOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterEventTopic request.
func (r *DeregisterEventTopicResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
