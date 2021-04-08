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

// Used to failover the primary region to a selected secondary region. The selected
// secondary region will become primary, and all other clusters will become
// secondary.
func (c *Client) FailoverGlobalReplicationGroup(ctx context.Context, params *FailoverGlobalReplicationGroupInput, optFns ...func(*Options)) (*FailoverGlobalReplicationGroupOutput, error) {
	if params == nil {
		params = &FailoverGlobalReplicationGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "FailoverGlobalReplicationGroup", params, optFns, addOperationFailoverGlobalReplicationGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*FailoverGlobalReplicationGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type FailoverGlobalReplicationGroupInput struct {

	// The name of the Global datastore
	//
	// This member is required.
	GlobalReplicationGroupId *string

	// The AWS region of the primary cluster of the Global datastore
	//
	// This member is required.
	PrimaryRegion *string

	// The name of the primary replication group
	//
	// This member is required.
	PrimaryReplicationGroupId *string
}

type FailoverGlobalReplicationGroupOutput struct {

	// Consists of a primary cluster that accepts writes and an associated secondary
	// cluster that resides in a different AWS region. The secondary cluster accepts
	// only reads. The primary cluster automatically replicates updates to the
	// secondary cluster.
	//
	// * The GlobalReplicationGroupIdSuffix represents the name of
	// the Global datastore, which is what you use to associate a secondary cluster.
	GlobalReplicationGroup *types.GlobalReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationFailoverGlobalReplicationGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpFailoverGlobalReplicationGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpFailoverGlobalReplicationGroup{}, middleware.After)
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
	if err = addOpFailoverGlobalReplicationGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opFailoverGlobalReplicationGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opFailoverGlobalReplicationGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "FailoverGlobalReplicationGroup",
	}
}
