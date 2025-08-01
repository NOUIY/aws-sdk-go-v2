// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets usage details (for example, data storage) about a particular identity pool.
//
// This API can only be called with developer credentials. You cannot call this
// API with the temporary user credentials provided by Cognito Identity.
//
// DescribeIdentityPoolUsage The following examples have been edited for
// readability. POST / HTTP/1.1 CONTENT-TYPE: application/json X-AMZN-REQUESTID:
// 8dc0e749-c8cd-48bd-8520-da6be00d528b X-AMZ-TARGET:
// com.amazonaws.cognito.sync.model.AWSCognitoSyncService.DescribeIdentityPoolUsage
// HOST: cognito-sync.us-east-1.amazonaws.com:443 X-AMZ-DATE: 20141111T205737Z
// AUTHORIZATION: AWS4-HMAC-SHA256 Credential=,
// SignedHeaders=content-type;host;x-amz-date;x-amz-target;x-amzn-requestid,
// Signature=
//
// { "Operation": "com.amazonaws.cognito.sync.model#DescribeIdentityPoolUsage",
// "Service": "com.amazonaws.cognito.sync.model#AWSCognitoSyncService", "Input": {
// "IdentityPoolId": "IDENTITY_POOL_ID" } } 1.1 200 OK x-amzn-requestid:
// 8dc0e749-c8cd-48bd-8520-da6be00d528b content-type: application/json
// content-length: 271 date: Tue, 11 Nov 2014 20:57:37 GMT
//
// { "Output": { "__type":
// "com.amazonaws.cognito.sync.model#DescribeIdentityPoolUsageResponse",
// "IdentityPoolUsage": { "DataStorage": 0, "IdentityPoolId": "IDENTITY_POOL_ID",
// "LastModifiedDate": 1.413231134115E9, "SyncSessionsCount": null } }, "Version":
// "1.0" }
func (c *Client) DescribeIdentityPoolUsage(ctx context.Context, params *DescribeIdentityPoolUsageInput, optFns ...func(*Options)) (*DescribeIdentityPoolUsageOutput, error) {
	if params == nil {
		params = &DescribeIdentityPoolUsageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeIdentityPoolUsage", params, optFns, c.addOperationDescribeIdentityPoolUsageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeIdentityPoolUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request for usage information about the identity pool.
type DescribeIdentityPoolUsageInput struct {

	// A name-spaced GUID (for example,
	// us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito. GUID
	// generation is unique within a region.
	//
	// This member is required.
	IdentityPoolId *string

	noSmithyDocumentSerde
}

// Response to a successful DescribeIdentityPoolUsage request.
type DescribeIdentityPoolUsageOutput struct {

	// Information about the usage of the identity pool.
	IdentityPoolUsage *types.IdentityPoolUsage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeIdentityPoolUsageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeIdentityPoolUsage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeIdentityPoolUsage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeIdentityPoolUsage"); err != nil {
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
	if err = addOpDescribeIdentityPoolUsageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeIdentityPoolUsage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeIdentityPoolUsage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeIdentityPoolUsage",
	}
}
