// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Submits an Batch job from a job definition. Parameters that are specified
// during SubmitJoboverride parameters defined in the job definition. vCPU and memory
// requirements that are specified in the resourceRequirements objects in the job
// definition are the exception. They can't be overridden this way using the memory
// and vcpus parameters. Rather, you must specify updates to job definition
// parameters in a resourceRequirements object that's included in the
// containerOverrides parameter.
//
// Job queues with a scheduling policy are limited to 500 active share identifiers
// at a time.
//
// Jobs that run on Fargate resources can't be guaranteed to run for more than 14
// days. This is because, after 14 days, Fargate resources might become unavailable
// and job might be terminated.
func (c *Client) SubmitJob(ctx context.Context, params *SubmitJobInput, optFns ...func(*Options)) (*SubmitJobOutput, error) {
	if params == nil {
		params = &SubmitJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SubmitJob", params, optFns, c.addOperationSubmitJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SubmitJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for SubmitJob .
type SubmitJobInput struct {

	// The job definition used by this job. This value can be one of definition-name ,
	// definition-name:revision , or the Amazon Resource Name (ARN) for the job
	// definition, with or without the revision (
	// arn:aws:batch:region:account:job-definition/definition-name:revision , or
	// arn:aws:batch:region:account:job-definition/definition-name ).
	//
	// If the revision is not specified, then the latest active revision is used.
	//
	// This member is required.
	JobDefinition *string

	// The name of the job. It can be up to 128 letters long. The first character must
	// be alphanumeric, can contain uppercase and lowercase letters, numbers, hyphens
	// (-), and underscores (_).
	//
	// This member is required.
	JobName *string

	// The job queue where the job is submitted. You can specify either the name or
	// the Amazon Resource Name (ARN) of the queue.
	//
	// This member is required.
	JobQueue *string

	// The array properties for the submitted job, such as the size of the array. The
	// array size can be between 2 and 10,000. If you specify array properties for a
	// job, it becomes an array job. For more information, see [Array Jobs]in the Batch User Guide.
	//
	// [Array Jobs]: https://docs.aws.amazon.com/batch/latest/userguide/array_jobs.html
	ArrayProperties *types.ArrayProperties

	// An object that contains overrides for the consumable resources of a job.
	ConsumableResourcePropertiesOverride *types.ConsumableResourceProperties

	// An object with properties that override the defaults for the job definition
	// that specify the name of a container in the specified job definition and the
	// overrides it should receive. You can override the default command for a
	// container, which is specified in the job definition or the Docker image, with a
	// command override. You can also override existing environment variables on a
	// container or add new environment variables to it with an environment override.
	ContainerOverrides *types.ContainerOverrides

	// A list of dependencies for the job. A job can depend upon a maximum of 20 jobs.
	// You can specify a SEQUENTIAL type dependency without specifying a job ID for
	// array jobs so that each child array job completes sequentially, starting at
	// index 0. You can also specify an N_TO_N type dependency with a job ID for array
	// jobs. In that case, each index child of this job must wait for the corresponding
	// index child of each dependency to complete before it can begin.
	DependsOn []types.JobDependency

	// An object, with properties that override defaults for the job definition, can
	// only be specified for jobs that are run on Amazon ECS resources.
	EcsPropertiesOverride *types.EcsPropertiesOverride

	// An object, with properties that override defaults for the job definition, can
	// only be specified for jobs that are run on Amazon EKS resources.
	EksPropertiesOverride *types.EksPropertiesOverride

	// A list of node overrides in JSON format that specify the node range to target
	// and the container overrides for that node range.
	//
	// This parameter isn't applicable to jobs that are running on Fargate resources;
	// use containerOverrides instead.
	NodeOverrides *types.NodeOverrides

	// Additional parameters passed to the job that replace parameter substitution
	// placeholders that are set in the job definition. Parameters are specified as a
	// key and value pair mapping. Parameters in a SubmitJob request override any
	// corresponding parameter defaults from the job definition.
	Parameters map[string]string

	// Specifies whether to propagate the tags from the job or job definition to the
	// corresponding Amazon ECS task. If no value is specified, the tags aren't
	// propagated. Tags can only be propagated to the tasks during task creation. For
	// tags with the same name, job tags are given priority over job definitions tags.
	// If the total number of combined tags from the job and job definition is over 50,
	// the job is moved to the FAILED state. When specified, this overrides the tag
	// propagation setting in the job definition.
	PropagateTags *bool

	// The retry strategy to use for failed jobs from this SubmitJob operation. When a retry
	// strategy is specified here, it overrides the retry strategy defined in the job
	// definition.
	RetryStrategy *types.RetryStrategy

	// The scheduling priority for the job. This only affects jobs in job queues with
	// a fair-share policy. Jobs with a higher scheduling priority are scheduled before
	// jobs with a lower scheduling priority. This overrides any scheduling priority in
	// the job definition and works only within a single share identifier.
	//
	// The minimum supported value is 0 and the maximum supported value is 9999.
	SchedulingPriorityOverride *int32

	// The share identifier for the job. Don't specify this parameter if the job queue
	// doesn't have a fair-share scheduling policy. If the job queue has a fair-share
	// scheduling policy, then this parameter must be specified.
	//
	// This string is limited to 255 alphanumeric characters, and can be followed by
	// an asterisk (*).
	ShareIdentifier *string

	// The tags that you apply to the job request to help you categorize and organize
	// your resources. Each tag consists of a key and an optional value. For more
	// information, see [Tagging Amazon Web Services Resources]in Amazon Web Services General Reference.
	//
	// [Tagging Amazon Web Services Resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
	Tags map[string]string

	// The timeout configuration for this SubmitJob operation. You can specify a timeout
	// duration after which Batch terminates your jobs if they haven't finished. If a
	// job is terminated due to a timeout, it isn't retried. The minimum value for the
	// timeout is 60 seconds. This configuration overrides any timeout configuration
	// specified in the job definition. For array jobs, child jobs have the same
	// timeout configuration as the parent job. For more information, see [Job Timeouts]in the
	// Amazon Elastic Container Service Developer Guide.
	//
	// [Job Timeouts]: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/job_timeouts.html
	Timeout *types.JobTimeout

	noSmithyDocumentSerde
}

type SubmitJobOutput struct {

	// The unique identifier for the job.
	//
	// This member is required.
	JobId *string

	// The name of the job.
	//
	// This member is required.
	JobName *string

	// The Amazon Resource Name (ARN) for the job.
	JobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSubmitJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSubmitJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSubmitJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SubmitJob"); err != nil {
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
	if err = addOpSubmitJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSubmitJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSubmitJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SubmitJob",
	}
}
