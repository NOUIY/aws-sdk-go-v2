// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves relevant passages or text excerpts given an input query.
//
// This API is similar to the [Query] API. However, by default, the Query API only
// returns excerpt passages of up to 100 token words. With the Retrieve API, you
// can retrieve longer passages of up to 200 token words and up to 100 semantically
// relevant passages. This doesn't include question-answer or FAQ type responses
// from your index. The passages are text excerpts that can be semantically
// extracted from multiple documents and multiple parts of the same document. If in
// extreme cases your documents produce zero passages using the Retrieve API, you
// can alternatively use the Query API and its types of responses.
//
// You can also do the following:
//
//   - Override boosting at the index level
//
//   - Filter based on document fields or attributes
//
//   - Filter based on the user or their group access to documents
//
//   - View the confidence score bucket for a retrieved passage result. The
//     confidence bucket provides a relative ranking that indicates how confident
//     Amazon Kendra is that the response is relevant to the query.
//
// Confidence score buckets are currently available only for English.
//
// You can also include certain fields in the response that might provide useful
// additional information.
//
// The Retrieve API shares the number of [query capacity units] that you set for your index. For more
// information on what's included in a single capacity unit and the default base
// capacity for an index, see [Adjusting capacity].
//
// If you're using an Amazon Kendra Gen AI Enterprise Edition index, you can only
// use ATTRIBUTE_FILTER to filter search results by user context. If you're using
// an Amazon Kendra Gen AI Enterprise Edition index and you try to use USER_TOKEN
// to configure user context policy, Amazon Kendra returns a ValidationException
// error.
//
// [Adjusting capacity]: https://docs.aws.amazon.com/kendra/latest/dg/adjusting-capacity.html
// [Query]: https://docs.aws.amazon.com/kendra/latest/APIReference/API_Query.html
// [query capacity units]: https://docs.aws.amazon.com/kendra/latest/APIReference/API_CapacityUnitsConfiguration.html
func (c *Client) Retrieve(ctx context.Context, params *RetrieveInput, optFns ...func(*Options)) (*RetrieveOutput, error) {
	if params == nil {
		params = &RetrieveInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Retrieve", params, optFns, c.addOperationRetrieveMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RetrieveOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RetrieveInput struct {

	// The identifier of the index to retrieve relevant passages for the search.
	//
	// This member is required.
	IndexId *string

	// The input query text to retrieve relevant passages for the search. Amazon
	// Kendra truncates queries at 30 token words, which excludes punctuation and stop
	// words. Truncation still applies if you use Boolean or more advanced, complex
	// queries. For example, Timeoff AND October AND Category:HR is counted as 3
	// tokens: timeoff , october , hr . For more information, see [Searching with advanced query syntax] in the Amazon
	// Kendra Developer Guide.
	//
	// [Searching with advanced query syntax]: https://docs.aws.amazon.com/kendra/latest/dg/searching-example.html#searching-index-query-syntax
	//
	// This member is required.
	QueryText *string

	// Filters search results by document fields/attributes. You can only provide one
	// attribute filter; however, the AndAllFilters , NotFilter , and OrAllFilters
	// parameters contain a list of other filters.
	//
	// The AttributeFilter parameter means you can create a set of filtering rules
	// that a document must satisfy to be included in the query results.
	//
	// For Amazon Kendra Gen AI Enterprise Edition indices use AttributeFilter to
	// enable document filtering for end users using _email_id or include public
	// documents ( _email_id=null ).
	AttributeFilter *types.AttributeFilter

	// Overrides relevance tuning configurations of fields/attributes set at the index
	// level.
	//
	// If you use this API to override the relevance tuning configured at the index
	// level, but there is no relevance tuning configured at the index level, then
	// Amazon Kendra does not apply any relevance tuning.
	//
	// If there is relevance tuning configured for fields at the index level, and you
	// use this API to override only some of these fields, then for the fields you did
	// not override, the importance is set to 1.
	DocumentRelevanceOverrideConfigurations []types.DocumentRelevanceConfiguration

	// Retrieved relevant passages are returned in pages the size of the PageSize
	// parameter. By default, Amazon Kendra returns the first page of results. Use this
	// parameter to get result pages after the first one.
	PageNumber *int32

	// Sets the number of retrieved relevant passages that are returned in each page
	// of results. The default page size is 10. The maximum number of results returned
	// is 100. If you ask for more than 100 results, only 100 are returned.
	PageSize *int32

	// A list of document fields/attributes to include in the response. You can limit
	// the response to include certain document fields. By default, all document fields
	// are included in the response.
	RequestedDocumentAttributes []string

	// The user context token or user and group information.
	UserContext *types.UserContext

	noSmithyDocumentSerde
}

type RetrieveOutput struct {

	// The identifier of query used for the search. You also use QueryId to identify
	// the search when using the [Submitfeedback]API.
	//
	// [Submitfeedback]: https://docs.aws.amazon.com/kendra/latest/APIReference/API_SubmitFeedback.html
	QueryId *string

	// The results of the retrieved relevant passages for the search.
	ResultItems []types.RetrieveResultItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRetrieveMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRetrieve{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRetrieve{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "Retrieve"); err != nil {
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
	if err = addOpRetrieveValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRetrieve(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRetrieve(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "Retrieve",
	}
}
