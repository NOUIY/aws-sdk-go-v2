// Code generated by smithy-go-codegen DO NOT EDIT.

package ivsrealtime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts replicating a publishing participant from a source stage to a
// destination stage.
func (c *Client) StartParticipantReplication(ctx context.Context, params *StartParticipantReplicationInput, optFns ...func(*Options)) (*StartParticipantReplicationOutput, error) {
	if params == nil {
		params = &StartParticipantReplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartParticipantReplication", params, optFns, c.addOperationStartParticipantReplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartParticipantReplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartParticipantReplicationInput struct {

	// ARN of the stage to which the participant will be replicated.
	//
	// This member is required.
	DestinationStageArn *string

	// Participant ID of the publisher that will be replicated. This is assigned by
	// IVS and returned by CreateParticipantTokenor the jti (JWT ID) used to [create a self signed token].
	//
	// [create a self signed token]: https://docs.aws.amazon.com/ivs/latest/RealTimeUserGuide/getting-started-distribute-tokens.html#getting-started-distribute-tokens-self-signed
	//
	// This member is required.
	ParticipantId *string

	// ARN of the stage where the participant is publishing.
	//
	// This member is required.
	SourceStageArn *string

	// Application-provided attributes to set on the replicated participant in the
	// destination stage. Map keys and values can contain UTF-8 encoded text. The
	// maximum length of this field is 1 KB total. This field is exposed to all stage
	// participants and should not be used for personally identifying, confidential, or
	// sensitive information.
	//
	// These attributes are merged with any attributes set for this participant when
	// creating the token. If there is overlap in keys, the values in these attributes
	// are replaced.
	Attributes map[string]string

	// If the participant disconnects and then reconnects within the specified
	// interval, replication will continue to be ACTIVE . Default: 0.
	ReconnectWindowSeconds *int32

	noSmithyDocumentSerde
}

type StartParticipantReplicationOutput struct {

	//
	AccessControlAllowOrigin *string

	//
	AccessControlExposeHeaders *string

	//
	CacheControl *string

	//
	ContentSecurityPolicy *string

	//
	StrictTransportSecurity *string

	//
	XContentTypeOptions *string

	//
	XFrameOptions *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartParticipantReplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartParticipantReplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartParticipantReplication{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartParticipantReplication"); err != nil {
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
	if err = addOpStartParticipantReplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartParticipantReplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartParticipantReplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartParticipantReplication",
	}
}
