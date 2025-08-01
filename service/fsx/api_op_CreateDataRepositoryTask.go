// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon FSx for Lustre data repository task. A
// CreateDataRepositoryTask operation will fail if a data repository is not linked
// to the FSx file system.
//
// You use import and export data repository tasks to perform bulk operations
// between your FSx for Lustre file system and its linked data repositories. An
// example of a data repository task is exporting any data and metadata changes,
// including POSIX metadata, to files, directories, and symbolic links (symlinks)
// from your FSx file system to a linked data repository.
//
// You use release data repository tasks to release data from your file system for
// files that are exported to S3. The metadata of released files remains on the
// file system so users or applications can still access released files by reading
// the files again, which will restore data from Amazon S3 to the FSx for Lustre
// file system.
//
// To learn more about data repository tasks, see [Data Repository Tasks]. To learn more about linking a
// data repository to your file system, see [Linking your file system to an S3 bucket].
//
// [Data Repository Tasks]: https://docs.aws.amazon.com/fsx/latest/LustreGuide/data-repository-tasks.html
// [Linking your file system to an S3 bucket]: https://docs.aws.amazon.com/fsx/latest/LustreGuide/create-dra-linked-data-repo.html
func (c *Client) CreateDataRepositoryTask(ctx context.Context, params *CreateDataRepositoryTaskInput, optFns ...func(*Options)) (*CreateDataRepositoryTaskOutput, error) {
	if params == nil {
		params = &CreateDataRepositoryTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataRepositoryTask", params, optFns, c.addOperationCreateDataRepositoryTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataRepositoryTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataRepositoryTaskInput struct {

	// The globally unique ID of the file system, assigned by Amazon FSx.
	//
	// This member is required.
	FileSystemId *string

	// Defines whether or not Amazon FSx provides a CompletionReport once the task has
	// completed. A CompletionReport provides a detailed report on the files that
	// Amazon FSx processed that meet the criteria specified by the Scope parameter.
	// For more information, see [Working with Task Completion Reports].
	//
	// [Working with Task Completion Reports]: https://docs.aws.amazon.com/fsx/latest/LustreGuide/task-completion-report.html
	//
	// This member is required.
	Report *types.CompletionReport

	// Specifies the type of data repository task to create.
	//
	//   - EXPORT_TO_REPOSITORY tasks export from your Amazon FSx for Lustre file
	//   system to a linked data repository.
	//
	//   - IMPORT_METADATA_FROM_REPOSITORY tasks import metadata changes from a linked
	//   S3 bucket to your Amazon FSx for Lustre file system.
	//
	//   - RELEASE_DATA_FROM_FILESYSTEM tasks release files in your Amazon FSx for
	//   Lustre file system that have been exported to a linked S3 bucket and that meet
	//   your specified release criteria.
	//
	//   - AUTO_RELEASE_DATA tasks automatically release files from an Amazon File
	//   Cache resource.
	//
	// This member is required.
	Type types.DataRepositoryTaskType

	// Specifies the amount of data to release, in GiB, by an Amazon File Cache
	// AUTO_RELEASE_DATA task that automatically releases files from the cache.
	CapacityToRelease *int64

	// (Optional) An idempotency token for resource creation, in a string of up to 63
	// ASCII characters. This token is automatically filled on your behalf when you use
	// the Command Line Interface (CLI) or an Amazon Web Services SDK.
	ClientRequestToken *string

	// A list of paths for the data repository task to use when the task is processed.
	// If a path that you provide isn't valid, the task fails. If you don't provide
	// paths, the default behavior is to export all files to S3 (for export tasks),
	// import all files from S3 (for import tasks), or release all exported files that
	// meet the last accessed time criteria (for release tasks).
	//
	//   - For export tasks, the list contains paths on the FSx for Lustre file system
	//   from which the files are exported to the Amazon S3 bucket. The default path is
	//   the file system root directory. The paths you provide need to be relative to the
	//   mount point of the file system. If the mount point is /mnt/fsx and
	//   /mnt/fsx/path1 is a directory or file on the file system you want to export,
	//   then the path to provide is path1 .
	//
	//   - For import tasks, the list contains paths in the Amazon S3 bucket from
	//   which POSIX metadata changes are imported to the FSx for Lustre file system. The
	//   path can be an S3 bucket or prefix in the format s3://bucket-name/prefix
	//   (where prefix is optional).
	//
	//   - For release tasks, the list contains directory or file paths on the FSx for
	//   Lustre file system from which to release exported files. If a directory is
	//   specified, files within the directory are released. If a file path is specified,
	//   only that file is released. To release all exported files in the file system,
	//   specify a forward slash (/) as the path.
	//
	// A file must also meet the last accessed time criteria specified in for the file
	//   to be released.
	Paths []string

	// The configuration that specifies the last accessed time criteria for files that
	// will be released from an Amazon FSx for Lustre file system.
	ReleaseConfiguration *types.ReleaseConfiguration

	// A list of Tag values, with a maximum of 50 elements.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDataRepositoryTaskOutput struct {

	// The description of the data repository task that you just created.
	DataRepositoryTask *types.DataRepositoryTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDataRepositoryTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDataRepositoryTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDataRepositoryTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDataRepositoryTask"); err != nil {
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
	if err = addIdempotencyToken_opCreateDataRepositoryTaskMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateDataRepositoryTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataRepositoryTask(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateDataRepositoryTask struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDataRepositoryTask) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDataRepositoryTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDataRepositoryTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDataRepositoryTaskInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateDataRepositoryTaskMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateDataRepositoryTask{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDataRepositoryTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDataRepositoryTask",
	}
}
