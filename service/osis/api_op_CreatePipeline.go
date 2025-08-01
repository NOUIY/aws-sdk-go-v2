// Code generated by smithy-go-codegen DO NOT EDIT.

package osis

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/osis/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an OpenSearch Ingestion pipeline. For more information, see [Creating Amazon OpenSearch Ingestion pipelines].
//
// [Creating Amazon OpenSearch Ingestion pipelines]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/creating-pipeline.html
func (c *Client) CreatePipeline(ctx context.Context, params *CreatePipelineInput, optFns ...func(*Options)) (*CreatePipelineOutput, error) {
	if params == nil {
		params = &CreatePipelineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePipeline", params, optFns, c.addOperationCreatePipelineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePipelineInput struct {

	// The maximum pipeline capacity, in Ingestion Compute Units (ICUs).
	//
	// This member is required.
	MaxUnits *int32

	// The minimum pipeline capacity, in Ingestion Compute Units (ICUs).
	//
	// This member is required.
	MinUnits *int32

	// The pipeline configuration in YAML format. The command accepts the pipeline
	// configuration as a string or within a .yaml file. If you provide the
	// configuration as a string, each new line must be escaped with \n .
	//
	// This member is required.
	PipelineConfigurationBody *string

	// The name of the OpenSearch Ingestion pipeline to create. Pipeline names are
	// unique across the pipelines owned by an account within an Amazon Web Services
	// Region.
	//
	// This member is required.
	PipelineName *string

	// Key-value pairs to configure persistent buffering for the pipeline.
	BufferOptions *types.BufferOptions

	// Key-value pairs to configure encryption for data that is written to a
	// persistent buffer.
	EncryptionAtRestOptions *types.EncryptionAtRestOptions

	// Key-value pairs to configure log publishing.
	LogPublishingOptions *types.LogPublishingOptions

	// The Amazon Resource Name (ARN) of an IAM role that provides the required
	// permissions for a pipeline to read from the source and write to the sink. For
	// more information, see [Setting up roles and users in Amazon OpenSearch Ingestion].
	//
	// [Setting up roles and users in Amazon OpenSearch Ingestion]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/pipeline-security-overview.html
	PipelineRoleArn *string

	// List of tags to add to the pipeline upon creation.
	Tags []types.Tag

	// Container for the values required to configure VPC access for the pipeline. If
	// you don't specify these values, OpenSearch Ingestion creates the pipeline with a
	// public endpoint.
	VpcOptions *types.VpcOptions

	noSmithyDocumentSerde
}

type CreatePipelineOutput struct {

	// Container for information about the created pipeline.
	Pipeline *types.Pipeline

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePipeline"); err != nil {
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
	if err = addOpCreatePipelineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePipeline(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePipeline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePipeline",
	}
}
