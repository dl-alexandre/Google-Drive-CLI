// Google Forms API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package forms

// The submitted answer for a question.
type Answer struct {
	FileUploadAnswers FileUploadAnswers `json:"fileUploadAnswers,omitempty\"` // Output only. The answers to a file upload question.

	Grade Grade `json:"grade,omitempty\"` // Output only. The grade for the answer if the form was a quiz.

	QuestionId string `json:"questionId,omitempty\"` // Output only. The question's ID. See also Question.question_id.

	TextAnswers TextAnswers `json:"textAnswers,omitempty\"` // Output only. The specific answers as text.

}

// A batch of updates to perform on a form. All the specified updates are made or none of them are.
type BatchUpdateFormRequest struct {
	IncludeFormInResponse bool `json:"includeFormInResponse,omitempty\"` // Whether to return an updated version of the model in the response.

	Requests []Request `json:"requests,omitempty\"` // Required. The update requests of this batch.

	WriteControl WriteControl `json:"writeControl,omitempty\"` // Provides control over how write requests are executed.

}

// Response to a BatchUpdateFormRequest.
type BatchUpdateFormResponse struct {
	Form Form `json:"form,omitempty\"` // Based on the bool request field `include_form_in_response`, a form with all applied mutations/updates is returned or not. This may be later than the revision ID created by these changes.

	Replies []Response `json:"replies,omitempty\"` // The reply of the updates. This maps 1:1 with the update requests, although replies to some requests may be empty.

	WriteControl WriteControl `json:"writeControl,omitempty\"` // The updated write control after applying the request.

}

// A radio/checkbox/dropdown question.
type ChoiceQuestion struct {
	Options []Option `json:"options,omitempty\"` // Required. List of options that a respondent must choose from.

	Shuffle bool `json:"shuffle,omitempty\"` // Whether the options should be displayed in random order for different instances of the quiz. This is often used to prevent cheating by respondents who might be looking at another respondent's screen, or to address bias in a survey that might be introduced by always putting the same options first or last.

	TypeValue string `json:"type,omitempty\"` // Required. The type of choice question.

}

// A Pub/Sub topic.
type CloudPubsubTopic struct {
	TopicName string `json:"topicName,omitempty\"` // Required. A fully qualified Pub/Sub topic name to publish the events to. This topic must be owned by the calling project and already exist in Pub/Sub.

}

// A single correct answer for a question. For multiple-valued (`CHECKBOX`) questions, several `CorrectAnswer`s may be needed to represent a single correct response option.
type CorrectAnswer struct {
	Value string `json:"value,omitempty\"` // Required. The correct answer value. See the documentation for TextAnswer.value for details on how various value types are formatted.

}

// The answer key for a question.
type CorrectAnswers struct {
	Answers []CorrectAnswer `json:"answers,omitempty\"` // A list of correct answers. A quiz response can be automatically graded based on these answers. For single-valued questions, a response is marked correct if it matches any value in this list (in other words, multiple correct answers are possible). For multiple-valued (`CHECKBOX`) questions, a response is marked correct if it contains exactly the values in this list.

}

// Create an item in a form.
type CreateItemRequest struct {
	Item Item `json:"item,omitempty\"` // Required. The item to create.

	Location Location `json:"location,omitempty\"` // Required. Where to place the new item.

}

// The result of creating an item.
type CreateItemResponse struct {
	ItemId string `json:"itemId,omitempty\"` // The ID of the created item.

	QuestionId []string `json:"questionId,omitempty\"` // The ID of the question created as part of this item, for a question group it lists IDs of all the questions created for this item.

}

// Create a new watch.
type CreateWatchRequest struct {
	Watch Watch `json:"watch,omitempty\"` // Required. The watch object. No ID should be set on this object; use `watch_id` instead.

	WatchId string `json:"watchId,omitempty\"` // The ID to use for the watch. If specified, the ID must not already be in use. If not specified, an ID is generated. This value should be 4-63 characters, and valid characters are /a-z-/.

}

// A date question. Date questions default to just month + day.
type DateQuestion struct {
	IncludeTime bool `json:"includeTime,omitempty\"` // Whether to include the time as part of the question.

	IncludeYear bool `json:"includeYear,omitempty\"` // Whether to include the year as part of the question.

}

// Delete an item in a form.
type DeleteItemRequest struct {
	Location Location `json:"location,omitempty\"` // Required. The location of the item to delete.

}

// A generic empty message that you can re-use to avoid defining duplicated empty messages in your APIs. A typical example is to use it as the request or the response type of an API method. For instance: service Foo { rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty); }
type Empty struct {
}

// Supplementary material to the feedback.
type ExtraMaterial struct {
	Link TextLink `json:"link,omitempty\"` // Text feedback.

	Video VideoLink `json:"video,omitempty\"` // Video feedback.

}

// Feedback for a respondent about their response to a question.
type Feedback struct {
	Material []ExtraMaterial `json:"material,omitempty\"` // Additional information provided as part of the feedback, often used to point the respondent to more reading and resources.

	Text string `json:"text,omitempty\"` // Required. The main text of the feedback.

}

// Info for a single file submitted to a file upload question.
type FileUploadAnswer struct {
	FileId string `json:"fileId,omitempty\"` // Output only. The ID of the Google Drive file.

	FileName string `json:"fileName,omitempty\"` // Output only. The file name, as stored in Google Drive on upload.

	MimeType string `json:"mimeType,omitempty\"` // Output only. The MIME type of the file, as stored in Google Drive on upload.

}

// All submitted files for a FileUpload question.
type FileUploadAnswers struct {
	Answers []FileUploadAnswer `json:"answers,omitempty\"` // Output only. All submitted files for a FileUpload question.

}

// A file upload question. The API currently does not support creating file upload questions.
type FileUploadQuestion struct {
	FolderId string `json:"folderId,omitempty\"` // Required. The ID of the Drive folder where uploaded files are stored.

	MaxFileSize int64 `json:"maxFileSize,omitempty\"` // Maximum number of bytes allowed for any single file uploaded to this question.

	MaxFiles int `json:"maxFiles,omitempty\"` // Maximum number of files that can be uploaded for this question in a single response.

	Types []string `json:"types,omitempty\"` // File types accepted by this question.

}

// A Google Forms document. A form is created in Drive, and deleting a form or changing its access protections is done via the [Drive API](https://developers.google.com/drive/api/v3/about-sdk).
type Form struct {
	FormId string `json:"formId,omitempty\"` // Output only. The form ID.

	Info Info `json:"info,omitempty\"` // Required. The title and description of the form.

	Items []Item `json:"items,omitempty\"` // Required. A list of the form's items, which can include section headers, questions, embedded media, etc.

	LinkedSheetId string `json:"linkedSheetId,omitempty\"` // Output only. The ID of the linked Google Sheet which is accumulating responses from this Form (if such a Sheet exists).

	PublishSettings PublishSettings `json:"publishSettings,omitempty\"` // Output only. The publishing settings for a form. This field isn't set for legacy forms because they don't have the publish_settings field. All newly created forms support publish settings. Forms with publish_settings value set can call SetPublishSettings API to publish or unpublish the form.

	ResponderUri string `json:"responderUri,omitempty\"` // Output only. The form URI to share with responders. This opens a page that allows the user to submit responses but not edit the questions. For forms that have publish_settings value set, this is the published form URI.

	RevisionId string `json:"revisionId,omitempty\"` // Output only. The revision ID of the form. Used in the WriteControl in update requests to identify the revision on which the changes are based. The format of the revision ID may change over time, so it should be treated opaquely. A returned revision ID is only guaranteed to be valid for 24 hours after it has been returned and cannot be shared across users. If the revision ID is unchanged between calls, then the form *content* has not changed. Conversely, a changed ID (for the same form and user) usually means the form *content* has been updated; however, a changed ID can also be due to internal factors such as ID format changes. Form content excludes form metadata, including: * sharing settings (who has access to the form) * publish_settings (if the form supports publishing and if it is published)

	Settings FormSettings `json:"settings,omitempty\"` // The form's settings. This must be updated with UpdateSettingsRequest; it is ignored during CreateForm and UpdateFormInfoRequest.

}

// A form response.
type FormResponse struct {
	Answers map[string]interface{} `json:"answers,omitempty\"` // Output only. The actual answers to the questions, keyed by question_id.

	CreateTime string `json:"createTime,omitempty\"` // Output only. Timestamp for the first time the response was submitted.

	FormId string `json:"formId,omitempty\"` // Output only. The form ID.

	LastSubmittedTime string `json:"lastSubmittedTime,omitempty\"` // Output only. Timestamp for the most recent time the response was submitted. Does not track changes to grades.

	RespondentEmail string `json:"respondentEmail,omitempty\"` // Output only. The email address (if collected) for the respondent.

	ResponseId string `json:"responseId,omitempty\"` // Output only. The response ID.

	TotalScore float64 `json:"totalScore,omitempty\"` // Output only. The total number of points the respondent received for their submission Only set if the form was a quiz and the response was graded. This includes points automatically awarded via autograding adjusted by any manual corrections entered by the form owner.

}

// A form's settings.
type FormSettings struct {
	EmailCollectionType string `json:"emailCollectionType,omitempty\"` // Optional. The setting that determines whether the form collects email addresses from respondents.

	QuizSettings QuizSettings `json:"quizSettings,omitempty\"` // Settings related to quiz forms and grading.

}

// Grade information associated with a respondent's answer to a question.
type Grade struct {
	Correct bool `json:"correct,omitempty\"` // Output only. Whether the question was answered correctly or not. A zero-point score is not enough to infer incorrectness, since a correctly answered question could be worth zero points.

	Feedback Feedback `json:"feedback,omitempty\"` // Output only. Additional feedback given for an answer.

	Score float64 `json:"score,omitempty\"` // Output only. The numeric score awarded for the answer.

}

// Grading for a single question
type Grading struct {
	CorrectAnswers CorrectAnswers `json:"correctAnswers,omitempty\"` // Required. The answer key for the question. Responses are automatically graded based on this field.

	GeneralFeedback Feedback `json:"generalFeedback,omitempty\"` // The feedback displayed for all answers. This is commonly used for short answer questions when a quiz owner wants to quickly give respondents some sense of whether they answered the question correctly before they've had a chance to officially grade the response. General feedback cannot be set for automatically graded multiple choice questions.

	PointValue int `json:"pointValue,omitempty\"` // Required. The maximum number of points a respondent can automatically get for a correct answer. This must not be negative.

	WhenRight Feedback `json:"whenRight,omitempty\"` // The feedback displayed for correct responses. This feedback can only be set for multiple choice questions that have correct answers provided.

	WhenWrong Feedback `json:"whenWrong,omitempty\"` // The feedback displayed for incorrect responses. This feedback can only be set for multiple choice questions that have correct answers provided.

}

// A grid of choices (radio or check boxes) with each row constituting a separate question. Each row has the same choices, which are shown as the columns.
type Grid struct {
	Columns ChoiceQuestion `json:"columns,omitempty\"` // Required. The choices shared by each question in the grid. In other words, the values of the columns. Only `CHECK_BOX` and `RADIO` choices are allowed.

	ShuffleQuestions bool `json:"shuffleQuestions,omitempty\"` // If `true`, the questions are randomly ordered. In other words, the rows appear in a different order for every respondent.

}

// Data representing an image.
type Image struct {
	AltText string `json:"altText,omitempty\"` // A description of the image that is shown on hover and read by screenreaders.

	ContentUri string `json:"contentUri,omitempty\"` // Output only. A URI from which you can download the image; this is valid only for a limited time.

	Properties MediaProperties `json:"properties,omitempty\"` // Properties of an image.

	SourceUri string `json:"sourceUri,omitempty\"` // Input only. The source URI is the URI used to insert the image. The source URI can be empty when fetched.

}

// An item containing an image.
type ImageItem struct {
	Image Image `json:"image,omitempty\"` // Required. The image displayed in the item.

}

// The general information for a form.
type Info struct {
	Description string `json:"description,omitempty\"` // The description of the form.

	DocumentTitle string `json:"documentTitle,omitempty\"` // Output only. The title of the document which is visible in Drive. If Info.title is empty, `document_title` may appear in its place in the Google Forms UI and be visible to responders. `document_title` can be set on create, but cannot be modified by a batchUpdate request. Please use the [Google Drive API](https://developers.google.com/drive/api/v3/reference/files/update) if you need to programmatically update `document_title`.

	Title string `json:"title,omitempty\"` // Required. The title of the form which is visible to responders.

}

// A single item of the form. `kind` defines which kind of item it is.
type Item struct {
	Description string `json:"description,omitempty\"` // The description of the item.

	ImageItem ImageItem `json:"imageItem,omitempty\"` // Displays an image on the page.

	ItemId string `json:"itemId,omitempty\"` // The item ID. On creation, it can be provided but the ID must not be already used in the form. If not provided, a new ID is assigned.

	PageBreakItem PageBreakItem `json:"pageBreakItem,omitempty\"` // Starts a new page with a title.

	QuestionGroupItem QuestionGroupItem `json:"questionGroupItem,omitempty\"` // Poses one or more questions to the user with a single major prompt.

	QuestionItem QuestionItem `json:"questionItem,omitempty\"` // Poses a question to the user.

	TextItem TextItem `json:"textItem,omitempty\"` // Displays a title and description on the page.

	Title string `json:"title,omitempty\"` // The title of the item.

	VideoItem VideoItem `json:"videoItem,omitempty\"` // Displays a video on the page.

}

// Response to a ListFormResponsesRequest.
type ListFormResponsesResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty\"` // If set, there are more responses. To get the next page of responses, provide this as `page_token` in a future request.

	Responses []FormResponse `json:"responses,omitempty\"` // The returned form responses. Note: The `formId` field is not returned in the `FormResponse` object for list requests.

}

// The response of a ListWatchesRequest.
type ListWatchesResponse struct {
	Watches []Watch `json:"watches,omitempty\"` // The returned watches.

}

// A specific location in a form.
type Location struct {
	Index int `json:"index,omitempty\"` // The index of an item in the form. This must be in the range [0..*N*), where *N* is the number of items in the form.

}

// Properties of the media.
type MediaProperties struct {
	Alignment string `json:"alignment,omitempty\"` // Position of the media.

	Width int `json:"width,omitempty\"` // The width of the media in pixels. When the media is displayed, it is scaled to the smaller of this value or the width of the displayed form. The original aspect ratio of the media is preserved. If a width is not specified when the media is added to the form, it is set to the width of the media source. Width must be between 0 and 740, inclusive. Setting width to 0 or unspecified is only permitted when updating the media source.

}

// Move an item in a form.
type MoveItemRequest struct {
	NewLocation Location `json:"newLocation,omitempty\"` // Required. The new location for the item.

	OriginalLocation Location `json:"originalLocation,omitempty\"` // Required. The location of the item to move.

}

// An option for a Choice question.
type Option struct {
	GoToAction string `json:"goToAction,omitempty\"` // Section navigation type.

	GoToSectionId string `json:"goToSectionId,omitempty\"` // Item ID of section header to go to.

	Image Image `json:"image,omitempty\"` // Display image as an option.

	IsOther bool `json:"isOther,omitempty\"` // Whether the option is "other". Currently only applies to `RADIO` and `CHECKBOX` choice types, but is not allowed in a QuestionGroupItem.

	Value string `json:"value,omitempty\"` // Required. The choice as presented to the user.

}

// A page break. The title and description of this item are shown at the top of the new page.
type PageBreakItem struct {
}

// The publishing settings of a form.
type PublishSettings struct {
	PublishState PublishState `json:"publishState,omitempty\"` // Optional. The publishing state of a form. When updating `publish_state`, both `is_published` and `is_accepting_responses` must be set. However, setting `is_accepting_responses` to `true` and `is_published` to `false` isn't supported and returns an error.

}

// The publishing state of a form.
type PublishState struct {
	IsAcceptingResponses bool `json:"isAcceptingResponses,omitempty\"` // Required. Whether the form accepts responses. If `is_published` is set to `false`, this field is forced to `false`.

	IsPublished bool `json:"isPublished,omitempty\"` // Required. Whether the form is published and visible to others.

}

// Any question. The specific type of question is known by its `kind`.
type Question struct {
	ChoiceQuestion ChoiceQuestion `json:"choiceQuestion,omitempty\"` // A respondent can choose from a pre-defined set of options.

	DateQuestion DateQuestion `json:"dateQuestion,omitempty\"` // A respondent can enter a date.

	FileUploadQuestion FileUploadQuestion `json:"fileUploadQuestion,omitempty\"` // A respondent can upload one or more files.

	Grading Grading `json:"grading,omitempty\"` // Grading setup for the question.

	QuestionId string `json:"questionId,omitempty\"` // Read only. The question ID. On creation, it can be provided but the ID must not be already used in the form. If not provided, a new ID is assigned.

	RatingQuestion RatingQuestion `json:"ratingQuestion,omitempty\"` // A respondent can choose a rating from a pre-defined set of icons.

	Required bool `json:"required,omitempty\"` // Whether the question must be answered in order for a respondent to submit their response.

	RowQuestion RowQuestion `json:"rowQuestion,omitempty\"` // A row of a QuestionGroupItem.

	ScaleQuestion ScaleQuestion `json:"scaleQuestion,omitempty\"` // A respondent can choose a number from a range.

	TextQuestion TextQuestion `json:"textQuestion,omitempty\"` // A respondent can enter a free text response.

	TimeQuestion TimeQuestion `json:"timeQuestion,omitempty\"` // A respondent can enter a time.

}

// Defines a question that comprises multiple questions grouped together.
type QuestionGroupItem struct {
	Grid Grid `json:"grid,omitempty\"` // The question group is a grid with rows of multiple choice questions that share the same options. When `grid` is set, all questions in the group must be of kind `row`.

	Image Image `json:"image,omitempty\"` // The image displayed within the question group above the specific questions.

	Questions []Question `json:"questions,omitempty\"` // Required. A list of questions that belong in this question group. A question must only belong to one group. The `kind` of the group may affect what types of questions are allowed.

}

// A form item containing a single question.
type QuestionItem struct {
	Image Image `json:"image,omitempty\"` // The image displayed within the question.

	Question Question `json:"question,omitempty\"` // Required. The displayed question.

}

// Settings related to quiz forms and grading. These must be updated with the UpdateSettingsRequest.
type QuizSettings struct {
	IsQuiz bool `json:"isQuiz,omitempty\"` // Whether this form is a quiz or not. When true, responses are graded based on question Grading. Upon setting to false, all question Grading is deleted.

}

// A rating question. The user has a range of icons to choose from.
type RatingQuestion struct {
	IconType string `json:"iconType,omitempty\"` // Required. The icon type to use for the rating.

	RatingScaleLevel int `json:"ratingScaleLevel,omitempty\"` // Required. The rating scale level of the rating question.

}

// Renew an existing Watch for seven days.
type RenewWatchRequest struct {
}

// The kinds of update requests that can be made.
type Request struct {
	CreateItem CreateItemRequest `json:"createItem,omitempty\"` // Create a new item.

	DeleteItem DeleteItemRequest `json:"deleteItem,omitempty\"` // Delete an item.

	MoveItem MoveItemRequest `json:"moveItem,omitempty\"` // Move an item to a specified location.

	UpdateFormInfo UpdateFormInfoRequest `json:"updateFormInfo,omitempty\"` // Update Form's Info.

	UpdateItem UpdateItemRequest `json:"updateItem,omitempty\"` // Update an item.

	UpdateSettings UpdateSettingsRequest `json:"updateSettings,omitempty\"` // Updates the Form's settings.

}

// A single response from an update.
type Response struct {
	CreateItem CreateItemResponse `json:"createItem,omitempty\"` // The result of creating an item.

}

// Configuration for a question that is part of a question group.
type RowQuestion struct {
	Title string `json:"title,omitempty\"` // Required. The title for the single row in the QuestionGroupItem.

}

// A scale question. The user has a range of numeric values to choose from.
type ScaleQuestion struct {
	High int `json:"high,omitempty\"` // Required. The highest possible value for the scale.

	HighLabel string `json:"highLabel,omitempty\"` // The label to display describing the highest point on the scale.

	Low int `json:"low,omitempty\"` // Required. The lowest possible value for the scale.

	LowLabel string `json:"lowLabel,omitempty\"` // The label to display describing the lowest point on the scale.

}

// Updates the publish settings of a Form.
type SetPublishSettingsRequest struct {
	PublishSettings PublishSettings `json:"publishSettings,omitempty\"` // Required. The desired publish settings to apply to the form.

	UpdateMask string `json:"updateMask,omitempty\"` // Optional. The `publish_settings` fields to update. This field mask accepts the following values: * `publish_state`: Updates or replaces all `publish_state` settings. * `"*"`: Updates or replaces all `publish_settings` fields.

}

// The response of a SetPublishSettings request.
type SetPublishSettingsResponse struct {
	FormId string `json:"formId,omitempty\"` // Required. The ID of the Form. This is same as the Form.form_id field.

	PublishSettings PublishSettings `json:"publishSettings,omitempty\"` // The publish settings of the form.

}

// An answer to a question represented as text.
type TextAnswer struct {
	Value string `json:"value,omitempty\"` // Output only. The answer value. Formatting used for different kinds of question: * ChoiceQuestion * `RADIO` or `DROP_DOWN`: A single string corresponding to the option that was selected. * `CHECKBOX`: Multiple strings corresponding to each option that was selected. * TextQuestion: The text that the user entered. * ScaleQuestion: A string containing the number that was selected. * DateQuestion * Without time or year: MM-DD e.g. "05-19" * With year: YYYY-MM-DD e.g. "1986-05-19" * With time: MM-DD HH:MM e.g. "05-19 14:51" * With year and time: YYYY-MM-DD HH:MM e.g. "1986-05-19 14:51" * TimeQuestion: String with time or duration in HH:MM format e.g. "14:51" * RowQuestion within QuestionGroupItem: The answer for each row of a QuestionGroupItem is represented as a separate Answer. Each will contain one string for `RADIO`-type choices or multiple strings for `CHECKBOX` choices.

}

// A question's answers as text.
type TextAnswers struct {
	Answers []TextAnswer `json:"answers,omitempty\"` // Output only. Answers to a question. For multiple-value ChoiceQuestions, each answer is a separate value.

}

// A text item.
type TextItem struct {
}

// Link for text.
type TextLink struct {
	DisplayText string `json:"displayText,omitempty\"` // Required. Display text for the URI.

	Uri string `json:"uri,omitempty\"` // Required. The URI.

}

// A text-based question.
type TextQuestion struct {
	Paragraph bool `json:"paragraph,omitempty\"` // Whether the question is a paragraph question or not. If not, the question is a short text question.

}

// A time question.
type TimeQuestion struct {
	Duration bool `json:"duration,omitempty\"` // `true` if the question is about an elapsed time. Otherwise it is about a time of day.

}

// Update Form's Info.
type UpdateFormInfoRequest struct {
	Info Info `json:"info,omitempty\"` // The info to update.

	UpdateMask string `json:"updateMask,omitempty\"` // Required. Only values named in this mask are changed. At least one field must be specified. The root `info` is implied and should not be specified. A single `"*"` can be used as short-hand for updating every field.

}

// Update an item in a form.
type UpdateItemRequest struct {
	Item Item `json:"item,omitempty\"` // Required. New values for the item. Note that item and question IDs are used if they are provided (and are in the field mask). If an ID is blank (and in the field mask) a new ID is generated. This means you can modify an item by getting the form via forms.get, modifying your local copy of that item to be how you want it, and using UpdateItemRequest to write it back, with the IDs being the same (or not in the field mask).

	Location Location `json:"location,omitempty\"` // Required. The location identifying the item to update.

	UpdateMask string `json:"updateMask,omitempty\"` // Required. Only values named in this mask are changed.

}

// Update Form's FormSettings.
type UpdateSettingsRequest struct {
	Settings FormSettings `json:"settings,omitempty\"` // Required. The settings to update with.

	UpdateMask string `json:"updateMask,omitempty\"` // Required. Only values named in this mask are changed. At least one field must be specified. The root `settings` is implied and should not be specified. A single `"*"` can be used as short-hand for updating every field.

}

// Data representing a video.
type Video struct {
	Properties MediaProperties `json:"properties,omitempty\"` // Properties of a video.

	YoutubeUri string `json:"youtubeUri,omitempty\"` // Required. A YouTube URI.

}

// An item containing a video.
type VideoItem struct {
	Caption string `json:"caption,omitempty\"` // The text displayed below the video.

	Video Video `json:"video,omitempty\"` // Required. The video displayed in the item.

}

// Link to a video.
type VideoLink struct {
	DisplayText string `json:"displayText,omitempty\"` // Required. The display text for the link.

	YoutubeUri string `json:"youtubeUri,omitempty\"` // The URI of a YouTube video.

}

// A watch for events for a form. When the designated event happens, a notification will be published to the specified target. The notification's attributes will include a `formId` key that has the ID of the watched form and an `eventType` key that has the string of the type. Messages are sent with at-least-once delivery and are only dropped in extraordinary circumstances. Typically all notifications should be reliably delivered within a few seconds; however, in some situations notifications may be delayed. A watch expires seven days after it is created unless it is renewed with watches.renew
type Watch struct {
	CreateTime string `json:"createTime,omitempty\"` // Output only. Timestamp of when this was created.

	ErrorType string `json:"errorType,omitempty\"` // Output only. The most recent error type for an attempted delivery. To begin watching the form again a call can be made to watches.renew which also clears this error information.

	EventType string `json:"eventType,omitempty\"` // Required. Which event type to watch for.

	ExpireTime string `json:"expireTime,omitempty\"` // Output only. Timestamp for when this will expire. Each watches.renew call resets this to seven days in the future.

	Id string `json:"id,omitempty\"` // Output only. The ID of this watch. See notes on CreateWatchRequest.watch_id.

	State string `json:"state,omitempty\"` // Output only. The current state of the watch. Additional details about suspended watches can be found by checking the `error_type`.

	Target WatchTarget `json:"target,omitempty\"` // Required. Where to send the notification.

}

// The target for notification delivery.
type WatchTarget struct {
	Topic CloudPubsubTopic `json:"topic,omitempty\"` // A Pub/Sub topic. To receive notifications, the topic must grant publish privileges to the Forms service account `serviceAccount:forms-notifications@system.gserviceaccount.com`. Only the project that owns a topic may create a watch with it. Pub/Sub delivery guarantees should be considered.

}

// Provides control over how write requests are executed.
type WriteControl struct {
	RequiredRevisionId string `json:"requiredRevisionId,omitempty\"` // The revision ID of the form that the write request is applied to. If this is not the latest revision of the form, the request is not processed and returns a 400 bad request error.

	TargetRevisionId string `json:"targetRevisionId,omitempty\"` // The target revision ID of the form that the write request is applied to. If changes have occurred after this revision, the changes in this update request are transformed against those changes. This results in a new revision of the form that incorporates both the changes in the request and the intervening changes, with the server resolving conflicting changes. The target revision ID may only be used to write to recent versions of a form. If the target revision is too far behind the latest revision, the request is not processed and returns a 400 (Bad Request Error). The request may be retried after reading the latest version of the form. In most cases a target revision ID remains valid for several minutes after it is read, but for frequently-edited forms this window may be shorter.

}
