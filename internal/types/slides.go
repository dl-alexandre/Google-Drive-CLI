package types

import (
	"fmt"
)

type Presentation struct {
	PresentationID string       `json:"presentationId"`
	Title          string       `json:"title"`
	SlideCount     int          `json:"slideCount"`
	Slides         []SlideBrief `json:"slides,omitempty"`
}

func (p *Presentation) Headers() []string {
	return []string{"Slide", "Object ID"}
}

func (p *Presentation) Rows() [][]string {
	rows := make([][]string, len(p.Slides))
	for i, slide := range p.Slides {
		rows[i] = []string{
			fmt.Sprintf("%d", i+1),
			slide.ObjectID,
		}
	}
	return rows
}

func (p *Presentation) EmptyMessage() string {
	return "No slides found"
}

type SlideBrief struct {
	ObjectID string `json:"objectId"`
}

type PresentationText struct {
	PresentationID string      `json:"presentationId"`
	Title          string      `json:"title"`
	SlideCount     int         `json:"slideCount"`
	TextBySlide    []SlideText `json:"textBySlide"`
}

func (p *PresentationText) Headers() []string {
	return []string{"Slide", "Object ID", "Text"}
}

func (p *PresentationText) Rows() [][]string {
	rows := make([][]string, len(p.TextBySlide))
	for i, st := range p.TextBySlide {
		rows[i] = []string{
			fmt.Sprintf("%d", st.SlideIndex),
			st.ObjectID,
			truncateSlideText(st.Text, 50),
		}
	}
	return rows
}

func (p *PresentationText) EmptyMessage() string {
	return "No text content found"
}

type SlideText struct {
	SlideIndex int    `json:"slideIndex"`
	ObjectID   string `json:"objectId"`
	Text       string `json:"text"`
}

type SlidesBatchUpdateResponse struct {
	PresentationID string `json:"presentationId"`
	RepliesCount   int    `json:"repliesCount"`
}

func (r *SlidesBatchUpdateResponse) Headers() []string {
	return []string{"Presentation ID", "Replies"}
}

func (r *SlidesBatchUpdateResponse) Rows() [][]string {
	return [][]string{{
		r.PresentationID,
		fmt.Sprintf("%d", r.RepliesCount),
	}}
}

func (r *SlidesBatchUpdateResponse) EmptyMessage() string {
	return "No update information available"
}

func truncateSlideText(s string, max int) string {
	if len(s) <= max {
		return s
	}
	if max <= 3 {
		return s[:max]
	}
	return s[:max-3] + "..."
}
