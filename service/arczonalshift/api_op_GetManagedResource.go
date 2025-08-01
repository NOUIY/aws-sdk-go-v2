// Code generated by smithy-go-codegen DO NOT EDIT.

package arczonalshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/arczonalshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get information about a resource that's been registered for zonal shifts with
// Amazon Application Recovery Controller in this Amazon Web Services Region.
// Resources that are registered for zonal shifts are managed resources in ARC. You
// can start zonal shifts and configure zonal autoshift for managed resources.
func (c *Client) GetManagedResource(ctx context.Context, params *GetManagedResourceInput, optFns ...func(*Options)) (*GetManagedResourceOutput, error) {
	if params == nil {
		params = &GetManagedResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetManagedResource", params, optFns, c.addOperationGetManagedResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetManagedResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetManagedResourceInput struct {

	// The identifier for the resource that Amazon Web Services shifts traffic for.
	// The identifier is the Amazon Resource Name (ARN) for the resource.
	//
	// Amazon Application Recovery Controller currently supports enabling the
	// following resources for zonal shift and zonal autoshift:
	//
	// [Amazon EC2 Auto Scaling groups]
	//
	// [Amazon Elastic Kubernetes Service]
	//
	// [Application Load Balancer]
	//
	// [Network Load Balancer]
	//
	// [Amazon EC2 Auto Scaling groups]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-shift.resource-types.ec2-auto-scaling-groups.html
	// [Amazon Elastic Kubernetes Service]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-shift.resource-types.eks.html
	// [Application Load Balancer]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-shift.resource-types.app-load-balancers.html
	// [Network Load Balancer]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-shift.resource-types.network-load-balancers.html
	//
	// This member is required.
	ResourceIdentifier *string

	noSmithyDocumentSerde
}

type GetManagedResourceOutput struct {

	// A collection of key-value pairs that indicate whether resources are active in
	// Availability Zones or not. The key name is the Availability Zone where the
	// resource is deployed. The value is 1 or 0.
	//
	// This member is required.
	AppliedWeights map[string]float32

	// The zonal shifts that are currently active for a resource.
	//
	// This member is required.
	ZonalShifts []types.ZonalShiftInResource

	// The Amazon Resource Name (ARN) for the resource.
	Arn *string

	// An array of the autoshifts that are active for the resource.
	Autoshifts []types.AutoshiftInResource

	// The name of the resource.
	Name *string

	// The practice run configuration for zonal autoshift that's associated with the
	// resource.
	PracticeRunConfiguration *types.PracticeRunConfiguration

	// The status for zonal autoshift for a resource. When the autoshift status is
	// ENABLED , Amazon Web Services shifts traffic for a resource away from an
	// Availability Zone, on your behalf, when Amazon Web Services determines that
	// there's an issue in the Availability Zone that could potentially affect
	// customers.
	ZonalAutoshiftStatus types.ZonalAutoshiftStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetManagedResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetManagedResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetManagedResource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetManagedResource"); err != nil {
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
	if err = addOpGetManagedResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetManagedResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetManagedResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetManagedResource",
	}
}
