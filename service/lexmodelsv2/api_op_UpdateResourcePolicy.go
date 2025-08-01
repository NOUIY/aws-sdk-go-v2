// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Replaces the existing resource policy for a bot or bot alias with a new one. If
// the policy doesn't exist, Amazon Lex returns an exception.
func (c *Client) UpdateResourcePolicy(ctx context.Context, params *UpdateResourcePolicyInput, optFns ...func(*Options)) (*UpdateResourcePolicyOutput, error) {
	if params == nil {
		params = &UpdateResourcePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateResourcePolicy", params, optFns, c.addOperationUpdateResourcePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateResourcePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateResourcePolicyInput struct {

	// A resource policy to add to the resource. The policy is a JSON structure that
	// contains one or more statements that define the policy. The policy must follow
	// the IAM syntax. For more information about the contents of a JSON policy
	// document, see [IAM JSON policy reference].
	//
	// If the policy isn't valid, Amazon Lex returns a validation exception.
	//
	// [IAM JSON policy reference]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html
	//
	// This member is required.
	Policy *string

	// The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy
	// is attached to.
	//
	// This member is required.
	ResourceArn *string

	// The identifier of the revision of the policy to update. If this revision ID
	// doesn't match the current revision ID, Amazon Lex throws an exception.
	//
	// If you don't specify a revision, Amazon Lex overwrites the contents of the
	// policy with the new values.
	ExpectedRevisionId *string

	noSmithyDocumentSerde
}

type UpdateResourcePolicyOutput struct {

	// The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy
	// is attached to.
	ResourceArn *string

	// The current revision of the resource policy. Use the revision ID to make sure
	// that you are updating the most current version of a resource policy when you add
	// a policy statement to a resource, delete a resource, or update a resource.
	RevisionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateResourcePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateResourcePolicy"); err != nil {
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
	if err = addOpUpdateResourcePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateResourcePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateResourcePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateResourcePolicy",
	}
}
