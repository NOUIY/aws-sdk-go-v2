// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about [custom key stores] in the account and Region.
//
// This operation is part of the custom key stores feature in KMS, which combines
// the convenience and extensive integration of KMS with the isolation and control
// of a key store that you own and manage.
//
// By default, this operation returns information about all custom key stores in
// the account and Region. To get only information about a particular custom key
// store, use either the CustomKeyStoreName or CustomKeyStoreId parameter (but not
// both).
//
// To determine whether the custom key store is connected to its CloudHSM cluster
// or external key store proxy, use the ConnectionState element in the response.
// If an attempt to connect the custom key store failed, the ConnectionState value
// is FAILED and the ConnectionErrorCode element in the response indicates the
// cause of the failure. For help interpreting the ConnectionErrorCode , see CustomKeyStoresListEntry.
//
// Custom key stores have a DISCONNECTED connection state if the key store has
// never been connected or you used the DisconnectCustomKeyStoreoperation to disconnect it. Otherwise, the
// connection state is CONNECTED. If your custom key store connection state is
// CONNECTED but you are having trouble using it, verify that the backing store is
// active and available. For an CloudHSM key store, verify that the associated
// CloudHSM cluster is active and contains the minimum number of HSMs required for
// the operation, if any. For an external key store, verify that the external key
// store proxy and its associated external key manager are reachable and enabled.
//
// For help repairing your CloudHSM key store, see the [Troubleshooting CloudHSM key stores]. For help repairing your
// external key store, see the [Troubleshooting external key stores]. Both topics are in the Key Management Service
// Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a custom key store
// in a different Amazon Web Services account.
//
// Required permissions: [kms:DescribeCustomKeyStores] (IAM policy)
//
// Related operations:
//
// # ConnectCustomKeyStore
//
// # CreateCustomKeyStore
//
// # DeleteCustomKeyStore
//
// # DisconnectCustomKeyStore
//
// # UpdateCustomKeyStore
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [kms:DescribeCustomKeyStores]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [custom key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html
// [Troubleshooting CloudHSM key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html
// [Troubleshooting external key stores]: https://docs.aws.amazon.com/kms/latest/developerguide/xks-troubleshooting.html
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/accessing-kms.html#programming-eventual-consistency
func (c *Client) DescribeCustomKeyStores(ctx context.Context, params *DescribeCustomKeyStoresInput, optFns ...func(*Options)) (*DescribeCustomKeyStoresOutput, error) {
	if params == nil {
		params = &DescribeCustomKeyStoresInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCustomKeyStores", params, optFns, c.addOperationDescribeCustomKeyStoresMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCustomKeyStoresOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCustomKeyStoresInput struct {

	// Gets only information about the specified custom key store. Enter the key store
	// ID.
	//
	// By default, this operation gets information about all custom key stores in the
	// account and Region. To limit the output to a particular custom key store,
	// provide either the CustomKeyStoreId or CustomKeyStoreName parameter, but not
	// both.
	CustomKeyStoreId *string

	// Gets only information about the specified custom key store. Enter the friendly
	// name of the custom key store.
	//
	// By default, this operation gets information about all custom key stores in the
	// account and Region. To limit the output to a particular custom key store,
	// provide either the CustomKeyStoreId or CustomKeyStoreName parameter, but not
	// both.
	CustomKeyStoreName *string

	// Use this parameter to specify the maximum number of items to return. When this
	// value is present, KMS does not return more than the specified number of items,
	// but it might return fewer.
	Limit *int32

	// Use this parameter in a subsequent request after you receive a response with
	// truncated results. Set it to the value of NextMarker from the truncated
	// response you just received.
	Marker *string

	noSmithyDocumentSerde
}

type DescribeCustomKeyStoresOutput struct {

	// Contains metadata about each custom key store.
	CustomKeyStores []types.CustomKeyStoresListEntry

	// When Truncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent request.
	NextMarker *string

	// A flag that indicates whether there are more items in the list. When this value
	// is true, the list in this response is truncated. To get more items, pass the
	// value of the NextMarker element in this response to the Marker parameter in a
	// subsequent request.
	Truncated bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCustomKeyStoresMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCustomKeyStores{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCustomKeyStores{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeCustomKeyStores"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCustomKeyStores(options.Region), middleware.Before); err != nil {
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

// DescribeCustomKeyStoresPaginatorOptions is the paginator options for
// DescribeCustomKeyStores
type DescribeCustomKeyStoresPaginatorOptions struct {
	// Use this parameter to specify the maximum number of items to return. When this
	// value is present, KMS does not return more than the specified number of items,
	// but it might return fewer.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeCustomKeyStoresPaginator is a paginator for DescribeCustomKeyStores
type DescribeCustomKeyStoresPaginator struct {
	options   DescribeCustomKeyStoresPaginatorOptions
	client    DescribeCustomKeyStoresAPIClient
	params    *DescribeCustomKeyStoresInput
	nextToken *string
	firstPage bool
}

// NewDescribeCustomKeyStoresPaginator returns a new
// DescribeCustomKeyStoresPaginator
func NewDescribeCustomKeyStoresPaginator(client DescribeCustomKeyStoresAPIClient, params *DescribeCustomKeyStoresInput, optFns ...func(*DescribeCustomKeyStoresPaginatorOptions)) *DescribeCustomKeyStoresPaginator {
	if params == nil {
		params = &DescribeCustomKeyStoresInput{}
	}

	options := DescribeCustomKeyStoresPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeCustomKeyStoresPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeCustomKeyStoresPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeCustomKeyStores page.
func (p *DescribeCustomKeyStoresPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeCustomKeyStoresOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeCustomKeyStores(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// DescribeCustomKeyStoresAPIClient is a client that implements the
// DescribeCustomKeyStores operation.
type DescribeCustomKeyStoresAPIClient interface {
	DescribeCustomKeyStores(context.Context, *DescribeCustomKeyStoresInput, ...func(*Options)) (*DescribeCustomKeyStoresOutput, error)
}

var _ DescribeCustomKeyStoresAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeCustomKeyStores(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeCustomKeyStores",
	}
}
