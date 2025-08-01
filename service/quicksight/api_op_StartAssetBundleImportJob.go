// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an Asset Bundle import job.
//
// An Asset Bundle import job imports specified Amazon QuickSight assets into an
// Amazon QuickSight account. You can also choose to import a naming prefix and
// specified configuration overrides. The assets that are contained in the bundle
// file that you provide are used to create or update a new or existing asset in
// your Amazon QuickSight account. Each Amazon QuickSight account can run up to 5
// import jobs concurrently.
//
// The API caller must have the necessary "create" , "describe" , and "update"
// permissions in their IAM role to access each resource type that is contained in
// the bundle file before the resources can be imported.
func (c *Client) StartAssetBundleImportJob(ctx context.Context, params *StartAssetBundleImportJobInput, optFns ...func(*Options)) (*StartAssetBundleImportJobOutput, error) {
	if params == nil {
		params = &StartAssetBundleImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartAssetBundleImportJob", params, optFns, c.addOperationStartAssetBundleImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartAssetBundleImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartAssetBundleImportJobInput struct {

	// The ID of the job. This ID is unique while the job is running. After the job is
	// completed, you can reuse this ID for another job.
	//
	// This member is required.
	AssetBundleImportJobId *string

	// The source of the asset bundle zip file that contains the data that you want to
	// import. The file must be in QUICKSIGHT_JSON format.
	//
	// This member is required.
	AssetBundleImportSource *types.AssetBundleImportSource

	// The ID of the Amazon Web Services account to import assets into.
	//
	// This member is required.
	AwsAccountId *string

	// The failure action for the import job.
	//
	// If you choose ROLLBACK , failed import jobs will attempt to undo any asset
	// changes caused by the failed job.
	//
	// If you choose DO_NOTHING , failed import jobs will not attempt to roll back any
	// asset changes caused by the failed job, possibly keeping the Amazon QuickSight
	// account in an inconsistent state.
	FailureAction types.AssetBundleImportFailureAction

	// Optional overrides that are applied to the resource configuration before import.
	OverrideParameters *types.AssetBundleImportJobOverrideParameters

	// Optional permission overrides that are applied to the resource configuration
	// before import.
	OverridePermissions *types.AssetBundleImportJobOverridePermissions

	// Optional tag overrides that are applied to the resource configuration before
	// import.
	OverrideTags *types.AssetBundleImportJobOverrideTags

	// An optional validation strategy override for all analyses and dashboards that
	// is applied to the resource configuration before import.
	OverrideValidationStrategy *types.AssetBundleImportJobOverrideValidationStrategy

	noSmithyDocumentSerde
}

type StartAssetBundleImportJobOutput struct {

	// The Amazon Resource Name (ARN) for the import job.
	Arn *string

	// The ID of the job. This ID is unique while the job is running. After the job is
	// completed, you can reuse this ID for another job.
	AssetBundleImportJobId *string

	// The Amazon Web Services response ID for this operation.
	RequestId *string

	// The HTTP status of the response.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartAssetBundleImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartAssetBundleImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartAssetBundleImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartAssetBundleImportJob"); err != nil {
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
	if err = addOpStartAssetBundleImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartAssetBundleImportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartAssetBundleImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartAssetBundleImportJob",
	}
}
