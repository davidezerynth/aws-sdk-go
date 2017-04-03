// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/restjson"
)

// The WorkDocs API is designed for the following use cases:
//
//    * File Migration: File migration applications are supported for users
//    who want to migrate their files from an on-premise or off-premise file
//    system or service. Users can insert files into a user directory structure,
//    as well as allow for basic metadata changes, such as modifications to
//    the permissions of files.
//
//    * Security: Support security applications are supported for users who
//    have additional security needs, such as anti-virus or data loss prevention.
//    The APIs, in conjunction with Amazon CloudTrail, allow these applications
//    to detect when changes occur in Amazon WorkDocs, so the application can
//    take the necessary actions and replace the target file. The application
//    can also choose to email the user if the target file violates the policy.
//
//    * eDiscovery/Analytics: General administrative applications are supported,
//    such as eDiscovery and analytics. These applications can choose to mimic
//    and/or record the actions in an Amazon WorkDocs site, in conjunction with
//    Amazon CloudTrails, to replicate data for eDiscovery, backup, or analytical
//    applications.
//
// All Amazon WorkDocs APIs are Amazon authenticated, certificate-signed APIs.
// They not only require the use of the AWS SDK, but also allow for the exclusive
// use of IAM users and roles to help facilitate access, trust, and permission
// policies. By creating a role and allowing an IAM user to access the Amazon
// WorkDocs site, the IAM user gains full administrative visibility into the
// entire Amazon WorkDocs site (or as set in the IAM policy). This includes,
// but is not limited to, the ability to modify file permissions and upload
// any file to any user. This allows developers to perform the three use cases
// above, as well as give users the ability to grant access on a selective basis
// using the IAM model.
// The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01
type WorkDocs struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "workdocs"  // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the WorkDocs client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a WorkDocs client from just a session.
//     svc := workdocs.New(mySession)
//
//     // Create a WorkDocs client with additional configuration
//     svc := workdocs.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *WorkDocs {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *WorkDocs {
	svc := &WorkDocs{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2016-05-01",
				JSONVersion:   "1.1",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restjson.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a WorkDocs operation and runs any
// custom request initialization.
func (c *WorkDocs) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}