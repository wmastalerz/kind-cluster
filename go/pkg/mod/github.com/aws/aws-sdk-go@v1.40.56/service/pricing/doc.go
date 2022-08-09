// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package pricing provides the client and types for making API
// requests to AWS Price List Service.
//
// Amazon Web Services Price List Service API (Amazon Web Services Price List
// Service) is a centralized and convenient way to programmatically query Amazon
// Web Services for services, products, and pricing information. The Amazon
// Web Services Price List Service uses standardized product attributes such
// as Location, Storage Class, and Operating System, and provides prices at
// the SKU level. You can use the Amazon Web Services Price List Service to
// build cost control and scenario planning tools, reconcile billing data, forecast
// future spend for budgeting purposes, and provide cost benefit analysis that
// compare your internal workloads with Amazon Web Services.
//
// Use GetServices without a service code to retrieve the service codes for
// all AWS services, then GetServices with a service code to retreive the attribute
// names for that service. After you have the service code and attribute names,
// you can use GetAttributeValues to see what values are available for an attribute.
// With the service code and an attribute name and value, you can use GetProducts
// to find specific products that you're interested in, such as an AmazonEC2
// instance, with a Provisioned IOPS volumeType.
//
// Service Endpoint
//
// Amazon Web Services Price List Service API provides the following two endpoints:
//
//    * https://api.pricing.us-east-1.amazonaws.com
//
//    * https://api.pricing.ap-south-1.amazonaws.com
//
// See https://docs.aws.amazon.com/goto/WebAPI/pricing-2017-10-15 for more information on this service.
//
// See pricing package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/pricing/
//
// Using the Client
//
// To contact AWS Price List Service with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Price List Service client Pricing for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/pricing/#New
package pricing
