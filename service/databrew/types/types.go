// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Represents an individual condition that evaluates to true or false. Conditions
// are used with recipe actions. The action is only performed for column values
// where the condition evaluates to true. If a recipe requires more than one
// condition, then the recipe must specify multiple ConditionExpression elements.
// Each condition is applied to the rows in a dataset first, before the recipe
// action is performed.
type ConditionExpression struct {

	// A specific condition to apply to a recipe action. For more information, see
	// Recipe structure
	// (https://docs.aws.amazon.com/databrew/latest/dg/recipes.html#recipes.structure)
	// in the AWS Glue DataBrew Developer Guide.
	//
	// This member is required.
	Condition *string

	// A column to apply this condition to.
	//
	// This member is required.
	TargetColumn *string

	// A value that the condition must evaluate to for the condition to succeed.
	Value *string
}

// Represents a set of options that define how DataBrew will read a comma-separated
// value (CSV) file when creating a dataset from that file.
type CsvOptions struct {

	// A single character that specifies the delimiter being used in the CSV file.
	Delimiter *string

	// A variable that specifies whether the first row in the file is parsed as the
	// header. If this value is false, column names are auto-generated.
	HeaderRow *bool
}

// Represents a set of options that define how DataBrew will write a
// comma-separated value (CSV) file.
type CsvOutputOptions struct {

	// A single character that specifies the delimiter used to create CSV job output.
	Delimiter *string
}

// Connection information for dataset input files stored in a database.
type DatabaseInputDefinition struct {

	// The table within the target database.
	//
	// This member is required.
	DatabaseTableName *string

	// The AWS Glue Connection that stores the connection information for the target
	// database.
	//
	// This member is required.
	GlueConnectionName *string

	// Represents an Amazon S3 location (bucket name and object key) where DataBrew can
	// read input data, or write output from a job.
	TempDirectory *S3Location
}

// Represents how metadata stored in the AWS Glue Data Catalog is defined in a
// DataBrew dataset.
type DataCatalogInputDefinition struct {

	// The name of a database in the Data Catalog.
	//
	// This member is required.
	DatabaseName *string

	// The name of a database table in the Data Catalog. This table corresponds to a
	// DataBrew dataset.
	//
	// This member is required.
	TableName *string

	// The unique identifier of the AWS account that holds the Data Catalog that stores
	// the data.
	CatalogId *string

	// An Amazon location that AWS Glue Data Catalog can use as a temporary directory.
	TempDirectory *S3Location
}

// Represents a dataset that can be processed by DataBrew.
type Dataset struct {

	// Information on how DataBrew can find the dataset, in either the AWS Glue Data
	// Catalog or Amazon S3.
	//
	// This member is required.
	Input *Input

	// The unique name of the dataset.
	//
	// This member is required.
	Name *string

	// The ID of the AWS account that owns the dataset.
	AccountId *string

	// The date and time that the dataset was created.
	CreateDate *time.Time

	// The Amazon Resource Name (ARN) of the user who created the dataset.
	CreatedBy *string

	// The file format of a dataset that is created from an S3 file or folder.
	Format InputFormat

	// A set of options that define how DataBrew interprets the data in the dataset.
	FormatOptions *FormatOptions

	// The Amazon Resource Name (ARN) of the user who last modified the dataset.
	LastModifiedBy *string

	// The last modification date and time of the dataset.
	LastModifiedDate *time.Time

	// A set of options that defines how DataBrew interprets an S3 path of the dataset.
	PathOptions *PathOptions

	// The unique Amazon Resource Name (ARN) for the dataset.
	ResourceArn *string

	// The location of the data for the dataset, either Amazon S3 or the AWS Glue Data
	// Catalog.
	Source Source

	// Metadata tags that have been applied to the dataset.
	Tags map[string]string
}

// Represents a dataset paramater that defines type and conditions for a parameter
// in the S3 path of the dataset.
type DatasetParameter struct {

	// The name of the parameter that is used in the dataset's S3 path.
	//
	// This member is required.
	Name *string

	// The type of the dataset parameter, can be one of a 'String', 'Number' or
	// 'Datetime'.
	//
	// This member is required.
	Type ParameterType

	// Optional boolean value that defines whether the captured value of this parameter
	// should be loaded as an additional column in the dataset.
	CreateColumn bool

	// Additional parameter options such as a format and a timezone. Required for
	// datetime parameters.
	DatetimeOptions *DatetimeOptions

	// The optional filter expression structure to apply additional matching criteria
	// to the parameter.
	Filter *FilterExpression
}

// Represents additional options for correct interpretation of datetime parameters
// used in the S3 path of a dataset.
type DatetimeOptions struct {

	// Required option, that defines the datetime format used for a date parameter in
	// the S3 path. Should use only supported datetime specifiers and separation
	// characters, all litera a-z or A-Z character should be escaped with single
	// quotes. E.g. "MM.dd.yyyy-'at'-HH:mm".
	//
	// This member is required.
	Format *string

	// Optional value for a non-US locale code, needed for correct interpretation of
	// some date formats.
	LocaleCode *string

	// Optional value for a timezone offset of the datetime parameter value in the S3
	// path. Shouldn't be used if Format for this parameter includes timezone fields.
	// If no offset specified, UTC is assumed.
	TimezoneOffset *string
}

// Represents a set of options that define how DataBrew will interpret a Microsoft
// Excel file when creating a dataset from that file.
type ExcelOptions struct {

	// A variable that specifies whether the first row in the file is parsed as the
	// header. If this value is false, column names are auto-generated.
	HeaderRow *bool

	// One or more sheet numbers in the Excel file that will be included in the
	// dataset.
	SheetIndexes []int32

	// One or more named sheets in the Excel file that will be included in the dataset.
	SheetNames []string
}

// Represents a limit imposed on number of S3 files that should be selected for a
// dataset from a connected S3 path.
type FilesLimit struct {

	// The number of S3 files to select.
	//
	// This member is required.
	MaxFiles int32

	// A criteria to use for S3 files sorting before their selection. By default uses
	// DESCENDING order, i.e. most recent files are selected first. Anotherpossible
	// value is ASCENDING.
	Order Order

	// A criteria to use for S3 files sorting before their selection. By default uses
	// LAST_MODIFIED_DATE as a sorting criteria. Currently it's the only allowed value.
	OrderedBy OrderedBy
}

// Represents a structure for defining parameter conditions.
type FilterExpression struct {

	// The expression which includes condition names followed by substitution
	// variables, possibly grouped and combined with other conditions. For example,
	// "(starts_with :prefix1 or starts_with :prefix2) and (ends_with :suffix1 or
	// ends_with :suffix2)". Substitution variables should start with ':' symbol.
	//
	// This member is required.
	Expression *string

	// The map of substitution variable names to their values used in this filter
	// expression.
	//
	// This member is required.
	ValuesMap map[string]string
}

// Represents a set of options that define the structure of either comma-separated
// value (CSV), Excel, or JSON input.
type FormatOptions struct {

	// Options that define how CSV input is to be interpreted by DataBrew.
	Csv *CsvOptions

	// Options that define how Excel input is to be interpreted by DataBrew.
	Excel *ExcelOptions

	// Options that define how JSON input is to be interpreted by DataBrew.
	Json *JsonOptions
}

// Represents information on how DataBrew can find data, in either the AWS Glue
// Data Catalog or Amazon S3.
type Input struct {

	// The AWS Glue Data Catalog parameters for the data.
	DataCatalogInputDefinition *DataCatalogInputDefinition

	// Connection information for dataset input files stored in a database.
	DatabaseInputDefinition *DatabaseInputDefinition

	// The Amazon S3 location where the data is stored.
	S3InputDefinition *S3Location
}

// Represents all of the attributes of a DataBrew job.
type Job struct {

	// The unique name of the job.
	//
	// This member is required.
	Name *string

	// The ID of the AWS account that owns the job.
	AccountId *string

	// The date and time that the job was created.
	CreateDate *time.Time

	// The Amazon Resource Name (ARN) of the user who created the job.
	CreatedBy *string

	// A dataset that the job is to process.
	DatasetName *string

	// The Amazon Resource Name (ARN) of an encryption key that is used to protect the
	// job output. For more information, see Encrypting data written by DataBrew jobs
	// (https://docs.aws.amazon.com/databrew/latest/dg/encryption-security-configuration.html)
	EncryptionKeyArn *string

	// The encryption mode for the job, which can be one of the following:
	//
	// * SSE-KMS -
	// Server-side encryption with keys managed by AWS KMS.
	//
	// * SSE-S3 - Server-side
	// encryption with keys managed by Amazon S3.
	EncryptionMode EncryptionMode

	// A sample configuration for profile jobs only, which determines the number of
	// rows on which the profile job is run. If a JobSample value isn't provided, the
	// default value is used. The default value is CUSTOM_ROWS for the mode parameter
	// and 20,000 for the size parameter.
	JobSample *JobSample

	// The Amazon Resource Name (ARN) of the user who last modified the job.
	LastModifiedBy *string

	// The modification date and time of the job.
	LastModifiedDate *time.Time

	// The current status of Amazon CloudWatch logging for the job.
	LogSubscription LogSubscription

	// The maximum number of nodes that can be consumed when the job processes data.
	MaxCapacity int32

	// The maximum number of times to retry the job after a job run fails.
	MaxRetries int32

	// One or more artifacts that represent output from running the job.
	Outputs []Output

	// The name of the project that the job is associated with.
	ProjectName *string

	// A set of steps that the job runs.
	RecipeReference *RecipeReference

	// The unique Amazon Resource Name (ARN) for the job.
	ResourceArn *string

	// The Amazon Resource Name (ARN) of the role to be assumed for this job.
	RoleArn *string

	// Metadata tags that have been applied to the job.
	Tags map[string]string

	// The job's timeout in minutes. A job that attempts to run longer than this
	// timeout period ends with a status of TIMEOUT.
	Timeout int32

	// The job type of the job, which must be one of the following:
	//
	// * PROFILE - A job
	// to analyze a dataset, to determine its size, data types, data distribution, and
	// more.
	//
	// * RECIPE - A job to apply one or more transformations to a dataset.
	Type JobType
}

// Represents one run of a DataBrew job.
type JobRun struct {

	// The number of times that DataBrew has attempted to run the job.
	Attempt int32

	// The date and time when the job completed processing.
	CompletedOn *time.Time

	// The name of the dataset for the job to process.
	DatasetName *string

	// A message indicating an error (if any) that was encountered when the job ran.
	ErrorMessage *string

	// The amount of time, in seconds, during which a job run consumed resources.
	ExecutionTime int32

	// The name of the job being processed during this run.
	JobName *string

	// A sample configuration for profile jobs only, which determines the number of
	// rows on which the profile job is run. If a JobSample value isn't provided, the
	// default is used. The default value is CUSTOM_ROWS for the mode parameter and
	// 20,000 for the size parameter.
	JobSample *JobSample

	// The name of an Amazon CloudWatch log group, where the job writes diagnostic
	// messages when it runs.
	LogGroupName *string

	// The current status of Amazon CloudWatch logging for the job run.
	LogSubscription LogSubscription

	// One or more output artifacts from a job run.
	Outputs []Output

	// The set of steps processed by the job.
	RecipeReference *RecipeReference

	// The unique identifier of the job run.
	RunId *string

	// The Amazon Resource Name (ARN) of the user who initiated the job run.
	StartedBy *string

	// The date and time when the job run began.
	StartedOn *time.Time

	// The current state of the job run entity itself.
	State JobRunState
}

// A sample configuration for profile jobs only, which determines the number of
// rows on which the profile job is run. If a JobSample value isn't provided, the
// default is used. The default value is CUSTOM_ROWS for the mode parameter and
// 20,000 for the size parameter.
type JobSample struct {

	// A value that determines whether the profile job is run on the entire dataset or
	// a specified number of rows. This value must be one of the following:
	//
	// *
	// FULL_DATASET - The profile job is run on the entire dataset.
	//
	// * CUSTOM_ROWS -
	// The profile job is run on the number of rows specified in the Size parameter.
	Mode SampleMode

	// The Size parameter is only required when the mode is CUSTOM_ROWS. The profile
	// job is run on the specified number of rows. The maximum value for size is
	// Long.MAX_VALUE. Long.MAX_VALUE = 9223372036854775807
	Size *int64
}

// Represents the JSON-specific options that define how input is to be interpreted
// by AWS Glue DataBrew.
type JsonOptions struct {

	// A value that specifies whether JSON input contains embedded new line characters.
	MultiLine bool
}

// Represents options that specify how and where DataBrew writes the output
// generated by recipe jobs or profile jobs.
type Output struct {

	// The location in Amazon S3 where the job writes its output.
	//
	// This member is required.
	Location *S3Location

	// The compression algorithm used to compress the output text of the job.
	CompressionFormat CompressionFormat

	// The data format of the output of the job.
	Format OutputFormat

	// Represents options that define how DataBrew formats job output files.
	FormatOptions *OutputFormatOptions

	// A value that, if true, means that any data in the location specified for output
	// is overwritten with new output.
	Overwrite bool

	// The names of one or more partition columns for the output of the job.
	PartitionColumns []string
}

// Represents a set of options that define the structure of comma-separated (CSV)
// job output.
type OutputFormatOptions struct {

	// Represents a set of options that define the structure of comma-separated value
	// (CSV) job output.
	Csv *CsvOutputOptions
}

// Represents a set of options that define how DataBrew selects files for a given
// S3 path in a dataset.
type PathOptions struct {

	// If provided, this structure imposes a limit on a number of files that should be
	// selected.
	FilesLimit *FilesLimit

	// If provided, this structure defines a date range for matching S3 objects based
	// on their LastModifiedDate attribute in S3.
	LastModifiedDateCondition *FilterExpression

	// A structure that maps names of parameters used in the S3 path of a dataset to
	// their definitions.
	Parameters map[string]DatasetParameter
}

// Represents all of the attributes of a DataBrew project.
type Project struct {

	// The unique name of a project.
	//
	// This member is required.
	Name *string

	// The name of a recipe that will be developed during a project session.
	//
	// This member is required.
	RecipeName *string

	// The ID of the AWS account that owns the project.
	AccountId *string

	// The date and time that the project was created.
	CreateDate *time.Time

	// The Amazon Resource Name (ARN) of the user who crated the project.
	CreatedBy *string

	// The dataset that the project is to act upon.
	DatasetName *string

	// The Amazon Resource Name (ARN) of the user who last modified the project.
	LastModifiedBy *string

	// The last modification date and time for the project.
	LastModifiedDate *time.Time

	// The date and time when the project was opened.
	OpenDate *time.Time

	// The Amazon Resource Name (ARN) of the user that opened the project for use.
	OpenedBy *string

	// The Amazon Resource Name (ARN) for the project.
	ResourceArn *string

	// The Amazon Resource Name (ARN) of the role that will be assumed for this
	// project.
	RoleArn *string

	// The sample size and sampling type to apply to the data. If this parameter isn't
	// specified, then the sample consists of the first 500 rows from the dataset.
	Sample *Sample

	// Metadata tags that have been applied to the project.
	Tags map[string]string
}

// Represents one or more actions to be performed on a DataBrew dataset.
type Recipe struct {

	// The unique name for the recipe.
	//
	// This member is required.
	Name *string

	// The date and time that the recipe was created.
	CreateDate *time.Time

	// The Amazon Resource Name (ARN) of the user who created the recipe.
	CreatedBy *string

	// The description of the recipe.
	Description *string

	// The Amazon Resource Name (ARN) of the user who last modified the recipe.
	LastModifiedBy *string

	// The last modification date and time of the recipe.
	LastModifiedDate *time.Time

	// The name of the project that the recipe is associated with.
	ProjectName *string

	// The Amazon Resource Name (ARN) of the user who published the recipe.
	PublishedBy *string

	// The date and time when the recipe was published.
	PublishedDate *time.Time

	// The identifier for the version for the recipe. Must be one of the following:
	//
	// *
	// Numeric version (X.Y) - X and Y stand for major and minor version numbers. The
	// maximum length of each is 6 digits, and neither can be negative values. Both X
	// and Y are required, and "0.0" isn't a valid version.
	//
	// * LATEST_WORKING - the
	// most recent valid version being developed in a DataBrew project.
	//
	// *
	// LATEST_PUBLISHED - the most recent published version.
	RecipeVersion *string

	// The Amazon Resource Name (ARN) for the recipe.
	ResourceArn *string

	// A list of steps that are defined by the recipe.
	Steps []RecipeStep

	// Metadata tags that have been applied to the recipe.
	Tags map[string]string
}

// Represents a transformation and associated parameters that are used to apply a
// change to a DataBrew dataset. For more information, see Recipe structure
// (https://docs.aws.amazon.com/databrew/latest/dg/recipe-structure.html) and
// Recipe actions reference
// (https://docs.aws.amazon.com/databrew/latest/dg/recipe-actions-reference.html).
type RecipeAction struct {

	// The name of a valid DataBrew transformation to be performed on the data.
	//
	// This member is required.
	Operation *string

	// Contextual parameters for the transformation.
	Parameters map[string]string
}

// Represents the name and version of a DataBrew recipe.
type RecipeReference struct {

	// The name of the recipe.
	//
	// This member is required.
	Name *string

	// The identifier for the version for the recipe.
	RecipeVersion *string
}

// Represents a single step from a DataBrew recipe to be performed.
type RecipeStep struct {

	// The particular action to be performed in the recipe step.
	//
	// This member is required.
	Action *RecipeAction

	// One or more conditions that must be met for the recipe step to succeed. All of
	// the conditions in the array must be met. In other words, all of the conditions
	// must be combined using a logical AND operation.
	ConditionExpressions []ConditionExpression
}

// Represents any errors encountered when attempting to delete multiple recipe
// versions.
type RecipeVersionErrorDetail struct {

	// The HTTP status code for the error.
	ErrorCode *string

	// The text of the error message.
	ErrorMessage *string

	// The identifier for the recipe version associated with this error.
	RecipeVersion *string
}

// Represents an Amazon S3 location (bucket name and object key) where DataBrew can
// read input data, or write output from a job.
type S3Location struct {

	// The S3 bucket name.
	//
	// This member is required.
	Bucket *string

	// The unique name of the object in the bucket.
	Key *string
}

// Represents the sample size and sampling type for DataBrew to use for interactive
// data analysis.
type Sample struct {

	// The way in which DataBrew obtains rows from a dataset.
	//
	// This member is required.
	Type SampleType

	// The number of rows in the sample.
	Size *int32
}

// Represents one or more dates and times when a job is to run.
type Schedule struct {

	// The name of the schedule.
	//
	// This member is required.
	Name *string

	// The ID of the AWS account that owns the schedule.
	AccountId *string

	// The date and time that the schedule was created.
	CreateDate *time.Time

	// The Amazon Resource Name (ARN) of the user who created the schedule.
	CreatedBy *string

	// The dates and times when the job is to run. For more information, see Cron
	// expressions (https://docs.aws.amazon.com/databrew/latest/dg/jobs.cron.html) in
	// the AWS Glue DataBrew Developer Guide.
	CronExpression *string

	// A list of jobs to be run, according to the schedule.
	JobNames []string

	// The Amazon Resource Name (ARN) of the user who last modified the schedule.
	LastModifiedBy *string

	// The date and time when the schedule was last modified.
	LastModifiedDate *time.Time

	// The Amazon Resource Name (ARN) of the schedule.
	ResourceArn *string

	// Metadata tags that have been applied to the schedule.
	Tags map[string]string
}

// Represents the data being transformed during an action.
type ViewFrame struct {

	// The starting index for the range of columns to return in the view frame.
	//
	// This member is required.
	StartColumnIndex *int32

	// The number of columns to include in the view frame, beginning with the
	// StartColumnIndex value and ignoring any columns in the HiddenColumns list.
	ColumnRange *int32

	// A list of columns to hide in the view frame.
	HiddenColumns []string
}
