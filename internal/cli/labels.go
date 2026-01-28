package cli

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/auth"
	"github.com/dl-alexandre/gdrv/internal/labels"
	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
	"github.com/spf13/cobra"
)

var labelsCmd = &cobra.Command{
	Use:   "labels",
	Short: "Drive Labels API operations",
	Long:  "Manage Drive labels and apply labels to files",
}

var labelsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available labels",
	Long: `List available Drive labels.

Examples:
  # List all labels
  gdrv labels list --json

  # List labels with full details
  gdrv labels list --view LABEL_VIEW_FULL --json

  # List labels for a specific customer (admin)
  gdrv labels list --customer C01234567 --json

  # List only published labels
  gdrv labels list --published-only --json`,
	RunE: runLabelsList,
}

var labelsGetCmd = &cobra.Command{
	Use:   "get <label-id>",
	Short: "Get label schema",
	Long: `Get the schema for a specific label.

Examples:
  # Get label details
  gdrv labels get labels/abc123 --json

  # Get label with full details
  gdrv labels get labels/abc123 --view LABEL_VIEW_FULL --json

  # Get label with admin access
  gdrv labels get labels/abc123 --use-admin-access --json`,
	Args: cobra.ExactArgs(1),
	RunE: runLabelsGet,
}

var labelsCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new label (admin)",
	Long: `Create a new Drive label (requires admin access).

Examples:
  # Create a simple label
  gdrv labels create "Document Type" --json

  # Create a label with admin access
  gdrv labels create "Project Status" --use-admin-access --json

  # Create a label with language code
  gdrv labels create "Status" --language-code en --json`,
	Args: cobra.ExactArgs(1),
	RunE: runLabelsCreate,
}

var labelsPublishCmd = &cobra.Command{
	Use:   "publish <label-id>",
	Short: "Publish a label (admin)",
	Long: `Publish a label to make it available for use.

Examples:
  # Publish a label
  gdrv labels publish labels/abc123 --json

  # Publish with admin access
  gdrv labels publish labels/abc123 --use-admin-access --json`,
	Args: cobra.ExactArgs(1),
	RunE: runLabelsPublish,
}

var labelsDisableCmd = &cobra.Command{
	Use:   "disable <label-id>",
	Short: "Disable a label (admin)",
	Long: `Disable a label to prevent it from being applied to new files.

Examples:
  # Disable a label
  gdrv labels disable labels/abc123 --json

  # Disable with admin access
  gdrv labels disable labels/abc123 --use-admin-access --json

  # Disable and hide in search
  gdrv labels disable labels/abc123 --hide-in-search --json`,
	Args: cobra.ExactArgs(1),
	RunE: runLabelsDisable,
}

var labelsFileCmd = &cobra.Command{
	Use:   "file",
	Short: "File label operations",
	Long:  "Manage labels on files",
}

var labelsFileListCmd = &cobra.Command{
	Use:   "list <file-id>",
	Short: "List labels on a file",
	Long: `List all labels applied to a file.

Examples:
  # List labels on a file
  gdrv labels file list 1abc123... --json

  # List with full details
  gdrv labels file list 1abc123... --view LABEL_VIEW_FULL --json`,
	Args: cobra.ExactArgs(1),
	RunE: runLabelsFileList,
}

var labelsFileApplyCmd = &cobra.Command{
	Use:   "apply <file-id> <label-id>",
	Short: "Apply a label to a file",
	Long: `Apply a label to a file with optional field values.

Examples:
  # Apply a label without field values
  gdrv labels file apply 1abc123... labels/abc123 --json

  # Apply a label with field values
  gdrv labels file apply 1abc123... labels/abc123 --fields '{"field1":{"valueType":"text","text":["value1"]}}' --json`,
	Args: cobra.ExactArgs(2),
	RunE: runLabelsFileApply,
}

var labelsFileUpdateCmd = &cobra.Command{
	Use:   "update <file-id> <label-id>",
	Short: "Update label fields on a file",
	Long: `Update the field values of a label applied to a file.

Examples:
  # Update label field values
  gdrv labels file update 1abc123... labels/abc123 --fields '{"field1":{"valueType":"text","text":["new value"]}}' --json`,
	Args: cobra.ExactArgs(2),
	RunE: runLabelsFileUpdate,
}

var labelsFileRemoveCmd = &cobra.Command{
	Use:   "remove <file-id> <label-id>",
	Short: "Remove a label from a file",
	Long: `Remove a label from a file.

Examples:
  # Remove a label from a file
  gdrv labels file remove 1abc123... labels/abc123`,
	Args: cobra.ExactArgs(2),
	RunE: runLabelsFileRemove,
}

var (
	labelsCustomer       string
	labelsView           string
	labelsMinimumRole    string
	labelsPublishedOnly  bool
	labelsLimit          int
	labelsPageToken      string
	labelsFields         string
	labelsUseAdminAccess bool
	labelsLanguageCode   string
	labelsHideInSearch   bool
	labelsShowInApply    bool
	labelsFieldValues    string
)

func init() {
	labelsListCmd.Flags().StringVar(&labelsCustomer, "customer", "", "Customer ID (for admin operations)")
	labelsListCmd.Flags().StringVar(&labelsView, "view", "", "View mode (LABEL_VIEW_BASIC, LABEL_VIEW_FULL)")
	labelsListCmd.Flags().StringVar(&labelsMinimumRole, "minimum-role", "", "Minimum role (READER, APPLIER, ORGANIZER, EDITOR)")
	labelsListCmd.Flags().BoolVar(&labelsPublishedOnly, "published-only", false, "Only return published labels")
	labelsListCmd.Flags().IntVar(&labelsLimit, "limit", 100, "Maximum results per page")
	labelsListCmd.Flags().StringVar(&labelsPageToken, "page-token", "", "Pagination token")
	labelsListCmd.Flags().StringVar(&labelsFields, "fields", "", "Fields to return")

	labelsGetCmd.Flags().StringVar(&labelsView, "view", "", "View mode (LABEL_VIEW_BASIC, LABEL_VIEW_FULL)")
	labelsGetCmd.Flags().BoolVar(&labelsUseAdminAccess, "use-admin-access", false, "Use admin access")
	labelsGetCmd.Flags().StringVar(&labelsFields, "fields", "", "Fields to return")

	labelsCreateCmd.Flags().BoolVar(&labelsUseAdminAccess, "use-admin-access", false, "Use admin access")
	labelsCreateCmd.Flags().StringVar(&labelsLanguageCode, "language-code", "", "Language code (e.g., en)")

	labelsPublishCmd.Flags().BoolVar(&labelsUseAdminAccess, "use-admin-access", false, "Use admin access")

	labelsDisableCmd.Flags().BoolVar(&labelsUseAdminAccess, "use-admin-access", false, "Use admin access")
	labelsDisableCmd.Flags().BoolVar(&labelsHideInSearch, "hide-in-search", false, "Hide in search")
	labelsDisableCmd.Flags().BoolVar(&labelsShowInApply, "show-in-apply", false, "Show in apply")

	labelsFileListCmd.Flags().StringVar(&labelsView, "view", "", "View mode (LABEL_VIEW_BASIC, LABEL_VIEW_FULL)")
	labelsFileListCmd.Flags().StringVar(&labelsFields, "fields", "", "Fields to return")

	labelsFileApplyCmd.Flags().StringVar(&labelsFieldValues, "fields", "", "Field values as JSON")

	labelsFileUpdateCmd.Flags().StringVar(&labelsFieldValues, "fields", "", "Field values as JSON")

	labelsFileCmd.AddCommand(labelsFileListCmd)
	labelsFileCmd.AddCommand(labelsFileApplyCmd)
	labelsFileCmd.AddCommand(labelsFileUpdateCmd)
	labelsFileCmd.AddCommand(labelsFileRemoveCmd)

	labelsCmd.AddCommand(labelsListCmd)
	labelsCmd.AddCommand(labelsGetCmd)
	labelsCmd.AddCommand(labelsCreateCmd)
	labelsCmd.AddCommand(labelsPublishCmd)
	labelsCmd.AddCommand(labelsDisableCmd)
	labelsCmd.AddCommand(labelsFileCmd)

	rootCmd.AddCommand(labelsCmd)
}

func getLabelsManager(ctx context.Context, flags types.GlobalFlags) (*labels.Manager, *api.Client, *types.RequestContext, *OutputWriter, error) {
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	configDir := getConfigDir()
	authMgr := auth.NewManager(configDir)

	creds, err := authMgr.GetValidCredentials(ctx, flags.Profile)
	if err != nil {
		return nil, nil, nil, out, err
	}

	service, err := authMgr.GetDriveService(ctx, creds)
	if err != nil {
		return nil, nil, nil, out, err
	}

	client := api.NewClient(service, utils.DefaultMaxRetries, utils.DefaultRetryDelayMs, GetLogger())
	mgr := labels.NewManager(client)
	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypeListOrSearch)

	return mgr, client, reqCtx, out, nil
}

func runLabelsList(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()

	mgr, _, reqCtx, out, err := getLabelsManager(ctx, flags)
	if err != nil {
		return out.WriteError("labels.list", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	opts := types.LabelListOptions{
		Customer:      labelsCustomer,
		View:          labelsView,
		MinimumRole:   labelsMinimumRole,
		PublishedOnly: labelsPublishedOnly,
		Limit:         labelsLimit,
		PageToken:     labelsPageToken,
		Fields:        labelsFields,
	}

	labelsList, nextPageToken, err := mgr.List(ctx, reqCtx, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("labels.list", appErr.CLIError)
		}
		return out.WriteError("labels.list", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	result := &LabelsListResult{
		Labels:        labelsList,
		NextPageToken: nextPageToken,
	}
	return out.WriteSuccess("labels.list", result)
}

func runLabelsGet(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()
	labelID := args[0]

	mgr, _, reqCtx, out, err := getLabelsManager(ctx, flags)
	if err != nil {
		return out.WriteError("labels.get", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	opts := types.LabelGetOptions{
		View:           labelsView,
		UseAdminAccess: labelsUseAdminAccess,
		Fields:         labelsFields,
	}

	label, err := mgr.Get(ctx, reqCtx, labelID, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("labels.get", appErr.CLIError)
		}
		return out.WriteError("labels.get", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	result := &LabelResult{Label: label}
	return out.WriteSuccess("labels.get", result)
}

func runLabelsCreate(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()
	name := args[0]

	mgr, _, reqCtx, out, err := getLabelsManager(ctx, flags)
	if err != nil {
		return out.WriteError("labels.create", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	label := &types.Label{
		LabelType: types.LabelTypeShared,
		Properties: &types.LabelProperties{
			Title: name,
		},
	}

	opts := types.LabelCreateOptions{
		UseAdminAccess: labelsUseAdminAccess,
		LanguageCode:   labelsLanguageCode,
	}

	createdLabel, err := mgr.CreateLabel(ctx, reqCtx, label, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("labels.create", appErr.CLIError)
		}
		return out.WriteError("labels.create", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	result := &LabelResult{Label: createdLabel}
	return out.WriteSuccess("labels.create", result)
}

func runLabelsPublish(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()
	labelID := args[0]

	mgr, _, reqCtx, out, err := getLabelsManager(ctx, flags)
	if err != nil {
		return out.WriteError("labels.publish", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	opts := types.LabelPublishOptions{
		UseAdminAccess: labelsUseAdminAccess,
	}

	publishedLabel, err := mgr.PublishLabel(ctx, reqCtx, labelID, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("labels.publish", appErr.CLIError)
		}
		return out.WriteError("labels.publish", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	result := &LabelResult{Label: publishedLabel}
	return out.WriteSuccess("labels.publish", result)
}

func runLabelsDisable(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()
	labelID := args[0]

	mgr, _, reqCtx, out, err := getLabelsManager(ctx, flags)
	if err != nil {
		return out.WriteError("labels.disable", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	opts := types.LabelDisableOptions{
		UseAdminAccess: labelsUseAdminAccess,
	}

	if labelsHideInSearch || labelsShowInApply {
		opts.DisabledPolicy = &types.LabelDisabledPolicy{
			HideInSearch: labelsHideInSearch,
			ShowInApply:  labelsShowInApply,
		}
	}

	disabledLabel, err := mgr.DisableLabel(ctx, reqCtx, labelID, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("labels.disable", appErr.CLIError)
		}
		return out.WriteError("labels.disable", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	result := &LabelResult{Label: disabledLabel}
	return out.WriteSuccess("labels.disable", result)
}

func runLabelsFileList(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()
	fileID := args[0]

	mgr, _, reqCtx, out, err := getLabelsManager(ctx, flags)
	if err != nil {
		return out.WriteError("labels.file.list", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	opts := types.FileLabelListOptions{
		View:   labelsView,
		Fields: labelsFields,
	}

	fileLabels, err := mgr.ListFileLabels(ctx, reqCtx, fileID, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("labels.file.list", appErr.CLIError)
		}
		return out.WriteError("labels.file.list", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	result := &FileLabelsListResult{FileLabels: fileLabels}
	return out.WriteSuccess("labels.file.list", result)
}

func runLabelsFileApply(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()
	fileID := args[0]
	labelID := args[1]

	mgr, _, reqCtx, out, err := getLabelsManager(ctx, flags)
	if err != nil {
		return out.WriteError("labels.file.apply", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	opts := types.FileLabelApplyOptions{
		Fields: make(map[string]*types.LabelFieldValue),
	}

	if labelsFieldValues != "" {
		if err := json.Unmarshal([]byte(labelsFieldValues), &opts.Fields); err != nil {
			return out.WriteError("labels.file.apply", utils.NewCLIError(utils.ErrCodeInvalidArgument,
				fmt.Sprintf("Invalid field values JSON: %s", err)).Build())
		}
	}

	fileLabel, err := mgr.ApplyLabel(ctx, reqCtx, fileID, labelID, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("labels.file.apply", appErr.CLIError)
		}
		return out.WriteError("labels.file.apply", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	result := &FileLabelResult{FileLabel: fileLabel}
	return out.WriteSuccess("labels.file.apply", result)
}

func runLabelsFileUpdate(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()
	fileID := args[0]
	labelID := args[1]

	mgr, _, reqCtx, out, err := getLabelsManager(ctx, flags)
	if err != nil {
		return out.WriteError("labels.file.update", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	opts := types.FileLabelUpdateOptions{
		Fields: make(map[string]*types.LabelFieldValue),
	}

	if labelsFieldValues != "" {
		if err := json.Unmarshal([]byte(labelsFieldValues), &opts.Fields); err != nil {
			return out.WriteError("labels.file.update", utils.NewCLIError(utils.ErrCodeInvalidArgument,
				fmt.Sprintf("Invalid field values JSON: %s", err)).Build())
		}
	}

	fileLabel, err := mgr.UpdateLabel(ctx, reqCtx, fileID, labelID, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("labels.file.update", appErr.CLIError)
		}
		return out.WriteError("labels.file.update", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	result := &FileLabelResult{FileLabel: fileLabel}
	return out.WriteSuccess("labels.file.update", result)
}

func runLabelsFileRemove(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()
	fileID := args[0]
	labelID := args[1]

	mgr, _, reqCtx, out, err := getLabelsManager(ctx, flags)
	if err != nil {
		return out.WriteError("labels.file.remove", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	err = mgr.RemoveLabel(ctx, reqCtx, fileID, labelID)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("labels.file.remove", appErr.CLIError)
		}
		return out.WriteError("labels.file.remove", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	result := &SuccessResult{Message: fmt.Sprintf("Label %s removed from file %s", labelID, fileID)}
	return out.WriteSuccess("labels.file.remove", result)
}

type LabelsListResult struct {
	Labels        []*types.Label
	NextPageToken string
}

func (r *LabelsListResult) Headers() []string {
	return []string{"ID", "Name", "Type", "State"}
}

func (r *LabelsListResult) Rows() [][]string {
	rows := make([][]string, len(r.Labels))
	for i, label := range r.Labels {
		state := ""
		if label.Lifecycle != nil {
			state = label.Lifecycle.State
		}
		rows[i] = []string{label.ID, label.Name, label.LabelType, state}
	}
	return rows
}

func (r *LabelsListResult) EmptyMessage() string {
	return "No labels found"
}

type LabelResult struct {
	Label *types.Label
}

func (r *LabelResult) Headers() []string {
	return []string{"ID", "Name", "Type", "State", "Fields"}
}

func (r *LabelResult) Rows() [][]string {
	state := ""
	if r.Label.Lifecycle != nil {
		state = r.Label.Lifecycle.State
	}
	fieldCount := fmt.Sprintf("%d", len(r.Label.Fields))
	return [][]string{{r.Label.ID, r.Label.Name, r.Label.LabelType, state, fieldCount}}
}

func (r *LabelResult) EmptyMessage() string {
	return "No label found"
}

type FileLabelsListResult struct {
	FileLabels []*types.FileLabel
}

func (r *FileLabelsListResult) Headers() []string {
	return []string{"Label ID", "Revision ID", "Fields"}
}

func (r *FileLabelsListResult) Rows() [][]string {
	rows := make([][]string, len(r.FileLabels))
	for i, fileLabel := range r.FileLabels {
		fieldCount := fmt.Sprintf("%d", len(fileLabel.Fields))
		rows[i] = []string{fileLabel.ID, fileLabel.RevisionID, fieldCount}
	}
	return rows
}

func (r *FileLabelsListResult) EmptyMessage() string {
	return "No labels found on file"
}

type FileLabelResult struct {
	FileLabel *types.FileLabel
}

func (r *FileLabelResult) Headers() []string {
	return []string{"Label ID", "Revision ID", "Fields"}
}

func (r *FileLabelResult) Rows() [][]string {
	fieldCount := fmt.Sprintf("%d", len(r.FileLabel.Fields))
	return [][]string{{r.FileLabel.ID, r.FileLabel.RevisionID, fieldCount}}
}

func (r *FileLabelResult) EmptyMessage() string {
	return "No label found"
}

type SuccessResult struct {
	Message string
}

func (r *SuccessResult) Headers() []string {
	return []string{"Status"}
}

func (r *SuccessResult) Rows() [][]string {
	return [][]string{{r.Message}}
}

func (r *SuccessResult) EmptyMessage() string {
	return ""
}
