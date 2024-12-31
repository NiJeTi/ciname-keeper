// Code generated by mockery. DO NOT EDIT.

package quote_test

import (
	context "context"

	discordgo "github.com/bwmarrin/discordgo"
	mock "github.com/stretchr/testify/mock"

	models "github.com/nijeti/cinema-keeper/internal/models"
)

// MockListQuotes is an autogenerated mock type for the listQuotes type
type MockListQuotes struct {
	mock.Mock
}

type MockListQuotes_Expecter struct {
	mock *mock.Mock
}

func (_m *MockListQuotes) EXPECT() *MockListQuotes_Expecter {
	return &MockListQuotes_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: ctx, i, authorID
func (_m *MockListQuotes) Exec(ctx context.Context, i *discordgo.Interaction, authorID models.ID) error {
	ret := _m.Called(ctx, i, authorID)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *discordgo.Interaction, models.ID) error); ok {
		r0 = rf(ctx, i, authorID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockListQuotes_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockListQuotes_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - i *discordgo.Interaction
//   - authorID models.ID
func (_e *MockListQuotes_Expecter) Exec(ctx interface{}, i interface{}, authorID interface{}) *MockListQuotes_Exec_Call {
	return &MockListQuotes_Exec_Call{Call: _e.mock.On("Exec", ctx, i, authorID)}
}

func (_c *MockListQuotes_Exec_Call) Run(run func(ctx context.Context, i *discordgo.Interaction, authorID models.ID)) *MockListQuotes_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*discordgo.Interaction), args[2].(models.ID))
	})
	return _c
}

func (_c *MockListQuotes_Exec_Call) Return(_a0 error) *MockListQuotes_Exec_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockListQuotes_Exec_Call) RunAndReturn(run func(context.Context, *discordgo.Interaction, models.ID) error) *MockListQuotes_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockListQuotes creates a new instance of MockListQuotes. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockListQuotes(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockListQuotes {
	mock := &MockListQuotes{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
