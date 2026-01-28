package cli

import (
	"context"

	"github.com/dl-alexandre/gdrv/internal/activity"
	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/auth"
	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
	"github.com/spf13/cobra"
)

var activityCmd = &cobra.Command{
	Use:   "activity",
	Short: "Drive Activity API operations",
	Long:  "Query and monitor file and folder activity across Google Drive",
}

var activityQueryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query Drive activity",
	Long: `Query Drive activity events with filtering options.

Examples:
  # Query recent activity for all accessible files
  gdrv activity query --json

  # Query activity for a specific file
  gdrv activity query --file-id 1abc123... --json

  # Query activity within a time range
  gdrv activity query --start-time "2026-01-01T00:00:00Z" --end-time "2026-01-31T23:59:59Z" --json

  # Query activity for a folder (including descendants)
  gdrv activity query --folder-id 0ABC123... --json

  # Filter by activity types
  gdrv activity query --action-types "edit,share,permission_change" --json

  # Get activity for a specific user
  gdrv activity query --user user@example.com --json

  # Paginate through activity results
  gdrv activity query --limit 100 --page-token "TOKEN" --json`,
	RunE: runActivityQuery,
}

var (
	activityFileID      string
	activityFolderID    string
	activityAncestor    string
	activityStartTime   string
	activityEndTime     string
	activityActionTypes string
	activityUser        string
	activityLimit       int
	activityPageToken   string
	activityFields      string
)

func init() {
	activityQueryCmd.Flags().StringVar(&activityFileID, "file-id", "", "Filter by specific file ID")
	activityQueryCmd.Flags().StringVar(&activityFolderID, "folder-id", "", "Filter by folder ID (includes descendants)")
	activityQueryCmd.Flags().StringVar(&activityAncestor, "ancestor-name", "", "Filter by ancestor folder (e.g., folders/123)")
	activityQueryCmd.Flags().StringVar(&activityStartTime, "start-time", "", "Start of time range (RFC3339 format)")
	activityQueryCmd.Flags().StringVar(&activityEndTime, "end-time", "", "End of time range (RFC3339 format)")
	activityQueryCmd.Flags().StringVar(&activityActionTypes, "action-types", "", "Comma-separated action types (edit,comment,share,permission_change,move,delete,restore,create,rename)")
	activityQueryCmd.Flags().StringVar(&activityUser, "user", "", "Filter by user email")
	activityQueryCmd.Flags().IntVar(&activityLimit, "limit", 100, "Maximum results per page")
	activityQueryCmd.Flags().StringVar(&activityPageToken, "page-token", "", "Pagination token")
	activityQueryCmd.Flags().StringVar(&activityFields, "fields", "", "Fields to return")

	activityCmd.AddCommand(activityQueryCmd)
	rootCmd.AddCommand(activityCmd)
}

func getActivityManager(ctx context.Context, flags types.GlobalFlags) (*activity.Manager, *api.Client, *types.RequestContext, *OutputWriter, error) {
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
	mgr := activity.NewManager(client)
	reqCtx := api.NewRequestContext(flags.Profile, flags.DriveID, types.RequestTypeListOrSearch)

	return mgr, client, reqCtx, out, nil
}

func runActivityQuery(cmd *cobra.Command, args []string) error {
	flags := GetGlobalFlags()
	ctx := context.Background()

	mgr, _, reqCtx, out, err := getActivityManager(ctx, flags)
	if err != nil {
		return out.WriteError("activity.query", utils.NewCLIError(utils.ErrCodeAuthRequired, err.Error()).Build())
	}

	opts := types.QueryOptions{
		FileID:       activityFileID,
		FolderID:     activityFolderID,
		AncestorName: activityAncestor,
		StartTime:    activityStartTime,
		EndTime:      activityEndTime,
		ActionTypes:  activityActionTypes,
		User:         activityUser,
		Limit:        activityLimit,
		PageToken:    activityPageToken,
		Fields:       activityFields,
	}

	activities, err := mgr.Query(ctx, reqCtx, opts)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			return out.WriteError("activity.query", appErr.CLIError)
		}
		return out.WriteError("activity.query", utils.NewCLIError(utils.ErrCodeUnknown, err.Error()).Build())
	}

	result := &ActivityQueryResult{Activities: activities}
	return out.WriteSuccess("activity.query", result)
}

type ActivityQueryResult struct {
	Activities []types.Activity
}

func (r *ActivityQueryResult) Headers() []string {
	return []string{"Timestamp", "Action", "Actor", "Target"}
}

func (r *ActivityQueryResult) Rows() [][]string {
	rows := make([][]string, len(r.Activities))
	for i, activity := range r.Activities {
		timestamp := activity.Timestamp.Format("2006-01-02 15:04:05")
		action := activity.PrimaryActionDetail.Type

		actor := "unknown"
		if len(activity.Actors) > 0 {
			if activity.Actors[0].User != nil {
				actor = activity.Actors[0].User.Email
			} else {
				actor = activity.Actors[0].Type
			}
		}

		target := "unknown"
		if len(activity.Targets) > 0 {
			if activity.Targets[0].DriveItem != nil {
				target = activity.Targets[0].DriveItem.Title
				if target == "" {
					target = activity.Targets[0].DriveItem.Name
				}
			} else {
				target = activity.Targets[0].Type
			}
		}

		rows[i] = []string{timestamp, action, actor, target}
	}
	return rows
}

func (r *ActivityQueryResult) EmptyMessage() string {
	return "No activity found"
}
