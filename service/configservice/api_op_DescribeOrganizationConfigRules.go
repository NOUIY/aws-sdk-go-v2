// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of organization Config rules.
//
// When you specify the limit and the next token, you receive a paginated response.
//
// Limit and next token are not applicable if you specify organization Config rule
// names. It is only applicable, when you request all the organization Config
// rules.
//
// # For accounts within an organization
//
// If you deploy an organizational rule or conformance pack in an organization
// administrator account, and then establish a delegated administrator and deploy
// an organizational rule or conformance pack in the delegated administrator
// account, you won't be able to see the organizational rule or conformance pack in
// the organization administrator account from the delegated administrator account
// or see the organizational rule or conformance pack in the delegated
// administrator account from organization administrator account. The
// DescribeOrganizationConfigRules and DescribeOrganizationConformancePacks APIs
// can only see and interact with the organization-related resource that were
// deployed from within the account calling those APIs.
func (c *Client) DescribeOrganizationConfigRules(ctx context.Context, params *DescribeOrganizationConfigRulesInput, optFns ...func(*Options)) (*DescribeOrganizationConfigRulesOutput, error) {
	if params == nil {
		params = &DescribeOrganizationConfigRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrganizationConfigRules", params, optFns, c.addOperationDescribeOrganizationConfigRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrganizationConfigRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationConfigRulesInput struct {

	// The maximum number of organization Config rules returned on each page. If you
	// do no specify a number, Config uses the default. The default is 100.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// The names of organization Config rules for which you want details. If you do
	// not specify any names, Config returns details for all your organization Config
	// rules.
	OrganizationConfigRuleNames []string

	noSmithyDocumentSerde
}

type DescribeOrganizationConfigRulesOutput struct {

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Returns a list of OrganizationConfigRule objects.
	OrganizationConfigRules []types.OrganizationConfigRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOrganizationConfigRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeOrganizationConfigRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeOrganizationConfigRules{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeOrganizationConfigRules"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrganizationConfigRules(options.Region), middleware.Before); err != nil {
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

// DescribeOrganizationConfigRulesPaginatorOptions is the paginator options for
// DescribeOrganizationConfigRules
type DescribeOrganizationConfigRulesPaginatorOptions struct {
	// The maximum number of organization Config rules returned on each page. If you
	// do no specify a number, Config uses the default. The default is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeOrganizationConfigRulesPaginator is a paginator for
// DescribeOrganizationConfigRules
type DescribeOrganizationConfigRulesPaginator struct {
	options   DescribeOrganizationConfigRulesPaginatorOptions
	client    DescribeOrganizationConfigRulesAPIClient
	params    *DescribeOrganizationConfigRulesInput
	nextToken *string
	firstPage bool
}

// NewDescribeOrganizationConfigRulesPaginator returns a new
// DescribeOrganizationConfigRulesPaginator
func NewDescribeOrganizationConfigRulesPaginator(client DescribeOrganizationConfigRulesAPIClient, params *DescribeOrganizationConfigRulesInput, optFns ...func(*DescribeOrganizationConfigRulesPaginatorOptions)) *DescribeOrganizationConfigRulesPaginator {
	if params == nil {
		params = &DescribeOrganizationConfigRulesInput{}
	}

	options := DescribeOrganizationConfigRulesPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeOrganizationConfigRulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeOrganizationConfigRulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeOrganizationConfigRules page.
func (p *DescribeOrganizationConfigRulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeOrganizationConfigRulesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeOrganizationConfigRules(ctx, &params, optFns...)
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

// DescribeOrganizationConfigRulesAPIClient is a client that implements the
// DescribeOrganizationConfigRules operation.
type DescribeOrganizationConfigRulesAPIClient interface {
	DescribeOrganizationConfigRules(context.Context, *DescribeOrganizationConfigRulesInput, ...func(*Options)) (*DescribeOrganizationConfigRulesOutput, error)
}

var _ DescribeOrganizationConfigRulesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeOrganizationConfigRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeOrganizationConfigRules",
	}
}
