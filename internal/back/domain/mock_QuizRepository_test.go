// Code generated by mockery v2.20.0. DO NOT EDIT.

package domain

import (
	context "context"

	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
)

// MockQuizRepository is an autogenerated mock type for the QuizRepository type
type MockQuizRepository struct {
	mock.Mock
}

type MockQuizRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockQuizRepository) EXPECT() *MockQuizRepository_Expecter {
	return &MockQuizRepository_Expecter{mock: &_m.Mock}
}

// ActivateOnlyVersion provides a mock function with given fields: ctx, filename, version
func (_m *MockQuizRepository) ActivateOnlyVersion(ctx context.Context, filename string, version int) error {
	ret := _m.Called(ctx, filename, version)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) error); ok {
		r0 = rf(ctx, filename, version)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQuizRepository_ActivateOnlyVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ActivateOnlyVersion'
type MockQuizRepository_ActivateOnlyVersion_Call struct {
	*mock.Call
}

// ActivateOnlyVersion is a helper method to define mock.On call
//   - ctx context.Context
//   - filename string
//   - version int
func (_e *MockQuizRepository_Expecter) ActivateOnlyVersion(ctx interface{}, filename interface{}, version interface{}) *MockQuizRepository_ActivateOnlyVersion_Call {
	return &MockQuizRepository_ActivateOnlyVersion_Call{Call: _e.mock.On("ActivateOnlyVersion", ctx, filename, version)}
}

func (_c *MockQuizRepository_ActivateOnlyVersion_Call) Run(run func(ctx context.Context, filename string, version int)) *MockQuizRepository_ActivateOnlyVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int))
	})
	return _c
}

func (_c *MockQuizRepository_ActivateOnlyVersion_Call) Return(_a0 error) *MockQuizRepository_ActivateOnlyVersion_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQuizRepository_ActivateOnlyVersion_Call) RunAndReturn(run func(context.Context, string, int) error) *MockQuizRepository_ActivateOnlyVersion_Call {
	_c.Call.Return(run)
	return _c
}

// AddSessionAnswer provides a mock function with given fields: ctx, sessionUuid, userId, questionSha1, answerSha1, checked
func (_m *MockQuizRepository) AddSessionAnswer(ctx context.Context, sessionUuid uuid.UUID, userId string, questionSha1 string, answerSha1 string, checked bool) error {
	ret := _m.Called(ctx, sessionUuid, userId, questionSha1, answerSha1, checked)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string, string, string, bool) error); ok {
		r0 = rf(ctx, sessionUuid, userId, questionSha1, answerSha1, checked)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQuizRepository_AddSessionAnswer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddSessionAnswer'
type MockQuizRepository_AddSessionAnswer_Call struct {
	*mock.Call
}

// AddSessionAnswer is a helper method to define mock.On call
//   - ctx context.Context
//   - sessionUuid uuid.UUID
//   - userId string
//   - questionSha1 string
//   - answerSha1 string
//   - checked bool
func (_e *MockQuizRepository_Expecter) AddSessionAnswer(ctx interface{}, sessionUuid interface{}, userId interface{}, questionSha1 interface{}, answerSha1 interface{}, checked interface{}) *MockQuizRepository_AddSessionAnswer_Call {
	return &MockQuizRepository_AddSessionAnswer_Call{Call: _e.mock.On("AddSessionAnswer", ctx, sessionUuid, userId, questionSha1, answerSha1, checked)}
}

func (_c *MockQuizRepository_AddSessionAnswer_Call) Run(run func(ctx context.Context, sessionUuid uuid.UUID, userId string, questionSha1 string, answerSha1 string, checked bool)) *MockQuizRepository_AddSessionAnswer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(string), args[3].(string), args[4].(string), args[5].(bool))
	})
	return _c
}

func (_c *MockQuizRepository_AddSessionAnswer_Call) Return(_a0 error) *MockQuizRepository_AddSessionAnswer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQuizRepository_AddSessionAnswer_Call) RunAndReturn(run func(context.Context, uuid.UUID, string, string, string, bool) error) *MockQuizRepository_AddSessionAnswer_Call {
	_c.Call.Return(run)
	return _c
}

// CountAllActive provides a mock function with given fields: ctx
func (_m *MockQuizRepository) CountAllActive(ctx context.Context) (uint32, error) {
	ret := _m.Called(ctx)

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint32, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint32); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuizRepository_CountAllActive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountAllActive'
type MockQuizRepository_CountAllActive_Call struct {
	*mock.Call
}

// CountAllActive is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockQuizRepository_Expecter) CountAllActive(ctx interface{}) *MockQuizRepository_CountAllActive_Call {
	return &MockQuizRepository_CountAllActive_Call{Call: _e.mock.On("CountAllActive", ctx)}
}

func (_c *MockQuizRepository_CountAllActive_Call) Run(run func(ctx context.Context)) *MockQuizRepository_CountAllActive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockQuizRepository_CountAllActive_Call) Return(_a0 uint32, _a1 error) *MockQuizRepository_CountAllActive_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuizRepository_CountAllActive_Call) RunAndReturn(run func(context.Context) (uint32, error)) *MockQuizRepository_CountAllActive_Call {
	_c.Call.Return(run)
	return _c
}

// CountAllSessions provides a mock function with given fields: ctx, quizActive, userId
func (_m *MockQuizRepository) CountAllSessions(ctx context.Context, quizActive bool, userId string) (uint32, error) {
	ret := _m.Called(ctx, quizActive, userId)

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bool, string) (uint32, error)); ok {
		return rf(ctx, quizActive, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bool, string) uint32); ok {
		r0 = rf(ctx, quizActive, userId)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, bool, string) error); ok {
		r1 = rf(ctx, quizActive, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuizRepository_CountAllSessions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountAllSessions'
type MockQuizRepository_CountAllSessions_Call struct {
	*mock.Call
}

// CountAllSessions is a helper method to define mock.On call
//   - ctx context.Context
//   - quizActive bool
//   - userId string
func (_e *MockQuizRepository_Expecter) CountAllSessions(ctx interface{}, quizActive interface{}, userId interface{}) *MockQuizRepository_CountAllSessions_Call {
	return &MockQuizRepository_CountAllSessions_Call{Call: _e.mock.On("CountAllSessions", ctx, quizActive, userId)}
}

func (_c *MockQuizRepository_CountAllSessions_Call) Run(run func(ctx context.Context, quizActive bool, userId string)) *MockQuizRepository_CountAllSessions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(bool), args[2].(string))
	})
	return _c
}

func (_c *MockQuizRepository_CountAllSessions_Call) Return(_a0 uint32, _a1 error) *MockQuizRepository_CountAllSessions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuizRepository_CountAllSessions_Call) RunAndReturn(run func(context.Context, bool, string) (uint32, error)) *MockQuizRepository_CountAllSessions_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx, quiz
func (_m *MockQuizRepository) Create(ctx context.Context, quiz *Quiz) error {
	ret := _m.Called(ctx, quiz)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *Quiz) error); ok {
		r0 = rf(ctx, quiz)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQuizRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockQuizRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - quiz *Quiz
func (_e *MockQuizRepository_Expecter) Create(ctx interface{}, quiz interface{}) *MockQuizRepository_Create_Call {
	return &MockQuizRepository_Create_Call{Call: _e.mock.On("Create", ctx, quiz)}
}

func (_c *MockQuizRepository_Create_Call) Run(run func(ctx context.Context, quiz *Quiz)) *MockQuizRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Quiz))
	})
	return _c
}

func (_c *MockQuizRepository_Create_Call) Return(_a0 error) *MockQuizRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQuizRepository_Create_Call) RunAndReturn(run func(context.Context, *Quiz) error) *MockQuizRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// FindAllActive provides a mock function with given fields: ctx, limit, offset
func (_m *MockQuizRepository) FindAllActive(ctx context.Context, limit uint16, offset uint16) ([]*Quiz, error) {
	ret := _m.Called(ctx, limit, offset)

	var r0 []*Quiz
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint16, uint16) ([]*Quiz, error)); ok {
		return rf(ctx, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint16, uint16) []*Quiz); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Quiz)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint16, uint16) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuizRepository_FindAllActive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAllActive'
type MockQuizRepository_FindAllActive_Call struct {
	*mock.Call
}

// FindAllActive is a helper method to define mock.On call
//   - ctx context.Context
//   - limit uint16
//   - offset uint16
func (_e *MockQuizRepository_Expecter) FindAllActive(ctx interface{}, limit interface{}, offset interface{}) *MockQuizRepository_FindAllActive_Call {
	return &MockQuizRepository_FindAllActive_Call{Call: _e.mock.On("FindAllActive", ctx, limit, offset)}
}

func (_c *MockQuizRepository_FindAllActive_Call) Run(run func(ctx context.Context, limit uint16, offset uint16)) *MockQuizRepository_FindAllActive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint16), args[2].(uint16))
	})
	return _c
}

func (_c *MockQuizRepository_FindAllActive_Call) Return(_a0 []*Quiz, _a1 error) *MockQuizRepository_FindAllActive_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuizRepository_FindAllActive_Call) RunAndReturn(run func(context.Context, uint16, uint16) ([]*Quiz, error)) *MockQuizRepository_FindAllActive_Call {
	_c.Call.Return(run)
	return _c
}

// FindAllSessions provides a mock function with given fields: ctx, quizActive, userId, limit, offset
func (_m *MockQuizRepository) FindAllSessions(ctx context.Context, quizActive bool, userId string, limit uint16, offset uint16) ([]*Session, error) {
	ret := _m.Called(ctx, quizActive, userId, limit, offset)

	var r0 []*Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, bool, string, uint16, uint16) ([]*Session, error)); ok {
		return rf(ctx, quizActive, userId, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, bool, string, uint16, uint16) []*Session); ok {
		r0 = rf(ctx, quizActive, userId, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Session)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, bool, string, uint16, uint16) error); ok {
		r1 = rf(ctx, quizActive, userId, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuizRepository_FindAllSessions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAllSessions'
type MockQuizRepository_FindAllSessions_Call struct {
	*mock.Call
}

// FindAllSessions is a helper method to define mock.On call
//   - ctx context.Context
//   - quizActive bool
//   - userId string
//   - limit uint16
//   - offset uint16
func (_e *MockQuizRepository_Expecter) FindAllSessions(ctx interface{}, quizActive interface{}, userId interface{}, limit interface{}, offset interface{}) *MockQuizRepository_FindAllSessions_Call {
	return &MockQuizRepository_FindAllSessions_Call{Call: _e.mock.On("FindAllSessions", ctx, quizActive, userId, limit, offset)}
}

func (_c *MockQuizRepository_FindAllSessions_Call) Run(run func(ctx context.Context, quizActive bool, userId string, limit uint16, offset uint16)) *MockQuizRepository_FindAllSessions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(bool), args[2].(string), args[3].(uint16), args[4].(uint16))
	})
	return _c
}

func (_c *MockQuizRepository_FindAllSessions_Call) Return(_a0 []*Session, _a1 error) *MockQuizRepository_FindAllSessions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuizRepository_FindAllSessions_Call) RunAndReturn(run func(context.Context, bool, string, uint16, uint16) ([]*Session, error)) *MockQuizRepository_FindAllSessions_Call {
	_c.Call.Return(run)
	return _c
}

// FindBySha1 provides a mock function with given fields: ctx, sha1
func (_m *MockQuizRepository) FindBySha1(ctx context.Context, sha1 string) (*Quiz, error) {
	ret := _m.Called(ctx, sha1)

	var r0 *Quiz
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*Quiz, error)); ok {
		return rf(ctx, sha1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *Quiz); ok {
		r0 = rf(ctx, sha1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Quiz)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, sha1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuizRepository_FindBySha1_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindBySha1'
type MockQuizRepository_FindBySha1_Call struct {
	*mock.Call
}

// FindBySha1 is a helper method to define mock.On call
//   - ctx context.Context
//   - sha1 string
func (_e *MockQuizRepository_Expecter) FindBySha1(ctx interface{}, sha1 interface{}) *MockQuizRepository_FindBySha1_Call {
	return &MockQuizRepository_FindBySha1_Call{Call: _e.mock.On("FindBySha1", ctx, sha1)}
}

func (_c *MockQuizRepository_FindBySha1_Call) Run(run func(ctx context.Context, sha1 string)) *MockQuizRepository_FindBySha1_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockQuizRepository_FindBySha1_Call) Return(_a0 *Quiz, _a1 error) *MockQuizRepository_FindBySha1_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuizRepository_FindBySha1_Call) RunAndReturn(run func(context.Context, string) (*Quiz, error)) *MockQuizRepository_FindBySha1_Call {
	_c.Call.Return(run)
	return _c
}

// FindFullBySha1 provides a mock function with given fields: ctx, sha1
func (_m *MockQuizRepository) FindFullBySha1(ctx context.Context, sha1 string) (*Quiz, error) {
	ret := _m.Called(ctx, sha1)

	var r0 *Quiz
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*Quiz, error)); ok {
		return rf(ctx, sha1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *Quiz); ok {
		r0 = rf(ctx, sha1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Quiz)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, sha1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuizRepository_FindFullBySha1_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindFullBySha1'
type MockQuizRepository_FindFullBySha1_Call struct {
	*mock.Call
}

// FindFullBySha1 is a helper method to define mock.On call
//   - ctx context.Context
//   - sha1 string
func (_e *MockQuizRepository_Expecter) FindFullBySha1(ctx interface{}, sha1 interface{}) *MockQuizRepository_FindFullBySha1_Call {
	return &MockQuizRepository_FindFullBySha1_Call{Call: _e.mock.On("FindFullBySha1", ctx, sha1)}
}

func (_c *MockQuizRepository_FindFullBySha1_Call) Run(run func(ctx context.Context, sha1 string)) *MockQuizRepository_FindFullBySha1_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockQuizRepository_FindFullBySha1_Call) Return(_a0 *Quiz, _a1 error) *MockQuizRepository_FindFullBySha1_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuizRepository_FindFullBySha1_Call) RunAndReturn(run func(context.Context, string) (*Quiz, error)) *MockQuizRepository_FindFullBySha1_Call {
	_c.Call.Return(run)
	return _c
}

// FindLatestVersionByFilename provides a mock function with given fields: ctx, filename
func (_m *MockQuizRepository) FindLatestVersionByFilename(ctx context.Context, filename string) (*Quiz, error) {
	ret := _m.Called(ctx, filename)

	var r0 *Quiz
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*Quiz, error)); ok {
		return rf(ctx, filename)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *Quiz); ok {
		r0 = rf(ctx, filename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Quiz)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, filename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuizRepository_FindLatestVersionByFilename_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindLatestVersionByFilename'
type MockQuizRepository_FindLatestVersionByFilename_Call struct {
	*mock.Call
}

// FindLatestVersionByFilename is a helper method to define mock.On call
//   - ctx context.Context
//   - filename string
func (_e *MockQuizRepository_Expecter) FindLatestVersionByFilename(ctx interface{}, filename interface{}) *MockQuizRepository_FindLatestVersionByFilename_Call {
	return &MockQuizRepository_FindLatestVersionByFilename_Call{Call: _e.mock.On("FindLatestVersionByFilename", ctx, filename)}
}

func (_c *MockQuizRepository_FindLatestVersionByFilename_Call) Run(run func(ctx context.Context, filename string)) *MockQuizRepository_FindLatestVersionByFilename_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockQuizRepository_FindLatestVersionByFilename_Call) Return(_a0 *Quiz, _a1 error) *MockQuizRepository_FindLatestVersionByFilename_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuizRepository_FindLatestVersionByFilename_Call) RunAndReturn(run func(context.Context, string) (*Quiz, error)) *MockQuizRepository_FindLatestVersionByFilename_Call {
	_c.Call.Return(run)
	return _c
}

// StartSession provides a mock function with given fields: ctx, userId, quizSha1
func (_m *MockQuizRepository) StartSession(ctx context.Context, userId string, quizSha1 string) (uuid.UUID, error) {
	ret := _m.Called(ctx, userId, quizSha1)

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (uuid.UUID, error)); ok {
		return rf(ctx, userId, quizSha1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) uuid.UUID); ok {
		r0 = rf(ctx, userId, quizSha1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userId, quizSha1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuizRepository_StartSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartSession'
type MockQuizRepository_StartSession_Call struct {
	*mock.Call
}

// StartSession is a helper method to define mock.On call
//   - ctx context.Context
//   - userId string
//   - quizSha1 string
func (_e *MockQuizRepository_Expecter) StartSession(ctx interface{}, userId interface{}, quizSha1 interface{}) *MockQuizRepository_StartSession_Call {
	return &MockQuizRepository_StartSession_Call{Call: _e.mock.On("StartSession", ctx, userId, quizSha1)}
}

func (_c *MockQuizRepository_StartSession_Call) Run(run func(ctx context.Context, userId string, quizSha1 string)) *MockQuizRepository_StartSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockQuizRepository_StartSession_Call) Return(_a0 uuid.UUID, _a1 error) *MockQuizRepository_StartSession_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuizRepository_StartSession_Call) RunAndReturn(run func(context.Context, string, string) (uuid.UUID, error)) *MockQuizRepository_StartSession_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockQuizRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockQuizRepository creates a new instance of MockQuizRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockQuizRepository(t mockConstructorTestingTNewMockQuizRepository) *MockQuizRepository {
	mock := &MockQuizRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
