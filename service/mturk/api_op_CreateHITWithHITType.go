// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

//	The CreateHITWithHITType operation creates a new Human Intelligence Task (HIT)
//
// using an existing HITTypeID generated by the CreateHITType operation.
//
// This is an alternative way to create HITs from the CreateHIT operation. This is
// the recommended best practice for Requesters who are creating large numbers of
// HITs.
//
// CreateHITWithHITType also supports several ways to provide question data: by
// providing a value for the Question parameter that fully specifies the contents
// of the HIT, or by providing a HitLayoutId and associated HitLayoutParameters .
//
// If a HIT is created with 10 or more maximum assignments, there is an additional
// fee. For more information, see [Amazon Mechanical Turk Pricing].
//
// [Amazon Mechanical Turk Pricing]: https://requester.mturk.com/pricing
func (c *Client) CreateHITWithHITType(ctx context.Context, params *CreateHITWithHITTypeInput, optFns ...func(*Options)) (*CreateHITWithHITTypeOutput, error) {
	if params == nil {
		params = &CreateHITWithHITTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHITWithHITType", params, optFns, c.addOperationCreateHITWithHITTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHITWithHITTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHITWithHITTypeInput struct {

	// The HIT type ID you want to create this HIT with.
	//
	// This member is required.
	HITTypeId *string

	//  An amount of time, in seconds, after which the HIT is no longer available for
	// users to accept. After the lifetime of the HIT elapses, the HIT no longer
	// appears in HIT searches, even if not all of the assignments for the HIT have
	// been accepted.
	//
	// This member is required.
	LifetimeInSeconds *int64

	//  The Assignment-level Review Policy applies to the assignments under the HIT.
	// You can specify for Mechanical Turk to take various actions based on the policy.
	AssignmentReviewPolicy *types.ReviewPolicy

	//  The HITLayoutId allows you to use a pre-existing HIT design with placeholder
	// values and create an additional HIT by providing those values as
	// HITLayoutParameters.
	//
	// Constraints: Either a Question parameter or a HITLayoutId parameter must be
	// provided.
	HITLayoutId *string

	//  If the HITLayoutId is provided, any placeholder values must be filled in with
	// values using the HITLayoutParameter structure. For more information, see
	// HITLayout.
	HITLayoutParameters []types.HITLayoutParameter

	//  The HIT-level Review Policy applies to the HIT. You can specify for Mechanical
	// Turk to take various actions based on the policy.
	HITReviewPolicy *types.ReviewPolicy

	//  The number of times the HIT can be accepted and completed before the HIT
	// becomes unavailable.
	MaxAssignments *int32

	//  The data the person completing the HIT uses to produce the results.
	//
	// Constraints: Must be a QuestionForm data structure, an ExternalQuestion data
	// structure, or an HTMLQuestion data structure. The XML question data must not be
	// larger than 64 kilobytes (65,535 bytes) in size, including whitespace.
	//
	// Either a Question parameter or a HITLayoutId parameter must be provided.
	Question *string

	//  An arbitrary data field. The RequesterAnnotation parameter lets your
	// application attach arbitrary data to the HIT for tracking purposes. For example,
	// this parameter could be an identifier internal to the Requester's application
	// that corresponds with the HIT.
	//
	// The RequesterAnnotation parameter for a HIT is only visible to the Requester
	// who created the HIT. It is not shown to the Worker, or any other Requester.
	//
	// The RequesterAnnotation parameter may be different for each HIT you submit. It
	// does not affect how your HITs are grouped.
	RequesterAnnotation *string

	//  A unique identifier for this request which allows you to retry the call on
	// error without creating duplicate HITs. This is useful in cases such as network
	// timeouts where it is unclear whether or not the call succeeded on the server. If
	// the HIT already exists in the system from a previous call using the same
	// UniqueRequestToken, subsequent calls will return a
	// AWS.MechanicalTurk.HitAlreadyExists error with a message containing the HITId.
	//
	// Note: It is your responsibility to ensure uniqueness of the token. The unique
	// token expires after 24 hours. Subsequent calls using the same UniqueRequestToken
	// made after the 24 hour limit could create duplicate HITs.
	UniqueRequestToken *string

	noSmithyDocumentSerde
}

type CreateHITWithHITTypeOutput struct {

	//  Contains the newly created HIT data. For a description of the HIT data
	// structure as it appears in responses, see the HIT Data Structure documentation.
	HIT *types.HIT

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateHITWithHITTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateHITWithHITType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateHITWithHITType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateHITWithHITType"); err != nil {
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
	if err = addOpCreateHITWithHITTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHITWithHITType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateHITWithHITType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateHITWithHITType",
	}
}
