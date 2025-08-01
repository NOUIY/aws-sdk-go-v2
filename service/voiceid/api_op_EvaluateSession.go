// Code generated by smithy-go-codegen DO NOT EDIT.

package voiceid

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/voiceid/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Evaluates a specified session based on audio data accumulated during a
// streaming Amazon Connect Voice ID call.
func (c *Client) EvaluateSession(ctx context.Context, params *EvaluateSessionInput, optFns ...func(*Options)) (*EvaluateSessionOutput, error) {
	if params == nil {
		params = &EvaluateSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EvaluateSession", params, optFns, c.addOperationEvaluateSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EvaluateSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EvaluateSessionInput struct {

	// The identifier of the domain where the session started.
	//
	// This member is required.
	DomainId *string

	// The session identifier, or name of the session, that you want to evaluate. In
	// Voice ID integration, this is the Contact-Id.
	//
	// This member is required.
	SessionNameOrId *string

	noSmithyDocumentSerde
}

type EvaluateSessionOutput struct {

	// Details resulting from the authentication process, such as authentication
	// decision and authentication score.
	AuthenticationResult *types.AuthenticationResult

	// The identifier of the domain that contains the session.
	DomainId *string

	// Details resulting from the fraud detection process, such as fraud detection
	// decision and risk score.
	FraudDetectionResult *types.FraudDetectionResult

	// The service-generated identifier of the session.
	SessionId *string

	// The client-provided name of the session.
	SessionName *string

	// The current status of audio streaming for this session. This field is useful to
	// infer next steps when the Authentication or Fraud Detection results are empty or
	// the decision is NOT_ENOUGH_SPEECH . In this situation, if the StreamingStatus
	// is ONGOING/PENDING_CONFIGURATION , it can mean that the client should call the
	// API again later, after Voice ID has enough audio to produce a result. If the
	// decision remains NOT_ENOUGH_SPEECH even after StreamingStatus is ENDED , it
	// means that the previously streamed session did not have enough speech to perform
	// evaluation, and a new streaming session is needed to try again.
	StreamingStatus types.StreamingStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEvaluateSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpEvaluateSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpEvaluateSession{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "EvaluateSession"); err != nil {
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
	if err = addOpEvaluateSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEvaluateSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEvaluateSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "EvaluateSession",
	}
}
