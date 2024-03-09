// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	settings "github.com/maks112v/photomanager/pkg/settings"
	mock "github.com/stretchr/testify/mock"
)

// SettingsProvider is an autogenerated mock type for the SettingsProvider type
type SettingsProvider struct {
	mock.Mock
}

// GetSettings provides a mock function with given fields:
func (_m *SettingsProvider) GetSettings() (*settings.Settings, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSettings")
	}

	var r0 *settings.Settings
	var r1 error
	if rf, ok := ret.Get(0).(func() (*settings.Settings, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *settings.Settings); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.Settings)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveSettings provides a mock function with given fields: _a0
func (_m *SettingsProvider) SaveSettings(_a0 *settings.Settings) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SaveSettings")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*settings.Settings) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSettingsProvider creates a new instance of SettingsProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSettingsProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *SettingsProvider {
	mock := &SettingsProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}