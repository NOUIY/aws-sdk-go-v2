// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing protection group. A protection group is a grouping of
// protected resources so they can be handled as a collective. This resource
// grouping improves the accuracy of detection and reduces false positives.
func (c *Client) UpdateProtectionGroup(ctx context.Context, params *UpdateProtectionGroupInput, optFns ...func(*Options)) (*UpdateProtectionGroupOutput, error) {
	if params == nil {
		params = &UpdateProtectionGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateProtectionGroup", params, optFns, c.addOperationUpdateProtectionGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateProtectionGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProtectionGroupInput struct {

	// Defines how Shield combines resource data for the group in order to detect,
	// mitigate, and report events.
	//
	//   - Sum - Use the total traffic across the group. This is a good choice for
	//   most cases. Examples include Elastic IP addresses for EC2 instances that scale
	//   manually or automatically.
	//
	//   - Mean - Use the average of the traffic across the group. This is a good
	//   choice for resources that share traffic uniformly. Examples include accelerators
	//   and load balancers.
	//
	//   - Max - Use the highest traffic from each resource. This is useful for
	//   resources that don't share traffic and for resources that share that traffic in
	//   a non-uniform way. Examples include Amazon CloudFront distributions and origin
	//   resources for CloudFront distributions.
	//
	// This member is required.
	Aggregation types.ProtectionGroupAggregation

	// The criteria to use to choose the protected resources for inclusion in the
	// group. You can include all resources that have protections, provide a list of
	// resource Amazon Resource Names (ARNs), or include all resources of a specified
	// resource type.
	//
	// This member is required.
	Pattern types.ProtectionGroupPattern

	// The name of the protection group. You use this to identify the protection group
	// in lists and to manage the protection group, for example to update, delete, or
	// describe it.
	//
	// This member is required.
	ProtectionGroupId *string

	// The Amazon Resource Names (ARNs) of the resources to include in the protection
	// group. You must set this when you set Pattern to ARBITRARY and you must not set
	// it for any other Pattern setting.
	Members []string

	// The resource type to include in the protection group. All protected resources
	// of this type are included in the protection group. You must set this when you
	// set Pattern to BY_RESOURCE_TYPE and you must not set it for any other Pattern
	// setting.
	ResourceType types.ProtectedResourceType

	noSmithyDocumentSerde
}

type UpdateProtectionGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateProtectionGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateProtectionGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateProtectionGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateProtectionGroup"); err != nil {
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
	if err = addOpUpdateProtectionGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProtectionGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateProtectionGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateProtectionGroup",
	}
}
