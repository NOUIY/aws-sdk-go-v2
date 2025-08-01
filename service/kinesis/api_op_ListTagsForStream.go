// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the tags for the specified Kinesis data stream. This operation has a
// limit of five transactions per second per account.
//
// When invoking this API, you must use either the StreamARN or the StreamName
// parameter, or both. It is recommended that you use the StreamARN input
// parameter when you invoke this API.
func (c *Client) ListTagsForStream(ctx context.Context, params *ListTagsForStreamInput, optFns ...func(*Options)) (*ListTagsForStreamOutput, error) {
	if params == nil {
		params = &ListTagsForStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTagsForStream", params, optFns, c.addOperationListTagsForStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagsForStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for ListTagsForStream .
type ListTagsForStreamInput struct {

	// The key to use as the starting point for the list of tags. If this parameter is
	// set, ListTagsForStream gets all tags that occur after ExclusiveStartTagKey .
	ExclusiveStartTagKey *string

	// The number of tags to return. If this number is less than the total number of
	// tags associated with the stream, HasMoreTags is set to true . To list additional
	// tags, set ExclusiveStartTagKey to the last key in the response.
	Limit *int32

	// The ARN of the stream.
	StreamARN *string

	// The name of the stream.
	StreamName *string

	noSmithyDocumentSerde
}

func (in *ListTagsForStreamInput) bindEndpointParams(p *EndpointParameters) {

	p.StreamARN = in.StreamARN
	p.OperationType = ptr.String("control")
}

// Represents the output for ListTagsForStream .
type ListTagsForStreamOutput struct {

	// If set to true , more tags are available. To request additional tags, set
	// ExclusiveStartTagKey to the key of the last tag returned.
	//
	// This member is required.
	HasMoreTags *bool

	// A list of tags associated with StreamName , starting with the first tag after
	// ExclusiveStartTagKey and up to the specified Limit .
	//
	// This member is required.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTagsForStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTagsForStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTagsForStream{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTagsForStream"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTagsForStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListTagsForStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTagsForStream",
	}
}
