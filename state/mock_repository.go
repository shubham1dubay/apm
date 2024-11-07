// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: state/repository.go

// Package state is a generated GoMock package.
package state

import (
	reflect "reflect"

	types "github.com/shubham1dubay/apm/types"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetPath mocks base method.
func (m *MockRepository) GetPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPath indicates an expected call of GetPath.
func (mr *MockRepositoryMockRecorder) GetPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPath", reflect.TypeOf((*MockRepository)(nil).GetPath))
}

// GetSubnet mocks base method.
func (m *MockRepository) GetSubnet(name string) (Definition[types.Subnet], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnet", name)
	ret0, _ := ret[0].(Definition[types.Subnet])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnet indicates an expected call of GetSubnet.
func (mr *MockRepositoryMockRecorder) GetSubnet(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnet", reflect.TypeOf((*MockRepository)(nil).GetSubnet), name)
}

// GetVM mocks base method.
func (m *MockRepository) GetVM(name string) (Definition[types.VM], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVM", name)
	ret0, _ := ret[0].(Definition[types.VM])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVM indicates an expected call of GetVM.
func (mr *MockRepositoryMockRecorder) GetVM(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVM", reflect.TypeOf((*MockRepository)(nil).GetVM), name)
}
