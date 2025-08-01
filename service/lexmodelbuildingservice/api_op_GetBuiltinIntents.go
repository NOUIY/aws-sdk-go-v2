// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of built-in intents that meet the specified criteria.
//
// This operation requires permission for the lex:GetBuiltinIntents action.
func (c *Client) GetBuiltinIntents(ctx context.Context, params *GetBuiltinIntentsInput, optFns ...func(*Options)) (*GetBuiltinIntentsOutput, error) {
	if params == nil {
		params = &GetBuiltinIntentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBuiltinIntents", params, optFns, c.addOperationGetBuiltinIntentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBuiltinIntentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBuiltinIntentsInput struct {

	// A list of locales that the intent supports.
	Locale types.Locale

	// The maximum number of intents to return in the response. The default is 10.
	MaxResults *int32

	// A pagination token that fetches the next page of intents. If this API call is
	// truncated, Amazon Lex returns a pagination token in the response. To fetch the
	// next page of intents, use the pagination token in the next request.
	NextToken *string

	// Substring to match in built-in intent signatures. An intent will be returned if
	// any part of its signature matches the substring. For example, "xyz" matches both
	// "xyzabc" and "abcxyz." To find the signature for an intent, see [Standard Built-in Intents]in the Alexa
	// Skills Kit.
	//
	// [Standard Built-in Intents]: https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/standard-intents
	SignatureContains *string

	noSmithyDocumentSerde
}

type GetBuiltinIntentsOutput struct {

	// An array of builtinIntentMetadata objects, one for each intent in the response.
	Intents []types.BuiltinIntentMetadata

	// A pagination token that fetches the next page of intents. If the response to
	// this API call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of intents, specify the pagination token in the
	// next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBuiltinIntentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBuiltinIntents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBuiltinIntents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetBuiltinIntents"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBuiltinIntents(options.Region), middleware.Before); err != nil {
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

// GetBuiltinIntentsPaginatorOptions is the paginator options for GetBuiltinIntents
type GetBuiltinIntentsPaginatorOptions struct {
	// The maximum number of intents to return in the response. The default is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetBuiltinIntentsPaginator is a paginator for GetBuiltinIntents
type GetBuiltinIntentsPaginator struct {
	options   GetBuiltinIntentsPaginatorOptions
	client    GetBuiltinIntentsAPIClient
	params    *GetBuiltinIntentsInput
	nextToken *string
	firstPage bool
}

// NewGetBuiltinIntentsPaginator returns a new GetBuiltinIntentsPaginator
func NewGetBuiltinIntentsPaginator(client GetBuiltinIntentsAPIClient, params *GetBuiltinIntentsInput, optFns ...func(*GetBuiltinIntentsPaginatorOptions)) *GetBuiltinIntentsPaginator {
	if params == nil {
		params = &GetBuiltinIntentsInput{}
	}

	options := GetBuiltinIntentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetBuiltinIntentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetBuiltinIntentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetBuiltinIntents page.
func (p *GetBuiltinIntentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetBuiltinIntentsOutput, error) {
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
	result, err := p.client.GetBuiltinIntents(ctx, &params, optFns...)
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

// GetBuiltinIntentsAPIClient is a client that implements the GetBuiltinIntents
// operation.
type GetBuiltinIntentsAPIClient interface {
	GetBuiltinIntents(context.Context, *GetBuiltinIntentsInput, ...func(*Options)) (*GetBuiltinIntentsOutput, error)
}

var _ GetBuiltinIntentsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetBuiltinIntents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetBuiltinIntents",
	}
}
