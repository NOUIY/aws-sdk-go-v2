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

// Lists datasets for an identity. With Amazon Cognito Sync, each identity has
// access only to its own data. Thus, the credentials used to make this API call
// need to have access to the identity data.
//
// ListDatasets can be called with temporary user credentials provided by Cognito
// Identity or with developer credentials. You should use the Cognito Identity
// credentials to make this API call.
//
// ListDatasets The following examples have been edited for readability. POST /
// HTTP/1.1 CONTENT-TYPE: application/json X-AMZN-REQUESTID:
// 15225768-209f-4078-aaed-7494ace9f2db X-AMZ-TARGET:
// com.amazonaws.cognito.sync.model.AWSCognitoSyncService.ListDatasets HOST:
// cognito-sync.us-east-1.amazonaws.com:443 X-AMZ-DATE: 20141111T215640Z
// AUTHORIZATION: AWS4-HMAC-SHA256 Credential=,
// SignedHeaders=content-type;host;x-amz-date;x-amz-target;x-amzn-requestid,
// Signature= { "Operation": "com.amazonaws.cognito.sync.model#ListDatasets",
// "Service": "com.amazonaws.cognito.sync.model#AWSCognitoSyncService", "Input": {
// "IdentityPoolId": "IDENTITY_POOL_ID", "IdentityId": "IDENTITY_ID", "MaxResults":
// "3" } } 1.1 200 OK x-amzn-requestid: 15225768-209f-4078-aaed-7494ace9f2db,
// 15225768-209f-4078-aaed-7494ace9f2db content-type: application/json
// content-length: 355 date: Tue, 11 Nov 2014 21:56:40 GMT
//
// { "Output": { "__type":
// "com.amazonaws.cognito.sync.model#ListDatasetsResponse", "Count": 1, "Datasets":
// [ { "CreationDate": 1.412974057151E9, "DataStorage": 16, "DatasetName":
// "my_list", "IdentityId": "IDENTITY_ID", "LastModifiedBy": "123456789012",
// "LastModifiedDate": 1.412974057244E9, "NumRecords": 1 }], "NextToken": null },
// "Version": "1.0" }
func (c *Client) ListDatasets(ctx context.Context, params *ListDatasetsInput, optFns ...func(*Options)) (*ListDatasetsOutput, error) {
	if params == nil {
		params = &ListDatasetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDatasets", params, optFns, c.addOperationListDatasetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDatasetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request for a list of datasets for an identity.
type ListDatasetsInput struct {

	// A name-spaced GUID (for example,
	// us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito. GUID
	// generation is unique within a region.
	//
	// This member is required.
	IdentityId *string

	// A name-spaced GUID (for example,
	// us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito. GUID
	// generation is unique within a region.
	//
	// This member is required.
	IdentityPoolId *string

	// The maximum number of results to be returned.
	MaxResults *int32

	// A pagination token for obtaining the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

// Returned for a successful ListDatasets request.
type ListDatasetsOutput struct {

	// Number of datasets returned.
	Count int32

	// A set of datasets.
	Datasets []types.Dataset

	// A pagination token for obtaining the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDatasetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDatasets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDatasets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDatasets"); err != nil {
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
	if err = addOpListDatasetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDatasets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDatasets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDatasets",
	}
}
