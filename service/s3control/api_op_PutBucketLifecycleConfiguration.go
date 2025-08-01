// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// This action puts a lifecycle configuration to an Amazon S3 on Outposts bucket.
// To put a lifecycle configuration to an S3 bucket, see [PutBucketLifecycleConfiguration]in the Amazon S3 API
// Reference.
//
// Creates a new lifecycle configuration for the S3 on Outposts bucket or replaces
// an existing lifecycle configuration. Outposts buckets only support lifecycle
// configurations that delete/expire objects after a certain period of time and
// abort incomplete multipart uploads.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to PutBucketLifecycleConfiguration :
//
// [GetBucketLifecycleConfiguration]
//
// [DeleteBucketLifecycleConfiguration]
//
// [PutBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycleConfiguration.html
// [GetBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketLifecycleConfiguration.html
// [DeleteBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketLifecycleConfiguration.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketLifecycleConfiguration.html#API_control_PutBucketLifecycleConfiguration_Examples
func (c *Client) PutBucketLifecycleConfiguration(ctx context.Context, params *PutBucketLifecycleConfigurationInput, optFns ...func(*Options)) (*PutBucketLifecycleConfigurationOutput, error) {
	if params == nil {
		params = &PutBucketLifecycleConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketLifecycleConfiguration", params, optFns, c.addOperationPutBucketLifecycleConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketLifecycleConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketLifecycleConfigurationInput struct {

	// The Amazon Web Services account ID of the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// The name of the bucket for which to set the configuration.
	//
	// This member is required.
	Bucket *string

	// Container for lifecycle rules. You can add as many as 1,000 rules.
	LifecycleConfiguration *types.LifecycleConfiguration

	noSmithyDocumentSerde
}

func (in *PutBucketLifecycleConfigurationInput) bindEndpointParams(p *EndpointParameters) {

	p.AccountId = in.AccountId
	p.Bucket = in.Bucket
	p.RequiresAccountId = ptr.Bool(true)
}

type PutBucketLifecycleConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBucketLifecycleConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketLifecycleConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketLifecycleConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutBucketLifecycleConfiguration"); err != nil {
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
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
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opPutBucketLifecycleConfigurationMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutBucketLifecycleConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketLifecycleConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addPutBucketLifecycleConfigurationUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addStashOperationInput(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3controlcust.AddDisableHostPrefixMiddleware(stack); err != nil {
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

func (m *PutBucketLifecycleConfigurationInput) GetARNMember() (*string, bool) {
	if m.Bucket == nil {
		return nil, false
	}
	return m.Bucket, true
}

func (m *PutBucketLifecycleConfigurationInput) SetARNMember(v string) error {
	m.Bucket = &v
	return nil
}

type endpointPrefix_opPutBucketLifecycleConfigurationMiddleware struct {
}

func (*endpointPrefix_opPutBucketLifecycleConfigurationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opPutBucketLifecycleConfigurationMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	opaqueInput := getOperationInput(ctx)
	input, ok := opaqueInput.(*PutBucketLifecycleConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", opaqueInput)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opPutBucketLifecycleConfigurationMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opPutBucketLifecycleConfigurationMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opPutBucketLifecycleConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutBucketLifecycleConfiguration",
	}
}

func copyPutBucketLifecycleConfigurationInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*PutBucketLifecycleConfigurationInput)
	if !ok {
		return nil, fmt.Errorf("expect *PutBucketLifecycleConfigurationInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *PutBucketLifecycleConfigurationInput) copy() interface{} {
	v := *in
	return &v
}
func getPutBucketLifecycleConfigurationARNMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketLifecycleConfigurationInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func setPutBucketLifecycleConfigurationARNMember(input interface{}, v string) error {
	in := input.(*PutBucketLifecycleConfigurationInput)
	in.Bucket = &v
	return nil
}
func backFillPutBucketLifecycleConfigurationAccountID(input interface{}, v string) error {
	in := input.(*PutBucketLifecycleConfigurationInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addPutBucketLifecycleConfigurationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getPutBucketLifecycleConfigurationARNMember,
			BackfillAccountID: backFillPutBucketLifecycleConfigurationAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setPutBucketLifecycleConfigurationARNMember,
			CopyInput:         copyPutBucketLifecycleConfigurationInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
