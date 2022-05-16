// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an application.
func (c *Client) UpdateApp(ctx context.Context, params *UpdateAppInput, optFns ...func(*Options)) (*UpdateAppOutput, error) {
	if params == nil {
		params = &UpdateAppInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateApp", params, optFns, c.addOperationUpdateAppMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAppOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAppInput struct {

	// The Amazon Resource Name (ARN) of the application. The format for this ARN is:
	// arn:partition:resiliencehub:region:account:app/app-id. For more information
	// about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	//
	// This member is required.
	AppArn *string

	// Assessment execution schedule with 'Daily' or 'Disabled' values.
	AssessmentSchedule types.AppAssessmentScheduleType

	// Specifies if the resiliency policy ARN should be cleared.
	ClearResiliencyPolicyArn *bool

	// The optional description for an app.
	Description *string

	// The Amazon Resource Name (ARN) of the resiliency policy. The format for this ARN
	// is: arn:partition:resiliencehub:region:account:resiliency-policy/policy-id. For
	// more information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	PolicyArn *string

	noSmithyDocumentSerde
}

type UpdateAppOutput struct {

	// The specified application, returned as an object with details including
	// compliance status, creation time, description, resiliency score, and more.
	//
	// This member is required.
	App *types.App

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAppMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateApp{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateApp{}, middleware.After)
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
	if err = addOpUpdateAppValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApp(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateApp(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resiliencehub",
		OperationName: "UpdateApp",
	}
}
