// Code generated by smithy-go-codegen DO NOT EDIT.

package arczonalshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/arczonalshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Start an on-demand practice run zonal shift in Amazon Application Recovery
// Controller. With zonal autoshift enabled, you can start an on-demand practice
// run to verify preparedness at any time. Amazon Web Services also runs automated
// practice runs about weekly when you have enabled zonal autoshift.
//
// For more information, see [Considerations when you configure zonal autoshift] in the Amazon Application Recovery Controller
// Developer Guide.
//
// [Considerations when you configure zonal autoshift]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-autoshift.considerations.html
func (c *Client) StartPracticeRun(ctx context.Context, params *StartPracticeRunInput, optFns ...func(*Options)) (*StartPracticeRunOutput, error) {
	if params == nil {
		params = &StartPracticeRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartPracticeRun", params, optFns, c.addOperationStartPracticeRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartPracticeRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartPracticeRunInput struct {

	// The Availability Zone (for example, use1-az1 ) that traffic is shifted away from
	// for the resource that you specify for the practice run.
	//
	// This member is required.
	AwayFrom *string

	// The initial comment that you enter about the practice run. Be aware that this
	// comment can be overwritten by Amazon Web Services if the automatic check for
	// balanced capacity fails. For more information, see [Capacity checks for practice runs]in the Amazon Application
	// Recovery Controller Developer Guide.
	//
	// [Capacity checks for practice runs]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-autoshift.how-it-works.capacity-check.html
	//
	// This member is required.
	Comment *string

	// The identifier for the resource that you want to start a practice run zonal
	// shift for. The identifier is the Amazon Resource Name (ARN) for the resource.
	//
	// This member is required.
	ResourceIdentifier *string

	noSmithyDocumentSerde
}

type StartPracticeRunOutput struct {

	// The Availability Zone (for example, use1-az1 ) that traffic is shifted away from
	// for the resource that you specify for the practice run.
	//
	// This member is required.
	AwayFrom *string

	// The initial comment that you enter about the practice run. Be aware that this
	// comment can be overwritten by Amazon Web Services if the automatic check for
	// balanced capacity fails. For more information, see [Capacity checks for practice runs]in the Amazon Application
	// Recovery Controller Developer Guide.
	//
	// [Capacity checks for practice runs]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-autoshift.how-it-works.capacity-check.html
	//
	// This member is required.
	Comment *string

	// The expiry time (expiration time) for an on-demand practice run zonal shift is
	// 30 minutes from the time when you start the practice run, unless you cancel it
	// before that time. However, be aware that the expiryTime field for practice run
	// zonal shifts always has a value of 1 minute.
	//
	// This member is required.
	ExpiryTime *time.Time

	// The identifier for the resource that you want to shift traffic for. The
	// identifier is the Amazon Resource Name (ARN) for the resource.
	//
	// This member is required.
	ResourceIdentifier *string

	// The time (UTC) when the zonal shift starts.
	//
	// This member is required.
	StartTime *time.Time

	// A status for the practice run (expected status is ACTIVE).
	//
	// This member is required.
	Status types.ZonalShiftStatus

	// The identifier of a practice run zonal shift.
	//
	// This member is required.
	ZonalShiftId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartPracticeRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartPracticeRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartPracticeRun{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartPracticeRun"); err != nil {
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
	if err = addOpStartPracticeRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartPracticeRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartPracticeRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartPracticeRun",
	}
}
