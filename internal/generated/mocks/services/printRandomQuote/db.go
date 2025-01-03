// Code generated by mockery. DO NOT EDIT.

package printRandomQuote_test

import (
	context "context"

	models "github.com/nijeti/cinema-keeper/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// MockDb is an autogenerated mock type for the db type
type MockDb struct {
	mock.Mock
}

type MockDb_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDb) EXPECT() *MockDb_Expecter {
	return &MockDb_Expecter{mock: &_m.Mock}
}

// GetRandomQuoteInGuild provides a mock function with given fields: ctx, guildID
func (_m *MockDb) GetRandomQuoteInGuild(ctx context.Context, guildID models.ID) (*models.Quote, error) {
	ret := _m.Called(ctx, guildID)

	if len(ret) == 0 {
		panic("no return value specified for GetRandomQuoteInGuild")
	}

	var r0 *models.Quote
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.ID) (*models.Quote, error)); ok {
		return rf(ctx, guildID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.ID) *models.Quote); ok {
		r0 = rf(ctx, guildID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Quote)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.ID) error); ok {
		r1 = rf(ctx, guildID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDb_GetRandomQuoteInGuild_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRandomQuoteInGuild'
type MockDb_GetRandomQuoteInGuild_Call struct {
	*mock.Call
}

// GetRandomQuoteInGuild is a helper method to define mock.On call
//   - ctx context.Context
//   - guildID models.ID
func (_e *MockDb_Expecter) GetRandomQuoteInGuild(ctx interface{}, guildID interface{}) *MockDb_GetRandomQuoteInGuild_Call {
	return &MockDb_GetRandomQuoteInGuild_Call{Call: _e.mock.On("GetRandomQuoteInGuild", ctx, guildID)}
}

func (_c *MockDb_GetRandomQuoteInGuild_Call) Run(run func(ctx context.Context, guildID models.ID)) *MockDb_GetRandomQuoteInGuild_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.ID))
	})
	return _c
}

func (_c *MockDb_GetRandomQuoteInGuild_Call) Return(_a0 *models.Quote, _a1 error) *MockDb_GetRandomQuoteInGuild_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDb_GetRandomQuoteInGuild_Call) RunAndReturn(run func(context.Context, models.ID) (*models.Quote, error)) *MockDb_GetRandomQuoteInGuild_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDb creates a new instance of MockDb. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDb(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDb {
	mock := &MockDb{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
