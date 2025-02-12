// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs provide specific management operations for Lakeview dashboards.
package dashboards

import (
	"context"

	"github.com/databricks/databricks-sdk-go/client"
)

type LakeviewInterface interface {
	// WithImpl could be used to override low-level API implementations for unit
	// testing purposes with [github.com/golang/mock] or other mocking frameworks.
	// Deprecated: use MockLakeviewInterface instead.
	WithImpl(impl LakeviewService) LakeviewInterface

	// Impl returns low-level Lakeview API implementation
	// Deprecated: use MockLakeviewInterface instead.
	Impl() LakeviewService

	// Create dashboard.
	//
	// Create a draft dashboard.
	Create(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error)

	// Get dashboard.
	//
	// Get a draft dashboard.
	Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error)

	// Get dashboard.
	//
	// Get a draft dashboard.
	GetByDashboardId(ctx context.Context, dashboardId string) (*Dashboard, error)

	// Get published dashboard.
	//
	// Get the current published dashboard.
	GetPublished(ctx context.Context, request GetPublishedDashboardRequest) (*PublishedDashboard, error)

	// Get published dashboard.
	//
	// Get the current published dashboard.
	GetPublishedByDashboardId(ctx context.Context, dashboardId string) (*PublishedDashboard, error)

	// Migrate dashboard.
	//
	// Migrates a classic SQL dashboard to Lakeview.
	Migrate(ctx context.Context, request MigrateDashboardRequest) (*Dashboard, error)

	// Publish dashboard.
	//
	// Publish the current draft dashboard.
	Publish(ctx context.Context, request PublishRequest) (*PublishedDashboard, error)

	// Trash dashboard.
	//
	// Trash a dashboard.
	Trash(ctx context.Context, request TrashDashboardRequest) error

	// Trash dashboard.
	//
	// Trash a dashboard.
	TrashByDashboardId(ctx context.Context, dashboardId string) error

	// Unpublish dashboard.
	//
	// Unpublish the dashboard.
	Unpublish(ctx context.Context, request UnpublishDashboardRequest) error

	// Unpublish dashboard.
	//
	// Unpublish the dashboard.
	UnpublishByDashboardId(ctx context.Context, dashboardId string) error

	// Update dashboard.
	//
	// Update a draft dashboard.
	Update(ctx context.Context, request UpdateDashboardRequest) (*Dashboard, error)
}

func NewLakeview(client *client.DatabricksClient) *LakeviewAPI {
	return &LakeviewAPI{
		impl: &lakeviewImpl{
			client: client,
		},
	}
}

// These APIs provide specific management operations for Lakeview dashboards.
// Generic resource management can be done with Workspace API (import, export,
// get-status, list, delete).
type LakeviewAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(LakeviewService)
	impl LakeviewService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
// Deprecated: use MockLakeviewInterface instead.
func (a *LakeviewAPI) WithImpl(impl LakeviewService) LakeviewInterface {
	a.impl = impl
	return a
}

// Impl returns low-level Lakeview API implementation
// Deprecated: use MockLakeviewInterface instead.
func (a *LakeviewAPI) Impl() LakeviewService {
	return a.impl
}

// Create dashboard.
//
// Create a draft dashboard.
func (a *LakeviewAPI) Create(ctx context.Context, request CreateDashboardRequest) (*Dashboard, error) {
	return a.impl.Create(ctx, request)
}

// Get dashboard.
//
// Get a draft dashboard.
func (a *LakeviewAPI) Get(ctx context.Context, request GetDashboardRequest) (*Dashboard, error) {
	return a.impl.Get(ctx, request)
}

// Get dashboard.
//
// Get a draft dashboard.
func (a *LakeviewAPI) GetByDashboardId(ctx context.Context, dashboardId string) (*Dashboard, error) {
	return a.impl.Get(ctx, GetDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Get published dashboard.
//
// Get the current published dashboard.
func (a *LakeviewAPI) GetPublished(ctx context.Context, request GetPublishedDashboardRequest) (*PublishedDashboard, error) {
	return a.impl.GetPublished(ctx, request)
}

// Get published dashboard.
//
// Get the current published dashboard.
func (a *LakeviewAPI) GetPublishedByDashboardId(ctx context.Context, dashboardId string) (*PublishedDashboard, error) {
	return a.impl.GetPublished(ctx, GetPublishedDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Migrate dashboard.
//
// Migrates a classic SQL dashboard to Lakeview.
func (a *LakeviewAPI) Migrate(ctx context.Context, request MigrateDashboardRequest) (*Dashboard, error) {
	return a.impl.Migrate(ctx, request)
}

// Publish dashboard.
//
// Publish the current draft dashboard.
func (a *LakeviewAPI) Publish(ctx context.Context, request PublishRequest) (*PublishedDashboard, error) {
	return a.impl.Publish(ctx, request)
}

// Trash dashboard.
//
// Trash a dashboard.
func (a *LakeviewAPI) Trash(ctx context.Context, request TrashDashboardRequest) error {
	return a.impl.Trash(ctx, request)
}

// Trash dashboard.
//
// Trash a dashboard.
func (a *LakeviewAPI) TrashByDashboardId(ctx context.Context, dashboardId string) error {
	return a.impl.Trash(ctx, TrashDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Unpublish dashboard.
//
// Unpublish the dashboard.
func (a *LakeviewAPI) Unpublish(ctx context.Context, request UnpublishDashboardRequest) error {
	return a.impl.Unpublish(ctx, request)
}

// Unpublish dashboard.
//
// Unpublish the dashboard.
func (a *LakeviewAPI) UnpublishByDashboardId(ctx context.Context, dashboardId string) error {
	return a.impl.Unpublish(ctx, UnpublishDashboardRequest{
		DashboardId: dashboardId,
	})
}

// Update dashboard.
//
// Update a draft dashboard.
func (a *LakeviewAPI) Update(ctx context.Context, request UpdateDashboardRequest) (*Dashboard, error) {
	return a.impl.Update(ctx, request)
}
