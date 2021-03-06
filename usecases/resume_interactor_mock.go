// Code generated by MockGen. DO NOT EDIT.
// Source: resume_interactor.go

// Package usecases is a generated GoMock package.
package usecases

import (
	domain "github.com/antonve/portfolio-api/domain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockResumeInteractor is a mock of ResumeInteractor interface
type MockResumeInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockResumeInteractorMockRecorder
}

// MockResumeInteractorMockRecorder is the mock recorder for MockResumeInteractor
type MockResumeInteractorMockRecorder struct {
	mock *MockResumeInteractor
}

// NewMockResumeInteractor creates a new mock instance
func NewMockResumeInteractor(ctrl *gomock.Controller) *MockResumeInteractor {
	mock := &MockResumeInteractor{ctrl: ctrl}
	mock.recorder = &MockResumeInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResumeInteractor) EXPECT() *MockResumeInteractorMockRecorder {
	return m.recorder
}

// TrackedFind mocks base method
func (m *MockResumeInteractor) TrackedFind(slug, ipAddress, userAgent string) (*domain.Resume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrackedFind", slug, ipAddress, userAgent)
	ret0, _ := ret[0].(*domain.Resume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TrackedFind indicates an expected call of TrackedFind
func (mr *MockResumeInteractorMockRecorder) TrackedFind(slug, ipAddress, userAgent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrackedFind", reflect.TypeOf((*MockResumeInteractor)(nil).TrackedFind), slug, ipAddress, userAgent)
}
