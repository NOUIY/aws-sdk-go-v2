// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The SetVisibleToAllUsers parameter is no longer supported. Your cluster may be
// visible to all users in your account. To restrict cluster access using an IAM
// policy, see [Identity and Access Management for Amazon EMR].
//
// Sets the Cluster$VisibleToAllUsers value for an Amazon EMR cluster. When true , IAM principals in the
// Amazon Web Services account can perform Amazon EMR cluster actions that their
// IAM policies allow. When false , only the IAM principal that created the cluster
// and the Amazon Web Services account root user can perform Amazon EMR actions on
// the cluster, regardless of IAM permissions policies attached to other IAM
// principals.
//
// This action works on running clusters. When you create a cluster, use the RunJobFlowInput$VisibleToAllUsers
// parameter.
//
// For more information, see [Understanding the Amazon EMR Cluster VisibleToAllUsers Setting] in the Amazon EMR Management Guide.
//
// [Understanding the Amazon EMR Cluster VisibleToAllUsers Setting]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/security_IAM_emr-with-IAM.html#security_set_visible_to_all_users
// [Identity and Access Management for Amazon EMR]: https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-access-IAM.html
func (c *Client) SetVisibleToAllUsers(ctx context.Context, params *SetVisibleToAllUsersInput, optFns ...func(*Options)) (*SetVisibleToAllUsersOutput, error) {
	if params == nil {
		params = &SetVisibleToAllUsersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetVisibleToAllUsers", params, optFns, c.addOperationSetVisibleToAllUsersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetVisibleToAllUsersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input to the SetVisibleToAllUsers action.
type SetVisibleToAllUsersInput struct {

	// The unique identifier of the job flow (cluster).
	//
	// This member is required.
	JobFlowIds []string

	// A value of true indicates that an IAM principal in the Amazon Web Services
	// account can perform Amazon EMR actions on the cluster that the IAM policies
	// attached to the principal allow. A value of false indicates that only the IAM
	// principal that created the cluster and the Amazon Web Services root user can
	// perform Amazon EMR actions on the cluster.
	//
	// This member is required.
	VisibleToAllUsers *bool

	noSmithyDocumentSerde
}

type SetVisibleToAllUsersOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetVisibleToAllUsersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetVisibleToAllUsers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetVisibleToAllUsers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetVisibleToAllUsers"); err != nil {
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
	if err = addOpSetVisibleToAllUsersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetVisibleToAllUsers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetVisibleToAllUsers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetVisibleToAllUsers",
	}
}
