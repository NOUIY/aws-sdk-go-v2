// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new, empty Amazon FSx file system. You can create the following
// supported Amazon FSx file systems using the CreateFileSystem API operation:
//
//   - Amazon FSx for Lustre
//
//   - Amazon FSx for NetApp ONTAP
//
//   - Amazon FSx for OpenZFS
//
//   - Amazon FSx for Windows File Server
//
// This operation requires a client request token in the request that Amazon FSx
// uses to ensure idempotent creation. This means that calling the operation
// multiple times with the same client request token has no effect. By using the
// idempotent operation, you can retry a CreateFileSystem operation without the
// risk of creating an extra file system. This approach can be useful when an
// initial call fails in a way that makes it unclear whether a file system was
// created. Examples are if a transport level timeout occurred, or your connection
// was reset. If you use the same client request token and the initial call created
// a file system, the client receives success as long as the parameters are the
// same.
//
// If a file system with the specified client request token exists and the
// parameters match, CreateFileSystem returns the description of the existing file
// system. If a file system with the specified client request token exists and the
// parameters don't match, this call returns IncompatibleParameterError . If a file
// system with the specified client request token doesn't exist, CreateFileSystem
// does the following:
//
//   - Creates a new, empty Amazon FSx file system with an assigned ID, and an
//     initial lifecycle state of CREATING .
//
//   - Returns the description of the file system in JSON format.
//
// The CreateFileSystem call returns while the file system's lifecycle state is
// still CREATING . You can check the file-system creation status by calling the [DescribeFileSystems]
// operation, which returns the file system state along with other information.
//
// [DescribeFileSystems]: https://docs.aws.amazon.com/fsx/latest/APIReference/API_DescribeFileSystems.html
func (c *Client) CreateFileSystem(ctx context.Context, params *CreateFileSystemInput, optFns ...func(*Options)) (*CreateFileSystemOutput, error) {
	if params == nil {
		params = &CreateFileSystemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFileSystem", params, optFns, c.addOperationCreateFileSystemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFileSystemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request object used to create a new Amazon FSx file system.
type CreateFileSystemInput struct {

	// The type of Amazon FSx file system to create. Valid values are WINDOWS , LUSTRE
	// , ONTAP , and OPENZFS .
	//
	// This member is required.
	FileSystemType types.FileSystemType

	// Specifies the IDs of the subnets that the file system will be accessible from.
	// For Windows and ONTAP MULTI_AZ_1 deployment types,provide exactly two subnet
	// IDs, one for the preferred file server and one for the standby file server. You
	// specify one of these subnets as the preferred subnet using the
	// WindowsConfiguration > PreferredSubnetID or OntapConfiguration >
	// PreferredSubnetID properties. For more information about Multi-AZ file system
	// configuration, see [Availability and durability: Single-AZ and Multi-AZ file systems]in the Amazon FSx for Windows User Guide and [Availability and durability] in the Amazon
	// FSx for ONTAP User Guide.
	//
	// For Windows SINGLE_AZ_1 and SINGLE_AZ_2 and all Lustre deployment types,
	// provide exactly one subnet ID. The file server is launched in that subnet's
	// Availability Zone.
	//
	// [Availability and durability: Single-AZ and Multi-AZ file systems]: https://docs.aws.amazon.com/fsx/latest/WindowsGuide/high-availability-multiAZ.html
	// [Availability and durability]: https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/high-availability-multiAZ.html
	//
	// This member is required.
	SubnetIds []string

	// A string of up to 63 ASCII characters that Amazon FSx uses to ensure idempotent
	// creation. This string is automatically filled on your behalf when you use the
	// Command Line Interface (CLI) or an Amazon Web Services SDK.
	ClientRequestToken *string

	// For FSx for Lustre file systems, sets the Lustre version for the file system
	// that you're creating. Valid values are 2.10 , 2.12 , and 2.15 :
	//
	//   - 2.10 is supported by the Scratch and Persistent_1 Lustre deployment types.
	//
	//   - 2.12 is supported by all Lustre deployment types, except for PERSISTENT_2
	//   with a metadata configuration mode.
	//
	//   - 2.15 is supported by all Lustre deployment types and is recommended for all
	//   new file systems.
	//
	// Default value is 2.10 , except for the following deployments:
	//
	//   - Default value is 2.12 when DeploymentType is set to PERSISTENT_2 without a
	//   metadata configuration mode.
	//
	//   - Default value is 2.15 when DeploymentType is set to PERSISTENT_2 with a
	//   metadata configuration mode.
	FileSystemTypeVersion *string

	// Specifies the ID of the Key Management Service (KMS) key to use for encrypting
	// data on Amazon FSx file systems, as follows:
	//
	//   - Amazon FSx for Lustre PERSISTENT_1 and PERSISTENT_2 deployment types only.
	//
	// SCRATCH_1 and SCRATCH_2 types are encrypted using the Amazon FSx service KMS key
	//   for your account.
	//
	//   - Amazon FSx for NetApp ONTAP
	//
	//   - Amazon FSx for OpenZFS
	//
	//   - Amazon FSx for Windows File Server
	//
	// If a KmsKeyId isn't specified, the Amazon FSx-managed KMS key for your account
	// is used. For more information, see [Encrypt]in the Key Management Service API Reference.
	//
	// [Encrypt]: https://docs.aws.amazon.com/kms/latest/APIReference/API_Encrypt.html
	KmsKeyId *string

	// The Lustre configuration for the file system being created.
	//
	// The following parameters are not supported for file systems with a data
	// repository association created with .
	//
	//   - AutoImportPolicy
	//
	//   - ExportPath
	//
	//   - ImportedFileChunkSize
	//
	//   - ImportPath
	LustreConfiguration *types.CreateFileSystemLustreConfiguration

	// The ONTAP configuration properties of the FSx for ONTAP file system that you
	// are creating.
	OntapConfiguration *types.CreateFileSystemOntapConfiguration

	// The OpenZFS configuration for the file system that's being created.
	OpenZFSConfiguration *types.CreateFileSystemOpenZFSConfiguration

	// A list of IDs specifying the security groups to apply to all network interfaces
	// created for file system access. This list isn't returned in later requests to
	// describe the file system.
	//
	// You must specify a security group if you are creating a Multi-AZ FSx for ONTAP
	// file system in a VPC subnet that has been shared with you.
	SecurityGroupIds []string

	// Sets the storage capacity of the file system that you're creating, in gibibytes
	// (GiB).
	//
	// FSx for Lustre file systems - The amount of storage capacity that you can
	// configure depends on the value that you set for StorageType and the Lustre
	// DeploymentType , as follows:
	//
	//   - For SCRATCH_2 , PERSISTENT_2 , and PERSISTENT_1 deployment types using SSD
	//   storage type, the valid values are 1200 GiB, 2400 GiB, and increments of 2400
	//   GiB.
	//
	//   - For PERSISTENT_1 HDD file systems, valid values are increments of 6000 GiB
	//   for 12 MB/s/TiB file systems and increments of 1800 GiB for 40 MB/s/TiB file
	//   systems.
	//
	//   - For SCRATCH_1 deployment type, valid values are 1200 GiB, 2400 GiB, and
	//   increments of 3600 GiB.
	//
	// FSx for ONTAP file systems - The amount of storage capacity that you can
	// configure depends on the value of the HAPairs property. The minimum value is
	// calculated as 1,024 * HAPairs and the maximum is calculated as 524,288 * HAPairs
	// .
	//
	// FSx for OpenZFS file systems - The amount of storage capacity that you can
	// configure is from 64 GiB up to 524,288 GiB (512 TiB).
	//
	// FSx for Windows File Server file systems - The amount of storage capacity that
	// you can configure depends on the value that you set for StorageType as follows:
	//
	//   - For SSD storage, valid values are 32 GiB-65,536 GiB (64 TiB).
	//
	//   - For HDD storage, valid values are 2000 GiB-65,536 GiB (64 TiB).
	StorageCapacity *int32

	// Sets the storage class for the file system that you're creating. Valid values
	// are SSD , HDD , and INTELLIGENT_TIERING .
	//
	//   - Set to SSD to use solid state drive storage. SSD is supported on all
	//   Windows, Lustre, ONTAP, and OpenZFS deployment types.
	//
	//   - Set to HDD to use hard disk drive storage, which is supported on SINGLE_AZ_2
	//   and MULTI_AZ_1 Windows file system deployment types, and on PERSISTENT_1
	//   Lustre file system deployment types.
	//
	//   - Set to INTELLIGENT_TIERING to use fully elastic, intelligently-tiered
	//   storage. Intelligent-Tiering is only available for OpenZFS file systems with the
	//   Multi-AZ deployment type and for Lustre file systems with the Persistent_2
	//   deployment type.
	//
	// Default value is SSD . For more information, see [Storage type options] in the FSx for Windows File
	// Server User Guide, [FSx for Lustre storage classes]in the FSx for Lustre User Guide, and [Working with Intelligent-Tiering] in the Amazon FSx for
	// OpenZFS User Guide.
	//
	// [FSx for Lustre storage classes]: https://docs.aws.amazon.com/fsx/latest/LustreGuide/using-fsx-lustre.html#lustre-storage-classes
	// [Storage type options]: https://docs.aws.amazon.com/fsx/latest/WindowsGuide/optimize-fsx-costs.html#storage-type-options
	// [Working with Intelligent-Tiering]: https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/performance-intelligent-tiering
	StorageType types.StorageType

	// The tags to apply to the file system that's being created. The key value of the
	// Name tag appears in the console as the file system name.
	Tags []types.Tag

	// The Microsoft Windows configuration for the file system that's being created.
	WindowsConfiguration *types.CreateFileSystemWindowsConfiguration

	noSmithyDocumentSerde
}

// The response object returned after the file system is created.
type CreateFileSystemOutput struct {

	// The configuration of the file system that was created.
	FileSystem *types.FileSystem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFileSystemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFileSystem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFileSystem{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateFileSystem"); err != nil {
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
	if err = addIdempotencyToken_opCreateFileSystemMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateFileSystemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFileSystem(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateFileSystem struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateFileSystem) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateFileSystem) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateFileSystemInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateFileSystemInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateFileSystemMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateFileSystem{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateFileSystem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateFileSystem",
	}
}
