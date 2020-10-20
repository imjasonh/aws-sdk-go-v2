// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// The example tests how requests serialize different timestamp formats in the URI
// path.
func (c *Client) HttpRequestWithLabelsAndTimestampFormat(ctx context.Context, params *HttpRequestWithLabelsAndTimestampFormatInput, optFns ...func(*Options)) (*HttpRequestWithLabelsAndTimestampFormatOutput, error) {
	if params == nil {
		params = &HttpRequestWithLabelsAndTimestampFormatInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "HttpRequestWithLabelsAndTimestampFormat", params, optFns, addOperationHttpRequestWithLabelsAndTimestampFormatMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*HttpRequestWithLabelsAndTimestampFormatOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpRequestWithLabelsAndTimestampFormatInput struct {
	DefaultFormat *time.Time

	MemberDateTime *time.Time

	MemberEpochSeconds *time.Time

	MemberHttpDate *time.Time

	TargetDateTime *time.Time

	TargetEpochSeconds *time.Time

	TargetHttpDate *time.Time
}

type HttpRequestWithLabelsAndTimestampFormatOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationHttpRequestWithLabelsAndTimestampFormatMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpHttpRequestWithLabelsAndTimestampFormat{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpHttpRequestWithLabelsAndTimestampFormat{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpHttpRequestWithLabelsAndTimestampFormatValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opHttpRequestWithLabelsAndTimestampFormat(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opHttpRequestWithLabelsAndTimestampFormat(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpRequestWithLabelsAndTimestampFormat",
	}
}