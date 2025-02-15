// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	"github.com/aws/amazon-ssm-agent/agent/setupcli/managers/registermanager"
	mock "github.com/stretchr/testify/mock"
)

// IRegisterManager is an autogenerated mock type for the IRegisterManager type
type IRegisterManager struct {
	mock.Mock
}

// RegisterAgent provides a mock function with given fields: region, role, tags
func (_m *IRegisterManager) RegisterAgent(registerAgentInpModel *registermanager.RegisterAgentInputModel) error {
	ret := _m.Called(registerAgentInpModel)

	var r0 error
	if rf, ok := ret.Get(0).(func(inp1 *registermanager.RegisterAgentInputModel) error); ok {
		r0 = rf(registerAgentInpModel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
