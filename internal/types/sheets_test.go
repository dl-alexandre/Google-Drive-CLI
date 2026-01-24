package types

import "testing"

func TestColumnLetter(t *testing.T) {
	cases := map[int]string{
		0:  "A",
		25: "Z",
		26: "AA",
		27: "AB",
		51: "AZ",
		52: "BA",
	}
	for input, expected := range cases {
		if got := columnLetter(input); got != expected {
			t.Fatalf("columnLetter(%d) = %q, want %q", input, got, expected)
		}
	}
}

func TestSheetValuesHeadersAndRows(t *testing.T) {
	values := &SheetValues{
		Range: "Sheet1!A1:B2",
		Values: [][]interface{}{
			{1, "a"},
			{2, "b"},
		},
	}
	headers := values.Headers()
	if len(headers) != 2 || headers[0] != "A" || headers[1] != "B" {
		t.Fatalf("unexpected headers: %#v", headers)
	}
	rows := values.Rows()
	if len(rows) != 2 || rows[0][0] != "1" || rows[0][1] != "a" {
		t.Fatalf("unexpected rows: %#v", rows)
	}
}

func TestUpdateValuesResponseRowsTruncate(t *testing.T) {
	resp := &UpdateValuesResponse{
		SpreadsheetID: "abcdefghijklmnopqrstuvwxyz",
		UpdatedRange:  "Sheet1!A1",
		UpdatedRows:   1,
		UpdatedColumns: 1,
		UpdatedCells:  1,
	}
	rows := resp.Rows()
	if len(rows) != 1 {
		t.Fatalf("expected 1 row")
	}
	if len(rows[0][0]) != 20 || rows[0][0][17:] != "..." {
		t.Fatalf("expected truncated id, got %q", rows[0][0])
	}
}
