package mocks

import "github.com/stretchr/testify/mock"

type MockUserProvider struct {
	mock.Mock
}

func (m *MockUserProvider) GetUserInfo(token string, userId string) (string, error) {
	args := m.Called(token, userId)
	return args.Get(0).(string), args.Error(1)
}
