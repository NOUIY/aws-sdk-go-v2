// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a set of featured results. Features results are placed above all other
// results for certain queries. You map specific queries to specific documents for
// featuring in the results. If a query contains an exact match of a query, then
// one or more specific documents are featured in the search results.
func (c *Client) UpdateFeaturedResultsSet(ctx context.Context, params *UpdateFeaturedResultsSetInput, optFns ...func(*Options)) (*UpdateFeaturedResultsSetOutput, error) {
	if params == nil {
		params = &UpdateFeaturedResultsSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFeaturedResultsSet", params, optFns, c.addOperationUpdateFeaturedResultsSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFeaturedResultsSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFeaturedResultsSetInput struct {

	// The identifier of the set of featured results that you want to update.
	//
	// This member is required.
	FeaturedResultsSetId *string

	// The identifier of the index used for featuring results.
	//
	// This member is required.
	IndexId *string

	// A new description for the set of featured results.
	Description *string

	// A list of document IDs for the documents you want to feature at the top of the
	// search results page. For more information on the list of featured documents, see
	// [FeaturedResultsSet].
	//
	// [FeaturedResultsSet]: https://docs.aws.amazon.com/kendra/latest/dg/API_FeaturedResultsSet.html
	FeaturedDocuments []types.FeaturedDocument

	// A new name for the set of featured results.
	FeaturedResultsSetName *string

	// A list of queries for featuring results. For more information on the list of
	// queries, see [FeaturedResultsSet].
	//
	// [FeaturedResultsSet]: https://docs.aws.amazon.com/kendra/latest/dg/API_FeaturedResultsSet.html
	QueryTexts []string

	// You can set the status to ACTIVE or INACTIVE . When the value is ACTIVE ,
	// featured results are ready for use. You can still configure your settings before
	// setting the status to ACTIVE . The queries you specify for featured results must
	// be unique per featured results set for each index, whether the status is ACTIVE
	// or INACTIVE .
	Status types.FeaturedResultsSetStatus

	noSmithyDocumentSerde
}

type UpdateFeaturedResultsSetOutput struct {

	// Information on the set of featured results. This includes the identifier of the
	// featured results set, whether the featured results set is active or inactive,
	// when the featured results set was last updated, and more.
	FeaturedResultsSet *types.FeaturedResultsSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFeaturedResultsSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateFeaturedResultsSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateFeaturedResultsSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateFeaturedResultsSet"); err != nil {
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
	if err = addOpUpdateFeaturedResultsSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFeaturedResultsSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFeaturedResultsSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateFeaturedResultsSet",
	}
}
