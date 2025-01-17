// Code generated by smithy-go-codegen DO NOT EDIT.

// Package pricing provides the API client, operations, and parameter types for
// AWS Price List Service.
//
// The Amazon Web Services Price List API is a centralized and convenient way to
// programmatically query Amazon Web Services for services, products, and pricing
// information. The Amazon Web Services Price List uses standardized product
// attributes such as Location , Storage Class , and Operating System , and
// provides prices at the SKU level. You can use the Amazon Web Services Price List
// to do the following:
//
//   - Build cost control and scenario planning tools
//
//   - Reconcile billing data
//
//   - Forecast future spend for budgeting purposes
//
//   - Provide cost benefit analysis that compare your internal workloads with
//     Amazon Web Services
//
// Use GetServices without a service code to retrieve the service codes for all
// Amazon Web Services services, then GetServices with a service code to retrieve
// the attribute names for that service. After you have the service code and
// attribute names, you can use GetAttributeValues to see what values are
// available for an attribute. With the service code and an attribute name and
// value, you can use GetProducts to find specific products that you're interested
// in, such as an AmazonEC2 instance, with a Provisioned IOPS volumeType .
//
// For more information, see [Using the Amazon Web Services Price List API] in the Billing User Guide.
//
// [Using the Amazon Web Services Price List API]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/price-changes.html
package pricing
