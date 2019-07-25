// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/ListBucketInventoryConfigurationsRequest
type ListBucketInventoryConfigurationsInput struct {
	_ struct{} `type:"structure"`

	// The name of the bucket containing the inventory configurations to retrieve.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The marker used to continue an inventory configuration listing that has been
	// truncated. Use the NextContinuationToken from a previously truncated list
	// response to continue the listing. The continuation token is an opaque value
	// that Amazon S3 understands.
	ContinuationToken *string `location:"querystring" locationName:"continuation-token" type:"string"`
}

// String returns the string representation
func (s ListBucketInventoryConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBucketInventoryConfigurationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListBucketInventoryConfigurationsInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *ListBucketInventoryConfigurationsInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListBucketInventoryConfigurationsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.ContinuationToken != nil {
		v := *s.ContinuationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "continuation-token", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/ListBucketInventoryConfigurationsOutput
type ListBucketInventoryConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	// If sent in the request, the marker that is used as a starting point for this
	// inventory configuration list response.
	ContinuationToken *string `type:"string"`

	// The list of inventory configurations for a bucket.
	InventoryConfigurationList []InventoryConfiguration `locationName:"InventoryConfiguration" type:"list" flattened:"true"`

	// Indicates whether the returned list of inventory configurations is truncated
	// in this response. A value of true indicates that the list is truncated.
	IsTruncated *bool `type:"boolean"`

	// The marker used to continue this inventory configuration listing. Use the
	// NextContinuationToken from this response to continue the listing in a subsequent
	// request. The continuation token is an opaque value that Amazon S3 understands.
	NextContinuationToken *string `type:"string"`
}

// String returns the string representation
func (s ListBucketInventoryConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListBucketInventoryConfigurationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ContinuationToken != nil {
		v := *s.ContinuationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ContinuationToken", protocol.StringValue(v), metadata)
	}
	if s.InventoryConfigurationList != nil {
		v := s.InventoryConfigurationList

		metadata := protocol.Metadata{Flatten: true}
		ls0 := e.List(protocol.BodyTarget, "InventoryConfiguration", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.IsTruncated != nil {
		v := *s.IsTruncated

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IsTruncated", protocol.BoolValue(v), metadata)
	}
	if s.NextContinuationToken != nil {
		v := *s.NextContinuationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextContinuationToken", protocol.StringValue(v), metadata)
	}
	return nil
}

// UnmarshalAWSXML decodes the AWS API shape using the passed in *xml.Decoder.
func (s *ListBucketInventoryConfigurationsOutput) UnmarshalAWSXML(d *xml.Decoder) (err error) {
	defer func() {
		if err != nil {
			*s = ListBucketInventoryConfigurationsOutput{}
		}
	}()
	for {
		tok, err := d.Token()
		if tok == nil || err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("fail to UnmarshalAWSXML ListBucketInventoryConfigurationsOutput, %s", err)
		}
		start, ok := tok.(xml.StartElement)
		if !ok {
			continue
		}
		err = s.unmarshalAWSXML(d, start)
		if err != nil {
			return fmt.Errorf("fail to UnmarshalAWSXML ListBucketInventoryConfigurationsOutput, %s", err)
		}
		return nil
	}
}

func (s *ListBucketInventoryConfigurationsOutput) unmarshalAWSXML(d *xml.Decoder, head xml.StartElement) (err error) {
	defer func() {
		if err != nil {
			*s = ListBucketInventoryConfigurationsOutput{}
		}
	}()
	name := ""
	for {
		tok, err := d.Token()
		if tok == nil || err != nil {
			if err == io.EOF {
				return nil
			}
		}
		if end, ok := tok.(xml.EndElement); ok {
			name = end.Name.Local
			if name == head.Name.Local {
				return nil
			}
		}
		if start, ok := tok.(xml.StartElement); ok {
			switch name = start.Name.Local; name {
			case "ContinuationToken":
				tok, err = d.Token()
				if tok == nil || err != nil {
					return fmt.Errorf("fail to UnmarshalAWSXML ListBucketInventoryConfigurationsOutput.%s, %s", name, err)
				}
				v, _ := tok.(xml.CharData)
				value := string(v)
				s.ContinuationToken = &value
			case "InventoryConfiguration":
				if s.InventoryConfigurationList == nil {
					s.InventoryConfigurationList = make([]InventoryConfiguration, 0)
				}
				err := unmarshalAWSXMLFlattenedListInventoryConfigurationList(&s.InventoryConfigurationList, d, start)
				if err != nil {
					return fmt.Errorf("fail to UnmarshalAWSXML ListBucketInventoryConfigurationsOutput.%s, %s", name, err)
				}
			case "IsTruncated":
				tok, err = d.Token()
				if tok == nil || err != nil {
					return fmt.Errorf("fail to UnmarshalAWSXML ListBucketInventoryConfigurationsOutput.%s, %s", name, err)
				}
				v, _ := tok.(xml.CharData)
				value, _ := strconv.ParseBool(string(v))
				s.IsTruncated = &value
			case "NextContinuationToken":
				tok, err = d.Token()
				if tok == nil || err != nil {
					return fmt.Errorf("fail to UnmarshalAWSXML ListBucketInventoryConfigurationsOutput.%s, %s", name, err)
				}
				v, _ := tok.(xml.CharData)
				value := string(v)
				s.NextContinuationToken = &value
			default:
				err := d.Skip()
				if err != nil {
					return fmt.Errorf("fail to UnmarshalAWSXML ListBucketInventoryConfigurationsOutput.%s, %s", name, err)
				}
			}
		}
	}
}

const opListBucketInventoryConfigurations = "ListBucketInventoryConfigurations"

// ListBucketInventoryConfigurationsRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Returns a list of inventory configurations for the bucket.
//
//    // Example sending a request using ListBucketInventoryConfigurationsRequest.
//    req := client.ListBucketInventoryConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/ListBucketInventoryConfigurations
func (c *Client) ListBucketInventoryConfigurationsRequest(input *ListBucketInventoryConfigurationsInput) ListBucketInventoryConfigurationsRequest {
	op := &aws.Operation{
		Name:       opListBucketInventoryConfigurations,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?inventory",
	}

	if input == nil {
		input = &ListBucketInventoryConfigurationsInput{}
	}

	req := c.newRequest(op, input, &ListBucketInventoryConfigurationsOutput{})
	return ListBucketInventoryConfigurationsRequest{Request: req, Input: input, Copy: c.ListBucketInventoryConfigurationsRequest}
}

// ListBucketInventoryConfigurationsRequest is the request type for the
// ListBucketInventoryConfigurations API operation.
type ListBucketInventoryConfigurationsRequest struct {
	*aws.Request
	Input *ListBucketInventoryConfigurationsInput
	Copy  func(*ListBucketInventoryConfigurationsInput) ListBucketInventoryConfigurationsRequest
}

// Send marshals and sends the ListBucketInventoryConfigurations API request.
func (r ListBucketInventoryConfigurationsRequest) Send(ctx context.Context) (*ListBucketInventoryConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListBucketInventoryConfigurationsResponse{
		ListBucketInventoryConfigurationsOutput: r.Request.Data.(*ListBucketInventoryConfigurationsOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListBucketInventoryConfigurationsResponse is the response type for the
// ListBucketInventoryConfigurations API operation.
type ListBucketInventoryConfigurationsResponse struct {
	*ListBucketInventoryConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListBucketInventoryConfigurations request.
func (r *ListBucketInventoryConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
