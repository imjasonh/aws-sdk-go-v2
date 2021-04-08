// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows the user to set the SourceServer.LifeCycle.state property for specific
// Source Server IDs to one of the following: READY_FOR_TEST or READY_FOR_CUTOVER.
// This command only works if the Source Server is already launchable
// (dataReplicationInfo.lagDuration is not null.)
func (c *Client) ChangeServerLifeCycleState(ctx context.Context, params *ChangeServerLifeCycleStateInput, optFns ...func(*Options)) (*ChangeServerLifeCycleStateOutput, error) {
	if params == nil {
		params = &ChangeServerLifeCycleStateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ChangeServerLifeCycleState", params, optFns, addOperationChangeServerLifeCycleStateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ChangeServerLifeCycleStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ChangeServerLifeCycleStateInput struct {

	// The request to change the source server migration lifecycle state.
	//
	// This member is required.
	LifeCycle *types.ChangeServerLifeCycleStateSourceServerLifecycle

	// The request to change the source server migration lifecycle state by source
	// server ID.
	//
	// This member is required.
	SourceServerID *string
}

type ChangeServerLifeCycleStateOutput struct {

	// Source server ARN.
	Arn *string

	// Source server data replication info.
	DataReplicationInfo *types.DataReplicationInfo

	// Source server archived status.
	IsArchived *bool

	// Source server launched instance.
	LaunchedInstance *types.LaunchedInstance

	// Source server lifecycle state.
	LifeCycle *types.LifeCycle

	// Source server properties.
	SourceProperties *types.SourceProperties

	// Source server ID.
	SourceServerID *string

	// Source server Tags.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationChangeServerLifeCycleStateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpChangeServerLifeCycleState{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpChangeServerLifeCycleState{}, middleware.After)
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
	if err = addOpChangeServerLifeCycleStateValidationMiddleware(stack); err != nil {
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
