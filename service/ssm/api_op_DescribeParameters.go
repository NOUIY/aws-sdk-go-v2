// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the parameters in your Amazon Web Services account or the parameters
// shared with you when you enable the [Shared]option.
//
// Request results are returned on a best-effort basis. If you specify MaxResults
// in the request, the response includes information up to the limit specified. The
// number of items returned, however, can be between zero and the value of
// MaxResults . If the service reaches an internal limit while processing the
// results, it stops the operation and returns the matching values up to that point
// and a NextToken . You can specify the NextToken in a subsequent call to get the
// next set of results.
//
// Parameter names can't contain spaces. The service removes any spaces specified
// for the beginning or end of a parameter name. If the specified name for a
// parameter contains spaces between characters, the request fails with a
// ValidationException error.
//
// If you change the KMS key alias for the KMS key used to encrypt a parameter,
// then you must also update the key alias the parameter uses to reference KMS.
// Otherwise, DescribeParameters retrieves whatever the original key alias was
// referencing.
//
// [Shared]: https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_DescribeParameters.html#systemsmanager-DescribeParameters-request-Shared
func (c *Client) DescribeParameters(ctx context.Context, params *DescribeParametersInput, optFns ...func(*Options)) (*DescribeParametersOutput, error) {
	if params == nil {
		params = &DescribeParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeParameters", params, optFns, c.addOperationDescribeParametersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeParametersInput struct {

	// This data type is deprecated. Instead, use ParameterFilters .
	Filters []types.ParametersFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// Filters to limit the request results.
	ParameterFilters []types.ParameterStringFilter

	// Lists parameters that are shared with you.
	//
	// By default when using this option, the command returns parameters that have
	// been shared using a standard Resource Access Manager Resource Share. In order
	// for a parameter that was shared using the PutResourcePolicycommand to be returned, the
	// associated RAM Resource Share Created From Policy must have been promoted to a
	// standard Resource Share using the RAM [PromoteResourceShareCreatedFromPolicy]API operation.
	//
	// For more information about sharing parameters, see [Working with shared parameters] in the Amazon Web Services
	// Systems Manager User Guide.
	//
	// [PromoteResourceShareCreatedFromPolicy]: https://docs.aws.amazon.com/ram/latest/APIReference/API_PromoteResourceShareCreatedFromPolicy.html
	// [Working with shared parameters]: https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-shared-parameters.html
	Shared *bool

	noSmithyDocumentSerde
}

type DescribeParametersOutput struct {

	// The token to use when requesting the next set of items.
	NextToken *string

	// Parameters returned by the request.
	Parameters []types.ParameterMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeParametersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeParameters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeParameters{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeParameters"); err != nil {
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
	if err = addOpDescribeParametersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeParameters(options.Region), middleware.Before); err != nil {
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

// DescribeParametersPaginatorOptions is the paginator options for
// DescribeParameters
type DescribeParametersPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeParametersPaginator is a paginator for DescribeParameters
type DescribeParametersPaginator struct {
	options   DescribeParametersPaginatorOptions
	client    DescribeParametersAPIClient
	params    *DescribeParametersInput
	nextToken *string
	firstPage bool
}

// NewDescribeParametersPaginator returns a new DescribeParametersPaginator
func NewDescribeParametersPaginator(client DescribeParametersAPIClient, params *DescribeParametersInput, optFns ...func(*DescribeParametersPaginatorOptions)) *DescribeParametersPaginator {
	if params == nil {
		params = &DescribeParametersInput{}
	}

	options := DescribeParametersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeParametersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeParametersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeParameters page.
func (p *DescribeParametersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeParametersOutput, error) {
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
	result, err := p.client.DescribeParameters(ctx, &params, optFns...)
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

// DescribeParametersAPIClient is a client that implements the DescribeParameters
// operation.
type DescribeParametersAPIClient interface {
	DescribeParameters(context.Context, *DescribeParametersInput, ...func(*Options)) (*DescribeParametersOutput, error)
}

var _ DescribeParametersAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeParameters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeParameters",
	}
}
