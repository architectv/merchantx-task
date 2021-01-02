// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/architectv/merchantx-task/pkg/repository (interfaces: Offer)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	model "github.com/architectv/merchantx-task/pkg/model"
	gomock "github.com/golang/mock/gomock"
)

// MockOffer is a mock of Offer interface.
type MockOffer struct {
	ctrl     *gomock.Controller
	recorder *MockOfferMockRecorder
}

// MockOfferMockRecorder is the mock recorder for MockOffer.
type MockOfferMockRecorder struct {
	mock *MockOffer
}

// NewMockOffer creates a new mock instance.
func NewMockOffer(ctrl *gomock.Controller) *MockOffer {
	mock := &MockOffer{ctrl: ctrl}
	mock.recorder = &MockOfferMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOffer) EXPECT() *MockOfferMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOffer) Create(arg0 *model.Offer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockOfferMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOffer)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockOffer) Delete(arg0, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOfferMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOffer)(nil).Delete), arg0, arg1)
}

// GetAllByParams mocks base method.
func (m *MockOffer) GetAllByParams(arg0, arg1 int, arg2 string) ([]*model.Offer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByParams", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*model.Offer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByParams indicates an expected call of GetAllByParams.
func (mr *MockOfferMockRecorder) GetAllByParams(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByParams", reflect.TypeOf((*MockOffer)(nil).GetAllByParams), arg0, arg1, arg2)
}

// GetByTuple mocks base method.
func (m *MockOffer) GetByTuple(arg0, arg1 int) (model.Offer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTuple", arg0, arg1)
	ret0, _ := ret[0].(model.Offer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTuple indicates an expected call of GetByTuple.
func (mr *MockOfferMockRecorder) GetByTuple(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTuple", reflect.TypeOf((*MockOffer)(nil).GetByTuple), arg0, arg1)
}

// Update mocks base method.
func (m *MockOffer) Update(arg0 *model.Offer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockOfferMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOffer)(nil).Update), arg0)
}