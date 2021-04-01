// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a geofence collection from your AWS account. This action deletes the
// resource permanently. You can't undo this action. If the geofence collection is
// the target of a tracker resource, the devices will no longer be monitored.
func (c *Client) DeleteGeofenceCollection(ctx context.Context, params *DeleteGeofenceCollectionInput, optFns ...func(*Options)) (*DeleteGeofenceCollectionOutput, error) {
	if params == nil {
		params = &DeleteGeofenceCollectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteGeofenceCollection", params, optFns, addOperationDeleteGeofenceCollectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteGeofenceCollectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteGeofenceCollectionInput struct {

	// The name of the geofence collection to be deleted.
	//
	// This member is required.
	CollectionName *string
}

type DeleteGeofenceCollectionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteGeofenceCollectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteGeofenceCollection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteGeofenceCollection{}, middleware.After)
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
	if err = addOpDeleteGeofenceCollectionValidationMiddleware(stack); err != nil {
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
