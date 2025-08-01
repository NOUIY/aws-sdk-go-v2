// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about workflow types in the specified domain. The results
// may be split into multiple pages that can be retrieved by making the call
// repeatedly.
//
// # Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF
// resources as follows:
//
//   - Use a Resource element with the domain name to limit the action to only
//     specified domains.
//
//   - Use an Action element to allow or deny permission to call this action.
//
//   - You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see [Using IAM to Manage Access to Amazon SWF Workflows]in the Amazon SWF Developer Guide.
//
// [Using IAM to Manage Access to Amazon SWF Workflows]: https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html
func (c *Client) ListWorkflowTypes(ctx context.Context, params *ListWorkflowTypesInput, optFns ...func(*Options)) (*ListWorkflowTypesOutput, error) {
	if params == nil {
		params = &ListWorkflowTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkflowTypes", params, optFns, c.addOperationListWorkflowTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkflowTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkflowTypesInput struct {

	// The name of the domain in which the workflow types have been registered.
	//
	// This member is required.
	Domain *string

	// Specifies the registration status of the workflow types to list.
	//
	// This member is required.
	RegistrationStatus types.RegistrationStatus

	// The maximum number of results that are returned per call. Use nextPageToken to
	// obtain further pages of results.
	MaximumPageSize int32

	// If specified, lists the workflow type with this name.
	Name *string

	// If NextPageToken is returned there are more results available. The value of
	// NextPageToken is a unique pagination token for each page. Make the call again
	// using the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return a 400 error: " Specified token has exceeded its
	// maximum lifetime ".
	//
	// The configured maximumPageSize determines how many results can be returned in a
	// single call.
	NextPageToken *string

	// When set to true , returns the results in reverse order. By default the results
	// are returned in ascending alphabetical order of the name of the workflow types.
	ReverseOrder bool

	noSmithyDocumentSerde
}

// Contains a paginated list of information structures about workflow types.
type ListWorkflowTypesOutput struct {

	// The list of workflow type information.
	//
	// This member is required.
	TypeInfos []types.WorkflowTypeInfo

	// If a NextPageToken was returned by a previous call, there are more results
	// available. To retrieve the next page of results, make the call again using the
	// returned token in nextPageToken . Keep all other arguments unchanged.
	//
	// The configured maximumPageSize determines how many results can be returned in a
	// single call.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkflowTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListWorkflowTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListWorkflowTypes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListWorkflowTypes"); err != nil {
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
	if err = addOpListWorkflowTypesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkflowTypes(options.Region), middleware.Before); err != nil {
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

// ListWorkflowTypesPaginatorOptions is the paginator options for ListWorkflowTypes
type ListWorkflowTypesPaginatorOptions struct {
	// The maximum number of results that are returned per call. Use nextPageToken to
	// obtain further pages of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkflowTypesPaginator is a paginator for ListWorkflowTypes
type ListWorkflowTypesPaginator struct {
	options   ListWorkflowTypesPaginatorOptions
	client    ListWorkflowTypesAPIClient
	params    *ListWorkflowTypesInput
	nextToken *string
	firstPage bool
}

// NewListWorkflowTypesPaginator returns a new ListWorkflowTypesPaginator
func NewListWorkflowTypesPaginator(client ListWorkflowTypesAPIClient, params *ListWorkflowTypesInput, optFns ...func(*ListWorkflowTypesPaginatorOptions)) *ListWorkflowTypesPaginator {
	if params == nil {
		params = &ListWorkflowTypesInput{}
	}

	options := ListWorkflowTypesPaginatorOptions{}
	if params.MaximumPageSize != 0 {
		options.Limit = params.MaximumPageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkflowTypesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextPageToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkflowTypesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkflowTypes page.
func (p *ListWorkflowTypesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkflowTypesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextPageToken = p.nextToken

	params.MaximumPageSize = p.options.Limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListWorkflowTypes(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListWorkflowTypesAPIClient is a client that implements the ListWorkflowTypes
// operation.
type ListWorkflowTypesAPIClient interface {
	ListWorkflowTypes(context.Context, *ListWorkflowTypesInput, ...func(*Options)) (*ListWorkflowTypesOutput, error)
}

var _ ListWorkflowTypesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListWorkflowTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListWorkflowTypes",
	}
}
