// Code generated by smithy-go-codegen DO NOT EDIT.

package securityir

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/securityir/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Grants permission to get details of a designated service membership.
func (c *Client) GetMembership(ctx context.Context, params *GetMembershipInput, optFns ...func(*Options)) (*GetMembershipOutput, error) {
	if params == nil {
		params = &GetMembershipInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMembership", params, optFns, c.addOperationGetMembershipMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMembershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMembershipInput struct {

	// Required element for GetMembership to identify the membership ID to query.
	//
	// This member is required.
	MembershipId *string

	noSmithyDocumentSerde
}

type GetMembershipOutput struct {

	// Response element for GetMembership that provides the queried membership ID.
	//
	// This member is required.
	MembershipId *string

	// Response element for GetMembership that provides the configured account for
	// managing the membership.
	AccountId *string

	// Response element for GetMembership that provides the configured membership
	// type. Options include Standalone | Organizations .
	CustomerType types.CustomerType

	// Response element for GetMembership that provides the configured membership
	// incident response team members.
	IncidentResponseTeam []types.IncidentResponder

	// Response element for GetMembership that provides the configured membership
	// activation timestamp.
	MembershipActivationTimestamp *time.Time

	// Response element for GetMembership that provides the membership ARN.
	MembershipArn *string

	// Response element for GetMembership that provides the configured membership name
	// deactivation timestamp.
	MembershipDeactivationTimestamp *time.Time

	// Response element for GetMembership that provides the configured membership name.
	MembershipName *string

	// Response element for GetMembership that provides the current membership status.
	MembershipStatus types.MembershipStatus

	// Response element for GetMembership that provides the number of accounts in the
	// membership.
	NumberOfAccountsCovered *int64

	// Response element for GetMembership that provides the if opt-in features have
	// been enabled.
	OptInFeatures []types.OptInFeature

	// Response element for GetMembership that provides the configured region for
	// managing the membership.
	Region types.AwsRegion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMembershipMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMembership{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMembership{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMembership"); err != nil {
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
	if err = addOpGetMembershipValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMembership(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMembership(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMembership",
	}
}
