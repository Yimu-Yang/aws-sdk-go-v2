![Build Status](https://codebuild.us-west-2.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoiNmlHN1RaaXBIc3RmZzFCYjgydENqSENIaTZJazF0QTBWUkxhR2JoWnZLdG9BdU9nblpXbDk5S2xoYUhRcWl5dERFVklaMDRrUy9rY3l4cmJTRzJnNHJZPSIsIml2UGFyYW1ldGVyU3BlYyI6Inc4bW5GZzZNN1MreGl1Y3giLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master) [![Documentation](https://godoc.org/github.com/aws/aws-sdk-go-v2?status.svg)](https://godoc.org/github.com/aws/aws-sdk-go-v2)

# AWS SDK for Go v2

aws-sdk-go-v2 is the Developer Preview for the v2 of the AWS SDK for the Go programming language. 

Check out the [Issues] and [Projects] for design and updates being made to the SDK. The v2 SDK requires a minimum version of Go 1.9.

We'll be expanding out the [Issues] and [Projects] sections with additional changes to the SDK based on your feedback, and SDK's core's improvements. Check the the SDK's [CHANGE_LOG] for information about the latest updates to the SDK.

## Getting started

The best way to get started working with the SDK is to use `go get` to add the SDK to your Go Workspace or application using Go modules.

```sh
go get github.com/aws/aws-sdk-go-v2
```

Without Go Modules, or in a GOPATH with Go 1.11 or 1.12 use the `/...` suffix on the `go get` to retrieve all of the SDK's depdnencies.

```sh
go get github.com/aws/aws-sdk-go-v2/...
```

### Hello AWS

This example shows how you can use the v2 SDK to make an API request using the SDK's [Amazon DynamoDB] client.

```go
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	// Set the AWS Region that the service clients should use
	cfg.Region = endpoints.UsWest2RegionID

	// Using the Config value, create the DynamoDB client
	svc := dynamodb.New(cfg)

	// Build the request with its input parameters
	req := svc.DescribeTableRequest(&dynamodb.DescribeTableInput{
		TableName: aws.String("myTable"),
	})

	// Send the request, and get the response or error back
	resp, err := req.Send(context.Background())
	if err != nil {
		panic("failed to describe table, "+err.Error())
	}

	fmt.Println("Response", resp)
}
```

## Feedback and contributing

The v2 SDK will use GitHub [Issues] to track feature requests and issues with the SDK. In addition, we'll use GitHub [Projects] to track large tasks spanning multiple pull requests, such as refactoring the SDK's internal request lifecycle. You can provide feedback to us in several ways. 

**GitHub issues**. To provide feedback or report bugs, file GitHub [Issues] on the SDK. This is the preferred mechanism to give feedback so that other users can engage in the conversation, +1 issues, etc. Issues you open will be evaluated, and included in our roadmap for the GA launch.

**Gitter channel**. For more informal discussions or general feedback, check out our [Gitter channel] for the SDK. The [Gitter channel] is also a great place to ask general questions, and find help to get started with the 2.0 SDK Developer Preview.

**Contributing**. You can open pull requests for fixes or additions to the AWS SDK for Go 2.0 Developer Preview release. All pull requests must be submitted under the Apache 2.0 license and will be reviewed by an SDK team member before being merged in. Accompanying unit tests, where possible, are appreciated.

## License

This SDK is distributed under the
[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0),
see LICENSE.txt and NOTICE.txt for more information.

[Dep]: https://github.com/golang/dep
[Issues]: https://github.com/aws/aws-sdk-go-v2/issues
[Projects]: https://github.com/aws/aws-sdk-go-v2/projects
[CHANGE_LOG]: https://github.com/aws/aws-sdk-go-v2/blob/master/CHANGELOG.md
[Amazon DynamoDB]: https://aws.amazon.com/dynamodb/
[Gitter channel]: https://gitter.im/aws/aws-sdk-go-v2
