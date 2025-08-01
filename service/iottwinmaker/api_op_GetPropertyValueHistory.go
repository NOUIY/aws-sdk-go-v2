// Code generated by smithy-go-codegen DO NOT EDIT.

package iottwinmaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iottwinmaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about the history of a time series property value for a
// component, component type, entity, or workspace.
//
// You must specify a value for workspaceId . For entity-specific queries, specify
// values for componentName and entityId . For cross-entity quries, specify a value
// for componentTypeId .
func (c *Client) GetPropertyValueHistory(ctx context.Context, params *GetPropertyValueHistoryInput, optFns ...func(*Options)) (*GetPropertyValueHistoryOutput, error) {
	if params == nil {
		params = &GetPropertyValueHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPropertyValueHistory", params, optFns, c.addOperationGetPropertyValueHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPropertyValueHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPropertyValueHistoryInput struct {

	// A list of properties whose value histories the request retrieves.
	//
	// This member is required.
	SelectedProperties []string

	// The ID of the workspace.
	//
	// This member is required.
	WorkspaceId *string

	// The name of the component.
	ComponentName *string

	// This string specifies the path to the composite component, starting from the
	// top-level component.
	ComponentPath *string

	// The ID of the component type.
	ComponentTypeId *string

	// The date and time of the latest property value to return.
	//
	// Deprecated: This field is deprecated and will throw an error in the future. Use
	// endTime instead.
	EndDateTime *time.Time

	// The ISO8601 DateTime of the latest property value to return.
	//
	// For more information about the ISO8601 DateTime format, see the data type [PropertyValue].
	//
	// [PropertyValue]: https://docs.aws.amazon.com/iot-twinmaker/latest/apireference/API_PropertyValue.html
	EndTime *string

	// The ID of the entity.
	EntityId *string

	// An object that specifies the interpolation type and the interval over which to
	// interpolate data.
	Interpolation *types.InterpolationParameters

	// The maximum number of results to return at one time. The default is 25.
	//
	// Valid Range: Minimum value of 1. Maximum value of 250.
	MaxResults *int32

	// The string that specifies the next page of results.
	NextToken *string

	// The time direction to use in the result order.
	OrderByTime types.OrderByTime

	// A list of objects that filter the property value history request.
	PropertyFilters []types.PropertyFilter

	// The date and time of the earliest property value to return.
	//
	// Deprecated: This field is deprecated and will throw an error in the future. Use
	// startTime instead.
	StartDateTime *time.Time

	// The ISO8601 DateTime of the earliest property value to return.
	//
	// For more information about the ISO8601 DateTime format, see the data type [PropertyValue].
	//
	// [PropertyValue]: https://docs.aws.amazon.com/iot-twinmaker/latest/apireference/API_PropertyValue.html
	StartTime *string

	noSmithyDocumentSerde
}

type GetPropertyValueHistoryOutput struct {

	// An object that maps strings to the property definitions in the component type.
	// Each string in the mapping must be unique to this object.
	//
	// This member is required.
	PropertyValues []types.PropertyValueHistory

	// The string that specifies the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPropertyValueHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPropertyValueHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPropertyValueHistory{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetPropertyValueHistory"); err != nil {
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
	if err = addEndpointPrefix_opGetPropertyValueHistoryMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetPropertyValueHistoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPropertyValueHistory(options.Region), middleware.Before); err != nil {
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

// GetPropertyValueHistoryPaginatorOptions is the paginator options for
// GetPropertyValueHistory
type GetPropertyValueHistoryPaginatorOptions struct {
	// The maximum number of results to return at one time. The default is 25.
	//
	// Valid Range: Minimum value of 1. Maximum value of 250.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetPropertyValueHistoryPaginator is a paginator for GetPropertyValueHistory
type GetPropertyValueHistoryPaginator struct {
	options   GetPropertyValueHistoryPaginatorOptions
	client    GetPropertyValueHistoryAPIClient
	params    *GetPropertyValueHistoryInput
	nextToken *string
	firstPage bool
}

// NewGetPropertyValueHistoryPaginator returns a new
// GetPropertyValueHistoryPaginator
func NewGetPropertyValueHistoryPaginator(client GetPropertyValueHistoryAPIClient, params *GetPropertyValueHistoryInput, optFns ...func(*GetPropertyValueHistoryPaginatorOptions)) *GetPropertyValueHistoryPaginator {
	if params == nil {
		params = &GetPropertyValueHistoryInput{}
	}

	options := GetPropertyValueHistoryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetPropertyValueHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetPropertyValueHistoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetPropertyValueHistory page.
func (p *GetPropertyValueHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetPropertyValueHistoryOutput, error) {
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
	result, err := p.client.GetPropertyValueHistory(ctx, &params, optFns...)
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

type endpointPrefix_opGetPropertyValueHistoryMiddleware struct {
}

func (*endpointPrefix_opGetPropertyValueHistoryMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetPropertyValueHistoryMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "data." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetPropertyValueHistoryMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetPropertyValueHistoryMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// GetPropertyValueHistoryAPIClient is a client that implements the
// GetPropertyValueHistory operation.
type GetPropertyValueHistoryAPIClient interface {
	GetPropertyValueHistory(context.Context, *GetPropertyValueHistoryInput, ...func(*Options)) (*GetPropertyValueHistoryOutput, error)
}

var _ GetPropertyValueHistoryAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opGetPropertyValueHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetPropertyValueHistory",
	}
}
