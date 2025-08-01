// Code generated by smithy-go-codegen DO NOT EDIT.

// Package kms provides the API client, operations, and parameter types for AWS
// Key Management Service.
//
// # Key Management Service
//
// Key Management Service (KMS) is an encryption and key management web service.
// This guide describes the KMS operations that you can call programmatically. For
// general information about KMS, see the [Key Management Service Developer Guide].
//
// KMS has replaced the term customer master key (CMK) with Key Management Service
// key and KMS key. The concept has not changed. To prevent breaking changes, KMS
// is keeping some variations of this term.
//
// Amazon Web Services provides SDKs that consist of libraries and sample code for
// various programming languages and platforms (Java, Rust, Python, Ruby, .Net,
// macOS, Android, etc.). The SDKs provide a convenient way to create programmatic
// access to KMS and other Amazon Web Services services. For example, the SDKs take
// care of tasks such as signing requests (see below), managing errors, and
// retrying requests automatically. For more information about the Amazon Web
// Services SDKs, including how to download and install them, see [Tools for Amazon Web Services].
//
// We recommend that you use the Amazon Web Services SDKs to make programmatic API
// calls to KMS.
//
// If you need to use FIPS 140-2 validated cryptographic modules when
// communicating with Amazon Web Services, use one of the FIPS endpoints in your
// preferred Amazon Web Services Region. If you need communicate over IPv6, use the
// dual-stack endpoint in your preferred Amazon Web Services Region. For more
// information see [Service endpoints]in the Key Management Service topic of the Amazon Web Services
// General Reference and [Dual-stack endpoint support]in the KMS Developer Guide.
//
// All KMS API calls must be signed and be transmitted using Transport Layer
// Security (TLS). KMS recommends you always use the latest supported TLS version.
// Clients must also support cipher suites with Perfect Forward Secrecy (PFS) such
// as Ephemeral Diffie-Hellman (DHE) or Elliptic Curve Ephemeral Diffie-Hellman
// (ECDHE). Most modern systems such as Java 7 and later support these modes.
//
// # Signing Requests
//
// Requests must be signed using an access key ID and a secret access key. We
// strongly recommend that you do not use your Amazon Web Services account root
// access key ID and secret access key for everyday work. You can use the access
// key ID and secret access key for an IAM user or you can use the Security Token
// Service (STS) to generate temporary security credentials and use those to sign
// requests.
//
// All KMS requests must be signed with [Signature Version 4].
//
// # Logging API Requests
//
// KMS supports CloudTrail, a service that logs Amazon Web Services API calls and
// related events for your Amazon Web Services account and delivers them to an
// Amazon S3 bucket that you specify. By using the information collected by
// CloudTrail, you can determine what requests were made to KMS, who made the
// request, when it was made, and so on. To learn more about CloudTrail, including
// how to turn it on and find your log files, see the [CloudTrail User Guide].
//
// # Additional Resources
//
// For more information about credentials and request signing, see the following:
//
// [Amazon Web Services Security Credentials]
//   - - This topic provides general information about the types of credentials
//     used to access Amazon Web Services.
//
// [Temporary Security Credentials]
//   - - This section of the IAM User Guide describes how to create and use
//     temporary security credentials.
//
// [Signature Version 4 Signing Process]
//   - - This set of topics walks you through the process of signing a request
//     using an access key ID and a secret access key.
//
// # Commonly Used API Operations
//
// Of the API operations discussed in this guide, the following will prove the
// most useful for most applications. You will likely perform operations other than
// these, such as creating keys and assigning policies, by using the console.
//
// # Encrypt
//
// # Decrypt
//
// # GenerateDataKey
//
// # GenerateDataKeyWithoutPlaintext
//
// [Dual-stack endpoint support]: https://docs.aws.amazon.com/kms/latest/developerguide/ipv6-kms.html
// [Signature Version 4]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
// [Temporary Security Credentials]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp.html
// [Tools for Amazon Web Services]: http://aws.amazon.com/tools/
// [Amazon Web Services Security Credentials]: https://docs.aws.amazon.com/general/latest/gr/aws-security-credentials.html
// [Key Management Service Developer Guide]: https://docs.aws.amazon.com/kms/latest/developerguide/
// [Service endpoints]: https://docs.aws.amazon.com/general/latest/gr/kms.html#kms_region
// [CloudTrail User Guide]: https://docs.aws.amazon.com/awscloudtrail/latest/userguide/
// [Signature Version 4 Signing Process]: https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html
package kms
