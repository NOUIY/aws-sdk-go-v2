// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets an Amazon DataZone asset.
func (c *Client) GetAsset(ctx context.Context, params *GetAssetInput, optFns ...func(*Options)) (*GetAssetOutput, error) {
	if params == nil {
		params = &GetAssetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAsset", params, optFns, c.addOperationGetAssetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAssetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssetInput struct {

	// The ID of the Amazon DataZone domain to which the asset belongs.
	//
	// This member is required.
	DomainIdentifier *string

	// The ID of the Amazon DataZone asset.
	//
	// This parameter supports either the value of assetId or externalIdentifier as
	// input. If you are passing the value of externalIdentifier , you must prefix this
	// value with externalIdentifer%2F .
	//
	// This member is required.
	Identifier *string

	// The revision of the Amazon DataZone asset.
	Revision *string

	noSmithyDocumentSerde
}

type GetAssetOutput struct {

	// The ID of the Amazon DataZone domain to which the asset belongs.
	//
	// This member is required.
	DomainId *string

	// The metadata forms attached to the asset.
	//
	// This member is required.
	FormsOutput []types.FormOutput

	// The ID of the asset.
	//
	// This member is required.
	Id *string

	// The name of the asset.
	//
	// This member is required.
	Name *string

	// The ID of the project that owns the asset.
	//
	// This member is required.
	OwningProjectId *string

	// The revision of the asset.
	//
	// This member is required.
	Revision *string

	// The ID of the asset type.
	//
	// This member is required.
	TypeIdentifier *string

	// The revision of the asset type.
	//
	// This member is required.
	TypeRevision *string

	// The timestamp of when the asset was created.
	CreatedAt *time.Time

	// The Amazon DataZone user who created the asset.
	CreatedBy *string

	// The description of the Amazon DataZone asset.
	Description *string

	// The external ID of the asset.
	ExternalIdentifier *string

	// The timestamp of when the first revision of the asset was created.
	FirstRevisionCreatedAt *time.Time

	// The Amazon DataZone user who created the first revision of the asset.
	FirstRevisionCreatedBy *string

	// The business glossary terms attached to the asset.
	GlossaryTerms []string

	// The latest data point that was imported into the time series form for the
	// asset.
	LatestTimeSeriesDataPointFormsOutput []types.TimeSeriesDataPointSummaryFormOutput

	// The listing of the asset.
	Listing *types.AssetListingDetails

	// The read-only metadata forms attached to the asset.
	ReadOnlyFormsOutput []types.FormOutput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAssetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAsset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAsset{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAsset"); err != nil {
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
	if err = addOpGetAssetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAsset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAsset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAsset",
	}
}
