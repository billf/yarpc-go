// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: go.uber.org/yarpc/transport (interfaces: UnaryOutbound,OnewayOutbound)

package transporttest

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	transport "go.uber.org/yarpc/transport"
)

// Mock of UnaryOutbound interface
type MockUnaryOutbound struct {
	ctrl     *gomock.Controller
	recorder *_MockUnaryOutboundRecorder
}

// Recorder for MockUnaryOutbound (not exported)
type _MockUnaryOutboundRecorder struct {
	mock *MockUnaryOutbound
}

func NewMockUnaryOutbound(ctrl *gomock.Controller) *MockUnaryOutbound {
	mock := &MockUnaryOutbound{ctrl: ctrl}
	mock.recorder = &_MockUnaryOutboundRecorder{mock}
	return mock
}

func (_m *MockUnaryOutbound) EXPECT() *_MockUnaryOutboundRecorder {
	return _m.recorder
}

func (_m *MockUnaryOutbound) Call(_param0 context.Context, _param1 *transport.Request) (*transport.Response, error) {
	ret := _m.ctrl.Call(_m, "Call", _param0, _param1)
	ret0, _ := ret[0].(*transport.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockUnaryOutboundRecorder) Call(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Call", arg0, arg1)
}

func (_m *MockUnaryOutbound) Start(_param0 transport.Deps) error {
	ret := _m.ctrl.Call(_m, "Start", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockUnaryOutboundRecorder) Start(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start", arg0)
}

func (_m *MockUnaryOutbound) Stop() error {
	ret := _m.ctrl.Call(_m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockUnaryOutboundRecorder) Stop() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop")
}

// Mock of OnewayOutbound interface
type MockOnewayOutbound struct {
	ctrl     *gomock.Controller
	recorder *_MockOnewayOutboundRecorder
}

// Recorder for MockOnewayOutbound (not exported)
type _MockOnewayOutboundRecorder struct {
	mock *MockOnewayOutbound
}

func NewMockOnewayOutbound(ctrl *gomock.Controller) *MockOnewayOutbound {
	mock := &MockOnewayOutbound{ctrl: ctrl}
	mock.recorder = &_MockOnewayOutboundRecorder{mock}
	return mock
}

func (_m *MockOnewayOutbound) EXPECT() *_MockOnewayOutboundRecorder {
	return _m.recorder
}

func (_m *MockOnewayOutbound) CallOneway(_param0 context.Context, _param1 *transport.Request) (transport.Ack, error) {
	ret := _m.ctrl.Call(_m, "CallOneway", _param0, _param1)
	ret0, _ := ret[0].(transport.Ack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockOnewayOutboundRecorder) CallOneway(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CallOneway", arg0, arg1)
}

func (_m *MockOnewayOutbound) Start(_param0 transport.Deps) error {
	ret := _m.ctrl.Call(_m, "Start", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockOnewayOutboundRecorder) Start(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start", arg0)
}

func (_m *MockOnewayOutbound) Stop() error {
	ret := _m.ctrl.Call(_m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockOnewayOutboundRecorder) Stop() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop")
}
