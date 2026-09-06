// Apps Script API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package script

// The Content resource.
type Content struct {
	Files []File `json:"files,omitempty\"` // The list of script project files. One of the files is a script manifest; it must be named "appsscript", must have type of JSON, and include the manifest configurations for the project.

	ScriptId string `json:"scriptId,omitempty\"` // The script project's Drive ID.

}

// Request to create a script project.
type CreateProjectRequest struct {
	ParentId string `json:"parentId,omitempty\"` // The Drive ID of a parent file that the created script project is bound to. This is usually the ID of a Google Doc, Google Sheet, Google Form, or Google Slides file. If not set, a standalone script project is created.

	Title string `json:"title,omitempty\"` // The title for the project.

}

// Representation of a single script deployment.
type Deployment struct {
	DeploymentConfig DeploymentConfig `json:"deploymentConfig,omitempty\"` // The deployment configuration.

	DeploymentId string `json:"deploymentId,omitempty\"` // The deployment ID for this deployment.

	EntryPoints []EntryPoint `json:"entryPoints,omitempty\"` // The deployment's entry points.

	UpdateTime string `json:"updateTime,omitempty\"` // Last modified date time stamp.

}

// Metadata the defines how a deployment is configured.
type DeploymentConfig struct {
	Description string `json:"description,omitempty\"` // The description for this deployment.

	ManifestFileName string `json:"manifestFileName,omitempty\"` // The manifest file name for this deployment.

	ScriptId string `json:"scriptId,omitempty\"` // The script project's Drive ID.

	VersionNumber int `json:"versionNumber,omitempty\"` // The version number on which this deployment is based.

}

// A generic empty message that you can re-use to avoid defining duplicated empty messages in your APIs. A typical example is to use it as the request or the response type of an API method. For instance: service Foo { rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty); }
type Empty struct {
}

// A configuration that defines how a deployment is accessed externally.
type EntryPoint struct {
	AddOn GoogleAppsScriptTypeAddOnEntryPoint `json:"addOn,omitempty\"` // Add-on properties.

	EntryPointType string `json:"entryPointType,omitempty\"` // The type of the entry point.

	ExecutionApi GoogleAppsScriptTypeExecutionApiEntryPoint `json:"executionApi,omitempty\"` // An entry point specification for Apps Script API execution calls.

	WebApp GoogleAppsScriptTypeWebAppEntryPoint `json:"webApp,omitempty\"` // An entry point specification for web apps.

}

// An object that provides information about the nature of an error resulting from an attempted execution of a script function using the Apps Script API. If a run call succeeds but the script function (or Apps Script itself) throws an exception, the response body's error field contains a Status object. The `Status` object's `details` field contains an array with a single one of these `ExecutionError` objects.
type ExecutionError struct {
	ErrorMessage string `json:"errorMessage,omitempty\"` // The error message thrown by Apps Script, usually localized into the user's language.

	ErrorType string `json:"errorType,omitempty\"` // The error type, for example `TypeError` or `ReferenceError`. If the error type is unavailable, this field is not included.

	ScriptStackTraceElements []ScriptStackTraceElement `json:"scriptStackTraceElements,omitempty\"` // An array of objects that provide a stack trace through the script to show where the execution failed, with the deepest call first.

}

// A request to run the function in a script. The script is identified by the specified `script_id`. Executing a function on a script returns results based on the implementation of the script.
type ExecutionRequest struct {
	DevMode bool `json:"devMode,omitempty\"` // If `true` and the user is an owner of the script, the script runs at the most recently saved version rather than the version deployed for use with the Apps Script API. Optional; default is `false`.

	Function string `json:"function,omitempty\"` // The name of the function to execute in the given script. The name does not include parentheses or parameters. It can reference a function in an included library such as `Library.libFunction1`.

	Parameters []interface{} `json:"parameters,omitempty\"` // The parameters to be passed to the function being executed. The object type for each parameter should match the expected type in Apps Script. Parameters cannot be Apps Script-specific object types (such as a `Document` or a `Calendar`); they can only be primitive types such as `string`, `number`, `array`, `object`, or `boolean`. Optional.

	SessionState string `json:"sessionState,omitempty\"` // *Deprecated*. For use with Android add-ons only. An ID that represents the user's current session in the Android app for Google Docs or Sheets, included as extra data in the [Intent](https://developer.android.com/guide/components/intents-filters.html) that launches the add-on. When an Android add-on is run with a session state, it gains the privileges of a [bound](https://developers.google.com/apps-script/guides/bound) script—that is, it can access information like the user's current cursor position (in Docs) or selected cell (in Sheets). To retrieve the state, call `Intent.getStringExtra("com.google.android.apps.docs.addons.SessionState")`. Optional.

}

// An object that provides the return value of a function executed using the Apps Script API. If the script function returns successfully, the response body's response field contains this `ExecutionResponse` object.
type ExecutionResponse struct {
	Result interface{} `json:"result,omitempty\"` // The return value of the script function. The type matches the object type returned in Apps Script. Functions called using the Apps Script API cannot return Apps Script-specific objects (such as a `Document` or a `Calendar`); they can only return primitive types such as a `string`, `number`, `array`, `object`, or `boolean`.

}

// An individual file within a script project. A file is a third-party source code created by one or more developers. It can be a server-side JS code, HTML, or a configuration file. Each script project can contain multiple files.
type File struct {
	CreateTime string `json:"createTime,omitempty\"` // Creation date timestamp.

	FunctionSet GoogleAppsScriptTypeFunctionSet `json:"functionSet,omitempty\"` // The defined set of functions in the script file, if any.

	LastModifyUser GoogleAppsScriptTypeUser `json:"lastModifyUser,omitempty\"` // The user who modified the file most recently. The details visible in this object are controlled by the profile visibility settings of the last modifying user.

	Name string `json:"name,omitempty\"` // The name of the file. The file extension is not part of the file name, which can be identified from the type field.

	Source string `json:"source,omitempty\"` // The file content.

	TypeValue string `json:"type,omitempty\"` // The type of the file.

	UpdateTime string `json:"updateTime,omitempty\"` // Last modified date timestamp.

}

// An add-on entry point.
type GoogleAppsScriptTypeAddOnEntryPoint struct {
	AddOnType string `json:"addOnType,omitempty\"` // The add-on's required list of supported container types.

	Description string `json:"description,omitempty\"` // The add-on's optional description.

	HelpUrl string `json:"helpUrl,omitempty\"` // The add-on's optional help URL.

	PostInstallTipUrl string `json:"postInstallTipUrl,omitempty\"` // The add-on's required post install tip URL.

	ReportIssueUrl string `json:"reportIssueUrl,omitempty\"` // The add-on's optional report issue URL.

	Title string `json:"title,omitempty\"` // The add-on's required title.

}

// API executable entry point configuration.
type GoogleAppsScriptTypeExecutionApiConfig struct {
	Access string `json:"access,omitempty\"` // Who has permission to run the API executable.

}

// An API executable entry point.
type GoogleAppsScriptTypeExecutionApiEntryPoint struct {
	EntryPointConfig GoogleAppsScriptTypeExecutionApiConfig `json:"entryPointConfig,omitempty\"` // The entry point's configuration.

}

// Represents a function in a script project.
type GoogleAppsScriptTypeFunction struct {
	Name string `json:"name,omitempty\"` // The function name in the script project.

	Parameters []string `json:"parameters,omitempty\"` // The ordered list of parameter names of the function in the script project.

}

// A set of functions. No duplicates are permitted.
type GoogleAppsScriptTypeFunctionSet struct {
	Values []GoogleAppsScriptTypeFunction `json:"values,omitempty\"` // A list of functions composing the set.

}

// Representation of a single script process execution that was started from the script editor, a trigger, an application, or using the Apps Script API. This is distinct from the `Operation` resource, which only represents executions started via the Apps Script API.
type GoogleAppsScriptTypeProcess struct {
	Duration string `json:"duration,omitempty\"` // Duration the execution spent executing.

	FunctionName string `json:"functionName,omitempty\"` // Name of the function the started the execution.

	ProcessStatus string `json:"processStatus,omitempty\"` // The executions status.

	ProcessType string `json:"processType,omitempty\"` // The executions type.

	ProjectName string `json:"projectName,omitempty\"` // Name of the script being executed.

	RuntimeVersion string `json:"runtimeVersion,omitempty\"` // Which version of maestro to use to execute the script.

	StartTime string `json:"startTime,omitempty\"` // Time the execution started.

	UserAccessLevel string `json:"userAccessLevel,omitempty\"` // The executing users access level to the script.

}

// A simple user profile resource.
type GoogleAppsScriptTypeUser struct {
	Domain string `json:"domain,omitempty\"` // The user's domain.

	Email string `json:"email,omitempty\"` // The user's identifying email address.

	Name string `json:"name,omitempty\"` // The user's display name.

	PhotoUrl string `json:"photoUrl,omitempty\"` // The user's photo.

}

// Web app entry point configuration.
type GoogleAppsScriptTypeWebAppConfig struct {
	Access string `json:"access,omitempty\"` // Who has permission to run the web app.

	ExecuteAs string `json:"executeAs,omitempty\"` // Who to execute the web app as.

}

// A web application entry point.
type GoogleAppsScriptTypeWebAppEntryPoint struct {
	EntryPointConfig GoogleAppsScriptTypeWebAppConfig `json:"entryPointConfig,omitempty\"` // The entry point's configuration.

	Url string `json:"url,omitempty\"` // The URL for the web application.

}

// Response with the list of deployments for the specified Apps Script project.
type ListDeploymentsResponse struct {
	Deployments []Deployment `json:"deployments,omitempty\"` // The list of deployments.

	NextPageToken string `json:"nextPageToken,omitempty\"` // The token that can be used in the next call to get the next page of results.

}

// Response with the list of Process resources.
type ListScriptProcessesResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty\"` // Token for the next page of results. If empty, there are no more pages remaining.

	Processes []GoogleAppsScriptTypeProcess `json:"processes,omitempty\"` // List of processes matching request parameters.

}

// Response with the list of Process resources.
type ListUserProcessesResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty\"` // Token for the next page of results. If empty, there are no more pages remaining.

	Processes []GoogleAppsScriptTypeProcess `json:"processes,omitempty\"` // List of processes matching request parameters.

}

// Response with the list of the versions for the specified script project.
type ListVersionsResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty\"` // The token use to fetch the next page of records. if not exist in the response, that means no more versions to list.

	Versions []Version `json:"versions,omitempty\"` // The list of versions.

}

// Resource containing usage stats for a given script, based on the supplied filter and mask present in the request.
type Metrics struct {
	ActiveUsers []MetricsValue `json:"activeUsers,omitempty\"` // Number of active users.

	FailedExecutions []MetricsValue `json:"failedExecutions,omitempty\"` // Number of failed executions.

	TotalExecutions []MetricsValue `json:"totalExecutions,omitempty\"` // Number of total executions.

}

// Metrics value that holds number of executions counted.
type MetricsValue struct {
	EndTime string `json:"endTime,omitempty\"` // Required field indicating the end time of the interval.

	StartTime string `json:"startTime,omitempty\"` // Required field indicating the start time of the interval.

	Value uint64 `json:"value,omitempty\"` // Indicates the number of executions counted.

}

// A representation of an execution of an Apps Script function started with run. The execution response does not arrive until the function finishes executing. The maximum execution runtime is listed in the [Apps Script quotas guide](/apps-script/guides/services/quotas#current_limitations). After execution has started, it can have one of four outcomes: - If the script function returns successfully, the response field contains an ExecutionResponse object with the function's return value in the object's `result` field. - If the script function (or Apps Script itself) throws an exception, the error field contains a Status object. The `Status` object's `details` field contains an array with a single ExecutionError object that provides information about the nature of the error. - If the execution has not yet completed, the done field is `false` and the neither the `response` nor `error` fields are present. - If the `run` call itself fails (for example, because of a malformed request or an authorization error), the method returns an HTTP response code in the 4XX range with a different format for the response body. Client libraries automatically convert a 4XX response into an exception class.
type Operation struct {
	Done bool `json:"done,omitempty\"` // This field indicates whether the script execution has completed. A completed execution has a populated `response` field containing the ExecutionResponse from function that was executed.

	Error Status `json:"error,omitempty\"` // If a `run` call succeeds but the script function (or Apps Script itself) throws an exception, this field contains a Status object. The `Status` object's `details` field contains an array with a single ExecutionError object that provides information about the nature of the error.

	Response map[string]interface{} `json:"response,omitempty\"` // If the script function returns successfully, this field contains an ExecutionResponse object with the function's return value.

}

// The script project resource.
type Project struct {
	CreateTime string `json:"createTime,omitempty\"` // When the script was created.

	Creator GoogleAppsScriptTypeUser `json:"creator,omitempty\"` // User who originally created the script.

	LastModifyUser GoogleAppsScriptTypeUser `json:"lastModifyUser,omitempty\"` // User who last modified the script.

	ParentId string `json:"parentId,omitempty\"` // The parent's Drive ID that the script will be attached to. This is usually the ID of a Google Document or Google Sheet. This field is optional, and if not set, a stand-alone script will be created.

	ScriptId string `json:"scriptId,omitempty\"` // The script project's Drive ID.

	Title string `json:"title,omitempty\"` // The title for the project.

	UpdateTime string `json:"updateTime,omitempty\"` // When the script was last updated.

}

// A stack trace through the script that shows where the execution failed.
type ScriptStackTraceElement struct {
	Function string `json:"function,omitempty\"` // The name of the function that failed.

	LineNumber int `json:"lineNumber,omitempty\"` // The line number where the script failed.

}

// If a `run` call succeeds but the script function (or Apps Script itself) throws an exception, the response body's error field contains this `Status` object.
type Status struct {
	Code int `json:"code,omitempty\"` // The status code. For this API, this value either: - 10, indicating a `SCRIPT_TIMEOUT` error, - 3, indicating an `INVALID_ARGUMENT` error, or - 1, indicating a `CANCELLED` execution.

	Details []map[string]interface{} `json:"details,omitempty\"` // An array that contains a single ExecutionError object that provides information about the nature of the error.

	Message string `json:"message,omitempty\"` // A developer-facing error message, which is in English. Any user-facing error message is localized and sent in the details field, or localized by the client.

}

// Request with deployment information to update an existing deployment.
type UpdateDeploymentRequest struct {
	DeploymentConfig DeploymentConfig `json:"deploymentConfig,omitempty\"` // The deployment configuration.

}

// A resource representing a script project version. A version is a "snapshot" of a script project and is similar to a read-only branched release. When creating deployments, the version to use must be specified.
type Version struct {
	CreateTime string `json:"createTime,omitempty\"` // When the version was created.

	Description string `json:"description,omitempty\"` // The description for this version.

	ScriptId string `json:"scriptId,omitempty\"` // The script project's Drive ID.

	VersionNumber int `json:"versionNumber,omitempty\"` // The incremental ID that is created by Apps Script when a version is created. This is system assigned number and is immutable once created.

}
