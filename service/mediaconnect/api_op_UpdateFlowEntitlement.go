// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Updates an entitlement. You can change an entitlement's description,
//
// subscribers, and encryption. If you change the subscribers, the service will
// remove the outputs that are are used by the subscribers that are removed.
func (c *Client) UpdateFlowEntitlement(ctx context.Context, params *UpdateFlowEntitlementInput, optFns ...func(*Options)) (*UpdateFlowEntitlementOutput, error) {
	if params == nil {
		params = &UpdateFlowEntitlementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFlowEntitlement", params, optFns, c.addOperationUpdateFlowEntitlementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFlowEntitlementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFlowEntitlementInput struct {

	//  The Amazon Resource Name (ARN) of the entitlement that you want to update.
	//
	// This member is required.
	EntitlementArn *string

	//  The ARN of the flow that is associated with the entitlement that you want to
	// update.
	//
	// This member is required.
	FlowArn *string

	//  A description of the entitlement. This description appears only on the
	// MediaConnect console and will not be seen by the subscriber or end user.
	Description *string

	//  The type of encryption that will be used on the output associated with this
	// entitlement. Allowable encryption types: static-key, speke.
	Encryption *types.UpdateEncryption

	//  An indication of whether you want to enable the entitlement to allow access,
	// or disable it to stop streaming content to the subscriber’s flow temporarily. If
	// you don’t specify the entitlementStatus field in your request, MediaConnect
	// leaves the value unchanged.
	EntitlementStatus types.EntitlementStatus

	//  The Amazon Web Services account IDs that you want to share your content with.
	// The receiving accounts (subscribers) will be allowed to create their own flow
	// using your content as the source.
	Subscribers []string

	noSmithyDocumentSerde
}

type UpdateFlowEntitlementOutput struct {

	//  The new configuration of the entitlement that you updated.
	Entitlement *types.Entitlement

	//  The ARN of the flow that this entitlement was granted on.
	FlowArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFlowEntitlementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFlowEntitlement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFlowEntitlement{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateFlowEntitlement"); err != nil {
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
	if err = addOpUpdateFlowEntitlementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFlowEntitlement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFlowEntitlement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateFlowEntitlement",
	}
}
