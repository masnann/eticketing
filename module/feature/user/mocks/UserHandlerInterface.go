// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	fiber "github.com/gofiber/fiber/v2"
	mock "github.com/stretchr/testify/mock"
)

// UserHandlerInterface is an autogenerated mock type for the UserHandlerInterface type
type UserHandlerInterface struct {
	mock.Mock
}

// GetUserByID provides a mock function with given fields: c
func (_m *UserHandlerInterface) GetUserByID(c *fiber.Ctx) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserHandlerInterface creates a new instance of UserHandlerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserHandlerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserHandlerInterface {
	mock := &UserHandlerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}