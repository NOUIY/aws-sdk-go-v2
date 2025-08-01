// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an EFS access point. An access point is an application-specific view
// into an EFS file system that applies an operating system user and group, and a
// file system path, to any file system request made through the access point. The
// operating system user and group override any identity information provided by
// the NFS client. The file system path is exposed as the access point's root
// directory. Applications using the access point can only access data in the
// application's own directory and any subdirectories. A file system can have a
// maximum of 10,000 access points unless you request an increase. To learn more,
// see [Mounting a file system using EFS access points].
//
// If multiple requests to create access points on the same file system are sent
// in quick succession, and the file system is near the limit of access points, you
// may experience a throttling response for these requests. This is to ensure that
// the file system does not exceed the stated access point limit.
//
// This operation requires permissions for the elasticfilesystem:CreateAccessPoint
// action.
//
// Access points can be tagged on creation. If tags are specified in the creation
// action, IAM performs additional authorization on the
// elasticfilesystem:TagResource action to verify if users have permissions to
// create tags. Therefore, you must grant explicit permissions to use the
// elasticfilesystem:TagResource action. For more information, see [Granting permissions to tag resources during creation].
//
// [Mounting a file system using EFS access points]: https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html
// [Granting permissions to tag resources during creation]: https://docs.aws.amazon.com/efs/latest/ug/using-tags-efs.html#supported-iam-actions-tagging.html
func (c *Client) CreateAccessPoint(ctx context.Context, params *CreateAccessPointInput, optFns ...func(*Options)) (*CreateAccessPointOutput, error) {
	if params == nil {
		params = &CreateAccessPointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAccessPoint", params, optFns, c.addOperationCreateAccessPointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAccessPointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccessPointInput struct {

	// A string of up to 64 ASCII characters that Amazon EFS uses to ensure idempotent
	// creation.
	//
	// This member is required.
	ClientToken *string

	// The ID of the EFS file system that the access point provides access to.
	//
	// This member is required.
	FileSystemId *string

	// The operating system user and group applied to all file system requests made
	// using the access point.
	PosixUser *types.PosixUser

	// Specifies the directory on the EFS file system that the access point exposes as
	// the root directory of your file system to NFS clients using the access point.
	// The clients using the access point can only access the root directory and below.
	// If the RootDirectory > Path specified does not exist, Amazon EFS creates it and
	// applies the CreationInfo settings when a client connects to an access point.
	// When specifying a RootDirectory , you must provide the Path , and the
	// CreationInfo .
	//
	// Amazon EFS creates a root directory only if you have provided the CreationInfo:
	// OwnUid, OwnGID, and permissions for the directory. If you do not provide this
	// information, Amazon EFS does not create the root directory. If the root
	// directory does not exist, attempts to mount using the access point will fail.
	RootDirectory *types.RootDirectory

	// Creates tags associated with the access point. Each tag is a key-value pair,
	// each key must be unique. For more information, see [Tagging Amazon Web Services resources]in the Amazon Web Services
	// General Reference Guide.
	//
	// [Tagging Amazon Web Services resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Provides a description of an EFS file system access point.
type CreateAccessPointOutput struct {

	// The unique Amazon Resource Name (ARN) associated with the access point.
	AccessPointArn *string

	// The ID of the access point, assigned by Amazon EFS.
	AccessPointId *string

	// The opaque string specified in the request to ensure idempotent creation.
	ClientToken *string

	// The ID of the EFS file system that the access point applies to.
	FileSystemId *string

	// Identifies the lifecycle phase of the access point.
	LifeCycleState types.LifeCycleState

	// The name of the access point. This is the value of the Name tag.
	Name *string

	// Identifies the Amazon Web Services account that owns the access point resource.
	OwnerId *string

	// The full POSIX identity, including the user ID, group ID, and secondary group
	// IDs on the access point that is used for all file operations by NFS clients
	// using the access point.
	PosixUser *types.PosixUser

	// The directory on the EFS file system that the access point exposes as the root
	// directory to NFS clients using the access point.
	RootDirectory *types.RootDirectory

	// The tags associated with the access point, presented as an array of Tag objects.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAccessPointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAccessPoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAccessPoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAccessPoint"); err != nil {
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
	if err = addIdempotencyToken_opCreateAccessPointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAccessPointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccessPoint(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateAccessPoint struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAccessPoint) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAccessPoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAccessPointInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAccessPointInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateAccessPointMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateAccessPoint{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAccessPoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAccessPoint",
	}
}
