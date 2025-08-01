// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List all permissions on a private CA, if any, granted to the Certificate
// Manager (ACM) service principal (acm.amazonaws.com).
//
// These permissions allow ACM to issue and renew ACM certificates that reside in
// the same Amazon Web Services account as the CA.
//
// Permissions can be granted with the [CreatePermission] action and revoked with the [DeletePermission] action.
//
// About Permissions
//
//   - If the private CA and the certificates it issues reside in the same
//     account, you can use CreatePermission to grant permissions for ACM to carry
//     out automatic certificate renewals.
//
//   - For automatic certificate renewal to succeed, the ACM service principal
//     needs permissions to create, retrieve, and list certificates.
//
//   - If the private CA and the ACM certificates reside in different accounts,
//     then permissions cannot be used to enable automatic renewals. Instead, the ACM
//     certificate owner must set up a resource-based policy to enable cross-account
//     issuance and renewals. For more information, see [Using a Resource Based Policy with Amazon Web Services Private CA].
//
// [Using a Resource Based Policy with Amazon Web Services Private CA]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-rbp.html
// [CreatePermission]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreatePermission.html
// [DeletePermission]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_DeletePermission.html
func (c *Client) ListPermissions(ctx context.Context, params *ListPermissionsInput, optFns ...func(*Options)) (*ListPermissionsOutput, error) {
	if params == nil {
		params = &ListPermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPermissions", params, optFns, c.addOperationListPermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPermissionsInput struct {

	// The Amazon Resource Number (ARN) of the private CA to inspect. You can find the
	// ARN by calling the [ListCertificateAuthorities]action. This must be of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	// You can get a private CA's ARN by running the [ListCertificateAuthorities]action.
	//
	// [ListCertificateAuthorities]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListCertificateAuthorities.html
	//
	// This member is required.
	CertificateAuthorityArn *string

	// When paginating results, use this parameter to specify the maximum number of
	// items to return in the response. If additional items exist beyond the number you
	// specify, the NextToken element is sent in the response. Use this NextToken value
	// in a subsequent request to retrieve additional items.
	MaxResults *int32

	// When paginating results, use this parameter in a subsequent request after you
	// receive a response with truncated results. Set it to the value of NextToken from
	// the response you just received.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPermissionsOutput struct {

	// When the list is truncated, this value is present and should be used for the
	// NextToken parameter in a subsequent pagination request.
	NextToken *string

	// Summary information about each permission assigned by the specified private CA,
	// including the action enabled, the policy provided, and the time of creation.
	Permissions []types.Permission

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPermissions"); err != nil {
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
	if err = addOpListPermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPermissions(options.Region), middleware.Before); err != nil {
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

// ListPermissionsPaginatorOptions is the paginator options for ListPermissions
type ListPermissionsPaginatorOptions struct {
	// When paginating results, use this parameter to specify the maximum number of
	// items to return in the response. If additional items exist beyond the number you
	// specify, the NextToken element is sent in the response. Use this NextToken value
	// in a subsequent request to retrieve additional items.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPermissionsPaginator is a paginator for ListPermissions
type ListPermissionsPaginator struct {
	options   ListPermissionsPaginatorOptions
	client    ListPermissionsAPIClient
	params    *ListPermissionsInput
	nextToken *string
	firstPage bool
}

// NewListPermissionsPaginator returns a new ListPermissionsPaginator
func NewListPermissionsPaginator(client ListPermissionsAPIClient, params *ListPermissionsInput, optFns ...func(*ListPermissionsPaginatorOptions)) *ListPermissionsPaginator {
	if params == nil {
		params = &ListPermissionsInput{}
	}

	options := ListPermissionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPermissionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPermissionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPermissions page.
func (p *ListPermissionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPermissionsOutput, error) {
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
	result, err := p.client.ListPermissions(ctx, &params, optFns...)
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

// ListPermissionsAPIClient is a client that implements the ListPermissions
// operation.
type ListPermissionsAPIClient interface {
	ListPermissions(context.Context, *ListPermissionsInput, ...func(*Options)) (*ListPermissionsOutput, error)
}

var _ ListPermissionsAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListPermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPermissions",
	}
}
