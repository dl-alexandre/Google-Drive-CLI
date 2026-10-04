// Google Tasks API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package tasks

// Information about the source of the task assignment (Document, Chat Space).
type AssignmentInfo struct {
	DriveResourceInfo DriveResourceInfo `json:"driveResourceInfo,omitempty\"` // Output only. Information about the Drive file where this task originates from. Currently, the Drive file can only be a document. This field is read-only.

	LinkToTask string `json:"linkToTask,omitempty\"` // Output only. An absolute link to the original task in the surface of assignment (Docs, Chat spaces, etc.).

	SpaceInfo SpaceInfo `json:"spaceInfo,omitempty\"` // Output only. Information about the Chat Space where this task originates from. This field is read-only.

	SurfaceType string `json:"surfaceType,omitempty\"` // Output only. The type of surface this assigned task originates from. Currently limited to DOCUMENT or SPACE.

}

// Information about the Drive resource where a task was assigned from (the document, sheet, etc.).
type DriveResourceInfo struct {
	DriveFileId string `json:"driveFileId,omitempty\"` // Output only. Identifier of the file in the Drive API.

	ResourceKey string `json:"resourceKey,omitempty\"` // Output only. Resource key required to access files shared via a shared link. Not required for all files. See also developers.google.com/drive/api/guides/resource-keys.

}

// Information about the Chat Space where a task was assigned from.
type SpaceInfo struct {
	Space string `json:"space,omitempty\"` // Output only. The Chat space where this task originates from. The format is "spaces/{space}".

}

type Task struct {
	AssignmentInfo AssignmentInfo `json:"assignmentInfo,omitempty\"` // Output only. Context information for assigned tasks. A task can be assigned to a user, currently possible from surfaces like Docs and Chat Spaces. This field is populated for tasks assigned to the current user and identifies where the task was assigned from. This field is read-only.

	Completed string `json:"completed,omitempty\"` // Completion date of the task (as a RFC 3339 timestamp). This field is omitted if the task has not been completed.

	Deleted bool `json:"deleted,omitempty\"` // Flag indicating whether the task has been deleted. For assigned tasks this field is read-only. They can only be deleted by calling tasks.delete, in which case both the assigned task and the original task (in Docs or Chat Spaces) are deleted. To delete the assigned task only, navigate to the assignment surface and unassign the task from there. The default is False.

	Due string `json:"due,omitempty\"` // Scheduled date for the task (as an RFC 3339 timestamp). Optional. This represents the day that the task should be done, or that the task is visible on the calendar grid. It doesn't represent the deadline of the task. Only date information is recorded; the time portion of the timestamp is discarded when setting this field. It isn't possible to read or write the time that a task is scheduled for using the API.

	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Hidden bool `json:"hidden,omitempty\"` // Flag indicating whether the task is hidden. This is the case if the task had been marked completed when the task list was last cleared. The default is False. This field is read-only.

	Id string `json:"id,omitempty\"` // Task identifier.

	Kind string `json:"kind,omitempty\"` // Output only. Type of the resource. This is always "tasks#task".

	Links []map[string]interface{} `json:"links,omitempty\"` // Output only. Collection of links. This collection is read-only.

	Notes string `json:"notes,omitempty\"` // Notes describing the task. Tasks assigned from Google Docs cannot have notes. Optional. Maximum length allowed: 8192 characters.

	Parent string `json:"parent,omitempty\"` // Output only. Parent task identifier. This field is omitted if it is a top-level task. Use the "move" method to move the task under a different parent or to the top level. A parent task can never be an assigned task (from Chat Spaces, Docs). This field is read-only.

	Position string `json:"position,omitempty\"` // Output only. String indicating the position of the task among its sibling tasks under the same parent task or at the top level. If this string is greater than another task's corresponding position string according to lexicographical ordering, the task is positioned after the other task under the same parent task (or at the top level). Use the "move" method to move the task to another position.

	SelfLink string `json:"selfLink,omitempty\"` // Output only. URL pointing to this task. Used to retrieve, update, or delete this task.

	Status string `json:"status,omitempty\"` // Status of the task. This is either "needsAction" or "completed".

	Title string `json:"title,omitempty\"` // Title of the task. Maximum length allowed: 1024 characters.

	Updated string `json:"updated,omitempty\"` // Output only. Last modification time of the task (as a RFC 3339 timestamp).

	WebViewLink string `json:"webViewLink,omitempty\"` // Output only. An absolute link to the task in the Google Tasks Web UI.

}

type TaskList struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Id string `json:"id,omitempty\"` // Task list identifier.

	Kind string `json:"kind,omitempty\"` // Output only. Type of the resource. This is always "tasks#taskList".

	SelfLink string `json:"selfLink,omitempty\"` // Output only. URL pointing to this task list. Used to retrieve, update, or delete this task list.

	Title string `json:"title,omitempty\"` // Title of the task list. Maximum length allowed: 1024 characters.

	Updated string `json:"updated,omitempty\"` // Output only. Last modification time of the task list (as a RFC 3339 timestamp).

}

type TaskLists struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Items []TaskList `json:"items,omitempty\"` // Collection of task lists.

	Kind string `json:"kind,omitempty\"` // Type of the resource. This is always "tasks#taskLists".

	NextPageToken string `json:"nextPageToken,omitempty\"` // Token that can be used to request the next page of this result.

}

type Tasks struct {
	Etag string `json:"etag,omitempty\"` // ETag of the resource.

	Items []Task `json:"items,omitempty\"` // Collection of tasks.

	Kind string `json:"kind,omitempty\"` // Type of the resource. This is always "tasks#tasks".

	NextPageToken string `json:"nextPageToken,omitempty\"` // Token used to access the next page of this result.

}
