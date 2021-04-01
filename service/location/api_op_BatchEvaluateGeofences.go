// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Evaluates device positions against the geofence geometries from a given geofence
// collection. The evaluation determines if the device has entered or exited a
// geofenced area, which publishes ENTER or EXIT geofence events to Amazon
// EventBridge. The last geofence that a device was observed within, if any, is
// tracked for 30 days after the most recent device position update
func (c *Client) BatchEvaluateGeofences(ctx context.Context, params *BatchEvaluateGeofencesInput, optFns ...func(*Options)) (*BatchEvaluateGeofencesOutput, error) {
	if params == nil {
		params = &BatchEvaluateGeofencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchEvaluateGeofences", params, optFns, addOperationBatchEvaluateGeofencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchEvaluateGeofencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchEvaluateGeofencesInput struct {

	// The geofence collection used in evaluating the position of devices against its
	// geofences.
	//
	// This member is required.
	CollectionName *string

	// Contains device details for each device to be evaluated against the given
	// geofence collection.
	//
	// This member is required.
	DevicePositionUpdates []types.DevicePositionUpdate
}

type BatchEvaluateGeofencesOutput struct {

	// Contains error details for each device that failed to evaluate its position
	// against the given geofence collection.
	//
	// This member is required.
	Errors []types.BatchEvaluateGeofencesError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchEvaluateGeofencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchEvaluateGeofences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchEvaluateGeofences{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = addOpBatchEvaluateGeofencesValidationMiddleware(stack); err != nil {
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
