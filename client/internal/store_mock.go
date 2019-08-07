// Code generated by MockGen. DO NOT EDIT.
// Source: vendor/github.com/vacp2p/mvds/store/messagestore.go

// Package internal is a generated GoMock package.
package internal

import (
	gomock "github.com/golang/mock/gomock"
	protobuf "github.com/vacp2p/mvds/protobuf"
	state "github.com/vacp2p/mvds/state"
	reflect "reflect"
)

// MockMessageStore is a mock of MessageStore interface
type MockMessageStore struct {
	ctrl     *gomock.Controller
	recorder *MockMessageStoreMockRecorder
}

// MockMessageStoreMockRecorder is the mock recorder for MockMessageStore
type MockMessageStoreMockRecorder struct {
	mock *MockMessageStore
}

// NewMockMessageStore creates a new mock instance
func NewMockMessageStore(ctrl *gomock.Controller) *MockMessageStore {
	mock := &MockMessageStore{ctrl: ctrl}
	mock.recorder = &MockMessageStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessageStore) EXPECT() *MockMessageStoreMockRecorder {
	return m.recorder
}

// Has mocks base method
func (m *MockMessageStore) Has(id state.MessageID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Has indicates an expected call of Has
func (mr *MockMessageStoreMockRecorder) Has(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockMessageStore)(nil).Has), id)
}

// Get mocks base method
func (m *MockMessageStore) Get(id state.MessageID) (*protobuf.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*protobuf.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockMessageStoreMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMessageStore)(nil).Get), id)
}

// Add mocks base method
func (m *MockMessageStore) Add(message *protobuf.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockMessageStoreMockRecorder) Add(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockMessageStore)(nil).Add), message)
}
