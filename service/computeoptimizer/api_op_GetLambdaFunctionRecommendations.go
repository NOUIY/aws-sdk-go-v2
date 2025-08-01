// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns Lambda function recommendations.
//
// Compute Optimizer generates recommendations for functions that meet a specific
// set of requirements. For more information, see the [Supported resources and requirements]in the Compute Optimizer
// User Guide.
//
// [Supported resources and requirements]: https://docs.aws.amazon.com/compute-optimizer/latest/ug/requirements.html
func (c *Client) GetLambdaFunctionRecommendations(ctx context.Context, params *GetLambdaFunctionRecommendationsInput, optFns ...func(*Options)) (*GetLambdaFunctionRecommendationsOutput, error) {
	if params == nil {
		params = &GetLambdaFunctionRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLambdaFunctionRecommendations", params, optFns, c.addOperationGetLambdaFunctionRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLambdaFunctionRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLambdaFunctionRecommendationsInput struct {

	// The ID of the Amazon Web Services account for which to return function
	// recommendations.
	//
	// If your account is the management account of an organization, use this
	// parameter to specify the member account for which you want to return function
	// recommendations.
	//
	// Only one account ID can be specified per request.
	AccountIds []string

	// An array of objects to specify a filter that returns a more specific list of
	// function recommendations.
	Filters []types.LambdaFunctionRecommendationFilter

	// The Amazon Resource Name (ARN) of the functions for which to return
	// recommendations.
	//
	// You can specify a qualified or unqualified ARN. If you specify an unqualified
	// ARN without a function version suffix, Compute Optimizer will return
	// recommendations for the latest ( $LATEST ) version of the function. If you
	// specify a qualified ARN with a version suffix, Compute Optimizer will return
	// recommendations for the specified function version. For more information about
	// using function versions, see [Using versions]in the Lambda Developer Guide.
	//
	// [Using versions]: https://docs.aws.amazon.com/lambda/latest/dg/configuration-versions.html#versioning-versions-using
	FunctionArns []string

	// The maximum number of function recommendations to return with a single request.
	//
	// To retrieve the remaining results, make another request with the returned
	// nextToken value.
	MaxResults *int32

	// The token to advance to the next page of function recommendations.
	NextToken *string

	noSmithyDocumentSerde
}

type GetLambdaFunctionRecommendationsOutput struct {

	// An array of objects that describe function recommendations.
	LambdaFunctionRecommendations []types.LambdaFunctionRecommendation

	// The token to use to advance to the next page of function recommendations.
	//
	// This value is null when there are no more pages of function recommendations to
	// return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLambdaFunctionRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetLambdaFunctionRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetLambdaFunctionRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetLambdaFunctionRecommendations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLambdaFunctionRecommendations(options.Region), middleware.Before); err != nil {
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

// GetLambdaFunctionRecommendationsPaginatorOptions is the paginator options for
// GetLambdaFunctionRecommendations
type GetLambdaFunctionRecommendationsPaginatorOptions struct {
	// The maximum number of function recommendations to return with a single request.
	//
	// To retrieve the remaining results, make another request with the returned
	// nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetLambdaFunctionRecommendationsPaginator is a paginator for
// GetLambdaFunctionRecommendations
type GetLambdaFunctionRecommendationsPaginator struct {
	options   GetLambdaFunctionRecommendationsPaginatorOptions
	client    GetLambdaFunctionRecommendationsAPIClient
	params    *GetLambdaFunctionRecommendationsInput
	nextToken *string
	firstPage bool
}

// NewGetLambdaFunctionRecommendationsPaginator returns a new
// GetLambdaFunctionRecommendationsPaginator
func NewGetLambdaFunctionRecommendationsPaginator(client GetLambdaFunctionRecommendationsAPIClient, params *GetLambdaFunctionRecommendationsInput, optFns ...func(*GetLambdaFunctionRecommendationsPaginatorOptions)) *GetLambdaFunctionRecommendationsPaginator {
	if params == nil {
		params = &GetLambdaFunctionRecommendationsInput{}
	}

	options := GetLambdaFunctionRecommendationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetLambdaFunctionRecommendationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetLambdaFunctionRecommendationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetLambdaFunctionRecommendations page.
func (p *GetLambdaFunctionRecommendationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetLambdaFunctionRecommendationsOutput, error) {
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
	result, err := p.client.GetLambdaFunctionRecommendations(ctx, &params, optFns...)
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

// GetLambdaFunctionRecommendationsAPIClient is a client that implements the
// GetLambdaFunctionRecommendations operation.
type GetLambdaFunctionRecommendationsAPIClient interface {
	GetLambdaFunctionRecommendations(context.Context, *GetLambdaFunctionRecommendationsInput, ...func(*Options)) (*GetLambdaFunctionRecommendationsOutput, error)
}

var _ GetLambdaFunctionRecommendationsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetLambdaFunctionRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetLambdaFunctionRecommendations",
	}
}
