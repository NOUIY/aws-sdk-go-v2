// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoverycluster

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycluster/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get the state for a routing control. A routing control is a simple on/off
// switch that you can use to route traffic to cells. When a routing control state
// is set to ON, traffic flows to a cell. When the state is set to OFF, traffic
// does not flow.
//
// Before you can create a routing control, you must first create a cluster, and
// then host the control in a control panel on the cluster. For more information,
// see [Create routing control structures]in the Amazon Route 53 Application Recovery Controller Developer Guide. You
// access one of the endpoints for the cluster to get or update the routing control
// state to redirect traffic for your application.
//
// You must specify Regional endpoints when you work with API cluster operations
// to get or update routing control states in Route 53 ARC.
//
// To see a code example for getting a routing control state, including accessing
// Regional cluster endpoints in sequence, see [API examples]in the Amazon Route 53 Application
// Recovery Controller Developer Guide.
//
// Learn more about working with routing controls in the following topics in the
// Amazon Route 53 Application Recovery Controller Developer Guide:
//
// [Viewing and updating routing control states]
//
// [Working with routing controls in Route 53 ARC]
//
// [API examples]: https://docs.aws.amazon.com/r53recovery/latest/dg/service_code_examples_actions.html
// [Working with routing controls in Route 53 ARC]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.html
// [Create routing control structures]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.create.html
// [Viewing and updating routing control states]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.update.html
func (c *Client) GetRoutingControlState(ctx context.Context, params *GetRoutingControlStateInput, optFns ...func(*Options)) (*GetRoutingControlStateOutput, error) {
	if params == nil {
		params = &GetRoutingControlStateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRoutingControlState", params, optFns, c.addOperationGetRoutingControlStateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRoutingControlStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRoutingControlStateInput struct {

	// The Amazon Resource Name (ARN) for the routing control that you want to get the
	// state for.
	//
	// This member is required.
	RoutingControlArn *string

	noSmithyDocumentSerde
}

type GetRoutingControlStateOutput struct {

	// The Amazon Resource Name (ARN) of the response.
	//
	// This member is required.
	RoutingControlArn *string

	// The state of the routing control.
	//
	// This member is required.
	RoutingControlState types.RoutingControlState

	// The routing control name.
	RoutingControlName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRoutingControlStateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetRoutingControlState{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetRoutingControlState{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetRoutingControlState"); err != nil {
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
	if err = addOpGetRoutingControlStateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRoutingControlState(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRoutingControlState(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetRoutingControlState",
	}
}
