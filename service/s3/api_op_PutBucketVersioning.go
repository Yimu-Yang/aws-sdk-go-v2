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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketVersioningRequest
type PutBucketVersioningInput struct {
	_ struct{} `type:"structure" payload:"VersioningConfiguration"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The concatenation of the authentication device's serial number, a space,
	// and the value that is displayed on your authentication device.
	MFA *string `location:"header" locationName:"x-amz-mfa" type:"string"`

	// Describes the versioning state of an Amazon S3 bucket. For more information,
	// see PUT Bucket versioning (https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTVersioningStatus.html)
	// in the Amazon Simple Storage Service API Reference.
	//
	// VersioningConfiguration is a required field
	VersioningConfiguration *VersioningConfiguration `locationName:"VersioningConfiguration" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`
}

// String returns the string representation
func (s PutBucketVersioningInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBucketVersioningInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBucketVersioningInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.VersioningConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("VersioningConfiguration"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutBucketVersioningInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketVersioningInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.MFA != nil {
		v := *s.MFA

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-mfa", protocol.StringValue(v), metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.VersioningConfiguration != nil {
		v := s.VersioningConfiguration

		metadata := protocol.Metadata{XMLNamespaceURI: "http://s3.amazonaws.com/doc/2006-03-01/"}
		e.SetFields(protocol.PayloadTarget, "VersioningConfiguration", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketVersioningOutput
type PutBucketVersioningOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBucketVersioningOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBucketVersioningOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}
func (s *PutBucketVersioningOutput) unmarshalAWSXML(d *xml.Decoder, head xml.StartElement) (err error) {
	defer func() {
		if err != nil {
			*s = PutBucketVersioningOutput{}
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
					return fmt.Errorf("fail to UnmarshalAWSXML PutBucketVersioningOutput.%s, %s", name, err)
				}
			}
		}
	}
}

const opPutBucketVersioning = "PutBucketVersioning"

// PutBucketVersioningRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Sets the versioning state of an existing bucket. To set the versioning state,
// you must be the bucket owner.
//
//    // Example sending a request using PutBucketVersioningRequest.
//    req := client.PutBucketVersioningRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutBucketVersioning
func (c *Client) PutBucketVersioningRequest(input *PutBucketVersioningInput) PutBucketVersioningRequest {
	op := &aws.Operation{
		Name:       opPutBucketVersioning,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?versioning",
	}

	if input == nil {
		input = &PutBucketVersioningInput{}
	}

	req := c.newRequest(op, input, &PutBucketVersioningOutput{})
	req.Handlers.Unmarshal.Remove(restxml.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutBucketVersioningRequest{Request: req, Input: input, Copy: c.PutBucketVersioningRequest}
}

// PutBucketVersioningRequest is the request type for the
// PutBucketVersioning API operation.
type PutBucketVersioningRequest struct {
	*aws.Request
	Input *PutBucketVersioningInput
	Copy  func(*PutBucketVersioningInput) PutBucketVersioningRequest
}

// Send marshals and sends the PutBucketVersioning API request.
func (r PutBucketVersioningRequest) Send(ctx context.Context) (*PutBucketVersioningResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBucketVersioningResponse{
		PutBucketVersioningOutput: r.Request.Data.(*PutBucketVersioningOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBucketVersioningResponse is the response type for the
// PutBucketVersioning API operation.
type PutBucketVersioningResponse struct {
	*PutBucketVersioningOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBucketVersioning request.
func (r *PutBucketVersioningResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
