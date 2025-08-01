// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a stack using the input information that was provided when the
// specified change set was created. After the call successfully completes,
// CloudFormation starts updating the stack. Use the DescribeStacksaction to view the status of
// the update.
//
// When you execute a change set, CloudFormation deletes all other change sets
// associated with the stack because they aren't valid for the updated stack.
//
// If a stack policy is associated with the stack, CloudFormation enforces the
// policy during the update. You can't specify a temporary stack policy that
// overrides the current policy.
//
// To create a change set for the entire stack hierarchy, IncludeNestedStacks must
// have been set to True .
func (c *Client) ExecuteChangeSet(ctx context.Context, params *ExecuteChangeSetInput, optFns ...func(*Options)) (*ExecuteChangeSetOutput, error) {
	if params == nil {
		params = &ExecuteChangeSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExecuteChangeSet", params, optFns, c.addOperationExecuteChangeSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExecuteChangeSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ExecuteChangeSet action.
type ExecuteChangeSetInput struct {

	// The name or Amazon Resource Name (ARN) of the change set that you want use to
	// update the specified stack.
	//
	// This member is required.
	ChangeSetName *string

	// A unique identifier for this ExecuteChangeSet request. Specify this token if
	// you plan to retry requests so that CloudFormation knows that you're not
	// attempting to execute a change set to update a stack with the same name. You
	// might retry ExecuteChangeSet requests to ensure that CloudFormation
	// successfully received them.
	ClientRequestToken *string

	// Preserves the state of previously provisioned resources when an operation
	// fails. This parameter can't be specified when the OnStackFailure parameter to
	// the [CreateChangeSet]API operation was specified.
	//
	//   - True - if the stack creation fails, do nothing. This is equivalent to
	//   specifying DO_NOTHING for the OnStackFailure parameter to the [CreateChangeSet]API operation.
	//
	//   - False - if the stack creation fails, roll back the stack. This is equivalent
	//   to specifying ROLLBACK for the OnStackFailure parameter to the [CreateChangeSet]API operation.
	//
	// Default: True
	//
	// [CreateChangeSet]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_CreateChangeSet.html
	DisableRollback *bool

	// When set to true , newly created resources are deleted when the operation rolls
	// back. This includes newly created resources marked with a deletion policy of
	// Retain .
	//
	// Default: false
	RetainExceptOnCreate *bool

	// If you specified the name of a change set, specify the stack name or Amazon
	// Resource Name (ARN) that's associated with the change set you want to execute.
	StackName *string

	noSmithyDocumentSerde
}

// The output for the ExecuteChangeSet action.
type ExecuteChangeSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExecuteChangeSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpExecuteChangeSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpExecuteChangeSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ExecuteChangeSet"); err != nil {
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
	if err = addOpExecuteChangeSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExecuteChangeSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExecuteChangeSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ExecuteChangeSet",
	}
}
