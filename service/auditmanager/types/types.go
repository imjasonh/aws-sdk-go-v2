// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// An entity that defines the scope of audit evidence collected by AWS Audit
// Manager. An AWS Audit Manager assessment is an implementation of an AWS Audit
// Manager framework.
type Assessment struct {

	// The Amazon Resource Name (ARN) of the assessment.
	Arn *string

	// The AWS account associated with the assessment.
	AwsAccount *AWSAccount

	// The framework from which the assessment was created.
	Framework *AssessmentFramework

	// The metadata for the specified assessment.
	Metadata *AssessmentMetadata

	// The tags associated with the assessment.
	Tags map[string]string
}

// The control entity that represents a standard or custom control used in an AWS
// Audit Manager assessment.
type AssessmentControl struct {

	// The amount of evidence in the assessment report.
	AssessmentReportEvidenceCount int32

	// The list of comments attached to the specified control.
	Comments []ControlComment

	// The description of the specified control.
	Description *string

	// The amount of evidence generated for the control.
	EvidenceCount int32

	// The list of data sources for the specified evidence.
	EvidenceSources []string

	// The identifier for the specified control.
	Id *string

	// The name of the specified control.
	Name *string

	// The response of the specified control.
	Response ControlResponse

	// The status of the specified control.
	Status ControlStatus
}

// Represents a set of controls in an AWS Audit Manager assessment.
type AssessmentControlSet struct {

	// The list of controls contained with the control set.
	Controls []AssessmentControl

	// The delegations associated with the control set.
	Delegations []Delegation

	// The description for the control set.
	Description *string

	// The identifier of the control set in the assessment. This is the control set
	// name in a plain string format.
	Id *string

	// The total number of evidence objects uploaded manually to the control set.
	ManualEvidenceCount int32

	// The roles associated with the control set.
	Roles []Role

	// Specifies the current status of the control set.
	Status ControlSetStatus

	// The total number of evidence objects retrieved automatically for the control
	// set.
	SystemEvidenceCount int32
}

// The folder in which AWS Audit Manager stores evidence for an assessment.
type AssessmentEvidenceFolder struct {

	// The identifier for the specified assessment.
	AssessmentId *string

	// The total count of evidence included in the assessment report.
	AssessmentReportSelectionCount int32

	// The name of the user who created the evidence folder.
	Author *string

	// The unique identifier for the specified control.
	ControlId *string

	// The name of the control.
	ControlName *string

	// The identifier for the control set.
	ControlSetId *string

	// The AWS service from which the evidence was collected.
	DataSource *string

	// The date when the first evidence was added to the evidence folder.
	Date *time.Time

	// The total number of AWS resources assessed to generate the evidence.
	EvidenceAwsServiceSourceCount int32

	// The number of evidence that falls under the compliance check category. This
	// evidence is collected from AWS Config or AWS Security Hub.
	EvidenceByTypeComplianceCheckCount int32

	// The total number of issues that were reported directly from AWS Security Hub,
	// AWS Config, or both.
	EvidenceByTypeComplianceCheckIssuesCount int32

	// The number of evidence that falls under the configuration data category. This
	// evidence is collected from configuration snapshots of other AWS services such as
	// Amazon EC2, Amazon S3, or IAM.
	EvidenceByTypeConfigurationDataCount int32

	// The number of evidence that falls under the manual category. This evidence is
	// imported manually.
	EvidenceByTypeManualCount int32

	// The number of evidence that falls under the user activity category. This
	// evidence is collected from AWS CloudTrail logs.
	EvidenceByTypeUserActivityCount int32

	// The amount of evidence included in the evidence folder.
	EvidenceResourcesIncludedCount int32

	// The identifier for the folder in which evidence is stored.
	Id *string

	// The name of the specified evidence folder.
	Name *string

	// The total amount of evidence in the evidence folder.
	TotalEvidence int32
}

// The file used to structure and automate AWS Audit Manager assessments for a
// given compliance standard.
type AssessmentFramework struct {

	// The Amazon Resource Name (ARN) of the specified framework.
	Arn *string

	// The control sets associated with the framework.
	ControlSets []AssessmentControlSet

	// The unique identifier for the framework.
	Id *string

	// The metadata of a framework, such as the name, ID, description, and so on.
	Metadata *FrameworkMetadata
}

// The metadata associated with a standard or custom framework.
type AssessmentFrameworkMetadata struct {

	// The Amazon Resource Name (ARN) of the framework.
	Arn *string

	// The compliance type that the new custom framework supports, such as CIS or
	// HIPAA.
	ComplianceType *string

	// The number of control sets associated with the specified framework.
	ControlSetsCount int32

	// The number of controls associated with the specified framework.
	ControlsCount int32

	// Specifies when the framework was created.
	CreatedAt *time.Time

	// The description of the specified framework.
	Description *string

	// The unique identified for the specified framework.
	Id *string

	// Specifies when the framework was most recently updated.
	LastUpdatedAt *time.Time

	// The logo associated with the framework.
	Logo *string

	// The name of the specified framework.
	Name *string

	// The framework type, such as standard or custom.
	Type FrameworkType
}

// The metadata associated with the specified assessment.
type AssessmentMetadata struct {

	// The destination in which evidence reports are stored for the specified
	// assessment.
	AssessmentReportsDestination *AssessmentReportsDestination

	// The name of a compliance standard related to the assessment, such as PCI-DSS.
	ComplianceType *string

	// Specifies when the assessment was created.
	CreationTime *time.Time

	// The delegations associated with the assessment.
	Delegations []Delegation

	// The description of the assessment.
	Description *string

	// The unique identifier for the assessment.
	Id *string

	// The time of the most recent update.
	LastUpdated *time.Time

	// The name of the assessment.
	Name *string

	// The roles associated with the assessment.
	Roles []Role

	// The wrapper of AWS accounts and services in scope for the assessment.
	Scope *Scope

	// The overall status of the assessment.
	Status AssessmentStatus
}

// A metadata object associated with an assessment in AWS Audit Manager.
type AssessmentMetadataItem struct {

	// The name of the compliance standard related to the assessment, such as PCI-DSS.
	ComplianceType *string

	// Specifies when the assessment was created.
	CreationTime *time.Time

	// The delegations associated with the assessment.
	Delegations []Delegation

	// The unique identifier for the assessment.
	Id *string

	// The time of the most recent update.
	LastUpdated *time.Time

	// The name of the assessment.
	Name *string

	// The roles associated with the assessment.
	Roles []Role

	// The current status of the assessment.
	Status AssessmentStatus
}

// A finalized document generated from an AWS Audit Manager assessment. These
// reports summarize the relevant evidence collected for your audit, and link to
// the relevant evidence folders which are named and organized according to the
// controls specified in your assessment.
type AssessmentReport struct {

	// The identifier for the specified assessment.
	AssessmentId *string

	// The name of the associated assessment.
	AssessmentName *string

	// The name of the user who created the assessment report.
	Author *string

	// The identifier for the specified AWS account.
	AwsAccountId *string

	// Specifies when the assessment report was created.
	CreationTime *time.Time

	// The description of the specified assessment report.
	Description *string

	// The unique identifier for the specified assessment report.
	Id *string

	// The name given to the assessment report.
	Name *string

	// The current status of the specified assessment report.
	Status AssessmentReportStatus
}

// An error entity for the AssessmentReportEvidence API. This is used to provide
// more meaningful errors than a simple string message.
type AssessmentReportEvidenceError struct {

	// The error code returned by the AssessmentReportEvidence API.
	ErrorCode *string

	// The error message returned by the AssessmentReportEvidence API.
	ErrorMessage *string

	// The identifier for the evidence.
	EvidenceId *string
}

// The metadata objects associated with the specified assessment report.
type AssessmentReportMetadata struct {

	// The unique identifier for the associated assessment.
	AssessmentId *string

	// The name of the associated assessment.
	AssessmentName *string

	// The name of the user who created the assessment report.
	Author *string

	// Specifies when the assessment report was created.
	CreationTime *time.Time

	// The description of the specified assessment report.
	Description *string

	// The unique identifier for the assessment report.
	Id *string

	// The name of the assessment report.
	Name *string

	// The current status of the assessment report.
	Status AssessmentReportStatus
}

// The location in which AWS Audit Manager saves assessment reports for the given
// assessment.
type AssessmentReportsDestination struct {

	// The destination of the assessment report.
	Destination *string

	// The destination type, such as Amazon S3.
	DestinationType AssessmentReportDestinationType
}

// The wrapper of AWS account details, such as account ID, email address, and so
// on.
type AWSAccount struct {

	// The email address associated with the specified AWS account.
	EmailAddress *string

	// The identifier for the specified AWS account.
	Id *string

	// The name of the specified AWS account.
	Name *string
}

// An AWS service such as Amazon S3, AWS CloudTrail, and so on.
type AWSService struct {

	// The name of the AWS service.
	ServiceName *string
}

// An error entity for the BatchCreateDelegationByAssessment API. This is used to
// provide more meaningful errors than a simple string message.
type BatchCreateDelegationByAssessmentError struct {

	// The API request to batch create delegations in AWS Audit Manager.
	CreateDelegationRequest *CreateDelegationRequest

	// The error code returned by the BatchCreateDelegationByAssessment API.
	ErrorCode *string

	// The error message returned by the BatchCreateDelegationByAssessment API.
	ErrorMessage *string
}

// An error entity for the BatchDeleteDelegationByAssessment API. This is used to
// provide more meaningful errors than a simple string message.
type BatchDeleteDelegationByAssessmentError struct {

	// The identifier for the specified delegation.
	DelegationId *string

	// The error code returned by the BatchDeleteDelegationByAssessment API.
	ErrorCode *string

	// The error message returned by the BatchDeleteDelegationByAssessment API.
	ErrorMessage *string
}

// An error entity for the BatchImportEvidenceToAssessmentControl API. This is used
// to provide more meaningful errors than a simple string message.
type BatchImportEvidenceToAssessmentControlError struct {

	// The error code returned by the BatchImportEvidenceToAssessmentControl API.
	ErrorCode *string

	// The error message returned by the BatchImportEvidenceToAssessmentControl API.
	ErrorMessage *string

	// Manual evidence that cannot be collected automatically by AWS Audit Manager.
	ManualEvidence *ManualEvidence
}

// The record of a change within AWS Audit Manager, such as a modified assessment,
// a delegated control set, and so on.
type ChangeLog struct {

	// The action performed.
	Action ActionEnum

	// The time of creation for the changelog object.
	CreatedAt *time.Time

	// The IAM user or role that performed the action.
	CreatedBy *string

	// The name of the changelog object.
	ObjectName *string

	// The changelog object type, such as an assessment, control, or control set.
	ObjectType ObjectTypeEnum
}

// A control in AWS Audit Manager.
type Control struct {

	// The recommended actions to carry out if the control is not fulfilled.
	ActionPlanInstructions *string

	// The title of the action plan for remediating the control.
	ActionPlanTitle *string

	// The Amazon Resource Name (ARN) of the specified control.
	Arn *string

	// The data mapping sources for the specified control.
	ControlMappingSources []ControlMappingSource

	// The data source that determines from where AWS Audit Manager collects evidence
	// for the control.
	ControlSources *string

	// Specifies when the control was created.
	CreatedAt *time.Time

	// The IAM user or role that created the control.
	CreatedBy *string

	// The description of the specified control.
	Description *string

	// The unique identifier for the control.
	Id *string

	// Specifies when the control was most recently updated.
	LastUpdatedAt *time.Time

	// The IAM user or role that most recently updated the control.
	LastUpdatedBy *string

	// The name of the specified control.
	Name *string

	// The tags associated with the control.
	Tags map[string]string

	// The steps to follow to determine if the control has been satisfied.
	TestingInformation *string

	// The type of control, such as custom or standard.
	Type ControlType
}

// A comment posted by a user on a control. This includes the author's name, the
// comment text, and a timestamp.
type ControlComment struct {

	// The name of the user who authored the comment.
	AuthorName *string

	// The body text of a control comment.
	CommentBody *string

	// The time when the comment was posted.
	PostedDate *time.Time
}

// The data source that determines from where AWS Audit Manager collects evidence
// for the control.
type ControlMappingSource struct {

	// The description of the specified source.
	SourceDescription *string

	// The frequency of evidence collection for the specified control mapping source.
	SourceFrequency SourceFrequency

	// The unique identifier for the specified source.
	SourceId *string

	// The keyword to search for in AWS CloudTrail logs, AWS Config rules, AWS Security
	// Hub checks, and AWS API names.
	SourceKeyword *SourceKeyword

	// The name of the specified source.
	SourceName *string

	// The setup option for the data source, which reflects if the evidence collection
	// is automated or manual.
	SourceSetUpOption SourceSetUpOption

	// Specifies one of the five types of data sources for evidence collection.
	SourceType SourceType

	// The instructions for troubleshooting the specified control.
	TroubleshootingText *string
}

// The metadata associated with the specified standard or custom control.
type ControlMetadata struct {

	// The Amazon Resource Name (ARN) of the specified control.
	Arn *string

	// The data source that determines from where AWS Audit Manager collects evidence
	// for the control.
	ControlSources *string

	// Specifies when the control was created.
	CreatedAt *time.Time

	// The unique identifier for the specified control.
	Id *string

	// Specifies when the control was most recently updated.
	LastUpdatedAt *time.Time

	// The name of the specified control.
	Name *string
}

// A set of controls in AWS Audit Manager.
type ControlSet struct {

	// The list of controls within the control set.
	Controls []Control

	// The identifier of the control set in the assessment. This is the control set
	// name in a plain string format.
	Id *string

	// The name of the control set.
	Name *string
}

// Control entity attributes that uniquely identify an existing control to be added
// to a framework in AWS Audit Manager.
type CreateAssessmentFrameworkControl struct {

	// The unique identifier of the control.
	Id *string
}

// A controlSet entity that represents a collection of controls in AWS Audit
// Manager. This does not contain the control set ID.
type CreateAssessmentFrameworkControlSet struct {

	// The list of controls within the control set. This does not contain the control
	// set ID.
	Controls []CreateAssessmentFrameworkControl

	// The name of the specified control set.
	Name *string
}

// Control mapping fields that represent the source for evidence collection, along
// with related parameters and metadata. This does not contain mappingID.
type CreateControlMappingSource struct {

	// The description of the data source that determines from where AWS Audit Manager
	// collects evidence for the control.
	SourceDescription *string

	// The frequency of evidence collection for the specified control mapping source.
	SourceFrequency SourceFrequency

	// The keyword to search for in AWS CloudTrail logs, AWS Config rules, AWS Security
	// Hub checks, and AWS API names.
	SourceKeyword *SourceKeyword

	// The name of the control mapping data source.
	SourceName *string

	// The setup option for the data source, which reflects if the evidence collection
	// is automated or manual.
	SourceSetUpOption SourceSetUpOption

	// Specifies one of the five types of data sources for evidence collection.
	SourceType SourceType

	// The instructions for troubleshooting the specified control.
	TroubleshootingText *string
}

// A collection of attributes used to create a delegation for an assessment in AWS
// Audit Manager.
type CreateDelegationRequest struct {

	// A comment related to the delegation request.
	Comment *string

	// The unique identifier for the control set.
	ControlSetId *string

	// The Amazon Resource Name (ARN) of the IAM role.
	RoleArn *string

	// The type of customer persona. In CreateAssessment, roleType can only be
	// PROCESS_OWNER. In UpdateSettings, roleType can only be PROCESS_OWNER. In
	// BatchCreateDelegationByAssessment, roleType can only be RESOURCE_OWNER.
	RoleType RoleType
}

// The assignment of a control set to a delegate for review.
type Delegation struct {

	// The identifier for the associated assessment.
	AssessmentId *string

	// The name of the associated assessment.
	AssessmentName *string

	// The comment related to the delegation.
	Comment *string

	// The identifier for the associated control set.
	ControlSetId *string

	// The IAM user or role that created the delegation.
	CreatedBy *string

	// Specifies when the delegation was created.
	CreationTime *time.Time

	// The unique identifier for the delegation.
	Id *string

	// Specifies when the delegation was last updated.
	LastUpdated *time.Time

	// The Amazon Resource Name (ARN) of the IAM role.
	RoleArn *string

	// The type of customer persona. In CreateAssessment, roleType can only be
	// PROCESS_OWNER. In UpdateSettings, roleType can only be PROCESS_OWNER. In
	// BatchCreateDelegationByAssessment, roleType can only be RESOURCE_OWNER.
	RoleType RoleType

	// The status of the delegation.
	Status DelegationStatus
}

// The metadata associated with the specified delegation.
type DelegationMetadata struct {

	// The unique identifier for the specified assessment.
	AssessmentId *string

	// The name of the associated assessment.
	AssessmentName *string

	// Specifies the name of the control set delegated for review.
	ControlSetName *string

	// Specifies when the delegation was created.
	CreationTime *time.Time

	// The unique identifier for the delegation.
	Id *string

	// The Amazon Resource Name (ARN) of the IAM role.
	RoleArn *string

	// The current status of the delgation.
	Status DelegationStatus
}

// A record that contains the information needed to demonstrate compliance with the
// requirements specified by a control. Examples of evidence include change
// activity triggered by a user, or a system configuration snapshot.
type Evidence struct {

	// Specifies whether the evidence is inclded in the assessment report.
	AssessmentReportSelection *string

	// The names and values used by the evidence event, including an attribute name
	// (such as allowUsersToChangePassword) and value (such as true or false).
	Attributes map[string]string

	// The identifier for the specified AWS account.
	AwsAccountId *string

	// The AWS account from which the evidence is collected, and its AWS organization
	// path.
	AwsOrganization *string

	// The evaluation status for evidence that falls under the compliance check
	// category. For evidence collected from AWS Security Hub, a Pass or Fail result is
	// shown. For evidence collected from AWS Config, a Compliant or Noncompliant
	// result is shown.
	ComplianceCheck *string

	// The data source from which the specified evidence was collected.
	DataSource *string

	// The name of the specified evidence event.
	EventName *string

	// The AWS service from which the evidence is collected.
	EventSource *string

	// The identifier for the specified AWS account.
	EvidenceAwsAccountId *string

	// The type of automated evidence.
	EvidenceByType *string

	// The identifier for the folder in which the evidence is stored.
	EvidenceFolderId *string

	// The unique identifier for the IAM user or role associated with the evidence.
	IamId *string

	// The identifier for the evidence.
	Id *string

	// The list of resources assessed to generate the evidence.
	ResourcesIncluded []Resource

	// The timestamp that represents when the evidence was collected.
	Time *time.Time
}

// The file used to structure and automate AWS Audit Manager assessments for a
// given compliance standard.
type Framework struct {

	// The Amazon Resource Name (ARN) of the specified framework.
	Arn *string

	// The compliance type that the new custom framework supports, such as CIS or
	// HIPAA.
	ComplianceType *string

	// The control sets associated with the framework.
	ControlSets []ControlSet

	// The sources from which AWS Audit Manager collects evidence for the control.
	ControlSources *string

	// Specifies when the framework was created.
	CreatedAt *time.Time

	// The IAM user or role that created the framework.
	CreatedBy *string

	// The description of the specified framework.
	Description *string

	// The unique identifier for the specified framework.
	Id *string

	// Specifies when the framework was most recently updated.
	LastUpdatedAt *time.Time

	// The IAM user or role that most recently updated the framework.
	LastUpdatedBy *string

	// The logo associated with the framework.
	Logo *string

	// The name of the specified framework.
	Name *string

	// The tags associated with the framework.
	Tags map[string]string

	// The framework type, such as custom or standard.
	Type FrameworkType
}

// The metadata of a framework, such as the name, ID, description, and so on.
type FrameworkMetadata struct {

	// The compliance standard associated with the framework, such as PCI-DSS or HIPAA.
	ComplianceType *string

	// The description of the framework.
	Description *string

	// The logo associated with the framework.
	Logo *string

	// The name of the framework.
	Name *string
}

// Evidence that is uploaded to AWS Audit Manager manually.
type ManualEvidence struct {

	// The Amazon S3 URL that points to a manual evidence object.
	S3ResourcePath *string
}

// The notification used to inform a user of an update in AWS Audit Manager. For
// example, this includes the notification that is sent when a control set is
// delegated for review.
type Notification struct {

	// The identifier for the specified assessment.
	AssessmentId *string

	// The name of the related assessment.
	AssessmentName *string

	// The identifier for the specified control set.
	ControlSetId *string

	// Specifies the name of the control set that the notification is about.
	ControlSetName *string

	// The description of the notification.
	Description *string

	// The time when the notification was sent.
	EventTime *time.Time

	// The unique identifier for the notification.
	Id *string

	// The sender of the notification.
	Source *string
}

// A system asset that is evaluated in an AWS Audit Manager assessment.
type Resource struct {

	// The Amazon Resource Name (ARN) for the specified resource.
	Arn *string

	// The value of the specified resource.
	Value *string
}

// The wrapper that contains the AWS Audit Manager role information of the current
// user, such as the role type and IAM Amazon Resource Name (ARN).
type Role struct {

	// The Amazon Resource Name (ARN) of the IAM role.
	RoleArn *string

	// The type of customer persona. In CreateAssessment, roleType can only be
	// PROCESS_OWNER. In UpdateSettings, roleType can only be PROCESS_OWNER. In
	// BatchCreateDelegationByAssessment, roleType can only be RESOURCE_OWNER.
	RoleType RoleType
}

// The wrapper that contains the AWS accounts and AWS services in scope for the
// assessment.
type Scope struct {

	// The AWS accounts included in the scope of the assessment.
	AwsAccounts []AWSAccount

	// The AWS services included in the scope of the assessment.
	AwsServices []AWSService
}

// The metadata associated with the specified AWS service.
type ServiceMetadata struct {

	// The category in which the AWS service belongs, such as compute, storage,
	// database, and so on.
	Category *string

	// The description of the specified AWS service.
	Description *string

	// The display name of the AWS service.
	DisplayName *string

	// The name of the AWS service.
	Name *string
}

// The settings object that holds all supported AWS Audit Manager settings.
type Settings struct {

	// The default storage destination for assessment reports.
	DefaultAssessmentReportsDestination *AssessmentReportsDestination

	// The designated default audit owners.
	DefaultProcessOwners []Role

	// Specifies whether AWS Organizations is enabled.
	IsAwsOrgEnabled *bool

	// The AWS KMS key details.
	KmsKey *string

	// The designated Amazon Simple Notification Service (Amazon SNS) topic.
	SnsTopic *string
}

// The keyword to search for in AWS CloudTrail logs, AWS Config rules, AWS Security
// Hub checks, and AWS API names.
type SourceKeyword struct {

	// The method of input for the specified keyword.
	KeywordInputType KeywordInputType

	// The value of the keyword used to search AWS CloudTrail logs, AWS Config rules,
	// AWS Security Hub checks, and AWS API names when mapping a control data source.
	KeywordValue *string
}

// A controlSet entity that represents a collection of controls in AWS Audit
// Manager. This does not contain the control set ID.
type UpdateAssessmentFrameworkControlSet struct {

	// The list of controls contained within the control set.
	Controls []CreateAssessmentFrameworkControl

	// The unique identifier for the control set.
	Id *string

	// The name of the control set.
	Name *string
}

// A uniform resource locator, used as a unique identifier to locate a resource on
// the internet.
type URL struct {

	// The name or word used as a hyperlink to the URL.
	HyperlinkName *string

	// The unique identifier for the internet resource.
	Link *string
}

// Indicates that the request has invalid or missing parameters for the specified
// field.
type ValidationExceptionField struct {

	// The body of the error message.
	//
	// This member is required.
	Message *string

	// The name of the validation error.
	//
	// This member is required.
	Name *string
}
