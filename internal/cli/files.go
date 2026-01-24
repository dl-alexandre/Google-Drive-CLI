package cli

import (
	"context"

	"github.com/dl-alexandre/gdrive/internal/api"
	"github.com/dl-alexandre/gdrive/internal/auth"
	"github.com/dl-alexandre/gdrive/internal/export"
	"github.com/dl-alexandre/gdrive/internal/files"
	"github.com/dl-alexandre/gdrive/internal/revisions"
	"github.com/dl-alexandre/gdrive/internal/types"
	"github.com/dl-alexandre/gdrive/internal/utils"
	"github.com/spf13/cobra"
)

var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "File operations",
	Long:  "Manage files in Google Drive",
}

var filesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List files",
	RunE:  runFilesList,
}

var filesGetCmd = &cobra.Command{
	Use:   "get <file-id>",
	Short: "Get file metadata",
	Args:  cobra.ExactArgs(1),
	RunE:  runFilesGet,
}

var filesUploadCmd = &cobra.Command{
	Use:   "upload <local-path>",
	Short: "Upload a file",
	Args:  cobra.ExactArgs(1),
	RunE:  runFilesUpload,
}

var filesDownloadCmd = &cobra.Command{
	Use:   "download <file-id>",
	Short: "Download a file",
	Args:  cobra.ExactArgs(1),
	RunE:  runFilesDownload,
}

var filesDeleteCmd = &cobra.Command{
	Use:   "delete <file-id>",
	Short: "Delete a file",
	Args:  cobra.ExactArgs(1),
	RunE:  runFilesDelete,
}

var filesCopyCmd = &cobra.Command{
	Use:   "copy <file-id>",
	Short: "Copy a file",
	Args:  cobra.ExactArgs(1),
	RunE:  runFilesCopy,
}

var filesMoveCmd = &cobra.Command{
	Use:   "move <file-id>",
	Short: "Move a file",
	Args:  cobra.ExactArgs(1),
	RunE:  runFilesMove,
}

var filesTrashCmd = &cobra.Command{
	Use:   "trash <file-id>",
	Short: "Move file to trash",
	Args:  cobra.ExactArgs(1),
	RunE:  runFilesTrash,
}

var filesRestoreCmd = &cobra.Command{
	Use:   "restore <file-id>",
	Short: "Restore file from trash",
	Args:  cobra.ExactArgs(1),
	RunE:  runFilesRestore,
}

var filesRevisionsCmd = &cobra.Command{
	Use:   "revisions <file-id>",
	Short: "List file revisions",
	Args:  cobra.ExactArgs(1),
	RunE:  runFilesRevisions,
}

var filesRevisionsDownloadCmd = &cobra.Command{
	Use:   "revisions download <file-id> <revision-id>",
	Short: "Download a specific revision",
	Args:  cobra.ExactArgs(2),
	RunE:  runFilesRevisionsDownload,
}

var filesRevisionsRestoreCmd = &cobra.Command{
	Use:   "revisions restore <file-id> <revision-id>",
	Short: "Restore file to a specific revision",
	Args:  cobra.ExactArgs(2),
	RunE:  runFilesRevisionsRestore,
}

var filesListTrashedCmd = &cobra.Command{
	Use:   "list-trashed",
	Short: "List trashed files",
	RunE:  runFilesListTrashed,
}

var filesExportFormatsCmd = &cobra.Command{
	Use:   "export-formats <file-id>",
	Short: "Show available export formats for a file",
	Args:  cobra.ExactArgs(1),
	RunE:  runFilesExportFormats,
}

// Command flags
var (
	filesParentID       string
	filesQuery          string
	filesLimit          int
	filesPageToken      string
	filesOrderBy        string
	filesIncludeTrashed bool
	filesFields         string
	filesGetFields      string
	filesName           string
	filesMimeType       string
	filesOutput         string
	filesPermanent      bool
	filesForce          bool
	filesDownloadDoc    bool
	filesRevisionOutput string
	filesPaginate       bool
)

func init() {
	// List flags
	filesListCmd.Flags().StringVar(&filesParentID, "parent", "", "Parent folder ID")
	filesListCmd.Flags().StringVar(&filesQuery, "query", "", "Search query")
	filesListCmd.Flags().IntVar(&filesLimit, "limit", 100, "Maximum files to return per page")
	filesListCmd.Flags().StringVar(&filesPageToken, "page-token", "", "Page token for pagination")
	filesListCmd.Flags().StringVar(&filesOrderBy, "order-by", "", "Sort order")
	filesListCmd.Flags().BoolVar(&filesIncludeTrashed, "include-trashed", false, "Include trashed files")
	filesListCmd.Flags().StringVar(&filesFields, "fields", "", "Fields to return")
	filesListCmd.Flags().BoolVar(&filesPaginate, "paginate", false, "Automatically fetch all pages")

	// Get flags
	filesGetCmd.Flags().StringVar(&filesGetFields, "fields", "", "Fields to return")

	// Upload flags
	filesUploadCmd.Flags().StringVar(&filesParentID, "parent", "", "Parent folder ID")
	filesUploadCmd.Flags().StringVar(&filesName, "name", "", "File name")
	filesUploadCmd.Flags().StringVar(&filesMimeType, "mime-type", "", "MIME type")

	// Download flags
	filesDownloadCmd.Flags().StringVar(&filesOutput, "output", "", "Output path")
	filesDownloadCmd.Flags().StringVar(&filesMimeType, "mime-type", "", "Export MIME type")
	filesDownloadCmd.Flags().BoolVar(&filesDownloadDoc, "doc", false, "Export Google Docs as plain text")
	filesDownloadCmd.Flags().BoolVar(&filesDownloadDoc, "doc-text", false, "Export Google Docs as plain text")

	// Delete flags
	filesDeleteCmd.Flags().BoolVar(&filesPermanent, "permanent", false, "Permanently delete")
	filesDeleteCmd.Flags().BoolVar(&filesForce, "force", false, "Skip confirmation")

	// Copy flags
	filesCopyCmd.Flags().StringVar(&filesName, "name", "", "New file name")
	filesCopyCmd.Flags().StringVar(&filesParentID, "parent", "", "Destination folder ID")

	// Move flags
	filesMoveCmd.Flags().StringVar(&filesParentID, "parent", "", "New parent folder ID")
	_ = filesMoveCmd.MarkFlagRequired("parent")

	// Revision flags
	filesRevisionsDownloadCmd.Flags().StringVar(&filesRevisionOutput, "output", "", "Output path for revision download")
	_ = filesRevisionsDownloadCmd.MarkFlagRequired("output")

	filesRevisionsCmd.AddCommand(filesRevisionsDownloadCmd)
	filesRevisionsCmd.AddCommand(filesRevisionsRestoreCmd)

	// List trashed flags
	filesListTrashedCmd.Flags().StringVar(&filesQuery, "query", "", "Search query")
	filesListTrashedCmd.Flags().IntVar(&filesLimit, "limit", 100, "Maximum files to return per page")
	filesListTrashedCmd.Flags().StringVar(&filesPageToken, "page-token", "", "Page token for pagination")
	filesListTrashedCmd.Flags().StringVar(&filesOrderBy, "order-by", "", "Sort order")
	filesListTrashedCmd.Flags().StringVar(&filesFields, "fields", "", "Fields to return")
	filesListTrashedCmd.Flags().BoolVar(&filesPaginate, "paginate", false, "Automatically fetch all pages")

	filesCmd.AddCommand(filesListCmd)
	filesCmd.AddCommand(filesGetCmd)
	filesCmd.AddCommand(filesUploadCmd)
	filesCmd.AddCommand(filesDownloadCmd)
	filesCmd.AddCommand(filesDeleteCmd)
	filesCmd.AddCommand(filesCopyCmd)
	filesCmd.AddCommand(filesMoveCmd)
	filesCmd.AddCommand(filesTrashCmd)
	filesCmd.AddCommand(filesRestoreCmd)
	filesCmd.AddCommand(filesRevisionsCmd)
	filesCmd.AddCommand(filesListTrashedCmd)
	filesCmd.AddCommand(filesExportFormatsCmd)
	rootCmd.AddCommand(filesCmd)
}

func getFileManager(ctx context.Context, flags types.GlobalFlags) (*files.Manager, *api.Client, *types.RequestContext, *OutputWriter, error) {
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
	mgr := files.NewManager(client)
	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypeListOrSearch)

	return mgr, client, reqCtx, out, nil
}

func runFilesList(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()

	mgr, client, reqCtx, out, err := getFileManager(ctx, flags)
	if err != nil {
		return out.WriteError("files.list", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	// Resolve parent path if provided
	parentID := filesParentID
	if parentID != "" {
		resolvedID, err := ResolveFileID(ctx, client, flags, parentID)
		if err != nil {
			if appErr, ok := err.(*utils.AppError); ok {
				return out.WriteError("files.list", appErr.CLIError)
			}
			return out.WriteError("files.list", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
		}
		parentID = resolvedID
	}

	opts := files.ListOptions{
		ParentID:       parentID,
		Query:          filesQuery,
		PageSize:       filesLimit,
		PageToken:      filesPageToken,
		OrderBy:        filesOrderBy,
		IncludeTrashed: filesIncludeTrashed,
		Fields:         filesFields,
	}

	// If --paginate flag is set, fetch all pages
	if filesPaginate {
		allFiles, err := mgr.ListAll(ctx, reqCtx, opts)
		if err != nil {
			if appErr, ok := err.(*utils.AppError); ok {
				return out.WriteError("files.list", appErr.CLIError)
			}
			return out.WriteError("files.list", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
		}
		// Return result without nextPageToken (all pages fetched)
		return out.WriteSuccess("files.list", map[string]interface{}{
			"files": allFiles,
		})
	}

	result, err := mgr.List(ctx, reqCtx, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.list", appErr.CLIError)
		}
		return out.WriteError("files.list", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("files.list", result)
}

func runFilesGet(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()

	mgr, client, reqCtx, out, err := getFileManager(ctx, flags)
	if err != nil {
		return out.WriteError("files.get", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	// Resolve file ID from path if needed
	fileID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.get", appErr.CLIError)
		}
		return out.WriteError("files.get", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	reqCtx.RequestType = types.RequestTypeGetByID
	file, err := mgr.Get(ctx, reqCtx, fileID, filesGetFields)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.get", appErr.CLIError)
		}
		return out.WriteError("files.get", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("files.get", file)
}

func runFilesUpload(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()

	mgr, client, reqCtx, out, err := getFileManager(ctx, flags)
	if err != nil {
		return out.WriteError("files.upload", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	// Resolve parent path if provided
	parentID := filesParentID
	if parentID != "" {
		resolvedID, err := ResolveFileID(ctx, client, flags, parentID)
		if err != nil {
			if appErr, ok := err.(*utils.AppError); ok {
				return out.WriteError("files.upload", appErr.CLIError)
			}
			return out.WriteError("files.upload", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
		}
		parentID = resolvedID
	}

	reqCtx.RequestType = types.RequestTypeMutation
	file, err := mgr.Upload(ctx, reqCtx, args[0], files.UploadOptions{
		ParentID: parentID,
		Name:     filesName,
		MimeType: filesMimeType,
	})
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.upload", appErr.CLIError)
		}
		return out.WriteError("files.upload", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	out.Log("Uploaded: %s", file.Name)
	return out.WriteSuccess("files.upload", file)
}

func runFilesDownload(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()

	mgr, client, reqCtx, out, err := getFileManager(ctx, flags)
	if err != nil {
		return out.WriteError("files.download", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	// Resolve file ID from path if needed
	fileID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.download", appErr.CLIError)
		}
		return out.WriteError("files.download", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	reqCtx.RequestType = types.RequestTypeDownloadOrExport
	mimeType := filesMimeType
	if filesDownloadDoc && mimeType == "" {
		mimeType = "text/plain"
	}

	err = mgr.Download(ctx, reqCtx, fileID, files.DownloadOptions{
		OutputPath: filesOutput,
		MimeType:   mimeType,
	})
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.download", appErr.CLIError)
		}
		return out.WriteError("files.download", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	out.Log("Downloaded to: %s", filesOutput)
	return out.WriteSuccess("files.download", map[string]string{"path": filesOutput})
}

func runFilesDelete(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()

	mgr, client, reqCtx, out, err := getFileManager(ctx, flags)
	if err != nil {
		return out.WriteError("files.delete", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	// Resolve file ID from path if needed
	fileID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.delete", appErr.CLIError)
		}
		return out.WriteError("files.delete", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	reqCtx.RequestType = types.RequestTypeMutation
	err = mgr.Delete(ctx, reqCtx, fileID, filesPermanent)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.delete", appErr.CLIError)
		}
		return out.WriteError("files.delete", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	action := "trashed"
	if filesPermanent {
		action = "deleted"
	}
	out.Log("File %s: %s", action, fileID)
	return out.WriteSuccess("files.delete", map[string]string{"id": fileID, "status": action})
}

func runFilesCopy(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()

	mgr, client, reqCtx, out, err := getFileManager(ctx, flags)
	if err != nil {
		return out.WriteError("files.copy", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	// Resolve file ID from path if needed
	fileID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.copy", appErr.CLIError)
		}
		return out.WriteError("files.copy", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	// Resolve parent path if provided
	parentID := filesParentID
	if parentID != "" {
		resolvedID, err := ResolveFileID(ctx, client, flags, parentID)
		if err != nil {
			if appErr, ok := err.(*utils.AppError); ok {
				return out.WriteError("files.copy", appErr.CLIError)
			}
			return out.WriteError("files.copy", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
		}
		parentID = resolvedID
	}

	reqCtx.RequestType = types.RequestTypeMutation
	file, err := mgr.Copy(ctx, reqCtx, fileID, filesName, parentID)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.copy", appErr.CLIError)
		}
		return out.WriteError("files.copy", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	out.Log("Copied to: %s", file.Name)
	return out.WriteSuccess("files.copy", file)
}

func runFilesMove(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()

	mgr, client, reqCtx, out, err := getFileManager(ctx, flags)
	if err != nil {
		return out.WriteError("files.move", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	// Resolve file ID from path if needed
	fileID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.move", appErr.CLIError)
		}
		return out.WriteError("files.move", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	// Resolve parent path
	parentID, err := ResolveFileID(ctx, client, flags, filesParentID)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.move", appErr.CLIError)
		}
		return out.WriteError("files.move", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	reqCtx.RequestType = types.RequestTypeMutation
	file, err := mgr.Move(ctx, reqCtx, fileID, parentID)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.move", appErr.CLIError)
		}
		return out.WriteError("files.move", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	out.Log("Moved: %s", file.Name)
	return out.WriteSuccess("files.move", file)
}

func runFilesTrash(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()

	mgr, client, reqCtx, out, err := getFileManager(ctx, flags)
	if err != nil {
		return out.WriteError("files.trash", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	// Resolve file ID from path if needed
	fileID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.trash", appErr.CLIError)
		}
		return out.WriteError("files.trash", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	reqCtx.RequestType = types.RequestTypeMutation
	file, err := mgr.Trash(ctx, reqCtx, fileID)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.trash", appErr.CLIError)
		}
		return out.WriteError("files.trash", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	out.Log("Trashed: %s", file.Name)
	return out.WriteSuccess("files.trash", file)
}

func runFilesRestore(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()

	mgr, client, reqCtx, out, err := getFileManager(ctx, flags)
	if err != nil {
		return out.WriteError("files.restore", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	// Resolve file ID from path if needed
	fileID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.restore", appErr.CLIError)
		}
		return out.WriteError("files.restore", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	reqCtx.RequestType = types.RequestTypeMutation
	file, err := mgr.Restore(ctx, reqCtx, fileID)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.restore", appErr.CLIError)
		}
		return out.WriteError("files.restore", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	out.Log("Restored: %s", file.Name)
	return out.WriteSuccess("files.restore", file)
}

func runFilesRevisions(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()

	_, client, reqCtx, out, err := getFileManager(ctx, flags)
	if err != nil {
		return out.WriteError("files.revisions", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	// Resolve file ID from path if needed
	fileID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.revisions", appErr.CLIError)
		}
		return out.WriteError("files.revisions", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	// Create revisions manager
	revMgr := revisions.NewManager(client)
	reqCtx.RequestType = types.RequestTypeListOrSearch

	result, err := revMgr.List(ctx, reqCtx, fileID, revisions.ListOptions{})
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.revisions", appErr.CLIError)
		}
		return out.WriteError("files.revisions", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("files.revisions", result)
}

func runFilesRevisionsDownload(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()

	_, client, reqCtx, out, err := getFileManager(ctx, flags)
	if err != nil {
		return out.WriteError("files.revisions.download", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	// Resolve file ID from path if needed
	fileID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.revisions.download", appErr.CLIError)
		}
		return out.WriteError("files.revisions.download", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	revisionID := args[1]

	// Create revisions manager
	revMgr := revisions.NewManager(client)
	reqCtx.RequestType = types.RequestTypeDownloadOrExport

	err = revMgr.Download(ctx, reqCtx, fileID, revisionID, revisions.DownloadOptions{
		OutputPath: filesRevisionOutput,
	})
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.revisions.download", appErr.CLIError)
		}
		return out.WriteError("files.revisions.download", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	out.Log("Downloaded revision %s to: %s", revisionID, filesRevisionOutput)
	return out.WriteSuccess("files.revisions.download", map[string]string{"revisionId": revisionID, "path": filesRevisionOutput})
}

func runFilesRevisionsRestore(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()

	_, client, reqCtx, out, err := getFileManager(ctx, flags)
	if err != nil {
		return out.WriteError("files.revisions.restore", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	// Resolve file ID from path if needed
	fileID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.revisions.restore", appErr.CLIError)
		}
		return out.WriteError("files.revisions.restore", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	revisionID := args[1]

	// Create revisions manager
	revMgr := revisions.NewManager(client)
	reqCtx.RequestType = types.RequestTypeMutation

	file, err := revMgr.Restore(ctx, reqCtx, fileID, revisionID)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.revisions.restore", appErr.CLIError)
		}
		return out.WriteError("files.revisions.restore", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	out.Log("Restored file to revision: %s", revisionID)
	return out.WriteSuccess("files.revisions.restore", file)
}

func runFilesListTrashed(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()

	mgr, _, reqCtx, out, err := getFileManager(ctx, flags)
	if err != nil {
		return out.WriteError("files.list-trashed", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	opts := files.ListOptions{
		Query:     filesQuery,
		PageSize:  filesLimit,
		PageToken: filesPageToken,
		OrderBy:   filesOrderBy,
		Fields:    filesFields,
	}

	// If --paginate flag is set, fetch all pages
	if filesPaginate {
		// Use ListAll with trashed query
		opts.IncludeTrashed = true
		if opts.Query != "" {
			opts.Query = "trashed = true and (" + opts.Query + ")"
		} else {
			opts.Query = "trashed = true"
		}
		allFiles, err := mgr.ListAll(ctx, reqCtx, opts)
		if err != nil {
			if appErr, ok := err.(*utils.AppError); ok {
				return out.WriteError("files.list-trashed", appErr.CLIError)
			}
			return out.WriteError("files.list-trashed", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
		}
		return out.WriteSuccess("files.list-trashed", map[string]interface{}{
			"files": allFiles,
		})
	}

	result, err := mgr.ListTrashed(ctx, reqCtx, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.list-trashed", appErr.CLIError)
		}
		return out.WriteError("files.list-trashed", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	return out.WriteSuccess("files.list-trashed", result)
}

func runFilesExportFormats(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()

	mgr, client, reqCtx, out, err := getFileManager(ctx, flags)
	if err != nil {
		return out.WriteError("files.export-formats", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	// Resolve file ID from path if needed
	fileID, err := ResolveFileID(ctx, client, flags, args[0])
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.export-formats", appErr.CLIError)
		}
		return out.WriteError("files.export-formats", utils.NewCLIError(utils.ErrCodeInvalidPath, err.Error()).Build())
	}

	reqCtx.RequestType = types.RequestTypeGetByID
	file, err := mgr.Get(ctx, reqCtx, fileID, "mimeType,name")
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("files.export-formats", appErr.CLIError)
		}
		return out.WriteError("files.export-formats", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	formats, err := export.GetAvailableFormats(file.MimeType)
	if err != nil {
		return out.WriteError("files.export-formats", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	result := map[string]interface{}{
		"file": map[string]string{
			"id":       fileID,
			"name":     file.Name,
			"mimeType": file.MimeType,
		},
		"availableFormats": formats,
	}

	return out.WriteSuccess("files.export-formats", result)
}
