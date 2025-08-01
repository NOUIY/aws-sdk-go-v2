// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the latest version for every traffic policy that is
// associated with the current Amazon Web Services account. Policies are listed in
// the order that they were created in.
//
// For information about how of deleting a traffic policy affects the response
// from ListTrafficPolicies , see [DeleteTrafficPolicy].
//
// [DeleteTrafficPolicy]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_DeleteTrafficPolicy.html
func (c *Client) ListTrafficPolicies(ctx context.Context, params *ListTrafficPoliciesInput, optFns ...func(*Options)) (*ListTrafficPoliciesOutput, error) {
	if params == nil {
		params = &ListTrafficPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTrafficPolicies", params, optFns, c.addOperationListTrafficPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTrafficPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains the information about the request to list the
// traffic policies that are associated with the current Amazon Web Services
// account.
type ListTrafficPoliciesInput struct {

	// (Optional) The maximum number of traffic policies that you want Amazon Route 53
	// to return in response to this request. If you have more than MaxItems traffic
	// policies, the value of IsTruncated in the response is true , and the value of
	// TrafficPolicyIdMarker is the ID of the first traffic policy that Route 53 will
	// return if you submit another request.
	MaxItems *int32

	// (Conditional) For your first request to ListTrafficPolicies , don't include the
	// TrafficPolicyIdMarker parameter.
	//
	// If you have more traffic policies than the value of MaxItems ,
	// ListTrafficPolicies returns only the first MaxItems traffic policies. To get
	// the next group of policies, submit another request to ListTrafficPolicies . For
	// the value of TrafficPolicyIdMarker , specify the value of TrafficPolicyIdMarker
	// that was returned in the previous response.
	TrafficPolicyIdMarker *string

	noSmithyDocumentSerde
}

// A complex type that contains the response information for the request.
type ListTrafficPoliciesOutput struct {

	// A flag that indicates whether there are more traffic policies to be listed. If
	// the response was truncated, you can get the next group of traffic policies by
	// submitting another ListTrafficPolicies request and specifying the value of
	// TrafficPolicyIdMarker in the TrafficPolicyIdMarker request parameter.
	//
	// This member is required.
	IsTruncated bool

	// The value that you specified for the MaxItems parameter in the
	// ListTrafficPolicies request that produced the current response.
	//
	// This member is required.
	MaxItems *int32

	// If the value of IsTruncated is true , TrafficPolicyIdMarker is the ID of the
	// first traffic policy in the next group of MaxItems traffic policies.
	//
	// This member is required.
	TrafficPolicyIdMarker *string

	// A list that contains one TrafficPolicySummary element for each traffic policy
	// that was created by the current Amazon Web Services account.
	//
	// This member is required.
	TrafficPolicySummaries []types.TrafficPolicySummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTrafficPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListTrafficPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListTrafficPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTrafficPolicies"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTrafficPolicies(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListTrafficPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTrafficPolicies",
	}
}
