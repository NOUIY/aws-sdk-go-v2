package config

import (
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting"
	"github.com/aws/smithy-go/ptr"
)

var _ sharedConfigProfileProvider = (*EnvConfig)(nil)
var _ sharedConfigFilesProvider = (*EnvConfig)(nil)
var _ customCABundleProvider = (*EnvConfig)(nil)
var _ regionProvider = (*EnvConfig)(nil)

func TestNewEnvConfig_Creds(t *testing.T) {
	restoreEnv := awstesting.StashEnv()
	defer awstesting.PopEnv(restoreEnv)

	cases := []struct {
		Env map[string]string
		Val aws.Credentials
	}{
		{
			Env: map[string]string{
				"AWS_ACCESS_KEY": "AKID",
			},
			Val: aws.Credentials{},
		},
		{
			Env: map[string]string{
				"AWS_ACCESS_KEY_ID": "AKID",
			},
			Val: aws.Credentials{},
		},
		{
			Env: map[string]string{
				"AWS_SECRET_KEY": "SECRET",
			},
			Val: aws.Credentials{},
		},
		{
			Env: map[string]string{
				"AWS_SECRET_ACCESS_KEY": "SECRET",
			},
			Val: aws.Credentials{},
		},
		{
			Env: map[string]string{
				"AWS_ACCESS_KEY_ID":     "AKID",
				"AWS_SECRET_ACCESS_KEY": "SECRET",
			},
			Val: aws.Credentials{
				AccessKeyID: "AKID", SecretAccessKey: "SECRET",
				Source: CredentialsSourceName,
			},
		},
		{
			Env: map[string]string{
				"AWS_ACCESS_KEY": "AKID",
				"AWS_SECRET_KEY": "SECRET",
			},
			Val: aws.Credentials{
				AccessKeyID: "AKID", SecretAccessKey: "SECRET",
				Source: CredentialsSourceName,
			},
		},
		{
			Env: map[string]string{
				"AWS_ACCESS_KEY":    "AKID",
				"AWS_SECRET_KEY":    "SECRET",
				"AWS_SESSION_TOKEN": "TOKEN",
			},
			Val: aws.Credentials{
				AccessKeyID: "AKID", SecretAccessKey: "SECRET", SessionToken: "TOKEN",
				Source: CredentialsSourceName,
			},
		},
		{
			Env: map[string]string{
				"AWS_ACCESS_KEY_ID":     "AKID",
				"AWS_SECRET_ACCESS_KEY": "SECRET",
				"AWS_ACCOUNT_ID":        "012345678901",
			},
			Val: aws.Credentials{
				AccessKeyID: "AKID", SecretAccessKey: "SECRET", AccountID: "012345678901",
				Source: CredentialsSourceName,
			},
		},
	}

	for i, c := range cases {
		os.Clearenv()

		for k, v := range c.Env {
			os.Setenv(k, v)
		}

		cfg, err := NewEnvConfig()
		if err != nil {
			t.Fatalf("%d, expect no error, got %v", i, err)
		}

		if !reflect.DeepEqual(c.Val, cfg.Credentials) {
			t.Errorf("%d, expect aws to match.\n%s", i,
				awstesting.SprintExpectActual(c.Val, cfg.Credentials))
		}
	}
}

func TestNewEnvConfig(t *testing.T) {
	restoreEnv := awstesting.StashEnv()
	defer awstesting.PopEnv(restoreEnv)

	cases := []struct {
		Env     map[string]string
		Config  EnvConfig
		WantErr bool
	}{
		0: {
			Env: map[string]string{
				"AWS_REGION":  "region",
				"AWS_PROFILE": "profile",
			},
			Config: EnvConfig{
				Region: "region", SharedConfigProfile: "profile",
			},
		},
		1: {
			Env: map[string]string{
				"AWS_REGION":          "region",
				"AWS_DEFAULT_REGION":  "default_region",
				"AWS_PROFILE":         "profile",
				"AWS_DEFAULT_PROFILE": "default_profile",
			},
			Config: EnvConfig{
				Region: "region", SharedConfigProfile: "profile",
			},
		},
		2: {
			Env: map[string]string{
				"AWS_REGION":          "region",
				"AWS_DEFAULT_REGION":  "default_region",
				"AWS_PROFILE":         "profile",
				"AWS_DEFAULT_PROFILE": "default_profile",
			},
			Config: EnvConfig{
				Region: "region", SharedConfigProfile: "profile",
			},
		},
		3: {
			Env: map[string]string{
				"AWS_DEFAULT_REGION":  "default_region",
				"AWS_DEFAULT_PROFILE": "default_profile",
			},
			Config: EnvConfig{
				Region: "default_region", SharedConfigProfile: "default_profile",
			},
		},
		4: {
			Env: map[string]string{
				"AWS_REGION":  "region",
				"AWS_PROFILE": "profile",
			},
			Config: EnvConfig{
				Region: "region", SharedConfigProfile: "profile",
			},
		},
		5: {
			Env: map[string]string{
				"AWS_REGION":          "region",
				"AWS_DEFAULT_REGION":  "default_region",
				"AWS_PROFILE":         "profile",
				"AWS_DEFAULT_PROFILE": "default_profile",
			},
			Config: EnvConfig{
				Region: "region", SharedConfigProfile: "profile",
			},
		},
		6: {
			Env: map[string]string{
				"AWS_REGION":          "region",
				"AWS_DEFAULT_REGION":  "default_region",
				"AWS_PROFILE":         "profile",
				"AWS_DEFAULT_PROFILE": "default_profile",
			},
			Config: EnvConfig{
				Region: "region", SharedConfigProfile: "profile",
			},
		},
		7: {
			Env: map[string]string{
				"AWS_DEFAULT_REGION":  "default_region",
				"AWS_DEFAULT_PROFILE": "default_profile",
			},
			Config: EnvConfig{
				Region: "default_region", SharedConfigProfile: "default_profile",
			},
		},
		8: {
			Env: map[string]string{
				"AWS_CA_BUNDLE": "custom_ca_bundle",
			},
			Config: EnvConfig{
				CustomCABundle: "custom_ca_bundle",
			},
		},
		9: {
			Env: map[string]string{
				"AWS_CA_BUNDLE": "custom_ca_bundle",
			},
			Config: EnvConfig{
				CustomCABundle: "custom_ca_bundle",
			},
		},
		10: {
			Env: map[string]string{
				"AWS_SHARED_CREDENTIALS_FILE": "/path/to/aws/file",
				"AWS_CONFIG_FILE":             "/path/to/config/file",
			},
			Config: EnvConfig{
				SharedCredentialsFile: "/path/to/aws/file",
				SharedConfigFile:      "/path/to/config/file",
			},
		},
		11: {
			Env: map[string]string{
				"AWS_S3_USE_ARN_REGION": "true",
			},
			Config: EnvConfig{
				S3UseARNRegion: ptr.Bool(true),
			},
		},
		12: {
			Env: map[string]string{
				"AWS_ENABLE_ENDPOINT_DISCOVERY": "true",
			},
			Config: EnvConfig{
				EnableEndpointDiscovery: aws.EndpointDiscoveryEnabled,
			},
		},
		13: {
			Env: map[string]string{
				"AWS_ENABLE_ENDPOINT_DISCOVERY": "auto",
			},
			Config: EnvConfig{
				EnableEndpointDiscovery: aws.EndpointDiscoveryAuto,
			},
		},
		14: {
			Env: map[string]string{
				"AWS_ENABLE_ENDPOINT_DISCOVERY": "false",
			},
			Config: EnvConfig{
				EnableEndpointDiscovery: aws.EndpointDiscoveryDisabled,
			},
		},
		15: {
			Env:    map[string]string{},
			Config: EnvConfig{},
		},
		16: {
			Env: map[string]string{
				"AWS_EC2_METADATA_SERVICE_ENDPOINT_MODE": "IPv6",
			},
			Config: EnvConfig{
				EC2IMDSEndpointMode: imds.EndpointModeStateIPv6,
			},
		},
		17: {
			Env: map[string]string{
				"AWS_EC2_METADATA_SERVICE_ENDPOINT_MODE": "IPv4",
			},
			Config: EnvConfig{
				EC2IMDSEndpointMode: imds.EndpointModeStateIPv4,
			},
		},
		18: {
			Env: map[string]string{
				"AWS_EC2_METADATA_SERVICE_ENDPOINT_MODE": "foobar",
			},
			Config:  EnvConfig{},
			WantErr: true,
		},
		19: {
			Env: map[string]string{
				"AWS_EC2_METADATA_SERVICE_ENDPOINT": "http://endpoint.localhost",
			},
			Config: EnvConfig{
				EC2IMDSEndpoint: "http://endpoint.localhost",
			},
		},
		20: {
			Env: map[string]string{
				"AWS_EC2_METADATA_SERVICE_ENDPOINT_MODE": "IPv6",
				"AWS_EC2_METADATA_SERVICE_ENDPOINT":      "http://endpoint.localhost",
			},
			Config: EnvConfig{
				EC2IMDSEndpoint:     "http://endpoint.localhost",
				EC2IMDSEndpointMode: imds.EndpointModeStateIPv6,
			},
		},
		21: {
			Env: map[string]string{
				"AWS_EC2_METADATA_DISABLED": "false",
			},
			Config: EnvConfig{
				EC2IMDSClientEnableState: imds.ClientEnabled,
			},
		},
		22: {
			Env: map[string]string{
				"AWS_EC2_METADATA_DISABLED": "true",
			},
			Config: EnvConfig{
				EC2IMDSClientEnableState: imds.ClientDisabled,
			},
		},
		23: {
			Env: map[string]string{
				"AWS_EC2_METADATA_DISABLED": "foobar",
			},
			Config: EnvConfig{},
		},
		24: {
			Env: map[string]string{
				"AWS_S3_DISABLE_MULTIREGION_ACCESS_POINTS": "true",
			},
			Config: EnvConfig{
				S3DisableMultiRegionAccessPoints: ptr.Bool(true),
			},
		},
		25: {
			Env: map[string]string{
				"AWS_USE_DUALSTACK_ENDPOINT": "true",
			},
			Config: EnvConfig{
				UseDualStackEndpoint: aws.DualStackEndpointStateEnabled,
			},
		},
		26: {
			Env: map[string]string{
				"AWS_USE_DUALSTACK_ENDPOINT": "false",
			},
			Config: EnvConfig{
				UseDualStackEndpoint: aws.DualStackEndpointStateDisabled,
			},
		},
		27: {
			Env: map[string]string{
				"AWS_USE_DUALSTACK_ENDPOINT": "invalid",
			},
			WantErr: true,
		},
		28: {
			Env: map[string]string{
				"AWS_USE_FIPS_ENDPOINT": "true",
			},
			Config: EnvConfig{
				UseFIPSEndpoint: aws.FIPSEndpointStateEnabled,
			},
		},
		29: {
			Env: map[string]string{
				"AWS_USE_FIPS_ENDPOINT": "false",
			},
			Config: EnvConfig{
				UseFIPSEndpoint: aws.FIPSEndpointStateDisabled,
			},
		},
		30: {
			Env: map[string]string{
				"AWS_USE_FIPS_ENDPOINT": "invalid",
			},
			WantErr: true,
		},
		31: {
			Env: map[string]string{
				"AWS_DEFAULTS_MODE": "auto",
			},
			Config: EnvConfig{
				DefaultsMode: aws.DefaultsModeAuto,
			},
		},
		32: {
			Env: map[string]string{
				"AWS_DEFAULTS_MODE": "standard",
			},
			Config: EnvConfig{
				DefaultsMode: aws.DefaultsModeStandard,
			},
		},
		33: {
			Env: map[string]string{
				"AWS_DEFAULTS_MODE": "invalid",
			},
			Config: EnvConfig{
				DefaultsMode: aws.DefaultsMode("invalid"),
			},
			WantErr: true,
		},
		34: {
			Env: map[string]string{
				"AWS_MAX_ATTEMPTS": "2",
			},
			Config: EnvConfig{
				RetryMaxAttempts: 2,
			},
		},
		35: {
			Env: map[string]string{
				"AWS_MAX_ATTEMPTS": "invalid",
			},
			Config:  EnvConfig{},
			WantErr: true,
		},
		36: {
			Env: map[string]string{
				"AWS_RETRY_MODE": "adaptive",
			},
			Config: EnvConfig{
				RetryMode: aws.RetryModeAdaptive,
			},
		},
		37: {
			Env: map[string]string{
				"AWS_RETRY_MODE": "invalid",
			},
			Config:  EnvConfig{},
			WantErr: true,
		},
		38: {
			Env: map[string]string{
				"AWS_ENDPOINT_URL": "https://example.com",
			},
			Config: EnvConfig{
				BaseEndpoint: "https://example.com",
			},
		},
		39: {
			Env: map[string]string{
				"AWS_IGNORE_CONFIGURED_ENDPOINT_URLS": "true",
			},
			Config: EnvConfig{
				IgnoreConfiguredEndpoints: ptr.Bool(true),
			},
		},
		40: {
			Env: map[string]string{
				"AWS_EC2_METADATA_V1_DISABLED": "tRuE",
			},
			Config: EnvConfig{
				EC2IMDSv1Disabled: aws.Bool(true),
			},
		},
		41: {
			Env: map[string]string{
				"AWS_EC2_METADATA_V1_DISABLED": "invalid",
			},
			Config: EnvConfig{
				EC2IMDSv1Disabled: aws.Bool(false), // setBoolPtrFromEnvVal new()s the bool even if it errors
			},
			WantErr: true,
		},
		42: {
			Env: map[string]string{
				"AWS_DISABLE_REQUEST_COMPRESSION":        "true",
				"AWS_REQUEST_MIN_COMPRESSION_SIZE_BYTES": "12345",
			},
			Config: EnvConfig{
				DisableRequestCompression:   aws.Bool(true),
				RequestMinCompressSizeBytes: aws.Int64(12345),
			},
		},
		43: {
			Env: map[string]string{
				"AWS_DISABLE_REQUEST_COMPRESSION":        "blabla",
				"AWS_REQUEST_MIN_COMPRESSION_SIZE_BYTES": "12345",
			},
			Config: EnvConfig{
				DisableRequestCompression: aws.Bool(false),
			},
			WantErr: true,
		},
		44: {
			Env: map[string]string{
				"AWS_DISABLE_REQUEST_COMPRESSION":        "true",
				"AWS_REQUEST_MIN_COMPRESSION_SIZE_BYTES": "1.1",
			},
			Config: EnvConfig{
				DisableRequestCompression: aws.Bool(true),
			},
			WantErr: true,
		},
		// expect err detected due to AWS_REQUEST_MIN_COMPRESSION_SIZE_BYTES exceeding max 10485760
		45: {
			Env: map[string]string{
				"AWS_DISABLE_REQUEST_COMPRESSION":        "false",
				"AWS_REQUEST_MIN_COMPRESSION_SIZE_BYTES": "10485761",
			},
			Config: EnvConfig{
				DisableRequestCompression: aws.Bool(false),
			},
			WantErr: true,
		},
		46: {
			Env: map[string]string{
				"AWS_ACCOUNT_ID_ENDPOINT_MODE": "required",
			},
			Config: EnvConfig{
				AccountIDEndpointMode: aws.AccountIDEndpointModeRequired,
			},
		},
		47: {
			Env: map[string]string{
				"AWS_ACCOUNT_ID_ENDPOINT_MODE": "blabla",
			},
			Config:  EnvConfig{},
			WantErr: true,
		},
		48: {
			Env: map[string]string{
				"AWS_REQUEST_CHECKSUM_CALCULATION": "WHEN_SUPPORTED",
			},
			Config: EnvConfig{
				RequestChecksumCalculation: aws.RequestChecksumCalculationWhenSupported,
			},
		},
		49: {
			Env: map[string]string{
				"AWS_REQUEST_CHECKSUM_CALCULATION": "when_required",
			},
			Config: EnvConfig{
				RequestChecksumCalculation: aws.RequestChecksumCalculationWhenRequired,
			},
		},
		50: {
			Env: map[string]string{
				"AWS_REQUEST_CHECKSUM_CALCULATION": "blabla",
			},
			Config:  EnvConfig{},
			WantErr: true,
		},
		51: {
			Env: map[string]string{
				"AWS_RESPONSE_CHECKSUM_VALIDATION": "WHEN_SUPPORTED",
			},
			Config: EnvConfig{
				ResponseChecksumValidation: aws.ResponseChecksumValidationWhenSupported,
			},
		},
		52: {
			Env: map[string]string{
				"AWS_RESPONSE_CHECKSUM_VALIDATION": "when_Required",
			},
			Config: EnvConfig{
				ResponseChecksumValidation: aws.ResponseChecksumValidationWhenRequired,
			},
		},
		53: {
			Env: map[string]string{
				"AWS_RESPONSE_CHECKSUM_VALIDATION": "blabla",
			},
			Config:  EnvConfig{},
			WantErr: true,
		},
		54: {
			Env: map[string]string{
				"AWS_AUTH_SCHEME_PREFERENCE": "  \tsigv4a\t  ,sigv4 ",
			},
			Config: EnvConfig{
				AuthSchemePreference: []string{"sigv4a", "sigv4"},
			},
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			os.Clearenv()

			for k, v := range c.Env {
				os.Setenv(k, v)
			}

			cfg, err := NewEnvConfig()
			if (err != nil) != c.WantErr {
				t.Fatalf("WantErr=%v, got err=%v", c.WantErr, err)
			}

			if diff := cmpDiff(c.Config, cfg); len(diff) > 0 {
				t.Errorf("expect config to match.\n%s",
					diff)
			}
		})
	}
}

func TestSetEnvValue(t *testing.T) {
	restoreEnv := awstesting.StashEnv()
	defer awstesting.PopEnv(restoreEnv)

	os.Setenv("empty_key", "")
	os.Setenv("second_key", "2")
	os.Setenv("third_key", "3")

	var dst string
	setStringFromEnvVal(&dst, []string{
		"empty_key", "first_key", "second_key", "third_key",
	})

	if e, a := "2", dst; e != a {
		t.Errorf("expect %s value from environment, got %s", e, a)
	}
}
