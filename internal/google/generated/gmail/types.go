// Gmail API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package gmail

// Auto-forwarding settings for an account.
type AutoForwarding struct {
	Disposition string `json:"disposition,omitempty\"` // The state that a message should be left in after it has been forwarded.

	EmailAddress string `json:"emailAddress,omitempty\"` // Email address to which all incoming messages are forwarded. This email address must be a verified member of the forwarding addresses.

	Enabled bool `json:"enabled,omitempty\"` // Whether all incoming mail is automatically forwarded to another address.

}

type BatchDeleteMessagesRequest struct {
	Ids []string `json:"ids,omitempty\"` // The IDs of the messages to delete.

}

type BatchModifyMessagesRequest struct {
	AddClassificationLabels []ClassificationLabelValue `json:"addClassificationLabels,omitempty\"` // A list of Classification Label values to add. If a Classification Label with the same label ID is already applied to the message, fields with existing field IDs will be updated and fields with new field IDs will be added. There's a limit of 20 Classification Label values per request. If the message is already classified and the final total number of Classification Label values exceeds the maximum allowed number of Classification Label values per message, the modification fails.

	AddLabelIds []string `json:"addLabelIds,omitempty\"` // A list of label IDs to add to messages.

	Ids []string `json:"ids,omitempty\"` // The IDs of the messages to modify. There is a limit of 1000 ids per request.

	RemoveClassificationLabelIds []string `json:"removeClassificationLabelIds,omitempty\"` // A list of Classification Label values to remove from messages.

	RemoveLabelIds []string `json:"removeLabelIds,omitempty\"` // A list of label IDs to remove from messages.

}

// Field values for a classification label.
type ClassificationLabelFieldValue struct {
	FieldId string `json:"fieldId,omitempty\"` // Required. The field ID for the Classification Label Value. Maps to the ID field of the Google Drive `Label.Field` object.

	Selection string `json:"selection,omitempty\"` // Selection choice ID for the selection option. Should only be set if the field type is `SELECTION` in the Google Drive `Label.Field` object. Maps to the id field of the Google Drive `Label.Field.SelectionOptions` resource.

}

// Classification Labels applied to the email message. Classification Labels are different from Gmail inbox labels. Only used for Google Workspace accounts. [Learn more about classification labels](https://support.google.com/a/answer/9292382).
type ClassificationLabelValue struct {
	Fields []ClassificationLabelFieldValue `json:"fields,omitempty\"` // Field values for the given classification label ID.

	LabelId string `json:"labelId,omitempty\"` // Required. The canonical or raw alphanumeric classification label ID. Maps to the ID field of the Google Drive Label resource.

}

// The client-side encryption (CSE) configuration for the email address of an authenticated user. Gmail uses CSE configurations to save drafts of client-side encrypted email messages, and to sign and send encrypted email messages. For administrators managing identities and keypairs for users in their organization, requests require authorization with a [service account](https://developers.google.com/identity/protocols/OAuth2ServiceAccount) that has [domain-wide delegation authority](https://developers.google.com/identity/protocols/OAuth2ServiceAccount#delegatingauthority) to impersonate users with the `https://www.googleapis.com/auth/gmail.settings.basic` scope. For users managing their own identities and keypairs, requests require [hardware key encryption](https://support.google.com/a/answer/14153163) turned on and configured.
type CseIdentity struct {
	EmailAddress string `json:"emailAddress,omitempty\"` // The email address for the sending identity. The email address must be the primary email address of the authenticated user.

	PrimaryKeyPairId string `json:"primaryKeyPairId,omitempty\"` // If a key pair is associated, the ID of the key pair, CseKeyPair.

	SignAndEncryptKeyPairs SignAndEncryptKeyPairs `json:"signAndEncryptKeyPairs,omitempty\"` // The configuration of a CSE identity that uses different key pairs for signing and encryption.

}

// A client-side encryption S/MIME key pair, which is comprised of a public key, its certificate chain, and metadata for its paired private key. Gmail uses the key pair to complete the following tasks: - Sign outgoing client-side encrypted messages. - Save and reopen drafts of client-side encrypted messages. - Save and reopen sent messages. - Decrypt incoming or archived S/MIME messages. For administrators managing identities and keypairs for users in their organization, requests require authorization with a [service account](https://developers.google.com/identity/protocols/OAuth2ServiceAccount) that has [domain-wide delegation authority](https://developers.google.com/identity/protocols/OAuth2ServiceAccount#delegatingauthority) to impersonate users with the `https://www.googleapis.com/auth/gmail.settings.basic` scope. For users managing their own identities and keypairs, requests require [hardware key encryption](https://support.google.com/a/answer/14153163) turned on and configured.
type CseKeyPair struct {
	DisableTime string `json:"disableTime,omitempty\"` // Output only. If a key pair is set to `DISABLED`, the time that the key pair's state changed from `ENABLED` to `DISABLED`. This field is present only when the key pair is in state `DISABLED`.

	EnablementState string `json:"enablementState,omitempty\"` // Output only. The current state of the key pair.

	KeyPairId string `json:"keyPairId,omitempty\"` // Output only. The immutable ID for the client-side encryption S/MIME key pair.

	Pem string `json:"pem,omitempty\"` // Output only. The public key and its certificate chain, in [PEM](https://en.wikipedia.org/wiki/Privacy-Enhanced_Mail) format.

	Pkcs7 string `json:"pkcs7,omitempty\"` // Input only. The public key and its certificate chain. The chain must be in [PKCS#7](https://en.wikipedia.org/wiki/PKCS_7) format and use PEM encoding and ASCII armor.

	PrivateKeyMetadata []CsePrivateKeyMetadata `json:"privateKeyMetadata,omitempty\"` // Metadata for instances of this key pair's private key.

	SubjectEmailAddresses []string `json:"subjectEmailAddresses,omitempty\"` // Output only. The email address identities that are specified on the leaf certificate.

}

// Metadata for a private key instance.
type CsePrivateKeyMetadata struct {
	HardwareKeyMetadata HardwareKeyMetadata `json:"hardwareKeyMetadata,omitempty\"` // Metadata for hardware keys.

	KaclsKeyMetadata KaclsKeyMetadata `json:"kaclsKeyMetadata,omitempty\"` // Metadata for a private key instance managed by an external key access control list service.

	PrivateKeyMetadataId string `json:"privateKeyMetadataId,omitempty\"` // Output only. The immutable ID for the private key metadata instance.

}

// Settings for a delegate. Delegates can read, send, and delete messages, as well as view and add contacts, for the delegator's account. See "Set up mail delegation" for more information about delegates.
type Delegate struct {
	DelegateEmail string `json:"delegateEmail,omitempty\"` // The email address of the delegate.

	VerificationStatus string `json:"verificationStatus,omitempty\"` // Indicates whether this address has been verified and can act as a delegate for the account. Read-only.

}

// Requests to turn off a client-side encryption key pair.
type DisableCseKeyPairRequest struct {
}

// A draft email in the user's mailbox.
type Draft struct {
	Id string `json:"id,omitempty\"` // The immutable ID of the draft.

	Message Message `json:"message,omitempty\"` // The message content of the draft.

}

// Requests to turn on a client-side encryption key pair.
type EnableCseKeyPairRequest struct {
}

// Resource definition for Gmail filters. Filters apply to specific messages instead of an entire email thread.
type Filter struct {
	Action FilterAction `json:"action,omitempty\"` // Action that the filter performs.

	Criteria FilterCriteria `json:"criteria,omitempty\"` // Matching criteria for the filter.

	Id string `json:"id,omitempty\"` // The server assigned ID of the filter.

}

// A set of actions to perform on a message.
type FilterAction struct {
	AddLabelIds []string `json:"addLabelIds,omitempty\"` // List of labels to add to the message.

	Forward string `json:"forward,omitempty\"` // Email address that the message should be forwarded to. This effectively redirects the message to the address specified in this field, maintaining the original sender in the "From" field.

	RemoveLabelIds []string `json:"removeLabelIds,omitempty\"` // List of labels to remove from the message.

}

// Message matching criteria.
type FilterCriteria struct {
	ExcludeChats bool `json:"excludeChats,omitempty\"` // Whether the response should exclude chats.

	From string `json:"from,omitempty\"` // The sender's display name or email address.

	HasAttachment bool `json:"hasAttachment,omitempty\"` // Whether the message has any attachment.

	NegatedQuery string `json:"negatedQuery,omitempty\"` // Only return messages not matching the specified query. Supports the same query format as the Gmail search box. For example, `"from:someuser@example.com rfc822msgid: is:unread"`.

	Query string `json:"query,omitempty\"` // Only return messages matching the specified query. Supports the same query format as the Gmail search box. For example, `"from:someuser@example.com rfc822msgid: is:unread"`.

	Size int `json:"size,omitempty\"` // The size of the entire RFC822 message in bytes, including all headers and attachments.

	SizeComparison string `json:"sizeComparison,omitempty\"` // How the message size in bytes should be in relation to the size field.

	Subject string `json:"subject,omitempty\"` // Case-insensitive phrase found in the message's subject. Trailing and leading whitespace are be trimmed and adjacent spaces are collapsed.

	To string `json:"to,omitempty\"` // The recipient's display name or email address. Includes recipients in the "to", "cc", and "bcc" header fields. You can use simply the local part of the email address. For example, "example" and "example@" both match "example@gmail.com". This field is case-insensitive.

}

// Settings for a forwarding address.
type ForwardingAddress struct {
	ForwardingEmail string `json:"forwardingEmail,omitempty\"` // An email address to which messages can be forwarded.

	VerificationStatus string `json:"verificationStatus,omitempty\"` // Indicates whether this address has been verified and is usable for forwarding. Read-only.

}

// Metadata for hardware keys. If [hardware key encryption](https://support.google.com/a/answer/14153163) is set up for the Google Workspace organization, users can optionally store their private key on their smart card and use it to sign and decrypt email messages in Gmail by inserting their smart card into a reader attached to their Windows device.
type HardwareKeyMetadata struct {
	Description string `json:"description,omitempty\"` // Description about the hardware key.

}

// A record of a change to the user's mailbox. Each history change may affect multiple messages in multiple ways.
type History struct {
	Id uint64 `json:"id,omitempty\"` // The mailbox sequence ID.

	LabelsAdded []HistoryLabelAdded `json:"labelsAdded,omitempty\"` // Labels added to messages in this history record.

	LabelsRemoved []HistoryLabelRemoved `json:"labelsRemoved,omitempty\"` // Labels removed from messages in this history record.

	Messages []Message `json:"messages,omitempty\"` // List of messages changed in this history record. The fields for specific change types, such as `messagesAdded` may duplicate messages in this field. We recommend using the specific change-type fields instead of this.

	MessagesAdded []HistoryMessageAdded `json:"messagesAdded,omitempty\"` // Messages added to the mailbox in this history record.

	MessagesDeleted []HistoryMessageDeleted `json:"messagesDeleted,omitempty\"` // Messages deleted (not Trashed) from the mailbox in this history record.

}

type HistoryLabelAdded struct {
	LabelIds []string `json:"labelIds,omitempty\"` // Label IDs added to the message.

	Message Message `json:"message,omitempty\"`
}

type HistoryLabelRemoved struct {
	LabelIds []string `json:"labelIds,omitempty\"` // Label IDs removed from the message.

	Message Message `json:"message,omitempty\"`
}

type HistoryMessageAdded struct {
	Message Message `json:"message,omitempty\"`
}

type HistoryMessageDeleted struct {
	Message Message `json:"message,omitempty\"`
}

// IMAP settings for an account.
type ImapSettings struct {
	AutoExpunge bool `json:"autoExpunge,omitempty\"` // If this value is true, Gmail will immediately expunge a message when it is marked as deleted in IMAP. Otherwise, Gmail will wait for an update from the client before expunging messages marked as deleted.

	Enabled bool `json:"enabled,omitempty\"` // Whether IMAP is enabled for the account.

	ExpungeBehavior string `json:"expungeBehavior,omitempty\"` // The action that will be executed on a message when it is marked as deleted and expunged from the last visible IMAP folder.

	MaxFolderSize int `json:"maxFolderSize,omitempty\"` // An optional limit on the number of messages that an IMAP folder may contain. Legal values are 0, 1000, 2000, 5000 or 10000. A value of zero is interpreted to mean that there is no limit.

}

// Metadata for private keys managed by an external key access control list service. For details about managing key access, see [Google Workspace CSE API Reference](https://developers.google.com/workspace/cse/reference).
type KaclsKeyMetadata struct {
	KaclsData string `json:"kaclsData,omitempty\"` // Opaque data generated and used by the key access control list service. Maximum size: 8 KiB.

	KaclsUri string `json:"kaclsUri,omitempty\"` // The URI of the key access control list service that manages the private key.

}

// Labels are used to categorize messages and threads within the user's mailbox. The maximum number of labels supported for a user's mailbox is 10,000.
type Label struct {
	Color LabelColor `json:"color,omitempty\"` // The color to assign to the label. Color is only available for labels that have their `type` set to `user`.

	Id string `json:"id,omitempty\"` // The immutable ID of the label.

	LabelListVisibility string `json:"labelListVisibility,omitempty\"` // The visibility of the label in the label list in the Gmail web interface.

	MessageListVisibility string `json:"messageListVisibility,omitempty\"` // The visibility of messages with this label in the message list in the Gmail web interface.

	MessagesTotal int `json:"messagesTotal,omitempty\"` // The total number of messages with the label.

	MessagesUnread int `json:"messagesUnread,omitempty\"` // The number of unread messages with the label.

	Name string `json:"name,omitempty\"` // The display name of the label.

	ThreadsTotal int `json:"threadsTotal,omitempty\"` // The total number of threads with the label.

	ThreadsUnread int `json:"threadsUnread,omitempty\"` // The number of unread threads with the label.

	TypeValue string `json:"type,omitempty\"` // The owner type for the label. User labels are created by the user and can be modified and deleted by the user and can be applied to any message or thread. System labels are internally created and cannot be added, modified, or deleted. System labels may be able to be applied to or removed from messages and threads under some circumstances but this is not guaranteed. For example, users can apply and remove the `INBOX` and `UNREAD` labels from messages and threads, but cannot apply or remove the `DRAFTS` or `SENT` labels from messages or threads.

}

type LabelColor struct {
	BackgroundColor string `json:"backgroundColor,omitempty\"` // The background color represented as hex string #RRGGBB (ex #000000). This field is required in order to set the color of a label. Only the following predefined set of color values are allowed: \#000000, #434343, #666666, #999999, #cccccc, #efefef, #f3f3f3, #ffffff, \#fb4c2f, #ffad47, #fad165, #16a766, #43d692, #4a86e8, #a479e2, #f691b3, \#f6c5be, #ffe6c7, #fef1d1, #b9e4d0, #c6f3de, #c9daf8, #e4d7f5, #fcdee8, \#efa093, #ffd6a2, #fce8b3, #89d3b2, #a0eac9, #a4c2f4, #d0bcf1, #fbc8d9, \#e66550, #ffbc6b, #fcda83, #44b984, #68dfa9, #6d9eeb, #b694e8, #f7a7c0, \#cc3a21, #eaa041, #f2c960, #149e60, #3dc789, #3c78d8, #8e63ce, #e07798, \#ac2b16, #cf8933, #d5ae49, #0b804b, #2a9c68, #285bac, #653e9b, #b65775, \#822111, #a46a21, #aa8831, #076239, #1a764d, #1c4587, #41236d, #83334c, \#464646, #e7e7e7, #0d3472, #b6cff5, #0d3b44, #98d7e4, #3d188e, #e3d7ff, \#711a36, #fbd3e0, #8a1c0a, #f2b2a8, #7a2e0b, #ffc8af, #7a4706, #ffdeb5, \#594c05, #fbe983, #684e07, #fdedc1, #0b4f30, #b3efd3, #04502e, #a2dcc1, \#c2c2c2, #4986e7, #2da2bb, #b99aff, #994a64, #f691b2, #ff7537, #ffad46, \#662e37, #ebdbde, #cca6ac, #094228, #42d692, #16a765, #757575, #1e53b8, \#007286, #7858c3, #c2185b, #d93025, #54240e, #633e04, #521d28, #202124, \#083018

	TextColor string `json:"textColor,omitempty\"` // The text color of the label, represented as hex string. This field is required in order to set the color of a label. Only the following predefined set of color values are allowed: \#000000, #434343, #666666, #999999, #cccccc, #efefef, #f3f3f3, #ffffff, \#fb4c2f, #ffad47, #fad165, #16a766, #43d692, #4a86e8, #a479e2, #f691b3, \#f6c5be, #ffe6c7, #fef1d1, #b9e4d0, #c6f3de, #c9daf8, #e4d7f5, #fcdee8, \#efa093, #ffd6a2, #fce8b3, #89d3b2, #a0eac9, #a4c2f4, #d0bcf1, #fbc8d9, \#e66550, #ffbc6b, #fcda83, #44b984, #68dfa9, #6d9eeb, #b694e8, #f7a7c0, \#cc3a21, #eaa041, #f2c960, #149e60, #3dc789, #3c78d8, #8e63ce, #e07798, \#ac2b16, #cf8933, #d5ae49, #0b804b, #2a9c68, #285bac, #653e9b, #b65775, \#822111, #a46a21, #aa8831, #076239, #1a764d, #1c4587, #41236d, #83334c, \#464646, #e7e7e7, #0d3472, #b6cff5, #0d3b44, #98d7e4, #3d188e, #e3d7ff, \#711a36, #fbd3e0, #8a1c0a, #f2b2a8, #7a2e0b, #ffc8af, #7a4706, #ffdeb5, \#594c05, #fbe983, #684e07, #fdedc1, #0b4f30, #b3efd3, #04502e, #a2dcc1, \#c2c2c2, #4986e7, #2da2bb, #b99aff, #994a64, #f691b2, #ff7537, #ffad46, \#662e37, #ebdbde, #cca6ac, #094228, #42d692, #16a765, #757575, #1e53b8, \#007286, #7858c3, #c2185b, #d93025, #54240e, #633e04, #521d28, #202124, \#083018

}

// Language settings for an account. These settings correspond to the "Language settings" feature in the web interface.
type LanguageSettings struct {
	DisplayLanguage string `json:"displayLanguage,omitempty\"` // The language to display Gmail in, formatted as an RFC 3066 Language Tag (for example `en-GB`, `fr` or `ja` for British English, French, or Japanese respectively). The set of languages supported by Gmail evolves over time, so please refer to the "Language" dropdown in the Gmail settings for all available options, as described in the language settings help article. For a table of sample values, see [Manage language settings](https://developers.google.com/workspace/gmail/api/guides/language-settings). Not all Gmail clients can display the same set of languages. In the case that a user's display language is not available for use on a particular client, said client automatically chooses to display in the closest supported variant (or a reasonable default).

}

type ListCseIdentitiesResponse struct {
	CseIdentities []CseIdentity `json:"cseIdentities,omitempty\"` // One page of the list of CSE identities configured for the user.

	NextPageToken string `json:"nextPageToken,omitempty\"` // Pagination token to be passed to a subsequent ListCseIdentities call in order to retrieve the next page of identities. If this value is not returned or is the empty string, then no further pages remain.

}

type ListCseKeyPairsResponse struct {
	CseKeyPairs []CseKeyPair `json:"cseKeyPairs,omitempty\"` // One page of the list of CSE key pairs installed for the user.

	NextPageToken string `json:"nextPageToken,omitempty\"` // Pagination token to be passed to a subsequent ListCseKeyPairs call in order to retrieve the next page of key pairs. If this value is not returned, then no further pages remain.

}

// Response for the ListDelegates method.
type ListDelegatesResponse struct {
	Delegates []Delegate `json:"delegates,omitempty\"` // List of the user's delegates (with any verification status). If an account doesn't have delegates, this field doesn't appear.

}

type ListDraftsResponse struct {
	Drafts []Draft `json:"drafts,omitempty\"` // List of drafts. Note that the `Message` property in each `Draft` resource only contains an `id` and a `threadId`. The [`messages.get`](https://developers.google.com/workspace/gmail/api/v1/reference/users/messages/get) method can fetch additional message details.

	NextPageToken string `json:"nextPageToken,omitempty\"` // Token to retrieve the next page of results in the list.

	ResultSizeEstimate int `json:"resultSizeEstimate,omitempty\"` // Estimated total number of results.

}

// Response for the ListFilters method.
type ListFiltersResponse struct {
	Filter []Filter `json:"filter,omitempty\"` // List of a user's filters.

}

// Response for the ListForwardingAddresses method.
type ListForwardingAddressesResponse struct {
	ForwardingAddresses []ForwardingAddress `json:"forwardingAddresses,omitempty\"` // List of addresses that may be used for forwarding.

}

type ListHistoryResponse struct {
	History []History `json:"history,omitempty\"` // List of history records. Any `messages` contained in the response will typically only have `id` and `threadId` fields populated.

	HistoryId uint64 `json:"historyId,omitempty\"` // The ID of the mailbox's current history record.

	NextPageToken string `json:"nextPageToken,omitempty\"` // Page token to retrieve the next page of results in the list.

}

type ListLabelsResponse struct {
	Labels []Label `json:"labels,omitempty\"` // List of labels. Note that each label resource only contains an `id`, `name`, `messageListVisibility`, `labelListVisibility`, and `type`. The [`labels.get`](https://developers.google.com/workspace/gmail/api/v1/reference/users/labels/get) method can fetch additional label details.

}

type ListMessagesResponse struct {
	Messages []Message `json:"messages,omitempty\"` // List of messages. Note that each message resource contains only an `id` and a `threadId`. Additional message details can be fetched using the messages.get method.

	NextPageToken string `json:"nextPageToken,omitempty\"` // Token to retrieve the next page of results in the list.

	ResultSizeEstimate int `json:"resultSizeEstimate,omitempty\"` // Estimated total number of results.

}

// Response for the ListSendAs method.
type ListSendAsResponse struct {
	SendAs []SendAs `json:"sendAs,omitempty\"` // List of send-as aliases.

}

type ListSmimeInfoResponse struct {
	SmimeInfo []SmimeInfo `json:"smimeInfo,omitempty\"` // List of SmimeInfo.

}

type ListThreadsResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty\"` // Page token to retrieve the next page of results in the list.

	ResultSizeEstimate int `json:"resultSizeEstimate,omitempty\"` // Estimated total number of results.

	Threads []Thread `json:"threads,omitempty\"` // List of threads. Note that each thread resource does not contain a list of `messages`. The list of `messages` for a given thread can be fetched using the [`threads.get`](https://developers.google.com/workspace/gmail/api/v1/reference/users/threads/get) method.

}

// An email message.
type Message struct {
	ClassificationLabelValues []ClassificationLabelValue `json:"classificationLabelValues,omitempty\"` // Classification Label values on the message. Available Classification Label schemas can be queried using the Google Drive Labels API. Each classification label ID must be unique. If duplicate IDs are provided, only one will be retained, and the selection is arbitrary. Only used for Google Workspace accounts. There's a limit of 20 Classification Label values per request. If the Classification Label values exceeds the maximum allowed number, the request fails.

	HistoryId uint64 `json:"historyId,omitempty\"` // The ID of the last history record that modified this message.

	Id string `json:"id,omitempty\"` // The immutable ID of the message.

	InternalDate int64 `json:"internalDate,omitempty\"` // The internal message creation timestamp (epoch ms), which determines ordering in the inbox. For normal SMTP-received email, this represents the time the message was originally accepted by Google, which is more reliable than the `Date` header. However, for API-migrated mail, it can be configured by client to be based on the `Date` header.

	LabelIds []string `json:"labelIds,omitempty\"` // List of IDs of labels applied to this message.

	Payload MessagePart `json:"payload,omitempty\"` // The parsed email structure in the message parts.

	Raw string `json:"raw,omitempty\"` // The entire email message in an RFC 2822 formatted and base64url encoded string. Returned in `messages.get` and `drafts.get` responses when the `format=RAW` parameter is supplied. @required gmail.users.drafts.create gmail.users.drafts.update

	SizeEstimate int `json:"sizeEstimate,omitempty\"` // Estimated size in bytes of the message.

	Snippet string `json:"snippet,omitempty\"` // A short part of the message text.

	ThreadId string `json:"threadId,omitempty\"` // The ID of the thread the message belongs to. To add a message or draft to a thread, the following criteria must be met: 1. The requested `threadId` must be specified on the `Message` or `Draft.Message` you supply with your request. 2. The `References` and `In-Reply-To` headers must be set in compliance with the [RFC 2822](https://tools.ietf.org/html/rfc2822) standard. 3. The `Subject` headers must match.

}

// A single MIME message part.
type MessagePart struct {
	Body MessagePartBody `json:"body,omitempty\"` // The message part body for this part, which may be empty for container MIME message parts.

	Filename string `json:"filename,omitempty\"` // The filename of the attachment. Only present if this message part represents an attachment.

	Headers []MessagePartHeader `json:"headers,omitempty\"` // List of headers on this message part. For the top-level message part, representing the entire message payload, it will contain the standard RFC 2822 email headers such as `To`, `From`, and `Subject`.

	MimeType string `json:"mimeType,omitempty\"` // The MIME type of the message part.

	PartId string `json:"partId,omitempty\"` // The immutable ID of the message part.

	Parts []MessagePart `json:"parts,omitempty\"` // The child MIME message parts of this part. This only applies to container MIME message parts, for example `multipart/*`. For non- container MIME message part types, such as `text/plain`, this field is empty. For more information, see RFC 1521.

}

// The body of a single MIME message part.
type MessagePartBody struct {
	AttachmentId string `json:"attachmentId,omitempty\"` // When present, contains the ID of an external attachment that can be retrieved in a separate `messages.attachments.get` request. When not present, the entire content of the message part body is contained in the data field.

	Data string `json:"data,omitempty\"` // The body data of a MIME message part as a base64url encoded string. May be empty for MIME container types that have no message body or when the body data is sent as a separate attachment. An attachment ID is present if the body data is contained in a separate attachment.

	Size int `json:"size,omitempty\"` // Number of bytes for the message part data (encoding notwithstanding).

}

type MessagePartHeader struct {
	Name string `json:"name,omitempty\"` // The name of the header before the `:` separator. For example, `To`.

	Value string `json:"value,omitempty\"` // The value of the header after the `:` separator. For example, `someuser@example.com`.

}

type ModifyMessageRequest struct {
	AddClassificationLabels []ClassificationLabelValue `json:"addClassificationLabels,omitempty\"` // A list of classification label values to add. If a Classification Label with the same label ID is already applied to the message, fields with existing field IDs will be updated and fields with new field IDs will be added. There's a limit of 20 Classification Label values per request. If the message is already classified and the final total number of Classification Label values exceeds the maximum allowed number of Classification Label values per message, the modification fails.

	AddLabelIds []string `json:"addLabelIds,omitempty\"` // A list of IDs of labels to add to this message. You can add up to 100 labels with each update.

	RemoveClassificationLabelIds []string `json:"removeClassificationLabelIds,omitempty\"` // A list of Classification Label values to remove from this message.

	RemoveLabelIds []string `json:"removeLabelIds,omitempty\"` // A list IDs of labels to remove from this message. You can remove up to 100 labels with each update.

}

type ModifyThreadRequest struct {
	AddLabelIds []string `json:"addLabelIds,omitempty\"` // A list of IDs of labels to add to this thread. You can add up to 100 labels with each update.

	RemoveLabelIds []string `json:"removeLabelIds,omitempty\"` // A list of IDs of labels to remove from this thread. You can remove up to 100 labels with each update.

}

// Request to obliterate a CSE key pair.
type ObliterateCseKeyPairRequest struct {
}

// POP settings for an account.
type PopSettings struct {
	AccessWindow string `json:"accessWindow,omitempty\"` // The range of messages which are accessible via POP.

	Disposition string `json:"disposition,omitempty\"` // The action that will be executed on a message after it has been fetched via POP.

}

// Profile for a Gmail user.
type Profile struct {
	EmailAddress string `json:"emailAddress,omitempty\"` // The user's email address.

	HistoryId uint64 `json:"historyId,omitempty\"` // The ID of the mailbox's current history record.

	MessagesTotal int `json:"messagesTotal,omitempty\"` // The total number of messages in the mailbox.

	ThreadsTotal int `json:"threadsTotal,omitempty\"` // The total number of threads in the mailbox.

}

// Settings associated with a send-as alias, which can be either the primary login address associated with the account or a custom "from" address. Send-as aliases correspond to the "Send Mail As" feature in the web interface. The send-as alias must be a valid email address.
type SendAs struct {
	DisplayName string `json:"displayName,omitempty\"` // A name that appears in the "From:" header for mail sent using this alias. For custom "from" addresses, when this is empty, Gmail will populate the "From:" header with the name that is used for the primary address associated with the account. If the admin has disabled the ability for users to update their name format, requests to update this field for the primary login will silently fail.

	IsDefault bool `json:"isDefault,omitempty\"` // Whether this address is selected as the default "From:" address in situations such as composing a new message or sending a vacation auto-reply. Every Gmail account has exactly one default send-as address, so the only legal value that clients may write to this field is `true`. Changing this from `false` to `true` for an address will result in this field becoming `false` for the other previous default address.

	IsPrimary bool `json:"isPrimary,omitempty\"` // Whether this address is the primary address used to login to the account. Every Gmail account has exactly one primary address, and it cannot be deleted from the collection of send-as aliases. This field is read-only.

	ReplyToAddress string `json:"replyToAddress,omitempty\"` // An optional email address that is included in a "Reply-To:" header for mail sent using this alias. If this is empty, Gmail will not generate a "Reply-To:" header.

	SendAsEmail string `json:"sendAsEmail,omitempty\"` // The email address that appears in the "From:" header for mail sent using this alias. This is read-only for all operations except create.

	Signature string `json:"signature,omitempty\"` // An optional HTML signature that is included in messages composed with this alias in the Gmail web UI. This signature is added to new emails only.

	SmtpMsa SmtpMsa `json:"smtpMsa,omitempty\"` // An optional SMTP service that will be used as an outbound relay for mail sent using this alias. If this is empty, outbound mail will be sent directly from Gmail's servers to the destination SMTP service. This setting only applies to custom "from" aliases.

	TreatAsAlias bool `json:"treatAsAlias,omitempty\"` // Whether Gmail should treat this address as an alias for the user's primary email address. This setting only applies to custom "from" aliases.

	VerificationStatus string `json:"verificationStatus,omitempty\"` // Indicates whether this address has been verified for use as a send-as alias. Read-only. This setting only applies to custom "from" aliases.

}

// The configuration of a CSE identity that uses different key pairs for signing and encryption.
type SignAndEncryptKeyPairs struct {
	EncryptionKeyPairId string `json:"encryptionKeyPairId,omitempty\"` // The ID of the CseKeyPair that encrypts signed outgoing mail.

	SigningKeyPairId string `json:"signingKeyPairId,omitempty\"` // The ID of the CseKeyPair that signs outgoing mail.

}

// An S/MIME email config.
type SmimeInfo struct {
	EncryptedKeyPassword string `json:"encryptedKeyPassword,omitempty\"` // Encrypted key password, when key is encrypted.

	Expiration int64 `json:"expiration,omitempty\"` // When the certificate expires (in milliseconds since epoch).

	Id string `json:"id,omitempty\"` // The immutable ID for the SmimeInfo.

	IsDefault bool `json:"isDefault,omitempty\"` // Whether this SmimeInfo is the default one for this user's send-as address.

	IssuerCn string `json:"issuerCn,omitempty\"` // The S/MIME certificate issuer's common name.

	Pem string `json:"pem,omitempty\"` // PEM formatted X509 concatenated certificate string (standard base64 encoding). Format used for returning key, which includes public key as well as certificate chain (not private key).

	Pkcs12 string `json:"pkcs12,omitempty\"` // PKCS#12 format containing a single private/public key pair and certificate chain. This format is only accepted from client for creating a new SmimeInfo and is never returned, because the private key is not intended to be exported. PKCS#12 may be encrypted, in which case encryptedKeyPassword should be set appropriately.

}

// Configuration for communication with an SMTP service.
type SmtpMsa struct {
	Host string `json:"host,omitempty\"` // The hostname of the SMTP service. Required.

	Password string `json:"password,omitempty\"` // The password that will be used for authentication with the SMTP service. This is a write-only field that can be specified in requests to create or update SendAs settings; it is never populated in responses.

	Port int `json:"port,omitempty\"` // The port of the SMTP service. Required.

	SecurityMode string `json:"securityMode,omitempty\"` // The protocol that will be used to secure communication with the SMTP service. Required.

	Username string `json:"username,omitempty\"` // The username that will be used for authentication with the SMTP service. This is a write-only field that can be specified in requests to create or update SendAs settings; it is never populated in responses.

}

// A collection of messages representing a conversation.
type Thread struct {
	HistoryId uint64 `json:"historyId,omitempty\"` // The ID of the last history record that modified this thread.

	Id string `json:"id,omitempty\"` // The unique ID of the thread.

	Messages []Message `json:"messages,omitempty\"` // The list of messages in the thread.

	Snippet string `json:"snippet,omitempty\"` // A short part of the message text.

}

// Vacation auto-reply settings for an account. These settings correspond to the "Vacation responder" feature in the web interface.
type VacationSettings struct {
	EnableAutoReply bool `json:"enableAutoReply,omitempty\"` // Flag that controls whether Gmail automatically replies to messages.

	EndTime int64 `json:"endTime,omitempty\"` // An optional end time for sending auto-replies (epoch ms). When this is specified, Gmail will automatically reply only to messages that it receives before the end time. If both `startTime` and `endTime` are specified, `startTime` must precede `endTime`.

	ResponseBodyHtml string `json:"responseBodyHtml,omitempty\"` // Response body in HTML format. Gmail will sanitize the HTML before storing it. If both `response_body_plain_text` and `response_body_html` are specified, `response_body_html` will be used.

	ResponseBodyPlainText string `json:"responseBodyPlainText,omitempty\"` // Response body in plain text format. If both `response_body_plain_text` and `response_body_html` are specified, `response_body_html` will be used.

	ResponseSubject string `json:"responseSubject,omitempty\"` // Optional text to prepend to the subject line in vacation responses. In order to enable auto-replies, either the response subject or the response body must be nonempty.

	RestrictToContacts bool `json:"restrictToContacts,omitempty\"` // Flag that determines whether responses are sent to recipients who are not in the user's list of contacts.

	RestrictToDomain bool `json:"restrictToDomain,omitempty\"` // Flag that determines whether responses are sent to recipients who are outside of the user's domain. This feature is only available for Google Workspace users.

	StartTime int64 `json:"startTime,omitempty\"` // An optional start time for sending auto-replies (epoch ms). When this is specified, Gmail will automatically reply only to messages that it receives after the start time. If both `startTime` and `endTime` are specified, `startTime` must precede `endTime`.

}

// Set up or update a new push notification watch on this user's mailbox.
type WatchRequest struct {
	LabelFilterAction string `json:"labelFilterAction,omitempty\"` // Filtering behavior of `labelIds list` specified. This field is deprecated because it caused incorrect behavior in some cases; use `label_filter_behavior` instead.

	LabelFilterBehavior string `json:"labelFilterBehavior,omitempty\"` // Filtering behavior of `labelIds list` specified. This field replaces `label_filter_action`; if set, `label_filter_action` is ignored.

	LabelIds []string `json:"labelIds,omitempty\"` // List of label_ids to restrict notifications about. By default, if unspecified, all changes are pushed out. If specified then dictates which labels are required for a push notification to be generated.

	TopicName string `json:"topicName,omitempty\"` // A fully qualified Google Cloud Pub/Sub API topic name to publish the events to. This topic name **must** already exist in Cloud Pub/Sub and you **must** have already granted gmail "publish" permission on it. For example, "projects/my-project-identifier/topics/my-topic-name" (using the Cloud Pub/Sub "v1" topic naming format). Note that the "my-project-identifier" portion must exactly match your Google developer project id (the one executing this watch request).

}

// Push notification watch response.
type WatchResponse struct {
	Expiration int64 `json:"expiration,omitempty\"` // When Gmail will stop sending notifications for mailbox updates (epoch millis). Call `watch` again before this time to renew the watch.

	HistoryId uint64 `json:"historyId,omitempty\"` // The ID of the mailbox's current history record.

}
