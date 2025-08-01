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

//	Updates the throughput or capacity of a volume. During the update process, the
//
// filesystem might be unavailable for a few minutes. You can retry any operations
// after the update is complete.
func (c *Client) UpdateKxVolume(ctx context.Context, params *UpdateKxVolumeInput, optFns ...func(*Options)) (*UpdateKxVolumeOutput, error) {
	if params == nil {
		params = &UpdateKxVolumeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateKxVolume", params, optFns, c.addOperationUpdateKxVolumeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateKxVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateKxVolumeInput struct {

	// A unique identifier for the kdb environment where you created the storage
	// volume.
	//
	// This member is required.
	EnvironmentId *string

	//  A unique identifier for the volume.
	//
	// This member is required.
	VolumeName *string

	// A token that ensures idempotency. This token expires in 10 minutes.
	ClientToken *string

	//  A description of the volume.
	Description *string

	//  Specifies the configuration for the Network attached storage (NAS_1) file
	// system volume.
	Nas1Configuration *types.KxNAS1Configuration

	noSmithyDocumentSerde
}

type UpdateKxVolumeOutput struct {

	//  Specifies the clusters that a volume is attached to.
	AttachedClusters []types.KxAttachedCluster

	// The identifier of the availability zones.
	AvailabilityZoneIds []string

	// The number of availability zones you want to assign per volume. Currently,
	// FinSpace only supports SINGLE for volumes. This places dataview in a single AZ.
	AzMode types.KxAzMode

	//  The timestamp at which the volume was created in FinSpace. The value is
	// determined as epoch time in milliseconds. For example, the value for Monday,
	// November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
	CreatedTimestamp *time.Time

	//  The description for the volume.
	Description *string

	// A unique identifier for the kdb environment where you want to update the
	// volume.
	EnvironmentId *string

	// The last time that the volume was updated in FinSpace. The value is determined
	// as epoch time in milliseconds. For example, the value for Monday, November 1,
	// 2021 12:00:00 PM UTC is specified as 1635768000000.
	LastModifiedTimestamp *time.Time

	//  Specifies the configuration for the Network attached storage (NAS_1) file
	// system volume.
	Nas1Configuration *types.KxNAS1Configuration

	// The status of the volume.
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

	// A unique identifier for the volume that you want to update.
	VolumeName *string

	//  The type of file system volume. Currently, FinSpace only supports NAS_1 volume
	// type.
	VolumeType types.KxVolumeType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateKxVolumeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateKxVolume{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateKxVolume{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateKxVolume"); err != nil {
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
	if err = addIdempotencyToken_opUpdateKxVolumeMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateKxVolumeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateKxVolume(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateKxVolume struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateKxVolume) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateKxVolume) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateKxVolumeInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateKxVolumeInput ")
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
func addIdempotencyToken_opUpdateKxVolumeMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateKxVolume{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateKxVolume(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateKxVolume",
	}
}
