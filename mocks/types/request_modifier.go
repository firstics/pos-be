// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// RequestModifier is an autogenerated mock type for the RequestModifier type
type RequestModifier struct {
	mock.Mock
}

type RequestModifier_Expecter struct {
	mock *mock.Mock
}

func (_m *RequestModifier) EXPECT() *RequestModifier_Expecter {
	return &RequestModifier_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: r
func (_m *RequestModifier) Execute(r *http.Request) {
	_m.Called(r)
}

// RequestModifier_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type RequestModifier_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - r *http.Request
func (_e *RequestModifier_Expecter) Execute(r interface{}) *RequestModifier_Execute_Call {
	return &RequestModifier_Execute_Call{Call: _e.mock.On("Execute", r)}
}

func (_c *RequestModifier_Execute_Call) Run(run func(r *http.Request)) *RequestModifier_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*http.Request))
	})
	return _c
}

func (_c *RequestModifier_Execute_Call) Return() *RequestModifier_Execute_Call {
	_c.Call.Return()
	return _c
}

func (_c *RequestModifier_Execute_Call) RunAndReturn(run func(*http.Request)) *RequestModifier_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewRequestModifier creates a new instance of RequestModifier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequestModifier(t interface {
	mock.TestingT
	Cleanup(func())
}) *RequestModifier {
	mock := &RequestModifier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
