package docs

import (
	"strings"
	"testing"

	docsapi "google.golang.org/api/docs/v1"
)

func TestExtractTextFromBody(t *testing.T) {
	t.Run("nil body", func(t *testing.T) {
		if got := extractTextFromBody(nil); got != "" {
			t.Fatalf("expected empty, got %q", got)
		}
	})

	t.Run("paragraphs", func(t *testing.T) {
		body := &docsapi.Body{Content: []*docsapi.StructuralElement{
			{Paragraph: &docsapi.Paragraph{Elements: []*docsapi.ParagraphElement{
				{TextRun: &docsapi.TextRun{Content: "Hello "}},
				{TextRun: &docsapi.TextRun{Content: "World"}},
			}}},
		}}
		if got := extractTextFromBody(body); got != "Hello World" {
			t.Fatalf("expected text, got %q", got)
		}
	})

	t.Run("tables and breaks", func(t *testing.T) {
		body := &docsapi.Body{Content: []*docsapi.StructuralElement{
			{Table: &docsapi.Table{TableRows: []*docsapi.TableRow{
				{TableCells: []*docsapi.TableCell{
					{Content: []*docsapi.StructuralElement{
						{Paragraph: &docsapi.Paragraph{Elements: []*docsapi.ParagraphElement{
							{TextRun: &docsapi.TextRun{Content: "A1"}},
						}}},
					}},
					{Content: []*docsapi.StructuralElement{
						{Paragraph: &docsapi.Paragraph{Elements: []*docsapi.ParagraphElement{
							{TextRun: &docsapi.TextRun{Content: "B1"}},
						}}},
					}},
				}},
			}}},
			{SectionBreak: &docsapi.SectionBreak{}},
		}}
		got := extractTextFromBody(body)
		if !strings.Contains(got, "A1\tB1\t\n") || !strings.Contains(got, "\n\n") {
			t.Fatalf("unexpected table text: %q", got)
		}
	})
}

func TestCountWords(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		if got := countWords(""); got != 0 {
			t.Fatalf("expected 0, got %d", got)
		}
	})

	t.Run("whitespace", func(t *testing.T) {
		if got := countWords("   "); got != 0 {
			t.Fatalf("expected 0, got %d", got)
		}
	})

	t.Run("words", func(t *testing.T) {
		if got := countWords("one two three"); got != 3 {
			t.Fatalf("expected 3, got %d", got)
		}
	})
}
