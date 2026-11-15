// People API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package people

// A person's physical address. May be a P.O. box or street address. All fields are optional.
type Address struct {
	City string `json:"city,omitempty\"` // The city of the address.

	Country string `json:"country,omitempty\"` // The country of the address.

	CountryCode string `json:"countryCode,omitempty\"` // The [ISO 3166-1 alpha-2](http://www.iso.org/iso/country_codes.htm) country code of the address.

	ExtendedAddress string `json:"extendedAddress,omitempty\"` // The extended address of the address; for example, the apartment number.

	FormattedType string `json:"formattedType,omitempty\"` // Output only. The type of the address translated and formatted in the viewer's account locale or the `Accept-Language` HTTP header locale.

	FormattedValue string `json:"formattedValue,omitempty\"` // The unstructured value of the address. If this is not set by the user it will be automatically constructed from structured values.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the address.

	PoBox string `json:"poBox,omitempty\"` // The P.O. box of the address.

	PostalCode string `json:"postalCode,omitempty\"` // The postal code of the address.

	Region string `json:"region,omitempty\"` // The region of the address; for example, the state or province.

	StreetAddress string `json:"streetAddress,omitempty\"` // The street address.

	TypeValue string `json:"type,omitempty\"` // The type of the address. The type can be custom or one of these predefined values: * `home` * `work` * `other`

}

// A person's age range.
type AgeRangeType struct {
	AgeRange string `json:"ageRange,omitempty\"` // The age range.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the age range.

}

// A request to create a batch of contacts.
type BatchCreateContactsRequest struct {
	Contacts []ContactToCreate `json:"contacts,omitempty\"` // Required. The contact to create. Allows up to 200 contacts in a single request.

	ReadMask string `json:"readMask,omitempty\"` // Required. A field mask to restrict which fields on each person are returned in the response. Multiple fields can be specified by separating them with commas. If read mask is left empty, the post-mutate-get is skipped and no data will be returned in the response. Valid values are: * addresses * ageRanges * biographies * birthdays * calendarUrls * clientData * coverPhotos * emailAddresses * events * externalIds * genders * imClients * interests * locales * locations * memberships * metadata * miscKeywords * names * nicknames * occupations * organizations * phoneNumbers * photos * relations * sipAddresses * skills * urls * userDefined

	Sources []string `json:"sources,omitempty\"` // Optional. A mask of what source types to return in the post mutate read. Defaults to READ_SOURCE_TYPE_CONTACT and READ_SOURCE_TYPE_PROFILE if not set.

}

// If not successful, returns BatchCreateContactsErrorDetails which contains a list of errors for each invalid contact. The response to a request to create a batch of contacts.
type BatchCreateContactsResponse struct {
	CreatedPeople []PersonResponse `json:"createdPeople,omitempty\"` // The contacts that were created, unless the request `read_mask` is empty.

}

// A request to delete a batch of existing contacts.
type BatchDeleteContactsRequest struct {
	ResourceNames []string `json:"resourceNames,omitempty\"` // Required. The resource names of the contact to delete. It's repeatable. Allows up to 500 resource names in a single request.

}

// The response to a batch get contact groups request.
type BatchGetContactGroupsResponse struct {
	Responses []ContactGroupResponse `json:"responses,omitempty\"` // The list of responses for each requested contact group resource.

}

// A request to update a batch of contacts.
type BatchUpdateContactsRequest struct {
	Contacts map[string]interface{} `json:"contacts,omitempty\"` // Required. A map of resource names to the person data to be updated. Allows up to 200 contacts in a single request.

	ReadMask string `json:"readMask,omitempty\"` // Required. A field mask to restrict which fields on each person are returned. Multiple fields can be specified by separating them with commas. If read mask is left empty, the post-mutate-get is skipped and no data will be returned in the response. Valid values are: * addresses * ageRanges * biographies * birthdays * calendarUrls * clientData * coverPhotos * emailAddresses * events * externalIds * genders * imClients * interests * locales * locations * memberships * metadata * miscKeywords * names * nicknames * occupations * organizations * phoneNumbers * photos * relations * sipAddresses * skills * urls * userDefined

	Sources []string `json:"sources,omitempty\"` // Optional. A mask of what source types to return. Defaults to READ_SOURCE_TYPE_CONTACT and READ_SOURCE_TYPE_PROFILE if not set.

	UpdateMask string `json:"updateMask,omitempty\"` // Required. A field mask to restrict which fields on the person are updated. Multiple fields can be specified by separating them with commas. All specified fields will be replaced, or cleared if left empty for each person. Valid values are: * addresses * biographies * birthdays * calendarUrls * clientData * emailAddresses * events * externalIds * genders * imClients * interests * locales * locations * memberships * miscKeywords * names * nicknames * occupations * organizations * phoneNumbers * relations * sipAddresses * urls * userDefined

}

// If not successful, returns BatchUpdateContactsErrorDetails, a list of errors corresponding to each contact. The response to a request to update a batch of contacts.
type BatchUpdateContactsResponse struct {
	UpdateResult map[string]interface{} `json:"updateResult,omitempty\"` // A map of resource names to the contacts that were updated, unless the request `read_mask` is empty.

}

// A person's short biography.
type Biography struct {
	ContentType string `json:"contentType,omitempty\"` // The content type of the biography.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the biography.

	Value string `json:"value,omitempty\"` // The short biography.

}

// A person's birthday. At least one of the `date` and `text` fields are specified. The `date` and `text` fields typically represent the same date, but are not guaranteed to. Clients should always set the `date` field when mutating birthdays.
type Birthday struct {
	Date Date `json:"date,omitempty\"` // The structured date of the birthday.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the birthday.

	Text string `json:"text,omitempty\"` // Prefer to use the `date` field if set. A free-form string representing the user's birthday. This value is not validated.

}

// **DEPRECATED**: No data will be returned A person's bragging rights.
type BraggingRights struct {
	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the bragging rights.

	Value string `json:"value,omitempty\"` // The bragging rights; for example, `climbed mount everest`.

}

// A person's calendar URL.
type CalendarUrl struct {
	FormattedType string `json:"formattedType,omitempty\"` // Output only. The type of the calendar URL translated and formatted in the viewer's account locale or the `Accept-Language` HTTP header locale.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the calendar URL.

	TypeValue string `json:"type,omitempty\"` // The type of the calendar URL. The type can be custom or one of these predefined values: * `home` * `freeBusy` * `work`

	Url string `json:"url,omitempty\"` // The calendar URL.

}

// Arbitrary client data that is populated by clients. Duplicate keys and values are allowed.
type ClientData struct {
	Key string `json:"key,omitempty\"` // The client specified key of the client data.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the client data.

	Value string `json:"value,omitempty\"` // The client specified value of the client data.

}

// A contact group.
type ContactGroup struct {
	ClientData []GroupClientData `json:"clientData,omitempty\"` // The group's client data.

	Etag string `json:"etag,omitempty\"` // The [HTTP entity tag](https://en.wikipedia.org/wiki/HTTP_ETag) of the resource. Used for web cache validation.

	FormattedName string `json:"formattedName,omitempty\"` // Output only. The name translated and formatted in the viewer's account locale or the `Accept-Language` HTTP header locale for system groups names. Group names set by the owner are the same as name.

	GroupType string `json:"groupType,omitempty\"` // Output only. The contact group type.

	MemberCount int `json:"memberCount,omitempty\"` // Output only. The total number of contacts in the group irrespective of max members in specified in the request.

	MemberResourceNames []string `json:"memberResourceNames,omitempty\"` // Output only. The list of contact person resource names that are members of the contact group. The field is only populated for GET requests and will only return as many members as `maxMembers` in the get request.

	Metadata ContactGroupMetadata `json:"metadata,omitempty\"` // Output only. Metadata about the contact group.

	Name string `json:"name,omitempty\"` // The contact group name set by the group owner or a system provided name for system groups. For [`contactGroups.create`](/people/api/rest/v1/contactGroups/create) or [`contactGroups.update`](/people/api/rest/v1/contactGroups/update) the name must be unique to the users contact groups. Attempting to create a group with a duplicate name will return a HTTP 409 error.

	ResourceName string `json:"resourceName,omitempty\"` // The resource name for the contact group, assigned by the server. An ASCII string, in the form of `contactGroups/{contact_group_id}`.

}

// A Google contact group membership.
type ContactGroupMembership struct {
	ContactGroupId string `json:"contactGroupId,omitempty\"` // Output only. The contact group ID for the contact group membership.

	ContactGroupResourceName string `json:"contactGroupResourceName,omitempty\"` // The resource name for the contact group, assigned by the server. An ASCII string, in the form of `contactGroups/{contact_group_id}`. Only contact_group_resource_name can be used for modifying memberships. Any contact group membership can be removed, but only user group or "myContacts" or "starred" system groups memberships can be added. A contact must always have at least one contact group membership.

}

// The metadata about a contact group.
type ContactGroupMetadata struct {
	Deleted bool `json:"deleted,omitempty\"` // Output only. True if the contact group resource has been deleted. Populated only for [`ListContactGroups`](/people/api/rest/v1/contactgroups/list) requests that include a sync token.

	UpdateTime string `json:"updateTime,omitempty\"` // Output only. The time the group was last updated.

}

// The response for a specific contact group.
type ContactGroupResponse struct {
	ContactGroup ContactGroup `json:"contactGroup,omitempty\"` // The contact group.

	RequestedResourceName string `json:"requestedResourceName,omitempty\"` // The original requested resource name.

	Status Status `json:"status,omitempty\"` // The status of the response.

}

// A wrapper that contains the person data to populate a newly created source.
type ContactToCreate struct {
	ContactPerson Person `json:"contactPerson,omitempty\"` // Required. The person data to populate a newly created source.

}

// A request to copy an "Other contact" to my contacts group.
type CopyOtherContactToMyContactsGroupRequest struct {
	CopyMask string `json:"copyMask,omitempty\"` // Required. A field mask to restrict which fields are copied into the new contact. Valid values are: * emailAddresses * names * phoneNumbers

	ReadMask string `json:"readMask,omitempty\"` // Optional. A field mask to restrict which fields on the person are returned. Multiple fields can be specified by separating them with commas. Defaults to the copy mask with metadata and membership fields if not set. Valid values are: * addresses * ageRanges * biographies * birthdays * calendarUrls * clientData * coverPhotos * emailAddresses * events * externalIds * genders * imClients * interests * locales * locations * memberships * metadata * miscKeywords * names * nicknames * occupations * organizations * phoneNumbers * photos * relations * sipAddresses * skills * urls * userDefined

	Sources []string `json:"sources,omitempty\"` // Optional. A mask of what source types to return. Defaults to READ_SOURCE_TYPE_CONTACT and READ_SOURCE_TYPE_PROFILE if not set.

}

// A person's cover photo. A large image shown on the person's profile page that represents who they are or what they care about.
type CoverPhoto struct {
	DefaultValue bool `json:"default,omitempty\"` // True if the cover photo is the default cover photo; false if the cover photo is a user-provided cover photo.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the cover photo.

	Url string `json:"url,omitempty\"` // The URL of the cover photo.

}

// A request to create a new contact group.
type CreateContactGroupRequest struct {
	ContactGroup ContactGroup `json:"contactGroup,omitempty\"` // Required. The contact group to create.

	ReadGroupFields string `json:"readGroupFields,omitempty\"` // Optional. A field mask to restrict which fields on the group are returned. Defaults to `metadata`, `groupType`, and `name` if not set or set to empty. Valid fields are: * clientData * groupType * metadata * name

}

// Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following: * A full date, with non-zero year, month, and day values. * A month and day, with a zero year (for example, an anniversary). * A year on its own, with a zero month and a zero day. * A year and month, with a zero day (for example, a credit card expiration date). Related types: * google.type.TimeOfDay * google.type.DateTime * google.protobuf.Timestamp
type Date struct {
	Day int `json:"day,omitempty\"` // Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.

	Month int `json:"month,omitempty\"` // Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.

	Year int `json:"year,omitempty\"` // Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.

}

// The response for deleting a contact's photo.
type DeleteContactPhotoResponse struct {
	Person Person `json:"person,omitempty\"` // The updated person, if person_fields is set in the DeleteContactPhotoRequest; otherwise this will be unset.

}

// A Google Workspace Domain membership.
type DomainMembership struct {
	InViewerDomain bool `json:"inViewerDomain,omitempty\"` // True if the person is in the viewer's Google Workspace domain.

}

// A person's email address.
type EmailAddress struct {
	DisplayName string `json:"displayName,omitempty\"` // The display name of the email.

	FormattedType string `json:"formattedType,omitempty\"` // Output only. The type of the email address translated and formatted in the viewer's account locale or the `Accept-Language` HTTP header locale.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the email address.

	TypeValue string `json:"type,omitempty\"` // The type of the email address. The type can be custom or one of these predefined values: * `home` * `work` * `other`

	Value string `json:"value,omitempty\"` // The email address.

}

// A generic empty message that you can re-use to avoid defining duplicated empty messages in your APIs. A typical example is to use it as the request or the response type of an API method. For instance: service Foo { rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty); }
type Empty struct {
}

// An event related to the person.
type Event struct {
	Date Date `json:"date,omitempty\"` // The date of the event.

	FormattedType string `json:"formattedType,omitempty\"` // Output only. The type of the event translated and formatted in the viewer's account locale or the `Accept-Language` HTTP header locale.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the event.

	TypeValue string `json:"type,omitempty\"` // The type of the event. The type can be custom or one of these predefined values: * `anniversary` * `other`

}

// An identifier from an external entity related to the person.
type ExternalId struct {
	FormattedType string `json:"formattedType,omitempty\"` // Output only. The type of the event translated and formatted in the viewer's account locale or the `Accept-Language` HTTP header locale.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the external ID.

	TypeValue string `json:"type,omitempty\"` // The type of the external ID. The type can be custom or one of these predefined values: * `account` * `customer` * `loginId` * `network` * `organization`

	Value string `json:"value,omitempty\"` // The value of the external ID.

}

// Metadata about a field.
type FieldMetadata struct {
	Primary bool `json:"primary,omitempty\"` // Output only. True if the field is the primary field for all sources in the person. Each person will have at most one field with `primary` set to true.

	Source Source `json:"source,omitempty\"` // The source of the field.

	SourcePrimary bool `json:"sourcePrimary,omitempty\"` // True if the field is the primary field for the source. Each source must have at most one field with `source_primary` set to true.

	Verified bool `json:"verified,omitempty\"` // Output only. True if the field is verified; false if the field is unverified. A verified field is typically a name, email address, phone number, or website that has been confirmed to be owned by the person.

}

// The name that should be used to sort the person in a list.
type FileAs struct {
	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the file-as.

	Value string `json:"value,omitempty\"` // The file-as value

}

// A person's gender.
type Gender struct {
	AddressMeAs string `json:"addressMeAs,omitempty\"` // Free form text field for pronouns that should be used to address the person. Common values are: * `he`/`him` * `she`/`her` * `they`/`them`

	FormattedValue string `json:"formattedValue,omitempty\"` // Output only. The value of the gender translated and formatted in the viewer's account locale or the `Accept-Language` HTTP header locale. Unspecified or custom value are not localized.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the gender.

	Value string `json:"value,omitempty\"` // The gender for the person. The gender can be custom or one of these predefined values: * `male` * `female` * `unspecified`

}

// The response to a get request for a list of people by resource name.
type GetPeopleResponse struct {
	Responses []PersonResponse `json:"responses,omitempty\"` // The response for each requested resource name.

}

// Arbitrary client data that is populated by clients. Duplicate keys and values are allowed.
type GroupClientData struct {
	Key string `json:"key,omitempty\"` // The client specified key of the client data.

	Value string `json:"value,omitempty\"` // The client specified value of the client data.

}

// A person's instant messaging client.
type ImClient struct {
	FormattedProtocol string `json:"formattedProtocol,omitempty\"` // Output only. The protocol of the IM client formatted in the viewer's account locale or the `Accept-Language` HTTP header locale.

	FormattedType string `json:"formattedType,omitempty\"` // Output only. The type of the IM client translated and formatted in the viewer's account locale or the `Accept-Language` HTTP header locale.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the IM client.

	Protocol string `json:"protocol,omitempty\"` // The protocol of the IM client. The protocol can be custom or one of these predefined values: * `aim` * `msn` * `yahoo` * `skype` * `qq` * `googleTalk` * `icq` * `jabber` * `netMeeting`

	TypeValue string `json:"type,omitempty\"` // The type of the IM client. The type can be custom or one of these predefined values: * `home` * `work` * `other`

	Username string `json:"username,omitempty\"` // The user name used in the IM client.

}

// One of the person's interests.
type Interest struct {
	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the interest.

	Value string `json:"value,omitempty\"` // The interest; for example, `stargazing`.

}

// The response to a request for the authenticated user's connections.
type ListConnectionsResponse struct {
	Connections []Person `json:"connections,omitempty\"` // The list of people that the requestor is connected to.

	NextPageToken string `json:"nextPageToken,omitempty\"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.

	NextSyncToken string `json:"nextSyncToken,omitempty\"` // A token, which can be sent as `sync_token` to retrieve changes since the last request. Request must set `request_sync_token` to return the sync token. When the response is paginated, only the last page will contain `nextSyncToken`.

	TotalItems int `json:"totalItems,omitempty\"` // The total number of items in the list without pagination.

	TotalPeople int `json:"totalPeople,omitempty\"` // **DEPRECATED** (Please use totalItems) The total number of people in the list without pagination.

}

// The response to a list contact groups request.
type ListContactGroupsResponse struct {
	ContactGroups []ContactGroup `json:"contactGroups,omitempty\"` // The list of contact groups. Members of the contact groups are not populated.

	NextPageToken string `json:"nextPageToken,omitempty\"` // The token that can be used to retrieve the next page of results.

	NextSyncToken string `json:"nextSyncToken,omitempty\"` // The token that can be used to retrieve changes since the last request.

	TotalItems int `json:"totalItems,omitempty\"` // The total number of items in the list without pagination.

}

// The response to a request for the authenticated user's domain directory.
type ListDirectoryPeopleResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty\"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.

	NextSyncToken string `json:"nextSyncToken,omitempty\"` // A token, which can be sent as `sync_token` to retrieve changes since the last request. Request must set `request_sync_token` to return the sync token.

	People []Person `json:"people,omitempty\"` // The list of people in the domain directory.

}

// The response to a request for the authenticated user's "Other contacts".
type ListOtherContactsResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty\"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.

	NextSyncToken string `json:"nextSyncToken,omitempty\"` // A token, which can be sent as `sync_token` to retrieve changes since the last request. Request must set `request_sync_token` to return the sync token.

	OtherContacts []Person `json:"otherContacts,omitempty\"` // The list of "Other contacts" returned as Person resources. "Other contacts" support a limited subset of fields. See ListOtherContactsRequest.request_mask for more detailed information.

	TotalSize int `json:"totalSize,omitempty\"` // The total number of other contacts in the list without pagination.

}

// A person's locale preference.
type Locale struct {
	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the locale.

	Value string `json:"value,omitempty\"` // The well-formed [IETF BCP 47](https://tools.ietf.org/html/bcp47) language tag representing the locale.

}

// A person's location.
type Location struct {
	BuildingId string `json:"buildingId,omitempty\"` // The building identifier.

	Current bool `json:"current,omitempty\"` // Whether the location is the current location.

	DeskCode string `json:"deskCode,omitempty\"` // The individual desk location.

	Floor string `json:"floor,omitempty\"` // The floor name or number.

	FloorSection string `json:"floorSection,omitempty\"` // The floor section in `floor_name`.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the location.

	TypeValue string `json:"type,omitempty\"` // The type of the location. The type can be custom or one of these predefined values: * `desk` * `grewUp`

	Value string `json:"value,omitempty\"` // The free-form value of the location.

}

// A person's membership in a group. Only contact group memberships can be modified.
type Membership struct {
	ContactGroupMembership ContactGroupMembership `json:"contactGroupMembership,omitempty\"` // The contact group membership.

	DomainMembership DomainMembership `json:"domainMembership,omitempty\"` // Output only. The domain membership.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the membership.

}

// A person's miscellaneous keyword.
type MiscKeyword struct {
	FormattedType string `json:"formattedType,omitempty\"` // Output only. The type of the miscellaneous keyword translated and formatted in the viewer's account locale or the `Accept-Language` HTTP header locale.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the miscellaneous keyword.

	TypeValue string `json:"type,omitempty\"` // The miscellaneous keyword type.

	Value string `json:"value,omitempty\"` // The value of the miscellaneous keyword.

}

// A request to modify an existing contact group's members. Contacts can be removed from any group but they can only be added to a user group or "myContacts" or "starred" system groups.
type ModifyContactGroupMembersRequest struct {
	ResourceNamesToAdd []string `json:"resourceNamesToAdd,omitempty\"` // Optional. The resource names of the contact people to add in the form of `people/{person_id}`. The total number of resource names in `resource_names_to_add` and `resource_names_to_remove` must be less than or equal to 1000.

	ResourceNamesToRemove []string `json:"resourceNamesToRemove,omitempty\"` // Optional. The resource names of the contact people to remove in the form of `people/{person_id}`. The total number of resource names in `resource_names_to_add` and `resource_names_to_remove` must be less than or equal to 1000.

}

// The response to a modify contact group members request.
type ModifyContactGroupMembersResponse struct {
	CanNotRemoveLastContactGroupResourceNames []string `json:"canNotRemoveLastContactGroupResourceNames,omitempty\"` // The contact people resource names that cannot be removed from their last contact group.

	NotFoundResourceNames []string `json:"notFoundResourceNames,omitempty\"` // The contact people resource names that were not found.

}

// A person's name. If the name is a mononym, the family name is empty.
type Name struct {
	DisplayName string `json:"displayName,omitempty\"` // Output only. The display name formatted according to the locale specified by the viewer's account or the `Accept-Language` HTTP header.

	DisplayNameLastFirst string `json:"displayNameLastFirst,omitempty\"` // Output only. The display name with the last name first formatted according to the locale specified by the viewer's account or the `Accept-Language` HTTP header.

	FamilyName string `json:"familyName,omitempty\"` // The family name.

	GivenName string `json:"givenName,omitempty\"` // The given name.

	HonorificPrefix string `json:"honorificPrefix,omitempty\"` // The honorific prefixes, such as `Mrs.` or `Dr.`

	HonorificSuffix string `json:"honorificSuffix,omitempty\"` // The honorific suffixes, such as `Jr.`

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the name.

	MiddleName string `json:"middleName,omitempty\"` // The middle name(s).

	PhoneticFamilyName string `json:"phoneticFamilyName,omitempty\"` // The family name spelled as it sounds.

	PhoneticFullName string `json:"phoneticFullName,omitempty\"` // The full name spelled as it sounds.

	PhoneticGivenName string `json:"phoneticGivenName,omitempty\"` // The given name spelled as it sounds.

	PhoneticHonorificPrefix string `json:"phoneticHonorificPrefix,omitempty\"` // The honorific prefixes spelled as they sound.

	PhoneticHonorificSuffix string `json:"phoneticHonorificSuffix,omitempty\"` // The honorific suffixes spelled as they sound.

	PhoneticMiddleName string `json:"phoneticMiddleName,omitempty\"` // The middle name(s) spelled as they sound.

	UnstructuredName string `json:"unstructuredName,omitempty\"` // The free form name value.

}

// A person's nickname.
type Nickname struct {
	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the nickname.

	TypeValue string `json:"type,omitempty\"` // The type of the nickname.

	Value string `json:"value,omitempty\"` // The nickname.

}

// A person's occupation.
type Occupation struct {
	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the occupation.

	Value string `json:"value,omitempty\"` // The occupation; for example, `carpenter`.

}

// A person's past or current organization. Overlapping date ranges are permitted.
type Organization struct {
	CostCenter string `json:"costCenter,omitempty\"` // The person's cost center at the organization.

	Current bool `json:"current,omitempty\"` // True if the organization is the person's current organization; false if the organization is a past organization.

	Department string `json:"department,omitempty\"` // The person's department at the organization.

	Domain string `json:"domain,omitempty\"` // The domain name associated with the organization; for example, `google.com`.

	EndDate Date `json:"endDate,omitempty\"` // The end date when the person left the organization.

	FormattedType string `json:"formattedType,omitempty\"` // Output only. The type of the organization translated and formatted in the viewer's account locale or the `Accept-Language` HTTP header locale.

	FullTimeEquivalentMillipercent int `json:"fullTimeEquivalentMillipercent,omitempty\"` // The person's full-time equivalent millipercent within the organization (100000 = 100%).

	JobDescription string `json:"jobDescription,omitempty\"` // The person's job description at the organization.

	Location string `json:"location,omitempty\"` // The location of the organization office the person works at.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the organization.

	Name string `json:"name,omitempty\"` // The name of the organization.

	PhoneticName string `json:"phoneticName,omitempty\"` // The phonetic name of the organization.

	StartDate Date `json:"startDate,omitempty\"` // The start date when the person joined the organization.

	Symbol string `json:"symbol,omitempty\"` // The symbol associated with the organization; for example, a stock ticker symbol, abbreviation, or acronym.

	Title string `json:"title,omitempty\"` // The person's job title at the organization.

	TypeValue string `json:"type,omitempty\"` // The type of the organization. The type can be custom or one of these predefined values: * `work` * `school`

}

// Information about a person merged from various data sources such as the authenticated user's contacts and profile data. Most fields can have multiple items. The items in a field have no guaranteed order, but each non-empty field is guaranteed to have exactly one field with `metadata.primary` set to true.
type Person struct {
	Addresses []Address `json:"addresses,omitempty\"` // The person's street addresses.

	AgeRange string `json:"ageRange,omitempty\"` // Output only. **DEPRECATED** (Please use `person.ageRanges` instead) The person's age range.

	AgeRanges []AgeRangeType `json:"ageRanges,omitempty\"` // Output only. The person's age ranges.

	Biographies []Biography `json:"biographies,omitempty\"` // The person's biographies. This field is a singleton for contact sources.

	Birthdays []Birthday `json:"birthdays,omitempty\"` // The person's birthdays. This field is a singleton for contact sources.

	BraggingRights []BraggingRights `json:"braggingRights,omitempty\"` // **DEPRECATED**: No data will be returned The person's bragging rights.

	CalendarUrls []CalendarUrl `json:"calendarUrls,omitempty\"` // The person's calendar URLs.

	ClientData []ClientData `json:"clientData,omitempty\"` // The person's client data.

	CoverPhotos []CoverPhoto `json:"coverPhotos,omitempty\"` // Output only. The person's cover photos.

	EmailAddresses []EmailAddress `json:"emailAddresses,omitempty\"` // The person's email addresses. For `people.connections.list` and `otherContacts.list` the number of email addresses is limited to 100. If a Person has more email addresses the entire set can be obtained by calling GetPeople.

	Etag string `json:"etag,omitempty\"` // The [HTTP entity tag](https://en.wikipedia.org/wiki/HTTP_ETag) of the resource. Used for web cache validation.

	Events []Event `json:"events,omitempty\"` // The person's events.

	ExternalIds []ExternalId `json:"externalIds,omitempty\"` // The person's external IDs.

	FileAses []FileAs `json:"fileAses,omitempty\"` // The person's file-ases.

	Genders []Gender `json:"genders,omitempty\"` // The person's genders. This field is a singleton for contact sources.

	ImClients []ImClient `json:"imClients,omitempty\"` // The person's instant messaging clients.

	Interests []Interest `json:"interests,omitempty\"` // The person's interests.

	Locales []Locale `json:"locales,omitempty\"` // The person's locale preferences.

	Locations []Location `json:"locations,omitempty\"` // The person's locations.

	Memberships []Membership `json:"memberships,omitempty\"` // The person's group memberships.

	Metadata PersonMetadata `json:"metadata,omitempty\"` // Output only. Metadata about the person.

	MiscKeywords []MiscKeyword `json:"miscKeywords,omitempty\"` // The person's miscellaneous keywords.

	Names []Name `json:"names,omitempty\"` // The person's names. This field is a singleton for contact sources.

	Nicknames []Nickname `json:"nicknames,omitempty\"` // The person's nicknames.

	Occupations []Occupation `json:"occupations,omitempty\"` // The person's occupations.

	Organizations []Organization `json:"organizations,omitempty\"` // The person's past or current organizations.

	PhoneNumbers []PhoneNumber `json:"phoneNumbers,omitempty\"` // The person's phone numbers. For `people.connections.list` and `otherContacts.list` the number of phone numbers is limited to 100. If a Person has more phone numbers the entire set can be obtained by calling GetPeople.

	Photos []Photo `json:"photos,omitempty\"` // Output only. The person's photos.

	Relations []Relation `json:"relations,omitempty\"` // The person's relations.

	RelationshipInterests []RelationshipInterest `json:"relationshipInterests,omitempty\"` // Output only. **DEPRECATED**: No data will be returned The person's relationship interests.

	RelationshipStatuses []RelationshipStatus `json:"relationshipStatuses,omitempty\"` // Output only. **DEPRECATED**: No data will be returned The person's relationship statuses.

	Residences []Residence `json:"residences,omitempty\"` // **DEPRECATED**: (Please use `person.locations` instead) The person's residences.

	ResourceName string `json:"resourceName,omitempty\"` // The resource name for the person, assigned by the server. An ASCII string in the form of `people/{person_id}`.

	SipAddresses []SipAddress `json:"sipAddresses,omitempty\"` // The person's SIP addresses.

	Skills []Skill `json:"skills,omitempty\"` // The person's skills.

	Taglines []Tagline `json:"taglines,omitempty\"` // Output only. **DEPRECATED**: No data will be returned The person's taglines.

	Urls []Url `json:"urls,omitempty\"` // The person's associated URLs.

	UserDefined []UserDefined `json:"userDefined,omitempty\"` // The person's user defined data.

}

// The metadata about a person.
type PersonMetadata struct {
	Deleted bool `json:"deleted,omitempty\"` // Output only. True if the person resource has been deleted. Populated only for `people.connections.list` and `otherContacts.list` sync requests.

	LinkedPeopleResourceNames []string `json:"linkedPeopleResourceNames,omitempty\"` // Output only. Resource names of people linked to this resource.

	ObjectType string `json:"objectType,omitempty\"` // Output only. **DEPRECATED** (Please use `person.metadata.sources.profileMetadata.objectType` instead) The type of the person object.

	PreviousResourceNames []string `json:"previousResourceNames,omitempty\"` // Output only. Any former resource names this person has had. Populated only for `people.connections.list` requests that include a sync token. The resource name may change when adding or removing fields that link a contact and profile such as a verified email, verified phone number, or profile URL.

	Sources []Source `json:"sources,omitempty\"` // The sources of data for the person.

}

// The response for a single person
type PersonResponse struct {
	HttpStatusCode int `json:"httpStatusCode,omitempty\"` // **DEPRECATED** (Please use status instead) [HTTP 1.1 status code] (http://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html).

	Person Person `json:"person,omitempty\"` // The person.

	RequestedResourceName string `json:"requestedResourceName,omitempty\"` // The original requested resource name. May be different than the resource name on the returned person. The resource name can change when adding or removing fields that link a contact and profile such as a verified email, verified phone number, or a profile URL.

	Status Status `json:"status,omitempty\"` // The status of the response.

}

// A person's phone number.
type PhoneNumber struct {
	CanonicalForm string `json:"canonicalForm,omitempty\"` // Output only. The canonicalized [ITU-T E.164](https://law.resource.org/pub/us/cfr/ibr/004/itu-t.E.164.1.2008.pdf) form of the phone number.

	FormattedType string `json:"formattedType,omitempty\"` // Output only. The type of the phone number translated and formatted in the viewer's account locale or the `Accept-Language` HTTP header locale.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the phone number.

	TypeValue string `json:"type,omitempty\"` // The type of the phone number. The type can be custom or one of these predefined values: * `home` * `work` * `mobile` * `homeFax` * `workFax` * `otherFax` * `pager` * `workMobile` * `workPager` * `main` * `googleVoice` * `other`

	Value string `json:"value,omitempty\"` // The phone number.

}

// A person's photo. A picture shown next to the person's name to help others recognize the person.
type Photo struct {
	DefaultValue bool `json:"default,omitempty\"` // True if the photo is a default photo; false if the photo is a user-provided photo.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the photo.

	Url string `json:"url,omitempty\"` // The URL of the photo. You can change the desired size by appending a query parameter `sz={size}` at the end of the url, where {size} is the size in pixels. Example: https://lh3.googleusercontent.com/-T_wVWLlmg7w/AAAAAAAAAAI/AAAAAAAABa8/00gzXvDBYqw/s100/photo.jpg?sz=50

}

// The metadata about a profile.
type ProfileMetadata struct {
	ObjectType string `json:"objectType,omitempty\"` // Output only. The profile object type.

	UserTypes []string `json:"userTypes,omitempty\"` // Output only. The user types.

}

// A person's relation to another person.
type Relation struct {
	FormattedType string `json:"formattedType,omitempty\"` // Output only. The type of the relation translated and formatted in the viewer's account locale or the locale specified in the Accept-Language HTTP header.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the relation.

	Person string `json:"person,omitempty\"` // The name of the other person this relation refers to.

	TypeValue string `json:"type,omitempty\"` // The person's relation to the other person. The type can be custom or one of these predefined values: * `spouse` * `child` * `mother` * `father` * `parent` * `brother` * `sister` * `friend` * `relative` * `domesticPartner` * `manager` * `assistant` * `referredBy` * `partner`

}

// **DEPRECATED**: No data will be returned A person's relationship interest .
type RelationshipInterest struct {
	FormattedValue string `json:"formattedValue,omitempty\"` // Output only. The value of the relationship interest translated and formatted in the viewer's account locale or the locale specified in the Accept-Language HTTP header.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the relationship interest.

	Value string `json:"value,omitempty\"` // The kind of relationship the person is looking for. The value can be custom or one of these predefined values: * `friend` * `date` * `relationship` * `networking`

}

// **DEPRECATED**: No data will be returned A person's relationship status.
type RelationshipStatus struct {
	FormattedValue string `json:"formattedValue,omitempty\"` // Output only. The value of the relationship status translated and formatted in the viewer's account locale or the `Accept-Language` HTTP header locale.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the relationship status.

	Value string `json:"value,omitempty\"` // The relationship status. The value can be custom or one of these predefined values: * `single` * `inARelationship` * `engaged` * `married` * `itsComplicated` * `openRelationship` * `widowed` * `inDomesticPartnership` * `inCivilUnion`

}

// **DEPRECATED**: Please use `person.locations` instead. A person's past or current residence.
type Residence struct {
	Current bool `json:"current,omitempty\"` // True if the residence is the person's current residence; false if the residence is a past residence.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the residence.

	Value string `json:"value,omitempty\"` // The address of the residence.

}

// The response to a request for people in the authenticated user's domain directory that match the specified query.
type SearchDirectoryPeopleResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty\"` // A token, which can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.

	People []Person `json:"people,omitempty\"` // The list of people in the domain directory that match the query.

	TotalSize int `json:"totalSize,omitempty\"` // The total number of items in the list without pagination.

}

// The response to a search request for the authenticated user, given a query.
type SearchResponse struct {
	Results []SearchResult `json:"results,omitempty\"` // The results of the request.

}

// A result of a search query.
type SearchResult struct {
	Person Person `json:"person,omitempty\"` // The matched Person.

}

// A person's SIP address. Session Initial Protocol addresses are used for VoIP communications to make voice or video calls over the internet.
type SipAddress struct {
	FormattedType string `json:"formattedType,omitempty\"` // Output only. The type of the SIP address translated and formatted in the viewer's account locale or the `Accept-Language` HTTP header locale.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the SIP address.

	TypeValue string `json:"type,omitempty\"` // The type of the SIP address. The type can be custom or or one of these predefined values: * `home` * `work` * `mobile` * `other`

	Value string `json:"value,omitempty\"` // The SIP address in the [RFC 3261 19.1](https://tools.ietf.org/html/rfc3261#section-19.1) SIP URI format.

}

// A skill that the person has.
type Skill struct {
	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the skill.

	Value string `json:"value,omitempty\"` // The skill; for example, `underwater basket weaving`.

}

// The source of a field.
type Source struct {
	Etag string `json:"etag,omitempty\"` // **Only populated in `person.metadata.sources`.** The [HTTP entity tag](https://en.wikipedia.org/wiki/HTTP_ETag) of the source. Used for web cache validation.

	Id string `json:"id,omitempty\"` // The unique identifier within the source type generated by the server.

	ProfileMetadata ProfileMetadata `json:"profileMetadata,omitempty\"` // Output only. **Only populated in `person.metadata.sources`.** Metadata about a source of type PROFILE.

	TypeValue string `json:"type,omitempty\"` // The source type.

	UpdateTime string `json:"updateTime,omitempty\"` // Output only. **Only populated in `person.metadata.sources`.** Last update timestamp of this source.

}

// The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
type Status struct {
	Code int `json:"code,omitempty\"` // The status code, which should be an enum value of google.rpc.Code.

	Details []map[string]interface{} `json:"details,omitempty\"` // A list of messages that carry the error details. There is a common set of message types for APIs to use.

	Message string `json:"message,omitempty\"` // A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.

}

// **DEPRECATED**: No data will be returned A brief one-line description of the person.
type Tagline struct {
	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the tagline.

	Value string `json:"value,omitempty\"` // The tagline.

}

// A request to update an existing user contact group. All updated fields will be replaced.
type UpdateContactGroupRequest struct {
	ContactGroup ContactGroup `json:"contactGroup,omitempty\"` // Required. The contact group to update.

	ReadGroupFields string `json:"readGroupFields,omitempty\"` // Optional. A field mask to restrict which fields on the group are returned. Defaults to `metadata`, `groupType`, and `name` if not set or set to empty. Valid fields are: * clientData * groupType * memberCount * metadata * name

	UpdateGroupFields string `json:"updateGroupFields,omitempty\"` // Optional. A field mask to restrict which fields on the group are updated. Multiple fields can be specified by separating them with commas. Defaults to `name` if not set or set to empty. Updated fields are replaced. Valid values are: * clientData * name

}

// A request to update an existing contact's photo. All requests must have a valid photo format: JPEG or PNG.
type UpdateContactPhotoRequest struct {
	PersonFields string `json:"personFields,omitempty\"` // Optional. A field mask to restrict which fields on the person are returned. Multiple fields can be specified by separating them with commas. Defaults to empty if not set, which will skip the post mutate get. Valid values are: * addresses * ageRanges * biographies * birthdays * calendarUrls * clientData * coverPhotos * emailAddresses * events * externalIds * genders * imClients * interests * locales * locations * memberships * metadata * miscKeywords * names * nicknames * occupations * organizations * phoneNumbers * photos * relations * sipAddresses * skills * urls * userDefined

	PhotoBytes string `json:"photoBytes,omitempty\"` // Required. Raw photo bytes

	Sources []string `json:"sources,omitempty\"` // Optional. A mask of what source types to return. Defaults to READ_SOURCE_TYPE_CONTACT and READ_SOURCE_TYPE_PROFILE if not set.

}

// The response for updating a contact's photo.
type UpdateContactPhotoResponse struct {
	Person Person `json:"person,omitempty\"` // The updated person, if person_fields is set in the UpdateContactPhotoRequest; otherwise this will be unset.

}

// A person's associated URLs.
type Url struct {
	FormattedType string `json:"formattedType,omitempty\"` // Output only. The type of the URL translated and formatted in the viewer's account locale or the `Accept-Language` HTTP header locale.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the URL.

	TypeValue string `json:"type,omitempty\"` // The type of the URL. The type can be custom or one of these predefined values: * `home` * `work` * `blog` * `profile` * `homePage` * `ftp` * `reservations` * `appInstallPage`: website for a Currents application. * `other`

	Value string `json:"value,omitempty\"` // The URL.

}

// Arbitrary user data that is populated by the end users.
type UserDefined struct {
	Key string `json:"key,omitempty\"` // The end user specified key of the user defined data.

	Metadata FieldMetadata `json:"metadata,omitempty\"` // Metadata about the user defined data.

	Value string `json:"value,omitempty\"` // The end user specified value of the user defined data.

}
