// Code generated by mockery v2.52.2. DO NOT EDIT.

package dashboards

import (
	context "context"

	identity "github.com/grafana/grafana/pkg/apimachinery/identity"
	mock "github.com/stretchr/testify/mock"

	model "github.com/grafana/grafana/pkg/services/search/model"
)

// FakeDashboardService is an autogenerated mock type for the DashboardService type
type FakeDashboardService struct {
	mock.Mock
}

// BuildSaveDashboardCommand provides a mock function with given fields: ctx, dto, validateProvisionedDashboard
func (_m *FakeDashboardService) BuildSaveDashboardCommand(ctx context.Context, dto *SaveDashboardDTO, validateProvisionedDashboard bool) (*SaveDashboardCommand, error) {
	ret := _m.Called(ctx, dto, validateProvisionedDashboard)

	if len(ret) == 0 {
		panic("no return value specified for BuildSaveDashboardCommand")
	}

	var r0 *SaveDashboardCommand
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *SaveDashboardDTO, bool) (*SaveDashboardCommand, error)); ok {
		return rf(ctx, dto, validateProvisionedDashboard)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *SaveDashboardDTO, bool) *SaveDashboardCommand); ok {
		r0 = rf(ctx, dto, validateProvisionedDashboard)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*SaveDashboardCommand)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *SaveDashboardDTO, bool) error); ok {
		r1 = rf(ctx, dto, validateProvisionedDashboard)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CleanUpDeletedDashboards provides a mock function with given fields: ctx
func (_m *FakeDashboardService) CleanUpDeletedDashboards(ctx context.Context) (int64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CleanUpDeletedDashboards")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (int64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountDashboardsInOrg provides a mock function with given fields: ctx, orgID
func (_m *FakeDashboardService) CountDashboardsInOrg(ctx context.Context, orgID int64) (int64, error) {
	ret := _m.Called(ctx, orgID)

	if len(ret) == 0 {
		panic("no return value specified for CountDashboardsInOrg")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (int64, error)); ok {
		return rf(ctx, orgID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) int64); ok {
		r0 = rf(ctx, orgID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, orgID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountInFolders provides a mock function with given fields: ctx, orgID, folderUIDs, user
func (_m *FakeDashboardService) CountInFolders(ctx context.Context, orgID int64, folderUIDs []string, user identity.Requester) (int64, error) {
	ret := _m.Called(ctx, orgID, folderUIDs, user)

	if len(ret) == 0 {
		panic("no return value specified for CountInFolders")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, []string, identity.Requester) (int64, error)); ok {
		return rf(ctx, orgID, folderUIDs, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, []string, identity.Requester) int64); ok {
		r0 = rf(ctx, orgID, folderUIDs, user)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, []string, identity.Requester) error); ok {
		r1 = rf(ctx, orgID, folderUIDs, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAllDashboards provides a mock function with given fields: ctx, orgID
func (_m *FakeDashboardService) DeleteAllDashboards(ctx context.Context, orgID int64) error {
	ret := _m.Called(ctx, orgID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAllDashboards")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, orgID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDashboard provides a mock function with given fields: ctx, dashboardId, dashboardUID, orgId
func (_m *FakeDashboardService) DeleteDashboard(ctx context.Context, dashboardId int64, dashboardUID string, orgId int64) error {
	ret := _m.Called(ctx, dashboardId, dashboardUID, orgId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDashboard")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, int64) error); ok {
		r0 = rf(ctx, dashboardId, dashboardUID, orgId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindDashboards provides a mock function with given fields: ctx, query
func (_m *FakeDashboardService) FindDashboards(ctx context.Context, query *FindPersistedDashboardsQuery) ([]DashboardSearchProjection, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for FindDashboards")
	}

	var r0 []DashboardSearchProjection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *FindPersistedDashboardsQuery) ([]DashboardSearchProjection, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *FindPersistedDashboardsQuery) []DashboardSearchProjection); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]DashboardSearchProjection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *FindPersistedDashboardsQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllDashboards provides a mock function with given fields: ctx
func (_m *FakeDashboardService) GetAllDashboards(ctx context.Context) ([]*Dashboard, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllDashboards")
	}

	var r0 []*Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*Dashboard, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*Dashboard); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllDashboardsByOrgId provides a mock function with given fields: ctx, orgID
func (_m *FakeDashboardService) GetAllDashboardsByOrgId(ctx context.Context, orgID int64) ([]*Dashboard, error) {
	ret := _m.Called(ctx, orgID)

	if len(ret) == 0 {
		panic("no return value specified for GetAllDashboardsByOrgId")
	}

	var r0 []*Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]*Dashboard, error)); ok {
		return rf(ctx, orgID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*Dashboard); ok {
		r0 = rf(ctx, orgID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, orgID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDashboard provides a mock function with given fields: ctx, query
func (_m *FakeDashboardService) GetDashboard(ctx context.Context, query *GetDashboardQuery) (*Dashboard, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for GetDashboard")
	}

	var r0 *Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardQuery) (*Dashboard, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardQuery) *Dashboard); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetDashboardQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDashboardTags provides a mock function with given fields: ctx, query
func (_m *FakeDashboardService) GetDashboardTags(ctx context.Context, query *GetDashboardTagsQuery) ([]*DashboardTagCloudItem, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for GetDashboardTags")
	}

	var r0 []*DashboardTagCloudItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardTagsQuery) ([]*DashboardTagCloudItem, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardTagsQuery) []*DashboardTagCloudItem); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*DashboardTagCloudItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetDashboardTagsQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDashboardUIDByID provides a mock function with given fields: ctx, query
func (_m *FakeDashboardService) GetDashboardUIDByID(ctx context.Context, query *GetDashboardRefByIDQuery) (*DashboardRef, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for GetDashboardUIDByID")
	}

	var r0 *DashboardRef
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardRefByIDQuery) (*DashboardRef, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardRefByIDQuery) *DashboardRef); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DashboardRef)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetDashboardRefByIDQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDashboards provides a mock function with given fields: ctx, query
func (_m *FakeDashboardService) GetDashboards(ctx context.Context, query *GetDashboardsQuery) ([]*Dashboard, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for GetDashboards")
	}

	var r0 []*Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardsQuery) ([]*Dashboard, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetDashboardsQuery) []*Dashboard); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetDashboardsQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSoftDeletedDashboard provides a mock function with given fields: ctx, orgID, uid
func (_m *FakeDashboardService) GetSoftDeletedDashboard(ctx context.Context, orgID int64, uid string) (*Dashboard, error) {
	ret := _m.Called(ctx, orgID, uid)

	if len(ret) == 0 {
		panic("no return value specified for GetSoftDeletedDashboard")
	}

	var r0 *Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) (*Dashboard, error)); ok {
		return rf(ctx, orgID, uid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) *Dashboard); ok {
		r0 = rf(ctx, orgID, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string) error); ok {
		r1 = rf(ctx, orgID, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportDashboard provides a mock function with given fields: ctx, dto
func (_m *FakeDashboardService) ImportDashboard(ctx context.Context, dto *SaveDashboardDTO) (*Dashboard, error) {
	ret := _m.Called(ctx, dto)

	if len(ret) == 0 {
		panic("no return value specified for ImportDashboard")
	}

	var r0 *Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *SaveDashboardDTO) (*Dashboard, error)); ok {
		return rf(ctx, dto)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *SaveDashboardDTO) *Dashboard); ok {
		r0 = rf(ctx, dto)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *SaveDashboardDTO) error); ok {
		r1 = rf(ctx, dto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestoreDashboard provides a mock function with given fields: ctx, dashboard, user, optionalFolderUID
func (_m *FakeDashboardService) RestoreDashboard(ctx context.Context, dashboard *Dashboard, user identity.Requester, optionalFolderUID string) error {
	ret := _m.Called(ctx, dashboard, user, optionalFolderUID)

	if len(ret) == 0 {
		panic("no return value specified for RestoreDashboard")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *Dashboard, identity.Requester, string) error); ok {
		r0 = rf(ctx, dashboard, user, optionalFolderUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveDashboard provides a mock function with given fields: ctx, dto, allowUiUpdate
func (_m *FakeDashboardService) SaveDashboard(ctx context.Context, dto *SaveDashboardDTO, allowUiUpdate bool) (*Dashboard, error) {
	ret := _m.Called(ctx, dto, allowUiUpdate)

	if len(ret) == 0 {
		panic("no return value specified for SaveDashboard")
	}

	var r0 *Dashboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *SaveDashboardDTO, bool) (*Dashboard, error)); ok {
		return rf(ctx, dto, allowUiUpdate)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *SaveDashboardDTO, bool) *Dashboard); ok {
		r0 = rf(ctx, dto, allowUiUpdate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Dashboard)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *SaveDashboardDTO, bool) error); ok {
		r1 = rf(ctx, dto, allowUiUpdate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchDashboards provides a mock function with given fields: ctx, query
func (_m *FakeDashboardService) SearchDashboards(ctx context.Context, query *FindPersistedDashboardsQuery) (model.HitList, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for SearchDashboards")
	}

	var r0 model.HitList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *FindPersistedDashboardsQuery) (model.HitList, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *FindPersistedDashboardsQuery) model.HitList); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.HitList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *FindPersistedDashboardsQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SoftDeleteDashboard provides a mock function with given fields: ctx, orgID, dashboardUid
func (_m *FakeDashboardService) SoftDeleteDashboard(ctx context.Context, orgID int64, dashboardUid string) error {
	ret := _m.Called(ctx, orgID, dashboardUid)

	if len(ret) == 0 {
		panic("no return value specified for SoftDeleteDashboard")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) error); ok {
		r0 = rf(ctx, orgID, dashboardUid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CleanUpDashboard provides a mock function with given fields: ctx, dashboardUID, orgId
func (_m *FakeDashboardService) CleanUpDashboard(ctx context.Context, dashboardUID string, orgId int64) error {
	ret := _m.Called(ctx, dashboardUID, orgId)

	if len(ret) == 0 {
		panic("no return value specified for CleanUpDashboard")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) error); ok {
		r0 = rf(ctx, dashboardUID, orgId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewFakeDashboardService creates a new instance of FakeDashboardService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFakeDashboardService(t interface {
	mock.TestingT
	Cleanup(func())
}) *FakeDashboardService {
	mock := &FakeDashboardService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
