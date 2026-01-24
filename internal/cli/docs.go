package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/dl-alexandre/gdrive/internal/api"
	"github.com/dl-alexandre/gdrive/internal/auth"
	docsmgr "github.com/dl-alexandre/gdrive/internal/docs"
	"github.com/dl-alexandre/gdrive/internal/files"
	"github.com/dl-alexandre/gdrive/internal/types"
	"github.com/dl-alexandre/gdrive/internal/utils"
	"github.com/spf13/cobra"
	docsapi "google.golang.org/api/docs/v1"
	"google.golang.org/api/drive/v3"
)

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Google Docs operations",
	Long:  "Commands for managing Google Docs files",
}

var docsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Google Docs files",
	Long:  "List Google Docs files using Drive API with MIME type filter",
	RunE:  runDocsList,
}

var docsGetCmd = &cobra.Command{
	Use:   "get <document-id>",
	Short: "Get document metadata",
	Args:  cobra.ExactArgs(1),
	RunE:  runDocsGet,
}

var docsReadCmd = &cobra.Command{
	Use:   "read <document-id>",
	Short: "Read document content",
	Args:  cobra.ExactArgs(1),
	RunE:  runDocsRead,
}

var docsCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new document",
	Args:  cobra.ExactArgs(1),
	RunE:  runDocsCreate,
}

var docsUpdateCmd = &cobra.Command{
	Use:   "update <document-id>",
	Short: "Update document content",
	Args:  cobra.ExactArgs(1),
	RunE:  runDocsUpdate,
}

var (
	docsListParentID   string
	docsListQuery      string
	docsListLimit      int
	docsListPageToken  string
	docsListOrderBy    string
	docsListFields     string
	docsListPaginate   bool
	docsCreateParentID string
	docsUpdateRequests string
	docsUpdateFile     string
)

func init() {
	docsListCmd.Flags().StringVar(&docsListParentID, "parent", "", "Parent folder ID")
	docsListCmd.Flags().StringVar(&docsListQuery, "query", "", "Additional search query")
	docsListCmd.Flags().IntVar(&docsListLimit, "limit", 100, "Maximum files to return per page")
	docsListCmd.Flags().StringVar(&docsListPageToken, "page-token", "", "Page token for pagination")
	docsListCmd.Flags().StringVar(&docsListOrderBy, "order-by", "", "Sort order")
	docsListCmd.Flags().StringVar(&docsListFields, "fields", "", "Fields to return")
	docsListCmd.Flags().BoolVar(&docsListPaginate, "paginate", false, "Automatically fetch all pages")

	docsCreateCmd.Flags().StringVar(&docsCreateParentID, "parent", "", "Parent folder ID")
	docsUpdateCmd.Flags().StringVar(&docsUpdateRequests, "requests", "", "Batch update requests JSON")
	docsUpdateCmd.Flags().StringVar(&docsUpdateFile, "requests-file", "", "Path to JSON file with batch update requests")

	docsCmd.AddCommand(docsListCmd)
	docsCmd.AddCommand(docsGetCmd)
	docsCmd.AddCommand(docsReadCmd)
	docsCmd.AddCommand(docsCreateCmd)
	docsCmd.AddCommand(docsUpdateCmd)
	rootCmd.AddCommand(docsCmd)
}

func runDocsList(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	_, client, reqCtx, err := getDocsService(ctx, flags)
	if err != nil {
		return out.WriteError("docs.list", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	parentID := docsListParentID
	if parentID != "" {
		resolvedID, err := ResolveFileID(ctx, client, flags, parentID)
		if err != nil {
			if appErr, ok := err.(*utils.AppError); ok {
				return out.WriteError("docs.list", appErr.CLIError)
			}
			return out.WriteError("docs.list", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
		}
		parentID = resolvedID
	}

	query := fmt.Sprintf("mimeType = '%s'", utils.MimeTypeDocument)
	if parentID != "" {
		query = fmt.Sprintf("'%s' in parents and %s", parentID, query)
	}
	if docsListQuery != "" {
		query = fmt.Sprintf("%s and (%s)", query, docsListQuery)
	}

	opts := files.ListOptions{
		ParentID:       parentID,
		Query:          query,
		PageSize:       docsListLimit,
		PageToken:      docsListPageToken,
		OrderBy:        docsListOrderBy,
		IncludeTrashed: false,
		Fields:         docsListFields,
	}

	mgr := files.NewManager(client)
	reqCtx.RequestType = types.RequestTypeListOrSearch

	if docsListPaginate {
		allFiles, err := mgr.ListAll(ctx, reqCtx, opts)
		if err != nil {
			if appErr, ok := err.(*utils.AppError); ok {
				return out.WriteError("docs.list", appErr.CLIError)
			}
			return out.WriteError("docs.list", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
		}
		if flags.OutputFormat == types.OutputFormatTable {
			return out.WriteSuccess("docs.list", allFiles)
		}
		return out.WriteSuccess("docs.list", map[string]interface{}{
			"files": allFiles,
		})
	}

	result, err := mgr.List(ctx, reqCtx, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("docs.list", appErr.CLIError)
		}
		return out.WriteError("docs.list", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	if flags.OutputFormat == types.OutputFormatTable {
		return out.WriteSuccess("docs.list", result.Files)
	}
	return out.WriteSuccess("docs.list", result)
}

func runDocsGet(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getDocsService(ctx, flags)
	if err != nil {
		return out.WriteError("docs.get", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	documentID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("docs.get", appErr.CLIError)
		}
		return out.WriteError("docs.get", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	mgr := docsmgr.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeGetByID
	result, err := mgr.GetDocument(ctx, reqCtx, documentID)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("docs.get", appErr.CLIError)
		}
		return out.WriteError("docs.get", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("docs.get", result)
}

func runDocsRead(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getDocsService(ctx, flags)
	if err != nil {
		return out.WriteError("docs.read", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	documentID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("docs.read", appErr.CLIError)
		}
		return out.WriteError("docs.read", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	mgr := docsmgr.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeGetByID
	result, err := mgr.ReadDocument(ctx, reqCtx, documentID)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("docs.read", appErr.CLIError)
		}
		return out.WriteError("docs.read", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("docs.read", result)
}

func runDocsCreate(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	_, client, reqCtx, err := getDocsService(ctx, flags)
	if err != nil {
		return out.WriteError("docs.create", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	parentID := docsCreateParentID
	if parentID != "" {
		resolvedID, err := ResolveFileID(ctx, client, flags, parentID)
		if err != nil {
			if appErr, ok := err.(*utils.AppError); ok {
				return out.WriteError("docs.create", appErr.CLIError)
			}
			return out.WriteError("docs.create", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
		}
		parentID = resolvedID
	}

	reqCtx.RequestType = types.RequestTypeMutation
	metadata := &drive.File{
		Name:     args[0],
		MimeType: utils.MimeTypeDocument,
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
			return out.WriteError("docs.create", appErr.CLIError)
		}
		return out.WriteError("docs.create", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	if result.ResourceKey != "" {
		client.ResourceKeys().UpdateFromAPIResponse(result.Id, result.ResourceKey)
	}

	return out.WriteSuccess("docs.create", convertDriveFile(result))
}

func runDocsUpdate(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	out := NewOutputWriter(flags.OutputFormat, flags.Quiet, flags.Verbose)
	ctx := context.Background()

	svc, client, reqCtx, err := getDocsService(ctx, flags)
	if err != nil {
		return out.WriteError("docs.update", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	requests, err := readDocsRequests()
	if err != nil {
		return out.WriteError("docs.update", utils.NewCLIError(utils.ErrCodeInvalidArgument, err.Error()).Build())
	}

	documentID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("docs.update", appErr.CLIError)
		}
		return out.WriteError("docs.update", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	mgr := docsmgr.NewManager(client, svc)
	reqCtx.RequestType = types.RequestTypeMutation
	result, err := mgr.UpdateDocument(ctx, reqCtx, documentID, requests)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("docs.update", appErr.CLIError)
		}
		return out.WriteError("docs.update", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("docs.update", result)
}

func getDocsService(ctx context.Context, flags types.GlobalFlags) (*docsapi.Service, *api.Client, *types.RequestContext, error) {
	configDir := getConfigDir()
	authMgr := auth.NewManager(configDir)

	creds, err := authMgr.GetValidCredentials(ctx, flags.Profile)
	if err != nil {
		return nil, nil, nil, err
	}

	if err := authMgr.ValidateServiceScopes(creds, auth.ServiceDocs); err != nil {
		return nil, nil, nil, err
	}

	svc, err := authMgr.GetDocsService(ctx, creds)
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

func readDocsRequests() ([]*docsapi.Request, error) {
	if docsUpdateRequests == "" && docsUpdateFile == "" {
		return nil, fmt.Errorf("requests required via --requests or --requests-file")
	}

	var raw []byte
	if docsUpdateFile != "" {
		data, err := os.ReadFile(docsUpdateFile)
		if err != nil {
			return nil, err
		}
		raw = data
	} else {
		raw = []byte(docsUpdateRequests)
	}

	var requests []*docsapi.Request
	if err := json.Unmarshal(raw, &requests); err != nil {
		return nil, err
	}
	if len(requests) == 0 {
		return nil, fmt.Errorf("at least one request is required")
	}
	return requests, nil
}

