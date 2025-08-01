// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the metadata for the rule groups that you have defined. Depending on
// your setting for max results and the number of rule groups, a single call might
// not return the full list.
func (c *Client) ListRuleGroups(ctx context.Context, params *ListRuleGroupsInput, optFns ...func(*Options)) (*ListRuleGroupsOutput, error) {
	if params == nil {
		params = &ListRuleGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRuleGroups", params, optFns, c.addOperationListRuleGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRuleGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRuleGroupsInput struct {

	// Indicates the general category of the Amazon Web Services managed rule group.
	ManagedType types.ResourceManagedType

	// The maximum number of objects that you want Network Firewall to return for this
	// request. If more objects are available, in the response, Network Firewall
	// provides a NextToken value that you can use in a subsequent call to get the
	// next batch of objects.
	MaxResults *int32

	// When you request a list of objects with a MaxResults setting, if the number of
	// objects that are still available for retrieval exceeds the maximum you
	// requested, Network Firewall returns a NextToken value in the response. To
	// retrieve the next batch of objects, use the token returned from the prior
	// request in your next request.
	NextToken *string

	// The scope of the request. The default setting of ACCOUNT or a setting of NULL
	// returns all of the rule groups in your account. A setting of MANAGED returns
	// all available managed rule groups.
	Scope types.ResourceManagedStatus

	// Indicates whether the rule group is stateless or stateful. If the rule group is
	// stateless, it contains stateless rules. If it is stateful, it contains stateful
	// rules.
	Type types.RuleGroupType

	noSmithyDocumentSerde
}

type ListRuleGroupsOutput struct {

	// When you request a list of objects with a MaxResults setting, if the number of
	// objects that are still available for retrieval exceeds the maximum you
	// requested, Network Firewall returns a NextToken value in the response. To
	// retrieve the next batch of objects, use the token returned from the prior
	// request in your next request.
	NextToken *string

	// The rule group metadata objects that you've defined. Depending on your setting
	// for max results and the number of rule groups, this might not be the full list.
	RuleGroups []types.RuleGroupMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRuleGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListRuleGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListRuleGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRuleGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRuleGroups(options.Region), middleware.Before); err != nil {
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

// ListRuleGroupsPaginatorOptions is the paginator options for ListRuleGroups
type ListRuleGroupsPaginatorOptions struct {
	// The maximum number of objects that you want Network Firewall to return for this
	// request. If more objects are available, in the response, Network Firewall
	// provides a NextToken value that you can use in a subsequent call to get the
	// next batch of objects.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRuleGroupsPaginator is a paginator for ListRuleGroups
type ListRuleGroupsPaginator struct {
	options   ListRuleGroupsPaginatorOptions
	client    ListRuleGroupsAPIClient
	params    *ListRuleGroupsInput
	nextToken *string
	firstPage bool
}

// NewListRuleGroupsPaginator returns a new ListRuleGroupsPaginator
func NewListRuleGroupsPaginator(client ListRuleGroupsAPIClient, params *ListRuleGroupsInput, optFns ...func(*ListRuleGroupsPaginatorOptions)) *ListRuleGroupsPaginator {
	if params == nil {
		params = &ListRuleGroupsInput{}
	}

	options := ListRuleGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRuleGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRuleGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRuleGroups page.
func (p *ListRuleGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRuleGroupsOutput, error) {
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
	result, err := p.client.ListRuleGroups(ctx, &params, optFns...)
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

// ListRuleGroupsAPIClient is a client that implements the ListRuleGroups
// operation.
type ListRuleGroupsAPIClient interface {
	ListRuleGroups(context.Context, *ListRuleGroupsInput, ...func(*Options)) (*ListRuleGroupsOutput, error)
}

var _ ListRuleGroupsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListRuleGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRuleGroups",
	}
}
