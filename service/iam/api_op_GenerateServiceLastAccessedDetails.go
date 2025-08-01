// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates a report that includes details about when an IAM resource (user,
// group, role, or policy) was last used in an attempt to access Amazon Web
// Services services. Recent activity usually appears within four hours. IAM
// reports activity for at least the last 400 days, or less if your Region began
// supporting this feature within the last year. For more information, see [Regions where data is tracked]. For
// more information about services and actions for which action last accessed
// information is displayed, see [IAM action last accessed information services and actions].
//
// The service last accessed data includes all attempts to access an Amazon Web
// Services API, not just the successful ones. This includes all attempts that were
// made using the Amazon Web Services Management Console, the Amazon Web Services
// API through any of the SDKs, or any of the command line tools. An unexpected
// entry in the service last accessed data does not mean that your account has been
// compromised, because the request might have been denied. Refer to your
// CloudTrail logs as the authoritative source for information about all API calls
// and whether they were successful or denied access. For more information, see [Logging IAM events with CloudTrail]in
// the IAM User Guide.
//
// The GenerateServiceLastAccessedDetails operation returns a JobId . Use this
// parameter in the following operations to retrieve the following details from
// your report:
//
// [GetServiceLastAccessedDetails]
//   - – Use this operation for users, groups, roles, or policies to list every
//     Amazon Web Services service that the resource could access using permissions
//     policies. For each service, the response includes information about the most
//     recent access attempt.
//
// The JobId returned by GenerateServiceLastAccessedDetail must be used by the same
//
//	role within a session, or by the same user when used to call
//	GetServiceLastAccessedDetail .
//
// [GetServiceLastAccessedDetailsWithEntities]
//   - – Use this operation for groups and policies to list information about the
//     associated entities (users or roles) that attempted to access a specific Amazon
//     Web Services service.
//
// To check the status of the GenerateServiceLastAccessedDetails request, use the
// JobId parameter in the same operations and test the JobStatus response
// parameter.
//
// For additional information about the permissions policies that allow an
// identity (user, group, or role) to access specific services, use the [ListPoliciesGrantingServiceAccess]operation.
//
// Service last accessed data does not use other policy types when determining
// whether a resource could access a service. These other policy types include
// resource-based policies, access control lists, Organizations policies, IAM
// permissions boundaries, and STS assume role policies. It only applies
// permissions policy logic. For more about the evaluation of policy types, see [Evaluating policies]in
// the IAM User Guide.
//
// For more information about service and action last accessed data, see [Reducing permissions using service last accessed data] in the
// IAM User Guide.
//
// [Logging IAM events with CloudTrail]: https://docs.aws.amazon.com/IAM/latest/UserGuide/cloudtrail-integration.html
// [GetServiceLastAccessedDetails]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetServiceLastAccessedDetails.html
// [ListPoliciesGrantingServiceAccess]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_ListPoliciesGrantingServiceAccess.html
// [Reducing permissions using service last accessed data]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html
// [Regions where data is tracked]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html#access-advisor_tracking-period
// [Evaluating policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-basics
// [GetServiceLastAccessedDetailsWithEntities]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetServiceLastAccessedDetailsWithEntities.html
// [IAM action last accessed information services and actions]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor-action-last-accessed.html
func (c *Client) GenerateServiceLastAccessedDetails(ctx context.Context, params *GenerateServiceLastAccessedDetailsInput, optFns ...func(*Options)) (*GenerateServiceLastAccessedDetailsOutput, error) {
	if params == nil {
		params = &GenerateServiceLastAccessedDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GenerateServiceLastAccessedDetails", params, optFns, c.addOperationGenerateServiceLastAccessedDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GenerateServiceLastAccessedDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateServiceLastAccessedDetailsInput struct {

	// The ARN of the IAM resource (user, group, role, or managed policy) used to
	// generate information about when the resource was last used in an attempt to
	// access an Amazon Web Services service.
	//
	// This member is required.
	Arn *string

	// The level of detail that you want to generate. You can specify whether you want
	// to generate information about the last attempt to access services or actions. If
	// you specify service-level granularity, this operation generates only service
	// data. If you specify action-level granularity, it generates service and action
	// data. If you don't include this optional parameter, the operation generates
	// service data.
	Granularity types.AccessAdvisorUsageGranularityType

	noSmithyDocumentSerde
}

type GenerateServiceLastAccessedDetailsOutput struct {

	// The JobId that you can use in the [GetServiceLastAccessedDetails] or [GetServiceLastAccessedDetailsWithEntities] operations. The JobId returned by
	// GenerateServiceLastAccessedDetail must be used by the same role within a
	// session, or by the same user when used to call GetServiceLastAccessedDetail .
	//
	// [GetServiceLastAccessedDetails]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetServiceLastAccessedDetails.html
	// [GetServiceLastAccessedDetailsWithEntities]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetServiceLastAccessedDetailsWithEntities.html
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGenerateServiceLastAccessedDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGenerateServiceLastAccessedDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGenerateServiceLastAccessedDetails{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GenerateServiceLastAccessedDetails"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = addOpGenerateServiceLastAccessedDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateServiceLastAccessedDetails(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addInterceptAttempt(stack, options); err != nil {
		return err
	}
	if err = addInterceptExecution(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeSerialization(stack, options); err != nil {
		return err
	}
	if err = addInterceptAfterSerialization(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeSigning(stack, options); err != nil {
		return err
	}
	if err = addInterceptAfterSigning(stack, options); err != nil {
		return err
	}
	if err = addInterceptTransmit(stack, options); err != nil {
		return err
	}
	if err = addInterceptBeforeDeserialization(stack, options); err != nil {
		return err
	}
	if err = addInterceptAfterDeserialization(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGenerateServiceLastAccessedDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GenerateServiceLastAccessedDetails",
	}
}
