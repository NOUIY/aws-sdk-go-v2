// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Given a user pool ID, returns a list of users and their basic details in a user
// pool.
//
// Amazon Cognito evaluates Identity and Access Management (IAM) policies in
// requests for this API operation. For this operation, you must use IAM
// credentials to authorize requests, and you must grant yourself the corresponding
// IAM permission in a policy.
//
// # Learn more
//
// [Signing Amazon Web Services API Requests]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]
//
// [Using the Amazon Cognito user pools API and user pool endpoints]: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html
// [Signing Amazon Web Services API Requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html
func (c *Client) ListUsers(ctx context.Context, params *ListUsersInput, optFns ...func(*Options)) (*ListUsersOutput, error) {
	if params == nil {
		params = &ListUsersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUsers", params, optFns, c.addOperationListUsersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUsersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to list users.
type ListUsersInput struct {

	// The ID of the user pool where you want to display or search for users.
	//
	// This member is required.
	UserPoolId *string

	// A JSON array of user attribute names, for example given_name , that you want
	// Amazon Cognito to include in the response for each user. When you don't provide
	// an AttributesToGet parameter, Amazon Cognito returns all attributes for each
	// user.
	//
	// Use AttributesToGet with required attributes in your user pool, or in
	// conjunction with Filter . Amazon Cognito returns an error if not all users in
	// the results have set a value for the attribute you request. Attributes that you
	// can't filter on, including custom attributes, must have a value set in every
	// user profile before an AttributesToGet parameter returns results.
	AttributesToGet []string

	// A filter string of the form "AttributeName Filter-Type "AttributeValue" .
	// Quotation marks within the filter string must be escaped using the backslash ( \
	// ) character. For example, "family_name = \"Reddy\"" .
	//
	//   - AttributeName: The name of the attribute to search for. You can only search
	//   for one attribute at a time.
	//
	//   - Filter-Type: For an exact match, use = , for example, " given_name = \"Jon\"
	//   ". For a prefix ("starts with") match, use ^= , for example, " given_name ^=
	//   \"Jon\" ".
	//
	//   - AttributeValue: The attribute value that must be matched for each user.
	//
	// If the filter string is empty, ListUsers returns all users in the user pool.
	//
	// You can only search for the following standard attributes:
	//
	//   - username (case-sensitive)
	//
	//   - email
	//
	//   - phone_number
	//
	//   - name
	//
	//   - given_name
	//
	//   - family_name
	//
	//   - preferred_username
	//
	//   - cognito:user_status (called Status in the Console) (case-insensitive)
	//
	//   - status (called Enabled in the Console) (case-sensitive)
	//
	//   - sub
	//
	// Custom attributes aren't searchable.
	//
	// You can also list users with a client-side filter. The server-side filter
	// matches no more than one attribute. For an advanced search, use a client-side
	// filter with the --query parameter of the list-users action in the CLI. When you
	// use a client-side filter, ListUsers returns a paginated list of zero or more
	// users. You can receive multiple pages in a row with zero results. Repeat the
	// query with each pagination token that is returned until you receive a null
	// pagination token value, and then review the combined result.
	//
	// For more information about server-side and client-side filtering, see [FilteringCLI output] in the [Command Line Interface User Guide].
	//
	// For more information, see [Searching for Users Using the ListUsers API] and [Examples of Using the ListUsers API] in the Amazon Cognito Developer Guide.
	//
	// [Command Line Interface User Guide]: https://docs.aws.amazon.com/cli/latest/userguide/cli-usage-filter.html
	// [Searching for Users Using the ListUsers API]: https://docs.aws.amazon.com/cognito/latest/developerguide/how-to-manage-user-accounts.html#cognito-user-pools-searching-for-users-using-listusers-api
	// [FilteringCLI output]: https://docs.aws.amazon.com/cli/latest/userguide/cli-usage-filter.html
	// [Examples of Using the ListUsers API]: https://docs.aws.amazon.com/cognito/latest/developerguide/how-to-manage-user-accounts.html#cognito-user-pools-searching-for-users-listusers-api-examples
	Filter *string

	// The maximum number of users that you want Amazon Cognito to return in the
	// response.
	Limit *int32

	// This API operation returns a limited number of results. The pagination token is
	// an identifier that you can present in an additional API request with the same
	// parameters. When you include the pagination token, Amazon Cognito returns the
	// next set of items after the current list. Subsequent requests return a new
	// pagination token. By use of this token, you can paginate through the full list
	// of items.
	PaginationToken *string

	noSmithyDocumentSerde
}

// The response from the request to list users.
type ListUsersOutput struct {

	// The identifier that Amazon Cognito returned with the previous request to this
	// operation. When you include a pagination token in your request, Amazon Cognito
	// returns the next set of items in the list. By use of this token, you can
	// paginate through the full list of items.
	PaginationToken *string

	// An array of user pool users who match your query, and their attributes.
	Users []types.UserType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListUsersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListUsers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListUsers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListUsers"); err != nil {
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
	if err = addOpListUsersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUsers(options.Region), middleware.Before); err != nil {
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

// ListUsersPaginatorOptions is the paginator options for ListUsers
type ListUsersPaginatorOptions struct {
	// The maximum number of users that you want Amazon Cognito to return in the
	// response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUsersPaginator is a paginator for ListUsers
type ListUsersPaginator struct {
	options   ListUsersPaginatorOptions
	client    ListUsersAPIClient
	params    *ListUsersInput
	nextToken *string
	firstPage bool
}

// NewListUsersPaginator returns a new ListUsersPaginator
func NewListUsersPaginator(client ListUsersAPIClient, params *ListUsersInput, optFns ...func(*ListUsersPaginatorOptions)) *ListUsersPaginator {
	if params == nil {
		params = &ListUsersInput{}
	}

	options := ListUsersPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListUsersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.PaginationToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUsersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListUsers page.
func (p *ListUsersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUsersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PaginationToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListUsers(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.PaginationToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListUsersAPIClient is a client that implements the ListUsers operation.
type ListUsersAPIClient interface {
	ListUsers(context.Context, *ListUsersInput, ...func(*Options)) (*ListUsersOutput, error)
}

var _ ListUsersAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListUsers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListUsers",
	}
}
