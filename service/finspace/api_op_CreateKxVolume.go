// Code generated by smithy-go-codegen DO NOT EDIT.

package finspace

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/finspace/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

//	Creates a new volume with a specific amount of throughput and storage
//
// capacity.
func (c *Client) CreateKxVolume(ctx context.Context, params *CreateKxVolumeInput, optFns ...func(*Options)) (*CreateKxVolumeOutput, error) {
	if params == nil {
		params = &CreateKxVolumeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKxVolume", params, optFns, c.addOperationCreateKxVolumeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKxVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateKxVolumeInput struct {

	// The identifier of the availability zones.
	//
	// This member is required.
	AvailabilityZoneIds []string

	// The number of availability zones you want to assign per volume. Currently,
	// FinSpace only supports SINGLE for volumes. This places dataview in a single AZ.
	//
	// This member is required.
	AzMode types.KxAzMode

	// A unique identifier for the kdb environment, whose clusters can attach to the
	// volume.
	//
	// This member is required.
	EnvironmentId *string

	// A unique identifier for the volume.
	//
	// This member is required.
	VolumeName *string

	//  The type of file system volume. Currently, FinSpace only supports NAS_1 volume
	// type. When you select NAS_1 volume type, you must also provide nas1Configuration
	// .
	//
	// This member is required.
	VolumeType types.KxVolumeType

	// A token that ensures idempotency. This token expires in 10 minutes.
	ClientToken *string

	//  A description of the volume.
	Description *string

	//  Specifies the configuration for the Network attached storage (NAS_1) file
	// system volume. This parameter is required when you choose volumeType as NAS_1.
	Nas1Configuration *types.KxNAS1Configuration

	//  A list of key-value pairs to label the volume. You can add up to 50 tags to a
	// volume.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateKxVolumeOutput struct {

	// The identifier of the availability zones.
	AvailabilityZoneIds []string

	// The number of availability zones you want to assign per volume. Currently,
	// FinSpace only supports SINGLE for volumes. This places dataview in a single AZ.
	AzMode types.KxAzMode

	// The timestamp at which the volume was created in FinSpace. The value is
	// determined as epoch time in milliseconds. For example, the value for Monday,
	// November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
	CreatedTimestamp *time.Time

	//  A description of the volume.
	Description *string

	// A unique identifier for the kdb environment, whose clusters can attach to the
	// volume.
	EnvironmentId *string

	//  Specifies the configuration for the Network attached storage (NAS_1) file
	// system volume.
	Nas1Configuration *types.KxNAS1Configuration

	// The status of volume creation.
	//
	//   - CREATING – The volume creation is in progress.
	//
	//   - CREATE_FAILED – The volume creation has failed.
	//
	//   - ACTIVE – The volume is active.
	//
	//   - UPDATING – The volume is in the process of being updated.
	//
	//   - UPDATE_FAILED – The update action failed.
	//
	//   - UPDATED – The volume is successfully updated.
	//
	//   - DELETING – The volume is in the process of being deleted.
	//
	//   - DELETE_FAILED – The system failed to delete the volume.
	//
	//   - DELETED – The volume is successfully deleted.
	Status types.KxVolumeStatus

	// The error message when a failed state occurs.
	StatusReason *string

	//  The ARN identifier of the volume.
	VolumeArn *string

	// A unique identifier for the volume.
	VolumeName *string

	//  The type of file system volume. Currently, FinSpace only supports NAS_1 volume
	// type.
	VolumeType types.KxVolumeType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateKxVolumeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateKxVolume{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateKxVolume{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateKxVolume"); err != nil {
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
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
	if err = addIdempotencyToken_opCreateKxVolumeMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateKxVolumeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKxVolume(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateKxVolume struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateKxVolume) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateKxVolume) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateKxVolumeInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateKxVolumeInput ")
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
func addIdempotencyToken_opCreateKxVolumeMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateKxVolume{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateKxVolume(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateKxVolume",
	}
}
