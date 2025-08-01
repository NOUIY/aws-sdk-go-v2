// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the specified environment profile in Amazon DataZone.
func (c *Client) UpdateEnvironmentProfile(ctx context.Context, params *UpdateEnvironmentProfileInput, optFns ...func(*Options)) (*UpdateEnvironmentProfileOutput, error) {
	if params == nil {
		params = &UpdateEnvironmentProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEnvironmentProfile", params, optFns, c.addOperationUpdateEnvironmentProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEnvironmentProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEnvironmentProfileInput struct {

	// The identifier of the Amazon DataZone domain in which an environment profile is
	// to be updated.
	//
	// This member is required.
	DomainIdentifier *string

	// The identifier of the environment profile that is to be updated.
	//
	// This member is required.
	Identifier *string

	// The Amazon Web Services account in which a specified environment profile is to
	// be udpated.
	AwsAccountId *string

	// The Amazon Web Services Region in which a specified environment profile is to
	// be updated.
	AwsAccountRegion *string

	// The description to be updated as part of the UpdateEnvironmentProfile action.
	Description *string

	// The name to be updated as part of the UpdateEnvironmentProfile action.
	Name *string

	// The user parameters to be updated as part of the UpdateEnvironmentProfile
	// action.
	UserParameters []types.EnvironmentParameter

	noSmithyDocumentSerde
}

type UpdateEnvironmentProfileOutput struct {

	// The Amazon DataZone user who created the environment profile.
	//
	// This member is required.
	CreatedBy *string

	// The identifier of the Amazon DataZone domain in which the environment profile
	// is to be updated.
	//
	// This member is required.
	DomainId *string

	// The identifier of the blueprint of the environment profile that is to be
	// updated.
	//
	// This member is required.
	EnvironmentBlueprintId *string

	// The identifier of the environment profile that is to be udpated.
	//
	// This member is required.
	Id *string

	// The name to be updated as part of the UpdateEnvironmentProfile action.
	//
	// This member is required.
	Name *string

	// The Amazon Web Services account in which a specified environment profile is to
	// be udpated.
	AwsAccountId *string

	// The Amazon Web Services Region in which a specified environment profile is to
	// be updated.
	AwsAccountRegion *string

	// The timestamp of when the environment profile was created.
	CreatedAt *time.Time

	// The description to be updated as part of the UpdateEnvironmentProfile action.
	Description *string

	// The identifier of the project of the environment profile that is to be updated.
	ProjectId *string

	// The timestamp of when the environment profile was updated.
	UpdatedAt *time.Time

	// The user parameters to be updated as part of the UpdateEnvironmentProfile
	// action.
	UserParameters []types.CustomParameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEnvironmentProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateEnvironmentProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateEnvironmentProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateEnvironmentProfile"); err != nil {
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
	if err = addOpUpdateEnvironmentProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEnvironmentProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEnvironmentProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateEnvironmentProfile",
	}
}
