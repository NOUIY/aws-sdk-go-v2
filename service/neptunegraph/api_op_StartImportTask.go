// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunegraph

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/neptunegraph/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Import data into existing Neptune Analytics graph from Amazon Simple Storage
// Service (S3). The graph needs to be empty and in the AVAILABLE state.
func (c *Client) StartImportTask(ctx context.Context, params *StartImportTaskInput, optFns ...func(*Options)) (*StartImportTaskOutput, error) {
	if params == nil {
		params = &StartImportTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartImportTask", params, optFns, c.addOperationStartImportTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartImportTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartImportTaskInput struct {

	// The unique identifier of the Neptune Analytics graph.
	//
	// This member is required.
	GraphIdentifier *string

	// The ARN of the IAM role that will allow access to the data that is to be
	// imported.
	//
	// This member is required.
	RoleArn *string

	// A URL identifying the location of the data to be imported. This can be an
	// Amazon S3 path, or can point to a Neptune database endpoint or snapshot.
	//
	// This member is required.
	Source *string

	// The method to handle blank nodes in the dataset. Currently, only convertToIri
	// is supported, meaning blank nodes are converted to unique IRIs at load time.
	// Must be provided when format is ntriples . For more information, see [Handling RDF values].
	//
	// [Handling RDF values]: https://docs.aws.amazon.com/neptune-analytics/latest/userguide/using-rdf-data.html#rdf-handling
	BlankNodeHandling types.BlankNodeHandling

	// If set to true, the task halts when an import error is encountered. If set to
	// false, the task skips the data that caused the error and continues if possible.
	FailOnError *bool

	// Specifies the format of Amazon S3 data to be imported. Valid values are CSV,
	// which identifies the Gremlin CSV format or OPENCYPHER, which identies the
	// openCypher load format.
	Format types.Format

	// Options for how to perform an import.
	ImportOptions types.ImportOptions

	// The parquet type of the import task.
	ParquetType types.ParquetType

	noSmithyDocumentSerde
}

func (in *StartImportTaskInput) bindEndpointParams(p *EndpointParameters) {

	p.ApiType = ptr.String("ControlPlane")
}

type StartImportTaskOutput struct {

	// The ARN of the IAM role that will allow access to the data that is to be
	// imported.
	//
	// This member is required.
	RoleArn *string

	// A URL identifying the location of the data to be imported. This can be an
	// Amazon S3 path, or can point to a Neptune database endpoint or snapshot.
	//
	// This member is required.
	Source *string

	// The status of the import task.
	//
	// This member is required.
	Status types.ImportTaskStatus

	// The unique identifier of the import task.
	//
	// This member is required.
	TaskId *string

	// Specifies the format of Amazon S3 data to be imported. Valid values are CSV,
	// which identifies the Gremlin CSV format or OPENCYPHER, which identies the
	// openCypher load format.
	Format types.Format

	// The unique identifier of the Neptune Analytics graph.
	GraphId *string

	// Options for how to perform an import.
	ImportOptions types.ImportOptions

	// The parquet type of the import task.
	ParquetType types.ParquetType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartImportTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartImportTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartImportTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartImportTask"); err != nil {
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
	if err = addOpStartImportTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartImportTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartImportTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartImportTask",
	}
}
