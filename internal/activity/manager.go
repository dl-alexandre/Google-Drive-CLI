package activity

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
	"google.golang.org/api/driveactivity/v2"
)

type Manager struct {
	client *api.Client
}

func NewManager(client *api.Client) *Manager {
	return &Manager{
		client: client,
	}
}

func (m *Manager) Query(ctx context.Context, reqCtx *types.RequestContext, opts types.QueryOptions) ([]types.Activity, error) {
	service, err := driveactivity.NewService(ctx)
	if err != nil {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeUnknown,
			fmt.Sprintf("Failed to create Drive Activity service: %s", err)).Build())
	}

	req := &driveactivity.QueryDriveActivityRequest{}

	if opts.FileID != "" {
		req.ItemName = fmt.Sprintf("items/%s", opts.FileID)
		reqCtx.InvolvedFileIDs = append(reqCtx.InvolvedFileIDs, opts.FileID)
	}

	if opts.FolderID != "" {
		req.ItemName = fmt.Sprintf("items/%s", opts.FolderID)
		reqCtx.InvolvedParentIDs = append(reqCtx.InvolvedParentIDs, opts.FolderID)
	}

	if opts.AncestorName != "" {
		req.AncestorName = opts.AncestorName
	}

	if opts.StartTime != "" || opts.EndTime != "" {
		req.Filter = buildTimeFilter(opts.StartTime, opts.EndTime)
	}

	if opts.ActionTypes != "" {
		actionFilter := buildActionFilter(opts.ActionTypes)
		if req.Filter != "" {
			req.Filter = fmt.Sprintf("(%s) AND (%s)", req.Filter, actionFilter)
		} else {
			req.Filter = actionFilter
		}
	}

	if opts.User != "" {
		userFilter := fmt.Sprintf("actor.user.knownUser.personName = 'people/%s'", opts.User)
		if req.Filter != "" {
			req.Filter = fmt.Sprintf("(%s) AND (%s)", req.Filter, userFilter)
		} else {
			req.Filter = userFilter
		}
	}

	if opts.Limit > 0 {
		req.PageSize = int64(opts.Limit)
	}

	if opts.PageToken != "" {
		req.PageToken = opts.PageToken
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*driveactivity.QueryDriveActivityResponse, error) {
		return service.Activity.Query(req).Do()
	})
	if err != nil {
		return nil, err
	}

	activities := make([]types.Activity, 0, len(result.Activities))
	for _, act := range result.Activities {
		activities = append(activities, convertActivity(act))
	}

	return activities, nil
}

func buildTimeFilter(startTime, endTime string) string {
	var filters []string

	if startTime != "" {
		filters = append(filters, fmt.Sprintf("time >= '%s'", startTime))
	}

	if endTime != "" {
		filters = append(filters, fmt.Sprintf("time <= '%s'", endTime))
	}

	if len(filters) == 0 {
		return ""
	}

	return strings.Join(filters, " AND ")
}

func buildActionFilter(actionTypes string) string {
	types := strings.Split(actionTypes, ",")
	var filters []string

	for _, t := range types {
		t = strings.TrimSpace(t)
		switch t {
		case "edit":
			filters = append(filters, "detail.action_detail_case:EDIT")
		case "comment":
			filters = append(filters, "detail.action_detail_case:COMMENT")
		case "share":
			filters = append(filters, "detail.action_detail_case:PERMISSION_CHANGE")
		case "permission_change":
			filters = append(filters, "detail.action_detail_case:PERMISSION_CHANGE")
		case "move":
			filters = append(filters, "detail.action_detail_case:MOVE")
		case "delete":
			filters = append(filters, "detail.action_detail_case:DELETE")
		case "restore":
			filters = append(filters, "detail.action_detail_case:RESTORE")
		case "create":
			filters = append(filters, "detail.action_detail_case:CREATE")
		case "rename":
			filters = append(filters, "detail.action_detail_case:RENAME")
		}
	}

	if len(filters) == 0 {
		return ""
	}

	return strings.Join(filters, " OR ")
}

func convertActivity(act *driveactivity.DriveActivity) types.Activity {
	activity := types.Activity{
		Actors:  convertActors(act.Actors),
		Targets: convertTargets(act.Targets),
		Actions: convertActions(act.Actions),
	}

	if act.Timestamp != "" {
		activity.Timestamp, _ = parseTimestamp(act.Timestamp)
	}

	if act.PrimaryActionDetail != nil {
		activity.PrimaryActionDetail = convertActionDetail(act.PrimaryActionDetail)
	}

	return activity
}

func convertActors(actors []*driveactivity.Actor) []types.Actor {
	result := make([]types.Actor, 0, len(actors))
	for _, actor := range actors {
		a := types.Actor{}

		if actor.User != nil {
			a.Type = "user"
			if actor.User.KnownUser != nil {
				a.User = &types.ActivityUser{
					Email: actor.User.KnownUser.PersonName,
				}
			}
		} else if actor.Administrator != nil {
			a.Type = "administrator"
		} else if actor.System != nil {
			a.Type = "system"
		} else if actor.Anonymous != nil {
			a.Type = "anonymous"
		}

		result = append(result, a)
	}
	return result
}

func convertTargets(targets []*driveactivity.Target) []types.Target {
	result := make([]types.Target, 0, len(targets))
	for _, target := range targets {
		t := types.Target{}

		if target.DriveItem != nil {
			t.Type = "driveItem"
			t.DriveItem = &types.DriveItem{
				Name:     target.DriveItem.Name,
				Title:    target.DriveItem.Title,
				MimeType: target.DriveItem.MimeType,
			}
			if target.DriveItem.Owner != nil && target.DriveItem.Owner.User != nil {
				if target.DriveItem.Owner.User.KnownUser != nil {
					t.DriveItem.Owner = &types.ActivityUser{
						Email: target.DriveItem.Owner.User.KnownUser.PersonName,
					}
				}
			}
		} else if target.Drive != nil {
			t.Type = "drive"
		} else if target.FileComment != nil {
			t.Type = "fileComment"
		}

		result = append(result, t)
	}
	return result
}

func convertActions(actions []*driveactivity.Action) []types.Action {
	result := make([]types.Action, 0, len(actions))
	for _, action := range actions {
		a := types.Action{
			Detail: convertActionDetail(action.Detail),
		}

		if action.Detail != nil {
			a.Type = getActionType(action.Detail)
		}

		result = append(result, a)
	}
	return result
}

func convertActionDetail(detail *driveactivity.ActionDetail) types.ActionDetail {
	ad := types.ActionDetail{}

	if detail.Edit != nil {
		ad.Type = "edit"
		ad.Description = "File edited"
	} else if detail.Comment != nil {
		ad.Type = "comment"
		ad.Description = "Comment added"
	} else if detail.PermissionChange != nil {
		ad.Type = "permission_change"
		ad.Description = "Permissions changed"
	} else if detail.Move != nil {
		ad.Type = "move"
		ad.Description = "Item moved"
	} else if detail.Delete != nil {
		ad.Type = "delete"
		ad.Description = "Item deleted"
	} else if detail.Restore != nil {
		ad.Type = "restore"
		ad.Description = "Item restored"
	} else if detail.Create != nil {
		ad.Type = "create"
		ad.Description = "Item created"
	} else if detail.Rename != nil {
		ad.Type = "rename"
		ad.Description = "Item renamed"
	}

	return ad
}

func getActionType(detail *driveactivity.ActionDetail) string {
	if detail.Edit != nil {
		return "edit"
	} else if detail.Comment != nil {
		return "comment"
	} else if detail.PermissionChange != nil {
		return "permission_change"
	} else if detail.Move != nil {
		return "move"
	} else if detail.Delete != nil {
		return "delete"
	} else if detail.Restore != nil {
		return "restore"
	} else if detail.Create != nil {
		return "create"
	} else if detail.Rename != nil {
		return "rename"
	}
	return "unknown"
}

func parseTimestamp(ts string) (time.Time, error) {
	return time.Parse(time.RFC3339, ts)
}
