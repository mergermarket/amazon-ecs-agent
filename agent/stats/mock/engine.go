// Copyright 2015 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/amazon-ecs-agent/agent/stats (interfaces: Engine)

package mock_stats

import (
	gomock "code.google.com/p/gomock/gomock"
	ecstcs "github.com/aws/amazon-ecs-agent/agent/tcs/model/ecstcs"
)

// Mock of Engine interface
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *_MockEngineRecorder
}

// Recorder for MockEngine (not exported)
type _MockEngineRecorder struct {
	mock *MockEngine
}

func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &_MockEngineRecorder{mock}
	return mock
}

func (_m *MockEngine) EXPECT() *_MockEngineRecorder {
	return _m.recorder
}

func (_m *MockEngine) GetInstanceMetrics() (*ecstcs.MetricsMetadata, []*ecstcs.TaskMetric, error) {
	ret := _m.ctrl.Call(_m, "GetInstanceMetrics")
	ret0, _ := ret[0].(*ecstcs.MetricsMetadata)
	ret1, _ := ret[1].([]*ecstcs.TaskMetric)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockEngineRecorder) GetInstanceMetrics() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetInstanceMetrics")
}
