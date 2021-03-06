// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateImagePipelineInput struct {
	_ struct{} `type:"structure"`

	// The idempotency token used to make this request idempotent.
	//
	// ClientToken is a required field
	ClientToken *string `locationName:"clientToken" min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The description of the image pipeline.
	Description *string `locationName:"description" min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the distribution configuration that will
	// be used to configure and distribute images updated by this image pipeline.
	DistributionConfigurationArn *string `locationName:"distributionConfigurationArn" type:"string"`

	// Collects additional information about the image being created, including
	// the operating system (OS) version and package list. This information is used
	// to enhance the overall experience of using EC2 Image Builder. Enabled by
	// default.
	EnhancedImageMetadataEnabled *bool `locationName:"enhancedImageMetadataEnabled" type:"boolean"`

	// The Amazon Resource Name (ARN) of the image pipeline that you want to update.
	//
	// ImagePipelineArn is a required field
	ImagePipelineArn *string `locationName:"imagePipelineArn" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the image recipe that will be used to configure
	// images updated by this image pipeline.
	//
	// ImageRecipeArn is a required field
	ImageRecipeArn *string `locationName:"imageRecipeArn" type:"string" required:"true"`

	// The image test configuration of the image pipeline.
	ImageTestsConfiguration *ImageTestsConfiguration `locationName:"imageTestsConfiguration" type:"structure"`

	// The Amazon Resource Name (ARN) of the infrastructure configuration that will
	// be used to build images updated by this image pipeline.
	//
	// InfrastructureConfigurationArn is a required field
	InfrastructureConfigurationArn *string `locationName:"infrastructureConfigurationArn" type:"string" required:"true"`

	// The schedule of the image pipeline.
	Schedule *Schedule `locationName:"schedule" type:"structure"`

	// The status of the image pipeline.
	Status PipelineStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateImagePipelineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateImagePipelineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateImagePipelineInput"}

	if s.ClientToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientToken"))
	}
	if s.ClientToken != nil && len(*s.ClientToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 1))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.ImagePipelineArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImagePipelineArn"))
	}

	if s.ImageRecipeArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImageRecipeArn"))
	}

	if s.InfrastructureConfigurationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("InfrastructureConfigurationArn"))
	}
	if s.ImageTestsConfiguration != nil {
		if err := s.ImageTestsConfiguration.Validate(); err != nil {
			invalidParams.AddNested("ImageTestsConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.Schedule != nil {
		if err := s.Schedule.Validate(); err != nil {
			invalidParams.AddNested("Schedule", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateImagePipelineInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	var ClientToken string
	if s.ClientToken != nil {
		ClientToken = *s.ClientToken
	} else {
		ClientToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DistributionConfigurationArn != nil {
		v := *s.DistributionConfigurationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "distributionConfigurationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EnhancedImageMetadataEnabled != nil {
		v := *s.EnhancedImageMetadataEnabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enhancedImageMetadataEnabled", protocol.BoolValue(v), metadata)
	}
	if s.ImagePipelineArn != nil {
		v := *s.ImagePipelineArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "imagePipelineArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ImageRecipeArn != nil {
		v := *s.ImageRecipeArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "imageRecipeArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ImageTestsConfiguration != nil {
		v := s.ImageTestsConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "imageTestsConfiguration", v, metadata)
	}
	if s.InfrastructureConfigurationArn != nil {
		v := *s.InfrastructureConfigurationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "infrastructureConfigurationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Schedule != nil {
		v := s.Schedule

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "schedule", v, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

type UpdateImagePipelineOutput struct {
	_ struct{} `type:"structure"`

	// The idempotency token used to make this request idempotent.
	ClientToken *string `locationName:"clientToken" min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the image pipeline that was updated by
	// this request.
	ImagePipelineArn *string `locationName:"imagePipelineArn" type:"string"`

	// The request ID that uniquely identifies this request.
	RequestId *string `locationName:"requestId" min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateImagePipelineOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateImagePipelineOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ImagePipelineArn != nil {
		v := *s.ImagePipelineArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "imagePipelineArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateImagePipeline = "UpdateImagePipeline"

// UpdateImagePipelineRequest returns a request value for making API operation for
// EC2 Image Builder.
//
// Updates a new image pipeline. Image pipelines enable you to automate the
// creation and distribution of images.
//
//    // Example sending a request using UpdateImagePipelineRequest.
//    req := client.UpdateImagePipelineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/imagebuilder-2019-12-02/UpdateImagePipeline
func (c *Client) UpdateImagePipelineRequest(input *UpdateImagePipelineInput) UpdateImagePipelineRequest {
	op := &aws.Operation{
		Name:       opUpdateImagePipeline,
		HTTPMethod: "PUT",
		HTTPPath:   "/UpdateImagePipeline",
	}

	if input == nil {
		input = &UpdateImagePipelineInput{}
	}

	req := c.newRequest(op, input, &UpdateImagePipelineOutput{})
	return UpdateImagePipelineRequest{Request: req, Input: input, Copy: c.UpdateImagePipelineRequest}
}

// UpdateImagePipelineRequest is the request type for the
// UpdateImagePipeline API operation.
type UpdateImagePipelineRequest struct {
	*aws.Request
	Input *UpdateImagePipelineInput
	Copy  func(*UpdateImagePipelineInput) UpdateImagePipelineRequest
}

// Send marshals and sends the UpdateImagePipeline API request.
func (r UpdateImagePipelineRequest) Send(ctx context.Context) (*UpdateImagePipelineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateImagePipelineResponse{
		UpdateImagePipelineOutput: r.Request.Data.(*UpdateImagePipelineOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateImagePipelineResponse is the response type for the
// UpdateImagePipeline API operation.
type UpdateImagePipelineResponse struct {
	*UpdateImagePipelineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateImagePipeline request.
func (r *UpdateImagePipelineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
