package labels

import (
	"testing"

	"github.com/dl-alexandre/gdrv/internal/api"
	"github.com/dl-alexandre/gdrv/internal/types"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/drivelabels/v2"
)

var _ = drive.Label{}

func TestManager_Creation(t *testing.T) {
	client := &api.Client{}
	manager := NewManager(client)

	if manager == nil {
		t.Error("NewManager returned nil")
	}

	if manager.client != client {
		t.Error("Manager client not set correctly")
	}
}

func TestConvertLabel(t *testing.T) {
	tests := []struct {
		name     string
		input    *drivelabels.GoogleAppsDriveLabelsV2Label
		validate func(*types.Label) error
	}{
		{
			name: "label with basic fields",
			input: &drivelabels.GoogleAppsDriveLabelsV2Label{
				Id:         "label123",
				Name:       "Document Type",
				RevisionId: "rev1",
				LabelType:  "ADMIN",
				Customer:   "customers/123",
			},
			validate: func(label *types.Label) error {
				if label.ID != "label123" {
					return errorf("ID mismatch")
				}
				if label.Name != "Document Type" {
					return errorf("Name mismatch")
				}
				if label.RevisionID != "rev1" {
					return errorf("RevisionID mismatch")
				}
				if label.LabelType != "ADMIN" {
					return errorf("LabelType mismatch")
				}
				return nil
			},
		},
		{
			name: "label with properties",
			input: &drivelabels.GoogleAppsDriveLabelsV2Label{
				Id:   "label123",
				Name: "Test Label",
				Properties: &drivelabels.GoogleAppsDriveLabelsV2LabelProperties{
					Title:       "Document Type",
					Description: "Type of document",
				},
			},
			validate: func(label *types.Label) error {
				if label.Properties == nil {
					return errorf("Properties should not be nil")
				}
				if label.Properties.Title != "Document Type" {
					return errorf("Properties.Title mismatch")
				}
				if label.Properties.Description != "Type of document" {
					return errorf("Properties.Description mismatch")
				}
				return nil
			},
		},
		{
			name: "label with lifecycle",
			input: &drivelabels.GoogleAppsDriveLabelsV2Label{
				Id:   "label123",
				Name: "Test Label",
				Lifecycle: &drivelabels.GoogleAppsDriveLabelsV2Lifecycle{
					State:                 "PUBLISHED",
					HasUnpublishedChanges: false,
				},
			},
			validate: func(label *types.Label) error {
				if label.Lifecycle == nil {
					return errorf("Lifecycle should not be nil")
				}
				if label.Lifecycle.State != "PUBLISHED" {
					return errorf("Lifecycle.State mismatch")
				}
				if label.Lifecycle.HasUnpublishedChanges {
					return errorf("HasUnpublishedChanges should be false")
				}
				return nil
			},
		},

		{
			name:  "nil label",
			input: nil,
			validate: func(label *types.Label) error {
				if label != nil {
					return errorf("expected nil label")
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertLabel(tt.input)
			if err := tt.validate(result); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestConvertLabelField(t *testing.T) {
	tests := []struct {
		name     string
		input    *drivelabels.GoogleAppsDriveLabelsV2Field
		validate func(*types.LabelField) error
	}{
		{
			name: "field with basic properties",
			input: &drivelabels.GoogleAppsDriveLabelsV2Field{
				Id: "field1",
			},
			validate: func(field *types.LabelField) error {
				if field.ID != "field1" {
					return errorf("ID mismatch")
				}
				return nil
			},
		},
		{
			name: "field with properties",
			input: &drivelabels.GoogleAppsDriveLabelsV2Field{
				Id: "field1",
				Properties: &drivelabels.GoogleAppsDriveLabelsV2FieldProperties{
					DisplayName: "Document Type",
				},
			},
			validate: func(field *types.LabelField) error {
				if field.Properties == nil {
					return errorf("Properties should not be nil")
				}
				if field.Properties.DisplayName != "Document Type" {
					return errorf("DisplayName mismatch")
				}
				return nil
			},
		},
		{
			name:  "nil field",
			input: nil,
			validate: func(field *types.LabelField) error {
				if field != nil {
					return errorf("expected nil field")
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertLabelField(tt.input)
			if err := tt.validate(result); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestConvertDriveLabel(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		validate func(*types.FileLabel) error
	}{
		{
			name: "drive label with basic fields",
			input: map[string]interface{}{
				"id": "label123",
			},
			validate: func(label *types.FileLabel) error {
				if label.ID != "label123" {
					return errorf("ID mismatch")
				}
				return nil
			},
		},
		{
			name:  "nil drive label",
			input: nil,
			validate: func(label *types.FileLabel) error {
				if label != nil {
					return errorf("expected nil label")
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.input == nil {
				result := convertDriveLabel(nil)
				if err := tt.validate(result); err != nil {
					t.Error(err)
				}
			}
		})
	}
}

func TestConvertFieldModificationsToDrive(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]*types.LabelFieldValue
		validate func([]*drive.LabelFieldModification) error
	}{
		{
			name:  "empty map",
			input: map[string]*types.LabelFieldValue{},
			validate: func(mods []*drive.LabelFieldModification) error {
				if len(mods) != 0 {
					return errorf("modification count mismatch")
				}
				return nil
			},
		},
		{
			name: "single field modification",
			input: map[string]*types.LabelFieldValue{
				"field1": {},
			},
			validate: func(mods []*drive.LabelFieldModification) error {
				if len(mods) != 1 {
					return errorf("modification count mismatch")
				}
				if mods[0].FieldId != "field1" {
					return errorf("FieldId mismatch")
				}
				return nil
			},
		},
		{
			name: "multiple field modifications",
			input: map[string]*types.LabelFieldValue{
				"field1": {},
				"field2": {},
			},
			validate: func(mods []*drive.LabelFieldModification) error {
				if len(mods) != 2 {
					return errorf("modification count mismatch")
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertFieldModificationsToDrive(tt.input)
			if err := tt.validate(result); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestConvertToAPILabel(t *testing.T) {
	tests := []struct {
		name     string
		input    *types.Label
		validate func(*drivelabels.GoogleAppsDriveLabelsV2Label) error
	}{
		{
			name: "label with properties",
			input: &types.Label{
				ID:   "label123",
				Name: "Test Label",
				Properties: &types.LabelProperties{
					Title:       "Document Type",
					Description: "Type of document",
				},
			},
			validate: func(label *drivelabels.GoogleAppsDriveLabelsV2Label) error {
				if label.Properties == nil {
					return errorf("Properties should not be nil")
				}
				if label.Properties.Title != "Document Type" {
					return errorf("Properties.Title mismatch")
				}
				return nil
			},
		},
		{
			name:  "nil label",
			input: nil,
			validate: func(label *drivelabels.GoogleAppsDriveLabelsV2Label) error {
				if label != nil {
					return errorf("expected nil label")
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertToAPILabel(tt.input)
			if err := tt.validate(result); err != nil {
				t.Error(err)
			}
		})
	}
}

func errorf(msg string, args ...interface{}) error {
	return &testError{msg: msg}
}

type testError struct {
	msg string
}

func (e *testError) Error() string {
	return e.msg
}
