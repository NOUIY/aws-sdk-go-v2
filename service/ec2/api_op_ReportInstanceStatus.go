// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Submits feedback about the status of an instance. The instance must be in the
// running state. If your experience with the instance differs from the instance
// status returned by DescribeInstanceStatus, use ReportInstanceStatus to report your experience with the instance. Amazon
// EC2 collects this information to improve the accuracy of status checks.
//
// Use of this action does not change the value returned by DescribeInstanceStatus.
func (c *Client) ReportInstanceStatus(ctx context.Context, params *ReportInstanceStatusInput, optFns ...func(*Options)) (*ReportInstanceStatusOutput, error) {
	if params == nil {
		params = &ReportInstanceStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReportInstanceStatus", params, optFns, c.addOperationReportInstanceStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReportInstanceStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReportInstanceStatusInput struct {

	// The instances.
	//
	// This member is required.
	Instances []string

	// The reason codes that describe the health state of your instance.
	//
	//   - instance-stuck-in-state : My instance is stuck in a state.
	//
	//   - unresponsive : My instance is unresponsive.
	//
	//   - not-accepting-credentials : My instance is not accepting my credentials.
	//
	//   - password-not-available : A password is not available for my instance.
	//
	//   - performance-network : My instance is experiencing performance problems that
	//   I believe are network related.
	//
	//   - performance-instance-store : My instance is experiencing performance
	//   problems that I believe are related to the instance stores.
	//
	//   - performance-ebs-volume : My instance is experiencing performance problems
	//   that I believe are related to an EBS volume.
	//
	//   - performance-other : My instance is experiencing performance problems.
	//
	//   - other : [explain using the description parameter]
	//
	// This member is required.
	ReasonCodes []types.ReportInstanceReasonCodes

	// The status of all instances listed.
	//
	// This member is required.
	Status types.ReportStatusType

	// Descriptive text about the health state of your instance.
	//
	// Deprecated: This member has been deprecated
	Description *string

	// Checks whether you have the required permissions for the operation, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The time at which the reported instance health state ended.
	EndTime *time.Time

	// The time at which the reported instance health state began.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type ReportInstanceStatusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReportInstanceStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpReportInstanceStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpReportInstanceStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ReportInstanceStatus"); err != nil {
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
	if err = addOpReportInstanceStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReportInstanceStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opReportInstanceStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ReportInstanceStatus",
	}
}
