// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

//	Returns an asset (or file) that is in a package. For example, for a Maven
//
// package version, use GetPackageVersionAsset to download a JAR file, a POM file,
// or any other assets in the package version.
func (c *Client) GetPackageVersionAsset(ctx context.Context, params *GetPackageVersionAssetInput, optFns ...func(*Options)) (*GetPackageVersionAssetOutput, error) {
	if params == nil {
		params = &GetPackageVersionAssetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPackageVersionAsset", params, optFns, c.addOperationGetPackageVersionAssetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPackageVersionAssetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPackageVersionAssetInput struct {

	//  The name of the requested asset.
	//
	// This member is required.
	Asset *string

	//  The name of the domain that contains the repository that contains the package
	// version with the requested asset.
	//
	// This member is required.
	Domain *string

	//  A format that specifies the type of the package version with the requested
	// asset file.
	//
	// This member is required.
	Format types.PackageFormat

	//  The name of the package that contains the requested asset.
	//
	// This member is required.
	Package *string

	//  A string that contains the package version (for example, 3.5.2 ).
	//
	// This member is required.
	PackageVersion *string

	//  The repository that contains the package version with the requested asset.
	//
	// This member is required.
	Repository *string

	//  The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	// The namespace of the package version with the requested asset file. The package
	// component that specifies its namespace depends on its type. For example:
	//
	// The namespace is required when requesting assets from package versions of the
	// following formats:
	//
	//   - Maven
	//
	//   - Swift
	//
	//   - generic
	//
	//   - The namespace of a Maven package version is its groupId .
	//
	//   - The namespace of an npm or Swift package version is its scope .
	//
	//   - The namespace of a generic package is its namespace .
	//
	//   - Python, NuGet, Ruby, and Cargo package versions do not contain a
	//   corresponding component, package versions of those formats do not have a
	//   namespace.
	Namespace *string

	//  The name of the package version revision that contains the requested asset.
	PackageVersionRevision *string

	noSmithyDocumentSerde
}

type GetPackageVersionAssetOutput struct {

	//  The binary file, or asset, that is downloaded.
	Asset io.ReadCloser

	//  The name of the asset that is downloaded.
	AssetName *string

	//  A string that contains the package version (for example, 3.5.2 ).
	PackageVersion *string

	//  The name of the package version revision that contains the downloaded asset.
	PackageVersionRevision *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPackageVersionAssetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPackageVersionAsset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPackageVersionAsset{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetPackageVersionAsset"); err != nil {
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
	if err = addOpGetPackageVersionAssetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPackageVersionAsset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPackageVersionAsset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetPackageVersionAsset",
	}
}
