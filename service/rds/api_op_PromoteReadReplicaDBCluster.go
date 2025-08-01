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

// Promotes a read replica DB cluster to a standalone DB cluster.
func (c *Client) PromoteReadReplicaDBCluster(ctx context.Context, params *PromoteReadReplicaDBClusterInput, optFns ...func(*Options)) (*PromoteReadReplicaDBClusterOutput, error) {
	if params == nil {
		params = &PromoteReadReplicaDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PromoteReadReplicaDBCluster", params, optFns, c.addOperationPromoteReadReplicaDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PromoteReadReplicaDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PromoteReadReplicaDBClusterInput struct {

	// The identifier of the DB cluster read replica to promote. This parameter isn't
	// case-sensitive.
	//
	// Constraints:
	//
	//   - Must match the identifier of an existing DB cluster read replica.
	//
	// Example: my-cluster-replica1
	//
	// This member is required.
	DBClusterIdentifier *string

	noSmithyDocumentSerde
}

type PromoteReadReplicaDBClusterOutput struct {

	// Contains the details of an Amazon Aurora DB cluster or Multi-AZ DB cluster.
	//
	// For an Amazon Aurora DB cluster, this data type is used as a response element
	// in the operations CreateDBCluster , DeleteDBCluster , DescribeDBClusters ,
	// FailoverDBCluster , ModifyDBCluster , PromoteReadReplicaDBCluster ,
	// RestoreDBClusterFromS3 , RestoreDBClusterFromSnapshot ,
	// RestoreDBClusterToPointInTime , StartDBCluster , and StopDBCluster .
	//
	// For a Multi-AZ DB cluster, this data type is used as a response element in the
	// operations CreateDBCluster , DeleteDBCluster , DescribeDBClusters ,
	// FailoverDBCluster , ModifyDBCluster , RebootDBCluster ,
	// RestoreDBClusterFromSnapshot , and RestoreDBClusterToPointInTime .
	//
	// For more information on Amazon Aurora DB clusters, see [What is Amazon Aurora?] in the Amazon Aurora
	// User Guide.
	//
	// For more information on Multi-AZ DB clusters, see [Multi-AZ deployments with two readable standby DB instances] in the Amazon RDS User
	// Guide.
	//
	// [Multi-AZ deployments with two readable standby DB instances]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html
	// [What is Amazon Aurora?]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPromoteReadReplicaDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPromoteReadReplicaDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPromoteReadReplicaDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PromoteReadReplicaDBCluster"); err != nil {
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
	if err = addOpPromoteReadReplicaDBClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPromoteReadReplicaDBCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPromoteReadReplicaDBCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PromoteReadReplicaDBCluster",
	}
}
