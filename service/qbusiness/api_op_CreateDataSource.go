// Code generated by smithy-go-codegen DO NOT EDIT.

package qbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/document"
	"github.com/aws/aws-sdk-go-v2/service/qbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a data source connector for an Amazon Q Business application.
//
// CreateDataSource is a synchronous operation. The operation returns 200 if the
// data source was successfully created. Otherwise, an exception is raised.
func (c *Client) CreateDataSource(ctx context.Context, params *CreateDataSourceInput, optFns ...func(*Options)) (*CreateDataSourceOutput, error) {
	if params == nil {
		params = &CreateDataSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataSource", params, optFns, c.addOperationCreateDataSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataSourceInput struct {

	//  The identifier of the Amazon Q Business application the data source will be
	// attached to.
	//
	// This member is required.
	ApplicationId *string

	// Configuration information to connect your data source repository to Amazon Q
	// Business. Use this parameter to provide a JSON schema with configuration
	// information specific to your data source connector.
	//
	// Each data source has a JSON schema provided by Amazon Q Business that you must
	// use. For example, the Amazon S3 and Web Crawler connectors require the following
	// JSON schemas:
	//
	// [Amazon S3 JSON schema]
	//
	// [Web Crawler JSON schema]
	//
	// You can find configuration templates for your specific data source using the
	// following steps:
	//
	//   - Navigate to the [Supported connectors]page in the Amazon Q Business User Guide, and select the
	//   data source of your choice.
	//
	//   - Then, from your specific data source connector page, select Using the API.
	//   You will find the JSON schema for your data source, including parameter
	//   descriptions, in this section.
	//
	// [Supported connectors]: https://docs.aws.amazon.com/amazonq/latest/business-use-dg/connectors-list.html
	// [Web Crawler JSON schema]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/web-crawler-api.html
	// [Amazon S3 JSON schema]: https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/s3-api.html
	//
	// This member is required.
	Configuration document.Interface

	// A name for the data source connector.
	//
	// This member is required.
	DisplayName *string

	// The identifier of the index that you want to use with the data source connector.
	//
	// This member is required.
	IndexId *string

	// A token you provide to identify a request to create a data source connector.
	// Multiple calls to the CreateDataSource API with the same client token will
	// create only one data source connector.
	ClientToken *string

	// A description for the data source connector.
	Description *string

	// Provides the configuration information for altering document metadata and
	// content during the document ingestion process.
	//
	// For more information, see [Custom document enrichment].
	//
	// [Custom document enrichment]: https://docs.aws.amazon.com/amazonq/latest/business-use-dg/custom-document-enrichment.html
	DocumentEnrichmentConfiguration *types.DocumentEnrichmentConfiguration

	// The configuration for extracting information from media in documents during
	// ingestion.
	MediaExtractionConfiguration *types.MediaExtractionConfiguration

	// The Amazon Resource Name (ARN) of an IAM role with permission to access the
	// data source and required resources.
	RoleArn *string

	// Sets the frequency for Amazon Q Business to check the documents in your data
	// source repository and update your index. If you don't set a schedule, Amazon Q
	// Business won't periodically update the index.
	//
	// Specify a cron- format schedule string or an empty string to indicate that the
	// index is updated on demand. You can't specify the Schedule parameter when the
	// Type parameter is set to CUSTOM . If you do, you receive a ValidationException
	// exception.
	SyncSchedule *string

	// A list of key-value pairs that identify or categorize the data source
	// connector. You can also use tags to help control access to the data source
	// connector. Tag keys and values can consist of Unicode letters, digits, white
	// space, and any of the following symbols: _ . : / = + - @.
	Tags []types.Tag

	// Configuration information for an Amazon VPC (Virtual Private Cloud) to connect
	// to your data source. For more information, see [Using Amazon VPC with Amazon Q Business connectors].
	//
	// [Using Amazon VPC with Amazon Q Business connectors]: https://docs.aws.amazon.com/amazonq/latest/business-use-dg/connector-vpc.html
	VpcConfiguration *types.DataSourceVpcConfiguration

	noSmithyDocumentSerde
}

type CreateDataSourceOutput struct {

	//  The Amazon Resource Name (ARN) of a data source in an Amazon Q Business
	// application.
	DataSourceArn *string

	// The identifier of the data source connector.
	DataSourceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDataSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDataSource"); err != nil {
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
	if err = addIdempotencyToken_opCreateDataSourceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateDataSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataSource(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateDataSource struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDataSource) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDataSource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDataSourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDataSourceInput ")
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
func addIdempotencyToken_opCreateDataSourceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateDataSource{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDataSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDataSource",
	}
}
