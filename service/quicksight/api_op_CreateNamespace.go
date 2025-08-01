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

// (Enterprise edition only) Creates a new namespace for you to use with Amazon
// QuickSight.
//
// A namespace allows you to isolate the Amazon QuickSight users and groups that
// are registered for that namespace. Users that access the namespace can share
// assets only with other users or groups in the same namespace. They can't see
// users and groups in other namespaces. You can create a namespace after your
// Amazon Web Services account is subscribed to Amazon QuickSight. The namespace
// must be unique within the Amazon Web Services account. By default, there is a
// limit of 100 namespaces per Amazon Web Services account. To increase your limit,
// create a ticket with Amazon Web ServicesSupport.
func (c *Client) CreateNamespace(ctx context.Context, params *CreateNamespaceInput, optFns ...func(*Options)) (*CreateNamespaceOutput, error) {
	if params == nil {
		params = &CreateNamespaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNamespace", params, optFns, c.addOperationCreateNamespaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNamespaceInput struct {

	// The ID for the Amazon Web Services account that you want to create the Amazon
	// QuickSight namespace in.
	//
	// This member is required.
	AwsAccountId *string

	// Specifies the type of your user identity directory. Currently, this supports
	// users with an identity type of QUICKSIGHT .
	//
	// This member is required.
	IdentityStore types.IdentityStore

	// The name that you want to use to describe the new namespace.
	//
	// This member is required.
	Namespace *string

	// The tags that you want to associate with the namespace that you're creating.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateNamespaceOutput struct {

	// The ARN of the Amazon QuickSight namespace you created.
	Arn *string

	// The Amazon Web Services Region; that you want to use for the free SPICE
	// capacity for the new namespace. This is set to the region that you run
	// CreateNamespace in.
	CapacityRegion *string

	// The status of the creation of the namespace. This is an asynchronous process. A
	// status of CREATED means that your namespace is ready to use. If an error
	// occurs, it indicates if the process is retryable or non-retryable . In the case
	// of a non-retryable error, refer to the error message for follow-up tasks.
	CreationStatus types.NamespaceStatus

	// Specifies the type of your user identity directory. Currently, this supports
	// users with an identity type of QUICKSIGHT .
	IdentityStore types.IdentityStore

	// The name of the new namespace that you created.
	Name *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateNamespaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateNamespace"); err != nil {
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
	if err = addOpCreateNamespaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNamespace(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateNamespace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateNamespace",
	}
}
