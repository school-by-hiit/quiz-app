// Code generated by mockery v2.20.0. DO NOT EDIT.

package domain

import mock "github.com/stretchr/testify/mock"

// MockAuthRepository is an autogenerated mock type for the AuthRepository type
type MockAuthRepository struct {
	mock.Mock
}

type MockAuthRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAuthRepository) EXPECT() *MockAuthRepository_Expecter {
	return &MockAuthRepository_Expecter{mock: &_m.Mock}
}

// CacheToken provides a mock function with given fields: token
func (_m *MockAuthRepository) CacheToken(token *AccessToken) error {
	ret := _m.Called(token)

	var r0 error
	if rf, ok := ret.Get(0).(func(*AccessToken) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAuthRepository_CacheToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CacheToken'
type MockAuthRepository_CacheToken_Call struct {
	*mock.Call
}

// CacheToken is a helper method to define mock.On call
//   - token *AccessToken
func (_e *MockAuthRepository_Expecter) CacheToken(token interface{}) *MockAuthRepository_CacheToken_Call {
	return &MockAuthRepository_CacheToken_Call{Call: _e.mock.On("CacheToken", token)}
}

func (_c *MockAuthRepository_CacheToken_Call) Run(run func(token *AccessToken)) *MockAuthRepository_CacheToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*AccessToken))
	})
	return _c
}

func (_c *MockAuthRepository_CacheToken_Call) Return(_a0 error) *MockAuthRepository_CacheToken_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAuthRepository_CacheToken_Call) RunAndReturn(run func(*AccessToken) error) *MockAuthRepository_CacheToken_Call {
	_c.Call.Return(run)
	return _c
}

// FindTokenByTokenStr provides a mock function with given fields: tokenStr
func (_m *MockAuthRepository) FindTokenByTokenStr(tokenStr string) (*AccessToken, error) {
	ret := _m.Called(tokenStr)

	var r0 *AccessToken
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*AccessToken, error)); ok {
		return rf(tokenStr)
	}
	if rf, ok := ret.Get(0).(func(string) *AccessToken); ok {
		r0 = rf(tokenStr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AccessToken)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthRepository_FindTokenByTokenStr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindTokenByTokenStr'
type MockAuthRepository_FindTokenByTokenStr_Call struct {
	*mock.Call
}

// FindTokenByTokenStr is a helper method to define mock.On call
//   - tokenStr string
func (_e *MockAuthRepository_Expecter) FindTokenByTokenStr(tokenStr interface{}) *MockAuthRepository_FindTokenByTokenStr_Call {
	return &MockAuthRepository_FindTokenByTokenStr_Call{Call: _e.mock.On("FindTokenByTokenStr", tokenStr)}
}

func (_c *MockAuthRepository_FindTokenByTokenStr_Call) Run(run func(tokenStr string)) *MockAuthRepository_FindTokenByTokenStr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockAuthRepository_FindTokenByTokenStr_Call) Return(_a0 *AccessToken, _a1 error) *MockAuthRepository_FindTokenByTokenStr_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthRepository_FindTokenByTokenStr_Call) RunAndReturn(run func(string) (*AccessToken, error)) *MockAuthRepository_FindTokenByTokenStr_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockAuthRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAuthRepository creates a new instance of MockAuthRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAuthRepository(t mockConstructorTestingTNewMockAuthRepository) *MockAuthRepository {
	mock := &MockAuthRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
