// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// One or more parameters are not valid.
type DirectConnectClientException struct {
	Message *string
}

func (e *DirectConnectClientException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectConnectClientException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectConnectClientException) ErrorCode() string             { return "DirectConnectClientException" }
func (e *DirectConnectClientException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A server-side error occurred.
type DirectConnectServerException struct {
	Message *string
}

func (e *DirectConnectServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectConnectServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectConnectServerException) ErrorCode() string             { return "DirectConnectServerException" }
func (e *DirectConnectServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// A tag key was specified more than once.
type DuplicateTagKeysException struct {
	Message *string
}

func (e *DuplicateTagKeysException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateTagKeysException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateTagKeysException) ErrorCode() string             { return "DuplicateTagKeysException" }
func (e *DuplicateTagKeysException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have reached the limit on the number of tags that can be assigned.
type TooManyTagsException struct {
	Message *string
}

func (e *TooManyTagsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTagsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTagsException) ErrorCode() string             { return "TooManyTagsException" }
func (e *TooManyTagsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
