package sheets

import (
	"testing"

	sheetsapi "google.golang.org/api/sheets/v4"
)

func TestConvertSpreadsheet(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		got := convertSpreadsheet(nil)
		if got.ID != "" || got.Title != "" || got.SheetCount != 0 {
			t.Fatalf("expected empty spreadsheet")
		}
	})

	t.Run("missing properties", func(t *testing.T) {
		got := convertSpreadsheet(&sheetsapi.Spreadsheet{
			SpreadsheetId: "s1",
			Sheets:        []*sheetsapi.Sheet{{}},
		})
		if got.ID != "s1" {
			t.Fatalf("expected id s1")
		}
		if got.SheetCount != 1 {
			t.Fatalf("expected sheet count 1")
		}
		if got.Sheets[0].ID != 0 || got.Sheets[0].Title != "" {
			t.Fatalf("expected zero values for missing properties")
		}
	})

	t.Run("with properties", func(t *testing.T) {
		got := convertSpreadsheet(&sheetsapi.Spreadsheet{
			SpreadsheetId: "s1",
			Properties:    &sheetsapi.SpreadsheetProperties{Title: "Title", Locale: "en", TimeZone: "UTC"},
			Sheets: []*sheetsapi.Sheet{{
				Properties: &sheetsapi.SheetProperties{
					SheetId: 1,
					Title:   "Sheet1",
					Index:   0,
					SheetType: "GRID",
				},
			}},
		})
		if got.Title != "Title" || got.Locale != "en" || got.TimeZone != "UTC" {
			t.Fatalf("expected spreadsheet properties")
		}
		if got.Sheets[0].ID != 1 || got.Sheets[0].Title != "Sheet1" {
			t.Fatalf("expected sheet properties")
		}
	})
}
