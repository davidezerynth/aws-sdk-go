// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ebs provides the client and types for making API
// requests to Amazon Elastic Block Store.
//
// You can use the Amazon Elastic Block Store (EBS) direct APIs to directly
// read the data on your EBS snapshots, and identify the difference between
// two snapshots. You can view the details of blocks in an EBS snapshot, compare
// the block difference between two snapshots, and directly access the data
// in a snapshot. If you're an independent software vendor (ISV) who offers
// backup services for EBS, the EBS direct APIs make it easier and more cost-effective
// to track incremental changes on your EBS volumes via EBS snapshots. This
// can be done without having to create new volumes from EBS snapshots.
//
// This API reference provides detailed information about the actions, data
// types, parameters, and errors of the EBS direct APIs. For more information
// about the elements that make up the EBS direct APIs, and examples of how
// to use them effectively, see Accessing the Contents of an EBS Snapshot (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-accessing-snapshot.html)
// in the Amazon Elastic Compute Cloud User Guide. For more information about
// the supported AWS Regions, endpoints, and service quotas for the EBS direct
// APIs, see Amazon Elastic Block Store Endpoints and Quotas (https://docs.aws.amazon.com/general/latest/gr/ebs-service.html)
// in the AWS General Reference.
//
// See https://docs.aws.amazon.com/goto/WebAPI/ebs-2019-11-02 for more information on this service.
//
// See ebs package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/ebs/
//
// Using the Client
//
// To contact Amazon Elastic Block Store with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon Elastic Block Store client EBS for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/ebs/#New
package ebs
