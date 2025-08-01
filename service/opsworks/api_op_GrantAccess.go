// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action can be used only with Windows stacks.
//
// Grants RDP access to a Windows instance for a specified time period.
func (c *Client) GrantAccess(ctx context.Context, params *GrantAccessInput, optFns ...func(*Options)) (*GrantAccessOutput, error) {
	if params == nil {
		params = &GrantAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GrantAccess", params, optFns, c.addOperationGrantAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GrantAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GrantAccessInput struct {

	// The instance's OpsWorks Stacks ID.
	//
	// This member is required.
	InstanceId *string

	// The length of time (in minutes) that the grant is valid. When the grant expires
	// at the end of this period, the user will no longer be able to use the
	// credentials to log in. If the user is logged in at the time, they are logged
	// out.
	ValidForInMinutes *int32

	noSmithyDocumentSerde
}

// Contains the response to a GrantAccess request.
type GrantAccessOutput struct {

	// A TemporaryCredential object that contains the data needed to log in to the
	// instance by RDP clients, such as the Microsoft Remote Desktop Connection.
	TemporaryCredential *types.TemporaryCredential

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGrantAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGrantAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGrantAccess{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GrantAccess"); err != nil {
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
	if err = addOpGrantAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGrantAccess(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGrantAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GrantAccess",
	}
}
