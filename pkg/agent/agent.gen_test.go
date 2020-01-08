// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kcarretto/paragon/pkg/agent (interfaces: Sender,Receiver)

// Package agent_test is a generated GoMock package.
package agent_test

import (
	gomock "github.com/golang/mock/gomock"
	agent "github.com/kcarretto/paragon/pkg/agent"
	reflect "reflect"
)

// MockSender is a mock of Sender interface
type MockSender struct {
	ctrl     *gomock.Controller
	recorder *MockSenderMockRecorder
}

// MockSenderMockRecorder is the mock recorder for MockSender
type MockSenderMockRecorder struct {
	mock *MockSender
}

// NewMockSender creates a new mock instance
func NewMockSender(ctrl *gomock.Controller) *MockSender {
	mock := &MockSender{ctrl: ctrl}
	mock.recorder = &MockSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSender) EXPECT() *MockSenderMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockSender) Send(arg0 agent.ServerMessageWriter, arg1 agent.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockSenderMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockSender)(nil).Send), arg0, arg1)
}

// MockReceiver is a mock of Receiver interface
type MockReceiver struct {
	ctrl     *gomock.Controller
	recorder *MockReceiverMockRecorder
}

// MockReceiverMockRecorder is the mock recorder for MockReceiver
type MockReceiverMockRecorder struct {
	mock *MockReceiver
}

// NewMockReceiver creates a new mock instance
func NewMockReceiver(ctrl *gomock.Controller) *MockReceiver {
	mock := &MockReceiver{ctrl: ctrl}
	mock.recorder = &MockReceiverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReceiver) EXPECT() *MockReceiverMockRecorder {
	return m.recorder
}

// Receive mocks base method
func (m *MockReceiver) Receive(arg0 agent.MessageWriter, arg1 agent.ServerMessage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Receive", arg0, arg1)
}

// Receive indicates an expected call of Receive
func (mr *MockReceiverMockRecorder) Receive(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockReceiver)(nil).Receive), arg0, arg1)
}