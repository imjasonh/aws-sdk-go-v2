// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates an Amazon FSx file system from the specified gateway. After the
// disassociation process finishes, the gateway can no longer access the Amazon FSx
// file system. This operation is only supported in the Amazon FSx file gateway
// type.
func (c *Client) DisassociateFileSystem(ctx context.Context, params *DisassociateFileSystemInput, optFns ...func(*Options)) (*DisassociateFileSystemOutput, error) {
	if params == nil {
		params = &DisassociateFileSystemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateFileSystem", params, optFns, addOperationDisassociateFileSystemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateFileSystemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateFileSystemInput struct {

	// The Amazon Resource Name (ARN) of the file system association to be deleted.
	//
	// This member is required.
	FileSystemAssociationARN *string

	// If this value is set to true, the operation disassociates an Amazon FSx file
	// system immediately. It ends all data uploads to the file system, and the file
	// system association enters the FORCE_DELETING status. If this value is set to
	// false, the Amazon FSx file system does not disassociate until all data is
	// uploaded.
	ForceDelete bool
}

type DisassociateFileSystemOutput struct {

	// The Amazon Resource Name (ARN) of the deleted file system association.
	FileSystemAssociationARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateFileSystemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateFileSystem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateFileSystem{}, middleware.After)
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
	if err = addOpDisassociateFileSystemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateFileSystem(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateFileSystem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DisassociateFileSystem",
	}
}
