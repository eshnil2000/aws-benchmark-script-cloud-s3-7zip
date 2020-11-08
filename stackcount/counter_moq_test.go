// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package letsbuild13

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"sync"
)

// Ensure, that CounterInterfaceMock does implement CounterInterface.
// If this is not the case, regenerate this file with moq.
var _ CounterInterface = &CounterInterfaceMock{}

// CounterInterfaceMock is a mock implementation of CounterInterface.
//
//     func TestSomethingThatUsesCounterInterface(t *testing.T) {
//
//         // make and configure a mocked CounterInterface
//         mockedCounterInterface := &CounterInterfaceMock{
//             DescribeStacksFunc: func(ctx context.Context, params *cloudformation.DescribeStacksInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStacksOutput, error) {
// 	               panic("mock out the DescribeStacks method")
//             },
//         }
//
//         // use mockedCounterInterface in code that requires CounterInterface
//         // and then make assertions.
//
//     }
type CounterInterfaceMock struct {
	// DescribeStacksFunc mocks the DescribeStacks method.
	DescribeStacksFunc func(ctx context.Context, params *cloudformation.DescribeStacksInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStacksOutput, error)

	// calls tracks calls to the methods.
	calls struct {
		// DescribeStacks holds details about calls to the DescribeStacks method.
		DescribeStacks []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params *cloudformation.DescribeStacksInput
			// OptFns is the optFns argument value.
			OptFns []func(*cloudformation.Options)
		}
	}
	lockDescribeStacks sync.RWMutex
}

// DescribeStacks calls DescribeStacksFunc.
func (mock *CounterInterfaceMock) DescribeStacks(ctx context.Context, params *cloudformation.DescribeStacksInput, optFns ...func(*cloudformation.Options)) (*cloudformation.DescribeStacksOutput, error) {
	if mock.DescribeStacksFunc == nil {
		panic("CounterInterfaceMock.DescribeStacksFunc: method is nil but CounterInterface.DescribeStacks was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params *cloudformation.DescribeStacksInput
		OptFns []func(*cloudformation.Options)
	}{
		Ctx:    ctx,
		Params: params,
		OptFns: optFns,
	}
	mock.lockDescribeStacks.Lock()
	mock.calls.DescribeStacks = append(mock.calls.DescribeStacks, callInfo)
	mock.lockDescribeStacks.Unlock()
	return mock.DescribeStacksFunc(ctx, params, optFns...)
}

// DescribeStacksCalls gets all the calls that were made to DescribeStacks.
// Check the length with:
//     len(mockedCounterInterface.DescribeStacksCalls())
func (mock *CounterInterfaceMock) DescribeStacksCalls() []struct {
	Ctx    context.Context
	Params *cloudformation.DescribeStacksInput
	OptFns []func(*cloudformation.Options)
} {
	var calls []struct {
		Ctx    context.Context
		Params *cloudformation.DescribeStacksInput
		OptFns []func(*cloudformation.Options)
	}
	mock.lockDescribeStacks.RLock()
	calls = mock.calls.DescribeStacks
	mock.lockDescribeStacks.RUnlock()
	return calls
}
