/*******************************************************************************
 * Copyright 2020 Samsung Electronics All Rights Reserved.
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
// Source: cipher.go

// Package mock_cipher is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	cipher "restinterface/cipher"
)

// MockIEdgeCipherer is a mock of IEdgeCipherer interface
type MockIEdgeCipherer struct {
	ctrl     *gomock.Controller
	recorder *MockIEdgeCiphererMockRecorder
}

// MockIEdgeCiphererMockRecorder is the mock recorder for MockIEdgeCipherer
type MockIEdgeCiphererMockRecorder struct {
	mock *MockIEdgeCipherer
}

// NewMockIEdgeCipherer creates a new mock instance
func NewMockIEdgeCipherer(ctrl *gomock.Controller) *MockIEdgeCipherer {
	mock := &MockIEdgeCipherer{ctrl: ctrl}
	mock.recorder = &MockIEdgeCiphererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIEdgeCipherer) EXPECT() *MockIEdgeCiphererMockRecorder {
	return m.recorder
}

// EncryptByte mocks base method
func (m *MockIEdgeCipherer) EncryptByte(byteData []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncryptByte", byteData)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EncryptByte indicates an expected call of EncryptByte
func (mr *MockIEdgeCiphererMockRecorder) EncryptByte(byteData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptByte", reflect.TypeOf((*MockIEdgeCipherer)(nil).EncryptByte), byteData)
}

// EncryptJSONToByte mocks base method
func (m *MockIEdgeCipherer) EncryptJSONToByte(jsonMap map[string]interface{}) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncryptJSONToByte", jsonMap)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EncryptJSONToByte indicates an expected call of EncryptJSONToByte
func (mr *MockIEdgeCiphererMockRecorder) EncryptJSONToByte(jsonMap interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncryptJSONToByte", reflect.TypeOf((*MockIEdgeCipherer)(nil).EncryptJSONToByte), jsonMap)
}

// DecryptByte mocks base method
func (m *MockIEdgeCipherer) DecryptByte(byteData []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecryptByte", byteData)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecryptByte indicates an expected call of DecryptByte
func (mr *MockIEdgeCiphererMockRecorder) DecryptByte(byteData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecryptByte", reflect.TypeOf((*MockIEdgeCipherer)(nil).DecryptByte), byteData)
}

// DecryptByteToJSON mocks base method
func (m *MockIEdgeCipherer) DecryptByteToJSON(data []byte) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecryptByteToJSON", data)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecryptByteToJSON indicates an expected call of DecryptByteToJSON
func (mr *MockIEdgeCiphererMockRecorder) DecryptByteToJSON(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecryptByteToJSON", reflect.TypeOf((*MockIEdgeCipherer)(nil).DecryptByteToJSON), data)
}

// MockSetter is a mock of Setter interface
type MockSetter struct {
	ctrl     *gomock.Controller
	recorder *MockSetterMockRecorder
}

// MockSetterMockRecorder is the mock recorder for MockSetter
type MockSetterMockRecorder struct {
	mock *MockSetter
}

// NewMockSetter creates a new mock instance
func NewMockSetter(ctrl *gomock.Controller) *MockSetter {
	mock := &MockSetter{ctrl: ctrl}
	mock.recorder = &MockSetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSetter) EXPECT() *MockSetterMockRecorder {
	return m.recorder
}

// SetCipher mocks base method
func (m *MockSetter) SetCipher(cipher cipher.IEdgeCipherer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCipher", cipher)
}

// SetCipher indicates an expected call of SetCipher
func (mr *MockSetterMockRecorder) SetCipher(cipher interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCipher", reflect.TypeOf((*MockSetter)(nil).SetCipher), cipher)
}
