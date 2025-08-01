// Code generated by smithy-go-codegen DO NOT EDIT.

package arcregionswitch

import (
	"context"
	smithy "github.com/aws/smithy-go"
	smithyauth "github.com/aws/smithy-go/auth"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

// For custom endpoint with region not set and fips disabled
func TestEndpointCase0(t *testing.T) {
	var params = EndpointParameters{
		Endpoint: ptr.String("https://example.com"),
		UseFIPS:  ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://example.com")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For custom endpoint with fips enabled
func TestEndpointCase1(t *testing.T) {
	var params = EndpointParameters{
		Endpoint: ptr.String("https://example.com"),
		UseFIPS:  ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "Invalid Configuration: FIPS and custom endpoint are not supported", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// For region us-east-1 with FIPS enabled and DualStack enabled
func TestEndpointCase2(t *testing.T) {
	var params = EndpointParameters{
		Region:  ptr.String("us-east-1"),
		UseFIPS: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch-fips.us-east-1.api.aws")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-east-1 with FIPS disabled and DualStack enabled
func TestEndpointCase3(t *testing.T) {
	var params = EndpointParameters{
		Region:  ptr.String("us-east-1"),
		UseFIPS: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch.us-east-1.api.aws")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region cn-northwest-1 with FIPS enabled and DualStack enabled
func TestEndpointCase4(t *testing.T) {
	var params = EndpointParameters{
		Region:  ptr.String("cn-northwest-1"),
		UseFIPS: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch-fips.cn-northwest-1.api.amazonwebservices.com.cn")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region cn-northwest-1 with FIPS disabled and DualStack enabled
func TestEndpointCase5(t *testing.T) {
	var params = EndpointParameters{
		Region:  ptr.String("cn-northwest-1"),
		UseFIPS: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch.cn-northwest-1.api.amazonwebservices.com.cn")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-gov-west-1 with FIPS enabled and DualStack enabled
func TestEndpointCase6(t *testing.T) {
	var params = EndpointParameters{
		Region:  ptr.String("us-gov-west-1"),
		UseFIPS: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch-fips.us-gov-west-1.api.aws")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-gov-west-1 with FIPS disabled and DualStack enabled
func TestEndpointCase7(t *testing.T) {
	var params = EndpointParameters{
		Region:  ptr.String("us-gov-west-1"),
		UseFIPS: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch.us-gov-west-1.api.aws")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-iso-east-1 with FIPS enabled and DualStack enabled
func TestEndpointCase8(t *testing.T) {
	var params = EndpointParameters{
		Region:  ptr.String("us-iso-east-1"),
		UseFIPS: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch-fips.us-iso-east-1.c2s.ic.gov")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-iso-east-1 with FIPS disabled and DualStack enabled
func TestEndpointCase9(t *testing.T) {
	var params = EndpointParameters{
		Region:  ptr.String("us-iso-east-1"),
		UseFIPS: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch.us-iso-east-1.c2s.ic.gov")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-isob-east-1 with FIPS enabled and DualStack enabled
func TestEndpointCase10(t *testing.T) {
	var params = EndpointParameters{
		Region:  ptr.String("us-isob-east-1"),
		UseFIPS: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch-fips.us-isob-east-1.sc2s.sgov.gov")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-isob-east-1 with FIPS disabled and DualStack enabled
func TestEndpointCase11(t *testing.T) {
	var params = EndpointParameters{
		Region:  ptr.String("us-isob-east-1"),
		UseFIPS: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch.us-isob-east-1.sc2s.sgov.gov")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region eu-isoe-west-1 with FIPS enabled and DualStack enabled
func TestEndpointCase12(t *testing.T) {
	var params = EndpointParameters{
		Region:  ptr.String("eu-isoe-west-1"),
		UseFIPS: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch-fips.eu-isoe-west-1.cloud.adc-e.uk")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region eu-isoe-west-1 with FIPS disabled and DualStack enabled
func TestEndpointCase13(t *testing.T) {
	var params = EndpointParameters{
		Region:  ptr.String("eu-isoe-west-1"),
		UseFIPS: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch.eu-isoe-west-1.cloud.adc-e.uk")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-isof-south-1 with FIPS enabled and DualStack enabled
func TestEndpointCase14(t *testing.T) {
	var params = EndpointParameters{
		Region:  ptr.String("us-isof-south-1"),
		UseFIPS: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch-fips.us-isof-south-1.csp.hci.ic.gov")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// For region us-isof-south-1 with FIPS disabled and DualStack enabled
func TestEndpointCase15(t *testing.T) {
	var params = EndpointParameters{
		Region:  ptr.String("us-isof-south-1"),
		UseFIPS: ptr.Bool(false),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch.us-isof-south-1.csp.hci.ic.gov")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// Missing region
func TestEndpointCase16(t *testing.T) {
	var params = EndpointParameters{}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "Invalid Configuration: Missing Region", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// Control plane operation with DualStack in us-west-2 routes to us-east-1
// DualStack endpoint
func TestEndpointCase17(t *testing.T) {
	var params = EndpointParameters{
		Region:                  ptr.String("us-west-2"),
		UseControlPlaneEndpoint: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch-control-plane.us-east-1.api.aws")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")

						smithyhttp.SetSigV4SigningName(&sp, "arc-region-switch")
						smithyhttp.SetSigV4ASigningName(&sp, "arc-region-switch")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// Control plane operation with endpoint set in us-east-1 routes to provided
// endpoint
func TestEndpointCase18(t *testing.T) {
	var params = EndpointParameters{
		Region:                  ptr.String("us-east-1"),
		UseControlPlaneEndpoint: ptr.Bool(true),
		Endpoint:                ptr.String("https://amazonaws.com"),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://amazonaws.com")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// Control plane operation with endpoint set in us-west-2 routes to provided
// endpoint
func TestEndpointCase19(t *testing.T) {
	var params = EndpointParameters{
		Region:                  ptr.String("us-west-2"),
		UseControlPlaneEndpoint: ptr.Bool(true),
		Endpoint:                ptr.String("https://amazonaws.com"),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://amazonaws.com")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:        *uri,
		Headers:    http.Header{},
		Properties: smithy.Properties{},
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// Control plane operation in us-west-2 (standard partition) routes to us-east-1
func TestEndpointCase20(t *testing.T) {
	var params = EndpointParameters{
		Region:                  ptr.String("us-west-2"),
		UseControlPlaneEndpoint: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch-control-plane.us-east-1.api.aws")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")

						smithyhttp.SetSigV4SigningName(&sp, "arc-region-switch")
						smithyhttp.SetSigV4ASigningName(&sp, "arc-region-switch")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// Control plane operation in cn-north-1 (China partition) routes to cn-north-1
// with China DNS suffix
func TestEndpointCase21(t *testing.T) {
	var params = EndpointParameters{
		Region:                  ptr.String("cn-north-1"),
		UseControlPlaneEndpoint: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch-control-plane.cn-north-1.api.amazonwebservices.com.cn")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningRegion(&sp, "cn-north-1")

						smithyhttp.SetSigV4SigningName(&sp, "arc-region-switch")
						smithyhttp.SetSigV4ASigningName(&sp, "arc-region-switch")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// Control plane operation in cn-northwest-1 (China partition) routes to cn-north-1
// with China DNS suffix
func TestEndpointCase22(t *testing.T) {
	var params = EndpointParameters{
		Region:                  ptr.String("cn-northwest-1"),
		UseControlPlaneEndpoint: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch-control-plane.cn-north-1.api.amazonwebservices.com.cn")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningRegion(&sp, "cn-north-1")

						smithyhttp.SetSigV4SigningName(&sp, "arc-region-switch")
						smithyhttp.SetSigV4ASigningName(&sp, "arc-region-switch")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// Control plane operation in us-gov-west-1 (GovCloud partition) routes to
// us-gov-west-1 with GovCloud DNS suffix
func TestEndpointCase23(t *testing.T) {
	var params = EndpointParameters{
		Region:                  ptr.String("us-gov-west-1"),
		UseControlPlaneEndpoint: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch-control-plane.us-gov-west-1.api.aws")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningRegion(&sp, "us-gov-west-1")

						smithyhttp.SetSigV4SigningName(&sp, "arc-region-switch")
						smithyhttp.SetSigV4ASigningName(&sp, "arc-region-switch")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// Control plane operation in us-gov-east-1 (GovCloud partition) routes to
// us-gov-west-1 with GovCloud DNS suffix
func TestEndpointCase24(t *testing.T) {
	var params = EndpointParameters{
		Region:                  ptr.String("us-gov-east-1"),
		UseControlPlaneEndpoint: ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch-control-plane.us-gov-west-1.api.aws")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningRegion(&sp, "us-gov-west-1")

						smithyhttp.SetSigV4SigningName(&sp, "arc-region-switch")
						smithyhttp.SetSigV4ASigningName(&sp, "arc-region-switch")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// Control plane operation with FIPS in us-west-2 routes to us-east-1 FIPS endpoint
func TestEndpointCase25(t *testing.T) {
	var params = EndpointParameters{
		Region:                  ptr.String("us-west-2"),
		UseControlPlaneEndpoint: ptr.Bool(true),
		UseFIPS:                 ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch-control-plane-fips.us-east-1.api.aws")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")

						smithyhttp.SetSigV4SigningName(&sp, "arc-region-switch")
						smithyhttp.SetSigV4ASigningName(&sp, "arc-region-switch")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// Control plane operation with FIPS in us-east-1 routes to us-east-1 FIPS endpoint
func TestEndpointCase26(t *testing.T) {
	var params = EndpointParameters{
		Region:                  ptr.String("us-east-1"),
		UseControlPlaneEndpoint: ptr.Bool(true),
		UseFIPS:                 ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	uri, _ := url.Parse("https://arc-region-switch-control-plane-fips.us-east-1.api.aws")

	expectEndpoint := smithyendpoints.Endpoint{
		URI:     *uri,
		Headers: http.Header{},
		Properties: func() smithy.Properties {
			var out smithy.Properties
			smithyauth.SetAuthOptions(&out, []*smithyauth.Option{
				{
					SchemeID: "aws.auth#sigv4",
					SignerProperties: func() smithy.Properties {
						var sp smithy.Properties
						smithyhttp.SetSigV4SigningRegion(&sp, "us-east-1")

						smithyhttp.SetSigV4SigningName(&sp, "arc-region-switch")
						smithyhttp.SetSigV4ASigningName(&sp, "arc-region-switch")
						return sp
					}(),
				},
			})
			return out
		}(),
	}

	if e, a := expectEndpoint.URI, result.URI; e != a {
		t.Errorf("expect %v URI, got %v", e, a)
	}

	if !reflect.DeepEqual(expectEndpoint.Headers, result.Headers) {
		t.Errorf("expect headers to match\n%v != %v", expectEndpoint.Headers, result.Headers)
	}

	if !reflect.DeepEqual(expectEndpoint.Properties, result.Properties) {
		t.Errorf("expect properties to match\n%v != %v", expectEndpoint.Properties, result.Properties)
	}
}

// Control plane operation with FIPS in CN returns an error
func TestEndpointCase27(t *testing.T) {
	var params = EndpointParameters{
		Region:                  ptr.String("cn-north-1"),
		UseControlPlaneEndpoint: ptr.Bool(true),
		UseFIPS:                 ptr.Bool(true),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "Invalid Configuration: FIPS is not supported in this partition", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// Control plane operation with endpoint set using FIPS in us-east-1 errors
func TestEndpointCase28(t *testing.T) {
	var params = EndpointParameters{
		Region:                  ptr.String("us-east-1"),
		UseControlPlaneEndpoint: ptr.Bool(true),
		UseFIPS:                 ptr.Bool(true),
		Endpoint:                ptr.String("https://amazonaws.com"),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "Invalid Configuration: FIPS and custom endpoint are not supported", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}

// Control plane operation with endpoint set using FIPS in us-west-2 routes to
// provided endpoint
func TestEndpointCase29(t *testing.T) {
	var params = EndpointParameters{
		Region:                  ptr.String("us-west-2"),
		UseControlPlaneEndpoint: ptr.Bool(true),
		UseFIPS:                 ptr.Bool(true),
		Endpoint:                ptr.String("https://amazonaws.com"),
	}

	resolver := NewDefaultEndpointResolverV2()
	result, err := resolver.ResolveEndpoint(context.Background(), params)
	_, _ = result, err

	if err == nil {
		t.Fatalf("expect error, got none")
	}
	if e, a := "Invalid Configuration: FIPS and custom endpoint are not supported", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %v error in %v", e, a)
	}
}
