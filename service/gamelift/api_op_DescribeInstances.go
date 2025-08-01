// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about the EC2 instances in an Amazon GameLift Servers
// managed fleet, including instance ID, connection data, and status. You can use
// this operation with a multi-location fleet to get location-specific instance
// information. As an alternative, use the operations [https://docs.aws.amazon.com/gamelift/latest/apireference/API_ListCompute]and [https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeCompute] to retrieve information
// for compute resources, including EC2 and Anywhere fleets.
//
// You can call this operation in the following ways:
//
//   - To get information on all instances in a fleet's home Region, specify the
//     fleet ID.
//
//   - To get information on all instances in a fleet's remote location, specify
//     the fleet ID and location name.
//
//   - To get information on a specific instance in a fleet, specify the fleet ID
//     and instance ID.
//
// Use the pagination parameters to retrieve results as a set of sequential pages.
//
// If successful, this operation returns Instance objects for each requested
// instance, listed in no particular order. If you call this operation for an
// Anywhere fleet, you receive an InvalidRequestException.
//
// # Learn more
//
// [Remotely connect to fleet instances]
//
// [Debug fleet issues]
//
// # Related actions
//
// [All APIs by task]
//
// [https://docs.aws.amazon.com/gamelift/latest/apireference/API_ListCompute]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_ListCompute
// [Remotely connect to fleet instances]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-remote-access.html
// [Debug fleet issues]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating-debug.html
// [https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeCompute]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeCompute
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
func (c *Client) DescribeInstances(ctx context.Context, params *DescribeInstancesInput, optFns ...func(*Options)) (*DescribeInstancesOutput, error) {
	if params == nil {
		params = &DescribeInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstances", params, optFns, c.addOperationDescribeInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstancesInput struct {

	// A unique identifier for the fleet to retrieve instance information for. You can
	// use either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// A unique identifier for an instance to retrieve. Specify an instance ID or
	// leave blank to retrieve all instances in the fleet.
	InstanceId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// The name of a location to retrieve instance information for, in the form of an
	// Amazon Web Services Region code such as us-west-2 .
	Location *string

	// A token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this operation. To start at
	// the beginning of the result set, do not specify a value.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeInstancesOutput struct {

	// A collection of objects containing properties for each instance returned.
	Instances []types.Instance

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeInstances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeInstances"); err != nil {
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
	if err = addOpDescribeInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstances(options.Region), middleware.Before); err != nil {
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

// DescribeInstancesPaginatorOptions is the paginator options for DescribeInstances
type DescribeInstancesPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeInstancesPaginator is a paginator for DescribeInstances
type DescribeInstancesPaginator struct {
	options   DescribeInstancesPaginatorOptions
	client    DescribeInstancesAPIClient
	params    *DescribeInstancesInput
	nextToken *string
	firstPage bool
}

// NewDescribeInstancesPaginator returns a new DescribeInstancesPaginator
func NewDescribeInstancesPaginator(client DescribeInstancesAPIClient, params *DescribeInstancesInput, optFns ...func(*DescribeInstancesPaginatorOptions)) *DescribeInstancesPaginator {
	if params == nil {
		params = &DescribeInstancesInput{}
	}

	options := DescribeInstancesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeInstances page.
func (p *DescribeInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeInstancesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeInstances(ctx, &params, optFns...)
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

// DescribeInstancesAPIClient is a client that implements the DescribeInstances
// operation.
type DescribeInstancesAPIClient interface {
	DescribeInstances(context.Context, *DescribeInstancesInput, ...func(*Options)) (*DescribeInstancesOutput, error)
}

var _ DescribeInstancesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeInstances",
	}
}
