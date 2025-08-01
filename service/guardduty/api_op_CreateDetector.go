// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a single GuardDuty detector. A detector is a resource that represents
// the GuardDuty service. To start using GuardDuty, you must create a detector in
// each Region where you enable the service. You can have only one detector per
// account per Region. All data sources are enabled in a new detector by default.
//
//   - When you don't specify any features , with an exception to
//     RUNTIME_MONITORING , all the optional features are enabled by default.
//
//   - When you specify some of the features , any feature that is not specified in
//     the API call gets enabled by default, with an exception to RUNTIME_MONITORING
//     .
//
// Specifying both EKS Runtime Monitoring ( EKS_RUNTIME_MONITORING ) and Runtime
// Monitoring ( RUNTIME_MONITORING ) will cause an error. You can add only one of
// these two features because Runtime Monitoring already includes the threat
// detection for Amazon EKS resources. For more information, see [Runtime Monitoring].
//
// There might be regional differences because some data sources might not be
// available in all the Amazon Web Services Regions where GuardDuty is presently
// supported. For more information, see [Regions and endpoints].
//
// [Regions and endpoints]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html
// [Runtime Monitoring]: https://docs.aws.amazon.com/guardduty/latest/ug/runtime-monitoring.html
func (c *Client) CreateDetector(ctx context.Context, params *CreateDetectorInput, optFns ...func(*Options)) (*CreateDetectorOutput, error) {
	if params == nil {
		params = &CreateDetectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDetector", params, optFns, c.addOperationCreateDetectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDetectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDetectorInput struct {

	// A Boolean value that specifies whether the detector is to be enabled.
	//
	// This member is required.
	Enable *bool

	// The idempotency token for the create request.
	ClientToken *string

	// Describes which data sources will be enabled for the detector.
	//
	// There might be regional differences because some data sources might not be
	// available in all the Amazon Web Services Regions where GuardDuty is presently
	// supported. For more information, see [Regions and endpoints].
	//
	// [Regions and endpoints]: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_regions.html
	//
	// Deprecated: This parameter is deprecated, use Features instead
	DataSources *types.DataSourceConfigurations

	// A list of features that will be configured for the detector.
	Features []types.DetectorFeatureConfiguration

	// A value that specifies how frequently updated findings are exported.
	FindingPublishingFrequency types.FindingPublishingFrequency

	// The tags to be added to a new detector resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateDetectorOutput struct {

	// The unique ID of the created detector.
	DetectorId *string

	// Specifies the data sources that couldn't be enabled when GuardDuty was enabled
	// for the first time.
	UnprocessedDataSources *types.UnprocessedDataSourcesResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDetectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDetector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDetector{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDetector"); err != nil {
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
	if err = addIdempotencyToken_opCreateDetectorMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateDetectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDetector(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateDetector struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDetector) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDetector) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDetectorInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDetectorInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateDetectorMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateDetector{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDetector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDetector",
	}
}
