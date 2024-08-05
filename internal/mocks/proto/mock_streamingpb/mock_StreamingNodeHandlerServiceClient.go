// Code generated by mockery v2.32.4. DO NOT EDIT.

package mock_streamingpb

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	streamingpb "github.com/milvus-io/milvus/internal/proto/streamingpb"
)

// MockStreamingNodeHandlerServiceClient is an autogenerated mock type for the StreamingNodeHandlerServiceClient type
type MockStreamingNodeHandlerServiceClient struct {
	mock.Mock
}

type MockStreamingNodeHandlerServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStreamingNodeHandlerServiceClient) EXPECT() *MockStreamingNodeHandlerServiceClient_Expecter {
	return &MockStreamingNodeHandlerServiceClient_Expecter{mock: &_m.Mock}
}

// Consume provides a mock function with given fields: ctx, opts
func (_m *MockStreamingNodeHandlerServiceClient) Consume(ctx context.Context, opts ...grpc.CallOption) (streamingpb.StreamingNodeHandlerService_ConsumeClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 streamingpb.StreamingNodeHandlerService_ConsumeClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...grpc.CallOption) (streamingpb.StreamingNodeHandlerService_ConsumeClient, error)); ok {
		return rf(ctx, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...grpc.CallOption) streamingpb.StreamingNodeHandlerService_ConsumeClient); ok {
		r0 = rf(ctx, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(streamingpb.StreamingNodeHandlerService_ConsumeClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStreamingNodeHandlerServiceClient_Consume_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Consume'
type MockStreamingNodeHandlerServiceClient_Consume_Call struct {
	*mock.Call
}

// Consume is a helper method to define mock.On call
//   - ctx context.Context
//   - opts ...grpc.CallOption
func (_e *MockStreamingNodeHandlerServiceClient_Expecter) Consume(ctx interface{}, opts ...interface{}) *MockStreamingNodeHandlerServiceClient_Consume_Call {
	return &MockStreamingNodeHandlerServiceClient_Consume_Call{Call: _e.mock.On("Consume",
		append([]interface{}{ctx}, opts...)...)}
}

func (_c *MockStreamingNodeHandlerServiceClient_Consume_Call) Run(run func(ctx context.Context, opts ...grpc.CallOption)) *MockStreamingNodeHandlerServiceClient_Consume_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockStreamingNodeHandlerServiceClient_Consume_Call) Return(_a0 streamingpb.StreamingNodeHandlerService_ConsumeClient, _a1 error) *MockStreamingNodeHandlerServiceClient_Consume_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStreamingNodeHandlerServiceClient_Consume_Call) RunAndReturn(run func(context.Context, ...grpc.CallOption) (streamingpb.StreamingNodeHandlerService_ConsumeClient, error)) *MockStreamingNodeHandlerServiceClient_Consume_Call {
	_c.Call.Return(run)
	return _c
}

// Produce provides a mock function with given fields: ctx, opts
func (_m *MockStreamingNodeHandlerServiceClient) Produce(ctx context.Context, opts ...grpc.CallOption) (streamingpb.StreamingNodeHandlerService_ProduceClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 streamingpb.StreamingNodeHandlerService_ProduceClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...grpc.CallOption) (streamingpb.StreamingNodeHandlerService_ProduceClient, error)); ok {
		return rf(ctx, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...grpc.CallOption) streamingpb.StreamingNodeHandlerService_ProduceClient); ok {
		r0 = rf(ctx, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(streamingpb.StreamingNodeHandlerService_ProduceClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStreamingNodeHandlerServiceClient_Produce_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Produce'
type MockStreamingNodeHandlerServiceClient_Produce_Call struct {
	*mock.Call
}

// Produce is a helper method to define mock.On call
//   - ctx context.Context
//   - opts ...grpc.CallOption
func (_e *MockStreamingNodeHandlerServiceClient_Expecter) Produce(ctx interface{}, opts ...interface{}) *MockStreamingNodeHandlerServiceClient_Produce_Call {
	return &MockStreamingNodeHandlerServiceClient_Produce_Call{Call: _e.mock.On("Produce",
		append([]interface{}{ctx}, opts...)...)}
}

func (_c *MockStreamingNodeHandlerServiceClient_Produce_Call) Run(run func(ctx context.Context, opts ...grpc.CallOption)) *MockStreamingNodeHandlerServiceClient_Produce_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockStreamingNodeHandlerServiceClient_Produce_Call) Return(_a0 streamingpb.StreamingNodeHandlerService_ProduceClient, _a1 error) *MockStreamingNodeHandlerServiceClient_Produce_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStreamingNodeHandlerServiceClient_Produce_Call) RunAndReturn(run func(context.Context, ...grpc.CallOption) (streamingpb.StreamingNodeHandlerService_ProduceClient, error)) *MockStreamingNodeHandlerServiceClient_Produce_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStreamingNodeHandlerServiceClient creates a new instance of MockStreamingNodeHandlerServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStreamingNodeHandlerServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStreamingNodeHandlerServiceClient {
	mock := &MockStreamingNodeHandlerServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}