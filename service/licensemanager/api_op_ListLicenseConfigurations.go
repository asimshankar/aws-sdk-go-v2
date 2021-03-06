// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package licensemanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListLicenseConfigurationsInput struct {
	_ struct{} `type:"structure"`

	// Filters to scope the results. The following filters and logical operators
	// are supported:
	//
	//    * licenseCountingType - The dimension on which licenses are counted (vCPU).
	//    Logical operators are EQUALS | NOT_EQUALS.
	//
	//    * enforceLicenseCount - A Boolean value that indicates whether hard license
	//    enforcement is used. Logical operators are EQUALS | NOT_EQUALS.
	//
	//    * usagelimitExceeded - A Boolean value that indicates whether the available
	//    licenses have been exceeded. Logical operators are EQUALS | NOT_EQUALS.
	Filters []Filter `type:"list"`

	// Amazon Resource Names (ARN) of the license configurations.
	LicenseConfigurationArns []string `type:"list"`

	// Maximum number of results to return in a single call.
	MaxResults *int64 `type:"integer"`

	// Token for the next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListLicenseConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

type ListLicenseConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the license configurations.
	LicenseConfigurations []LicenseConfiguration `type:"list"`

	// Token for the next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListLicenseConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListLicenseConfigurations = "ListLicenseConfigurations"

// ListLicenseConfigurationsRequest returns a request value for making API operation for
// AWS License Manager.
//
// Lists the license configurations for your account.
//
//    // Example sending a request using ListLicenseConfigurationsRequest.
//    req := client.ListLicenseConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/license-manager-2018-08-01/ListLicenseConfigurations
func (c *Client) ListLicenseConfigurationsRequest(input *ListLicenseConfigurationsInput) ListLicenseConfigurationsRequest {
	op := &aws.Operation{
		Name:       opListLicenseConfigurations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListLicenseConfigurationsInput{}
	}

	req := c.newRequest(op, input, &ListLicenseConfigurationsOutput{})
	return ListLicenseConfigurationsRequest{Request: req, Input: input, Copy: c.ListLicenseConfigurationsRequest}
}

// ListLicenseConfigurationsRequest is the request type for the
// ListLicenseConfigurations API operation.
type ListLicenseConfigurationsRequest struct {
	*aws.Request
	Input *ListLicenseConfigurationsInput
	Copy  func(*ListLicenseConfigurationsInput) ListLicenseConfigurationsRequest
}

// Send marshals and sends the ListLicenseConfigurations API request.
func (r ListLicenseConfigurationsRequest) Send(ctx context.Context) (*ListLicenseConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListLicenseConfigurationsResponse{
		ListLicenseConfigurationsOutput: r.Request.Data.(*ListLicenseConfigurationsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListLicenseConfigurationsResponse is the response type for the
// ListLicenseConfigurations API operation.
type ListLicenseConfigurationsResponse struct {
	*ListLicenseConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListLicenseConfigurations request.
func (r *ListLicenseConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
