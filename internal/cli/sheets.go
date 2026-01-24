package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/dl-alexandre/gdrive/internal/api"
	"github.com/dl-alexandre/gdrive/internal/auth"
	"github.com/dl-alexandre/gdrive/internal/files"
	sheetsmgr "github.com/dl-alexandre/gdrive/internal/sheets"
	"github.com/dl-alexandre/gdrive/internal/types"
	"github.com/dl-alexandre/gdrive/internal/utils"
	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
	sheetsapi "google.golang.org/api/sheets/v4"
)

var sheetsCmd = &cobra.Command{
	Use:   "sheets",
	Short: "Google Sheets operations",
	Long:  "Commands for managing Google Sheets files",
}

var sheetsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Google Sheets files",
	Long:  "List Google Sheets files using Drive API with MIME type filter",
	RunE:  runSheetsList,
}

var sheetsCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new spreadsheet",
	Long:  "Create a new Google Sheets spreadsheet",
	Args:  cobra.ExactArgs(1),
	RunE:  runSheetsCreate,
}

var sheetsGetCmd = &cobra.Command{
	Use:   "get <spreadsheet-id>",
	Short: "Get spreadsheet metadata",
	Args:  cobra.ExactArgs(1),
	RunE:  runSheetsGet,
}

var sheetsValuesCmd = &cobra.Command{
	Use:   "values",
	Short: "Read or write spreadsheet values",
}

var sheetsValuesGetCmd = &cobra.Command{
	Use:   "get <spreadsheet-id> <range>",
	Short: "Get values from a range",
	Args:  cobra.ExactArgs(2),
	RunE:  runSheetsValuesGet,
}

var sheetsValuesUpdateCmd = &cobra.Command{
	Use:   "update <spreadsheet-id> <range>",
	Short: "Update values in a range",
	Args:  cobra.ExactArgs(2),
	RunE:  runSheetsValuesUpdate,
}

var sheetsValuesAppendCmd = &cobra.Command{
	Use:   "append <spreadsheet-id> <range>",
	Short: "Append values to a range",
	Args:  cobra.ExactArgs(2),
	RunE:  runSheetsValuesAppend,
}

var sheetsValuesClearCmd = &cobra.Command{
	Use:   "clear <spreadsheet-id> <range>",
	Short: "Clear values from a range",
	Args:  cobra.ExactArgs(2),
	RunE:  runSheetsValuesClear,
}

var sheetsBatchUpdateCmd = &cobra.Command{
	Use:   "batch-update <spreadsheet-id>",
	Short: "Batch update spreadsheet",
	Args:  cobra.ExactArgs(1),
	RunE:  runSheetsBatchUpdate,
}

var (
	sheetsValuesJSON        string
	sheetsValuesFile        string
	sheetsValueInputOption  string
	sheetsBatchUpdateJSON   string
	sheetsBatchUpdateFile   string
	sheetsListParentID      string
	sheetsListQuery         string
	sheetsListLimit         int
	sheetsListPageToken     string
	sheetsListOrderBy       string
	sheetsListFields        string
	sheetsListPaginate      bool
	sheetsCreateParentID    string
)

func init() {
	sheetsValuesUpdateCmd.Flags().StringVar(&sheetsValuesJSON, "values", "", "Values JSON (2D array)")
	sheetsValuesUpdateCmd.Flags().StringVar(&sheetsValuesFile, "values-file", "", "Path to JSON file with values (2D array)")
	sheetsValuesUpdateCmd.Flags().StringVar(&sheetsValueInputOption, "value-input-option", "USER_ENTERED", "Value input option (RAW or USER_ENTERED)")
	sheetsValuesAppendCmd.Flags().StringVar(&sheetsValuesJSON, "values", "", "Values JSON (2D array)")
	sheetsValuesAppendCmd.Flags().StringVar(&sheetsValuesFile, "values-file", "", "Path to JSON file with values (2D array)")
	sheetsValuesAppendCmd.Flags().StringVar(&sheetsValueInputOption, "value-input-option", "USER_ENTERED", "Value input option (RAW or USER_ENTERED)")
	sheetsBatchUpdateCmd.Flags().StringVar(&sheetsBatchUpdateJSON, "requests", "", "Batch update requests JSON")
	sheetsBatchUpdateCmd.Flags().StringVar(&sheetsBatchUpdateFile, "requests-file", "", "Path to JSON file with batch update requests")

	sheetsListCmd.Flags().StringVar(&sheetsListParentID, "parent", "", "Parent folder ID")
	sheetsListCmd.Flags().StringVar(&sheetsListQuery, "query", "", "Additional search query")
	sheetsListCmd.Flags().IntVar(&sheetsListLimit, "limit", 100, "Maximum files to return per page")
	sheetsListCmd.Flags().StringVar(&sheetsListPageToken, "page-token", "", "Page token for pagination")
	sheetsListCmd.Flags().StringVar(&sheetsListOrderBy, "order-by", "", "Sort order")
	sheetsListCmd.Flags().StringVar(&sheetsListFields, "fields", "", "Fields to return")
	sheetsListCmd.Flags().BoolVar(&sheetsListPaginate, "paginate", false, "Automatically fetch all pages")

	sheetsCreateCmd.Flags().StringVar(&sheetsCreateParentID, "parent", "", "Parent folder ID")

	sheetsValuesCmd.AddCommand(sheetsValuesGetCmd)
	sheetsValuesCmd.AddCommand(sheetsValuesUpdateCmd)
	sheetsValuesCmd.AddCommand(sheetsValuesAppendCmd)
	sheetsValuesCmd.AddCommand(sheetsValuesClearCmd)

	sheetsCmd.AddCommand(sheetsListCmd)
	sheetsCmd.AddCommand(sheetsCreateCmd)
	sheetsCmd.AddCommand(sheetsGetCmd)
	sheetsCmd.AddCommand(sheetsBatchUpdateCmd)
	sheetsCmd.AddCommand(sheetsValuesCmd)
	rootCmd.AddCommand(sheetsCmd)
}

func runSheetsGet(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getSheetsService(ctx, flags)
	if err != nil {
		return out.WriteError("sheets.get", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	spreadsheetID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("sheets.get", appErr.CLIError)
		}
		return out.WriteError("sheets.get", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	mgr := sheetsmgr.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeGetByID
	result, err := mgr.GetSpreadsheet(ctx, reqCtx, spreadsheetID)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("sheets.get", appErr.CLIError)
		}
		return out.WriteError("sheets.get", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("sheets.get", result)
}

func runSheetsValuesGet(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getSheetsService(ctx, flags)
	if err != nil {
		return out.WriteError("sheets.values.get", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	spreadsheetID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("sheets.values.get", appErr.CLIError)
		}
		return out.WriteError("sheets.values.get", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	mgr := sheetsmgr.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeGetByID
	result, err := mgr.GetValues(ctx, reqCtx, spreadsheetID, args[1])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("sheets.values.get", appErr.CLIError)
		}
		return out.WriteError("sheets.values.get", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("sheets.values.get", result)
}

func runSheetsValuesUpdate(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getSheetsService(ctx, flags)
	if err != nil {
		return out.WriteError("sheets.values.update", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	values, err := readSheetValues()
	if err != nil {
		return out.WriteError("sheets.values.update", utils.NewCLIError(utils.ErrCodeInvalidArgument, err.Error()).Build())
	}

	spreadsheetID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("sheets.values.update", appErr.CLIError)
		}
		return out.WriteError("sheets.values.update", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	mgr := sheetsmgr.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeMutation
	result, err := mgr.UpdateValues(ctx, reqCtx, spreadsheetID, args[1], values, sheetsValueInputOption)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("sheets.values.update", appErr.CLIError)
		}
		return out.WriteError("sheets.values.update", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("sheets.values.update", result)
}

func runSheetsValuesAppend(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getSheetsService(ctx, flags)
	if err != nil {
		return out.WriteError("sheets.values.append", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	values, err := readSheetValues()
	if err != nil {
		return out.WriteError("sheets.values.append", utils.NewCLIError(utils.ErrCodeInvalidArgument, err.Error()).Build())
	}

	spreadsheetID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("sheets.values.append", appErr.CLIError)
		}
		return out.WriteError("sheets.values.append", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	mgr := sheetsmgr.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeMutation
	result, err := mgr.AppendValues(ctx, reqCtx, spreadsheetID, args[1], values, sheetsValueInputOption)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("sheets.values.append", appErr.CLIError)
		}
		return out.WriteError("sheets.values.append", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("sheets.values.append", result)
}

func runSheetsValuesClear(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	if strings.TrimSpace(args[1]) == "" {
		return out.WriteError("sheets.values.clear", utils.NewCLIError(utils.ErrCodeInvalidArgument, "range is required").Build())
	}

	svc, client, reqCtx, err := getSheetsService(ctx, flags)
	if err != nil {
		return out.WriteError("sheets.values.clear", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	spreadsheetID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("sheets.values.clear", appErr.CLIError)
		}
		return out.WriteError("sheets.values.clear", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	mgr := sheetsmgr.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeMutation
	result, err := mgr.ClearValues(ctx, reqCtx, spreadsheetID, args[1])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("sheets.values.clear", appErr.CLIError)
		}
		return out.WriteError("sheets.values.clear", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("sheets.values.clear", result)
}

func runSheetsBatchUpdate(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getSheetsService(ctx, flags)
	if err != nil {
		return out.WriteError("sheets.batch-update", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	requests, err := readSheetsBatchRequests()
	if err != nil {
		return out.WriteError("sheets.batch-update", utils.NewCLIError(utils.ErrCodeInvalidArgument, err.Error()).Build())
	}

	spreadsheetID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("sheets.batch-update", appErr.CLIError)
		}
		return out.WriteError("sheets.batch-update", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	mgr := sheetsmgr.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeMutation
	result, err := mgr.BatchUpdate(ctx, reqCtx, spreadsheetID, requests)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("sheets.batch-update", appErr.CLIError)
		}
		return out.WriteError("sheets.batch-update", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("sheets.batch-update", result)
}
func runSheetsList(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	_, client, reqCtx, err := getSheetsService(ctx, flags)
	if err != nil {
		return out.WriteError("sheets.list", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	parentID := sheetsListParentID
	if parentID != "" {
		resolvedID, err := ResolveFileID(ctx, client, flags, parentID)
		if err != nil {
			if appErr, ok := err.(*utils.AppError); ok {
				return out.WriteError("sheets.list", appErr.CLIError)
			}
			return out.WriteError("sheets.list", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
		}
		parentID = resolvedID
	}

	query := fmt.Sprintf("mimeType = '%s'", utils.MimeTypeSpreadsheet)
	if parentID != "" {
		query = fmt.Sprintf("'%s' in parents and %s", parentID, query)
	}
	if sheetsListQuery != "" {
		query = fmt.Sprintf("%s and (%s)", query, sheetsListQuery)
	}

	opts := files.ListOptions{
		ParentID:       parentID,
		Query:          query,
		PageSize:       sheetsListLimit,
		PageToken:      sheetsListPageToken,
		OrderBy:        sheetsListOrderBy,
		IncludeTrashed: false,
		Fields:         sheetsListFields,
	}

	mgr := files.NewManager(client)
	reqCtx.RequestType = types.RequestTypeListOrSearch

	if sheetsListPaginate {
		allFiles, err := mgr.ListAll(ctx, reqCtx, opts)
		if err != nil {
			if appErr, ok := err.(*utils.AppError); ok {
				return out.WriteError("sheets.list", appErr.CLIError)
			}
			return out.WriteError("sheets.list", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
		}
		if flags.OutputFormat == types.OutputFormatTable {
			return out.WriteSuccess("sheets.list", allFiles)
		}
		return out.WriteSuccess("sheets.list", map[string]interface{}{
			"files": allFiles,
		})
	}

	result, err := mgr.List(ctx, reqCtx, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("sheets.list", appErr.CLIError)
		}
		return out.WriteError("sheets.list", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	if flags.OutputFormat == types.OutputFormatTable {
		return out.WriteSuccess("sheets.list", result.Files)
	}
	return out.WriteSuccess("sheets.list", result)
}

func runSheetsCreate(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	_, client, reqCtx, err := getSheetsService(ctx, flags)
	if err != nil {
		return out.WriteError("sheets.create", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	parentID := sheetsCreateParentID
	if parentID != "" {
		resolvedID, err := ResolveFileID(ctx, client, flags, parentID)
		if err != nil {
			if appErr, ok := err.(*utils.AppError); ok {
				return out.WriteError("sheets.create", appErr.CLIError)
			}
			return out.WriteError("sheets.create", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
		}
		parentID = resolvedID
	}

	reqCtx.RequestType = types.RequestTypeMutation
	metadata := &drive.File{
		Name:     args[0],
		MimeType: utils.MimeTypeSpreadsheet,
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
			return out.WriteError("sheets.create", appErr.CLIError)
		}
		return out.WriteError("sheets.create", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	if result.ResourceKey != "" {
		client.ResourceKeys().UpdateFromAPIResponse(result.Id, result.ResourceKey)
	}

	return out.WriteSuccess("sheets.create", convertDriveFile(result))
}

func getSheetsService(ctx context.Context, flags types.GlobalFlags) (*sheetsapi.Service, *api.Client, *types.RequestContext, error) {
	configDir := getConfigDir()
	authMgr := auth.NewManager(configDir)

	creds, err := authMgr.GetValidCredentials(ctx, flags.Profile)
	if err != nil {
		return nil, nil, nil, err
	}

	if err := authMgr.ValidateServiceScopes(creds, auth.ServiceSheets); err != nil {
		return nil, nil, nil, err
	}

	svc, err := authMgr.GetSheetsService(ctx, creds)
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

func readSheetValues() ([][]interface{}, error) {
	if sheetsValuesJSON == "" && sheetsValuesFile == "" {
		return nil, fmt.Errorf("values required via --values or --values-file")
	}

	var raw []byte
	if sheetsValuesFile != "" {
		data, err := os.ReadFile(sheetsValuesFile)
		if err != nil {
			return nil, err
		}
		raw = data
	} else {
		raw = []byte(sheetsValuesJSON)
	}

	var values [][]interface{}
	if err := json.Unmarshal(raw, &values); err != nil {
		return nil, err
	}
	return values, nil
}

func readSheetsBatchRequests() ([]*sheetsapi.Request, error) {
	if sheetsBatchUpdateJSON == "" && sheetsBatchUpdateFile == "" {
		return nil, fmt.Errorf("requests required via --requests or --requests-file")
	}

	var raw []byte
	if sheetsBatchUpdateFile != "" {
		data, err := os.ReadFile(sheetsBatchUpdateFile)
		if err != nil {
			return nil, err
		}
		raw = data
	} else {
		raw = []byte(sheetsBatchUpdateJSON)
	}

	var requests []*sheetsapi.Request
	if err := json.Unmarshal(raw, &requests); err != nil {
		return nil, err
	}
	if len(requests) == 0 {
		return nil, fmt.Errorf("at least one request is required")
	}
	return requests, nil
}

