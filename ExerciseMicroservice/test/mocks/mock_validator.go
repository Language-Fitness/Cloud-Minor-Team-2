package mocks

import "github.com/stretchr/testify/mock"

type MockValidator struct {
	mock.Mock
}

func (m *MockValidator) ClearErrors() {
}

func (m *MockValidator) Validate(input interface{}, arr []string, name string) {

}

func (m *MockValidator) GetErrors() []string {
	args := m.Called()
	return args.Get(0).([]string)
}
