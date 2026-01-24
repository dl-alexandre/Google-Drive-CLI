package types

import "fmt"

type Document struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	RevisionID string `json:"revisionId,omitempty"`
}

func (d *Document) Headers() []string {
	return []string{"Document ID", "Title", "Revision"}
}

func (d *Document) Rows() [][]string {
	return [][]string{{
		truncateDocID(d.ID, 20),
		truncateDocText(d.Title, 40),
		d.RevisionID,
	}}
}

func (d *Document) EmptyMessage() string {
	return "No document data"
}

type DocumentText struct {
	DocumentID string `json:"documentId"`
	Title      string `json:"title"`
	Text       string `json:"text"`
	WordCount  int    `json:"wordCount"`
	CharCount  int    `json:"charCount"`
}

func (t *DocumentText) Headers() []string {
	return []string{"Document ID", "Title", "Excerpt", "Words", "Characters"}
}

func (t *DocumentText) Rows() [][]string {
	return [][]string{{
		truncateDocID(t.DocumentID, 20),
		truncateDocText(t.Title, 40),
		truncateDocText(t.Text, 80),
		fmt.Sprintf("%d", t.WordCount),
		fmt.Sprintf("%d", t.CharCount),
	}}
}

func (t *DocumentText) EmptyMessage() string {
	return "No text content found"
}

type UpdateDocumentResponse struct {
	DocumentID string `json:"documentId"`
	RevisionID string `json:"revisionId"`
}

func (r *UpdateDocumentResponse) Headers() []string {
	return []string{"Document ID", "Revision"}
}

func (r *UpdateDocumentResponse) Rows() [][]string {
	return [][]string{{
		truncateDocID(r.DocumentID, 20),
		r.RevisionID,
	}}
}

func (r *UpdateDocumentResponse) EmptyMessage() string {
	return "No update information available"
}

func truncateDocID(id string, maxLen int) string {
	if len(id) <= maxLen {
		return id
	}
	if maxLen <= 3 {
		return id[:maxLen]
	}
	return id[:maxLen-3] + "..."
}

func truncateDocText(s string, max int) string {
	if len(s) <= max {
		return s
	}
	if max <= 3 {
		return s[:max]
	}
	return s[:max-3] + "..."
}

