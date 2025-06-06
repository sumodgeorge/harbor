// Code generated by mockery v2.53.3. DO NOT EDIT.

package scanner

import (
	context "context"

	controllerscanner "github.com/goharbor/harbor/src/controller/scanner"
	mock "github.com/stretchr/testify/mock"

	q "github.com/goharbor/harbor/src/lib/q"

	scanner "github.com/goharbor/harbor/src/pkg/scan/dao/scanner"

	v1 "github.com/goharbor/harbor/src/pkg/scan/rest/v1"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// CreateRegistration provides a mock function with given fields: ctx, registration
func (_m *Controller) CreateRegistration(ctx context.Context, registration *scanner.Registration) (string, error) {
	ret := _m.Called(ctx, registration)

	if len(ret) == 0 {
		panic("no return value specified for CreateRegistration")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *scanner.Registration) (string, error)); ok {
		return rf(ctx, registration)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *scanner.Registration) string); ok {
		r0 = rf(ctx, registration)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *scanner.Registration) error); ok {
		r1 = rf(ctx, registration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRegistration provides a mock function with given fields: ctx, registrationUUID
func (_m *Controller) DeleteRegistration(ctx context.Context, registrationUUID string) (*scanner.Registration, error) {
	ret := _m.Called(ctx, registrationUUID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRegistration")
	}

	var r0 *scanner.Registration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*scanner.Registration, error)); ok {
		return rf(ctx, registrationUUID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *scanner.Registration); ok {
		r0 = rf(ctx, registrationUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scanner.Registration)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, registrationUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMetadata provides a mock function with given fields: ctx, registrationUUID
func (_m *Controller) GetMetadata(ctx context.Context, registrationUUID string) (*v1.ScannerAdapterMetadata, error) {
	ret := _m.Called(ctx, registrationUUID)

	if len(ret) == 0 {
		panic("no return value specified for GetMetadata")
	}

	var r0 *v1.ScannerAdapterMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*v1.ScannerAdapterMetadata, error)); ok {
		return rf(ctx, registrationUUID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *v1.ScannerAdapterMetadata); ok {
		r0 = rf(ctx, registrationUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ScannerAdapterMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, registrationUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRegistration provides a mock function with given fields: ctx, registrationUUID
func (_m *Controller) GetRegistration(ctx context.Context, registrationUUID string) (*scanner.Registration, error) {
	ret := _m.Called(ctx, registrationUUID)

	if len(ret) == 0 {
		panic("no return value specified for GetRegistration")
	}

	var r0 *scanner.Registration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*scanner.Registration, error)); ok {
		return rf(ctx, registrationUUID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *scanner.Registration); ok {
		r0 = rf(ctx, registrationUUID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scanner.Registration)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, registrationUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRegistrationByProject provides a mock function with given fields: ctx, projectID, options
func (_m *Controller) GetRegistrationByProject(ctx context.Context, projectID int64, options ...controllerscanner.Option) (*scanner.Registration, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, projectID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetRegistrationByProject")
	}

	var r0 *scanner.Registration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...controllerscanner.Option) (*scanner.Registration, error)); ok {
		return rf(ctx, projectID, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...controllerscanner.Option) *scanner.Registration); ok {
		r0 = rf(ctx, projectID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scanner.Registration)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, ...controllerscanner.Option) error); ok {
		r1 = rf(ctx, projectID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalOfRegistrations provides a mock function with given fields: ctx, query
func (_m *Controller) GetTotalOfRegistrations(ctx context.Context, query *q.Query) (int64, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for GetTotalOfRegistrations")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) (int64, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) int64); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRegistrations provides a mock function with given fields: ctx, query
func (_m *Controller) ListRegistrations(ctx context.Context, query *q.Query) ([]*scanner.Registration, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for ListRegistrations")
	}

	var r0 []*scanner.Registration
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) ([]*scanner.Registration, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*scanner.Registration); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*scanner.Registration)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ping provides a mock function with given fields: ctx, registration
func (_m *Controller) Ping(ctx context.Context, registration *scanner.Registration) (*v1.ScannerAdapterMetadata, error) {
	ret := _m.Called(ctx, registration)

	if len(ret) == 0 {
		panic("no return value specified for Ping")
	}

	var r0 *v1.ScannerAdapterMetadata
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *scanner.Registration) (*v1.ScannerAdapterMetadata, error)); ok {
		return rf(ctx, registration)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *scanner.Registration) *v1.ScannerAdapterMetadata); ok {
		r0 = rf(ctx, registration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ScannerAdapterMetadata)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *scanner.Registration) error); ok {
		r1 = rf(ctx, registration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegistrationExists provides a mock function with given fields: ctx, registrationUUID
func (_m *Controller) RegistrationExists(ctx context.Context, registrationUUID string) bool {
	ret := _m.Called(ctx, registrationUUID)

	if len(ret) == 0 {
		panic("no return value specified for RegistrationExists")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, registrationUUID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RetrieveCap provides a mock function with given fields: ctx, r
func (_m *Controller) RetrieveCap(ctx context.Context, r *scanner.Registration) error {
	ret := _m.Called(ctx, r)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveCap")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *scanner.Registration) error); ok {
		r0 = rf(ctx, r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDefaultRegistration provides a mock function with given fields: ctx, registrationUUID
func (_m *Controller) SetDefaultRegistration(ctx context.Context, registrationUUID string) error {
	ret := _m.Called(ctx, registrationUUID)

	if len(ret) == 0 {
		panic("no return value specified for SetDefaultRegistration")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, registrationUUID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetRegistrationByProject provides a mock function with given fields: ctx, projectID, scannerID
func (_m *Controller) SetRegistrationByProject(ctx context.Context, projectID int64, scannerID string) error {
	ret := _m.Called(ctx, projectID, scannerID)

	if len(ret) == 0 {
		panic("no return value specified for SetRegistrationByProject")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string) error); ok {
		r0 = rf(ctx, projectID, scannerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRegistration provides a mock function with given fields: ctx, registration
func (_m *Controller) UpdateRegistration(ctx context.Context, registration *scanner.Registration) error {
	ret := _m.Called(ctx, registration)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRegistration")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *scanner.Registration) error); ok {
		r0 = rf(ctx, registration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewController creates a new instance of Controller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewController(t interface {
	mock.TestingT
	Cleanup(func())
}) *Controller {
	mock := &Controller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
