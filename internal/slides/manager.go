package slides

import (
	"context"
	"strings"

	"github.com/dl-alexandre/gdrive/internal/api"
	"github.com/dl-alexandre/gdrive/internal/types"
	"google.golang.org/api/slides/v1"
)

type Manager struct {
	client  *api.Client
	service *slides.Service
}

func NewManager(client *api.Client, service *slides.Service) *Manager {
	return &Manager{
		client:  client,
		service: service,
	}
}

func (m *Manager) GetPresentation(ctx context.Context, reqCtx *types.RequestContext, presentationID string) (*types.Presentation, error) {
	call := m.service.Presentations.Get(presentationID)
	presentation, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*slides.Presentation, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}
	return convertPresentation(presentation), nil
}

func (m *Manager) ReadPresentation(ctx context.Context, reqCtx *types.RequestContext, presentationID string) (*types.PresentationText, error) {
	call := m.service.Presentations.Get(presentationID)
	presentation, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*slides.Presentation, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}
	return extractTextFromPresentation(presentation), nil
}

func (m *Manager) UpdatePresentation(ctx context.Context, reqCtx *types.RequestContext, presentationID string, requests []*slides.Request) (*types.SlidesBatchUpdateResponse, error) {
	call := m.service.Presentations.BatchUpdate(presentationID, &slides.BatchUpdatePresentationRequest{
		Requests: requests,
	})
	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*slides.BatchUpdatePresentationResponse, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}
	return &types.SlidesBatchUpdateResponse{
		PresentationID: presentationID,
		RepliesCount:   len(result.Replies),
	}, nil
}

func (m *Manager) ReplaceAllText(ctx context.Context, reqCtx *types.RequestContext, presentationID string, replacements map[string]string) (*types.SlidesBatchUpdateResponse, error) {
	requests := make([]*slides.Request, 0, len(replacements))
	for find, replace := range replacements {
		requests = append(requests, &slides.Request{
			ReplaceAllText: &slides.ReplaceAllTextRequest{
				ContainsText: &slides.SubstringMatchCriteria{
					Text:      find,
					MatchCase: false,
				},
				ReplaceText: replace,
			},
		})
	}
	return m.UpdatePresentation(ctx, reqCtx, presentationID, requests)
}

func convertPresentation(pres *slides.Presentation) *types.Presentation {
	if pres == nil {
		return &types.Presentation{}
	}
	result := &types.Presentation{
		PresentationID: pres.PresentationId,
		Title:          pres.Title,
		SlideCount:     len(pres.Slides),
	}
	if pres.Slides != nil {
		result.Slides = make([]types.SlideBrief, len(pres.Slides))
		for i, slide := range pres.Slides {
			result.Slides[i] = types.SlideBrief{
				ObjectID: slide.ObjectId,
			}
		}
	}
	return result
}

func extractTextFromPresentation(pres *slides.Presentation) *types.PresentationText {
	if pres == nil {
		return &types.PresentationText{}
	}
	result := &types.PresentationText{
		PresentationID: pres.PresentationId,
		Title:          pres.Title,
		SlideCount:     len(pres.Slides),
		TextBySlide:    []types.SlideText{},
	}

	for i, slide := range pres.Slides {
		for _, element := range slide.PageElements {
			if element.Shape != nil && element.Shape.Text != nil {
				text := extractTextFromShape(element.Shape.Text)
				if text != "" {
					result.TextBySlide = append(result.TextBySlide, types.SlideText{
						SlideIndex: i + 1,
						ObjectID:   element.ObjectId,
						Text:       text,
					})
				}
			}
		}
	}
	return result
}

func extractTextFromShape(content *slides.TextContent) string {
	if content == nil || content.TextElements == nil {
		return ""
	}
	var text strings.Builder
	for _, elem := range content.TextElements {
		if elem.TextRun != nil && elem.TextRun.Content != "" {
			text.WriteString(elem.TextRun.Content)
		}
	}
	return text.String()
}
