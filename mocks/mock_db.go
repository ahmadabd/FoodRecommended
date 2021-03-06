// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ahmadabd/FoodRecommended.git/internal/repository (interfaces: DB)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
	gomock "github.com/golang/mock/gomock"
)

// MockDB is a mock of DB interface.
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB.
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance.
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// CreateFood mocks base method.
func (m *MockDB) CreateFood(arg0 context.Context, arg1 model.Food) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFood", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFood indicates an expected call of CreateFood.
func (mr *MockDBMockRecorder) CreateFood(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFood", reflect.TypeOf((*MockDB)(nil).CreateFood), arg0, arg1)
}

// GetFoods mocks base method.
func (m *MockDB) GetFoods(arg0 context.Context, arg1 int) ([]model.Food, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFoods", arg0, arg1)
	ret0, _ := ret[0].([]model.Food)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFoods indicates an expected call of GetFoods.
func (mr *MockDBMockRecorder) GetFoods(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFoods", reflect.TypeOf((*MockDB)(nil).GetFoods), arg0, arg1)
}

// GetRandomFood mocks base method.
func (m *MockDB) GetRandomFood(arg0 context.Context) (model.Food, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandomFood", arg0)
	ret0, _ := ret[0].(model.Food)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRandomFood indicates an expected call of GetRandomFood.
func (mr *MockDBMockRecorder) GetRandomFood(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandomFood", reflect.TypeOf((*MockDB)(nil).GetRandomFood), arg0)
}
