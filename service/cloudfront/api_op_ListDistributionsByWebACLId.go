// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the distributions that are associated with a specified WAF web ACL.
func (c *Client) ListDistributionsByWebACLId(ctx context.Context, params *ListDistributionsByWebACLIdInput, optFns ...func(*Options)) (*ListDistributionsByWebACLIdOutput, error) {
	if params == nil {
		params = &ListDistributionsByWebACLIdInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDistributionsByWebACLId", params, optFns, c.addOperationListDistributionsByWebACLIdMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDistributionsByWebACLIdOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to list distributions that are associated with a specified WAF web
// ACL.
type ListDistributionsByWebACLIdInput struct {

	// The ID of the WAF web ACL that you want to list the associated distributions.
	// If you specify "null" for the ID, the request returns a list of the
	// distributions that aren't associated with a web ACL.
	//
	// For WAFV2, this is the ARN of the web ACL, such as
	// arn:aws:wafv2:us-east-1:123456789012:global/webacl/ExampleWebACL/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111
	// .
	//
	// For WAF Classic, this is the ID of the web ACL, such as
	// a1b2c3d4-5678-90ab-cdef-EXAMPLE11111 .
	//
	// This member is required.
	WebACLId *string

	// Use Marker and MaxItems to control pagination of results. If you have more than
	// MaxItems distributions that satisfy the request, the response includes a
	// NextMarker element. To get the next page of results, submit another request. For
	// the value of Marker , specify the value of NextMarker from the last response.
	// (For the first request, omit Marker .)
	Marker *string

	// The maximum number of distributions that you want CloudFront to return in the
	// response body. The maximum and default values are both 100.
	MaxItems *int32

	noSmithyDocumentSerde
}

// The response to a request to list the distributions that are associated with a
// specified WAF web ACL.
type ListDistributionsByWebACLIdOutput struct {

	// The DistributionList type.
	DistributionList *types.DistributionList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDistributionsByWebACLIdMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListDistributionsByWebACLId{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListDistributionsByWebACLId{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDistributionsByWebACLId"); err != nil {
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
	if err = addOpListDistributionsByWebACLIdValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDistributionsByWebACLId(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDistributionsByWebACLId(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDistributionsByWebACLId",
	}
}
