package types

import "testing"

func TestDocumentTextRowsTruncate(t *testing.T) {
	text := &DocumentText{
		DocumentID: "abcdefghijklmnopqrstuvwxyz",
		Title:      "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		Text:       "content",
		WordCount:  1,
		CharCount:  7,
	}
	rows := text.Rows()
	if len(rows) != 1 {
		t.Fatalf("expected 1 row")
	}
	if len(rows[0][0]) != 20 || rows[0][0][17:] != "..." {
		t.Fatalf("expected truncated id, got %q", rows[0][0])
	}
	if len(rows[0][1]) != 40 || rows[0][1][37:] != "..." {
		t.Fatalf("expected truncated title, got %q", rows[0][1])
	}
}
