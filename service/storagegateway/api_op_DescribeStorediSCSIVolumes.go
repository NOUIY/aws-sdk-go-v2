// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the description of the gateway volumes specified in the request. The
// list of gateway volumes in the request must be from one gateway. In the
// response, Storage Gateway returns volume information sorted by volume ARNs. This
// operation is only supported in stored volume gateway type.
func (c *Client) DescribeStorediSCSIVolumes(ctx context.Context, params *DescribeStorediSCSIVolumesInput, optFns ...func(*Options)) (*DescribeStorediSCSIVolumesOutput, error) {
	if params == nil {
		params = &DescribeStorediSCSIVolumesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStorediSCSIVolumes", params, optFns, c.addOperationDescribeStorediSCSIVolumesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStorediSCSIVolumesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing a list of DescribeStorediSCSIVolumesInput$VolumeARNs.
type DescribeStorediSCSIVolumesInput struct {

	// An array of strings where each string represents the Amazon Resource Name (ARN)
	// of a stored volume. All of the specified stored volumes must be from the same
	// gateway. Use ListVolumesto get volume ARNs for a gateway.
	//
	// This member is required.
	VolumeARNs []string

	noSmithyDocumentSerde
}

type DescribeStorediSCSIVolumesOutput struct {

	// Describes a single unit of output from DescribeStorediSCSIVolumes. The following fields are returned:
	//
	//   - ChapEnabled : Indicates whether mutual CHAP is enabled for the iSCSI target.
	//
	//   - LunNumber : The logical disk number.
	//
	//   - NetworkInterfaceId : The network interface ID of the stored volume that
	//   initiator use to map the stored volume as an iSCSI target.
	//
	//   - NetworkInterfacePort : The port used to communicate with iSCSI targets.
	//
	//   - PreservedExistingData : Indicates when the stored volume was created,
	//   existing data on the underlying local disk was preserved.
	//
	//   - SourceSnapshotId : If the stored volume was created from a snapshot, this
	//   field contains the snapshot ID used, e.g. snap-1122aabb . Otherwise, this
	//   field is not included.
	//
	//   - StorediSCSIVolumes : An array of StorediSCSIVolume objects where each object
	//   contains metadata about one stored volume.
	//
	//   - TargetARN : The Amazon Resource Name (ARN) of the volume target.
	//
	//   - VolumeARN : The Amazon Resource Name (ARN) of the stored volume.
	//
	//   - VolumeDiskId : The disk ID of the local disk that was specified in the CreateStorediSCSIVolume
	//   operation.
	//
	//   - VolumeId : The unique identifier of the storage volume, e.g. vol-1122AABB .
	//
	//   - VolumeiSCSIAttributes : An VolumeiSCSIAttributesobject that represents a collection of iSCSI
	//   attributes for one stored volume.
	//
	//   - VolumeProgress : Represents the percentage complete if the volume is
	//   restoring or bootstrapping that represents the percent of data transferred. This
	//   field does not appear in the response if the stored volume is not restoring or
	//   bootstrapping.
	//
	//   - VolumeSizeInBytes : The size of the volume in bytes.
	//
	//   - VolumeStatus : One of the VolumeStatus values that indicates the state of
	//   the volume.
	//
	//   - VolumeType : One of the enumeration values describing the type of the
	//   volume. Currently, only STORED volumes are supported.
	StorediSCSIVolumes []types.StorediSCSIVolume

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStorediSCSIVolumesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStorediSCSIVolumes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStorediSCSIVolumes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeStorediSCSIVolumes"); err != nil {
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
	if err = addOpDescribeStorediSCSIVolumesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStorediSCSIVolumes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeStorediSCSIVolumes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeStorediSCSIVolumes",
	}
}
