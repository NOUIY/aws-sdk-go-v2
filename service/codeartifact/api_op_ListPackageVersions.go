// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	Returns a list of [PackageVersionSummary] objects for package versions in a repository that match the
//
// request parameters. Package versions of all statuses will be returned by default
// when calling list-package-versions with no --status parameter.
//
// [PackageVersionSummary]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageVersionSummary.html
func (c *Client) ListPackageVersions(ctx context.Context, params *ListPackageVersionsInput, optFns ...func(*Options)) (*ListPackageVersionsOutput, error) {
	if params == nil {
		params = &ListPackageVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPackageVersions", params, optFns, c.addOperationListPackageVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPackageVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPackageVersionsInput struct {

	//  The name of the domain that contains the repository that contains the
	// requested package versions.
	//
	// This member is required.
	Domain *string

	//  The format of the package versions you want to list.
	//
	// This member is required.
	Format types.PackageFormat

	//  The name of the package for which you want to request package versions.
	//
	// This member is required.
	Package *string

	//  The name of the repository that contains the requested package versions.
	//
	// This member is required.
	Repository *string

	//  The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	//  The maximum number of results to return per page.
	MaxResults *int32

	// The namespace of the package that contains the requested package versions. The
	// package component that specifies its namespace depends on its type. For example:
	//
	// The namespace is required when deleting package versions of the following
	// formats:
	//
	//   - Maven
	//
	//   - Swift
	//
	//   - generic
	//
	//   - The namespace of a Maven package version is its groupId .
	//
	//   - The namespace of an npm or Swift package version is its scope .
	//
	//   - The namespace of a generic package is its namespace .
	//
	//   - Python, NuGet, Ruby, and Cargo package versions do not contain a
	//   corresponding component, package versions of those formats do not have a
	//   namespace.
	Namespace *string

	//  The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The originType used to filter package versions. Only package versions with the
	// provided originType will be returned.
	OriginType types.PackageVersionOriginType

	//  How to sort the requested list of package versions.
	SortBy types.PackageVersionSortType

	//  A string that filters the requested package versions by status.
	Status types.PackageVersionStatus

	noSmithyDocumentSerde
}

type ListPackageVersionsOutput struct {

	//  The default package version to display. This depends on the package format:
	//
	//   - For Maven and PyPI packages, it's the most recently published package
	//   version.
	//
	//   - For npm packages, it's the version referenced by the latest tag. If the
	//   latest tag is not set, it's the most recently published package version.
	DefaultDisplayVersion *string

	//  A format of the package.
	Format types.PackageFormat

	// The namespace of the package that contains the requested package versions. The
	// package component that specifies its namespace depends on its type. For example:
	//
	//   - The namespace of a Maven package version is its groupId .
	//
	//   - The namespace of an npm or Swift package version is its scope .
	//
	//   - The namespace of a generic package is its namespace .
	//
	//   - Python, NuGet, Ruby, and Cargo package versions do not contain a
	//   corresponding component, package versions of those formats do not have a
	//   namespace.
	Namespace *string

	//  If there are additional results, this is the token for the next set of
	// results.
	NextToken *string

	//  The name of the package.
	Package *string

	//  The returned list of [PackageVersionSummary] objects.
	//
	// [PackageVersionSummary]: https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageVersionSummary.html
	Versions []types.PackageVersionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPackageVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPackageVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPackageVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPackageVersions"); err != nil {
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
	if err = addOpListPackageVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPackageVersions(options.Region), middleware.Before); err != nil {
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

// ListPackageVersionsPaginatorOptions is the paginator options for
// ListPackageVersions
type ListPackageVersionsPaginatorOptions struct {
	//  The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPackageVersionsPaginator is a paginator for ListPackageVersions
type ListPackageVersionsPaginator struct {
	options   ListPackageVersionsPaginatorOptions
	client    ListPackageVersionsAPIClient
	params    *ListPackageVersionsInput
	nextToken *string
	firstPage bool
}

// NewListPackageVersionsPaginator returns a new ListPackageVersionsPaginator
func NewListPackageVersionsPaginator(client ListPackageVersionsAPIClient, params *ListPackageVersionsInput, optFns ...func(*ListPackageVersionsPaginatorOptions)) *ListPackageVersionsPaginator {
	if params == nil {
		params = &ListPackageVersionsInput{}
	}

	options := ListPackageVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPackageVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPackageVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPackageVersions page.
func (p *ListPackageVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPackageVersionsOutput, error) {
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
	result, err := p.client.ListPackageVersions(ctx, &params, optFns...)
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

// ListPackageVersionsAPIClient is a client that implements the
// ListPackageVersions operation.
type ListPackageVersionsAPIClient interface {
	ListPackageVersions(context.Context, *ListPackageVersionsInput, ...func(*Options)) (*ListPackageVersionsOutput, error)
}

var _ ListPackageVersionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListPackageVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPackageVersions",
	}
}
