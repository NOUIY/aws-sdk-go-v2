// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the merge options available for merging two specified
// branches. For details about why a merge option is not available, use
// GetMergeConflicts or DescribeMergeConflicts.
func (c *Client) GetMergeOptions(ctx context.Context, params *GetMergeOptionsInput, optFns ...func(*Options)) (*GetMergeOptionsOutput, error) {
	if params == nil {
		params = &GetMergeOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMergeOptions", params, optFns, c.addOperationGetMergeOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMergeOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMergeOptionsInput struct {

	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	//
	// This member is required.
	DestinationCommitSpecifier *string

	// The name of the repository that contains the commits about which you want to
	// get merge options.
	//
	// This member is required.
	RepositoryName *string

	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	//
	// This member is required.
	SourceCommitSpecifier *string

	// The level of conflict detail to use. If unspecified, the default FILE_LEVEL is
	// used, which returns a not-mergeable result if the same file has differences in
	// both branches. If LINE_LEVEL is specified, a conflict is considered not
	// mergeable if the same file in both branches has differences on the same line.
	ConflictDetailLevel types.ConflictDetailLevelTypeEnum

	// Specifies which branch to use when resolving conflicts, or whether to attempt
	// automatically merging two versions of a file. The default is NONE, which
	// requires any conflicts to be resolved manually before the merge operation is
	// successful.
	ConflictResolutionStrategy types.ConflictResolutionStrategyTypeEnum

	noSmithyDocumentSerde
}

type GetMergeOptionsOutput struct {

	// The commit ID of the merge base.
	//
	// This member is required.
	BaseCommitId *string

	// The commit ID of the destination commit specifier that was used in the merge
	// evaluation.
	//
	// This member is required.
	DestinationCommitId *string

	// The merge option or strategy used to merge the code.
	//
	// This member is required.
	MergeOptions []types.MergeOptionTypeEnum

	// The commit ID of the source commit specifier that was used in the merge
	// evaluation.
	//
	// This member is required.
	SourceCommitId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMergeOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMergeOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMergeOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMergeOptions"); err != nil {
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
	if err = addOpGetMergeOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMergeOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMergeOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMergeOptions",
	}
}
