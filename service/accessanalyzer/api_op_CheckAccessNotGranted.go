// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Checks whether the specified access isn't allowed by a policy.
func (c *Client) CheckAccessNotGranted(ctx context.Context, params *CheckAccessNotGrantedInput, optFns ...func(*Options)) (*CheckAccessNotGrantedOutput, error) {
	if params == nil {
		params = &CheckAccessNotGrantedInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CheckAccessNotGranted", params, optFns, c.addOperationCheckAccessNotGrantedMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CheckAccessNotGrantedOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CheckAccessNotGrantedInput struct {

	// An access object containing the permissions that shouldn't be granted by the
	// specified policy. If only actions are specified, IAM Access Analyzer checks for
	// access to peform at least one of the actions on any resource in the policy. If
	// only resources are specified, then IAM Access Analyzer checks for access to
	// perform any action on at least one of the resources. If both actions and
	// resources are specified, IAM Access Analyzer checks for access to perform at
	// least one of the specified actions on at least one of the specified resources.
	//
	// This member is required.
	Access []types.Access

	// The JSON policy document to use as the content for the policy.
	//
	// This member is required.
	PolicyDocument *string

	// The type of policy. Identity policies grant permissions to IAM principals.
	// Identity policies include managed and inline policies for IAM roles, users, and
	// groups.
	//
	// Resource policies grant permissions on Amazon Web Services resources. Resource
	// policies include trust policies for IAM roles and bucket policies for Amazon S3
	// buckets.
	//
	// This member is required.
	PolicyType types.AccessCheckPolicyType

	noSmithyDocumentSerde
}

type CheckAccessNotGrantedOutput struct {

	// The message indicating whether the specified access is allowed.
	Message *string

	// A description of the reasoning of the result.
	Reasons []types.ReasonSummary

	// The result of the check for whether the access is allowed. If the result is PASS
	// , the specified policy doesn't allow any of the specified permissions in the
	// access object. If the result is FAIL , the specified policy might allow some or
	// all of the permissions in the access object.
	Result types.CheckAccessNotGrantedResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCheckAccessNotGrantedMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCheckAccessNotGranted{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCheckAccessNotGranted{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CheckAccessNotGranted"); err != nil {
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
	if err = addOpCheckAccessNotGrantedValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCheckAccessNotGranted(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCheckAccessNotGranted(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CheckAccessNotGranted",
	}
}
