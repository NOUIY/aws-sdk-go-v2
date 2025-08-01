// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// For a given input face ID, searches for matching faces in the collection the
// face belongs to. You get a face ID when you add a face to the collection using
// the IndexFacesoperation. The operation compares the features of the input face with faces
// in the specified collection.
//
// You can also search faces without indexing faces by using the SearchFacesByImage
// operation.
//
// The operation response returns an array of faces that match, ordered by
// similarity score with the highest similarity first. More specifically, it is an
// array of metadata for each face match that is found. Along with the metadata,
// the response also includes a confidence value for each face match, indicating
// the confidence that the specific face matches the input face.
//
// For an example, see Searching for a face using its face ID in the Amazon
// Rekognition Developer Guide.
//
// This operation requires permissions to perform the rekognition:SearchFaces
// action.
func (c *Client) SearchFaces(ctx context.Context, params *SearchFacesInput, optFns ...func(*Options)) (*SearchFacesOutput, error) {
	if params == nil {
		params = &SearchFacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchFaces", params, optFns, c.addOperationSearchFacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchFacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchFacesInput struct {

	// ID of the collection the face belongs to.
	//
	// This member is required.
	CollectionId *string

	// ID of a face to find matches for in the collection.
	//
	// This member is required.
	FaceId *string

	// Optional value specifying the minimum confidence in the face match to return.
	// For example, don't return any matches where confidence in matches is less than
	// 70%. The default value is 80%.
	FaceMatchThreshold *float32

	// Maximum number of faces to return. The operation returns the maximum number of
	// faces with the highest confidence in the match.
	MaxFaces *int32

	noSmithyDocumentSerde
}

type SearchFacesOutput struct {

	// An array of faces that matched the input face, along with the confidence in the
	// match.
	FaceMatches []types.FaceMatch

	// Version number of the face detection model associated with the input collection
	// ( CollectionId ).
	FaceModelVersion *string

	// ID of the face that was searched for matches in a collection.
	SearchedFaceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchFacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchFaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchFaces{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchFaces"); err != nil {
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
	if err = addOpSearchFacesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchFaces(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSearchFaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchFaces",
	}
}
