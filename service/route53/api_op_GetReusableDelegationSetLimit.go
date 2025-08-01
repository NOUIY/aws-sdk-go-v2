// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the maximum number of hosted zones that you can associate with the
// specified reusable delegation set.
//
// For the default limit, see [Limits] in the Amazon Route 53 Developer Guide. To request
// a higher limit, [open a case].
//
// [Limits]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html
// [open a case]: https://console.aws.amazon.com/support/home#/case/create?issueType=service-limit-increase&limitType=service-code-route53
func (c *Client) GetReusableDelegationSetLimit(ctx context.Context, params *GetReusableDelegationSetLimitInput, optFns ...func(*Options)) (*GetReusableDelegationSetLimitOutput, error) {
	if params == nil {
		params = &GetReusableDelegationSetLimitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReusableDelegationSetLimit", params, optFns, c.addOperationGetReusableDelegationSetLimitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReusableDelegationSetLimitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the request to create a hosted
// zone.
type GetReusableDelegationSetLimitInput struct {

	// The ID of the delegation set that you want to get the limit for.
	//
	// This member is required.
	DelegationSetId *string

	// Specify MAX_ZONES_BY_REUSABLE_DELEGATION_SET to get the maximum number of
	// hosted zones that you can associate with the specified reusable delegation set.
	//
	// This member is required.
	Type types.ReusableDelegationSetLimitType

	noSmithyDocumentSerde
}

// A complex type that contains the requested limit.
type GetReusableDelegationSetLimitOutput struct {

	// The current number of hosted zones that you can associate with the specified
	// reusable delegation set.
	//
	// This member is required.
	Count int64

	// The current setting for the limit on hosted zones that you can associate with
	// the specified reusable delegation set.
	//
	// This member is required.
	Limit *types.ReusableDelegationSetLimit

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetReusableDelegationSetLimitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetReusableDelegationSetLimit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetReusableDelegationSetLimit{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetReusableDelegationSetLimit"); err != nil {
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
	if err = addOpGetReusableDelegationSetLimitValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReusableDelegationSetLimit(options.Region), middleware.Before); err != nil {
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
	if err = addSanitizeURLMiddleware(stack); err != nil {
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

func newServiceMetadataMiddleware_opGetReusableDelegationSetLimit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetReusableDelegationSetLimit",
	}
}
