// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the map style descriptor from a map resource. The style descriptor
// contains speciﬁcations on how features render on a map. For example, what data
// to display, what order to display the data in, and the style for the data. Style
// descriptors follow the Mapbox Style Specification.
func (c *Client) GetMapStyleDescriptor(ctx context.Context, params *GetMapStyleDescriptorInput, optFns ...func(*Options)) (*GetMapStyleDescriptorOutput, error) {
	if params == nil {
		params = &GetMapStyleDescriptorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMapStyleDescriptor", params, optFns, addOperationGetMapStyleDescriptorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMapStyleDescriptorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMapStyleDescriptorInput struct {

	// The map resource to retrieve the style descriptor from.
	//
	// This member is required.
	MapName *string
}

type GetMapStyleDescriptorOutput struct {

	// Contains the body of the style descriptor.
	Blob []byte

	// The style descriptor's content type. For example, application/json.
	ContentType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetMapStyleDescriptorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMapStyleDescriptor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMapStyleDescriptor{}, middleware.After)
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
	if err = addOpGetMapStyleDescriptorValidationMiddleware(stack); err != nil {
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
