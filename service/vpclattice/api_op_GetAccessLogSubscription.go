// Code generated by smithy-go-codegen DO NOT EDIT.

package vpclattice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/vpclattice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about the specified access log subscription.
func (c *Client) GetAccessLogSubscription(ctx context.Context, params *GetAccessLogSubscriptionInput, optFns ...func(*Options)) (*GetAccessLogSubscriptionOutput, error) {
	if params == nil {
		params = &GetAccessLogSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAccessLogSubscription", params, optFns, c.addOperationGetAccessLogSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAccessLogSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAccessLogSubscriptionInput struct {

	// The ID or ARN of the access log subscription.
	//
	// This member is required.
	AccessLogSubscriptionIdentifier *string

	noSmithyDocumentSerde
}

type GetAccessLogSubscriptionOutput struct {

	// The Amazon Resource Name (ARN) of the access log subscription.
	//
	// This member is required.
	Arn *string

	// The date and time that the access log subscription was created, in ISO-8601
	// format.
	//
	// This member is required.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the access log destination.
	//
	// This member is required.
	DestinationArn *string

	// The ID of the access log subscription.
	//
	// This member is required.
	Id *string

	// The date and time that the access log subscription was last updated, in
	// ISO-8601 format.
	//
	// This member is required.
	LastUpdatedAt *time.Time

	// The Amazon Resource Name (ARN) of the service network or service.
	//
	// This member is required.
	ResourceArn *string

	// The ID of the service network or service.
	//
	// This member is required.
	ResourceId *string

	// The log type for the service network.
	ServiceNetworkLogType types.ServiceNetworkLogType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAccessLogSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAccessLogSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAccessLogSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAccessLogSubscription"); err != nil {
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
	if err = addOpGetAccessLogSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAccessLogSubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAccessLogSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAccessLogSubscription",
	}
}
