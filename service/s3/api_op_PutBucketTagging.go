// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the tags for a bucket. Use tags to organize your Amazon Web Services bill
// to reflect your own cost structure. To do this, sign up to get your Amazon Web
// Services account bill with tag key values included. Then, to see the cost of
// combined resources, organize your billing information according to resources
// with the same tag key values. For example, you can tag several resources with a
// specific application name, and then organize your billing information to see the
// total cost of that application across several services. For more information,
// see Cost Allocation and Tagging
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
// and Using Cost Allocation in Amazon S3 Bucket Tags
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/CostAllocTagging.html). When
// this operation sets the tags for a bucket, it will overwrite any current tags
// the bucket already has. You cannot use this operation to add tags to an existing
// list of tags. To use this operation, you must have permissions to perform the
// s3:PutBucketTagging action. The bucket owner has this permission by default and
// can grant this permission to others. For more information about permissions, see
// Permissions Related to Bucket Subresource Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html).
// PutBucketTagging has the following special errors:
//
// * Error code:
// InvalidTagError
//
// * Description: The tag provided was not a valid tag. This error
// can occur if the tag did not pass input validation. For information about tag
// restrictions, see User-Defined Tag Restrictions
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.html)
// and Amazon Web Services-Generated Cost Allocation Tag Restrictions
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/aws-tag-restrictions.html).
//
// *
// Error code: MalformedXMLError
//
// * Description: The XML provided does not match
// the schema.
//
// * Error code: OperationAbortedError
//
// * Description: A conflicting
// conditional action is currently in progress against this resource. Please try
// again.
//
// * Error code: InternalError
//
// * Description: The service was unable to
// apply the provided tag to the bucket.
//
// The following operations are related to
// PutBucketTagging:
//
// * GetBucketTagging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketTagging.html)
//
// *
// DeleteBucketTagging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketTagging.html)
func (c *Client) PutBucketTagging(ctx context.Context, params *PutBucketTaggingInput, optFns ...func(*Options)) (*PutBucketTaggingOutput, error) {
	if params == nil {
		params = &PutBucketTaggingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketTagging", params, optFns, c.addOperationPutBucketTaggingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketTaggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketTaggingInput struct {

	// The bucket name.
	//
	// This member is required.
	Bucket *string

	// Container for the TagSet and Tag elements.
	//
	// This member is required.
	Tagging *types.Tagging

	// The base64-encoded 128-bit MD5 digest of the data. You must use this header as a
	// message integrity check to verify that the request body was not corrupted in
	// transit. For more information, see RFC 1864
	// (http://www.ietf.org/rfc/rfc1864.txt). For requests made using the Amazon Web
	// Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is
	// calculated automatically.
	ContentMD5 *string

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

type PutBucketTaggingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBucketTaggingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketTagging{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketTagging{}, middleware.After)
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
	if err = swapWithCustomHTTPSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutBucketTaggingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketTagging(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addPutBucketTaggingUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutBucketTagging(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketTagging",
	}
}

// getPutBucketTaggingBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getPutBucketTaggingBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketTaggingInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutBucketTaggingUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutBucketTaggingBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseDualstack:                   options.UseDualstack,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
