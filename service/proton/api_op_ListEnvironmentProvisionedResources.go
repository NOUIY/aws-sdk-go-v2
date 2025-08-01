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

// List the provisioned resources for your environment.
func (c *Client) ListEnvironmentProvisionedResources(ctx context.Context, params *ListEnvironmentProvisionedResourcesInput, optFns ...func(*Options)) (*ListEnvironmentProvisionedResourcesOutput, error) {
	if params == nil {
		params = &ListEnvironmentProvisionedResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEnvironmentProvisionedResources", params, optFns, c.addOperationListEnvironmentProvisionedResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEnvironmentProvisionedResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEnvironmentProvisionedResourcesInput struct {

	// The environment name.
	//
	// This member is required.
	EnvironmentName *string

	// A token that indicates the location of the next environment provisioned
	// resource in the array of environment provisioned resources, after the list of
	// environment provisioned resources that was previously requested.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEnvironmentProvisionedResourcesOutput struct {

	// An array of environment provisioned resources.
	//
	// This member is required.
	ProvisionedResources []types.ProvisionedResource

	// A token that indicates the location of the next environment provisioned
	// resource in the array of provisioned resources, after the current requested list
	// of environment provisioned resources.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEnvironmentProvisionedResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListEnvironmentProvisionedResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListEnvironmentProvisionedResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEnvironmentProvisionedResources"); err != nil {
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
	if err = addOpListEnvironmentProvisionedResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEnvironmentProvisionedResources(options.Region), middleware.Before); err != nil {
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

// ListEnvironmentProvisionedResourcesPaginatorOptions is the paginator options
// for ListEnvironmentProvisionedResources
type ListEnvironmentProvisionedResourcesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEnvironmentProvisionedResourcesPaginator is a paginator for
// ListEnvironmentProvisionedResources
type ListEnvironmentProvisionedResourcesPaginator struct {
	options   ListEnvironmentProvisionedResourcesPaginatorOptions
	client    ListEnvironmentProvisionedResourcesAPIClient
	params    *ListEnvironmentProvisionedResourcesInput
	nextToken *string
	firstPage bool
}

// NewListEnvironmentProvisionedResourcesPaginator returns a new
// ListEnvironmentProvisionedResourcesPaginator
func NewListEnvironmentProvisionedResourcesPaginator(client ListEnvironmentProvisionedResourcesAPIClient, params *ListEnvironmentProvisionedResourcesInput, optFns ...func(*ListEnvironmentProvisionedResourcesPaginatorOptions)) *ListEnvironmentProvisionedResourcesPaginator {
	if params == nil {
		params = &ListEnvironmentProvisionedResourcesInput{}
	}

	options := ListEnvironmentProvisionedResourcesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEnvironmentProvisionedResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEnvironmentProvisionedResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEnvironmentProvisionedResources page.
func (p *ListEnvironmentProvisionedResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEnvironmentProvisionedResourcesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListEnvironmentProvisionedResources(ctx, &params, optFns...)
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

// ListEnvironmentProvisionedResourcesAPIClient is a client that implements the
// ListEnvironmentProvisionedResources operation.
type ListEnvironmentProvisionedResourcesAPIClient interface {
	ListEnvironmentProvisionedResources(context.Context, *ListEnvironmentProvisionedResourcesInput, ...func(*Options)) (*ListEnvironmentProvisionedResourcesOutput, error)
}

var _ ListEnvironmentProvisionedResourcesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListEnvironmentProvisionedResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEnvironmentProvisionedResources",
	}
}
