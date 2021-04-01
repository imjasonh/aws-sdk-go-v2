// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the object definitions for a set of objects associated with the pipeline.
// Object definitions are composed of a set of fields that define the properties of
// the object. POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1
// X-Amz-Target: DataPipeline.DescribeObjects Content-Length: 98 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams {"pipelineId": "df-06372391ZG65EXAMPLE", "objectIds":
// ["Schedule"], "evaluateExpressions": true} x-amzn-RequestId:
// 4c18ea5d-0777-11e2-8a14-21bb8a1f50ef Content-Type: application/x-amz-json-1.1
// Content-Length: 1488 Date: Mon, 12 Nov 2012 17:50:53 GMT {"hasMoreResults":
// false, "pipelineObjects": [ {"fields": [ {"key": "startDateTime", "stringValue":
// "2012-12-12T00:00:00"}, {"key": "parent", "refValue": "Default"}, {"key":
// "@sphere", "stringValue": "COMPONENT"}, {"key": "type", "stringValue":
// "Schedule"}, {"key": "period", "stringValue": "1 hour"}, {"key": "endDateTime",
// "stringValue": "2012-12-21T18:00:00"}, {"key": "@version", "stringValue": "1"},
// {"key": "@status", "stringValue": "PENDING"}, {"key": "@pipelineId",
// "stringValue": "df-06372391ZG65EXAMPLE"} ], "id": "Schedule", "name":
// "Schedule"} ] }
func (c *Client) DescribeObjects(ctx context.Context, params *DescribeObjectsInput, optFns ...func(*Options)) (*DescribeObjectsOutput, error) {
	if params == nil {
		params = &DescribeObjectsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeObjects", params, optFns, addOperationDescribeObjectsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeObjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeObjects.
type DescribeObjectsInput struct {

	// The IDs of the pipeline objects that contain the definitions to be described.
	// You can pass as many as 25 identifiers in a single call to DescribeObjects.
	//
	// This member is required.
	ObjectIds []string

	// The ID of the pipeline that contains the object definitions.
	//
	// This member is required.
	PipelineId *string

	// Indicates whether any expressions in the object should be evaluated when the
	// object descriptions are returned.
	EvaluateExpressions bool

	// The starting point for the results to be returned. For the first call, this
	// value should be empty. As long as there are more results, continue to call
	// DescribeObjects with the marker value from the previous call to retrieve the
	// next set of results.
	Marker *string
}

// Contains the output of DescribeObjects.
type DescribeObjectsOutput struct {

	// An array of object definitions.
	//
	// This member is required.
	PipelineObjects []types.PipelineObject

	// Indicates whether there are more results to return.
	HasMoreResults bool

	// The starting point for the next page of results. To view the next page of
	// results, call DescribeObjects again with this marker value. If the value is
	// null, there are no more results.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeObjectsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeObjects{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeObjects{}, middleware.After)
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
	if err = addOpDescribeObjectsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeObjects(options.Region), middleware.Before); err != nil {
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

// DescribeObjectsAPIClient is a client that implements the DescribeObjects
// operation.
type DescribeObjectsAPIClient interface {
	DescribeObjects(context.Context, *DescribeObjectsInput, ...func(*Options)) (*DescribeObjectsOutput, error)
}

var _ DescribeObjectsAPIClient = (*Client)(nil)

// DescribeObjectsPaginatorOptions is the paginator options for DescribeObjects
type DescribeObjectsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeObjectsPaginator is a paginator for DescribeObjects
type DescribeObjectsPaginator struct {
	options   DescribeObjectsPaginatorOptions
	client    DescribeObjectsAPIClient
	params    *DescribeObjectsInput
	nextToken *string
	firstPage bool
}

// NewDescribeObjectsPaginator returns a new DescribeObjectsPaginator
func NewDescribeObjectsPaginator(client DescribeObjectsAPIClient, params *DescribeObjectsInput, optFns ...func(*DescribeObjectsPaginatorOptions)) *DescribeObjectsPaginator {
	if params == nil {
		params = &DescribeObjectsInput{}
	}

	options := DescribeObjectsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeObjectsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeObjectsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeObjects page.
func (p *DescribeObjectsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeObjectsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	result, err := p.client.DescribeObjects(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeObjects(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "DescribeObjects",
	}
}
