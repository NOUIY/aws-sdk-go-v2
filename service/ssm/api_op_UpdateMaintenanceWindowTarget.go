// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the target of an existing maintenance window. You can change the
// following:
//
//   - Name
//
//   - Description
//
//   - Owner
//
//   - IDs for an ID target
//
//   - Tags for a Tag target
//
//   - From any supported tag type to another. The three supported tag types are
//     ID target, Tag target, and resource group. For more information, see Target.
//
// If a parameter is null, then the corresponding field isn't modified.
func (c *Client) UpdateMaintenanceWindowTarget(ctx context.Context, params *UpdateMaintenanceWindowTargetInput, optFns ...func(*Options)) (*UpdateMaintenanceWindowTargetOutput, error) {
	if params == nil {
		params = &UpdateMaintenanceWindowTargetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMaintenanceWindowTarget", params, optFns, c.addOperationUpdateMaintenanceWindowTargetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMaintenanceWindowTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMaintenanceWindowTargetInput struct {

	// The maintenance window ID with which to modify the target.
	//
	// This member is required.
	WindowId *string

	// The target ID to modify.
	//
	// This member is required.
	WindowTargetId *string

	// An optional description for the update.
	Description *string

	// A name for the update.
	Name *string

	// User-provided value that will be included in any Amazon CloudWatch Events
	// events raised while running tasks for these targets in this maintenance window.
	OwnerInformation *string

	// If True , then all fields that are required by the RegisterTargetWithMaintenanceWindow operation are also required
	// for this API request. Optional fields that aren't specified are set to null.
	Replace *bool

	// The targets to add or replace.
	Targets []types.Target

	noSmithyDocumentSerde
}

type UpdateMaintenanceWindowTargetOutput struct {

	// The updated description.
	Description *string

	// The updated name.
	Name *string

	// The updated owner.
	OwnerInformation *string

	// The updated targets.
	Targets []types.Target

	// The maintenance window ID specified in the update request.
	WindowId *string

	// The target ID specified in the update request.
	WindowTargetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMaintenanceWindowTargetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateMaintenanceWindowTarget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMaintenanceWindowTarget{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateMaintenanceWindowTarget"); err != nil {
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
	if err = addOpUpdateMaintenanceWindowTargetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMaintenanceWindowTarget(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateMaintenanceWindowTarget(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateMaintenanceWindowTarget",
	}
}
