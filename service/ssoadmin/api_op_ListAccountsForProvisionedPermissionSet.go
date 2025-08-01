// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the Amazon Web Services accounts where the specified permission set
// is provisioned.
func (c *Client) ListAccountsForProvisionedPermissionSet(ctx context.Context, params *ListAccountsForProvisionedPermissionSetInput, optFns ...func(*Options)) (*ListAccountsForProvisionedPermissionSetOutput, error) {
	if params == nil {
		params = &ListAccountsForProvisionedPermissionSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccountsForProvisionedPermissionSet", params, optFns, c.addOperationListAccountsForProvisionedPermissionSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccountsForProvisionedPermissionSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccountsForProvisionedPermissionSetInput struct {

	// The ARN of the IAM Identity Center instance under which the operation will be
	// executed. For more information about ARNs, see Amazon Resource Names (ARNs) and Amazon Web Services Service Namespacesin the Amazon Web Services
	// General Reference.
	//
	// This member is required.
	InstanceArn *string

	// The ARN of the PermissionSet from which the associated Amazon Web Services accounts will be
	// listed.
	//
	// This member is required.
	PermissionSetArn *string

	// The maximum number of results to display for the PermissionSet.
	MaxResults *int32

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string

	// The permission set provisioning status for an Amazon Web Services account.
	ProvisioningStatus types.ProvisioningStatus

	noSmithyDocumentSerde
}

type ListAccountsForProvisionedPermissionSetOutput struct {

	// The list of Amazon Web Services AccountIds .
	AccountIds []string

	// The pagination token for the list API. Initially the value is null. Use the
	// output of previous API calls to make subsequent calls.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAccountsForProvisionedPermissionSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAccountsForProvisionedPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAccountsForProvisionedPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAccountsForProvisionedPermissionSet"); err != nil {
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
	if err = addOpListAccountsForProvisionedPermissionSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccountsForProvisionedPermissionSet(options.Region), middleware.Before); err != nil {
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

// ListAccountsForProvisionedPermissionSetPaginatorOptions is the paginator
// options for ListAccountsForProvisionedPermissionSet
type ListAccountsForProvisionedPermissionSetPaginatorOptions struct {
	// The maximum number of results to display for the PermissionSet.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccountsForProvisionedPermissionSetPaginator is a paginator for
// ListAccountsForProvisionedPermissionSet
type ListAccountsForProvisionedPermissionSetPaginator struct {
	options   ListAccountsForProvisionedPermissionSetPaginatorOptions
	client    ListAccountsForProvisionedPermissionSetAPIClient
	params    *ListAccountsForProvisionedPermissionSetInput
	nextToken *string
	firstPage bool
}

// NewListAccountsForProvisionedPermissionSetPaginator returns a new
// ListAccountsForProvisionedPermissionSetPaginator
func NewListAccountsForProvisionedPermissionSetPaginator(client ListAccountsForProvisionedPermissionSetAPIClient, params *ListAccountsForProvisionedPermissionSetInput, optFns ...func(*ListAccountsForProvisionedPermissionSetPaginatorOptions)) *ListAccountsForProvisionedPermissionSetPaginator {
	if params == nil {
		params = &ListAccountsForProvisionedPermissionSetInput{}
	}

	options := ListAccountsForProvisionedPermissionSetPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccountsForProvisionedPermissionSetPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccountsForProvisionedPermissionSetPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccountsForProvisionedPermissionSet page.
func (p *ListAccountsForProvisionedPermissionSetPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccountsForProvisionedPermissionSetOutput, error) {
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
	result, err := p.client.ListAccountsForProvisionedPermissionSet(ctx, &params, optFns...)
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

// ListAccountsForProvisionedPermissionSetAPIClient is a client that implements
// the ListAccountsForProvisionedPermissionSet operation.
type ListAccountsForProvisionedPermissionSetAPIClient interface {
	ListAccountsForProvisionedPermissionSet(context.Context, *ListAccountsForProvisionedPermissionSetInput, ...func(*Options)) (*ListAccountsForProvisionedPermissionSetOutput, error)
}

var _ ListAccountsForProvisionedPermissionSetAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAccountsForProvisionedPermissionSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAccountsForProvisionedPermissionSet",
	}
}
