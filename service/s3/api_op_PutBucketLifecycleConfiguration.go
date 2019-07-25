// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketLifecycleConfigurationRequest
type PutBucketLifecycleConfigurationInput struct {
	_ struct{} `type:"structure" payload:"LifecycleConfiguration"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Specifies the lifecycle configuration for objects in an Amazon S3 bucket.
	// For more information, see Object Lifecycle Management (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html)
	// in the Amazon Simple Storage Service Developer Guide.
	LifecycleConfiguration *BucketLifecycleConfiguration `locationName:"LifecycleConfiguration" type:"structure" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`
}

// String returns the string representation
func (s PutBucketLifecycleConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketLifecycleConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketLifecycleConfigurationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}
	if s.LifecycleConfiguration != nil {
		if err := s.LifecycleConfiguration.Validate(); err != nil {
			invalidParams.AddNested("LifecycleConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketLifecycleConfigurationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketLifecycleConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.LifecycleConfiguration != nil {
		v := s.LifecycleConfiguration

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "LifecycleConfiguration", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketLifecycleConfigurationOutput
type PutBucketLifecycleConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketLifecycleConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketLifecycleConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}
func (s *PutBucketLifecycleConfigurationOutput) unmarshalAWSXML(d *xml.Decoder, head xml.StartElement) (err error) {
	defer func() {
		if err != nil {
			*s = PutBucketLifecycleConfigurationOutput{}
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
			default:
				err := d.Skip()
				if err != nil {
					return fmt.Errorf("fail to UnmarshalAWSXML PutBucketLifecycleConfigurationOutput.%s, %s", name, err)
				}
			}
		}
	}
}

const opPutBucketLifecycleConfiguration = "PutBucketLifecycleConfiguration"

// PutBucketLifecycleConfigurationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Sets lifecycle configuration for your bucket. If a lifecycle configuration
// exists, it replaces it.
//
//    // Example sending a request using PutBucketLifecycleConfigurationRequest.
//    req := client.PutBucketLifecycleConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketLifecycleConfiguration
func (c *Client) PutBucketLifecycleConfigurationRequest(input *PutBucketLifecycleConfigurationInput) PutBucketLifecycleConfigurationRequest {
	op := &aws.Operation{
		Name:       opPutBucketLifecycleConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?lifecycle",
	}

	if input == nil {
		input = &PutBucketLifecycleConfigurationInput{}
	}

	req := c.newRequest(op, input, &PutBucketLifecycleConfigurationOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutBucketLifecycleConfigurationRequest{Request: req, Input: input, Copy: c.PutBucketLifecycleConfigurationRequest}
}

// PutBucketLifecycleConfigurationRequest is the request type for the
// PutBucketLifecycleConfiguration API operation.
type PutBucketLifecycleConfigurationRequest struct {
	*aws.Request
	Input *PutBucketLifecycleConfigurationInput
	Copy  func(*PutBucketLifecycleConfigurationInput) PutBucketLifecycleConfigurationRequest
}

// Send marshals and sends the PutBucketLifecycleConfiguration API request.
func (r PutBucketLifecycleConfigurationRequest) Send(ctx context.Context) (*PutBucketLifecycleConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBucketLifecycleConfigurationResponse{
		PutBucketLifecycleConfigurationOutput: r.Request.Data.(*PutBucketLifecycleConfigurationOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBucketLifecycleConfigurationResponse is the response type for the
// PutBucketLifecycleConfiguration API operation.
type PutBucketLifecycleConfigurationResponse struct {
	*PutBucketLifecycleConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBucketLifecycleConfiguration request.
func (r *PutBucketLifecycleConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
