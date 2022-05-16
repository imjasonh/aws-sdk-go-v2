// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists agents or connectors as specified by ID or other filters. All
// agents/connectors associated with your user account can be listed if you call
// DescribeAgents as is without passing any parameters.
func (c *Client) DescribeAgents(ctx context.Context, params *DescribeAgentsInput, optFns ...func(*Options)) (*DescribeAgentsOutput, error) {
	if params == nil {
		params = &DescribeAgentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAgents", params, optFns, c.addOperationDescribeAgentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAgentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAgentsInput struct {

	// The agent or the Connector IDs for which you want information. If you specify no
	// IDs, the system returns information about all agents/Connectors associated with
	// your Amazon Web Services user account.
	AgentIds []string

	// You can filter the request using various logical operators and a key-value
	// format. For example: {"key": "collectionStatus", "value": "STARTED"}
	Filters []types.Filter

	// The total number of agents/Connectors to return in a single page of output. The
	// maximum value is 100.
	MaxResults int32

	// Token to retrieve the next set of results. For example, if you previously
	// specified 100 IDs for DescribeAgentsRequest$agentIds but set
	// DescribeAgentsRequest$maxResults to 10, you received a set of 10 results along
	// with a token. Use that token in this query to get the next set of 10.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeAgentsOutput struct {

	// Lists agents or the Connector by ID or lists all agents/Connectors associated
	// with your user account if you did not specify an agent/Connector ID. The output
	// includes agent/Connector IDs, IP addresses, media access control (MAC)
	// addresses, agent/Connector health, host name where the agent/Connector resides,
	// and the version number of each agent/Connector.
	AgentsInfo []types.AgentInfo

	// Token to retrieve the next set of results. For example, if you specified 100 IDs
	// for DescribeAgentsRequest$agentIds but set DescribeAgentsRequest$maxResults to
	// 10, you received a set of 10 results along with this token. Use this token in
	// the next query to retrieve the next set of 10.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAgentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAgents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAgents{}, middleware.After)
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
	if err = addOpDescribeAgentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAgents(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAgents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "DescribeAgents",
	}
}
