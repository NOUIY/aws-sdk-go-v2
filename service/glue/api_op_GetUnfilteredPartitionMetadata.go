// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves partition metadata from the Data Catalog that contains unfiltered
// metadata.
//
// For IAM authorization, the public IAM action associated with this API is
// glue:GetPartition .
func (c *Client) GetUnfilteredPartitionMetadata(ctx context.Context, params *GetUnfilteredPartitionMetadataInput, optFns ...func(*Options)) (*GetUnfilteredPartitionMetadataOutput, error) {
	if params == nil {
		params = &GetUnfilteredPartitionMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUnfilteredPartitionMetadata", params, optFns, c.addOperationGetUnfilteredPartitionMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUnfilteredPartitionMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUnfilteredPartitionMetadataInput struct {

	// The catalog ID where the partition resides.
	//
	// This member is required.
	CatalogId *string

	// (Required) Specifies the name of a database that contains the partition.
	//
	// This member is required.
	DatabaseName *string

	// (Required) A list of partition key values.
	//
	// This member is required.
	PartitionValues []string

	// (Required) A list of supported permission types.
	//
	// This member is required.
	SupportedPermissionTypes []types.PermissionType

	// (Required) Specifies the name of a table that contains the partition.
	//
	// This member is required.
	TableName *string

	// A structure containing Lake Formation audit context information.
	AuditContext *types.AuditContext

	// A structure used as a protocol between query engines and Lake Formation or
	// Glue. Contains both a Lake Formation generated authorization identifier and
	// information from the request's authorization context.
	QuerySessionContext *types.QuerySessionContext

	// Specified only if the base tables belong to a different Amazon Web Services
	// Region.
	Region *string

	noSmithyDocumentSerde
}

type GetUnfilteredPartitionMetadataOutput struct {

	// A list of column names that the user has been granted access to.
	AuthorizedColumns []string

	// A Boolean value that indicates whether the partition location is registered
	// with Lake Formation.
	IsRegisteredWithLakeFormation bool

	// A Partition object containing the partition metadata.
	Partition *types.Partition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUnfilteredPartitionMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetUnfilteredPartitionMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetUnfilteredPartitionMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetUnfilteredPartitionMetadata"); err != nil {
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
	if err = addOpGetUnfilteredPartitionMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUnfilteredPartitionMetadata(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetUnfilteredPartitionMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetUnfilteredPartitionMetadata",
	}
}
