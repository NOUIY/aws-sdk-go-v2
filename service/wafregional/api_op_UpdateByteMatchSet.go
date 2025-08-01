// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes ByteMatchTuple objects (filters) in a ByteMatchSet. For each ByteMatchTuple object,
// you specify the following values:
//
//   - Whether to insert or delete the object from the array. If you want to
//     change a ByteMatchSetUpdate object, you delete the existing object and add a
//     new one.
//
//   - The part of a web request that you want AWS WAF to inspect, such as a query
//     string or the value of the User-Agent header.
//
//   - The bytes (typically a string that corresponds with ASCII characters) that
//     you want AWS WAF to look for. For more information, including how you specify
//     the values for the AWS WAF API and the AWS CLI or SDKs, see TargetString in
//     the ByteMatchTupledata type.
//
//   - Where to look, such as at the beginning or the end of a query string.
//
//   - Whether to perform any conversions on the request, such as converting it to
//     lowercase, before inspecting it for the specified string.
//
// For example, you can add a ByteMatchSetUpdate object that matches web requests
// in which User-Agent headers contain the string BadBot . You can then configure
// AWS WAF to block those requests.
//
// To create and configure a ByteMatchSet , perform the following steps:
//
//   - Create a ByteMatchSet. For more information, see CreateByteMatchSet.
//
//   - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
//     an UpdateByteMatchSet request.
//
//   - Submit an UpdateByteMatchSet request to specify the part of the request that
//     you want AWS WAF to inspect (for example, the header or the URI) and the value
//     that you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
func (c *Client) UpdateByteMatchSet(ctx context.Context, params *UpdateByteMatchSetInput, optFns ...func(*Options)) (*UpdateByteMatchSetOutput, error) {
	if params == nil {
		params = &UpdateByteMatchSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateByteMatchSet", params, optFns, c.addOperationUpdateByteMatchSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateByteMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateByteMatchSetInput struct {

	// The ByteMatchSetId of the ByteMatchSet that you want to update. ByteMatchSetId is returned
	// by CreateByteMatchSetand by ListByteMatchSets.
	//
	// This member is required.
	ByteMatchSetId *string

	// The value returned by the most recent call to GetChangeToken.
	//
	// This member is required.
	ChangeToken *string

	// An array of ByteMatchSetUpdate objects that you want to insert into or delete
	// from a ByteMatchSet. For more information, see the applicable data types:
	//
	// ByteMatchSetUpdate
	//   - : Contains Action and ByteMatchTuple
	//
	// ByteMatchTuple
	//   - : Contains FieldToMatch , PositionalConstraint , TargetString , and
	//   TextTransformation
	//
	// FieldToMatch
	//   - : Contains Data and Type
	//
	// This member is required.
	Updates []types.ByteMatchSetUpdate

	noSmithyDocumentSerde
}

type UpdateByteMatchSetOutput struct {

	// The ChangeToken that you used to submit the UpdateByteMatchSet request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateByteMatchSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateByteMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateByteMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateByteMatchSet"); err != nil {
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
	if err = addOpUpdateByteMatchSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateByteMatchSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateByteMatchSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateByteMatchSet",
	}
}
