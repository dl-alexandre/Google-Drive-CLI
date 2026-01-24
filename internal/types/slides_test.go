package types

import "testing"

func TestPresentationTextRowsTruncate(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	pres := &PresentationText{
		PresentationID: "p1",
		Title:          "Title",
		SlideCount:     1,
		TextBySlide: []SlideText{{
			SlideIndex: 1,
			ObjectID:   "obj",
			Text:       text,
		}},
	}
	rows := pres.Rows()
	if len(rows) != 1 {
		t.Fatalf("expected 1 row")
	}
	if len(rows[0][2]) != 50 || rows[0][2][47:] != "..." {
		t.Fatalf("expected truncated text, got %q", rows[0][2])
	}
}
