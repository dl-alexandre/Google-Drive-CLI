package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/dl-alexandre/gdrive/internal/api"
	"github.com/dl-alexandre/gdrive/internal/auth"
	"github.com/dl-alexandre/gdrive/internal/files"
	slidesmgr "github.com/dl-alexandre/gdrive/internal/slides"
	"github.com/dl-alexandre/gdrive/internal/types"
	"github.com/dl-alexandre/gdrive/internal/utils"
	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
	slidesapi "google.golang.org/api/slides/v1"
)

var slidesCmd = &cobra.Command{
	Use:   "slides",
	Short: "Google Slides operations",
	Long:  "Commands for managing Google Slides files",
}

var slidesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Google Slides files",
	Long:  "List Google Slides files using Drive API with MIME type filter",
	RunE:  runSlidesList,
}

var slidesGetCmd = &cobra.Command{
	Use:   "get <presentation-id>",
	Short: "Get presentation metadata",
	Args:  cobra.ExactArgs(1),
	RunE:  runSlidesGet,
}

var slidesReadCmd = &cobra.Command{
	Use:   "read <presentation-id>",
	Short: "Read presentation content",
	Args:  cobra.ExactArgs(1),
	RunE:  runSlidesRead,
}

var slidesCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new presentation",
	Args:  cobra.ExactArgs(1),
	RunE:  runSlidesCreate,
}

var slidesUpdateCmd = &cobra.Command{
	Use:   "update <presentation-id>",
	Short: "Update presentation content",
	Args:  cobra.ExactArgs(1),
	RunE:  runSlidesUpdate,
}

var slidesReplaceCmd = &cobra.Command{
	Use:   "replace <presentation-id>",
	Short: "Replace text placeholders",
	Long:  "Replace text placeholders with values using ReplaceAllText requests",
	Args:  cobra.ExactArgs(1),
	RunE:  runSlidesReplace,
}

var (
	slidesListParentID   string
	slidesListQuery      string
	slidesListLimit      int
	slidesListPageToken  string
	slidesListOrderBy    string
	slidesListFields     string
	slidesListPaginate   bool
	slidesCreateParentID string
	slidesUpdateRequests string
	slidesUpdateFile     string
	slidesReplaceData    string
	slidesReplaceFile    string
)

func init() {
	slidesListCmd.Flags().StringVar(&slidesListParentID, "parent", "", "Parent folder ID")
	slidesListCmd.Flags().StringVar(&slidesListQuery, "query", "", "Additional search query")
	slidesListCmd.Flags().IntVar(&slidesListLimit, "limit", 100, "Maximum files to return per page")
	slidesListCmd.Flags().StringVar(&slidesListPageToken, "page-token", "", "Page token for pagination")
	slidesListCmd.Flags().StringVar(&slidesListOrderBy, "order-by", "", "Sort order")
	slidesListCmd.Flags().StringVar(&slidesListFields, "fields", "", "Fields to return")
	slidesListCmd.Flags().BoolVar(&slidesListPaginate, "paginate", false, "Automatically fetch all pages")

	slidesCreateCmd.Flags().StringVar(&slidesCreateParentID, "parent", "", "Parent folder ID")
	slidesUpdateCmd.Flags().StringVar(&slidesUpdateRequests, "requests", "", "Batch update requests JSON")
	slidesUpdateCmd.Flags().StringVar(&slidesUpdateFile, "requests-file", "", "Path to JSON file with batch update requests")
	slidesReplaceCmd.Flags().StringVar(&slidesReplaceData, "data", "", "JSON string with replacements map")
	slidesReplaceCmd.Flags().StringVar(&slidesReplaceFile, "file", "", "Path to JSON file with replacements map")

	slidesCmd.AddCommand(slidesListCmd)
	slidesCmd.AddCommand(slidesGetCmd)
	slidesCmd.AddCommand(slidesReadCmd)
	slidesCmd.AddCommand(slidesCreateCmd)
	slidesCmd.AddCommand(slidesUpdateCmd)
	slidesCmd.AddCommand(slidesReplaceCmd)
	rootCmd.AddCommand(slidesCmd)
}

func runSlidesList(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	_, client, reqCtx, err := getSlidesService(ctx, flags)
	if err != nil {
		return out.WriteError("slides.list", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	parentID := slidesListParentID
	if parentID != "" {
		resolvedID, err := ResolveFileID(ctx, client, flags, parentID)
		if err != nil {
			if appErr, ok := err.(*utils.AppError); ok {
				return out.WriteError("slides.list", appErr.CLIError)
			}
			return out.WriteError("slides.list", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
		}
		parentID = resolvedID
	}

	query := fmt.Sprintf("mimeType = '%s'", utils.MimeTypePresentation)
	if parentID != "" {
		query = fmt.Sprintf("'%s' in parents and %s", parentID, query)
	}
	if slidesListQuery != "" {
		query = fmt.Sprintf("%s and (%s)", query, slidesListQuery)
	}

	opts := files.ListOptions{
		ParentID:       parentID,
		Query:          query,
		PageSize:       slidesListLimit,
		PageToken:      slidesListPageToken,
		OrderBy:        slidesListOrderBy,
		IncludeTrashed: false,
		Fields:         slidesListFields,
	}

	mgr := files.NewManager(client)
	reqCtx.RequestType = types.RequestTypeListOrSearch

	if slidesListPaginate {
		allFiles, err := mgr.ListAll(ctx, reqCtx, opts)
		if err != nil {
			if appErr, ok := err.(*utils.AppError); ok {
				return out.WriteError("slides.list", appErr.CLIError)
			}
			return out.WriteError("slides.list", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
		}
		if flags.OutputFormat == types.OutputFormatTable {
			return out.WriteSuccess("slides.list", allFiles)
		}
		return out.WriteSuccess("slides.list", map[string]interface{}{
			"files": allFiles,
		})
	}

	result, err := mgr.List(ctx, reqCtx, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("slides.list", appErr.CLIError)
		}
		return out.WriteError("slides.list", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	if flags.OutputFormat == types.OutputFormatTable {
		return out.WriteSuccess("slides.list", result.Files)
	}
	return out.WriteSuccess("slides.list", result)
}

func runSlidesGet(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getSlidesService(ctx, flags)
	if err != nil {
		return out.WriteError("slides.get", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	presentationID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("slides.get", appErr.CLIError)
		}
		return out.WriteError("slides.get", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	mgr := slidesmgr.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeGetByID
	result, err := mgr.GetPresentation(ctx, reqCtx, presentationID)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("slides.get", appErr.CLIError)
		}
		return out.WriteError("slides.get", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("slides.get", result)
}

func runSlidesRead(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getSlidesService(ctx, flags)
	if err != nil {
		return out.WriteError("slides.read", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	presentationID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("slides.read", appErr.CLIError)
		}
		return out.WriteError("slides.read", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	mgr := slidesmgr.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeGetByID
	result, err := mgr.ReadPresentation(ctx, reqCtx, presentationID)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("slides.read", appErr.CLIError)
		}
		return out.WriteError("slides.read", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("slides.read", result)
}

func runSlidesCreate(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	_, client, reqCtx, err := getSlidesService(ctx, flags)
	if err != nil {
		return out.WriteError("slides.create", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	parentID := slidesCreateParentID
	if parentID != "" {
		resolvedID, err := ResolveFileID(ctx, client, flags, parentID)
		if err != nil {
			if appErr, ok := err.(*utils.AppError); ok {
				return out.WriteError("slides.create", appErr.CLIError)
			}
			return out.WriteError("slides.create", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
		}
		parentID = resolvedID
	}

	reqCtx.RequestType = types.RequestTypeMutation
	metadata := &drive.File{
		Name:     args[0],
		MimeType: utils.MimeTypePresentation,
	}
	if parentID != "" {
		metadata.Parents = []string{parentID}
		reqCtx.InvolvedParentIDs = append(reqCtx.InvolvedParentIDs, parentID)
	}

	call := client.Service().Files.Create(metadata)
	call = api.NewRequestShaper(client).ShapeFilesCreate(call, reqCtx)
	call = call.Fields("id,name,mimeType,size,createdTime,modifiedTime,parents,resourceKey,trashed,capabilities")

	result, err := api.ExecuteWithRetry(ctx, client, reqCtx, func() (*drive.File, error) {
		return call.Do()
	})
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("slides.create", appErr.CLIError)
		}
		return out.WriteError("slides.create", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	if result.ResourceKey != "" {
		client.ResourceKeys().UpdateFromAPIResponse(result.Id, result.ResourceKey)
	}

	return out.WriteSuccess("slides.create", convertDriveFile(result))
}

func runSlidesUpdate(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getSlidesService(ctx, flags)
	if err != nil {
		return out.WriteError("slides.update", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	requests, err := readSlidesRequests()
	if err != nil {
		return out.WriteError("slides.update", utils.NewCLIError(utils.ErrCodeInvalidArgument, err.Error()).Build())
	}

	presentationID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("slides.update", appErr.CLIError)
		}
		return out.WriteError("slides.update", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	mgr := slidesmgr.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeMutation
	result, err := mgr.UpdatePresentation(ctx, reqCtx, presentationID, requests)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("slides.update", appErr.CLIError)
		}
		return out.WriteError("slides.update", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("slides.update", result)
}

func runSlidesReplace(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getSlidesService(ctx, flags)
	if err != nil {
		return out.WriteError("slides.replace", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	replacements, err := readSlidesReplacements()
	if err != nil {
		return out.WriteError("slides.replace", utils.NewCLIError(utils.ErrCodeInvalidArgument, err.Error()).Build())
	}

	presentationID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("slides.replace", appErr.CLIError)
		}
		return out.WriteError("slides.replace", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	mgr := slidesmgr.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeMutation
	result, err := mgr.ReplaceAllText(ctx, reqCtx, presentationID, replacements)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("slides.replace", appErr.CLIError)
		}
		return out.WriteError("slides.replace", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("slides.replace", result)
}

func getSlidesService(ctx context.Context, flags types.GlobalFlags) (*slidesapi.Service, *api.Client, *types.RequestContext, error) {
	configDir := getConfigDir()
	authMgr := auth.NewManager(configDir)

	creds, err := authMgr.GetValidCredentials(ctx, flags.Profile)
	if err != nil {
		return nil, nil, nil, err
	}

	if err := authMgr.ValidateServiceScopes(creds, auth.ServiceSlides); err != nil {
		return nil, nil, nil, err
	}

	svc, err := authMgr.GetSlidesService(ctx, creds)
	if err != nil {
		return nil, nil, nil, err
	}

	driveSvc, err := authMgr.GetDriveService(ctx, creds)
	if err != nil {
		return nil, nil, nil, err
	}

	client := api.NewClient(driveSvc, utils.DefaultMaxRetries, utils.DefaultRetryDelayMs, GetLogger())
	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypeListOrSearch)
	return svc, client, reqCtx, nil
}

func readSlidesRequests() ([]*slidesapi.Request, error) {
	if slidesUpdateRequests == "" && slidesUpdateFile == "" {
		return nil, fmt.Errorf("requests required via --requests or --requests-file")
	}

	var raw []byte
	if slidesUpdateFile != "" {
		data, err := os.ReadFile(slidesUpdateFile)
		if err != nil {
			return nil, err
		}
		raw = data
	} else {
		raw = []byte(slidesUpdateRequests)
	}

	var requests []*slidesapi.Request
	if err := json.Unmarshal(raw, &requests); err != nil {
		return nil, err
	}
	if len(requests) == 0 {
		return nil, fmt.Errorf("at least one request is required")
	}
	return requests, nil
}

func readSlidesReplacements() (map[string]string, error) {
	if slidesReplaceData == "" && slidesReplaceFile == "" {
		return nil, fmt.Errorf("replacements required via --data or --file")
	}

	var raw []byte
	if slidesReplaceFile != "" {
		data, err := os.ReadFile(slidesReplaceFile)
		if err != nil {
			return nil, err
		}
		raw = data
	} else {
		raw = []byte(slidesReplaceData)
	}

	var replacements map[string]string
	if err := json.Unmarshal(raw, &replacements); err != nil {
		return nil, err
	}
	if len(replacements) == 0 {
		return nil, fmt.Errorf("at least one replacement is required")
	}
	return replacements, nil
}

