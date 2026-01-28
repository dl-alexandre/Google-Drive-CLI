package labels

import (
	"context"
	"fmt"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/types"
	"github.com/dl-alexandre/gdrv/internal/utils"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/drivelabels/v2"
	"google.golang.org/api/googleapi"
)

type Manager struct {
	client *api.Client
}

func NewManager(client *api.Client) *Manager {
	return &Manager{
		client: client,
	}
}

func (m *Manager) List(ctx context.Context, reqCtx *types.RequestContext, opts types.LabelListOptions) ([]*types.Label, string, error) {
	service, err := drivelabels.NewService(ctx)
	if err != nil {
		return nil, "", utils.NewAppError(utils.NewCLIError(utils.ErrCodeUnknown,
			fmt.Sprintf("Failed to create Drive Labels service: %s", err)).Build())
	}

	call := service.Labels.List()

	if opts.Customer != "" {
		call = call.Customer(opts.Customer)
	}

	if opts.View != "" {
		call = call.View(opts.View)
	}

	if opts.MinimumRole != "" {
		call = call.MinimumRole(opts.MinimumRole)
	}

	if opts.PublishedOnly {
		call = call.PublishedOnly(opts.PublishedOnly)
	}

	if opts.Limit > 0 {
		call = call.PageSize(int64(opts.Limit))
	}

	if opts.PageToken != "" {
		call = call.PageToken(opts.PageToken)
	}

	if opts.Fields != "" {
		call = call.Fields(googleapi.Field(opts.Fields))
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drivelabels.GoogleAppsDriveLabelsV2ListLabelsResponse, error) {
		return call.Do()
	})
	if err != nil {
		return nil, "", err
	}

	labels := make([]*types.Label, 0, len(result.Labels))
	for _, label := range result.Labels {
		labels = append(labels, convertLabel(label))
	}

	return labels, result.NextPageToken, nil
}

func (m *Manager) Get(ctx context.Context, reqCtx *types.RequestContext, labelID string, opts types.LabelGetOptions) (*types.Label, error) {
	service, err := drivelabels.NewService(ctx)
	if err != nil {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeUnknown,
			fmt.Sprintf("Failed to create Drive Labels service: %s", err)).Build())
	}

	call := service.Labels.Get(labelID)

	if opts.View != "" {
		call = call.View(opts.View)
	}

	if opts.UseAdminAccess {
		call = call.UseAdminAccess(opts.UseAdminAccess)
	}

	if opts.Fields != "" {
		call = call.Fields(googleapi.Field(opts.Fields))
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drivelabels.GoogleAppsDriveLabelsV2Label, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return convertLabel(result), nil
}

func (m *Manager) ListFileLabels(ctx context.Context, reqCtx *types.RequestContext, fileID string, opts types.FileLabelListOptions) ([]*types.FileLabel, error) {
	driveService := m.client.Service()

	call := driveService.Files.Get(fileID).SupportsAllDrives(true)

	fields := "labelInfo"
	if opts.Fields != "" {
		fields = opts.Fields
	}
	call = call.Fields(googleapi.Field(fields))

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.File, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	if result.LabelInfo == nil || len(result.LabelInfo.Labels) == 0 {
		return []*types.FileLabel{}, nil
	}

	fileLabels := make([]*types.FileLabel, 0, len(result.LabelInfo.Labels))
	for _, label := range result.LabelInfo.Labels {
		fileLabels = append(fileLabels, convertDriveLabel(label))
	}

	return fileLabels, nil
}

func (m *Manager) ApplyLabel(ctx context.Context, reqCtx *types.RequestContext, fileID string, labelID string, opts types.FileLabelApplyOptions) (*types.FileLabel, error) {
	driveService := m.client.Service()

	modifyRequest := &drive.ModifyLabelsRequest{
		LabelModifications: []*drive.LabelModification{
			{
				LabelId:            labelID,
				FieldModifications: convertFieldModificationsToDrive(opts.Fields),
			},
		},
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.ModifyLabelsResponse, error) {
		return driveService.Files.ModifyLabels(fileID, modifyRequest).Do()
	})
	if err != nil {
		return nil, err
	}

	if len(result.ModifiedLabels) == 0 {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeUnknown, "No labels were applied").Build())
	}

	return convertDriveLabel(result.ModifiedLabels[0]), nil
}

func (m *Manager) UpdateLabel(ctx context.Context, reqCtx *types.RequestContext, fileID string, labelID string, opts types.FileLabelUpdateOptions) (*types.FileLabel, error) {
	driveService := m.client.Service()

	modifyRequest := &drive.ModifyLabelsRequest{
		LabelModifications: []*drive.LabelModification{
			{
				LabelId:            labelID,
				FieldModifications: convertFieldModificationsToDrive(opts.Fields),
			},
		},
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.ModifyLabelsResponse, error) {
		return driveService.Files.ModifyLabels(fileID, modifyRequest).Do()
	})
	if err != nil {
		return nil, err
	}

	if len(result.ModifiedLabels) == 0 {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeUnknown, "No labels were updated").Build())
	}

	return convertDriveLabel(result.ModifiedLabels[0]), nil
}

func (m *Manager) RemoveLabel(ctx context.Context, reqCtx *types.RequestContext, fileID string, labelID string) error {
	driveService := m.client.Service()

	modifyRequest := &drive.ModifyLabelsRequest{
		LabelModifications: []*drive.LabelModification{
			{
				LabelId:     labelID,
				RemoveLabel: true,
			},
		},
	}

	_, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drive.ModifyLabelsResponse, error) {
		return driveService.Files.ModifyLabels(fileID, modifyRequest).Do()
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) CreateLabel(ctx context.Context, reqCtx *types.RequestContext, label *types.Label, opts types.LabelCreateOptions) (*types.Label, error) {
	service, err := drivelabels.NewService(ctx)
	if err != nil {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeUnknown,
			fmt.Sprintf("Failed to create Drive Labels service: %s", err)).Build())
	}

	apiLabel := convertToAPILabel(label)

	call := service.Labels.Create(apiLabel)

	if opts.UseAdminAccess {
		call = call.UseAdminAccess(opts.UseAdminAccess)
	}

	if opts.LanguageCode != "" {
		call = call.LanguageCode(opts.LanguageCode)
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drivelabels.GoogleAppsDriveLabelsV2Label, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return convertLabel(result), nil
}

func (m *Manager) PublishLabel(ctx context.Context, reqCtx *types.RequestContext, labelID string, opts types.LabelPublishOptions) (*types.Label, error) {
	service, err := drivelabels.NewService(ctx)
	if err != nil {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeUnknown,
			fmt.Sprintf("Failed to create Drive Labels service: %s", err)).Build())
	}

	publishRequest := &drivelabels.GoogleAppsDriveLabelsV2PublishLabelRequest{
		UseAdminAccess: opts.UseAdminAccess,
	}

	if opts.WriteControl != nil {
		publishRequest.WriteControl = &drivelabels.GoogleAppsDriveLabelsV2WriteControl{
			RequiredRevisionId: opts.WriteControl.RequiredRevisionID,
		}
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drivelabels.GoogleAppsDriveLabelsV2Label, error) {
		return service.Labels.Publish(labelID, publishRequest).Do()
	})
	if err != nil {
		return nil, err
	}

	return convertLabel(result), nil
}

func (m *Manager) DisableLabel(ctx context.Context, reqCtx *types.RequestContext, labelID string, opts types.LabelDisableOptions) (*types.Label, error) {
	service, err := drivelabels.NewService(ctx)
	if err != nil {
		return nil, utils.NewAppError(utils.NewCLIError(utils.ErrCodeUnknown,
			fmt.Sprintf("Failed to create Drive Labels service: %s", err)).Build())
	}

	disableRequest := &drivelabels.GoogleAppsDriveLabelsV2DisableLabelRequest{
		UseAdminAccess: opts.UseAdminAccess,
	}

	if opts.WriteControl != nil {
		disableRequest.WriteControl = &drivelabels.GoogleAppsDriveLabelsV2WriteControl{
			RequiredRevisionId: opts.WriteControl.RequiredRevisionID,
		}
	}

	if opts.DisabledPolicy != nil {
		disableRequest.DisabledPolicy = &drivelabels.GoogleAppsDriveLabelsV2LifecycleDisabledPolicy{
			HideInSearch: opts.DisabledPolicy.HideInSearch,
			ShowInApply:  opts.DisabledPolicy.ShowInApply,
		}
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*drivelabels.GoogleAppsDriveLabelsV2Label, error) {
		return service.Labels.Disable(labelID, disableRequest).Do()
	})
	if err != nil {
		return nil, err
	}

	return convertLabel(result), nil
}

func convertLabel(apiLabel *drivelabels.GoogleAppsDriveLabelsV2Label) *types.Label {
	if apiLabel == nil {
		return nil
	}

	label := &types.Label{
		ID:                 apiLabel.Id,
		Name:               apiLabel.Name,
		RevisionID:         apiLabel.RevisionId,
		LabelType:          apiLabel.LabelType,
		CreateTime:         apiLabel.CreateTime,
		RevisionCreateTime: apiLabel.RevisionCreateTime,
		PublishTime:        apiLabel.PublishTime,
		DisableTime:        apiLabel.DisableTime,
		Customer:           apiLabel.Customer,
	}

	if apiLabel.Creator != nil {
		label.Creator = &types.LabelUser{
			Person: apiLabel.Creator.Person,
		}
	}

	if apiLabel.RevisionCreator != nil {
		label.RevisionCreator = &types.LabelUser{
			Person: apiLabel.RevisionCreator.Person,
		}
	}

	if apiLabel.Publisher != nil {
		label.Publisher = &types.LabelUser{
			Person: apiLabel.Publisher.Person,
		}
	}

	if apiLabel.Disabler != nil {
		label.Disabler = &types.LabelUser{
			Person: apiLabel.Disabler.Person,
		}
	}

	if apiLabel.Properties != nil {
		label.Properties = &types.LabelProperties{
			Title:       apiLabel.Properties.Title,
			Description: apiLabel.Properties.Description,
		}
	}

	if apiLabel.Lifecycle != nil {
		label.Lifecycle = &types.LabelLifecycle{
			State:                 apiLabel.Lifecycle.State,
			HasUnpublishedChanges: apiLabel.Lifecycle.HasUnpublishedChanges,
		}
		if apiLabel.Lifecycle.DisabledPolicy != nil {
			label.Lifecycle.DisabledPolicy = &types.LabelDisabledPolicy{
				HideInSearch: apiLabel.Lifecycle.DisabledPolicy.HideInSearch,
				ShowInApply:  apiLabel.Lifecycle.DisabledPolicy.ShowInApply,
			}
		}
	}

	if apiLabel.AppliedCapabilities != nil {
		label.AppliedCapabilities = &types.LabelAppliedCapabilities{
			CanRead:   apiLabel.AppliedCapabilities.CanRead,
			CanApply:  apiLabel.AppliedCapabilities.CanApply,
			CanRemove: apiLabel.AppliedCapabilities.CanRemove,
		}
	}

	if apiLabel.SchemaCapabilities != nil {
		label.SchemaCapabilities = &types.LabelSchemaCapabilities{
			CanUpdate:  apiLabel.SchemaCapabilities.CanUpdate,
			CanDelete:  apiLabel.SchemaCapabilities.CanDelete,
			CanDisable: apiLabel.SchemaCapabilities.CanDisable,
			CanEnable:  apiLabel.SchemaCapabilities.CanEnable,
		}
	}

	if apiLabel.AppliedLabelPolicy != nil {
		label.AppliedLabelPolicy = &types.LabelAppliedLabelPolicy{
			CopyMode: apiLabel.AppliedLabelPolicy.CopyMode,
		}
	}

	if len(apiLabel.Fields) > 0 {
		label.Fields = make([]*types.LabelField, 0, len(apiLabel.Fields))
		for _, field := range apiLabel.Fields {
			label.Fields = append(label.Fields, convertLabelField(field))
		}
	}

	return label
}

func convertLabelField(apiField *drivelabels.GoogleAppsDriveLabelsV2Field) *types.LabelField {
	if apiField == nil {
		return nil
	}

	field := &types.LabelField{
		ID:       apiField.Id,
		QueryKey: apiField.QueryKey,
	}

	if apiField.Properties != nil {
		field.Properties = &types.LabelFieldProperties{
			DisplayName:       apiField.Properties.DisplayName,
			Required:          apiField.Properties.Required,
			InsertBeforeField: apiField.Properties.InsertBeforeField,
		}
	}

	if apiField.Lifecycle != nil {
		field.Lifecycle = &types.LabelFieldLifecycle{
			State: apiField.Lifecycle.State,
		}
		if apiField.Lifecycle.DisabledPolicy != nil {
			field.Lifecycle.DisabledPolicy = &types.LabelFieldDisabledPolicy{
				HideInSearch: apiField.Lifecycle.DisabledPolicy.HideInSearch,
				ShowInApply:  apiField.Lifecycle.DisabledPolicy.ShowInApply,
			}
		}
	}

	if apiField.DisplayHints != nil {
		field.DisplayHints = &types.LabelFieldDisplayHints{
			Required:       apiField.DisplayHints.Required,
			Disabled:       apiField.DisplayHints.Disabled,
			HiddenInSearch: apiField.DisplayHints.HiddenInSearch,
			ShownInApply:   apiField.DisplayHints.ShownInApply,
		}
	}

	if apiField.SchemaCapabilities != nil {
		field.SchemaCapabilities = &types.LabelFieldSchemaCapabilities{
			CanUpdate:  apiField.SchemaCapabilities.CanUpdate,
			CanDelete:  apiField.SchemaCapabilities.CanDelete,
			CanDisable: apiField.SchemaCapabilities.CanDisable,
			CanEnable:  apiField.SchemaCapabilities.CanEnable,
		}
	}

	if apiField.AppliedCapabilities != nil {
		field.AppliedCapabilities = &types.LabelFieldAppliedCapabilities{
			CanRead:   apiField.AppliedCapabilities.CanRead,
			CanSearch: apiField.AppliedCapabilities.CanSearch,
			CanWrite:  apiField.AppliedCapabilities.CanWrite,
		}
	}

	if apiField.TextOptions != nil {
		field.TextOptions = &types.LabelFieldTextOptions{
			MinLength: int(apiField.TextOptions.MinLength),
			MaxLength: int(apiField.TextOptions.MaxLength),
		}
	}

	if apiField.IntegerOptions != nil {
		field.IntegerOptions = &types.LabelFieldIntegerOptions{
			MinValue: apiField.IntegerOptions.MinValue,
			MaxValue: apiField.IntegerOptions.MaxValue,
		}
	}

	if apiField.DateOptions != nil {
		field.DateOptions = &types.LabelFieldDateOptions{
			DateFormatType: apiField.DateOptions.DateFormatType,
			DateFormat:     apiField.DateOptions.DateFormat,
		}
		if apiField.DateOptions.MinValue != nil {
			field.DateOptions.MinValue = &types.LabelFieldDateValue{
				Year:  int(apiField.DateOptions.MinValue.Year),
				Month: int(apiField.DateOptions.MinValue.Month),
				Day:   int(apiField.DateOptions.MinValue.Day),
			}
		}
		if apiField.DateOptions.MaxValue != nil {
			field.DateOptions.MaxValue = &types.LabelFieldDateValue{
				Year:  int(apiField.DateOptions.MaxValue.Year),
				Month: int(apiField.DateOptions.MaxValue.Month),
				Day:   int(apiField.DateOptions.MaxValue.Day),
			}
		}
	}

	if apiField.SelectionOptions != nil {
		field.SelectionOptions = &types.LabelFieldSelectionOptions{}
		if apiField.SelectionOptions.ListOptions != nil {
			field.SelectionOptions.ListOptions = &types.LabelFieldListOptions{
				MaxEntries: int(apiField.SelectionOptions.ListOptions.MaxEntries),
			}
		}
		if len(apiField.SelectionOptions.Choices) > 0 {
			field.SelectionOptions.Choices = make([]*types.LabelFieldChoice, 0, len(apiField.SelectionOptions.Choices))
			for _, choice := range apiField.SelectionOptions.Choices {
				field.SelectionOptions.Choices = append(field.SelectionOptions.Choices, convertLabelFieldChoice(choice))
			}
		}
	}

	if apiField.UserOptions != nil {
		field.UserOptions = &types.LabelFieldUserOptions{}
		if apiField.UserOptions.ListOptions != nil {
			field.UserOptions.ListOptions = &types.LabelFieldListOptions{
				MaxEntries: int(apiField.UserOptions.ListOptions.MaxEntries),
			}
		}
	}

	return field
}

func convertLabelFieldChoice(apiChoice *drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoice) *types.LabelFieldChoice {
	if apiChoice == nil {
		return nil
	}

	choice := &types.LabelFieldChoice{
		ID: apiChoice.Id,
	}

	if apiChoice.Properties != nil {
		choice.Properties = &types.LabelFieldChoiceProperties{
			DisplayName:        apiChoice.Properties.DisplayName,
			Description:        apiChoice.Properties.Description,
			InsertBeforeChoice: apiChoice.Properties.InsertBeforeChoice,
		}
		if apiChoice.Properties.BadgeConfig != nil {
			choice.Properties.BadgeConfig = &types.LabelFieldBadgeConfig{
				PriorityOverride: apiChoice.Properties.BadgeConfig.PriorityOverride,
			}
			if apiChoice.Properties.BadgeConfig.Color != nil {
				choice.Properties.BadgeConfig.Color = &types.LabelFieldBadgeColor{
					Red:   apiChoice.Properties.BadgeConfig.Color.Red,
					Green: apiChoice.Properties.BadgeConfig.Color.Green,
					Blue:  apiChoice.Properties.BadgeConfig.Color.Blue,
					Alpha: apiChoice.Properties.BadgeConfig.Color.Alpha,
				}
			}
		}
	}

	if apiChoice.Lifecycle != nil {
		choice.Lifecycle = &types.LabelFieldChoiceLifecycle{
			State: apiChoice.Lifecycle.State,
		}
		if apiChoice.Lifecycle.DisabledPolicy != nil {
			choice.Lifecycle.DisabledPolicy = &types.LabelFieldChoiceDisabledPolicy{
				HideInSearch: apiChoice.Lifecycle.DisabledPolicy.HideInSearch,
				ShowInApply:  apiChoice.Lifecycle.DisabledPolicy.ShowInApply,
			}
		}
	}

	if apiChoice.DisplayHints != nil {
		choice.DisplayHints = &types.LabelFieldChoiceDisplayHints{
			Disabled:       apiChoice.DisplayHints.Disabled,
			HiddenInSearch: apiChoice.DisplayHints.HiddenInSearch,
			ShownInApply:   apiChoice.DisplayHints.ShownInApply,
			BadgePriority:  apiChoice.DisplayHints.BadgePriority,
		}
	}

	if apiChoice.SchemaCapabilities != nil {
		choice.SchemaCapabilities = &types.LabelFieldChoiceSchemaCapabilities{
			CanUpdate:  apiChoice.SchemaCapabilities.CanUpdate,
			CanDelete:  apiChoice.SchemaCapabilities.CanDelete,
			CanDisable: apiChoice.SchemaCapabilities.CanDisable,
			CanEnable:  apiChoice.SchemaCapabilities.CanEnable,
		}
	}

	if apiChoice.AppliedCapabilities != nil {
		choice.AppliedCapabilities = &types.LabelFieldChoiceAppliedCapabilities{
			CanRead:   apiChoice.AppliedCapabilities.CanRead,
			CanSearch: apiChoice.AppliedCapabilities.CanSearch,
			CanSelect: apiChoice.AppliedCapabilities.CanSelect,
		}
	}

	return choice
}

func convertDriveLabel(driveLabel *drive.Label) *types.FileLabel {
	if driveLabel == nil {
		return nil
	}

	fileLabel := &types.FileLabel{
		ID:         driveLabel.Id,
		RevisionID: driveLabel.RevisionId,
		Fields:     make(map[string]*types.LabelFieldValue),
	}

	if driveLabel.Fields != nil {
		for fieldID, fieldValue := range driveLabel.Fields {
			fileLabel.Fields[fieldID] = convertDriveLabelFieldValue(&fieldValue)
		}
	}

	return fileLabel
}

func convertDriveLabelFieldValue(driveFieldValue *drive.LabelField) *types.LabelFieldValue {
	if driveFieldValue == nil {
		return nil
	}

	fieldValue := &types.LabelFieldValue{}

	if driveFieldValue.Text != nil {
		fieldValue.ValueType = "text"
		fieldValue.Text = driveFieldValue.Text
	} else if driveFieldValue.Integer != nil {
		fieldValue.ValueType = "integer"
		fieldValue.Integer = driveFieldValue.Integer
	} else if driveFieldValue.DateString != nil {
		fieldValue.ValueType = "date"
		dates := make([]*types.LabelFieldDateValue, 0, len(driveFieldValue.DateString))
		for _, dateStr := range driveFieldValue.DateString {
			dates = append(dates, parseDateString(dateStr))
		}
		fieldValue.Date = dates
	} else if driveFieldValue.Selection != nil {
		fieldValue.ValueType = "selection"
		fieldValue.Selection = driveFieldValue.Selection
	} else if driveFieldValue.User != nil {
		fieldValue.ValueType = "user"
		users := make([]*types.LabelUser, 0, len(driveFieldValue.User))
		for _, user := range driveFieldValue.User {
			users = append(users, &types.LabelUser{
				Email:       user.EmailAddress,
				DisplayName: user.DisplayName,
			})
		}
		fieldValue.User = users
	}

	return fieldValue
}

func parseDateString(dateStr string) *types.LabelFieldDateValue {
	return &types.LabelFieldDateValue{}
}

func convertFieldModificationsToDrive(fields map[string]*types.LabelFieldValue) []*drive.LabelFieldModification {
	if len(fields) == 0 {
		return nil
	}

	modifications := make([]*drive.LabelFieldModification, 0, len(fields))
	for fieldID, fieldValue := range fields {
		modification := &drive.LabelFieldModification{
			FieldId: fieldID,
		}

		if fieldValue != nil {
			switch fieldValue.ValueType {
			case "text":
				if len(fieldValue.Text) > 0 {
					modification.SetTextValues = fieldValue.Text
				}
			case "integer":
				if len(fieldValue.Integer) > 0 {
					modification.SetIntegerValues = fieldValue.Integer
				}
			case "date":
				if len(fieldValue.Date) > 0 {
					dates := make([]string, 0, len(fieldValue.Date))
					for _, date := range fieldValue.Date {
						dates = append(dates, fmt.Sprintf("%04d-%02d-%02d", date.Year, date.Month, date.Day))
					}
					modification.SetDateValues = dates
				}
			case "selection":
				if len(fieldValue.Selection) > 0 {
					modification.SetSelectionValues = fieldValue.Selection
				}
			case "user":
				if len(fieldValue.User) > 0 {
					users := make([]string, 0, len(fieldValue.User))
					for _, user := range fieldValue.User {
						users = append(users, user.Person)
					}
					modification.SetUserValues = users
				}
			}
		}

		modifications = append(modifications, modification)
	}

	return modifications
}

func convertToAPILabel(label *types.Label) *drivelabels.GoogleAppsDriveLabelsV2Label {
	if label == nil {
		return nil
	}

	apiLabel := &drivelabels.GoogleAppsDriveLabelsV2Label{
		LabelType: label.LabelType,
	}

	if label.Properties != nil {
		apiLabel.Properties = &drivelabels.GoogleAppsDriveLabelsV2LabelProperties{
			Title:       label.Properties.Title,
			Description: label.Properties.Description,
		}
	}

	if len(label.Fields) > 0 {
		apiLabel.Fields = make([]*drivelabels.GoogleAppsDriveLabelsV2Field, 0, len(label.Fields))
		for _, field := range label.Fields {
			apiLabel.Fields = append(apiLabel.Fields, convertToAPIField(field))
		}
	}

	return apiLabel
}

func convertToAPIField(field *types.LabelField) *drivelabels.GoogleAppsDriveLabelsV2Field {
	if field == nil {
		return nil
	}

	apiField := &drivelabels.GoogleAppsDriveLabelsV2Field{}

	if field.Properties != nil {
		apiField.Properties = &drivelabels.GoogleAppsDriveLabelsV2FieldProperties{
			DisplayName: field.Properties.DisplayName,
			Required:    field.Properties.Required,
		}
	}

	if field.TextOptions != nil {
		apiField.TextOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldTextOptions{
			MinLength: int64(field.TextOptions.MinLength),
			MaxLength: int64(field.TextOptions.MaxLength),
		}
	}

	if field.IntegerOptions != nil {
		apiField.IntegerOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldIntegerOptions{
			MinValue: field.IntegerOptions.MinValue,
			MaxValue: field.IntegerOptions.MaxValue,
		}
	}

	if field.DateOptions != nil {
		apiField.DateOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldDateOptions{
			DateFormatType: field.DateOptions.DateFormatType,
			DateFormat:     field.DateOptions.DateFormat,
		}
	}

	if field.SelectionOptions != nil {
		apiField.SelectionOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptions{}
		if len(field.SelectionOptions.Choices) > 0 {
			apiField.SelectionOptions.Choices = make([]*drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoice, 0, len(field.SelectionOptions.Choices))
			for _, choice := range field.SelectionOptions.Choices {
				apiField.SelectionOptions.Choices = append(apiField.SelectionOptions.Choices, convertToAPIChoice(choice))
			}
		}
	}

	if field.UserOptions != nil {
		apiField.UserOptions = &drivelabels.GoogleAppsDriveLabelsV2FieldUserOptions{}
	}

	return apiField
}

func convertToAPIChoice(choice *types.LabelFieldChoice) *drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoice {
	if choice == nil {
		return nil
	}

	apiChoice := &drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoice{}

	if choice.Properties != nil {
		apiChoice.Properties = &drivelabels.GoogleAppsDriveLabelsV2FieldSelectionOptionsChoiceProperties{
			DisplayName: choice.Properties.DisplayName,
			Description: choice.Properties.Description,
		}
		if choice.Properties.BadgeConfig != nil {
			apiChoice.Properties.BadgeConfig = &drivelabels.GoogleAppsDriveLabelsV2BadgeConfig{
				PriorityOverride: choice.Properties.BadgeConfig.PriorityOverride,
			}
			if choice.Properties.BadgeConfig.Color != nil {
				apiChoice.Properties.BadgeConfig.Color = &drivelabels.GoogleTypeColor{
					Red:   choice.Properties.BadgeConfig.Color.Red,
					Green: choice.Properties.BadgeConfig.Color.Green,
					Blue:  choice.Properties.BadgeConfig.Color.Blue,
					Alpha: choice.Properties.BadgeConfig.Color.Alpha,
				}
			}
		}
	}

	return apiChoice
}
