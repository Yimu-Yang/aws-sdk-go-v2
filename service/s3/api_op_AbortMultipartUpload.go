// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/AbortMultipartUploadRequest
type AbortMultipartUploadInput struct {
	_ struct{} `type:"structure"`

	// Name of the bucket to which the multipart upload was initiated.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Key of the object for which the multipart upload was initiated.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// Upload ID that identifies the multipart upload.
	//
	// UploadId is a required field
	UploadId *string `location:"querystring" locationName:"uploadId" type:"string" required:"true"`
}

// String returns the string representation
func (s AbortMultipartUploadInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AbortMultipartUploadInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AbortMultipartUploadInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.UploadId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UploadId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *AbortMultipartUploadInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AbortMultipartUploadInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.RequestPayer) > 0 {
		v := s.RequestPayer

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-payer", v, metadata)
	}
	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Key", protocol.StringValue(v), metadata)
	}
	if s.UploadId != nil {
		v := *s.UploadId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "uploadId", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/AbortMultipartUploadOutput
type AbortMultipartUploadOutput struct {
	_ struct{} `type:"structure"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`
}

// String returns the string representation
func (s AbortMultipartUploadOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AbortMultipartUploadOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.RequestCharged) > 0 {
		v := s.RequestCharged

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-request-charged", v, metadata)
	}
	return nil
}
func (s *AbortMultipartUploadOutput) unmarshalAWSXML(d *xml.Decoder, head xml.StartElement) (err error) {
	defer func() {
		if err != nil {
			*s = AbortMultipartUploadOutput{}
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
			case "x-amz-request-charged":
				tok, err = d.Token()
				if tok == nil || err != nil {
					return fmt.Errorf("fail to UnmarshalAWSXML AbortMultipartUploadOutput.%s, %s", name, err)
				}
				v, _ := tok.(xml.CharData)
				value := RequestCharged(v)
				s.RequestCharged = value
			default:
				err := d.Skip()
				if err != nil {
					return fmt.Errorf("fail to UnmarshalAWSXML AbortMultipartUploadOutput.%s, %s", name, err)
				}
			}
		}
	}
}

// UnmarshalAWSREST decodes the AWS API shape using the passed in *http.Response.
func (s *AbortMultipartUploadOutput) UnmarshalAWSREST(r *http.Response) (err error) {
	defer func() {
		if err != nil {
			*s = AbortMultipartUploadOutput{}
		}
	}()
	for k, v := range r.Header {
		switch {
		case strings.EqualFold(k, "x-amz-request-charged"):
			value := RequestCharged(v[0])
			s.RequestCharged = value
		}
	}
	return nil
}

const opAbortMultipartUpload = "AbortMultipartUpload"

// AbortMultipartUploadRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Aborts a multipart upload.
//
// To verify that all parts have been removed, so you don't get charged for
// the part storage, you should call the List Parts operation and ensure the
// parts list is empty.
//
//    // Example sending a request using AbortMultipartUploadRequest.
//    req := client.AbortMultipartUploadRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/AbortMultipartUpload
func (c *Client) AbortMultipartUploadRequest(input *AbortMultipartUploadInput) AbortMultipartUploadRequest {
	op := &aws.Operation{
		Name:       opAbortMultipartUpload,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}/{Key+}",
	}

	if input == nil {
		input = &AbortMultipartUploadInput{}
	}

	req := c.newRequest(op, input, &AbortMultipartUploadOutput{})
	return AbortMultipartUploadRequest{Request: req, Input: input, Copy: c.AbortMultipartUploadRequest}
}

// AbortMultipartUploadRequest is the request type for the
// AbortMultipartUpload API operation.
type AbortMultipartUploadRequest struct {
	*aws.Request
	Input *AbortMultipartUploadInput
	Copy  func(*AbortMultipartUploadInput) AbortMultipartUploadRequest
}

// Send marshals and sends the AbortMultipartUpload API request.
func (r AbortMultipartUploadRequest) Send(ctx context.Context) (*AbortMultipartUploadResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AbortMultipartUploadResponse{
		AbortMultipartUploadOutput: r.Request.Data.(*AbortMultipartUploadOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AbortMultipartUploadResponse is the response type for the
// AbortMultipartUpload API operation.
type AbortMultipartUploadResponse struct {
	*AbortMultipartUploadOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AbortMultipartUpload request.
func (r *AbortMultipartUploadResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
