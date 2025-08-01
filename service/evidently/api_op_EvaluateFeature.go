// Code generated by smithy-go-codegen DO NOT EDIT.

package evidently

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/evidently/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation assigns a feature variation to one given user session. You pass
// in an entityID that represents the user. Evidently then checks the evaluation
// rules and assigns the variation.
//
// The first rules that are evaluated are the override rules. If the user's
// entityID matches an override rule, the user is served the variation specified by
// that rule.
//
// If there is a current launch with this feature that uses segment overrides, and
// if the user session's evaluationContext matches a segment rule defined in a
// segment override, the configuration in the segment overrides is used. For more
// information about segments, see [CreateSegment]and [Use segments to focus your audience].
//
// If there is a launch with no segment overrides, the user might be assigned to a
// variation in the launch. The chance of this depends on the percentage of users
// that are allocated to that launch. If the user is enrolled in the launch, the
// variation they are served depends on the allocation of the various feature
// variations used for the launch.
//
// If the user is not assigned to a launch, and there is an ongoing experiment for
// this feature, the user might be assigned to a variation in the experiment. The
// chance of this depends on the percentage of users that are allocated to that
// experiment.
//
// If the experiment uses a segment, then only user sessions with evaluationContext
// values that match the segment rule are used in the experiment.
//
// If the user is enrolled in the experiment, the variation they are served
// depends on the allocation of the various feature variations used for the
// experiment.
//
// If the user is not assigned to a launch or experiment, they are served the
// default variation.
//
// [Use segments to focus your audience]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments.html
// [CreateSegment]: https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_CreateSegment.html
func (c *Client) EvaluateFeature(ctx context.Context, params *EvaluateFeatureInput, optFns ...func(*Options)) (*EvaluateFeatureOutput, error) {
	if params == nil {
		params = &EvaluateFeatureInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EvaluateFeature", params, optFns, c.addOperationEvaluateFeatureMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EvaluateFeatureOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EvaluateFeatureInput struct {

	// An internal ID that represents a unique user of the application. This entityID
	// is checked against any override rules assigned for this feature.
	//
	// This member is required.
	EntityId *string

	// The name of the feature being evaluated.
	//
	// This member is required.
	Feature *string

	// The name or ARN of the project that contains this feature.
	//
	// This member is required.
	Project *string

	// A JSON object of attributes that you can optionally pass in as part of the
	// evaluation event sent to Evidently from the user session. Evidently can use this
	// value to match user sessions with defined audience segments. For more
	// information, see [Use segments to focus your audience].
	//
	// If you include this parameter, the value must be a JSON object. A JSON array is
	// not supported.
	//
	// [Use segments to focus your audience]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-segments.html
	//
	// This value conforms to the media type: application/json
	EvaluationContext *string

	noSmithyDocumentSerde
}

type EvaluateFeatureOutput struct {

	// If this user was assigned to a launch or experiment, this field lists the
	// launch or experiment name.
	//
	// This value conforms to the media type: application/json
	Details *string

	// Specifies the reason that the user session was assigned this variation.
	// Possible values include DEFAULT , meaning the user was served the default
	// variation; LAUNCH_RULE_MATCH , if the user session was enrolled in a launch;
	// EXPERIMENT_RULE_MATCH , if the user session was enrolled in an experiment; or
	// ENTITY_OVERRIDES_MATCH , if the user's entityId matches an override rule.
	Reason *string

	// The value assigned to this variation to differentiate it from the other
	// variations of this feature.
	Value types.VariableValue

	// The name of the variation that was served to the user session.
	Variation *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEvaluateFeatureMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpEvaluateFeature{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpEvaluateFeature{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "EvaluateFeature"); err != nil {
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
	if err = addEndpointPrefix_opEvaluateFeatureMiddleware(stack); err != nil {
		return err
	}
	if err = addOpEvaluateFeatureValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEvaluateFeature(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opEvaluateFeatureMiddleware struct {
}

func (*endpointPrefix_opEvaluateFeatureMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opEvaluateFeatureMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "dataplane." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opEvaluateFeatureMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opEvaluateFeatureMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opEvaluateFeature(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "EvaluateFeature",
	}
}
