// Code generated by smithy-go-codegen DO NOT EDIT.

package keyspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/keyspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Adds a new Amazon Web Services Region to the keyspace. You can add a new
//
// Region to a keyspace that is either a single or a multi-Region keyspace. Amazon
// Keyspaces is going to replicate all tables in the keyspace to the new Region. To
// successfully replicate all tables to the new Region, they must use client-side
// timestamps for conflict resolution. To enable client-side timestamps, specify
// clientSideTimestamps.status = enabled when invoking the API. For more
// information about client-side timestamps, see [Client-side timestamps in Amazon Keyspaces]in the Amazon Keyspaces Developer
// Guide.
//
// To add a Region to a keyspace using the UpdateKeyspace API, the IAM principal
// needs permissions for the following IAM actions:
//
//   - cassandra:Alter
//
//   - cassandra:AlterMultiRegionResource
//
//   - cassandra:Create
//
//   - cassandra:CreateMultiRegionResource
//
//   - cassandra:Select
//
//   - cassandra:SelectMultiRegionResource
//
//   - cassandra:Modify
//
//   - cassandra:ModifyMultiRegionResource
//
// If the keyspace contains a table that is configured in provisioned mode with
// auto scaling enabled, the following additional IAM actions need to be allowed.
//
//   - application-autoscaling:RegisterScalableTarget
//
//   - application-autoscaling:DeregisterScalableTarget
//
//   - application-autoscaling:DescribeScalableTargets
//
//   - application-autoscaling:PutScalingPolicy
//
//   - application-autoscaling:DescribeScalingPolicies
//
// To use the UpdateKeyspace API, the IAM principal also needs permissions to
// create a service-linked role with the following elements:
//
//   - iam:CreateServiceLinkedRole - The action the principal can perform.
//
//     -
//     arn:aws:iam::*:role/aws-service-role/replication.cassandra.amazonaws.com/AWSServiceRoleForKeyspacesReplication
//
//   - The resource that the action can be performed on.
//
//   - iam:AWSServiceName: replication.cassandra.amazonaws.com - The only Amazon
//     Web Services service that this role can be attached to is Amazon Keyspaces.
//
// For more information, see [Configure the IAM permissions required to add an Amazon Web Services Region to a keyspace] in the Amazon Keyspaces Developer Guide.
//
// [Client-side timestamps in Amazon Keyspaces]: https://docs.aws.amazon.com/keyspaces/latest/devguide/client-side-timestamps.html
// [Configure the IAM permissions required to add an Amazon Web Services Region to a keyspace]: https://docs.aws.amazon.com/keyspaces/latest/devguide/howitworks_replication_permissions_addReplica.html
func (c *Client) UpdateKeyspace(ctx context.Context, params *UpdateKeyspaceInput, optFns ...func(*Options)) (*UpdateKeyspaceOutput, error) {
	if params == nil {
		params = &UpdateKeyspaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateKeyspace", params, optFns, c.addOperationUpdateKeyspaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateKeyspaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateKeyspaceInput struct {

	//  The name of the keyspace.
	//
	// This member is required.
	KeyspaceName *string

	//  The replication specification of the keyspace includes:
	//
	//   - regionList - the Amazon Web Services Regions where the keyspace is
	//   replicated in.
	//
	//   - replicationStrategy - the required value is SINGLE_REGION or MULTI_REGION .
	//
	// This member is required.
	ReplicationSpecification *types.ReplicationSpecification

	// The client-side timestamp setting of the table.
	//
	// For more information, see [How it works: Amazon Keyspaces client-side timestamps] in the Amazon Keyspaces Developer Guide.
	//
	// [How it works: Amazon Keyspaces client-side timestamps]: https://docs.aws.amazon.com/keyspaces/latest/devguide/client-side-timestamps-how-it-works.html
	ClientSideTimestamps *types.ClientSideTimestamps

	noSmithyDocumentSerde
}

type UpdateKeyspaceOutput struct {

	//  The unique identifier of the keyspace in the format of an Amazon Resource Name
	// (ARN).
	//
	// This member is required.
	ResourceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateKeyspaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateKeyspace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateKeyspace{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateKeyspace"); err != nil {
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
	if err = addOpUpdateKeyspaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateKeyspace(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateKeyspace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateKeyspace",
	}
}
