// Code generated by mockery v2.43.2. DO NOT EDIT.

package fcm

import mock "github.com/stretchr/testify/mock"

// mockOption is an autogenerated mock type for the Option type
type mockOption struct {
	mock.Mock
}

type mockOption_Expecter struct {
	mock *mock.Mock
}

func (_m *mockOption) EXPECT() *mockOption_Expecter {
	return &mockOption_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *mockOption) Execute(_a0 *Service) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*Service) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockOption_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type mockOption_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 *Service
func (_e *mockOption_Expecter) Execute(_a0 interface{}) *mockOption_Execute_Call {
	return &mockOption_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *mockOption_Execute_Call) Run(run func(_a0 *Service)) *mockOption_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Service))
	})
	return _c
}

func (_c *mockOption_Execute_Call) Return(_a0 error) *mockOption_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockOption_Execute_Call) RunAndReturn(run func(*Service) error) *mockOption_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// newMockOption creates a new instance of mockOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockOption {
	mock := &mockOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}