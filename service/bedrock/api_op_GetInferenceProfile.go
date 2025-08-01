// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrock

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about an inference profile. For more information, see [Increase throughput and resilience with cross-region inference in Amazon Bedrock]. in the
// Amazon Bedrock User Guide.
//
// [Increase throughput and resilience with cross-region inference in Amazon Bedrock]: https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html
func (c *Client) GetInferenceProfile(ctx context.Context, params *GetInferenceProfileInput, optFns ...func(*Options)) (*GetInferenceProfileOutput, error) {
	if params == nil {
		params = &GetInferenceProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInferenceProfile", params, optFns, c.addOperationGetInferenceProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInferenceProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInferenceProfileInput struct {

	// The ID or Amazon Resource Name (ARN) of the inference profile.
	//
	// This member is required.
	InferenceProfileIdentifier *string

	noSmithyDocumentSerde
}

type GetInferenceProfileOutput struct {

	// The Amazon Resource Name (ARN) of the inference profile.
	//
	// This member is required.
	InferenceProfileArn *string

	// The unique identifier of the inference profile.
	//
	// This member is required.
	InferenceProfileId *string

	// The name of the inference profile.
	//
	// This member is required.
	InferenceProfileName *string

	// A list of information about each model in the inference profile.
	//
	// This member is required.
	Models []types.InferenceProfileModel

	// The status of the inference profile. ACTIVE means that the inference profile is
	// ready to be used.
	//
	// This member is required.
	Status types.InferenceProfileStatus

	// The type of the inference profile. The following types are possible:
	//
	//   - SYSTEM_DEFINED – The inference profile is defined by Amazon Bedrock. You can
	//   route inference requests across regions with these inference profiles.
	//
	//   - APPLICATION – The inference profile was created by a user. This type of
	//   inference profile can track metrics and costs when invoking the model in it. The
	//   inference profile may route requests to one or multiple regions.
	//
	// This member is required.
	Type types.InferenceProfileType

	// The time at which the inference profile was created.
	CreatedAt *time.Time

	// The description of the inference profile.
	Description *string

	// The time at which the inference profile was last updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetInferenceProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetInferenceProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetInferenceProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetInferenceProfile"); err != nil {
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
	if err = addOpGetInferenceProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInferenceProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetInferenceProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetInferenceProfile",
	}
}
