// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the updates of a directory for a particular update type.
func (c *Client) DescribeUpdateDirectory(ctx context.Context, params *DescribeUpdateDirectoryInput, optFns ...func(*Options)) (*DescribeUpdateDirectoryOutput, error) {
	if params == nil {
		params = &DescribeUpdateDirectoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeUpdateDirectory", params, optFns, c.addOperationDescribeUpdateDirectoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeUpdateDirectoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUpdateDirectoryInput struct {

	//  The unique identifier of the directory.
	//
	// This member is required.
	DirectoryId *string

	//  The type of updates you want to describe for the directory.
	//
	// This member is required.
	UpdateType types.UpdateType

	//  The DescribeUpdateDirectoryResult . NextToken value from a previous call to DescribeUpdateDirectory.
	// Pass null if this is the first call.
	NextToken *string

	//  The name of the Region.
	RegionName *string

	noSmithyDocumentSerde
}

type DescribeUpdateDirectoryOutput struct {

	//  If not null, more results are available. Pass this value for the NextToken
	// parameter.
	NextToken *string

	//  The list of update activities on a directory for the requested update type.
	UpdateActivities []types.UpdateInfoEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeUpdateDirectoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeUpdateDirectory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeUpdateDirectory{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeUpdateDirectory"); err != nil {
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
	if err = addOpDescribeUpdateDirectoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUpdateDirectory(options.Region), middleware.Before); err != nil {
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

// DescribeUpdateDirectoryPaginatorOptions is the paginator options for
// DescribeUpdateDirectory
type DescribeUpdateDirectoryPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeUpdateDirectoryPaginator is a paginator for DescribeUpdateDirectory
type DescribeUpdateDirectoryPaginator struct {
	options   DescribeUpdateDirectoryPaginatorOptions
	client    DescribeUpdateDirectoryAPIClient
	params    *DescribeUpdateDirectoryInput
	nextToken *string
	firstPage bool
}

// NewDescribeUpdateDirectoryPaginator returns a new
// DescribeUpdateDirectoryPaginator
func NewDescribeUpdateDirectoryPaginator(client DescribeUpdateDirectoryAPIClient, params *DescribeUpdateDirectoryInput, optFns ...func(*DescribeUpdateDirectoryPaginatorOptions)) *DescribeUpdateDirectoryPaginator {
	if params == nil {
		params = &DescribeUpdateDirectoryInput{}
	}

	options := DescribeUpdateDirectoryPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeUpdateDirectoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeUpdateDirectoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeUpdateDirectory page.
func (p *DescribeUpdateDirectoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeUpdateDirectoryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.DescribeUpdateDirectory(ctx, &params, optFns...)
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

// DescribeUpdateDirectoryAPIClient is a client that implements the
// DescribeUpdateDirectory operation.
type DescribeUpdateDirectoryAPIClient interface {
	DescribeUpdateDirectory(context.Context, *DescribeUpdateDirectoryInput, ...func(*Options)) (*DescribeUpdateDirectoryOutput, error)
}

var _ DescribeUpdateDirectoryAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opDescribeUpdateDirectory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeUpdateDirectory",
	}
}
