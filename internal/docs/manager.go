package docs

import (
	"context"
	"strings"

	"github.com/dl-alexandre/gdrive/internal/api"
	"github.com/dl-alexandre/gdrive/internal/types"
	"google.golang.org/api/docs/v1"
)

type Manager struct {
	client  *api.Client
	service *docs.Service
}

func NewManager(client *api.Client, service *docs.Service) *Manager {
	return &Manager{
		client:  client,
		service: service,
	}
}

func (m *Manager) GetDocument(ctx context.Context, reqCtx *types.RequestContext, documentID string) (*types.Document, error) {
	call := m.service.Documents.Get(documentID)
	doc, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*docs.Document, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}
	return &types.Document{
		ID:         doc.DocumentId,
		Title:      doc.Title,
		RevisionID: doc.RevisionId,
	}, nil
}

func (m *Manager) ReadDocument(ctx context.Context, reqCtx *types.RequestContext, documentID string) (*types.DocumentText, error) {
	call := m.service.Documents.Get(documentID)
	doc, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*docs.Document, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}
	text := extractTextFromBody(doc.Body)
	return &types.DocumentText{
		DocumentID: documentID,
		Title:      doc.Title,
		Text:       text,
		WordCount:  countWords(text),
		CharCount:  len(text),
	}, nil
}

func (m *Manager) UpdateDocument(ctx context.Context, reqCtx *types.RequestContext, documentID string, requests []*docs.Request) (*types.UpdateDocumentResponse, error) {
	call := m.service.Documents.BatchUpdate(documentID, &docs.BatchUpdateDocumentRequest{
		Requests: requests,
	})
	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*docs.BatchUpdateDocumentResponse, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}
	revisionID := ""
	if result != nil && result.DocumentId != "" {
		documentID = result.DocumentId
	}
	if result != nil && result.WriteControl != nil {
		revisionID = result.WriteControl.RequiredRevisionId
	}
	return &types.UpdateDocumentResponse{
		DocumentID: documentID,
		RevisionID: revisionID,
	}, nil
}

func extractTextFromBody(body *docs.Body) string {
	if body == nil || body.Content == nil {
		return ""
	}
	var text strings.Builder
	for _, element := range body.Content {
		extractTextFromElement(element, &text)
	}
	return text.String()
}

func extractTextFromElement(element *docs.StructuralElement, text *strings.Builder) {
	if element == nil {
		return
	}
	if element.Paragraph != nil && element.Paragraph.Elements != nil {
		for _, paraElem := range element.Paragraph.Elements {
			if paraElem.TextRun != nil && paraElem.TextRun.Content != "" {
				text.WriteString(paraElem.TextRun.Content)
			}
		}
	}
	if element.Table != nil && element.Table.TableRows != nil {
		for _, row := range element.Table.TableRows {
			if row.TableCells != nil {
				for _, cell := range row.TableCells {
					if cell.Content != nil {
						for _, cellElem := range cell.Content {
							extractTextFromElement(cellElem, text)
						}
					}
					text.WriteString("\t")
				}
			}
			text.WriteString("\n")
		}
	}
	if element.SectionBreak != nil {
		text.WriteString("\n\n")
	}
}

func countWords(text string) int {
	if text == "" {
		return 0
	}
	return len(strings.Fields(text))
}
