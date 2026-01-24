package auth

import (
	"context"
	"fmt"

	"github.com/dl-alexandre/gdrive/internal/types"
	"google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/docs/v1"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"google.golang.org/api/slides/v1"
)

type ServiceType string

const (
	ServiceDrive    ServiceType = "drive"
	ServiceSheets   ServiceType = "sheets"
	ServiceDocs     ServiceType = "docs"
	ServiceSlides   ServiceType = "slides"
	ServiceAdminDir ServiceType = "admin_directory"
)

type ServiceFactory struct {
	manager *Manager
}

func NewServiceFactory(manager *Manager) *ServiceFactory {
	return &ServiceFactory{manager: manager}
}

func (f *ServiceFactory) CreateService(ctx context.Context, creds *types.Credentials, svcType ServiceType) (interface{}, error) {
	switch svcType {
	case ServiceDrive:
		return f.CreateDriveService(ctx, creds)
	case ServiceSheets:
		return f.CreateSheetsService(ctx, creds)
	case ServiceDocs:
		return f.CreateDocsService(ctx, creds)
	case ServiceSlides:
		return f.CreateSlidesService(ctx, creds)
	case ServiceAdminDir:
		return f.CreateAdminService(ctx, creds)
	default:
		return nil, fmt.Errorf("unknown service type: %s", svcType)
	}
}

func (f *ServiceFactory) CreateDriveService(ctx context.Context, creds *types.Credentials) (*drive.Service, error) {
	client := f.manager.GetHTTPClient(ctx, creds)
	return drive.NewService(ctx, option.WithHTTPClient(client))
}

func (f *ServiceFactory) CreateSheetsService(ctx context.Context, creds *types.Credentials) (*sheets.Service, error) {
	client := f.manager.GetHTTPClient(ctx, creds)
	return sheets.NewService(ctx, option.WithHTTPClient(client))
}

func (f *ServiceFactory) CreateDocsService(ctx context.Context, creds *types.Credentials) (*docs.Service, error) {
	client := f.manager.GetHTTPClient(ctx, creds)
	return docs.NewService(ctx, option.WithHTTPClient(client))
}

func (f *ServiceFactory) CreateSlidesService(ctx context.Context, creds *types.Credentials) (*slides.Service, error) {
	client := f.manager.GetHTTPClient(ctx, creds)
	return slides.NewService(ctx, option.WithHTTPClient(client))
}

func (f *ServiceFactory) CreateAdminService(ctx context.Context, creds *types.Credentials) (*admin.Service, error) {
	client := f.manager.GetHTTPClient(ctx, creds)
	return admin.NewService(ctx, option.WithHTTPClient(client))
}
