// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List deployments. You can filter the result list by environment, service, or a
// single service instance.
func (c *Client) ListDeployments(ctx context.Context, params *ListDeploymentsInput, optFns ...func(*Options)) (*ListDeploymentsOutput, error) {
	if params == nil {
		params = &ListDeploymentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeployments", params, optFns, c.addOperationListDeploymentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeploymentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDeploymentsInput struct {

	// The name of a component for result list filtering. Proton returns deployments
	// associated with that component.
	ComponentName *string

	// The name of an environment for result list filtering. Proton returns
	// deployments associated with the environment.
	EnvironmentName *string

	// The maximum number of deployments to list.
	MaxResults *int32

	// A token that indicates the location of the next deployment in the array of
	// deployment, after the list of deployment that was previously requested.
	NextToken *string

	// The name of a service instance for result list filtering. Proton returns the
	// deployments associated with the service instance.
	ServiceInstanceName *string

	// The name of a service for result list filtering. Proton returns deployments
	// associated with service instances of the service.
	ServiceName *string

	noSmithyDocumentSerde
}

type ListDeploymentsOutput struct {

	// An array of deployment with summary data.
	//
	// This member is required.
	Deployments []types.DeploymentSummary

	// A token that indicates the location of the next deployment in the array of
	// deployment, after the current requested list of deployment.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDeploymentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListDeployments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListDeployments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDeployments"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDeployments(options.Region), middleware.Before); err != nil {
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

// ListDeploymentsPaginatorOptions is the paginator options for ListDeployments
type ListDeploymentsPaginatorOptions struct {
	// The maximum number of deployments to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDeploymentsPaginator is a paginator for ListDeployments
type ListDeploymentsPaginator struct {
	options   ListDeploymentsPaginatorOptions
	client    ListDeploymentsAPIClient
	params    *ListDeploymentsInput
	nextToken *string
	firstPage bool
}

// NewListDeploymentsPaginator returns a new ListDeploymentsPaginator
func NewListDeploymentsPaginator(client ListDeploymentsAPIClient, params *ListDeploymentsInput, optFns ...func(*ListDeploymentsPaginatorOptions)) *ListDeploymentsPaginator {
	if params == nil {
		params = &ListDeploymentsInput{}
	}

	options := ListDeploymentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDeploymentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDeploymentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDeployments page.
func (p *ListDeploymentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDeploymentsOutput, error) {
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
	result, err := p.client.ListDeployments(ctx, &params, optFns...)
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

// ListDeploymentsAPIClient is a client that implements the ListDeployments
// operation.
type ListDeploymentsAPIClient interface {
	ListDeployments(context.Context, *ListDeploymentsInput, ...func(*Options)) (*ListDeploymentsOutput, error)
}

var _ ListDeploymentsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListDeployments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDeployments",
	}
}
