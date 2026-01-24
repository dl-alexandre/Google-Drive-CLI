package sheets

import (
	"context"

	"github.com/dl-alexandre/gdrive/internal/api"
	"github.com/dl-alexandre/gdrive/internal/types"
	"google.golang.org/api/sheets/v4"
)

type Manager struct {
	client  *api.Client
	service *sheets.Service
}

func NewManager(client *api.Client, service *sheets.Service) *Manager {
	return &Manager{
		client:  client,
		service: service,
	}
}

func (m *Manager) GetValues(ctx context.Context, reqCtx *types.RequestContext, spreadsheetID, rangeNotation string) (*types.SheetValues, error) {
	call := m.service.Spreadsheets.Values.Get(spreadsheetID, rangeNotation)

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*sheets.ValueRange, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return &types.SheetValues{
		SpreadsheetID:  spreadsheetID,
		Range:          result.Range,
		MajorDimension: result.MajorDimension,
		Values:         result.Values,
	}, nil
}

func (m *Manager) UpdateValues(ctx context.Context, reqCtx *types.RequestContext, spreadsheetID, rangeNotation string, values [][]interface{}, valueInputOption string) (*types.UpdateValuesResponse, error) {
	valueRange := &sheets.ValueRange{
		Values:         values,
		MajorDimension: "ROWS",
	}

	call := m.service.Spreadsheets.Values.Update(spreadsheetID, rangeNotation, valueRange)
	if valueInputOption != "" {
		call = call.ValueInputOption(valueInputOption)
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*sheets.UpdateValuesResponse, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateValuesResponse{
		SpreadsheetID:  result.SpreadsheetId,
		UpdatedRange:   result.UpdatedRange,
		UpdatedRows:    int(result.UpdatedRows),
		UpdatedColumns: int(result.UpdatedColumns),
		UpdatedCells:   int(result.UpdatedCells),
	}, nil
}

func (m *Manager) AppendValues(ctx context.Context, reqCtx *types.RequestContext, spreadsheetID, rangeNotation string, values [][]interface{}, valueInputOption string) (*types.UpdateValuesResponse, error) {
	valueRange := &sheets.ValueRange{
		Values:         values,
		MajorDimension: "ROWS",
	}

	call := m.service.Spreadsheets.Values.Append(spreadsheetID, rangeNotation, valueRange)
	if valueInputOption != "" {
		call = call.ValueInputOption(valueInputOption)
	}

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*sheets.AppendValuesResponse, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	if result.Updates == nil {
		return &types.UpdateValuesResponse{
			SpreadsheetID: spreadsheetID,
			UpdatedRange:  "",
		}, nil
	}

	return &types.UpdateValuesResponse{
		SpreadsheetID:  result.Updates.SpreadsheetId,
		UpdatedRange:   result.Updates.UpdatedRange,
		UpdatedRows:    int(result.Updates.UpdatedRows),
		UpdatedColumns: int(result.Updates.UpdatedColumns),
		UpdatedCells:   int(result.Updates.UpdatedCells),
	}, nil
}

func (m *Manager) ClearValues(ctx context.Context, reqCtx *types.RequestContext, spreadsheetID, rangeNotation string) (*types.ClearValuesResponse, error) {
	call := m.service.Spreadsheets.Values.Clear(spreadsheetID, rangeNotation, &sheets.ClearValuesRequest{})
	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*sheets.ClearValuesResponse, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}
	return &types.ClearValuesResponse{
		SpreadsheetID: result.SpreadsheetId,
		ClearedRange:  result.ClearedRange,
	}, nil
}

func (m *Manager) BatchUpdate(ctx context.Context, reqCtx *types.RequestContext, spreadsheetID string, requests []*sheets.Request) (*types.SheetsBatchUpdateResponse, error) {
	call := m.service.Spreadsheets.BatchUpdate(spreadsheetID, &sheets.BatchUpdateSpreadsheetRequest{
		Requests: requests,
	})
	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*sheets.BatchUpdateSpreadsheetResponse, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}
	return &types.SheetsBatchUpdateResponse{
		SpreadsheetID: result.SpreadsheetId,
		RepliesCount:  len(result.Replies),
	}, nil
}

func (m *Manager) GetSpreadsheet(ctx context.Context, reqCtx *types.RequestContext, spreadsheetID string) (*types.Spreadsheet, error) {
	call := m.service.Spreadsheets.Get(spreadsheetID)

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*sheets.Spreadsheet, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}

	return convertSpreadsheet(result), nil
}

func convertSpreadsheet(spreadsheet *sheets.Spreadsheet) *types.Spreadsheet {
	if spreadsheet == nil {
		return &types.Spreadsheet{}
	}
	title := ""
	locale := ""
	timeZone := ""
	if spreadsheet.Properties != nil {
		title = spreadsheet.Properties.Title
		locale = spreadsheet.Properties.Locale
		timeZone = spreadsheet.Properties.TimeZone
	}
	result := &types.Spreadsheet{
		ID:       spreadsheet.SpreadsheetId,
		Title:    title,
		Locale:   locale,
		TimeZone: timeZone,
	}

	if spreadsheet.Sheets != nil {
		result.SheetCount = len(spreadsheet.Sheets)
		result.Sheets = make([]types.Sheet, len(spreadsheet.Sheets))
		for i, sheet := range spreadsheet.Sheets {
			sheetType := ""
			sheetID := int64(0)
			sheetTitle := ""
			sheetIndex := int64(0)
			if sheet.Properties != nil {
				sheetType = sheet.Properties.SheetType
				sheetID = sheet.Properties.SheetId
				sheetTitle = sheet.Properties.Title
				sheetIndex = sheet.Properties.Index
			}
			result.Sheets[i] = types.Sheet{
				ID:    sheetID,
				Title: sheetTitle,
				Index: sheetIndex,
				Type:  sheetType,
			}
		}
	}

	return result
}
