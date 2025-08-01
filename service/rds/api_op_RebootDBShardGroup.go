// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// You might need to reboot your DB shard group, usually for maintenance reasons.
// For example, if you make certain modifications, reboot the DB shard group for
// the changes to take effect.
//
// This operation applies only to Aurora Limitless Database DBb shard groups.
func (c *Client) RebootDBShardGroup(ctx context.Context, params *RebootDBShardGroupInput, optFns ...func(*Options)) (*RebootDBShardGroupOutput, error) {
	if params == nil {
		params = &RebootDBShardGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RebootDBShardGroup", params, optFns, c.addOperationRebootDBShardGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RebootDBShardGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RebootDBShardGroupInput struct {

	// The name of the DB shard group to reboot.
	//
	// This member is required.
	DBShardGroupIdentifier *string

	noSmithyDocumentSerde
}

// Contains the details for an Amazon RDS DB shard group.
type RebootDBShardGroupOutput struct {

	// Specifies whether to create standby DB shard groups for the DB shard group.
	// Valid values are the following:
	//
	//   - 0 - Creates a DB shard group without a standby DB shard group. This is the
	//   default value.
	//
	//   - 1 - Creates a DB shard group with a standby DB shard group in a different
	//   Availability Zone (AZ).
	//
	//   - 2 - Creates a DB shard group with two standby DB shard groups in two
	//   different AZs.
	ComputeRedundancy *int32

	// The name of the primary DB cluster for the DB shard group.
	DBClusterIdentifier *string

	// The Amazon Resource Name (ARN) for the DB shard group.
	DBShardGroupArn *string

	// The name of the DB shard group.
	DBShardGroupIdentifier *string

	// The Amazon Web Services Region-unique, immutable identifier for the DB shard
	// group.
	DBShardGroupResourceId *string

	// The connection endpoint for the DB shard group.
	Endpoint *string

	// The maximum capacity of the DB shard group in Aurora capacity units (ACUs).
	MaxACU *float64

	// The minimum capacity of the DB shard group in Aurora capacity units (ACUs).
	MinACU *float64

	// Indicates whether the DB shard group is publicly accessible.
	//
	// When the DB shard group is publicly accessible, its Domain Name System (DNS)
	// endpoint resolves to the private IP address from within the DB shard group's
	// virtual private cloud (VPC). It resolves to the public IP address from outside
	// of the DB shard group's VPC. Access to the DB shard group is ultimately
	// controlled by the security group it uses. That public access isn't permitted if
	// the security group assigned to the DB shard group doesn't permit it.
	//
	// When the DB shard group isn't publicly accessible, it is an internal DB shard
	// group with a DNS name that resolves to a private IP address.
	//
	// For more information, see CreateDBShardGroup.
	//
	// This setting is only for Aurora Limitless Database.
	PubliclyAccessible *bool

	// The status of the DB shard group.
	Status *string

	// A list of tags.
	//
	// For more information, see [Tagging Amazon RDS resources] in the Amazon RDS User Guide or [Tagging Amazon Aurora and Amazon RDS resources] in the Amazon
	// Aurora User Guide.
	//
	// [Tagging Amazon RDS resources]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html
	// [Tagging Amazon Aurora and Amazon RDS resources]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Tagging.html
	TagList []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRebootDBShardGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRebootDBShardGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRebootDBShardGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RebootDBShardGroup"); err != nil {
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
	if err = addOpRebootDBShardGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRebootDBShardGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRebootDBShardGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RebootDBShardGroup",
	}
}
