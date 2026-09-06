// Drive Labels API
//
// Code generated from Google Discovery API. DO NOT EDIT.

package drivelabels

// The color derived from BadgeConfig and changed to the closest recommended supported color.
type GoogleAppsDriveLabelsV2BadgeColors struct {
	BackgroundColor GoogleTypeColor `json:"backgroundColor,omitempty\"` // Output only. Badge background that pairs with the foreground.

	ForegroundColor GoogleTypeColor `json:"foregroundColor,omitempty\"` // Output only. Badge foreground that pairs with the background.

	SoloColor GoogleTypeColor `json:"soloColor,omitempty\"` // Output only. Color that can be used for text without a background.

}

// Badge status of the label.
type GoogleAppsDriveLabelsV2BadgeConfig struct {
	Color GoogleTypeColor `json:"color,omitempty\"` // The color of the badge. When not specified, no badge is rendered. The background, foreground, and solo (light and dark mode) colors set here are changed in the Drive UI into the closest recommended supported color.

	PriorityOverride int64 `json:"priorityOverride,omitempty\"` // Override the default global priority of this badge. When set to 0, the default priority heuristic is used.

}

// Deletes one or more label permissions.
type GoogleAppsDriveLabelsV2BatchDeleteLabelPermissionsRequest struct {
	Requests []GoogleAppsDriveLabelsV2DeleteLabelPermissionRequest `json:"requests,omitempty\"` // Required. The request message specifying the resources to update.

	UseAdminAccess bool `json:"useAdminAccess,omitempty\"` // Set to `true` in order to use the user's admin credentials. The server will verify the user is an admin for the label before allowing access. If this is set, the `use_admin_access` field in the `DeleteLabelPermissionRequest` messages must either be empty or match this field.

}

// Updates one or more label permissions.
type GoogleAppsDriveLabelsV2BatchUpdateLabelPermissionsRequest struct {
	Requests []GoogleAppsDriveLabelsV2UpdateLabelPermissionRequest `json:"requests,omitempty\"` // Required. The request message specifying the resources to update.

	UseAdminAccess bool `json:"useAdminAccess,omitempty\"` // Set to `true` in order to use the user's admin credentials. The server will verify the user is an admin for the label before allowing access. If this is set, the `use_admin_access` field in the `UpdateLabelPermissionRequest` messages must either be empty or match this field.

}

// Response for updating one or more label permissions.
type GoogleAppsDriveLabelsV2BatchUpdateLabelPermissionsResponse struct {
	Permissions []GoogleAppsDriveLabelsV2LabelPermission `json:"permissions,omitempty\"` // Required. Permissions updated.

}

// Limits for date field type.
type GoogleAppsDriveLabelsV2DateLimits struct {
	MaxValue GoogleTypeDate `json:"maxValue,omitempty\"` // Maximum value for the date field type.

	MinValue GoogleTypeDate `json:"minValue,omitempty\"` // Minimum value for the date field type.

}

// Deletes a label permission. Permissions affect the label resource as a whole, aren't revisioned, and don't require publishing.
type GoogleAppsDriveLabelsV2DeleteLabelPermissionRequest struct {
	Name string `json:"name,omitempty\"` // Required. Label permission resource name.

	UseAdminAccess bool `json:"useAdminAccess,omitempty\"` // Set to `true` in order to use the user's admin credentials. The server will verify the user is an admin for the label before allowing access.

}

// The set of requests for updating aspects of a label. If any request isn't valid, no requests will be applied.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelRequest struct {
	LanguageCode string `json:"languageCode,omitempty\"` // The BCP-47 language code to use for evaluating localized field labels when `include_label_in_response` is `true`.

	Requests []GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestRequest `json:"requests,omitempty\"` // A list of updates to apply to the label. Requests will be applied in the order they are specified.

	UseAdminAccess bool `json:"useAdminAccess,omitempty\"` // Set to `true` in order to use the user's admin credentials. The server will verify the user is an admin for the label before allowing access.

	View string `json:"view,omitempty\"` // When specified, only certain fields belonging to the indicated view will be returned.

	WriteControl GoogleAppsDriveLabelsV2WriteControl `json:"writeControl,omitempty\"` // Provides control over how write requests are executed.

}

// Request to create a field within a label.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestCreateFieldRequest struct {
	Field GoogleAppsDriveLabelsV2Field `json:"field,omitempty\"` // Required. Field to create.

}

// Request to create a selection choice.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestCreateSelectionChoiceRequest struct {
	Choice GoogleAppsDriveLabelsV2FieldSelectionOptionsChoice `json:"choice,omitempty\"` // Required. The choice to create.

	FieldId string `json:"fieldId,omitempty\"` // Required. The selection field in which a choice will be created.

}

// Request to delete the field.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestDeleteFieldRequest struct {
	Id string `json:"id,omitempty\"` // Required. ID of the field to delete.

}

// Request to delete a choice.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestDeleteSelectionChoiceRequest struct {
	FieldId string `json:"fieldId,omitempty\"` // Required. The selection field from which a choice will be deleted.

	Id string `json:"id,omitempty\"` // Required. Choice to delete.

}

// Request to disable the field.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestDisableFieldRequest struct {
	DisabledPolicy GoogleAppsDriveLabelsV2LifecycleDisabledPolicy `json:"disabledPolicy,omitempty\"` // Required. Field disabled policy.

	Id string `json:"id,omitempty\"` // Required. Key of the field to disable.

	UpdateMask string `json:"updateMask,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `disabled_policy` is implied and should not be specified. A single `*` can be used as a short-hand for updating every field.

}

// Request to disable a choice.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestDisableSelectionChoiceRequest struct {
	DisabledPolicy GoogleAppsDriveLabelsV2LifecycleDisabledPolicy `json:"disabledPolicy,omitempty\"` // Required. The disabled policy to update.

	FieldId string `json:"fieldId,omitempty\"` // Required. The selection field in which a choice will be disabled.

	Id string `json:"id,omitempty\"` // Required. Choice to disable.

	UpdateMask string `json:"updateMask,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `disabled_policy` is implied and should not be specified. A single `*` can be used as a short-hand for updating every field.

}

// Request to enable the field.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestEnableFieldRequest struct {
	Id string `json:"id,omitempty\"` // Required. ID of the field to enable.

}

// Request to enable a choice.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestEnableSelectionChoiceRequest struct {
	FieldId string `json:"fieldId,omitempty\"` // Required. The selection field in which a choice will be enabled.

	Id string `json:"id,omitempty\"` // Required. Choice to enable.

}

// A single kind of update to apply to a label.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestRequest struct {
	CreateField GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestCreateFieldRequest `json:"createField,omitempty\"` // Creates a field.

	CreateSelectionChoice GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestCreateSelectionChoiceRequest `json:"createSelectionChoice,omitempty\"` // Create a choice within a selection field.

	DeleteField GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestDeleteFieldRequest `json:"deleteField,omitempty\"` // Deletes a field from the label.

	DeleteSelectionChoice GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestDeleteSelectionChoiceRequest `json:"deleteSelectionChoice,omitempty\"` // Delete a choice within a selection field.

	DisableField GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestDisableFieldRequest `json:"disableField,omitempty\"` // Disables the field.

	DisableSelectionChoice GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestDisableSelectionChoiceRequest `json:"disableSelectionChoice,omitempty\"` // Disable a choice within a selection field.

	EnableField GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestEnableFieldRequest `json:"enableField,omitempty\"` // Enables the field.

	EnableSelectionChoice GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestEnableSelectionChoiceRequest `json:"enableSelectionChoice,omitempty\"` // Enable a choice within a selection field.

	UpdateField GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestUpdateFieldPropertiesRequest `json:"updateField,omitempty\"` // Updates basic properties of a field.

	UpdateFieldType GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestUpdateFieldTypeRequest `json:"updateFieldType,omitempty\"` // Update field type and/or type options.

	UpdateLabel GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestUpdateLabelPropertiesRequest `json:"updateLabel,omitempty\"` // Updates the label properties.

	UpdateSelectionChoiceProperties GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestUpdateSelectionChoicePropertiesRequest `json:"updateSelectionChoiceProperties,omitempty\"` // Update a choice property within a selection field.

}

// Request to update field properties.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestUpdateFieldPropertiesRequest struct {
	Id string `json:"id,omitempty\"` // Required. The field to update.

	Properties GoogleAppsDriveLabelsV2FieldProperties `json:"properties,omitempty\"` // Required. Basic field properties.

	UpdateMask string `json:"updateMask,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `properties` is implied and should not be specified. A single `*` can be used as a short-hand for updating every field.

}

// Request to change the type of a field.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestUpdateFieldTypeRequest struct {
	DateOptions GoogleAppsDriveLabelsV2FieldDateOptions `json:"dateOptions,omitempty\"` // Update field to Date.

	Id string `json:"id,omitempty\"` // Required. The field to update.

	IntegerOptions GoogleAppsDriveLabelsV2FieldIntegerOptions `json:"integerOptions,omitempty\"` // Update field to Integer.

	SelectionOptions GoogleAppsDriveLabelsV2FieldSelectionOptions `json:"selectionOptions,omitempty\"` // Update field to Selection.

	TextOptions GoogleAppsDriveLabelsV2FieldTextOptions `json:"textOptions,omitempty\"` // Update field to Text.

	UpdateMask string `json:"updateMask,omitempty\"` // The fields that should be updated. At least one field must be specified. The root of `type_options` is implied and should not be specified. A single `*` can be used as a short-hand for updating every field.

	UserOptions GoogleAppsDriveLabelsV2FieldUserOptions `json:"userOptions,omitempty\"` // Update field to User.

}

// Updates basic properties of a label.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestUpdateLabelPropertiesRequest struct {
	Properties GoogleAppsDriveLabelsV2LabelProperties `json:"properties,omitempty\"` // Required. Label properties to update.

	UpdateMask string `json:"updateMask,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `label_properties` is implied and should not be specified. A single `*` can be used as a short-hand for updating every field.

}

// Request to update a choice property.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelRequestUpdateSelectionChoicePropertiesRequest struct {
	FieldId string `json:"fieldId,omitempty\"` // Required. The selection field to update.

	Id string `json:"id,omitempty\"` // Required. The choice to update.

	Properties GoogleAppsDriveLabelsV2FieldSelectionOptionsChoiceProperties `json:"properties,omitempty\"` // Required. The choice properties to update.

	UpdateMask string `json:"updateMask,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `properties` is implied and should not be specified. A single `*` can be used as a short-hand for updating every field.

}

// Response for label update.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelResponse struct {
	Responses []GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseResponse `json:"responses,omitempty\"` // The reply of the updates. This maps 1:1 with the updates, although responses to some requests may be empty.

	UpdatedLabel GoogleAppsDriveLabelsV2Label `json:"updatedLabel,omitempty\"` // The label after updates were applied. This is only set if `include_label_in_response` is `true` and there were no errors.

}

// Response following field create.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseCreateFieldResponse struct {
	Id string `json:"id,omitempty\"` // The field of the created field. When left blank in a create request, a key will be autogenerated and can be identified here.

	Priority int `json:"priority,omitempty\"` // The priority of the created field. The priority may change from what was specified to assure contiguous priorities between fields (1-n).

}

// Response following selection choice create.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseCreateSelectionChoiceResponse struct {
	FieldId string `json:"fieldId,omitempty\"` // The server-generated ID of the field.

	Id string `json:"id,omitempty\"` // The server-generated ID of the created choice within the field.

}

// Response following field delete.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseDeleteFieldResponse struct {
}

// Response following choice delete.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseDeleteSelectionChoiceResponse struct {
}

// Response following field disable.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseDisableFieldResponse struct {
}

// Response following choice disable.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseDisableSelectionChoiceResponse struct {
}

// Response following field enable.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseEnableFieldResponse struct {
}

// Response following choice enable.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseEnableSelectionChoiceResponse struct {
}

// A single response from an update.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseResponse struct {
	CreateField GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseCreateFieldResponse `json:"createField,omitempty\"` // Creates a field.

	CreateSelectionChoice GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseCreateSelectionChoiceResponse `json:"createSelectionChoice,omitempty\"` // Creates a selection list option to add to a selection field.

	DeleteField GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseDeleteFieldResponse `json:"deleteField,omitempty\"` // Deletes a field from the label.

	DeleteSelectionChoice GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseDeleteSelectionChoiceResponse `json:"deleteSelectionChoice,omitempty\"` // Deletes a choice from a selection field.

	DisableField GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseDisableFieldResponse `json:"disableField,omitempty\"` // Disables field.

	DisableSelectionChoice GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseDisableSelectionChoiceResponse `json:"disableSelectionChoice,omitempty\"` // Disables a choice within a selection field.

	EnableField GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseEnableFieldResponse `json:"enableField,omitempty\"` // Enables field.

	EnableSelectionChoice GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseEnableSelectionChoiceResponse `json:"enableSelectionChoice,omitempty\"` // Enables a choice within a selection field.

	UpdateField GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseUpdateFieldPropertiesResponse `json:"updateField,omitempty\"` // Updates basic properties of a field.

	UpdateFieldType GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseUpdateFieldTypeResponse `json:"updateFieldType,omitempty\"` // Updates field type and/or type options.

	UpdateLabel GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseUpdateLabelPropertiesResponse `json:"updateLabel,omitempty\"` // Updates basic properties of a label.

	UpdateSelectionChoiceProperties GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseUpdateSelectionChoicePropertiesResponse `json:"updateSelectionChoiceProperties,omitempty\"` // Updates a choice within a selection field.

}

// Response following update to field properties.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseUpdateFieldPropertiesResponse struct {
	Priority int `json:"priority,omitempty\"` // The priority of the updated field. The priority may change from what was specified to assure contiguous priorities between fields (1-n).

}

// Response following update to field type.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseUpdateFieldTypeResponse struct {
}

// Response following update to label properties.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseUpdateLabelPropertiesResponse struct {
}

// Response following update to selection choice properties.
type GoogleAppsDriveLabelsV2DeltaUpdateLabelResponseUpdateSelectionChoicePropertiesResponse struct {
	Priority int `json:"priority,omitempty\"` // The priority of the updated choice. The priority may change from what was specified to assure contiguous priorities between choices (1-n).

}

// Request to deprecate a published label.
type GoogleAppsDriveLabelsV2DisableLabelRequest struct {
	DisabledPolicy GoogleAppsDriveLabelsV2LifecycleDisabledPolicy `json:"disabledPolicy,omitempty\"` // Disabled policy to use.

	LanguageCode string `json:"languageCode,omitempty\"` // The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default configured language will be used.

	UpdateMask string `json:"updateMask,omitempty\"` // The fields that should be updated. At least one field must be specified. The root `disabled_policy` is implied and should not be specified. A single `*` can be used as a short-hand for updating every field.

	UseAdminAccess bool `json:"useAdminAccess,omitempty\"` // Set to `true` in order to use the user's admin credentials. The server will verify the user is an admin for the label before allowing access.

	WriteControl GoogleAppsDriveLabelsV2WriteControl `json:"writeControl,omitempty\"` // Provides control over how write requests are executed. Defaults to unset, which means the last write wins.

}

// Request to enable a label.
type GoogleAppsDriveLabelsV2EnableLabelRequest struct {
	LanguageCode string `json:"languageCode,omitempty\"` // The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default configured language will be used.

	UseAdminAccess bool `json:"useAdminAccess,omitempty\"` // Set to `true` in order to use the user's admin credentials. The server will verify the user is an admin for the label before allowing access.

	WriteControl GoogleAppsDriveLabelsV2WriteControl `json:"writeControl,omitempty\"` // Provides control over how write requests are executed. Defaults to unset, which means the last write wins.

}

// Defines a field that has a display name, data type, and other configuration options. This field defines the kind of metadata that may be set on a Drive item.
type GoogleAppsDriveLabelsV2Field struct {
	AppliedCapabilities GoogleAppsDriveLabelsV2FieldAppliedCapabilities `json:"appliedCapabilities,omitempty\"` // Output only. The capabilities this user has on this field and its value when the label is applied on Drive items.

	CreateTime string `json:"createTime,omitempty\"` // Output only. The time this field was created.

	Creator GoogleAppsDriveLabelsV2UserInfo `json:"creator,omitempty\"` // Output only. The user who created this field.

	DateOptions GoogleAppsDriveLabelsV2FieldDateOptions `json:"dateOptions,omitempty\"` // Date field options.

	DisableTime string `json:"disableTime,omitempty\"` // Output only. The time this field was disabled. This value has no meaning when the field is not disabled.

	Disabler GoogleAppsDriveLabelsV2UserInfo `json:"disabler,omitempty\"` // Output only. The user who disabled this field. This value has no meaning when the field is not disabled.

	DisplayHints GoogleAppsDriveLabelsV2FieldDisplayHints `json:"displayHints,omitempty\"` // Output only. UI display hints for rendering a field.

	Id string `json:"id,omitempty\"` // Output only. The key of a field, unique within a label or library. This value is autogenerated. Matches the regex: `([a-zA-Z0-9])+`.

	IntegerOptions GoogleAppsDriveLabelsV2FieldIntegerOptions `json:"integerOptions,omitempty\"` // Integer field options.

	Lifecycle GoogleAppsDriveLabelsV2Lifecycle `json:"lifecycle,omitempty\"` // Output only. The lifecycle of this field.

	LockStatus GoogleAppsDriveLabelsV2LockStatus `json:"lockStatus,omitempty\"` // Output only. The `LockStatus` of this field.

	Properties GoogleAppsDriveLabelsV2FieldProperties `json:"properties,omitempty\"` // The basic properties of the field.

	Publisher GoogleAppsDriveLabelsV2UserInfo `json:"publisher,omitempty\"` // Output only. The user who published this field. This value has no meaning when the field is not published.

	QueryKey string `json:"queryKey,omitempty\"` // Output only. The key to use when constructing Drive search queries to find files based on values defined for this field on files. For example, "`{query_key}` > 2001-01-01".

	SchemaCapabilities GoogleAppsDriveLabelsV2FieldSchemaCapabilities `json:"schemaCapabilities,omitempty\"` // Output only. The capabilities this user has when editing this field.

	SelectionOptions GoogleAppsDriveLabelsV2FieldSelectionOptions `json:"selectionOptions,omitempty\"` // Selection field options.

	TextOptions GoogleAppsDriveLabelsV2FieldTextOptions `json:"textOptions,omitempty\"` // Text field options.

	UpdateTime string `json:"updateTime,omitempty\"` // Output only. The time this field was updated.

	Updater GoogleAppsDriveLabelsV2UserInfo `json:"updater,omitempty\"` // Output only. The user who modified this field.

	UserOptions GoogleAppsDriveLabelsV2FieldUserOptions `json:"userOptions,omitempty\"` // User field options.

}

// The capabilities related to this field on applied metadata.
type GoogleAppsDriveLabelsV2FieldAppliedCapabilities struct {
	CanRead bool `json:"canRead,omitempty\"` // Whether the user can read related applied metadata on items.

	CanSearch bool `json:"canSearch,omitempty\"` // Whether the user can search for Drive items referencing this field.

	CanWrite bool `json:"canWrite,omitempty\"` // Whether the user can set this field on Drive items.

}

// Options for the date field type.
type GoogleAppsDriveLabelsV2FieldDateOptions struct {
	DateFormat string `json:"dateFormat,omitempty\"` // Output only. ICU date format.

	DateFormatType string `json:"dateFormatType,omitempty\"` // Localized date formatting option. Field values are rendered in this format according to their locale.

	MaxValue GoogleTypeDate `json:"maxValue,omitempty\"` // Output only. Maximum valid value (year, month, day).

	MinValue GoogleTypeDate `json:"minValue,omitempty\"` // Output only. Minimum valid value (year, month, day).

}

// UI display hints for rendering a field.
type GoogleAppsDriveLabelsV2FieldDisplayHints struct {
	Disabled bool `json:"disabled,omitempty\"` // Whether the field should be shown in the UI as disabled.

	HiddenInSearch bool `json:"hiddenInSearch,omitempty\"` // This field should be hidden in the search menu when searching for Drive items.

	Required bool `json:"required,omitempty\"` // Whether the field should be shown as required in the UI.

	ShownInApply bool `json:"shownInApply,omitempty\"` // This field should be shown in the apply menu when applying values to a Drive item.

}

// Options for the Integer field type.
type GoogleAppsDriveLabelsV2FieldIntegerOptions struct {
	MaxValue int64 `json:"maxValue,omitempty\"` // Output only. The maximum valid value for the integer field.

	MinValue int64 `json:"minValue,omitempty\"` // Output only. The minimum valid value for the integer field.

}

// Field constants governing the structure of a field; such as, the maximum title length, minimum and maximum field values or length, etc.
type GoogleAppsDriveLabelsV2FieldLimits struct {
	DateLimits GoogleAppsDriveLabelsV2DateLimits `json:"dateLimits,omitempty\"` // Date field limits.

	IntegerLimits GoogleAppsDriveLabelsV2IntegerLimits `json:"integerLimits,omitempty\"` // Integer field limits.

	LongTextLimits GoogleAppsDriveLabelsV2LongTextLimits `json:"longTextLimits,omitempty\"` // Long text field limits.

	MaxDescriptionLength int `json:"maxDescriptionLength,omitempty\"` // Limits for field description, also called help text.

	MaxDisplayNameLength int `json:"maxDisplayNameLength,omitempty\"` // Limits for field title.

	MaxIdLength int `json:"maxIdLength,omitempty\"` // Maximum length for the id.

	SelectionLimits GoogleAppsDriveLabelsV2SelectionLimits `json:"selectionLimits,omitempty\"` // Selection field limits.

	TextLimits GoogleAppsDriveLabelsV2TextLimits `json:"textLimits,omitempty\"` // The relevant limits for the specified Field.Type. Text field limits.

	UserLimits GoogleAppsDriveLabelsV2UserLimits `json:"userLimits,omitempty\"` // User field limits.

}

// Options for a multi-valued variant of an associated field type.
type GoogleAppsDriveLabelsV2FieldListOptions struct {
	MaxEntries int `json:"maxEntries,omitempty\"` // Maximum number of entries permitted.

}

// The basic properties of the field.
type GoogleAppsDriveLabelsV2FieldProperties struct {
	DisplayName string `json:"displayName,omitempty\"` // Required. The display text to show in the UI identifying this field.

	InsertBeforeField string `json:"insertBeforeField,omitempty\"` // Input only. Insert or move this field before the indicated field. If empty, the field is placed at the end of the list.

	Required bool `json:"required,omitempty\"` // Whether the field should be marked as required.

}

// The capabilities related to this field when editing the field.
type GoogleAppsDriveLabelsV2FieldSchemaCapabilities struct {
	CanDelete bool `json:"canDelete,omitempty\"` // Whether the user can delete this field. The user must have permission and the field must be deprecated.

	CanDisable bool `json:"canDisable,omitempty\"` // Whether the user can disable this field. The user must have permission and this field must not already be disabled.

	CanEnable bool `json:"canEnable,omitempty\"` // Whether the user can enable this field. The user must have permission and this field must be disabled.

	CanUpdate bool `json:"canUpdate,omitempty\"` // Whether the user can change this field.

}

// Options for the selection field type.
type GoogleAppsDriveLabelsV2FieldSelectionOptions struct {
	Choices []GoogleAppsDriveLabelsV2FieldSelectionOptionsChoice `json:"choices,omitempty\"` // The options available for this selection field. The list order is consistent, and modified with `insert_before_choice`.

	ListOptions GoogleAppsDriveLabelsV2FieldListOptions `json:"listOptions,omitempty\"` // When specified, indicates this field supports a list of values. Once the field is published, this cannot be changed.

}

// Selection field choice.
type GoogleAppsDriveLabelsV2FieldSelectionOptionsChoice struct {
	AppliedCapabilities GoogleAppsDriveLabelsV2FieldSelectionOptionsChoiceAppliedCapabilities `json:"appliedCapabilities,omitempty\"` // Output only. The capabilities related to this choice on applied metadata.

	CreateTime string `json:"createTime,omitempty\"` // Output only. The time this choice was created.

	Creator GoogleAppsDriveLabelsV2UserInfo `json:"creator,omitempty\"` // Output only. The user who created this choice.

	DisableTime string `json:"disableTime,omitempty\"` // Output only. The time this choice was disabled. This value has no meaning when the choice is not disabled.

	Disabler GoogleAppsDriveLabelsV2UserInfo `json:"disabler,omitempty\"` // Output only. The user who disabled this choice. This value has no meaning when the option is not disabled.

	DisplayHints GoogleAppsDriveLabelsV2FieldSelectionOptionsChoiceDisplayHints `json:"displayHints,omitempty\"` // Output only. UI display hints for rendering a choice.

	Id string `json:"id,omitempty\"` // The unique value of the choice. This ID is autogenerated. Matches the regex: `([a-zA-Z0-9_])+`.

	Lifecycle GoogleAppsDriveLabelsV2Lifecycle `json:"lifecycle,omitempty\"` // Output only. Lifecycle of the choice.

	LockStatus GoogleAppsDriveLabelsV2LockStatus `json:"lockStatus,omitempty\"` // Output only. The `LockStatus` of this choice.

	Properties GoogleAppsDriveLabelsV2FieldSelectionOptionsChoiceProperties `json:"properties,omitempty\"` // Basic properties of the choice.

	PublishTime string `json:"publishTime,omitempty\"` // Output only. The time this choice was published. This value has no meaning when the choice is not published.

	Publisher GoogleAppsDriveLabelsV2UserInfo `json:"publisher,omitempty\"` // Output only. The user who published this choice. This value has no meaning when the choice is not published.

	SchemaCapabilities GoogleAppsDriveLabelsV2FieldSelectionOptionsChoiceSchemaCapabilities `json:"schemaCapabilities,omitempty\"` // Output only. The capabilities related to this option when editing the option.

	UpdateTime string `json:"updateTime,omitempty\"` // Output only. The time this choice was updated last.

	Updater GoogleAppsDriveLabelsV2UserInfo `json:"updater,omitempty\"` // Output only. The user who updated this choice last.

}

// The capabilities related to this choice on applied metadata.
type GoogleAppsDriveLabelsV2FieldSelectionOptionsChoiceAppliedCapabilities struct {
	CanRead bool `json:"canRead,omitempty\"` // Whether the user can read related applied metadata on items.

	CanSearch bool `json:"canSearch,omitempty\"` // Whether the user can use this choice in search queries.

	CanSelect bool `json:"canSelect,omitempty\"` // Whether the user can select this choice on an item.

}

// UI display hints for rendering an option.
type GoogleAppsDriveLabelsV2FieldSelectionOptionsChoiceDisplayHints struct {
	BadgeColors GoogleAppsDriveLabelsV2BadgeColors `json:"badgeColors,omitempty\"` // The colors to use for the badge. Changed to Google Material colors based on the chosen `properties.badge_config.color`.

	BadgePriority int64 `json:"badgePriority,omitempty\"` // The priority of this badge. Used to compare and sort between multiple badges. A lower number means the badge should be shown first. When a badging configuration is not present, this will be 0. Otherwise, this will be set to `BadgeConfig.priority_override` or the default heuristic which prefers creation date of the label, and field and option priority.

	DarkBadgeColors GoogleAppsDriveLabelsV2BadgeColors `json:"darkBadgeColors,omitempty\"` // The dark-mode color to use for the badge. Changed to Google Material colors based on the chosen `properties.badge_config.color`.

	Disabled bool `json:"disabled,omitempty\"` // Whether the option should be shown in the UI as disabled.

	HiddenInSearch bool `json:"hiddenInSearch,omitempty\"` // This option should be hidden in the search menu when searching for Drive items.

	ShownInApply bool `json:"shownInApply,omitempty\"` // This option should be shown in the apply menu when applying values to a Drive item.

}

// Basic properties of the choice.
type GoogleAppsDriveLabelsV2FieldSelectionOptionsChoiceProperties struct {
	BadgeConfig GoogleAppsDriveLabelsV2BadgeConfig `json:"badgeConfig,omitempty\"` // The badge configuration for this choice. When set, the label that owns this choice is considered a "badged label".

	Description string `json:"description,omitempty\"` // The description of this label.

	DisplayName string `json:"displayName,omitempty\"` // Required. The display text to show in the UI identifying this field.

	InsertBeforeChoice string `json:"insertBeforeChoice,omitempty\"` // Input only. Insert or move this choice before the indicated choice. If empty, the choice is placed at the end of the list.

}

// The capabilities related to this choice when editing the choice.
type GoogleAppsDriveLabelsV2FieldSelectionOptionsChoiceSchemaCapabilities struct {
	CanDelete bool `json:"canDelete,omitempty\"` // Whether the user can delete this choice.

	CanDisable bool `json:"canDisable,omitempty\"` // Whether the user can disable this choice.

	CanEnable bool `json:"canEnable,omitempty\"` // Whether the user can enable this choice.

	CanUpdate bool `json:"canUpdate,omitempty\"` // Whether the user can update this choice.

}

// Options for the Text field type.
type GoogleAppsDriveLabelsV2FieldTextOptions struct {
	MaxLength int `json:"maxLength,omitempty\"` // Output only. The maximum valid length of values for the text field.

	MinLength int `json:"minLength,omitempty\"` // Output only. The minimum valid length of values for the text field.

}

// Options for the user field type.
type GoogleAppsDriveLabelsV2FieldUserOptions struct {
	ListOptions GoogleAppsDriveLabelsV2FieldListOptions `json:"listOptions,omitempty\"` // When specified, indicates that this field supports a list of values. Once the field is published, this cannot be changed.

}

// Limits for integer field type.
type GoogleAppsDriveLabelsV2IntegerLimits struct {
	MaxValue int64 `json:"maxValue,omitempty\"` // Maximum value for an integer field type.

	MinValue int64 `json:"minValue,omitempty\"` // Minimum value for an integer field type.

}

// A label defines a taxonomy that can be applied to Drive items in order to organize and search across items. Labels can be simple strings, or can contain fields that describe additional metadata that can be further used to organize and search Drive items.
type GoogleAppsDriveLabelsV2Label struct {
	AppliedCapabilities GoogleAppsDriveLabelsV2LabelAppliedCapabilities `json:"appliedCapabilities,omitempty\"` // Output only. The capabilities related to this label on applied metadata.

	AppliedLabelPolicy GoogleAppsDriveLabelsV2LabelAppliedLabelPolicy `json:"appliedLabelPolicy,omitempty\"` // Output only. Behavior of this label when it's applied to Drive items.

	CreateTime string `json:"createTime,omitempty\"` // Output only. The time this label was created.

	Creator GoogleAppsDriveLabelsV2UserInfo `json:"creator,omitempty\"` // Output only. The user who created this label.

	Customer string `json:"customer,omitempty\"` // Output only. The customer this label belongs to. For example: `customers/123abc789`.

	DisableTime string `json:"disableTime,omitempty\"` // Output only. The time this label was disabled. This value has no meaning when the label isn't disabled.

	Disabler GoogleAppsDriveLabelsV2UserInfo `json:"disabler,omitempty\"` // Output only. The user who disabled this label. This value has no meaning when the label isn't disabled.

	DisplayHints GoogleAppsDriveLabelsV2LabelDisplayHints `json:"displayHints,omitempty\"` // Output only. UI display hints for rendering the label.

	EnabledAppSettings GoogleAppsDriveLabelsV2LabelEnabledAppSettings `json:"enabledAppSettings,omitempty\"` // Optional. The `EnabledAppSettings` for this Label.

	Fields []GoogleAppsDriveLabelsV2Field `json:"fields,omitempty\"` // List of fields in descending priority order.

	Id string `json:"id,omitempty\"` // Output only. Globally unique identifier of this label. ID makes up part of the label `name`, but unlike `name`, ID is consistent between revisions. Matches the regex: `([a-zA-Z0-9])+`.

	LabelType string `json:"labelType,omitempty\"` // Required. The type of label.

	LearnMoreUri string `json:"learnMoreUri,omitempty\"` // Custom URL to present to users to allow them to learn more about this label and how it should be used.

	Lifecycle GoogleAppsDriveLabelsV2Lifecycle `json:"lifecycle,omitempty\"` // Output only. The lifecycle state of the label including whether it's published, deprecated, and has draft changes.

	LockStatus GoogleAppsDriveLabelsV2LockStatus `json:"lockStatus,omitempty\"` // Output only. The `LockStatus` of this label.

	Name string `json:"name,omitempty\"` // Output only. Resource name of the label. Will be in the form of either: `labels/{id}` or `labels/{id}@{revision_id}` depending on the request. See `id` and `revision_id` below.

	Properties GoogleAppsDriveLabelsV2LabelProperties `json:"properties,omitempty\"` // Required. The basic properties of the label.

	PublishTime string `json:"publishTime,omitempty\"` // Output only. The time this label was published. This value has no meaning when the label isn't published.

	Publisher GoogleAppsDriveLabelsV2UserInfo `json:"publisher,omitempty\"` // Output only. The user who published this label. This value has no meaning when the label isn't published.>>

	RevisionCreateTime string `json:"revisionCreateTime,omitempty\"` // Output only. The time this label revision was created.

	RevisionCreator GoogleAppsDriveLabelsV2UserInfo `json:"revisionCreator,omitempty\"` // Output only. The user who created this label revision.

	RevisionId string `json:"revisionId,omitempty\"` // Output only. Revision ID of the label. Revision ID might be part of the label `name` depending on the request issued. A new revision is created whenever revisioned properties of a label are changed. Matches the regex: `([a-zA-Z0-9])+`.

	SchemaCapabilities GoogleAppsDriveLabelsV2LabelSchemaCapabilities `json:"schemaCapabilities,omitempty\"` // Output only. The capabilities the user has on this label.

}

// The capabilities a user has on this label's applied metadata.
type GoogleAppsDriveLabelsV2LabelAppliedCapabilities struct {
	CanApply bool `json:"canApply,omitempty\"` // Whether the user can apply this label to items.

	CanRead bool `json:"canRead,omitempty\"` // Whether the user can read applied metadata related to this label.

	CanRemove bool `json:"canRemove,omitempty\"` // Whether the user can remove this label from items.

}

// Behavior of this label when it's applied to Drive items.
type GoogleAppsDriveLabelsV2LabelAppliedLabelPolicy struct {
	CopyMode string `json:"copyMode,omitempty\"` // Indicates how the applied label and field values should be copied when a Drive item is copied.

}

// The UI display hints for rendering the label.
type GoogleAppsDriveLabelsV2LabelDisplayHints struct {
	Disabled bool `json:"disabled,omitempty\"` // Whether the label should be shown in the UI as disabled.

	HiddenInSearch bool `json:"hiddenInSearch,omitempty\"` // This label should be hidden in the search menu when searching for Drive items.

	Priority int64 `json:"priority,omitempty\"` // The order to display labels in a list.

	ShownInApply bool `json:"shownInApply,omitempty\"` // This label should be shown in the apply menu when applying values to a Drive item.

}

// Describes the Google Workspace apps in which the label can be used.
type GoogleAppsDriveLabelsV2LabelEnabledAppSettings struct {
	EnabledApps []GoogleAppsDriveLabelsV2LabelEnabledAppSettingsEnabledApp `json:"enabledApps,omitempty\"` // Optional. The list of apps where the label can be used.

}

// An app where the label can be used.
type GoogleAppsDriveLabelsV2LabelEnabledAppSettingsEnabledApp struct {
	App string `json:"app,omitempty\"` // Optional. The name of the app.

}

// Label constraints governing the structure of a label; such as, the maximum number of fields allowed and maximum length of the label title.
type GoogleAppsDriveLabelsV2LabelLimits struct {
	FieldLimits GoogleAppsDriveLabelsV2FieldLimits `json:"fieldLimits,omitempty\"` // The limits for fields.

	MaxDeletedFields int `json:"maxDeletedFields,omitempty\"` // The maximum number of published fields that can be deleted.

	MaxDescriptionLength int `json:"maxDescriptionLength,omitempty\"` // The maximum number of characters allowed for the description.

	MaxDraftRevisions int `json:"maxDraftRevisions,omitempty\"` // The maximum number of draft revisions that will be kept before deleting old drafts.

	MaxFields int `json:"maxFields,omitempty\"` // The maximum number of fields allowed within the label.

	MaxTitleLength int `json:"maxTitleLength,omitempty\"` // The maximum number of characters allowed for the title.

	Name string `json:"name,omitempty\"` // Resource name.

}

// A lock that can be applied to a label, field, or choice.
type GoogleAppsDriveLabelsV2LabelLock struct {
	Capabilities GoogleAppsDriveLabelsV2LabelLockCapabilities `json:"capabilities,omitempty\"` // Output only. The user's capabilities on this label lock.

	ChoiceId string `json:"choiceId,omitempty\"` // The ID of the selection field choice that should be locked. If present, `field_id` must also be present.

	CreateTime string `json:"createTime,omitempty\"` // Output only. The time this label lock was created.

	Creator GoogleAppsDriveLabelsV2UserInfo `json:"creator,omitempty\"` // Output only. The user whose credentials were used to create the label lock. Not present if no user was responsible for creating the label lock.

	DeleteTime string `json:"deleteTime,omitempty\"` // Output only. A timestamp indicating when this label lock was scheduled for deletion. Present only if this label lock is in the `DELETING` state.

	FieldId string `json:"fieldId,omitempty\"` // The ID of the field that should be locked. Empty if the whole label should be locked.

	Name string `json:"name,omitempty\"` // Output only. Resource name of this label lock.

	State string `json:"state,omitempty\"` // Output only. This label lock's state.

}

// A description of a user's capabilities on a label lock.
type GoogleAppsDriveLabelsV2LabelLockCapabilities struct {
	CanViewPolicy bool `json:"canViewPolicy,omitempty\"` // True if the user is authorized to view the policy.

}

// The permission that applies to a principal (user, group, audience) on a label.
type GoogleAppsDriveLabelsV2LabelPermission struct {
	Audience string `json:"audience,omitempty\"` // Audience to grant a role to. The magic value of `audiences/default` may be used to apply the role to the default audience in the context of the organization that owns the label.

	Email string `json:"email,omitempty\"` // Specifies the email address for a user or group principal. Not populated for audience principals. User and group permissions may only be inserted using an email address. On update requests, if email address is specified, no principal should be specified.

	Group string `json:"group,omitempty\"` // Group resource name.

	Name string `json:"name,omitempty\"` // Resource name of this permission.

	Person string `json:"person,omitempty\"` // Person resource name.

	Role string `json:"role,omitempty\"` // The role the principal should have.

}

// Basic properties of the label.
type GoogleAppsDriveLabelsV2LabelProperties struct {
	Description string `json:"description,omitempty\"` // The description of the label.

	Title string `json:"title,omitempty\"` // Required. Title of the label.

}

// The capabilities related to this label when editing the label.
type GoogleAppsDriveLabelsV2LabelSchemaCapabilities struct {
	CanDelete bool `json:"canDelete,omitempty\"` // Whether the user can delete this label. The user must have permission and the label must be disabled.

	CanDisable bool `json:"canDisable,omitempty\"` // Whether the user can disable this label. The user must have permission and this label must not already be disabled.

	CanEnable bool `json:"canEnable,omitempty\"` // Whether the user can enable this label. The user must have permission and this label must be disabled.

	CanUpdate bool `json:"canUpdate,omitempty\"` // Whether the user can change this label.

}

// The lifecycle state of an object, such as label, field, or choice. For more information, see [Label lifecycle](https://developers.google.com/workspace/drive/labels/guides/label-lifecycle). The lifecycle enforces the following transitions: * `UNPUBLISHED_DRAFT` (starting state) * `UNPUBLISHED_DRAFT` -> `PUBLISHED` * `UNPUBLISHED_DRAFT` -> (Deleted) * `PUBLISHED` -> `DISABLED` * `DISABLED` -> `PUBLISHED` * `DISABLED` -> (Deleted) The published and disabled states have some distinct characteristics: * `Published`: Some kinds of changes might be made to an object in this state, in which case `has_unpublished_changes` will be true. Also, some kinds of changes aren't permitted. Generally, any change that would invalidate or cause new restrictions on existing metadata related to the label are rejected. * `Disabled`: When disabled, the configured `DisabledPolicy` takes effect.
type GoogleAppsDriveLabelsV2Lifecycle struct {
	DisabledPolicy GoogleAppsDriveLabelsV2LifecycleDisabledPolicy `json:"disabledPolicy,omitempty\"` // The policy that governs how to show a disabled label, field, or selection choice.

	HasUnpublishedChanges bool `json:"hasUnpublishedChanges,omitempty\"` // Output only. Whether the object associated with this lifecycle has unpublished changes.

	State string `json:"state,omitempty\"` // Output only. The state of the object associated with this lifecycle.

}

// The policy that governs how to treat a disabled label, field, or selection choice in different contexts.
type GoogleAppsDriveLabelsV2LifecycleDisabledPolicy struct {
	HideInSearch bool `json:"hideInSearch,omitempty\"` // Whether to hide this disabled object in the search menu for Drive items. * When `false`, the object is generally shown in the UI as disabled but it appears in the search results when searching for Drive items. * When `true`, the object is generally hidden in the UI when searching for Drive items.

	ShowInApply bool `json:"showInApply,omitempty\"` // Whether to show this disabled object in the apply menu on Drive items. * When `true`, the object is generally shown in the UI as disabled and is unselectable. * When `false`, the object is generally hidden in the UI.

}

// The response to a `ListLabelLocksRequest`.
type GoogleAppsDriveLabelsV2ListLabelLocksResponse struct {
	LabelLocks []GoogleAppsDriveLabelsV2LabelLock `json:"labelLocks,omitempty\"` // Label locks.

	NextPageToken string `json:"nextPageToken,omitempty\"` // The token of the next page in the response.

}

// Response for listing the permissions on a label.
type GoogleAppsDriveLabelsV2ListLabelPermissionsResponse struct {
	LabelPermissions []GoogleAppsDriveLabelsV2LabelPermission `json:"labelPermissions,omitempty\"` // Label permissions.

	NextPageToken string `json:"nextPageToken,omitempty\"` // The token of the next page in the response.

}

// Response for listing labels.
type GoogleAppsDriveLabelsV2ListLabelsResponse struct {
	Labels []GoogleAppsDriveLabelsV2Label `json:"labels,omitempty\"` // Labels.

	NextPageToken string `json:"nextPageToken,omitempty\"` // The token of the next page in the response.

}

// Limits for list-variant of a field type.
type GoogleAppsDriveLabelsV2ListLimits struct {
	MaxEntries int `json:"maxEntries,omitempty\"` // Maximum number of values allowed for the field type.

}

// Contains information about whether a label component should be considered locked.
type GoogleAppsDriveLabelsV2LockStatus struct {
	Locked bool `json:"locked,omitempty\"` // Output only. Indicates whether this label component is the (direct) target of a label lock. A label component can be implicitly locked even if it's not the direct target of a label lock, in which case this field is set to false.

}

// Limits for long text field type.
type GoogleAppsDriveLabelsV2LongTextLimits struct {
	MaxLength int `json:"maxLength,omitempty\"` // Maximum length allowed for a long text field type.

	MinLength int `json:"minLength,omitempty\"` // Minimum length allowed for a long text field type.

}

// Request to publish a label.
type GoogleAppsDriveLabelsV2PublishLabelRequest struct {
	LanguageCode string `json:"languageCode,omitempty\"` // The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default configured language will be used.

	UseAdminAccess bool `json:"useAdminAccess,omitempty\"` // Set to `true` in order to use the user's admin credentials. The server will verify the user is an admin for the label before allowing access.

	WriteControl GoogleAppsDriveLabelsV2WriteControl `json:"writeControl,omitempty\"` // Provides control over how write requests are executed. Defaults to unset, which means the last write wins.

}

// Limits for selection field type.
type GoogleAppsDriveLabelsV2SelectionLimits struct {
	ListLimits GoogleAppsDriveLabelsV2ListLimits `json:"listLimits,omitempty\"` // Limits for list-variant of a field type.

	MaxChoices int `json:"maxChoices,omitempty\"` // Maximum number of choices.

	MaxDeletedChoices int `json:"maxDeletedChoices,omitempty\"` // Maximum number of deleted choices.

	MaxDisplayNameLength int `json:"maxDisplayNameLength,omitempty\"` // Maximum length for display name.

	MaxIdLength int `json:"maxIdLength,omitempty\"` // Maximum ID length for a selection option.

}

// Limits for text field type.
type GoogleAppsDriveLabelsV2TextLimits struct {
	MaxLength int `json:"maxLength,omitempty\"` // Maximum length allowed for a text field type.

	MinLength int `json:"minLength,omitempty\"` // Minimum length allowed for a text field type.

}

// Request to update the `CopyMode` of the given label. Changes to this policy aren't revisioned, don't require publishing, and take effect immediately. \
type GoogleAppsDriveLabelsV2UpdateLabelCopyModeRequest struct {
	CopyMode string `json:"copyMode,omitempty\"` // Required. Indicates how the applied label and field values should be copied when a Drive item is copied.

	LanguageCode string `json:"languageCode,omitempty\"` // The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default configured language will be used.

	UseAdminAccess bool `json:"useAdminAccess,omitempty\"` // Set to `true` in order to use the user's admin credentials. The server will verify the user is an admin for the label before allowing access.

	View string `json:"view,omitempty\"` // When specified, only certain fields belonging to the indicated view will be returned.

}

// Request to update the `EnabledAppSettings` of the given label. This change is not revisioned, doesn't require publishing, and takes effect immediately. \
type GoogleAppsDriveLabelsV2UpdateLabelEnabledAppSettingsRequest struct {
	EnabledAppSettings GoogleAppsDriveLabelsV2LabelEnabledAppSettings `json:"enabledAppSettings,omitempty\"` // Required. The new `EnabledAppSettings` value for the label.

	LanguageCode string `json:"languageCode,omitempty\"` // Optional. The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default configured language will be used.

	UseAdminAccess bool `json:"useAdminAccess,omitempty\"` // Optional. Set to `true` in order to use the user's admin credentials. The server will verify the user is an admin for the label before allowing access.

	View string `json:"view,omitempty\"` // Optional. When specified, only certain fields belonging to the indicated view will be returned.

}

// Updates a label permission. Permissions affect the label resource as a whole, aren't revisioned, and don't require publishing.
type GoogleAppsDriveLabelsV2UpdateLabelPermissionRequest struct {
	LabelPermission GoogleAppsDriveLabelsV2LabelPermission `json:"labelPermission,omitempty\"` // Required. The permission to create or update on the label.

	Parent string `json:"parent,omitempty\"` // Required. The parent label resource name.

	UseAdminAccess bool `json:"useAdminAccess,omitempty\"` // Set to `true` in order to use the user's admin credentials. The server will verify the user is an admin for the label before allowing access.

}

// The capabilities of a user.
type GoogleAppsDriveLabelsV2UserCapabilities struct {
	CanAccessLabelManager bool `json:"canAccessLabelManager,omitempty\"` // Output only. Whether the user is allowed access to the label manager.

	CanAdministrateLabels bool `json:"canAdministrateLabels,omitempty\"` // Output only. Whether the user is an administrator for the shared labels feature.

	CanCreateAdminLabels bool `json:"canCreateAdminLabels,omitempty\"` // Output only. Whether the user is allowed to create admin labels.

	CanCreateSharedLabels bool `json:"canCreateSharedLabels,omitempty\"` // Output only. Whether the user is allowed to create shared labels.

	Name string `json:"name,omitempty\"` // Output only. Resource name for the user capabilities.

}

// Information about a user.
type GoogleAppsDriveLabelsV2UserInfo struct {
	Person string `json:"person,omitempty\"` // The identifier for this user that can be used with the [People API](https://developers.google.com/people) to get more information. For example, `people/12345678`.

}

// Limits for Field.Type.USER.
type GoogleAppsDriveLabelsV2UserLimits struct {
	ListLimits GoogleAppsDriveLabelsV2ListLimits `json:"listLimits,omitempty\"` // Limits for list-variant of a field type.

}

// Provides control over how write requests are executed. When not specified, the last write wins.
type GoogleAppsDriveLabelsV2WriteControl struct {
	RequiredRevisionId string `json:"requiredRevisionId,omitempty\"` // The revision ID of the label that the write request will be applied to. If this isn't the latest revision of the label, the request will not be processed and will return a 400 Bad Request error.

}

// A generic empty message that you can re-use to avoid defining duplicated empty messages in your APIs. A typical example is to use it as the request or the response type of an API method. For instance: service Foo { rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty); }
type GoogleProtobufEmpty struct {
}

// Represents a color in the RGBA color space. This representation is designed for simplicity of conversion to and from color representations in various languages over compactness. For example, the fields of this representation can be trivially provided to the constructor of `java.awt.Color` in Java; it can also be trivially provided to UIColor's `+colorWithRed:green:blue:alpha` method in iOS; and, with just a little work, it can be easily formatted into a CSS `rgba()` string in JavaScript. This reference page doesn't have information about the absolute color space that should be used to interpret the RGB value—for example, sRGB, Adobe RGB, DCI-P3, and BT.2020. By default, applications should assume the sRGB color space. When color equality needs to be decided, implementations, unless documented otherwise, treat two colors as equal if all their red, green, blue, and alpha values each differ by at most `1e-5`. Example (Java): import com.google.type.Color; // ... public static java.awt.Color fromProto(Color protocolor) { float alpha = protocolor.hasAlpha() ? protocolor.getAlpha().getValue() : 1.0; return new java.awt.Color( protocolor.getRed(), protocolor.getGreen(), protocolor.getBlue(), alpha); } public static Color toProto(java.awt.Color color) { float red = (float) color.getRed(); float green = (float) color.getGreen(); float blue = (float) color.getBlue(); float denominator = 255.0; Color.Builder resultBuilder = Color .newBuilder() .setRed(red / denominator) .setGreen(green / denominator) .setBlue(blue / denominator); int alpha = color.getAlpha(); if (alpha != 255) { result.setAlpha( FloatValue .newBuilder() .setValue(((float) alpha) / denominator) .build()); } return resultBuilder.build(); } // ... Example (iOS / Obj-C): // ... static UIColor* fromProto(Color* protocolor) { float red = [protocolor red]; float green = [protocolor green]; float blue = [protocolor blue]; FloatValue* alpha_wrapper = [protocolor alpha]; float alpha = 1.0; if (alpha_wrapper != nil) { alpha = [alpha_wrapper value]; } return [UIColor colorWithRed:red green:green blue:blue alpha:alpha]; } static Color* toProto(UIColor* color) { CGFloat red, green, blue, alpha; if (![color getRed:&red green:&green blue:&blue alpha:&alpha]) { return nil; } Color* result = [[Color alloc] init]; [result setRed:red]; [result setGreen:green]; [result setBlue:blue]; if (alpha <= 0.9999) { [result setAlpha:floatWrapperWithValue(alpha)]; } [result autorelease]; return result; } // ... Example (JavaScript): // ... var protoToCssColor = function(rgb_color) { var redFrac = rgb_color.red || 0.0; var greenFrac = rgb_color.green || 0.0; var blueFrac = rgb_color.blue || 0.0; var red = Math.floor(redFrac * 255); var green = Math.floor(greenFrac * 255); var blue = Math.floor(blueFrac * 255); if (!('alpha' in rgb_color)) { return rgbToCssColor(red, green, blue); } var alphaFrac = rgb_color.alpha.value || 0.0; var rgbParams = [red, green, blue].join(','); return ['rgba(', rgbParams, ',', alphaFrac, ')'].join(”); }; var rgbToCssColor = function(red, green, blue) { var rgbNumber = new Number((red << 16) | (green << 8) | blue); var hexString = rgbNumber.toString(16); var missingZeros = 6 - hexString.length; var resultBuilder = ['#']; for (var i = 0; i < missingZeros; i++) { resultBuilder.push('0'); } resultBuilder.push(hexString); return resultBuilder.join(”); }; // ...
type GoogleTypeColor struct {
	Alpha float64 `json:"alpha,omitempty\"` // The fraction of this color that should be applied to the pixel. That is, the final pixel color is defined by the equation: `pixel color = alpha * (this color) + (1.0 - alpha) * (background color)` This means that a value of 1.0 corresponds to a solid color, whereas a value of 0.0 corresponds to a completely transparent color. This uses a wrapper message rather than a simple float scalar so that it is possible to distinguish between a default value and the value being unset. If omitted, this color object is rendered as a solid color (as if the alpha value had been explicitly given a value of 1.0).

	Blue float64 `json:"blue,omitempty\"` // The amount of blue in the color as a value in the interval [0, 1].

	Green float64 `json:"green,omitempty\"` // The amount of green in the color as a value in the interval [0, 1].

	Red float64 `json:"red,omitempty\"` // The amount of red in the color as a value in the interval [0, 1].

}

// Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following: * A full date, with non-zero year, month, and day values. * A month and day, with a zero year (for example, an anniversary). * A year on its own, with a zero month and a zero day. * A year and month, with a zero day (for example, a credit card expiration date). Related types: * google.type.TimeOfDay * google.type.DateTime * google.protobuf.Timestamp
type GoogleTypeDate struct {
	Day int `json:"day,omitempty\"` // Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.

	Month int `json:"month,omitempty\"` // Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.

	Year int `json:"year,omitempty\"` // Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.

}
