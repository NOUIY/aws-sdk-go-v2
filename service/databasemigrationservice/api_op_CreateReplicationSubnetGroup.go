// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a replication subnet group given a list of the subnet IDs in a VPC.
//
// The VPC needs to have at least one subnet in at least two availability zones in
// the Amazon Web Services Region, otherwise the service will throw a
// ReplicationSubnetGroupDoesNotCoverEnoughAZs exception.
//
// If a replication subnet group exists in your Amazon Web Services account, the
// CreateReplicationSubnetGroup action returns the following error message: The
// Replication Subnet Group already exists. In this case, delete the existing
// replication subnet group. To do so, use the [DeleteReplicationSubnetGroup]action. Optionally, choose Subnet
// groups in the DMS console, then choose your subnet group. Next, choose Delete
// from Actions.
//
// [DeleteReplicationSubnetGroup]: https://docs.aws.amazon.com/en_us/dms/latest/APIReference/API_DeleteReplicationSubnetGroup.html
func (c *Client) CreateReplicationSubnetGroup(ctx context.Context, params *CreateReplicationSubnetGroupInput, optFns ...func(*Options)) (*CreateReplicationSubnetGroupOutput, error) {
	if params == nil {
		params = &CreateReplicationSubnetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReplicationSubnetGroup", params, optFns, c.addOperationCreateReplicationSubnetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReplicationSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateReplicationSubnetGroupInput struct {

	// The description for the subnet group.
	//
	// Constraints: This parameter Must not contain non-printable control characters.
	//
	// This member is required.
	ReplicationSubnetGroupDescription *string

	// The name for the replication subnet group. This value is stored as a lowercase
	// string.
	//
	// Constraints: Must contain no more than 255 alphanumeric characters, periods,
	// underscores, or hyphens. Must not be "default".
	//
	// Example: mySubnetgroup
	//
	// This member is required.
	ReplicationSubnetGroupIdentifier *string

	// Two or more subnet IDs to be assigned to the subnet group.
	//
	// This member is required.
	SubnetIds []string

	// One or more tags to be assigned to the subnet group.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateReplicationSubnetGroupOutput struct {

	// The replication subnet group that was created.
	ReplicationSubnetGroup *types.ReplicationSubnetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateReplicationSubnetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateReplicationSubnetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateReplicationSubnetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateReplicationSubnetGroup"); err != nil {
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
	if err = addOpCreateReplicationSubnetGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReplicationSubnetGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateReplicationSubnetGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateReplicationSubnetGroup",
	}
}
