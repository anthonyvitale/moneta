// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anthonyvitale/moneta (interfaces: S3API)
//
// Generated by this command:
//
//	mockgen -destination=mocks/mock_store.auto.go -package mocks github.com/anthonyvitale/moneta S3API
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	gomock "go.uber.org/mock/gomock"
)

// MockS3API is a mock of S3API interface.
type MockS3API struct {
	ctrl     *gomock.Controller
	recorder *MockS3APIMockRecorder
}

// MockS3APIMockRecorder is the mock recorder for MockS3API.
type MockS3APIMockRecorder struct {
	mock *MockS3API
}

// NewMockS3API creates a new mock instance.
func NewMockS3API(ctrl *gomock.Controller) *MockS3API {
	mock := &MockS3API{ctrl: ctrl}
	mock.recorder = &MockS3APIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3API) EXPECT() *MockS3APIMockRecorder {
	return m.recorder
}

// HeadBucket mocks base method.
func (m *MockS3API) HeadBucket(arg0 context.Context, arg1 *s3.HeadBucketInput, arg2 ...func(*s3.Options)) (*s3.HeadBucketOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HeadBucket", varargs...)
	ret0, _ := ret[0].(*s3.HeadBucketOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeadBucket indicates an expected call of HeadBucket.
func (mr *MockS3APIMockRecorder) HeadBucket(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadBucket", reflect.TypeOf((*MockS3API)(nil).HeadBucket), varargs...)
}

// PutObject mocks base method.
func (m *MockS3API) PutObject(arg0 context.Context, arg1 *s3.PutObjectInput, arg2 ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutObject", varargs...)
	ret0, _ := ret[0].(*s3.PutObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObject indicates an expected call of PutObject.
func (mr *MockS3APIMockRecorder) PutObject(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockS3API)(nil).PutObject), varargs...)
}
