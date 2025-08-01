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

// Creates an access point and associates it to a specified bucket. For more
// information, see [Managing access to shared datasets with access points]or [Managing access to shared datasets in directory buckets with access points] in the Amazon S3 User Guide.
//
// To create an access point and attach it to a volume on an Amazon FSx file
// system, see [CreateAndAttachS3AccessPoint]in the Amazon FSx API Reference.
//
// S3 on Outposts only supports VPC-style access points.
//
// For more information, see [Accessing Amazon S3 on Outposts using virtual private cloud (VPC) only access points] in the Amazon S3 User Guide.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to CreateAccessPoint :
//
// [GetAccessPoint]
//
// [DeleteAccessPoint]
//
// [ListAccessPoints]
//
// [ListAccessPointsForDirectoryBuckets]
//
// [ListAccessPointsForDirectoryBuckets]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPointsForDirectoryBuckets.html
// [ListAccessPoints]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPoints.html
// [CreateAndAttachS3AccessPoint]: https://docs.aws.amazon.com/fsx/latest/APIReference/API_CreateAndAttachS3AccessPoint.html
// [Managing access to shared datasets with access points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points.html
// [GetAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html
// [Accessing Amazon S3 on Outposts using virtual private cloud (VPC) only access points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [DeleteAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPoint.html
// [Managing access to shared datasets in directory buckets with access points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-directory-buckets.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html#API_control_CreateAccessPoint_Examples
func (c *Client) CreateAccessPoint(ctx context.Context, params *CreateAccessPointInput, optFns ...func(*Options)) (*CreateAccessPointOutput, error) {
	if params == nil {
		params = &CreateAccessPointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAccessPoint", params, optFns, c.addOperationCreateAccessPointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAccessPointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccessPointInput struct {

	// The Amazon Web Services account ID for the account that owns the specified
	// access point.
	//
	// This member is required.
	AccountId *string

	// The name of the bucket that you want to associate this access point with.
	//
	// For using this parameter with Amazon S3 on Outposts with the REST API, you must
	// specify the name and the x-amz-outpost-id as well.
	//
	// For using this parameter with S3 on Outposts with the Amazon Web Services SDK
	// and CLI, you must specify the ARN of the bucket accessed in the format
	// arn:aws:s3-outposts:::outpost//bucket/ . For example, to access the bucket
	// reports through Outpost my-outpost owned by account 123456789012 in Region
	// us-west-2 , use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports .
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string

	// The name you want to assign to this access point.
	//
	// For directory buckets, the access point name must consist of a base name that
	// you provide and suffix that includes the ZoneID (Amazon Web Services
	// Availability Zone or Local Zone) of your bucket location, followed by --xa-s3 .
	// For more information, see [Managing access to shared datasets in directory buckets with access points]in the Amazon S3 User Guide.
	//
	// [Managing access to shared datasets in directory buckets with access points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-directory-buckets.html
	//
	// This member is required.
	Name *string

	// The Amazon Web Services account ID associated with the S3 bucket associated
	// with this access point.
	//
	// For same account access point when your bucket and access point belong to the
	// same account owner, the BucketAccountId is not required. For cross-account
	// access point when your bucket and access point are not in the same account, the
	// BucketAccountId is required.
	BucketAccountId *string

	//  The PublicAccessBlock configuration that you want to apply to the access
	// point.
	PublicAccessBlockConfiguration *types.PublicAccessBlockConfiguration

	// For directory buckets, you can filter access control to specific prefixes, API
	// operations, or a combination of both. For more information, see [Managing access to shared datasets in directory buckets with access points]in the Amazon
	// S3 User Guide.
	//
	// Scope is only supported for access points attached to directory buckets.
	//
	// [Managing access to shared datasets in directory buckets with access points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-directory-buckets.html
	Scope *types.Scope

	// An array of tags that you can apply to an access point. Tags are key-value
	// pairs of metadata used to control access to your access points. For more
	// information about tags, see [Using tags with Amazon S3]. For information about tagging access points, see [Using tags for attribute-based access control (ABAC)].
	//
	// [Using tags with Amazon S3]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/tagging.html
	// [Using tags for attribute-based access control (ABAC)]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/tagging.html#using-tags-for-abac
	Tags []types.Tag

	// If you include this field, Amazon S3 restricts access to this access point to
	// requests from the specified virtual private cloud (VPC).
	//
	// This is required for creating an access point for Amazon S3 on Outposts buckets.
	VpcConfiguration *types.VpcConfiguration

	noSmithyDocumentSerde
}

func (in *CreateAccessPointInput) bindEndpointParams(p *EndpointParameters) {

	p.AccountId = in.AccountId
	p.AccessPointName = in.Name
	p.Bucket = in.Bucket
	p.RequiresAccountId = ptr.Bool(true)
}

type CreateAccessPointOutput struct {

	// The ARN of the access point.
	//
	// This is only supported by Amazon S3 on Outposts.
	AccessPointArn *string

	// The name or alias of the access point.
	Alias *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAccessPointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateAccessPoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateAccessPoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAccessPoint"); err != nil {
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
	if err = addEndpointPrefix_opCreateAccessPointMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateAccessPointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccessPoint(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addCreateAccessPointUpdateEndpoint(stack, options); err != nil {
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

func (m *CreateAccessPointInput) GetARNMember() (*string, bool) {
	if m.Bucket == nil {
		return nil, false
	}
	return m.Bucket, true
}

func (m *CreateAccessPointInput) SetARNMember(v string) error {
	m.Bucket = &v
	return nil
}

type endpointPrefix_opCreateAccessPointMiddleware struct {
}

func (*endpointPrefix_opCreateAccessPointMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateAccessPointMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
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
	input, ok := opaqueInput.(*CreateAccessPointInput)
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
func addEndpointPrefix_opCreateAccessPointMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opCreateAccessPointMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opCreateAccessPoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAccessPoint",
	}
}

func copyCreateAccessPointInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*CreateAccessPointInput)
	if !ok {
		return nil, fmt.Errorf("expect *CreateAccessPointInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *CreateAccessPointInput) copy() interface{} {
	v := *in
	return &v
}
func getCreateAccessPointARNMember(input interface{}) (*string, bool) {
	in := input.(*CreateAccessPointInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func setCreateAccessPointARNMember(input interface{}, v string) error {
	in := input.(*CreateAccessPointInput)
	in.Bucket = &v
	return nil
}
func backFillCreateAccessPointAccountID(input interface{}, v string) error {
	in := input.(*CreateAccessPointInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addCreateAccessPointUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getCreateAccessPointARNMember,
			BackfillAccountID: backFillCreateAccessPointAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setCreateAccessPointARNMember,
			CopyInput:         copyCreateAccessPointInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
