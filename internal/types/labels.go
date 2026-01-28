package types

// Label represents a Drive label schema
type Label struct {
	// ID is the unique identifier for the label
	ID string `json:"id"`

	// Name is the display name of the label
	Name string `json:"name"`

	// RevisionID is the revision identifier for the label
	RevisionID string `json:"revisionId,omitempty"`

	// LabelType is the type of label (SHARED or ADMIN)
	LabelType string `json:"labelType,omitempty"`

	// Creator is the user who created the label
	Creator *LabelUser `json:"creator,omitempty"`

	// RevisionCreator is the user who created this revision
	RevisionCreator *LabelUser `json:"revisionCreator,omitempty"`

	// CreateTime is when the label was created
	CreateTime string `json:"createTime,omitempty"`

	// RevisionCreateTime is when this revision was created
	RevisionCreateTime string `json:"revisionCreateTime,omitempty"`

	// Publisher is the user who published the label
	Publisher *LabelUser `json:"publisher,omitempty"`

	// PublishTime is when the label was published
	PublishTime string `json:"publishTime,omitempty"`

	// DisableTime is when the label was disabled
	DisableTime string `json:"disableTime,omitempty"`

	// Disabler is the user who disabled the label
	Disabler *LabelUser `json:"disabler,omitempty"`

	// Customer is the customer this label belongs to
	Customer string `json:"customer,omitempty"`

	// Properties contains label properties
	Properties *LabelProperties `json:"properties,omitempty"`

	// Lifecycle describes the lifecycle state of the label
	Lifecycle *LabelLifecycle `json:"lifecycle,omitempty"`

	// Fields are the fields defined in the label
	Fields []*LabelField `json:"fields,omitempty"`

	// AppliedCapabilities describes what can be done with this label
	AppliedCapabilities *LabelAppliedCapabilities `json:"appliedCapabilities,omitempty"`

	// SchemaCapabilities describes what can be done with the label schema
	SchemaCapabilities *LabelSchemaCapabilities `json:"schemaCapabilities,omitempty"`

	// AppliedLabelPolicy describes the policy for applying this label
	AppliedLabelPolicy *LabelAppliedLabelPolicy `json:"appliedLabelPolicy,omitempty"`
}

// LabelUser represents a user in the label context
type LabelUser struct {
	// Person is the identifier for the user
	Person string `json:"person,omitempty"`

	// Email is the email address of the user
	Email string `json:"email,omitempty"`

	// DisplayName is the display name of the user
	DisplayName string `json:"displayName,omitempty"`
}

// LabelProperties contains label properties
type LabelProperties struct {
	// Title is the title of the label
	Title string `json:"title,omitempty"`

	// Description is the description of the label
	Description string `json:"description,omitempty"`
}

// LabelLifecycle describes the lifecycle state of a label
type LabelLifecycle struct {
	// State is the lifecycle state (UNPUBLISHED_DRAFT, PUBLISHED, DISABLED, DELETED)
	State string `json:"state,omitempty"`

	// HasUnpublishedChanges indicates if there are unpublished changes
	HasUnpublishedChanges bool `json:"hasUnpublishedChanges,omitempty"`

	// DisabledPolicy describes why the label is disabled
	DisabledPolicy *LabelDisabledPolicy `json:"disabledPolicy,omitempty"`
}

// LabelDisabledPolicy describes why a label is disabled
type LabelDisabledPolicy struct {
	// HideInSearch indicates if the label should be hidden in search
	HideInSearch bool `json:"hideInSearch,omitempty"`

	// ShowInApply indicates if the label should be shown in apply
	ShowInApply bool `json:"showInApply,omitempty"`
}

// LabelAppliedCapabilities describes what can be done with an applied label
type LabelAppliedCapabilities struct {
	// CanRead indicates if the label can be read
	CanRead bool `json:"canRead,omitempty"`

	// CanApply indicates if the label can be applied
	CanApply bool `json:"canApply,omitempty"`

	// CanRemove indicates if the label can be removed
	CanRemove bool `json:"canRemove,omitempty"`
}

// LabelSchemaCapabilities describes what can be done with a label schema
type LabelSchemaCapabilities struct {
	// CanUpdate indicates if the schema can be updated
	CanUpdate bool `json:"canUpdate,omitempty"`

	// CanDelete indicates if the schema can be deleted
	CanDelete bool `json:"canDelete,omitempty"`

	// CanDisable indicates if the schema can be disabled
	CanDisable bool `json:"canDisable,omitempty"`

	// CanEnable indicates if the schema can be enabled
	CanEnable bool `json:"canEnable,omitempty"`
}

// LabelAppliedLabelPolicy describes the policy for applying a label
type LabelAppliedLabelPolicy struct {
	// CopyMode describes how the label is copied (DO_NOT_COPY, ALWAYS_COPY, COPY_APPLIABLE)
	CopyMode string `json:"copyMode,omitempty"`
}

// LabelField represents a field in a label
type LabelField struct {
	// ID is the unique identifier for the field
	ID string `json:"id"`

	// QueryKey is the key to use in queries
	QueryKey string `json:"queryKey,omitempty"`

	// Properties contains field properties
	Properties *LabelFieldProperties `json:"properties,omitempty"`

	// Lifecycle describes the lifecycle state of the field
	Lifecycle *LabelFieldLifecycle `json:"lifecycle,omitempty"`

	// DisplayHints provides display hints for the field
	DisplayHints *LabelFieldDisplayHints `json:"displayHints,omitempty"`

	// SchemaCapabilities describes what can be done with the field schema
	SchemaCapabilities *LabelFieldSchemaCapabilities `json:"schemaCapabilities,omitempty"`

	// AppliedCapabilities describes what can be done with the applied field
	AppliedCapabilities *LabelFieldAppliedCapabilities `json:"appliedCapabilities,omitempty"`

	// TextOptions for text fields
	TextOptions *LabelFieldTextOptions `json:"textOptions,omitempty"`

	// IntegerOptions for integer fields
	IntegerOptions *LabelFieldIntegerOptions `json:"integerOptions,omitempty"`

	// DateOptions for date fields
	DateOptions *LabelFieldDateOptions `json:"dateOptions,omitempty"`

	// SelectionOptions for selection fields
	SelectionOptions *LabelFieldSelectionOptions `json:"selectionOptions,omitempty"`

	// UserOptions for user fields
	UserOptions *LabelFieldUserOptions `json:"userOptions,omitempty"`
}

// LabelFieldProperties contains field properties
type LabelFieldProperties struct {
	// DisplayName is the display name of the field
	DisplayName string `json:"displayName,omitempty"`

	// Required indicates if the field is required
	Required bool `json:"required,omitempty"`

	// InsertBeforeField is the field to insert before
	InsertBeforeField string `json:"insertBeforeField,omitempty"`
}

// LabelFieldLifecycle describes the lifecycle state of a field
type LabelFieldLifecycle struct {
	// State is the lifecycle state (UNPUBLISHED_DRAFT, PUBLISHED, DISABLED, DELETED)
	State string `json:"state,omitempty"`

	// DisabledPolicy describes why the field is disabled
	DisabledPolicy *LabelFieldDisabledPolicy `json:"disabledPolicy,omitempty"`
}

// LabelFieldDisabledPolicy describes why a field is disabled
type LabelFieldDisabledPolicy struct {
	// HideInSearch indicates if the field should be hidden in search
	HideInSearch bool `json:"hideInSearch,omitempty"`

	// ShowInApply indicates if the field should be shown in apply
	ShowInApply bool `json:"showInApply,omitempty"`
}

// LabelFieldDisplayHints provides display hints for a field
type LabelFieldDisplayHints struct {
	// Required indicates if the field is required
	Required bool `json:"required,omitempty"`

	// Disabled indicates if the field is disabled
	Disabled bool `json:"disabled,omitempty"`

	// HiddenInSearch indicates if the field is hidden in search
	HiddenInSearch bool `json:"hiddenInSearch,omitempty"`

	// ShownInApply indicates if the field is shown in apply
	ShownInApply bool `json:"shownInApply,omitempty"`
}

// LabelFieldSchemaCapabilities describes what can be done with a field schema
type LabelFieldSchemaCapabilities struct {
	// CanUpdate indicates if the field can be updated
	CanUpdate bool `json:"canUpdate,omitempty"`

	// CanDelete indicates if the field can be deleted
	CanDelete bool `json:"canDelete,omitempty"`

	// CanDisable indicates if the field can be disabled
	CanDisable bool `json:"canDisable,omitempty"`

	// CanEnable indicates if the field can be enabled
	CanEnable bool `json:"canEnable,omitempty"`
}

// LabelFieldAppliedCapabilities describes what can be done with an applied field
type LabelFieldAppliedCapabilities struct {
	// CanRead indicates if the field can be read
	CanRead bool `json:"canRead,omitempty"`

	// CanSearch indicates if the field can be searched
	CanSearch bool `json:"canSearch,omitempty"`

	// CanWrite indicates if the field can be written
	CanWrite bool `json:"canWrite,omitempty"`
}

// LabelFieldTextOptions for text fields
type LabelFieldTextOptions struct {
	// MinLength is the minimum length
	MinLength int `json:"minLength,omitempty"`

	// MaxLength is the maximum length
	MaxLength int `json:"maxLength,omitempty"`
}

// LabelFieldIntegerOptions for integer fields
type LabelFieldIntegerOptions struct {
	// MinValue is the minimum value
	MinValue int64 `json:"minValue,omitempty"`

	// MaxValue is the maximum value
	MaxValue int64 `json:"maxValue,omitempty"`
}

// LabelFieldDateOptions for date fields
type LabelFieldDateOptions struct {
	// DateFormatType is the date format type (LONG_DATE, SHORT_DATE)
	DateFormatType string `json:"dateFormatType,omitempty"`

	// DateFormat is the date format string
	DateFormat string `json:"dateFormat,omitempty"`

	// MinValue is the minimum date value
	MinValue *LabelFieldDateValue `json:"minValue,omitempty"`

	// MaxValue is the maximum date value
	MaxValue *LabelFieldDateValue `json:"maxValue,omitempty"`
}

// LabelFieldDateValue represents a date value
type LabelFieldDateValue struct {
	// Year is the year
	Year int `json:"year,omitempty"`

	// Month is the month (1-12)
	Month int `json:"month,omitempty"`

	// Day is the day (1-31)
	Day int `json:"day,omitempty"`
}

// LabelFieldSelectionOptions for selection fields
type LabelFieldSelectionOptions struct {
	// ListOptions contains the list of choices
	ListOptions *LabelFieldListOptions `json:"listOptions,omitempty"`

	// Choices are the available choices
	Choices []*LabelFieldChoice `json:"choices,omitempty"`
}

// LabelFieldListOptions contains list options
type LabelFieldListOptions struct {
	// MaxEntries is the maximum number of entries
	MaxEntries int `json:"maxEntries,omitempty"`
}

// LabelFieldChoice represents a choice in a selection field
type LabelFieldChoice struct {
	// ID is the unique identifier for the choice
	ID string `json:"id"`

	// Properties contains choice properties
	Properties *LabelFieldChoiceProperties `json:"properties,omitempty"`

	// Lifecycle describes the lifecycle state of the choice
	Lifecycle *LabelFieldChoiceLifecycle `json:"lifecycle,omitempty"`

	// DisplayHints provides display hints for the choice
	DisplayHints *LabelFieldChoiceDisplayHints `json:"displayHints,omitempty"`

	// SchemaCapabilities describes what can be done with the choice schema
	SchemaCapabilities *LabelFieldChoiceSchemaCapabilities `json:"schemaCapabilities,omitempty"`

	// AppliedCapabilities describes what can be done with the applied choice
	AppliedCapabilities *LabelFieldChoiceAppliedCapabilities `json:"appliedCapabilities,omitempty"`
}

// LabelFieldChoiceProperties contains choice properties
type LabelFieldChoiceProperties struct {
	// DisplayName is the display name of the choice
	DisplayName string `json:"displayName,omitempty"`

	// Description is the description of the choice
	Description string `json:"description,omitempty"`

	// BadgeConfig contains badge configuration
	BadgeConfig *LabelFieldBadgeConfig `json:"badgeConfig,omitempty"`

	// InsertBeforeChoice is the choice to insert before
	InsertBeforeChoice string `json:"insertBeforeChoice,omitempty"`
}

// LabelFieldBadgeConfig contains badge configuration
type LabelFieldBadgeConfig struct {
	// Color is the badge color
	Color *LabelFieldBadgeColor `json:"color,omitempty"`

	// PriorityOverride is the priority override
	PriorityOverride int64 `json:"priorityOverride,omitempty"`
}

// LabelFieldBadgeColor represents a badge color
type LabelFieldBadgeColor struct {
	// Red is the red component (0-1)
	Red float64 `json:"red,omitempty"`

	// Green is the green component (0-1)
	Green float64 `json:"green,omitempty"`

	// Blue is the blue component (0-1)
	Blue float64 `json:"blue,omitempty"`

	// Alpha is the alpha component (0-1)
	Alpha float64 `json:"alpha,omitempty"`
}

// LabelFieldChoiceLifecycle describes the lifecycle state of a choice
type LabelFieldChoiceLifecycle struct {
	// State is the lifecycle state (UNPUBLISHED_DRAFT, PUBLISHED, DISABLED, DELETED)
	State string `json:"state,omitempty"`

	// DisabledPolicy describes why the choice is disabled
	DisabledPolicy *LabelFieldChoiceDisabledPolicy `json:"disabledPolicy,omitempty"`
}

// LabelFieldChoiceDisabledPolicy describes why a choice is disabled
type LabelFieldChoiceDisabledPolicy struct {
	// HideInSearch indicates if the choice should be hidden in search
	HideInSearch bool `json:"hideInSearch,omitempty"`

	// ShowInApply indicates if the choice should be shown in apply
	ShowInApply bool `json:"showInApply,omitempty"`
}

// LabelFieldChoiceDisplayHints provides display hints for a choice
type LabelFieldChoiceDisplayHints struct {
	// Disabled indicates if the choice is disabled
	Disabled bool `json:"disabled,omitempty"`

	// HiddenInSearch indicates if the choice is hidden in search
	HiddenInSearch bool `json:"hiddenInSearch,omitempty"`

	// ShownInApply indicates if the choice is shown in apply
	ShownInApply bool `json:"shownInApply,omitempty"`

	// BadgeColors contains badge colors
	BadgeColors *LabelFieldBadgeColor `json:"badgeColors,omitempty"`

	// DarkBadgeColors contains dark badge colors
	DarkBadgeColors *LabelFieldBadgeColor `json:"darkBadgeColors,omitempty"`

	// BadgePriority is the badge priority
	BadgePriority int64 `json:"badgePriority,omitempty"`
}

// LabelFieldChoiceSchemaCapabilities describes what can be done with a choice schema
type LabelFieldChoiceSchemaCapabilities struct {
	// CanUpdate indicates if the choice can be updated
	CanUpdate bool `json:"canUpdate,omitempty"`

	// CanDelete indicates if the choice can be deleted
	CanDelete bool `json:"canDelete,omitempty"`

	// CanDisable indicates if the choice can be disabled
	CanDisable bool `json:"canDisable,omitempty"`

	// CanEnable indicates if the choice can be enabled
	CanEnable bool `json:"canEnable,omitempty"`
}

// LabelFieldChoiceAppliedCapabilities describes what can be done with an applied choice
type LabelFieldChoiceAppliedCapabilities struct {
	// CanRead indicates if the choice can be read
	CanRead bool `json:"canRead,omitempty"`

	// CanSearch indicates if the choice can be searched
	CanSearch bool `json:"canSearch,omitempty"`

	// CanSelect indicates if the choice can be selected
	CanSelect bool `json:"canSelect,omitempty"`
}

// LabelFieldUserOptions for user fields
type LabelFieldUserOptions struct {
	// ListOptions contains the list options
	ListOptions *LabelFieldListOptions `json:"listOptions,omitempty"`
}

// FileLabel represents a label applied to a file
type FileLabel struct {
	// ID is the label ID
	ID string `json:"id"`

	// RevisionID is the revision ID of the label
	RevisionID string `json:"revisionId,omitempty"`

	// Fields are the field values
	Fields map[string]*LabelFieldValue `json:"fields,omitempty"`
}

// LabelFieldValue represents a field value
type LabelFieldValue struct {
	// ValueType is the type of value (text, integer, date, selection, user)
	ValueType string `json:"valueType,omitempty"`

	// Text is the text value
	Text []string `json:"text,omitempty"`

	// Integer is the integer value
	Integer []int64 `json:"integer,omitempty"`

	// Date is the date value
	Date []*LabelFieldDateValue `json:"date,omitempty"`

	// Selection is the selection value (choice IDs)
	Selection []string `json:"selection,omitempty"`

	// User is the user value (person IDs)
	User []*LabelUser `json:"user,omitempty"`
}

// LabelListOptions configures label list operations
type LabelListOptions struct {
	// Customer is the customer ID (for admin operations)
	Customer string

	// View is the view mode (LABEL_VIEW_BASIC, LABEL_VIEW_FULL)
	View string

	// MinimumRole is the minimum role required (READER, APPLIER, ORGANIZER, EDITOR)
	MinimumRole string

	// PublishedOnly indicates if only published labels should be returned
	PublishedOnly bool

	// Limit is the maximum number of results per page
	Limit int

	// PageToken for pagination
	PageToken string

	// Fields to return (optional)
	Fields string
}

// LabelGetOptions configures label get operations
type LabelGetOptions struct {
	// View is the view mode (LABEL_VIEW_BASIC, LABEL_VIEW_FULL)
	View string

	// UseAdminAccess indicates if admin access should be used
	UseAdminAccess bool

	// Fields to return (optional)
	Fields string
}

// LabelCreateOptions configures label create operations
type LabelCreateOptions struct {
	// UseAdminAccess indicates if admin access should be used
	UseAdminAccess bool

	// LanguageCode is the language code for the label
	LanguageCode string
}

// LabelPublishOptions configures label publish operations
type LabelPublishOptions struct {
	// UseAdminAccess indicates if admin access should be used
	UseAdminAccess bool

	// WriteControl contains write control settings
	WriteControl *LabelWriteControl
}

// LabelWriteControl contains write control settings
type LabelWriteControl struct {
	// RequiredRevisionID is the required revision ID for the operation
	RequiredRevisionID string `json:"requiredRevisionId,omitempty"`
}

// LabelDisableOptions configures label disable operations
type LabelDisableOptions struct {
	// UseAdminAccess indicates if admin access should be used
	UseAdminAccess bool

	// WriteControl contains write control settings
	WriteControl *LabelWriteControl

	// DisabledPolicy describes the disabled policy
	DisabledPolicy *LabelDisabledPolicy
}

// FileLabelListOptions configures file label list operations
type FileLabelListOptions struct {
	// View is the view mode (LABEL_VIEW_BASIC, LABEL_VIEW_FULL)
	View string

	// Fields to return (optional)
	Fields string
}

// FileLabelApplyOptions configures file label apply operations
type FileLabelApplyOptions struct {
	// Fields are the field values to set
	Fields map[string]*LabelFieldValue
}

// FileLabelUpdateOptions configures file label update operations
type FileLabelUpdateOptions struct {
	// Fields are the field values to update
	Fields map[string]*LabelFieldValue
}

// Label view modes
const (
	LabelViewBasic = "LABEL_VIEW_BASIC"
	LabelViewFull  = "LABEL_VIEW_FULL"
)

// Label lifecycle states
const (
	LabelStateUnpublishedDraft = "UNPUBLISHED_DRAFT"
	LabelStatePublished        = "PUBLISHED"
	LabelStateDisabled         = "DISABLED"
	LabelStateDeleted          = "DELETED"
)

// Label types
const (
	LabelTypeShared = "SHARED"
	LabelTypeAdmin  = "ADMIN"
)

// Label copy modes
const (
	LabelCopyModeDoNotCopy     = "DO_NOT_COPY"
	LabelCopyModeAlwaysCopy    = "ALWAYS_COPY"
	LabelCopyModeCopyAppliable = "COPY_APPLIABLE"
)

// Label minimum roles
const (
	LabelRoleReader    = "READER"
	LabelRoleApplier   = "APPLIER"
	LabelRoleOrganizer = "ORGANIZER"
	LabelRoleEditor    = "EDITOR"
)
