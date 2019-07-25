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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketNotificationConfigurationRequest
type PutBucketNotificationConfigurationInput struct {
	_ struct{} `type:"structure" payload:"NotificationConfiguration"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// A container for specifying the notification configuration of the bucket.
	// If this element is empty, notifications are turned off for the bucket.
	//
	// NotificationConfiguration is a required field
	NotificationConfiguration *NotificationConfiguration `locationName:"NotificationConfiguration" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`
}

// String returns the string representation
func (s PutBucketNotificationConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketNotificationConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketNotificationConfigurationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.NotificationConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("NotificationConfiguration"))
	}
	if s.NotificationConfiguration != nil {
		if err := s.NotificationConfiguration.Validate(); err != nil {
			invalidParams.AddNested("NotificationConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketNotificationConfigurationInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketNotificationConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.NotificationConfiguration != nil {
		v := s.NotificationConfiguration

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "NotificationConfiguration", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketNotificationConfigurationOutput
type PutBucketNotificationConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketNotificationConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketNotificationConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}
func (s *PutBucketNotificationConfigurationOutput) unmarshalAWSXML(d *xml.Decoder, head xml.StartElement) (err error) {
	defer func() {
		if err != nil {
			*s = PutBucketNotificationConfigurationOutput{}
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
					return fmt.Errorf("fail to UnmarshalAWSXML PutBucketNotificationConfigurationOutput.%s, %s", name, err)
				}
			}
		}
	}
}

const opPutBucketNotificationConfiguration = "PutBucketNotificationConfiguration"

// PutBucketNotificationConfigurationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Enables notifications of specified events for a bucket.
//
//    // Example sending a request using PutBucketNotificationConfigurationRequest.
//    req := client.PutBucketNotificationConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketNotificationConfiguration
func (c *Client) PutBucketNotificationConfigurationRequest(input *PutBucketNotificationConfigurationInput) PutBucketNotificationConfigurationRequest {
	op := &aws.Operation{
		Name:       opPutBucketNotificationConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?notification",
	}

	if input == nil {
		input = &PutBucketNotificationConfigurationInput{}
	}

	req := c.newRequest(op, input, &PutBucketNotificationConfigurationOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutBucketNotificationConfigurationRequest{Request: req, Input: input, Copy: c.PutBucketNotificationConfigurationRequest}
}

// PutBucketNotificationConfigurationRequest is the request type for the
// PutBucketNotificationConfiguration API operation.
type PutBucketNotificationConfigurationRequest struct {
	*aws.Request
	Input *PutBucketNotificationConfigurationInput
	Copy  func(*PutBucketNotificationConfigurationInput) PutBucketNotificationConfigurationRequest
}

// Send marshals and sends the PutBucketNotificationConfiguration API request.
func (r PutBucketNotificationConfigurationRequest) Send(ctx context.Context) (*PutBucketNotificationConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBucketNotificationConfigurationResponse{
		PutBucketNotificationConfigurationOutput: r.Request.Data.(*PutBucketNotificationConfigurationOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBucketNotificationConfigurationResponse is the response type for the
// PutBucketNotificationConfiguration API operation.
type PutBucketNotificationConfigurationResponse struct {
	*PutBucketNotificationConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBucketNotificationConfiguration request.
func (r *PutBucketNotificationConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
