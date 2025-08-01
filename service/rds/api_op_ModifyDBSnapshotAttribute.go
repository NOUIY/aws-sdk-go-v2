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

// Adds an attribute and values to, or removes an attribute and values from, a
// manual DB snapshot.
//
// To share a manual DB snapshot with other Amazon Web Services accounts, specify
// restore as the AttributeName and use the ValuesToAdd parameter to add a list of
// IDs of the Amazon Web Services accounts that are authorized to restore the
// manual DB snapshot. Uses the value all to make the manual DB snapshot public,
// which means it can be copied or restored by all Amazon Web Services accounts.
//
// Don't add the all value for any manual DB snapshots that contain private
// information that you don't want available to all Amazon Web Services accounts.
//
// If the manual DB snapshot is encrypted, it can be shared, but only by
// specifying a list of authorized Amazon Web Services account IDs for the
// ValuesToAdd parameter. You can't use all as a value for that parameter in this
// case.
//
// To view which Amazon Web Services accounts have access to copy or restore a
// manual DB snapshot, or whether a manual DB snapshot public or private, use the DescribeDBSnapshotAttributes
// API operation. The accounts are returned as values for the restore attribute.
func (c *Client) ModifyDBSnapshotAttribute(ctx context.Context, params *ModifyDBSnapshotAttributeInput, optFns ...func(*Options)) (*ModifyDBSnapshotAttributeOutput, error) {
	if params == nil {
		params = &ModifyDBSnapshotAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBSnapshotAttribute", params, optFns, c.addOperationModifyDBSnapshotAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBSnapshotAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDBSnapshotAttributeInput struct {

	// The name of the DB snapshot attribute to modify.
	//
	// To manage authorization for other Amazon Web Services accounts to copy or
	// restore a manual DB snapshot, set this value to restore .
	//
	// To view the list of attributes available to modify, use the DescribeDBSnapshotAttributes API operation.
	//
	// This member is required.
	AttributeName *string

	// The identifier for the DB snapshot to modify the attributes for.
	//
	// This member is required.
	DBSnapshotIdentifier *string

	// A list of DB snapshot attributes to add to the attribute specified by
	// AttributeName .
	//
	// To authorize other Amazon Web Services accounts to copy or restore a manual
	// snapshot, set this list to include one or more Amazon Web Services account IDs,
	// or all to make the manual DB snapshot restorable by any Amazon Web Services
	// account. Do not add the all value for any manual DB snapshots that contain
	// private information that you don't want available to all Amazon Web Services
	// accounts.
	ValuesToAdd []string

	// A list of DB snapshot attributes to remove from the attribute specified by
	// AttributeName .
	//
	// To remove authorization for other Amazon Web Services accounts to copy or
	// restore a manual snapshot, set this list to include one or more Amazon Web
	// Services account identifiers, or all to remove authorization for any Amazon Web
	// Services account to copy or restore the DB snapshot. If you specify all , an
	// Amazon Web Services account whose account ID is explicitly added to the restore
	// attribute can still copy or restore the manual DB snapshot.
	ValuesToRemove []string

	noSmithyDocumentSerde
}

type ModifyDBSnapshotAttributeOutput struct {

	// Contains the results of a successful call to the DescribeDBSnapshotAttributes
	// API action.
	//
	// Manual DB snapshot attributes are used to authorize other Amazon Web Services
	// accounts to copy or restore a manual DB snapshot. For more information, see the
	// ModifyDBSnapshotAttribute API action.
	DBSnapshotAttributesResult *types.DBSnapshotAttributesResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyDBSnapshotAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBSnapshotAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBSnapshotAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyDBSnapshotAttribute"); err != nil {
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
	if err = addOpModifyDBSnapshotAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBSnapshotAttribute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyDBSnapshotAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyDBSnapshotAttribute",
	}
}
