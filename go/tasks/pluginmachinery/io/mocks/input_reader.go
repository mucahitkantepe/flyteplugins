// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import context "context"
import core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

import mock "github.com/stretchr/testify/mock"
import storage "github.com/lyft/flytestdlib/storage"

// InputReader is an autogenerated mock type for the InputReader type
type InputReader struct {
	mock.Mock
}

type InputReader_Get struct {
	*mock.Call
}

func (_m InputReader_Get) Return(_a0 *core.LiteralMap, _a1 error) *InputReader_Get {
	return &InputReader_Get{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *InputReader) OnGet(ctx context.Context) *InputReader_Get {
	c := _m.On("Get", ctx)
	return &InputReader_Get{Call: c}
}

func (_m *InputReader) OnGetMatch(matchers ...interface{}) *InputReader_Get {
	c := _m.On("Get", matchers...)
	return &InputReader_Get{Call: c}
}

// Get provides a mock function with given fields: ctx
func (_m *InputReader) Get(ctx context.Context) (*core.LiteralMap, error) {
	ret := _m.Called(ctx)

	var r0 *core.LiteralMap
	if rf, ok := ret.Get(0).(func(context.Context) *core.LiteralMap); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.LiteralMap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type InputReader_GetInputPath struct {
	*mock.Call
}

func (_m InputReader_GetInputPath) Return(_a0 storage.DataReference) *InputReader_GetInputPath {
	return &InputReader_GetInputPath{Call: _m.Call.Return(_a0)}
}

func (_m *InputReader) OnGetInputPath() *InputReader_GetInputPath {
	c := _m.On("GetInputPath")
	return &InputReader_GetInputPath{Call: c}
}

func (_m *InputReader) OnGetInputPathMatch(matchers ...interface{}) *InputReader_GetInputPath {
	c := _m.On("GetInputPath", matchers...)
	return &InputReader_GetInputPath{Call: c}
}

// GetInputPath provides a mock function with given fields:
func (_m *InputReader) GetInputPath() storage.DataReference {
	ret := _m.Called()

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func() storage.DataReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	return r0
}

type InputReader_GetInputPrefixPath struct {
	*mock.Call
}

func (_m InputReader_GetInputPrefixPath) Return(_a0 storage.DataReference) *InputReader_GetInputPrefixPath {
	return &InputReader_GetInputPrefixPath{Call: _m.Call.Return(_a0)}
}

func (_m *InputReader) OnGetInputPrefixPath() *InputReader_GetInputPrefixPath {
	c := _m.On("GetInputPrefixPath")
	return &InputReader_GetInputPrefixPath{Call: c}
}

func (_m *InputReader) OnGetInputPrefixPathMatch(matchers ...interface{}) *InputReader_GetInputPrefixPath {
	c := _m.On("GetInputPrefixPath", matchers...)
	return &InputReader_GetInputPrefixPath{Call: c}
}

// GetInputPrefixPath provides a mock function with given fields:
func (_m *InputReader) GetInputPrefixPath() storage.DataReference {
	ret := _m.Called()

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func() storage.DataReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	return r0
}