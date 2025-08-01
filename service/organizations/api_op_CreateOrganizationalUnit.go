// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an organizational unit (OU) within a root or parent OU. An OU is a
// container for accounts that enables you to organize your accounts to apply
// policies according to your business requirements. The number of levels deep that
// you can nest OUs is dependent upon the policy types enabled for that root. For
// service control policies, the limit is five.
//
// For more information about OUs, see [Managing organizational units (OUs)] in the Organizations User Guide.
//
// If the request includes tags, then the requester must have the
// organizations:TagResource permission.
//
// This operation can be called only from the organization's management account.
//
// [Managing organizational units (OUs)]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_ous.html
func (c *Client) CreateOrganizationalUnit(ctx context.Context, params *CreateOrganizationalUnitInput, optFns ...func(*Options)) (*CreateOrganizationalUnitOutput, error) {
	if params == nil {
		params = &CreateOrganizationalUnitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateOrganizationalUnit", params, optFns, c.addOperationCreateOrganizationalUnitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateOrganizationalUnitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateOrganizationalUnitInput struct {

	// The friendly name to assign to the new OU.
	//
	// This member is required.
	Name *string

	// The unique identifier (ID) of the parent root or OU that you want to create the
	// new OU in.
	//
	// The [regex pattern] for a parent ID string requires one of the following:
	//
	//   - Root - A string that begins with "r-" followed by from 4 to 32 lowercase
	//   letters or digits.
	//
	//   - Organizational unit (OU) - A string that begins with "ou-" followed by from
	//   4 to 32 lowercase letters or digits (the ID of the root that the OU is in). This
	//   string is followed by a second "-" dash and from 8 to 32 additional lowercase
	//   letters or digits.
	//
	// [regex pattern]: http://wikipedia.org/wiki/regex
	//
	// This member is required.
	ParentId *string

	// A list of tags that you want to attach to the newly created OU. For each tag in
	// the list, you must specify both a tag key and a value. You can set the value to
	// an empty string, but you can't set it to null . For more information about
	// tagging, see [Tagging Organizations resources]in the Organizations User Guide.
	//
	// If any one of the tags is not valid or if you exceed the allowed number of tags
	// for an OU, then the entire request fails and the OU is not created.
	//
	// [Tagging Organizations resources]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_tagging.html
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateOrganizationalUnitOutput struct {

	// A structure that contains details about the newly created OU.
	OrganizationalUnit *types.OrganizationalUnit

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateOrganizationalUnitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateOrganizationalUnit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateOrganizationalUnit{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateOrganizationalUnit"); err != nil {
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
	if err = addOpCreateOrganizationalUnitValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOrganizationalUnit(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateOrganizationalUnit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateOrganizationalUnit",
	}
}
