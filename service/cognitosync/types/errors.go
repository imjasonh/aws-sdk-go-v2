// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// An exception thrown when a bulk publish operation is requested less than 24
// hours after a previous bulk publish operation completed successfully.
type AlreadyStreamedException struct {
	Message *string
}

func (e *AlreadyStreamedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AlreadyStreamedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AlreadyStreamedException) ErrorCode() string             { return "AlreadyStreamedException" }
func (e *AlreadyStreamedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Thrown if there are parallel requests to modify a resource.
type ConcurrentModificationException struct {
	Message *string
}

func (e *ConcurrentModificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModificationException) ErrorCode() string {
	return "ConcurrentModificationException"
}
func (e *ConcurrentModificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An exception thrown when there is an IN_PROGRESS bulk publish operation for the
// given identity pool.
type DuplicateRequestException struct {
	Message *string
}

func (e *DuplicateRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateRequestException) ErrorCode() string             { return "DuplicateRequestException" }
func (e *DuplicateRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates an internal service error.
type InternalErrorException struct {
	Message *string
}

func (e *InternalErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalErrorException) ErrorCode() string             { return "InternalErrorException" }
func (e *InternalErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

type InvalidConfigurationException struct {
	Message *string
}

func (e *InvalidConfigurationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidConfigurationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidConfigurationException) ErrorCode() string             { return "InvalidConfigurationException" }
func (e *InvalidConfigurationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The AWS Lambda function returned invalid output or an exception.
type InvalidLambdaFunctionOutputException struct {
	Message *string
}

func (e *InvalidLambdaFunctionOutputException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidLambdaFunctionOutputException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidLambdaFunctionOutputException) ErrorCode() string {
	return "InvalidLambdaFunctionOutputException"
}
func (e *InvalidLambdaFunctionOutputException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Thrown when a request parameter does not comply with the associated constraints.
type InvalidParameterException struct {
	Message *string
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string             { return "InvalidParameterException" }
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// AWS Lambda throttled your account, please contact AWS Support
type LambdaThrottledException struct {
	Message *string
}

func (e *LambdaThrottledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LambdaThrottledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LambdaThrottledException) ErrorCode() string             { return "LambdaThrottledException" }
func (e *LambdaThrottledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Thrown when the limit on the number of objects or operations has been exceeded.
type LimitExceededException struct {
	Message *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Thrown when a user is not authorized to access the requested resource.
type NotAuthorizedException struct {
	Message *string
}

func (e *NotAuthorizedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotAuthorizedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotAuthorizedException) ErrorCode() string             { return "NotAuthorizedException" }
func (e *NotAuthorizedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Thrown if an update can't be applied because the resource was changed by another
// call and this would result in a conflict.
type ResourceConflictException struct {
	Message *string
}

func (e *ResourceConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceConflictException) ErrorCode() string             { return "ResourceConflictException" }
func (e *ResourceConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Thrown if the resource doesn't exist.
type ResourceNotFoundException struct {
	Message *string
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Thrown if the request is throttled.
type TooManyRequestsException struct {
	Message *string
}

func (e *TooManyRequestsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyRequestsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyRequestsException) ErrorCode() string             { return "TooManyRequestsException" }
func (e *TooManyRequestsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
