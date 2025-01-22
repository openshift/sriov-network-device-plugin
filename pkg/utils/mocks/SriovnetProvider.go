// Code generated by mockery v2.49.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SriovnetProvider is an autogenerated mock type for the SriovnetProvider type
type SriovnetProvider struct {
	mock.Mock
}

// GetAuxNetDevicesFromPci provides a mock function with given fields: pciAddr
func (_m *SriovnetProvider) GetAuxNetDevicesFromPci(pciAddr string) ([]string, error) {
	ret := _m.Called(pciAddr)

	if len(ret) == 0 {
		panic("no return value specified for GetAuxNetDevicesFromPci")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]string, error)); ok {
		return rf(pciAddr)
	}
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(pciAddr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(pciAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDefaultPKeyFromPci provides a mock function with given fields: pciAddr
func (_m *SriovnetProvider) GetDefaultPKeyFromPci(pciAddr string) (string, error) {
	ret := _m.Called(pciAddr)

	if len(ret) == 0 {
		panic("no return value specified for GetDefaultPKeyFromPci")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(pciAddr)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(pciAddr)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(pciAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNetDevicesFromAux provides a mock function with given fields: auxDev
func (_m *SriovnetProvider) GetNetDevicesFromAux(auxDev string) ([]string, error) {
	ret := _m.Called(auxDev)

	if len(ret) == 0 {
		panic("no return value specified for GetNetDevicesFromAux")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]string, error)); ok {
		return rf(auxDev)
	}
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(auxDev)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(auxDev)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPfPciFromAux provides a mock function with given fields: auxDev
func (_m *SriovnetProvider) GetPfPciFromAux(auxDev string) (string, error) {
	ret := _m.Called(auxDev)

	if len(ret) == 0 {
		panic("no return value specified for GetPfPciFromAux")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(auxDev)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(auxDev)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(auxDev)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSfIndexByAuxDev provides a mock function with given fields: auxDev
func (_m *SriovnetProvider) GetSfIndexByAuxDev(auxDev string) (int, error) {
	ret := _m.Called(auxDev)

	if len(ret) == 0 {
		panic("no return value specified for GetSfIndexByAuxDev")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int, error)); ok {
		return rf(auxDev)
	}
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(auxDev)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(auxDev)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUplinkRepresentor provides a mock function with given fields: vfPciAddress
func (_m *SriovnetProvider) GetUplinkRepresentor(vfPciAddress string) (string, error) {
	ret := _m.Called(vfPciAddress)

	if len(ret) == 0 {
		panic("no return value specified for GetUplinkRepresentor")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(vfPciAddress)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(vfPciAddress)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(vfPciAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUplinkRepresentorFromAux provides a mock function with given fields: auxDev
func (_m *SriovnetProvider) GetUplinkRepresentorFromAux(auxDev string) (string, error) {
	ret := _m.Called(auxDev)

	if len(ret) == 0 {
		panic("no return value specified for GetUplinkRepresentorFromAux")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(auxDev)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(auxDev)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(auxDev)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSriovnetProvider creates a new instance of SriovnetProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSriovnetProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *SriovnetProvider {
	mock := &SriovnetProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
