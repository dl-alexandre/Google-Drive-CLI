package types

import "fmt"

type SheetValues struct {
	SpreadsheetID  string          `json:"spreadsheetId"`
	Range          string          `json:"range"`
	MajorDimension string          `json:"majorDimension"`
	Values         [][]interface{} `json:"values"`
}

func (v *SheetValues) Headers() []string {
	if len(v.Values) == 0 {
		return []string{}
	}
	columnCount := len(v.Values[0])
	if columnCount == 0 {
		return []string{}
	}
	headers := make([]string, columnCount)
	for i := range headers {
		headers[i] = columnLetter(i)
	}
	return headers
}

func (v *SheetValues) Rows() [][]string {
	rows := make([][]string, len(v.Values))
	for i, row := range v.Values {
		rows[i] = make([]string, len(row))
		for j, cell := range row {
			if cell == nil {
				rows[i][j] = ""
			} else {
				rows[i][j] = fmt.Sprintf("%v", cell)
			}
		}
	}
	return rows
}

func (v *SheetValues) EmptyMessage() string {
	if v.Range == "" {
		return "No values found"
	}
	return fmt.Sprintf("No values found in range %s", v.Range)
}

type UpdateValuesResponse struct {
	SpreadsheetID  string `json:"spreadsheetId"`
	UpdatedRange   string `json:"updatedRange"`
	UpdatedRows    int    `json:"updatedRows"`
	UpdatedColumns int    `json:"updatedColumns"`
	UpdatedCells   int    `json:"updatedCells"`
}

func (r *UpdateValuesResponse) Headers() []string {
	return []string{"Spreadsheet ID", "Range", "Rows", "Columns", "Cells"}
}

func (r *UpdateValuesResponse) Rows() [][]string {
	return [][]string{{
		truncateID(r.SpreadsheetID, 20),
		r.UpdatedRange,
		fmt.Sprintf("%d", r.UpdatedRows),
		fmt.Sprintf("%d", r.UpdatedColumns),
		fmt.Sprintf("%d", r.UpdatedCells),
	}}
}

func (r *UpdateValuesResponse) EmptyMessage() string {
	return "No update information available"
}

type ClearValuesResponse struct {
	SpreadsheetID string `json:"spreadsheetId"`
	ClearedRange  string `json:"clearedRange"`
}

func (r *ClearValuesResponse) Headers() []string {
	return []string{"Spreadsheet ID", "Cleared Range"}
}

func (r *ClearValuesResponse) Rows() [][]string {
	return [][]string{{
		truncateID(r.SpreadsheetID, 20),
		r.ClearedRange,
	}}
}

func (r *ClearValuesResponse) EmptyMessage() string {
	return "No clear information available"
}

type SheetsBatchUpdateResponse struct {
	SpreadsheetID string `json:"spreadsheetId"`
	RepliesCount  int    `json:"repliesCount"`
}

func (r *SheetsBatchUpdateResponse) Headers() []string {
	return []string{"Spreadsheet ID", "Replies"}
}

func (r *SheetsBatchUpdateResponse) Rows() [][]string {
	return [][]string{{
		truncateID(r.SpreadsheetID, 20),
		fmt.Sprintf("%d", r.RepliesCount),
	}}
}

func (r *SheetsBatchUpdateResponse) EmptyMessage() string {
	return "No update information available"
}

type Spreadsheet struct {
	ID         string  `json:"id"`
	Title      string  `json:"title"`
	Locale     string  `json:"locale,omitempty"`
	TimeZone   string  `json:"timeZone,omitempty"`
	SheetCount int     `json:"sheetCount"`
	Sheets     []Sheet `json:"sheets,omitempty"`
}

func (s *Spreadsheet) Headers() []string {
	return []string{"Sheet ID", "Title", "Index", "Type"}
}

func (s *Spreadsheet) Rows() [][]string {
	rows := make([][]string, len(s.Sheets))
	for i, sheet := range s.Sheets {
		rows[i] = []string{
			fmt.Sprintf("%d", sheet.ID),
			sheet.Title,
			fmt.Sprintf("%d", sheet.Index),
			sheet.Type,
		}
	}
	return rows
}

func (s *Spreadsheet) EmptyMessage() string {
	if s.Title == "" {
		return "No sheets found"
	}
	return fmt.Sprintf("Spreadsheet '%s' has no sheets", s.Title)
}

type Sheet struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Index int64  `json:"index"`
	Type  string `json:"type,omitempty"`
}

func columnLetter(col int) string {
	result := ""
	for col >= 0 {
		result = string(rune('A'+(col%26))) + result
		col = col/26 - 1
	}
	return result
}

func truncateID(id string, maxLen int) string {
	if len(id) <= maxLen {
		return id
	}
	if maxLen <= 3 {
		return id[:maxLen]
	}
	return id[:maxLen-3] + "..."
}
