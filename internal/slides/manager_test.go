package slides

import (
	"testing"

	slidesapi "google.golang.org/api/slides/v1"
)

func TestConvertPresentation(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		if got := convertPresentation(nil); got.PresentationID != "" || got.Title != "" || got.SlideCount != 0 {
			t.Fatalf("expected empty presentation")
		}
	})

	t.Run("with slides", func(t *testing.T) {
		pres := &slidesapi.Presentation{
			PresentationId: "p1",
			Title:          "Title",
			Slides: []*slidesapi.Page{
				{ObjectId: "s1"},
				{ObjectId: "s2"},
			},
		}
		got := convertPresentation(pres)
		if got.SlideCount != 2 {
			t.Fatalf("expected 2 slides, got %d", got.SlideCount)
		}
		if got.Slides[0].ObjectID != "s1" {
			t.Fatalf("expected slide id s1")
		}
	})
}

func TestExtractTextFromPresentation(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		got := extractTextFromPresentation(nil)
		if got.PresentationID != "" || got.Title != "" || got.SlideCount != 0 {
			t.Fatalf("expected empty presentation text")
		}
	})

	t.Run("text shapes", func(t *testing.T) {
		pres := &slidesapi.Presentation{
			PresentationId: "p1",
			Title:          "Title",
			Slides: []*slidesapi.Page{{
				ObjectId: "s1",
				PageElements: []*slidesapi.PageElement{{
					ObjectId: "e1",
					Shape: &slidesapi.Shape{
						Text: &slidesapi.TextContent{
							TextElements: []*slidesapi.TextElement{
								{TextRun: &slidesapi.TextRun{Content: "Hello"}},
								{TextRun: &slidesapi.TextRun{Content: " World"}},
							},
						},
					},
				}},
			}},
		}
		got := extractTextFromPresentation(pres)
		if len(got.TextBySlide) != 1 {
			t.Fatalf("expected one text element")
		}
		if got.TextBySlide[0].Text != "Hello World" {
			t.Fatalf("unexpected text: %q", got.TextBySlide[0].Text)
		}
	})
}

func TestExtractTextFromShape(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		if got := extractTextFromShape(nil); got != "" {
			t.Fatalf("expected empty string")
		}
	})

	t.Run("text runs", func(t *testing.T) {
		content := &slidesapi.TextContent{
			TextElements: []*slidesapi.TextElement{
				{TextRun: &slidesapi.TextRun{Content: "A"}},
				{TextRun: &slidesapi.TextRun{Content: "B"}},
			},
		}
		if got := extractTextFromShape(content); got != "AB" {
			t.Fatalf("expected AB, got %q", got)
		}
	})
}
