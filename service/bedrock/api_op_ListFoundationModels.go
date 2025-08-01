// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrock

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrock/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists Amazon Bedrock foundation models that you can use. You can filter the
// results with the request parameters. For more information, see [Foundation models]in the [Amazon Bedrock User Guide].
//
// [Foundation models]: https://docs.aws.amazon.com/bedrock/latest/userguide/foundation-models.html
// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
func (c *Client) ListFoundationModels(ctx context.Context, params *ListFoundationModelsInput, optFns ...func(*Options)) (*ListFoundationModelsOutput, error) {
	if params == nil {
		params = &ListFoundationModelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFoundationModels", params, optFns, c.addOperationListFoundationModelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFoundationModelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFoundationModelsInput struct {

	// Return models that support the customization type that you specify. For more
	// information, see [Custom models]in the [Amazon Bedrock User Guide].
	//
	// [Custom models]: https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models.html
	// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
	ByCustomizationType types.ModelCustomization

	// Return models that support the inference type that you specify. For more
	// information, see [Provisioned Throughput]in the [Amazon Bedrock User Guide].
	//
	// [Provisioned Throughput]: https://docs.aws.amazon.com/bedrock/latest/userguide/prov-throughput.html
	// [Amazon Bedrock User Guide]: https://docs.aws.amazon.com/bedrock/latest/userguide/what-is-service.html
	ByInferenceType types.InferenceType

	// Return models that support the output modality that you specify.
	ByOutputModality types.ModelModality

	// Return models belonging to the model provider that you specify.
	ByProvider *string

	noSmithyDocumentSerde
}

type ListFoundationModelsOutput struct {

	// A list of Amazon Bedrock foundation models.
	ModelSummaries []types.FoundationModelSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFoundationModelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFoundationModels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFoundationModels{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFoundationModels"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFoundationModels(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListFoundationModels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFoundationModels",
	}
}
