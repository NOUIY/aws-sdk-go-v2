// Code generated by smithy-go-codegen DO NOT EDIT.

package partnercentralselling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/partnercentralselling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Use this action to retrieve the engagement record for a given
// EngagementIdentifier .
func (c *Client) GetEngagement(ctx context.Context, params *GetEngagementInput, optFns ...func(*Options)) (*GetEngagementOutput, error) {
	if params == nil {
		params = &GetEngagementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEngagement", params, optFns, c.addOperationGetEngagementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEngagementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEngagementInput struct {

	// Specifies the catalog related to the engagement request. Valid values are AWS
	// and Sandbox .
	//
	// This member is required.
	Catalog *string

	// Specifies the identifier of the Engagement record to retrieve.
	//
	// This member is required.
	Identifier *string

	noSmithyDocumentSerde
}

type GetEngagementOutput struct {

	// The Amazon Resource Name (ARN) of the engagement retrieved.
	Arn *string

	// A list of context objects associated with the engagement. Each context provides
	// additional information related to the Engagement, such as customer projects or
	// documents.
	Contexts []types.EngagementContextDetails

	// The date and time when the Engagement was created, presented in ISO 8601 format
	// (UTC). For example: "2023-05-01T20:37:46Z". This timestamp helps track the
	// lifecycle of the Engagement.
	CreatedAt *time.Time

	// The AWS account ID of the user who originally created the engagement. This
	// field helps in tracking the origin of the engagement.
	CreatedBy *string

	// A more detailed description of the engagement. This provides additional context
	// or information about the engagement's purpose or scope.
	Description *string

	// The unique resource identifier of the engagement retrieved.
	Id *string

	// Specifies the current count of members participating in the Engagement. This
	// count includes all active members regardless of their roles or permissions
	// within the Engagement.
	MemberCount *int32

	// The title of the engagement. It provides a brief, descriptive name for the
	// engagement that is meaningful and easily recognizable.
	Title *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEngagementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetEngagement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetEngagement{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetEngagement"); err != nil {
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
	if err = addOpGetEngagementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEngagement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEngagement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetEngagement",
	}
}
