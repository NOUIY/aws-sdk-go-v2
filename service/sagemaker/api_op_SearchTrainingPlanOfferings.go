// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Searches for available training plan offerings based on specified criteria.
//
//   - Users search for available plan offerings based on their requirements
//     (e.g., instance type, count, start time, duration).
//
//   - And then, they create a plan that best matches their needs using the ID of
//     the plan offering they want to use.
//
// For more information about how to reserve GPU capacity for your SageMaker
// training jobs or SageMaker HyperPod clusters using Amazon SageMaker Training
// Plan , see [CreateTrainingPlan].
//
// [CreateTrainingPlan]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingPlan.html
func (c *Client) SearchTrainingPlanOfferings(ctx context.Context, params *SearchTrainingPlanOfferingsInput, optFns ...func(*Options)) (*SearchTrainingPlanOfferingsOutput, error) {
	if params == nil {
		params = &SearchTrainingPlanOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchTrainingPlanOfferings", params, optFns, c.addOperationSearchTrainingPlanOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchTrainingPlanOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchTrainingPlanOfferingsInput struct {

	// The desired duration in hours for the training plan offerings.
	//
	// This member is required.
	DurationHours *int64

	// The target resources (e.g., SageMaker Training Jobs, SageMaker HyperPod) to
	// search for in the offerings.
	//
	// Training plans are specific to their target resource.
	//
	//   - A training plan designed for SageMaker training jobs can only be used to
	//   schedule and run training jobs.
	//
	//   - A training plan for HyperPod clusters can be used exclusively to provide
	//   compute resources to a cluster's instance group.
	//
	// This member is required.
	TargetResources []types.SageMakerResourceName

	// A filter to search for reserved capacity offerings with an end time before a
	// specified date.
	EndTimeBefore *time.Time

	// The number of instances you want to reserve in the training plan offerings.
	// This allows you to specify the quantity of compute resources needed for your
	// SageMaker training jobs or SageMaker HyperPod clusters, helping you find
	// reserved capacity offerings that match your requirements.
	InstanceCount *int32

	// The type of instance you want to search for in the available training plan
	// offerings. This field allows you to filter the search results based on the
	// specific compute resources you require for your SageMaker training jobs or
	// SageMaker HyperPod clusters. When searching for training plan offerings,
	// specifying the instance type helps you find Reserved Instances that match your
	// computational needs.
	InstanceType types.ReservedCapacityInstanceType

	// A filter to search for training plan offerings with a start time after a
	// specified date.
	StartTimeAfter *time.Time

	noSmithyDocumentSerde
}

type SearchTrainingPlanOfferingsOutput struct {

	// A list of training plan offerings that match the search criteria.
	//
	// This member is required.
	TrainingPlanOfferings []types.TrainingPlanOffering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchTrainingPlanOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchTrainingPlanOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchTrainingPlanOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchTrainingPlanOfferings"); err != nil {
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
	if err = addOpSearchTrainingPlanOfferingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchTrainingPlanOfferings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSearchTrainingPlanOfferings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchTrainingPlanOfferings",
	}
}
