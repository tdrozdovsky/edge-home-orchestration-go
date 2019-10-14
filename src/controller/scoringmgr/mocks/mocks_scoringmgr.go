/*******************************************************************************
 * Copyright 2019 Samsung Electronics All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *******************************************************************************/

// Code generated by MockGen. DO NOT EDIT.
// Source: scoringmgr.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockScoring is a mock of Scoring interface
type MockScoring struct {
	ctrl     *gomock.Controller
	recorder *MockScoringMockRecorder
}

// MockScoringMockRecorder is the mock recorder for MockScoring
type MockScoringMockRecorder struct {
	mock *MockScoring
}

// NewMockScoring creates a new mock instance
func NewMockScoring(ctrl *gomock.Controller) *MockScoring {
	mock := &MockScoring{ctrl: ctrl}
	mock.recorder = &MockScoringMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScoring) EXPECT() *MockScoringMockRecorder {
	return m.recorder
}

// GetScore mocks base method
func (m *MockScoring) GetScore(ID string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScore", ID)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScore indicates an expected call of GetScore
func (mr *MockScoringMockRecorder) GetScore(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScore", reflect.TypeOf((*MockScoring)(nil).GetScore), ID)
}