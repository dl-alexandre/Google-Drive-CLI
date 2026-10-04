// Google Drive API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package drive

import "time"

// Information about the user, the user's Drive, and system capabilities.
type About struct {
	AppInstalled bool `json:"appInstalled,omitempty\"` // Whether the user has installed the requesting app.

	CanCreateDrives bool `json:"canCreateDrives,omitempty\"` // Whether the user can create shared drives.

	CanCreateTeamDrives bool `json:"canCreateTeamDrives,omitempty\"` // Deprecated: Use `canCreateDrives` instead.

	DriveThemes []map[string]interface{} `json:"driveThemes,omitempty\"` // A list of themes that are supported for shared drives.

	ExportFormats map[string]interface{} `json:"exportFormats,omitempty\"` // A map of source MIME type to possible targets for all supported exports.

	FolderColorPalette []string `json:"folderColorPalette,omitempty\"` // The currently supported folder colors as RGB hex strings.

	ImportFormats map[string]interface{} `json:"importFormats,omitempty\"` // A map of source MIME type to possible targets for all supported imports.

	Kind string `json:"kind,omitempty\"` // Identifies what kind of resource this is. Value: the fixed string `"drive#about"`.

	MaxImportSizes map[string]interface{} `json:"maxImportSizes,omitempty\"` // A map of maximum import sizes by MIME type, in bytes.

	MaxUploadSize int64 `json:"maxUploadSize,omitempty\"` // The maximum upload size in bytes.

	StorageQuota map[string]interface{} `json:"storageQuota,omitempty\"` // The user's storage quota limits and usage. For users that are part of an organization with pooled storage, information about the limit and usage across all services is for the organization, rather than the individual user. All fields are measured in bytes.

	TeamDriveThemes []map[string]interface{} `json:"teamDriveThemes,omitempty\"` // Deprecated: Use `driveThemes` instead.

	User User `json:"user,omitempty\"` // The authenticated user.

}

// Manage outstanding access proposals on a file.
type AccessProposal struct {
	CreateTime string `json:"createTime,omitempty\"` // The creation time.

	FileId string `json:"fileId,omitempty\"` // The file ID that the proposal for access is on.

	ProposalId string `json:"proposalId,omitempty\"` // The ID of the access proposal.

	RecipientEmailAddress string `json:"recipientEmailAddress,omitempty\"` // The email address of the user that will receive permissions, if accepted.

	RequestMessage string `json:"requestMessage,omitempty\"` // The message that the requester added to the proposal.

	RequesterEmailAddress string `json:"requesterEmailAddress,omitempty\"` // The email address of the requesting user.

	RolesAndViews []AccessProposalRoleAndView `json:"rolesAndViews,omitempty\"` // A wrapper for the role and view of an access proposal. For more information, see [Roles and permissions](https://developers.google.com/workspace/drive/api/guides/ref-roles).

}

// A wrapper for the role and view of an access proposal. For more information, see [Roles and permissions](https://developers.google.com/workspace/drive/api/guides/ref-roles).
type AccessProposalRoleAndView struct {
	Role string `json:"role,omitempty\"` // The role that was proposed by the requester. The supported values are: * `writer` * `commenter` * `reader`

	View string `json:"view,omitempty\"` // Indicates the view for this access proposal. Only populated for proposals that belong to a view. Only `published` is supported.

}

// Representation of a reviewer addition.
type AddReviewer struct {
	AddedReviewerEmail string `json:"addedReviewerEmail,omitempty\"` // Required. The email of the reviewer to add.

}

// The `apps` resource provides a list of apps that a user has installed, with information about each app's supported MIME types, file extensions, and other details. Some resource methods (such as `apps.get`) require an `appId`. Use the `apps.list` method to retrieve the ID for an installed application.
type App struct {
	Authorized bool `json:"authorized,omitempty\"` // Whether the app is authorized to access data on the user's Drive.

	CreateInFolderTemplate string `json:"createInFolderTemplate,omitempty\"` // The template URL to create a file with this app in a given folder. The template contains the {folderId} to be replaced by the folder ID house the new file.

	CreateUrl string `json:"createUrl,omitempty\"` // The URL to create a file with this app.

	HasDriveWideScope bool `json:"hasDriveWideScope,omitempty\"` // Whether the app has Drive-wide scope. An app with Drive-wide scope can access all files in the user's Drive.

	Icons []AppIcons `json:"icons,omitempty\"` // The various icons for the app.

	Id string `json:"id,omitempty\"` // The ID of the app.

	Installed bool `json:"installed,omitempty\"` // Whether the app is installed.

	Kind string `json:"kind,omitempty\"` // Output only. Identifies what kind of resource this is. Value: the fixed string "drive#app".

	LongDescription string `json:"longDescription,omitempty\"` // A long description of the app.

	Name string `json:"name,omitempty\"` // The name of the app.

	ObjectType string `json:"objectType,omitempty\"` // The type of object this app creates such as a Chart. If empty, the app name should be used instead.

	OpenUrlTemplate string `json:"openUrlTemplate,omitempty\"` // The template URL for opening files with this app. The template contains {ids} or {exportIds} to be replaced by the actual file IDs. For more information, see Open Files for the full documentation.

	PrimaryFileExtensions []string `json:"primaryFileExtensions,omitempty\"` // The list of primary file extensions.

	PrimaryMimeTypes []string `json:"primaryMimeTypes,omitempty\"` // The list of primary MIME types.

	ProductId string `json:"productId,omitempty\"` // The ID of the product listing for this app.

	ProductUrl string `json:"productUrl,omitempty\"` // A link to the product listing for this app.

	SecondaryFileExtensions []string `json:"secondaryFileExtensions,omitempty\"` // The list of secondary file extensions.

	SecondaryMimeTypes []string `json:"secondaryMimeTypes,omitempty\"` // The list of secondary MIME types.

	ShortDescription string `json:"shortDescription,omitempty\"` // A short description of the app.

	SupportsCreate bool `json:"supportsCreate,omitempty\"` // Whether this app supports creating objects.

	SupportsImport bool `json:"supportsImport,omitempty\"` // Whether this app supports importing from Google Docs.

	SupportsMultiOpen bool `json:"supportsMultiOpen,omitempty\"` // Whether this app supports opening more than one file.

	SupportsOfflineCreate bool `json:"supportsOfflineCreate,omitempty\"` // Whether this app supports creating files when offline.

	UseByDefault bool `json:"useByDefault,omitempty\"` // Whether the app is selected as the default handler for the types it supports.

}

type AppIcons struct {
	Category string `json:"category,omitempty\"` // Category of the icon. Allowed values are: * `application` - The icon for the application. * `document` - The icon for a file associated with the app. * `documentShared` - The icon for a shared file associated with the app.

	IconUrl string `json:"iconUrl,omitempty\"` // URL for the icon.

	Size int `json:"size,omitempty\"` // Size of the icon. Represented as the maximum of the width and height.

}

// A list of third-party applications which the user has installed or given access to Google Drive.
type AppList struct {
	DefaultAppIds []string `json:"defaultAppIds,omitempty\"` // The list of app IDs that the user has specified to use by default. The list is in reverse-priority order (lowest to highest).

	Items []App `json:"items,omitempty\"` // The list of apps.

	Kind string `json:"kind,omitempty\"` // Output only. Identifies what kind of resource this is. Value: the fixed string "drive#appList".

	SelfLink string `json:"selfLink,omitempty\"` // A link back to this list.

}

// Metadata for an approval. An approval is a review or approve process for a Drive item.
type Approval struct {
	ApprovalId string `json:"approvalId,omitempty\"` // The approval ID.

	CompleteTime string `json:"completeTime,omitempty\"` // Output only. The time the approval was completed.

	CreateTime string `json:"createTime,omitempty\"` // Output only. The time the approval was created.

	DueTime string `json:"dueTime,omitempty\"` // The time that the approval is due.

	FileContentChangeBehavior string `json:"fileContentChangeBehavior,omitempty\"` // Output only. The behavior of the approval when the file content changes.

	Initiator User `json:"initiator,omitempty\"` // The user that requested the approval.

	Kind string `json:"kind,omitempty\"` // This is always drive#approval.

	ModifyTime string `json:"modifyTime,omitempty\"` // Output only. The most recent time the approval was modified.

	ReviewerResponses []ReviewerResponse `json:"reviewerResponses,omitempty\"` // The responses made on the approval by reviewers.

	Status string `json:"status,omitempty\"` // Output only. The status of the approval at the time this resource was requested.

	TargetFileId string `json:"targetFileId,omitempty\"` // Target file id of the approval.

}

// The response of an approvals list request.
type ApprovalList struct {
	Items []Approval `json:"items,omitempty\"` // The list of approvals. If `nextPageToken` is populated, then this list may be incomplete and an additional page of results should be fetched.

	Kind string `json:"kind,omitempty\"` // This is always drive#approvalList

	NextPageToken string `json:"nextPageToken,omitempty\"` // The page token for the next page of approvals. This is absent if the end of the approvals list has been reached. If the token is rejected for any reason, it should be discarded, and pagination should be restarted from the first page of results.

}

// Request for approving an approval as a reviewer.
type ApproveApprovalRequest struct {
	Message string `json:"message,omitempty\"` // Optional. A message to accompany the reviewer response on the approval. This message is included in notifications for the action and in the approval activity log.

}

// Request for cancelling an approval as an initiator.
type CancelApprovalRequest struct {
	Message string `json:"message,omitempty\"` // Optional. A message to accompany the cancellation of the approval. This message is included in notifications for the action and in the approval activity log.

}

// A change to a file or shared drive.
type Change struct {
	ChangeType string `json:"changeType,omitempty\"` // The type of the change. Possible values are `file` and `drive`.

	Drive Drive `json:"drive,omitempty\"` // The updated state of the shared drive. Present if the changeType is drive, the user is still a member of the shared drive, and the shared drive has not been deleted.

	DriveId string `json:"driveId,omitempty\"` // The ID of the shared drive associated with this change.

	File File `json:"file,omitempty\"` // The updated state of the file. Present if the type is file and the file has not been removed from this list of changes.

	FileId string `json:"fileId,omitempty\"` // The ID of the file which has changed.

	Kind string `json:"kind,omitempty\"` // Identifies what kind of resource this is. Value: the fixed string `"drive#change"`.

	Removed bool `json:"removed,omitempty\"` // Whether the file or shared drive has been removed from this list of changes, for example by deletion or loss of access.

	TeamDrive TeamDrive `json:"teamDrive,omitempty\"` // Deprecated: Use `drive` instead.

	TeamDriveId string `json:"teamDriveId,omitempty\"` // Deprecated: Use `driveId` instead.

	Time time.Time `json:"time,omitempty\"` // The time of this change (RFC 3339 date-time).

	TypeValue string `json:"type,omitempty\"` // Deprecated: Use `changeType` instead.

}

// A list of changes for a user.
type ChangeList struct {
	Changes []Change `json:"changes,omitempty\"` // The list of changes. If nextPageToken is populated, then this list may be incomplete and an additional page of results should be fetched.

	Kind string `json:"kind,omitempty\"` // Identifies what kind of resource this is. Value: the fixed string `"drive#changeList"`.

	NewStartPageToken string `json:"newStartPageToken,omitempty\"` // The starting page token for future changes. This will be present only if the end of the current changes list has been reached. The page token doesn't expire.

	NextPageToken string `json:"nextPageToken,omitempty\"` // The page token for the next page of changes. This will be absent if the end of the changes list has been reached. The page token doesn't expire.

}

// A notification channel used to watch for resource changes.
type Channel struct {
	Address string `json:"address,omitempty\"` // The address where notifications are delivered for this channel.

	Expiration int64 `json:"expiration,omitempty\"` // Date and time of notification channel expiration, expressed as a Unix timestamp, in milliseconds. Optional.

	Id string `json:"id,omitempty\"` // A UUID or similar unique string that identifies this channel.

	Kind string `json:"kind,omitempty\"` // Identifies this as a notification channel used to watch for changes to a resource, which is `api#channel`.

	Params map[string]interface{} `json:"params,omitempty\"` // Additional parameters controlling delivery channel behavior. Optional.

	Payload bool `json:"payload,omitempty\"` // A Boolean value to indicate whether payload is wanted. Optional.

	ResourceId string `json:"resourceId,omitempty\"` // An opaque ID that identifies the resource being watched on this channel. Stable across different API versions.

	ResourceUri string `json:"resourceUri,omitempty\"` // A version-specific identifier for the watched resource.

	Token string `json:"token,omitempty\"` // An arbitrary string delivered to the target address with each notification delivered over this channel. Optional.

	TypeValue string `json:"type,omitempty\"` // The type of delivery mechanism used for this channel. Valid values are "web_hook" or "webhook".

}

// Details about the client-side encryption applied to the file.
type ClientEncryptionDetails struct {
	DecryptionMetadata DecryptionMetadata `json:"decryptionMetadata,omitempty\"` // The metadata used for client-side operations.

	EncryptionState string `json:"encryptionState,omitempty\"` // The encryption state of the file. The values expected here are: - encrypted - unencrypted

}

// A comment on a file. Some resource methods (such as `comments.update`) require a `commentId`. Use the `comments.list` method to retrieve the ID for a comment in a file.
type Comment struct {
	Anchor string `json:"anchor,omitempty\"` // A region of the document represented as a JSON string. For details on defining anchor properties, refer to [Manage comments and replies](https://developers.google.com/workspace/drive/api/v3/manage-comments).

	AssigneeEmailAddress string `json:"assigneeEmailAddress,omitempty\"` // Output only. The email address of the user assigned to this comment. If no user is assigned, the field is unset.

	Author User `json:"author,omitempty\"` // Output only. The author of the comment. The author's email address and permission ID will not be populated.

	Content string `json:"content,omitempty\"` // The plain text content of the comment. This field is used for setting the content, while `htmlContent` should be displayed.

	CreatedTime time.Time `json:"createdTime,omitempty\"` // The time at which the comment was created (RFC 3339 date-time).

	Deleted bool `json:"deleted,omitempty\"` // Output only. Whether the comment has been deleted. A deleted comment has no content.

	HtmlContent string `json:"htmlContent,omitempty\"` // Output only. The content of the comment with HTML formatting.

	Id string `json:"id,omitempty\"` // Output only. The ID of the comment.

	Kind string `json:"kind,omitempty\"` // Output only. Identifies what kind of resource this is. Value: the fixed string `"drive#comment"`.

	MentionedEmailAddresses []string `json:"mentionedEmailAddresses,omitempty\"` // Output only. A list of email addresses for users mentioned in this comment. If no users are mentioned, the list is empty.

	ModifiedTime time.Time `json:"modifiedTime,omitempty\"` // The last time the comment or any of its replies was modified (RFC 3339 date-time).

	QuotedFileContent map[string]interface{} `json:"quotedFileContent,omitempty\"` // The file content to which the comment refers, typically within the anchor region. For a text file, for example, this would be the text at the location of the comment.

	Replies []Reply `json:"replies,omitempty\"` // Output only. The full list of replies to the comment in chronological order.

	Resolved bool `json:"resolved,omitempty\"` // Output only. Whether the comment has been resolved by one of its replies.

}

// Request for commenting on an approval.
type CommentApprovalRequest struct {
	Message string `json:"message,omitempty\"` // Required. A message to comment on the approval. This message is included in notifications for the action and in the approval activity log.

}

// A list of comments on a file.
type CommentList struct {
	Comments []Comment `json:"comments,omitempty\"` // The list of comments. If nextPageToken is populated, then this list may be incomplete and an additional page of results should be fetched.

	Kind string `json:"kind,omitempty\"` // Identifies what kind of resource this is. Value: the fixed string `"drive#commentList"`.

	NextPageToken string `json:"nextPageToken,omitempty\"` // The page token for the next page of comments. This will be absent if the end of the comments list has been reached. If the token is rejected for any reason, it should be discarded, and pagination should be restarted from the first page of results. The page token is typically valid for several hours. However, if new items are added or removed, your expected results might differ.

}

// A restriction for accessing the content of the file.
type ContentRestriction struct {
	OwnerRestricted bool `json:"ownerRestricted,omitempty\"` // Whether the content restriction can only be modified or removed by a user who owns the file. For files in shared drives, any user with `organizer` capabilities can modify or remove this content restriction.

	ReadOnly bool `json:"readOnly,omitempty\"` // Whether the content of the file is read-only. If a file is read-only, a new revision of the file may not be added, comments may not be added or modified, and the title of the file may not be modified.

	Reason string `json:"reason,omitempty\"` // Reason for why the content of the file is restricted. This is only mutable on requests that also set `readOnly=true`.

	RestrictingUser User `json:"restrictingUser,omitempty\"` // Output only. The user who set the content restriction. Only populated if `readOnly=true`.

	RestrictionTime time.Time `json:"restrictionTime,omitempty\"` // The time at which the content restriction was set (formatted RFC 3339 timestamp). Only populated if readOnly is true.

	SystemRestricted bool `json:"systemRestricted,omitempty\"` // Output only. Whether the content restriction was applied by the system, for example due to an esignature. Users cannot modify or remove system restricted content restrictions.

	TypeValue string `json:"type,omitempty\"` // Output only. The type of the content restriction. Currently the only possible value is `globalContentRestriction`.

}

// Request for declining an approval as a reviewer.
type DeclineApprovalRequest struct {
	Message string `json:"message,omitempty\"` // Optional. A message to accompany the reviewer response on the approval. This message is included in notifications for the action and in the approval activity log.

}

// Representation of the CSE DecryptionMetadata.
type DecryptionMetadata struct {
	Aes256GcmChunkSize string `json:"aes256GcmChunkSize,omitempty\"` // Chunk size used if content was encrypted with the AES 256 GCM Cipher. Possible values are: - default - small

	EncryptionResourceKeyHash string `json:"encryptionResourceKeyHash,omitempty\"` // The URL-safe Base64 encoded HMAC-SHA256 digest of the resource metadata with its DEK (Data Encryption Key); see https://developers.google.com/workspace/cse/reference

	Jwt string `json:"jwt,omitempty\"` // The signed JSON Web Token (JWT) which can be used to authorize the requesting user with the Key ACL Service (KACLS). The JWT asserts that the requesting user has at least read permissions on the file.

	KaclsId int64 `json:"kaclsId,omitempty\"` // The ID of the KACLS (Key ACL Service) used to encrypt the file.

	KaclsName string `json:"kaclsName,omitempty\"` // The name of the KACLS (Key ACL Service) used to encrypt the file.

	KeyFormat string `json:"keyFormat,omitempty\"` // Key format for the unwrapped key. Must be `tinkAesGcmKey`.

	WrappedKey string `json:"wrappedKey,omitempty\"` // The URL-safe Base64 encoded wrapped key used to encrypt the contents of the file.

}

// A restriction for copy and download of the file.
type DownloadRestriction struct {
	RestrictedForReaders bool `json:"restrictedForReaders,omitempty\"` // Whether download and copy is restricted for readers.

	RestrictedForWriters bool `json:"restrictedForWriters,omitempty\"` // Whether download and copy is restricted for writers. If true, download is also restricted for readers.

}

// Download restrictions applied to the file.
type DownloadRestrictionsMetadata struct {
	EffectiveDownloadRestrictionWithContext DownloadRestriction `json:"effectiveDownloadRestrictionWithContext,omitempty\"` // Output only. The effective download restriction applied to this file. This considers all restriction settings and DLP rules.

	ItemDownloadRestriction DownloadRestriction `json:"itemDownloadRestriction,omitempty\"` // The download restriction of the file applied directly by the owner or organizer. This doesn't take into account shared drive settings or DLP rules.

}

// Representation of a shared drive. Some resource methods (such as `drives.update`) require a `driveId`. Use the `drives.list` method to retrieve the ID for a shared drive.
type Drive struct {
	BackgroundImageFile map[string]interface{} `json:"backgroundImageFile,omitempty\"` // An image file and cropping parameters from which a background image for this shared drive is set. This is a write only field; it can only be set on `drive.drives.update` requests that don't set `themeId`. When specified, all fields of the `backgroundImageFile` must be set.

	BackgroundImageLink string `json:"backgroundImageLink,omitempty\"` // Output only. A short-lived link to this shared drive's background image.

	Capabilities map[string]interface{} `json:"capabilities,omitempty\"` // Output only. Capabilities the current user has on this shared drive.

	ColorRgb string `json:"colorRgb,omitempty\"` // The color of this shared drive as an RGB hex string. It can only be set on a `drive.drives.update` request that does not set `themeId`.

	CreatedTime time.Time `json:"createdTime,omitempty\"` // The time at which the shared drive was created (RFC 3339 date-time).

	Hidden bool `json:"hidden,omitempty\"` // Whether the shared drive is hidden from default view.

	Id string `json:"id,omitempty\"` // Output only. The ID of this shared drive which is also the ID of the top level folder of this shared drive.

	Kind string `json:"kind,omitempty\"` // Output only. Identifies what kind of resource this is. Value: the fixed string `"drive#drive"`.

	Name string `json:"name,omitempty\"` // The name of this shared drive.

	OrgUnitId string `json:"orgUnitId,omitempty\"` // Output only. The organizational unit of this shared drive. This field is only populated on `drives.list` responses when the `useDomainAdminAccess` parameter is set to `true`.

	Restrictions map[string]interface{} `json:"restrictions,omitempty\"` // A set of restrictions that apply to this shared drive or items inside this shared drive. Note that restrictions can't be set when creating a shared drive. To add a restriction, first create a shared drive and then use `drives.update` to add restrictions.

	ThemeId string `json:"themeId,omitempty\"` // The ID of the theme from which the background image and color will be set. The set of possible `driveThemes` can be retrieved from a `drive.about.get` response. When not specified on a `drive.drives.create` request, a random theme is chosen from which the background image and color are set. This is a write-only field; it can only be set on requests that don't set `colorRgb` or `backgroundImageFile`.

}

// A list of shared drives.
type DriveList struct {
	Drives []Drive `json:"drives,omitempty\"` // The list of shared drives. If nextPageToken is populated, then this list may be incomplete and an additional page of results should be fetched.

	Kind string `json:"kind,omitempty\"` // Identifies what kind of resource this is. Value: the fixed string `"drive#driveList"`.

	NextPageToken string `json:"nextPageToken,omitempty\"` // The page token for the next page of shared drives. This will be absent if the end of the list has been reached. If the token is rejected for any reason, it should be discarded, and pagination should be restarted from the first page of results. The page token is typically valid for several hours. However, if new items are added or removed, your expected results might differ.

}

// The metadata for a file. Some resource methods (such as `files.update`) require a `fileId`. Use the `files.list` method to retrieve the ID for a file.
type File struct {
	AppProperties map[string]interface{} `json:"appProperties,omitempty\"` // A collection of arbitrary key-value pairs which are private to the requesting app.
	// Entries with null values are cleared in update and copy requests. These properties can only be retrieved using an authenticated request. An authenticated request uses an access token obtained with a OAuth 2 client ID. You cannot use an API key to retrieve private properties.

	Capabilities map[string]interface{} `json:"capabilities,omitempty\"` // Output only. Capabilities the current user has on this file. Each capability corresponds to a fine-grained action that a user may take. For more information, see [Understand file capabilities](https://developers.google.com/workspace/drive/api/guides/manage-sharing#capabilities).

	ClientEncryptionDetails ClientEncryptionDetails `json:"clientEncryptionDetails,omitempty\"` // Client Side Encryption related details. Contains details about the encryption state of the file and details regarding the encryption mechanism that clients need to use when decrypting the contents of this item. This will only be present on files and not on folders or shortcuts.

	ContentHints map[string]interface{} `json:"contentHints,omitempty\"` // Additional information about the content of the file. These fields are never populated in responses.

	ContentRestrictions []ContentRestriction `json:"contentRestrictions,omitempty\"` // Restrictions for accessing the content of the file. Only populated if such a restriction exists.

	CopyRequiresWriterPermission bool `json:"copyRequiresWriterPermission,omitempty\"` // Whether the options to copy, print, or download this file should be disabled for readers and commenters.

	CreatedTime time.Time `json:"createdTime,omitempty\"` // The time at which the file was created (RFC 3339 date-time).

	Description string `json:"description,omitempty\"` // A short description of the file.

	DownloadRestrictions DownloadRestrictionsMetadata `json:"downloadRestrictions,omitempty\"` // Download restrictions applied on the file.

	DriveId string `json:"driveId,omitempty\"` // Output only. ID of the shared drive the file resides in. Only populated for items in shared drives.

	ExplicitlyTrashed bool `json:"explicitlyTrashed,omitempty\"` // Output only. Whether the file has been explicitly trashed, as opposed to recursively trashed from a parent folder.

	ExportLinks map[string]interface{} `json:"exportLinks,omitempty\"` // Output only. Links for exporting Docs Editors files to specific formats.

	FileExtension string `json:"fileExtension,omitempty\"` // Output only. The final component of `fullFileExtension`. This is only available for files with binary content in Google Drive.

	FolderColorRgb string `json:"folderColorRgb,omitempty\"` // The color for a folder or a shortcut to a folder as an RGB hex string. The supported colors are published in the `folderColorPalette` field of the [`about`](/workspace/drive/api/reference/rest/v3/about) resource. If an unsupported color is specified, the closest color in the palette is used instead.

	FullFileExtension string `json:"fullFileExtension,omitempty\"` // Output only. The full file extension extracted from the `name` field. May contain multiple concatenated extensions, such as "tar.gz". This is only available for files with binary content in Google Drive. This is automatically updated when the `name` field changes, however it's not cleared if the new name doesn't contain a valid extension.

	HasAugmentedPermissions bool `json:"hasAugmentedPermissions,omitempty\"` // Output only. Whether there are permissions directly on this file. This field is only populated for items in shared drives.

	HasThumbnail bool `json:"hasThumbnail,omitempty\"` // Output only. Whether this file has a thumbnail. This doesn't indicate whether the requesting app has access to the thumbnail. To check access, look for the presence of the thumbnailLink field.

	HeadRevisionId string `json:"headRevisionId,omitempty\"` // Output only. The ID of the file's head revision. This is currently only available for files with binary content in Google Drive.

	IconLink string `json:"iconLink,omitempty\"` // Output only. A static, unauthenticated link to the file's icon.

	Id string `json:"id,omitempty\"` // The ID of the file.

	ImageMediaMetadata map[string]interface{} `json:"imageMediaMetadata,omitempty\"` // Output only. Additional metadata about image media, if available.

	InheritedPermissionsDisabled bool `json:"inheritedPermissionsDisabled,omitempty\"` // Whether this file has inherited permissions disabled. Inherited permissions are enabled by default.

	IsAppAuthorized bool `json:"isAppAuthorized,omitempty\"` // Output only. Whether the file was created or opened by the requesting app.

	Kind string `json:"kind,omitempty\"` // Output only. Identifies what kind of resource this is. Value: the fixed string `"drive#file"`.

	LabelInfo map[string]interface{} `json:"labelInfo,omitempty\"` // Label information on the file.

	LastModifyingUser User `json:"lastModifyingUser,omitempty\"` // Output only. The last user to modify the file. This field is only populated when the last modification was performed by a signed-in user.

	LinkShareMetadata map[string]interface{} `json:"linkShareMetadata,omitempty\"` // Contains details about the link URLs that clients are using to refer to this item.

	Md5Checksum string `json:"md5Checksum,omitempty\"` // Output only. The MD5 checksum for the content of the file. This is only applicable to files with binary content in Google Drive.

	MimeType string `json:"mimeType,omitempty\"` // The MIME type of the file. Google Drive attempts to automatically detect an appropriate value from uploaded content, if no value is provided. The value cannot be changed unless a new revision is uploaded. If a file is created with a Google Doc MIME type, the uploaded content is imported, if possible. The supported import formats are published in the [`about`](/workspace/drive/api/reference/rest/v3/about) resource.

	ModifiedByMe bool `json:"modifiedByMe,omitempty\"` // Output only. Whether the file has been modified by this user.

	ModifiedByMeTime time.Time `json:"modifiedByMeTime,omitempty\"` // The last time the file was modified by the user (RFC 3339 date-time).

	ModifiedTime time.Time `json:"modifiedTime,omitempty\"` // he last time the file was modified by anyone (RFC 3339 date-time). Note that setting modifiedTime will also update modifiedByMeTime for the user.

	Name string `json:"name,omitempty\"` // The name of the file. This isn't necessarily unique within a folder. Note that for immutable items such as the top-level folders of shared drives, the My Drive root folder, and the Application Data folder, the name is constant.

	OriginalFilename string `json:"originalFilename,omitempty\"` // The original filename of the uploaded content if available, or else the original value of the `name` field. This is only available for files with binary content in Google Drive.

	OwnedByMe bool `json:"ownedByMe,omitempty\"` // Output only. Whether the user owns the file. Not populated for items in shared drives.

	Owners []User `json:"owners,omitempty\"` // Output only. The owner of this file. Only certain legacy files may have more than one owner. This field isn't populated for items in shared drives.

	Parents []string `json:"parents,omitempty\"` // The ID of the parent folder containing the file. A file can only have one parent folder; specifying multiple parents isn't supported. If not specified as part of a create request, the file is placed directly in the user's My Drive folder. If not specified as part of a copy request, the file inherits any discoverable parent of the source file. Update requests must use the `addParents` and `removeParents` parameters to modify the parents list.

	PermissionIds []string `json:"permissionIds,omitempty\"` // Output only. List of permission IDs for users with access to this file.

	Permissions []Permission `json:"permissions,omitempty\"` // Output only. The full list of permissions for the file. This is only available if the requesting user can share the file. Not populated for items in shared drives.

	Properties map[string]interface{} `json:"properties,omitempty\"` // A collection of arbitrary key-value pairs which are visible to all apps.
	// Entries with null values are cleared in update and copy requests.

	QuotaBytesUsed int64 `json:"quotaBytesUsed,omitempty\"` // Output only. The number of storage quota bytes used by the file. This includes the head revision as well as previous revisions with `keepForever` enabled.

	ResourceKey string `json:"resourceKey,omitempty\"` // Output only. A key needed to access the item via a shared link.

	Sha1Checksum string `json:"sha1Checksum,omitempty\"` // Output only. The SHA1 checksum associated with this file, if available. This field is only populated for files with content stored in Google Drive; it's not populated for Docs Editors or shortcut files.

	Sha256Checksum string `json:"sha256Checksum,omitempty\"` // Output only. The SHA256 checksum associated with this file, if available. This field is only populated for files with content stored in Google Drive; it's not populated for Docs Editors or shortcut files.

	Shared bool `json:"shared,omitempty\"` // Output only. Whether the file has been shared. Not populated for items in shared drives.

	SharedWithMeTime time.Time `json:"sharedWithMeTime,omitempty\"` // The time at which the file was shared with the user, if applicable (RFC 3339 date-time).

	SharingUser User `json:"sharingUser,omitempty\"` // Output only. The user who shared the file with the requesting user, if applicable.

	ShortcutDetails map[string]interface{} `json:"shortcutDetails,omitempty\"` // Information about a shortcut file.

	Size int64 `json:"size,omitempty\"` // Output only. Size in bytes of blobs and Google Workspace editor files. Won't be populated for files that have no size, like shortcuts and folders.

	Spaces []string `json:"spaces,omitempty\"` // Output only. The list of spaces which contain the file. The currently supported values are `drive`, `appDataFolder`, and `photos`.

	Starred bool `json:"starred,omitempty\"` // Whether the user has starred the file.

	TeamDriveId string `json:"teamDriveId,omitempty\"` // Deprecated: Output only. Use `driveId` instead.

	ThumbnailLink string `json:"thumbnailLink,omitempty\"` // Output only. A short-lived link to the file's thumbnail, if available. Typically lasts on the order of hours. Not intended for direct usage on web applications due to [Cross-Origin Resource Sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) policies. Consider using a proxy server. Only populated when the requesting app can access the file's content. If the file isn't shared publicly, the URL returned in `files.thumbnailLink` must be fetched using a credentialed request.

	ThumbnailVersion int64 `json:"thumbnailVersion,omitempty\"` // Output only. The thumbnail version for use in thumbnail cache invalidation.

	Trashed bool `json:"trashed,omitempty\"` // Whether the file has been trashed, either explicitly or from a trashed parent folder. Only the owner may trash a file, but other users can still access the file in the owner's trash until it's permanently deleted.

	TrashedTime time.Time `json:"trashedTime,omitempty\"` // The time that the item was trashed (RFC 3339 date-time). Only populated for items in shared drives.

	TrashingUser User `json:"trashingUser,omitempty\"` // Output only. If the file has been explicitly trashed, the user who trashed it. Only populated for items in shared drives.

	Version int64 `json:"version,omitempty\"` // Output only. A monotonically increasing version number for the file. This reflects every change made to the file on the server, even those not visible to the user.

	VideoMediaMetadata map[string]interface{} `json:"videoMediaMetadata,omitempty\"` // Output only. Additional metadata about video media. This may not be available immediately upon upload.

	ViewedByMe bool `json:"viewedByMe,omitempty\"` // Output only. Whether the file has been viewed by this user.

	ViewedByMeTime time.Time `json:"viewedByMeTime,omitempty\"` // The last time the file was viewed by the user (RFC 3339 date-time).

	ViewersCanCopyContent bool `json:"viewersCanCopyContent,omitempty\"` // Deprecated: Use `copyRequiresWriterPermission` instead.

	WebContentLink string `json:"webContentLink,omitempty\"` // Output only. A link for downloading the content of the file in a browser. This is only available for files with binary content in Google Drive.

	WebViewLink string `json:"webViewLink,omitempty\"` // Output only. A link for opening the file in a relevant Google editor or viewer in a browser.

	WritersCanShare bool `json:"writersCanShare,omitempty\"` // Whether users with only `writer` permission can modify the file's permissions. Not populated for items in shared drives.

}

// A list of files.
type FileList struct {
	Files []File `json:"files,omitempty\"` // The list of files. If `nextPageToken` is populated, then this list may be incomplete and an additional page of results should be fetched.

	IncompleteSearch bool `json:"incompleteSearch,omitempty\"` // Whether the search process was incomplete. If true, then some search results might be missing, since all documents were not searched. This can occur when searching multiple drives with the `allDrives` corpora, but all corpora couldn't be searched. When this happens, it's suggested that clients narrow their query by choosing a different corpus such as `user` or `drive`.

	Kind string `json:"kind,omitempty\"` // Identifies what kind of resource this is. Value: the fixed string `"drive#fileList"`.

	NextPageToken string `json:"nextPageToken,omitempty\"` // The page token for the next page of files. This will be absent if the end of the files list has been reached. If the token is rejected for any reason, it should be discarded, and pagination should be restarted from the first page of results. The page token is typically valid for several hours. However, if new items are added or removed, your expected results might differ.

}

// JWT and associated metadata used to generate CSE files.
type GenerateCseTokenResponse struct {
	CurrentKaclsId int64 `json:"currentKaclsId,omitempty\"` // The current Key ACL Service (KACLS) ID associated with the JWT.

	CurrentKaclsName string `json:"currentKaclsName,omitempty\"` // Name of the KACLs that the returned KACLs ID points to.

	FileId string `json:"fileId,omitempty\"` // The fileId for which the JWT was generated.

	Jwt string `json:"jwt,omitempty\"` // The signed JSON Web Token (JWT) for the file.

	Kind string `json:"kind,omitempty\"` // Output only. Identifies what kind of resource this is. Value: the fixed string `"drive#generateCseTokenResponse"`.

}

// A list of generated file IDs which can be provided in create requests.
type GeneratedIds struct {
	Ids []string `json:"ids,omitempty\"` // The IDs generated for the requesting user in the specified space.

	Kind string `json:"kind,omitempty\"` // Identifies what kind of resource this is. Value: the fixed string `"drive#generatedIds"`.

	Space string `json:"space,omitempty\"` // The type of file that can be created with these IDs.

}

// Representation of label and label fields.
type Label struct {
	Fields map[string]interface{} `json:"fields,omitempty\"` // A map of the fields on the label, keyed by the field's ID.

	Id string `json:"id,omitempty\"` // The ID of the label.

	Kind string `json:"kind,omitempty\"` // This is always drive#label

	RevisionId string `json:"revisionId,omitempty\"` // The revision ID of the label.

}

// Representation of field, which is a typed key-value pair.
type LabelField struct {
	DateString []string `json:"dateString,omitempty\"` // Only present if valueType is dateString. RFC 3339 formatted date: YYYY-MM-DD.

	Id string `json:"id,omitempty\"` // The identifier of this label field.

	Integer []int64 `json:"integer,omitempty\"` // Only present if `valueType` is `integer`.

	Kind string `json:"kind,omitempty\"` // This is always drive#labelField.

	Selection []string `json:"selection,omitempty\"` // Only present if `valueType` is `selection`

	Text []string `json:"text,omitempty\"` // Only present if `valueType` is `text`.

	User []User `json:"user,omitempty\"` // Only present if `valueType` is `user`.

	ValueType string `json:"valueType,omitempty\"` // The field type. While new values may be supported in the future, the following are currently allowed: * `dateString` * `integer` * `selection` * `text` * `user`

}

// A modification to a label's field.
type LabelFieldModification struct {
	FieldId string `json:"fieldId,omitempty\"` // The ID of the field to be modified.

	Kind string `json:"kind,omitempty\"` // This is always `"drive#labelFieldModification"`.

	SetDateValues []string `json:"setDateValues,omitempty\"` // Replaces the value of a dateString Field with these new values. The string must be in the RFC 3339 full-date format: YYYY-MM-DD.

	SetIntegerValues []int64 `json:"setIntegerValues,omitempty\"` // Replaces the value of an `integer` field with these new values.

	SetSelectionValues []string `json:"setSelectionValues,omitempty\"` // Replaces a `selection` field with these new values.

	SetTextValues []string `json:"setTextValues,omitempty\"` // Sets the value of a `text` field.

	SetUserValues []string `json:"setUserValues,omitempty\"` // Replaces a `user` field with these new values. The values must be a valid email addresses.

	UnsetValues bool `json:"unsetValues,omitempty\"` // Unsets the values for this field.

}

// A list of labels applied to a file.
type LabelList struct {
	Kind string `json:"kind,omitempty\"` // This is always `"drive#labelList"`.

	Labels []Label `json:"labels,omitempty\"` // The list of labels.

	NextPageToken string `json:"nextPageToken,omitempty\"` // The page token for the next page of labels. This field will be absent if the end of the list has been reached. If the token is rejected for any reason, it should be discarded, and pagination should be restarted from the first page of results. The page token is typically valid for several hours. However, if new items are added or removed, your expected results might differ.

}

// A modification to a label on a file. A `LabelModification` can be used to apply a label to a file, update an existing label on a file, or remove a label from a file.
type LabelModification struct {
	FieldModifications []LabelFieldModification `json:"fieldModifications,omitempty\"` // The list of modifications to this label's fields.

	Kind string `json:"kind,omitempty\"` // This is always `"drive#labelModification"`.

	LabelId string `json:"labelId,omitempty\"` // The ID of the label to modify.

	RemoveLabel bool `json:"removeLabel,omitempty\"` // If true, the label will be removed from the file.

}

// The response to an access proposal list request.
type ListAccessProposalsResponse struct {
	AccessProposals []AccessProposal `json:"accessProposals,omitempty\"` // The list of access proposals. This field is only populated in Drive API v3.

	NextPageToken string `json:"nextPageToken,omitempty\"` // The continuation token for the next page of results. This will be absent if the end of the results list has been reached. If the token is rejected for any reason, it should be discarded, and pagination should be restarted from the first page of results.

}

// A request to modify the set of labels on a file. This request may contain many modifications that will either all succeed or all fail atomically.
type ModifyLabelsRequest struct {
	Kind string `json:"kind,omitempty\"` // This is always `"drive#modifyLabelsRequest"`.

	LabelModifications []LabelModification `json:"labelModifications,omitempty\"` // The list of modifications to apply to the labels on the file.

}

// Response to a `ModifyLabels` request. This contains only those labels which were added or updated by the request.
type ModifyLabelsResponse struct {
	Kind string `json:"kind,omitempty\"` // This is always `"drive#modifyLabelsResponse"`.

	ModifiedLabels []Label `json:"modifiedLabels,omitempty\"` // The list of labels which were added or updated by the request.

}

// This resource represents a long-running operation that is the result of a network API call.
type Operation struct {
	Done bool `json:"done,omitempty\"` // If the value is `false`, it means the operation is still in progress. If `true`, the operation is completed, and either `error` or `response` is available.

	Error Status `json:"error,omitempty\"` // The error result of the operation in case of failure or cancellation.

	Metadata map[string]interface{} `json:"metadata,omitempty\"` // Service-specific metadata associated with the operation. It typically contains progress information and common metadata such as create time. Some services might not provide such metadata. Any method that returns a long-running operation should document the metadata type, if any.

	Name string `json:"name,omitempty\"` // The server-assigned name, which is only unique within the same service that originally returns it. If you use the default HTTP mapping, the `name` should be a resource name ending with `operations/{unique_id}`.

	Response map[string]interface{} `json:"response,omitempty\"` // The normal, successful response of the operation. If the original method returns no data on success, such as `Delete`, the response is `google.protobuf.Empty`. If the original method is standard `Get`/`Create`/`Update`, the response should be the resource. For other methods, the response should have the type `XxxResponse`, where `Xxx` is the original method name. For example, if the original method name is `TakeSnapshot()`, the inferred response type is `TakeSnapshotResponse`.

}

// A permission for a file. A permission grants a user, group, domain, or the world access to a file or a folder hierarchy. For more information, see [Share files, folders, and drives](https://developers.google.com/workspace/drive/api/guides/manage-sharing). By default, permission requests only return a subset of fields. Permission `kind`, `ID`, `type`, and `role` are always returned. To retrieve specific fields, see [Return specific fields](https://developers.google.com/workspace/drive/api/guides/fields-parameter). Some resource methods (such as `permissions.update`) require a `permissionId`. Use the `permissions.list` method to retrieve the ID for a file, folder, or shared drive.
type Permission struct {
	AllowFileDiscovery bool `json:"allowFileDiscovery,omitempty\"` // Whether the permission allows the file to be discovered through search. This is only applicable for permissions of type `domain` or `anyone`.

	Deleted bool `json:"deleted,omitempty\"` // Output only. Whether the account associated with this permission has been deleted. This field only pertains to permissions of type `user` or `group`.

	DisplayName string `json:"displayName,omitempty\"` // Output only. The "pretty" name of the value of the permission. The following is a list of examples for each type of permission: * `user` - User's full name, as defined for their Google Account, such as "Dana A." * `group` - Name of the Google Group, such as "The Company Administrators." * `domain` - String domain name, such as "cymbalgroup.com." * `anyone` - No `displayName` is present.

	Domain string `json:"domain,omitempty\"` // Output only. The domain to which this permission refers.

	EmailAddress string `json:"emailAddress,omitempty\"` // Output only. The email address of the user or group to which this permission refers.

	ExpirationTime time.Time `json:"expirationTime,omitempty\"` // The time at which this permission will expire (RFC 3339 date-time). Expiration times have the following restrictions: - They can only be set on user and group permissions - The time must be in the future - The time cannot be more than a year in the future

	Id string `json:"id,omitempty\"` // Output only. The ID of this permission. This is a unique identifier for the grantee, and is published in the [User resource](https://developers.google.com/workspace/drive/api/reference/rest/v3/User) as `permissionId`. IDs should be treated as opaque values.

	InheritedPermissionsDisabled bool `json:"inheritedPermissionsDisabled,omitempty\"` // When `true`, only organizers, owners, and users with permissions added directly on the item can access it.

	Kind string `json:"kind,omitempty\"` // Output only. Identifies what kind of resource this is. Value: the fixed string `"drive#permission"`.

	PendingOwner bool `json:"pendingOwner,omitempty\"` // Whether the account associated with this permission is a pending owner. Only populated for permissions of type `user` for files that aren't in a shared drive.

	PermissionDetails []map[string]interface{} `json:"permissionDetails,omitempty\"` // Output only. Details of whether the permissions on this item are inherited or are directly on this item.

	PhotoLink string `json:"photoLink,omitempty\"` // Output only. A link to the user's profile photo, if available.

	Role string `json:"role,omitempty\"` // The role granted by this permission. Supported values include: * `owner` * `organizer` * `fileOrganizer` * `writer` * `commenter` * `reader` For more information, see [Roles and permissions](https://developers.google.com/workspace/drive/api/guides/ref-roles).

	TeamDrivePermissionDetails []map[string]interface{} `json:"teamDrivePermissionDetails,omitempty\"` // Output only. Deprecated: Output only. Use `permissionDetails` instead.

	TypeValue string `json:"type,omitempty\"` // The type of the grantee. Supported values include: * `user` * `group` * `domain` * `anyone` When creating a permission, if `type` is `user` or `group`, you must provide an `emailAddress` for the user or group. If `type` is `domain`, you must provide a `domain`. If `type` is `anyone`, no extra information is required.

	View string `json:"view,omitempty\"` // Indicates the view for this permission. Only populated for permissions that belong to a view. The only supported values are `published` and `metadata`: * `published`: The permission's role is `publishedReader`. * `metadata`: The item is only visible to the `metadata` view because the item has limited access and the scope has at least read access to the parent. The `metadata` view is only supported on folders. For more information, see [Views](https://developers.google.com/workspace/drive/api/guides/ref-roles#views).

}

// A list of permissions for a file.
type PermissionList struct {
	Kind string `json:"kind,omitempty\"` // Identifies what kind of resource this is. Value: the fixed string `"drive#permissionList"`.

	NextPageToken string `json:"nextPageToken,omitempty\"` // The page token for the next page of permissions. This field will be absent if the end of the permissions list has been reached. If the token is rejected for any reason, it should be discarded, and pagination should be restarted from the first page of results. The page token is typically valid for several hours. However, if new items are added or removed, your expected results might differ.

	Permissions []Permission `json:"permissions,omitempty\"` // The list of permissions. If `nextPageToken` is populated, then this list may be incomplete and an additional page of results should be fetched.

}

// Request for reassigning an approval. Reviewers can be added or replaced, but not removed.
type ReassignApprovalRequest struct {
	AddReviewers []AddReviewer `json:"addReviewers,omitempty\"` // Optional. The list of reviewers to add.

	Message string `json:"message,omitempty\"` // Optional. A message to send to the new reviewers. This message is included in notifications for the action and in the approval activity log.

	ReplaceReviewers []ReplaceReviewer `json:"replaceReviewers,omitempty\"` // Optional. The list of reviewer replacements.

}

// Representation of a reviewer replacement.
type ReplaceReviewer struct {
	AddedReviewerEmail string `json:"addedReviewerEmail,omitempty\"` // Required. The email of the reviewer to add.

	RemovedReviewerEmail string `json:"removedReviewerEmail,omitempty\"` // Required. The email of the reviewer to remove.

}

// A reply to a comment on a file. Some resource methods (such as `replies.update`) require a `replyId`. Use the `replies.list` method to retrieve the ID for a reply.
type Reply struct {
	Action string `json:"action,omitempty\"` // The action the reply performed to the parent comment. The supported values are: * `resolve` * `reopen`

	AssigneeEmailAddress string `json:"assigneeEmailAddress,omitempty\"` // Output only. The email address of the user assigned to this comment. If no user is assigned, the field is unset.

	Author User `json:"author,omitempty\"` // Output only. The author of the reply. The author's email address and permission ID won't be populated.

	Content string `json:"content,omitempty\"` // The plain text content of the reply. This field is used for setting the content, while `htmlContent` should be displayed. This field is required by the `create` method if no `action` value is specified.

	CreatedTime time.Time `json:"createdTime,omitempty\"` // The time at which the reply was created (RFC 3339 date-time).

	Deleted bool `json:"deleted,omitempty\"` // Output only. Whether the reply has been deleted. A deleted reply has no content.

	HtmlContent string `json:"htmlContent,omitempty\"` // Output only. The content of the reply with HTML formatting.

	Id string `json:"id,omitempty\"` // Output only. The ID of the reply.

	Kind string `json:"kind,omitempty\"` // Output only. Identifies what kind of resource this is. Value: the fixed string `"drive#reply"`.

	MentionedEmailAddresses []string `json:"mentionedEmailAddresses,omitempty\"` // Output only. A list of email addresses for users mentioned in this comment. If no users are mentioned, the list is empty.

	ModifiedTime time.Time `json:"modifiedTime,omitempty\"` // The last time the reply was modified (RFC 3339 date-time).

}

// A list of replies to a comment on a file.
type ReplyList struct {
	Kind string `json:"kind,omitempty\"` // Identifies what kind of resource this is. Value: the fixed string `"drive#replyList"`.

	NextPageToken string `json:"nextPageToken,omitempty\"` // The page token for the next page of replies. This will be absent if the end of the replies list has been reached. If the token is rejected for any reason, it should be discarded, and pagination should be restarted from the first page of results. The page token is typically valid for several hours. However, if new items are added or removed, your expected results might differ.

	Replies []Reply `json:"replies,omitempty\"` // The list of replies. If `nextPageToken` is populated, then this list may be incomplete and an additional page of results should be fetched.

}

// Request message for resolving an AccessProposal on a file.
type ResolveAccessProposalRequest struct {
	Action string `json:"action,omitempty\"` // Required. The action to take on the access proposal.

	Role []string `json:"role,omitempty\"` // Optional. The roles that the approver has allowed, if any. For more information, see [Roles and permissions](https://developers.google.com/workspace/drive/api/guides/ref-roles). Note: This field is required for the `ACCEPT` action.

	SendNotification bool `json:"sendNotification,omitempty\"` // Optional. Whether to send an email to the requester when the access proposal is denied or accepted.

	View string `json:"view,omitempty\"` // Optional. Indicates the view for this access proposal. This should only be set when the proposal belongs to a view. Only `published` is supported.

}

// A response on an approval made by a specific reviewer.
type ReviewerResponse struct {
	Kind string `json:"kind,omitempty\"` // This is always drive#reviewerResponse.

	Response string `json:"response,omitempty\"` // A reviewer’s response for the approval.

	Reviewer User `json:"reviewer,omitempty\"` // The user that's responsible for this response.

}

// The metadata for a revision to a file. Some resource methods (such as `revisions.update`) require a `revisionId`. Use the `revisions.list` method to retrieve the ID for a revision.
type Revision struct {
	ExportLinks map[string]interface{} `json:"exportLinks,omitempty\"` // Output only. Links for exporting Docs Editors files to specific formats.

	Id string `json:"id,omitempty\"` // Output only. The ID of the revision.

	KeepForever bool `json:"keepForever,omitempty\"` // Whether to keep this revision forever, even if it is no longer the head revision. If not set, the revision will be automatically purged 30 days after newer content is uploaded. This can be set on a maximum of 200 revisions for a file. This field is only applicable to files with binary content in Drive.

	Kind string `json:"kind,omitempty\"` // Output only. Identifies what kind of resource this is. Value: the fixed string `"drive#revision"`.

	LastModifyingUser User `json:"lastModifyingUser,omitempty\"` // Output only. The last user to modify this revision. This field is only populated when the last modification was performed by a signed-in user.

	Md5Checksum string `json:"md5Checksum,omitempty\"` // Output only. The MD5 checksum of the revision's content. This is only applicable to files with binary content in Drive.

	MimeType string `json:"mimeType,omitempty\"` // Output only. The MIME type of the revision.

	ModifiedTime time.Time `json:"modifiedTime,omitempty\"` // The last time the revision was modified (RFC 3339 date-time).

	OriginalFilename string `json:"originalFilename,omitempty\"` // Output only. The original filename used to create this revision. This is only applicable to files with binary content in Drive.

	PublishAuto bool `json:"publishAuto,omitempty\"` // Whether subsequent revisions will be automatically republished. This is only applicable to Docs Editors files.

	Published bool `json:"published,omitempty\"` // Whether this revision is published. This is only applicable to Docs Editors files.

	PublishedLink string `json:"publishedLink,omitempty\"` // Output only. A link to the published revision. This is only populated for Docs Editors files.

	PublishedOutsideDomain bool `json:"publishedOutsideDomain,omitempty\"` // Whether this revision is published outside the domain. This is only applicable to Docs Editors files.

	Size int64 `json:"size,omitempty\"` // Output only. The size of the revision's content in bytes. This is only applicable to files with binary content in Drive.

}

// A list of revisions of a file.
type RevisionList struct {
	Kind string `json:"kind,omitempty\"` // Identifies what kind of resource this is. Value: the fixed string `"drive#revisionList"`.

	NextPageToken string `json:"nextPageToken,omitempty\"` // The page token for the next page of revisions. This will be absent if the end of the revisions list has been reached. If the token is rejected for any reason, it should be discarded, and pagination should be restarted from the first page of results. The page token is typically valid for several hours. However, if new items are added or removed, your expected results might differ.

	Revisions []Revision `json:"revisions,omitempty\"` // The list of revisions. If nextPageToken is populated, then this list may be incomplete and an additional page of results should be fetched.

}

// Allows creating an approval on a file.
type StartApprovalRequest struct {
	DueTime string `json:"dueTime,omitempty\"` // Optional. The time that the approval is due.

	FileContentChangeBehavior string `json:"fileContentChangeBehavior,omitempty\"` // Optional. The behavior of the approval when the file content changes.

	LockFile bool `json:"lockFile,omitempty\"` // Optional. Whether to lock the file when starting the approval.

	Message string `json:"message,omitempty\"` // Optional. A message to send to reviewers when notifying them of the approval request.

	ReviewerEmails []string `json:"reviewerEmails,omitempty\"` // Required. The emails of the users who are set to review the approval.

}

type StartPageToken struct {
	Kind string `json:"kind,omitempty\"` // Identifies what kind of resource this is. Value: the fixed string `"drive#startPageToken"`.

	StartPageToken string `json:"startPageToken,omitempty\"` // The starting page token for listing future changes. The page token doesn't expire.

}

// The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
type Status struct {
	Code int `json:"code,omitempty\"` // The status code, which should be an enum value of google.rpc.Code.

	Details []map[string]interface{} `json:"details,omitempty\"` // A list of messages that carry the error details. There is a common set of message types for APIs to use.

	Message string `json:"message,omitempty\"` // A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.

}

// Deprecated: use the drive collection instead. Next ID: 33
type TeamDrive struct {
	BackgroundImageFile map[string]interface{} `json:"backgroundImageFile,omitempty\"` // The background image file for a Team Drive.

	BackgroundImageLink string `json:"backgroundImageLink,omitempty\"` // A short-lived link to this Team Drive's background image.

	Capabilities map[string]interface{} `json:"capabilities,omitempty\"` // Capabilities the current user has on this Team Drive.

	ColorRgb string `json:"colorRgb,omitempty\"` // The color of this Team Drive as an RGB hex string. It can only be set on a `drive.teamdrives.update` request that does not set `themeId`.

	CreatedTime time.Time `json:"createdTime,omitempty\"` // The time at which the Team Drive was created (RFC 3339 date-time).

	Id string `json:"id,omitempty\"` // The ID of this Team Drive which is also the ID of the top level folder of this Team Drive.

	Kind string `json:"kind,omitempty\"` // Identifies what kind of resource this is. Value: the fixed string `"drive#teamDrive"`.

	Name string `json:"name,omitempty\"` // The name of this Team Drive.

	OrgUnitId string `json:"orgUnitId,omitempty\"` // The organizational unit of this shared drive. This field is only populated on `drives.list` responses when the `useDomainAdminAccess` parameter is set to `true`.

	Restrictions map[string]interface{} `json:"restrictions,omitempty\"` // A set of restrictions that apply to this Team Drive or items inside this Team Drive.

	ThemeId string `json:"themeId,omitempty\"` // The ID of the theme from which the background image and color will be set. The set of possible `teamDriveThemes` can be retrieved from a `drive.about.get` response. When not specified on a `drive.teamdrives.create` request, a random theme is chosen from which the background image and color are set. This is a write-only field; it can only be set on requests that don't set `colorRgb` or `backgroundImageFile`.

}

// A list of Team Drives.
type TeamDriveList struct {
	Kind string `json:"kind,omitempty\"` // Identifies what kind of resource this is. Value: the fixed string `"drive#teamDriveList"`.

	NextPageToken string `json:"nextPageToken,omitempty\"` // The page token for the next page of Team Drives. This will be absent if the end of the Team Drives list has been reached. If the token is rejected for any reason, it should be discarded, and pagination should be restarted from the first page of results. The page token is typically valid for several hours. However, if new items are added or removed, your expected results might differ.

	TeamDrives []TeamDrive `json:"teamDrives,omitempty\"` // The list of Team Drives. If nextPageToken is populated, then this list may be incomplete and an additional page of results should be fetched.

}

// Information about a Drive user.
type User struct {
	DisplayName string `json:"displayName,omitempty\"` // Output only. A plain text displayable name for this user.

	EmailAddress string `json:"emailAddress,omitempty\"` // Output only. The email address of the user. This may not be present in certain contexts if the user has not made their email address visible to the requester.

	Kind string `json:"kind,omitempty\"` // Output only. Identifies what kind of resource this is. Value: the fixed string `drive#user`.

	Me bool `json:"me,omitempty\"` // Output only. Whether this user is the requesting user.

	PermissionId string `json:"permissionId,omitempty\"` // Output only. The user's ID as visible in Permission resources.

	PhotoLink string `json:"photoLink,omitempty\"` // Output only. A link to the user's profile photo, if available.

}
