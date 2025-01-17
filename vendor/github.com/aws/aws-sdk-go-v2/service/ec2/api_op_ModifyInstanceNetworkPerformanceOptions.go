// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Change the configuration of the network performance options for an existing
// instance.
func (c *Client) ModifyInstanceNetworkPerformanceOptions(ctx context.Context, params *ModifyInstanceNetworkPerformanceOptionsInput, optFns ...func(*Options)) (*ModifyInstanceNetworkPerformanceOptionsOutput, error) {
	if params == nil {
		params = &ModifyInstanceNetworkPerformanceOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyInstanceNetworkPerformanceOptions", params, optFns, c.addOperationModifyInstanceNetworkPerformanceOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyInstanceNetworkPerformanceOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyInstanceNetworkPerformanceOptionsInput struct {

	// Specify the bandwidth weighting option to boost the associated type of baseline
	// bandwidth, as follows:
	//
	// default This option uses the standard bandwidth configuration for your instance
	// type.
	//
	// vpc-1 This option boosts your networking baseline bandwidth and reduces your
	// EBS baseline bandwidth.
	//
	// ebs-1 This option boosts your EBS baseline bandwidth and reduces your
	// networking baseline bandwidth.
	//
	// This member is required.
	BandwidthWeighting types.InstanceBandwidthWeighting

	// The ID of the instance to update.
	//
	// This member is required.
	InstanceId *string

	// Checks whether you have the required permissions for the operation, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type ModifyInstanceNetworkPerformanceOptionsOutput struct {

	// Contains the updated configuration for bandwidth weighting on the specified
	// instance.
	BandwidthWeighting types.InstanceBandwidthWeighting

	// The instance ID that was updated.
	InstanceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyInstanceNetworkPerformanceOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyInstanceNetworkPerformanceOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyInstanceNetworkPerformanceOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyInstanceNetworkPerformanceOptions"); err != nil {
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
	if err = addOpModifyInstanceNetworkPerformanceOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyInstanceNetworkPerformanceOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyInstanceNetworkPerformanceOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyInstanceNetworkPerformanceOptions",
	}
}
