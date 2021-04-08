// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new image with the latest Windows operating system updates, driver
// updates, and AppStream 2.0 agent software. For more information, see the "Update
// an Image by Using Managed AppStream 2.0 Image Updates" section in Administer
// Your AppStream 2.0 Images
// (https://docs.aws.amazon.com/appstream2/latest/developerguide/administer-images.html),
// in the Amazon AppStream 2.0 Administration Guide.
func (c *Client) CreateUpdatedImage(ctx context.Context, params *CreateUpdatedImageInput, optFns ...func(*Options)) (*CreateUpdatedImageOutput, error) {
	if params == nil {
		params = &CreateUpdatedImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUpdatedImage", params, optFns, addOperationCreateUpdatedImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUpdatedImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUpdatedImageInput struct {

	// The name of the image to update.
	//
	// This member is required.
	ExistingImageName *string

	// The name of the new image. The name must be unique within the AWS account and
	// Region.
	//
	// This member is required.
	NewImageName *string

	// Indicates whether to display the status of image update availability before
	// AppStream 2.0 initiates the process of creating a new updated image. If this
	// value is set to true, AppStream 2.0 displays whether image updates are
	// available. If this value is set to false, AppStream 2.0 initiates the process of
	// creating a new updated image without displaying whether image updates are
	// available.
	DryRun bool

	// The description to display for the new image.
	NewImageDescription *string

	// The name to display for the new image.
	NewImageDisplayName *string

	// The tags to associate with the new image. A tag is a key-value pair, and the
	// value is optional. For example, Environment=Test. If you do not specify a value,
	// Environment=. Generally allowed characters are: letters, numbers, and spaces
	// representable in UTF-8, and the following special characters: _ . : / = + \ - @
	// If you do not specify a value, the value is set to an empty string. For more
	// information about tags, see Tagging Your Resources
	// (https://docs.aws.amazon.com/appstream2/latest/developerguide/tagging-basic.html)
	// in the Amazon AppStream 2.0 Administration Guide.
	NewImageTags map[string]string
}

type CreateUpdatedImageOutput struct {

	// Indicates whether a new image can be created.
	CanUpdateImage bool

	// Describes an image.
	Image *types.Image

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateUpdatedImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUpdatedImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUpdatedImage{}, middleware.After)
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
	if err = addOpCreateUpdatedImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUpdatedImage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateUpdatedImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "CreateUpdatedImage",
	}
}
