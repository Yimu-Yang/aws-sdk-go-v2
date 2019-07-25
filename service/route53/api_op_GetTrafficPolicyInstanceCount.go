// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

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

// Request to get the number of traffic policy instances that are associated
// with the current AWS account.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/GetTrafficPolicyInstanceCountRequest
type GetTrafficPolicyInstanceCountInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetTrafficPolicyInstanceCountInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetTrafficPolicyInstanceCountInput) MarshalFields(e protocol.FieldEncoder) error {

	return nil
}

// A complex type that contains information about the resource record sets that
// Amazon Route 53 created based on a specified traffic policy.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/GetTrafficPolicyInstanceCountResponse
type GetTrafficPolicyInstanceCountOutput struct {
	_ struct{} `type:"structure"`

	// The number of traffic policy instances that are associated with the current
	// AWS account.
	//
	// TrafficPolicyInstanceCount is a required field
	TrafficPolicyInstanceCount *int64 `type:"integer" required:"true"`
}

// String returns the string representation
func (s GetTrafficPolicyInstanceCountOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetTrafficPolicyInstanceCountOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.TrafficPolicyInstanceCount != nil {
		v := *s.TrafficPolicyInstanceCount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TrafficPolicyInstanceCount", protocol.Int64Value(v), metadata)
	}
	return nil
}

// UnmarshalAWSXML decodes the AWS API shape using the passed in *xml.Decoder.
func (s *GetTrafficPolicyInstanceCountOutput) UnmarshalAWSXML(d *xml.Decoder) (err error) {
	defer func() {
		if err != nil {
			*s = GetTrafficPolicyInstanceCountOutput{}
		}
	}()
	for {
		tok, err := d.Token()
		if tok == nil || err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("fail to UnmarshalAWSXML GetTrafficPolicyInstanceCountOutput, %s", err)
		}
		start, ok := tok.(xml.StartElement)
		if !ok {
			continue
		}
		err = s.unmarshalAWSXML(d, start)
		if err != nil {
			return fmt.Errorf("fail to UnmarshalAWSXML GetTrafficPolicyInstanceCountOutput, %s", err)
		}
		return nil
	}
}

func (s *GetTrafficPolicyInstanceCountOutput) unmarshalAWSXML(d *xml.Decoder, head xml.StartElement) (err error) {
	defer func() {
		if err != nil {
			*s = GetTrafficPolicyInstanceCountOutput{}
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
			case "TrafficPolicyInstanceCount":
				tok, err = d.Token()
				if tok == nil || err != nil {
					return fmt.Errorf("fail to UnmarshalAWSXML GetTrafficPolicyInstanceCountOutput.%s, %s", name, err)
				}
				v, _ := tok.(xml.CharData)
				value, _ := strconv.ParseInt(string(v), 10, 64)
				s.TrafficPolicyInstanceCount = &value
			default:
				err := d.Skip()
				if err != nil {
					return fmt.Errorf("fail to UnmarshalAWSXML GetTrafficPolicyInstanceCountOutput.%s, %s", name, err)
				}
			}
		}
	}
}

const opGetTrafficPolicyInstanceCount = "GetTrafficPolicyInstanceCount"

// GetTrafficPolicyInstanceCountRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Gets the number of traffic policy instances that are associated with the
// current AWS account.
//
//    // Example sending a request using GetTrafficPolicyInstanceCountRequest.
//    req := client.GetTrafficPolicyInstanceCountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/GetTrafficPolicyInstanceCount
func (c *Client) GetTrafficPolicyInstanceCountRequest(input *GetTrafficPolicyInstanceCountInput) GetTrafficPolicyInstanceCountRequest {
	op := &aws.Operation{
		Name:       opGetTrafficPolicyInstanceCount,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/trafficpolicyinstancecount",
	}

	if input == nil {
		input = &GetTrafficPolicyInstanceCountInput{}
	}

	req := c.newRequest(op, input, &GetTrafficPolicyInstanceCountOutput{})
	return GetTrafficPolicyInstanceCountRequest{Request: req, Input: input, Copy: c.GetTrafficPolicyInstanceCountRequest}
}

// GetTrafficPolicyInstanceCountRequest is the request type for the
// GetTrafficPolicyInstanceCount API operation.
type GetTrafficPolicyInstanceCountRequest struct {
	*aws.Request
	Input *GetTrafficPolicyInstanceCountInput
	Copy  func(*GetTrafficPolicyInstanceCountInput) GetTrafficPolicyInstanceCountRequest
}

// Send marshals and sends the GetTrafficPolicyInstanceCount API request.
func (r GetTrafficPolicyInstanceCountRequest) Send(ctx context.Context) (*GetTrafficPolicyInstanceCountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTrafficPolicyInstanceCountResponse{
		GetTrafficPolicyInstanceCountOutput: r.Request.Data.(*GetTrafficPolicyInstanceCountOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetTrafficPolicyInstanceCountResponse is the response type for the
// GetTrafficPolicyInstanceCount API operation.
type GetTrafficPolicyInstanceCountResponse struct {
	*GetTrafficPolicyInstanceCountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTrafficPolicyInstanceCount request.
func (r *GetTrafficPolicyInstanceCountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
