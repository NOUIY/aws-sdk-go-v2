// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the default values for the quotas for the specified Amazon Web Services
// service. A default value does not reflect any quota increases.
func (c *Client) ListAWSDefaultServiceQuotas(ctx context.Context, params *ListAWSDefaultServiceQuotasInput, optFns ...func(*Options)) (*ListAWSDefaultServiceQuotasOutput, error) {
	if params == nil {
		params = &ListAWSDefaultServiceQuotasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAWSDefaultServiceQuotas", params, optFns, c.addOperationListAWSDefaultServiceQuotasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAWSDefaultServiceQuotasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAWSDefaultServiceQuotasInput struct {

	// Specifies the service identifier. To find the service code value for an Amazon
	// Web Services service, use the ListServicesoperation.
	//
	// This member is required.
	ServiceCode *string

	// Specifies the maximum number of results that you want included on each page of
	// the response. If you do not include this parameter, it defaults to a value
	// appropriate to the operation. If additional items exist beyond those included in
	// the current response, the NextToken response element is present and has a value
	// (is not null). Include that value as the NextToken request parameter in the
	// next call to the operation to get the next part of the results.
	//
	// An API operation can return fewer results than the maximum even when there are
	// more results available. You should check NextToken after every operation to
	// ensure that you receive all of the results.
	MaxResults *int32

	// Specifies a value for receiving additional results after you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAWSDefaultServiceQuotasOutput struct {

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null .
	NextToken *string

	// Information about the quotas.
	Quotas []types.ServiceQuota

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAWSDefaultServiceQuotasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAWSDefaultServiceQuotas{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAWSDefaultServiceQuotas{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAWSDefaultServiceQuotas"); err != nil {
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
	if err = addOpListAWSDefaultServiceQuotasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAWSDefaultServiceQuotas(options.Region), middleware.Before); err != nil {
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

// ListAWSDefaultServiceQuotasPaginatorOptions is the paginator options for
// ListAWSDefaultServiceQuotas
type ListAWSDefaultServiceQuotasPaginatorOptions struct {
	// Specifies the maximum number of results that you want included on each page of
	// the response. If you do not include this parameter, it defaults to a value
	// appropriate to the operation. If additional items exist beyond those included in
	// the current response, the NextToken response element is present and has a value
	// (is not null). Include that value as the NextToken request parameter in the
	// next call to the operation to get the next part of the results.
	//
	// An API operation can return fewer results than the maximum even when there are
	// more results available. You should check NextToken after every operation to
	// ensure that you receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAWSDefaultServiceQuotasPaginator is a paginator for
// ListAWSDefaultServiceQuotas
type ListAWSDefaultServiceQuotasPaginator struct {
	options   ListAWSDefaultServiceQuotasPaginatorOptions
	client    ListAWSDefaultServiceQuotasAPIClient
	params    *ListAWSDefaultServiceQuotasInput
	nextToken *string
	firstPage bool
}

// NewListAWSDefaultServiceQuotasPaginator returns a new
// ListAWSDefaultServiceQuotasPaginator
func NewListAWSDefaultServiceQuotasPaginator(client ListAWSDefaultServiceQuotasAPIClient, params *ListAWSDefaultServiceQuotasInput, optFns ...func(*ListAWSDefaultServiceQuotasPaginatorOptions)) *ListAWSDefaultServiceQuotasPaginator {
	if params == nil {
		params = &ListAWSDefaultServiceQuotasInput{}
	}

	options := ListAWSDefaultServiceQuotasPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAWSDefaultServiceQuotasPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAWSDefaultServiceQuotasPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAWSDefaultServiceQuotas page.
func (p *ListAWSDefaultServiceQuotasPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAWSDefaultServiceQuotasOutput, error) {
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
	result, err := p.client.ListAWSDefaultServiceQuotas(ctx, &params, optFns...)
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

// ListAWSDefaultServiceQuotasAPIClient is a client that implements the
// ListAWSDefaultServiceQuotas operation.
type ListAWSDefaultServiceQuotasAPIClient interface {
	ListAWSDefaultServiceQuotas(context.Context, *ListAWSDefaultServiceQuotasInput, ...func(*Options)) (*ListAWSDefaultServiceQuotasOutput, error)
}

var _ ListAWSDefaultServiceQuotasAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListAWSDefaultServiceQuotas(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAWSDefaultServiceQuotas",
	}
}
