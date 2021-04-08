// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows you to purchase a reserved cache node offering. Reserved nodes are not
// eligible for cancellation and are non-refundable. For more information, see
// Managing Costs with Reserved Nodes
// (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/reserved-nodes.html)
// for Redis or Managing Costs with Reserved Nodes
// (https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/reserved-nodes.html)
// for Memcached.
func (c *Client) PurchaseReservedCacheNodesOffering(ctx context.Context, params *PurchaseReservedCacheNodesOfferingInput, optFns ...func(*Options)) (*PurchaseReservedCacheNodesOfferingOutput, error) {
	if params == nil {
		params = &PurchaseReservedCacheNodesOfferingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PurchaseReservedCacheNodesOffering", params, optFns, addOperationPurchaseReservedCacheNodesOfferingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PurchaseReservedCacheNodesOfferingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a PurchaseReservedCacheNodesOffering operation.
type PurchaseReservedCacheNodesOfferingInput struct {

	// The ID of the reserved cache node offering to purchase. Example:
	// 438012d3-4052-4cc7-b2e3-8d3372e0e706
	//
	// This member is required.
	ReservedCacheNodesOfferingId *string

	// The number of cache node instances to reserve. Default: 1
	CacheNodeCount *int32

	// A customer-specified identifier to track this reservation. The Reserved Cache
	// Node ID is an unique customer-specified identifier to track this reservation. If
	// this parameter is not specified, ElastiCache automatically generates an
	// identifier for the reservation. Example: myreservationID
	ReservedCacheNodeId *string

	// A list of tags to be added to this resource. A tag is a key-value pair. A tag
	// key must be accompanied by a tag value, although null is accepted.
	Tags []types.Tag
}

type PurchaseReservedCacheNodesOfferingOutput struct {

	// Represents the output of a PurchaseReservedCacheNodesOffering operation.
	ReservedCacheNode *types.ReservedCacheNode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPurchaseReservedCacheNodesOfferingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPurchaseReservedCacheNodesOffering{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPurchaseReservedCacheNodesOffering{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPurchaseReservedCacheNodesOfferingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPurchaseReservedCacheNodesOffering(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opPurchaseReservedCacheNodesOffering(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "PurchaseReservedCacheNodesOffering",
	}
}
