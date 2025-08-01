// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservicedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/directoryservicedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about a specific user.
func (c *Client) DescribeUser(ctx context.Context, params *DescribeUserInput, optFns ...func(*Options)) (*DescribeUserOutput, error) {
	if params == nil {
		params = &DescribeUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeUser", params, optFns, c.addOperationDescribeUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUserInput struct {

	//  The identifier (ID) of the directory that's associated with the user.
	//
	// This member is required.
	DirectoryId *string

	//  The name of the user.
	//
	// This member is required.
	SAMAccountName *string

	//  One or more attribute names to be returned for the user. A key is an attribute
	// name, and the value is a list of maps. For a list of supported attributes, see [Directory Service Data Attributes]
	// .
	//
	// [Directory Service Data Attributes]: https://docs.aws.amazon.com/directoryservice/latest/admin-guide/ad_data_attributes.html
	OtherAttributes []string

	//  The domain name that's associated with the user.
	//
	// This parameter is optional, so you can return users outside your Managed
	// Microsoft AD domain. When no value is defined, only your Managed Microsoft AD
	// users are returned.
	//
	// This value is case insensitive.
	Realm *string

	noSmithyDocumentSerde
}

type DescribeUserOutput struct {

	//  The identifier (ID) of the directory that's associated with the user.
	DirectoryId *string

	//  The [distinguished name] of the object.
	//
	// [distinguished name]: https://learn.microsoft.com/en-us/windows/win32/ad/object-names-and-identities#distinguished-name
	DistinguishedName *string

	//  The email address of the user.
	EmailAddress *string

	//  Indicates whether the user account is active.
	Enabled *bool

	//  The first name of the user.
	GivenName *string

	//  The attribute values that are returned for the attribute names that are
	// included in the request.
	//
	// Attribute names are case insensitive.
	OtherAttributes map[string]types.AttributeValue

	//  The domain name that's associated with the user.
	Realm *string

	//  The name of the user.
	SAMAccountName *string

	//  The unique security identifier (SID) of the user.
	SID *string

	//  The last name of the user.
	Surname *string

	//  The UPN that is an Internet-style login name for a user and is based on the
	// Internet standard [RFC 822]. The UPN is shorter than the distinguished name and easier
	// to remember.
	//
	// [RFC 822]: https://datatracker.ietf.org/doc/html/rfc822
	UserPrincipalName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeUser{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeUser"); err != nil {
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
	if err = addOpDescribeUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeUser",
	}
}
