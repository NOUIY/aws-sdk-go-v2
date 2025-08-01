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

// Deletes the specified import task.
func (c *Client) CancelImportTask(ctx context.Context, params *CancelImportTaskInput, optFns ...func(*Options)) (*CancelImportTaskOutput, error) {
	if params == nil {
		params = &CancelImportTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelImportTask", params, optFns, c.addOperationCancelImportTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelImportTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelImportTaskInput struct {

	// The unique identifier of the import task.
	//
	// This member is required.
	TaskIdentifier *string

	noSmithyDocumentSerde
}

func (in *CancelImportTaskInput) bindEndpointParams(p *EndpointParameters) {

	p.ApiType = ptr.String("ControlPlane")
}

type CancelImportTaskOutput struct {

	// The ARN of the IAM role that will allow access to the data that is to be
	// imported.
	//
	// This member is required.
	RoleArn *string

	// A URL identifying to the location of the data to be imported. This can be an
	// Amazon S3 path, or can point to a Neptune database endpoint or snapshot.
	//
	// This member is required.
	Source *string

	// Current status of the task. Status is CANCELLING when the import task is
	// cancelled.
	//
	// This member is required.
	Status types.ImportTaskStatus

	// The unique identifier of the import task.
	//
	// This member is required.
	TaskId *string

	// Specifies the format of S3 data to be imported. Valid values are CSV , which
	// identifies the [Gremlin CSV format]or OPENCYPHER , which identies the [openCypher load format].
	//
	// [Gremlin CSV format]: https://docs.aws.amazon.com/neptune/latest/userguide/bulk-load-tutorial-format-gremlin.html
	// [openCypher load format]: https://docs.aws.amazon.com/neptune/latest/userguide/bulk-load-tutorial-format-opencypher.html
	Format types.Format

	// The unique identifier of the Neptune Analytics graph.
	GraphId *string

	// The parquet type of the cancelled import task.
	ParquetType types.ParquetType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelImportTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelImportTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelImportTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CancelImportTask"); err != nil {
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
	if err = addOpCancelImportTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelImportTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelImportTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CancelImportTask",
	}
}
