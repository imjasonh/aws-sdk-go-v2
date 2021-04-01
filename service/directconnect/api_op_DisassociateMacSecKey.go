// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the association between a MAC Security (MACsec) security key and an AWS
// Direct Connect dedicated connection.
func (c *Client) DisassociateMacSecKey(ctx context.Context, params *DisassociateMacSecKeyInput, optFns ...func(*Options)) (*DisassociateMacSecKeyOutput, error) {
	if params == nil {
		params = &DisassociateMacSecKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateMacSecKey", params, optFns, addOperationDisassociateMacSecKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateMacSecKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateMacSecKeyInput struct {

	// The ID of the dedicated connection (dxcon-xxxx), or the ID of the LAG
	// (dxlag-xxxx). You can use DescribeConnections or DescribeLags to retrieve
	// connection ID.
	//
	// This member is required.
	ConnectionId *string

	// The Amazon Resource Name (ARN) of the MAC Security (MACsec) secret key. You can
	// use DescribeConnections to retrieve the ARN of the MAC Security (MACsec) secret
	// key.
	//
	// This member is required.
	SecretARN *string
}

type DisassociateMacSecKeyOutput struct {

	// The ID of the dedicated connection (dxcon-xxxx), or the ID of the LAG
	// (dxlag-xxxx).
	ConnectionId *string

	// The MAC Security (MACsec) security keys no longer associated with the dedicated
	// connection.
	MacSecKeys []types.MacSecKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateMacSecKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateMacSecKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateMacSecKey{}, middleware.After)
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
	if err = addOpDisassociateMacSecKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateMacSecKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateMacSecKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DisassociateMacSecKey",
	}
}
