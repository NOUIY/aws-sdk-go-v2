// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists jobs.
//
// Requires permission to access the [ListJobs] action.
//
// [ListJobs]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func (c *Client) ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) {
	if params == nil {
		params = &ListJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJobs", params, optFns, c.addOperationListJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJobsInput struct {

	// The maximum number of results to return per request.
	MaxResults *int32

	// The namespace used to indicate that a job is a customer-managed job.
	//
	// When you specify a value for this parameter, Amazon Web Services IoT Core sends
	// jobs notifications to MQTT topics that contain the value in the following
	// format.
	//
	//     $aws/things/THING_NAME/jobs/JOB_ID/notify-namespace-NAMESPACE_ID/
	//
	// The namespaceId feature is only supported by IoT Greengrass at this time. For
	// more information, see [Setting up IoT Greengrass core devices.]
	//
	// [Setting up IoT Greengrass core devices.]: https://docs.aws.amazon.com/greengrass/v2/developerguide/setting-up.html
	NamespaceId *string

	// The token to retrieve the next set of results.
	NextToken *string

	// An optional filter that lets you search for jobs that have the specified status.
	Status types.JobStatus

	// Specifies whether the job will continue to run (CONTINUOUS), or will be
	// complete after all those things specified as targets have completed the job
	// (SNAPSHOT). If continuous, the job may also be run on a thing when a change is
	// detected in a target. For example, a job will run on a thing when the thing is
	// added to a target group, even after the job was completed by all things
	// originally in the group.
	//
	// We recommend that you use continuous jobs instead of snapshot jobs for dynamic
	// thing group targets. By using continuous jobs, devices that join the group
	// receive the job execution even after the job has been created.
	TargetSelection types.TargetSelection

	// A filter that limits the returned jobs to those for the specified group.
	ThingGroupId *string

	// A filter that limits the returned jobs to those for the specified group.
	ThingGroupName *string

	noSmithyDocumentSerde
}

type ListJobsOutput struct {

	// A list of jobs.
	Jobs []types.JobSummary

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListJobs(options.Region), middleware.Before); err != nil {
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

// ListJobsPaginatorOptions is the paginator options for ListJobs
type ListJobsPaginatorOptions struct {
	// The maximum number of results to return per request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListJobsPaginator is a paginator for ListJobs
type ListJobsPaginator struct {
	options   ListJobsPaginatorOptions
	client    ListJobsAPIClient
	params    *ListJobsInput
	nextToken *string
	firstPage bool
}

// NewListJobsPaginator returns a new ListJobsPaginator
func NewListJobsPaginator(client ListJobsAPIClient, params *ListJobsInput, optFns ...func(*ListJobsPaginatorOptions)) *ListJobsPaginator {
	if params == nil {
		params = &ListJobsInput{}
	}

	options := ListJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListJobs page.
func (p *ListJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListJobsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListJobs(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListJobsAPIClient is a client that implements the ListJobs operation.
type ListJobsAPIClient interface {
	ListJobs(context.Context, *ListJobsInput, ...func(*Options)) (*ListJobsOutput, error)
}

var _ ListJobsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListJobs",
	}
}
