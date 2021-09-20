// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	"github.com/keptn/keptn/go-sdk/pkg/sdk"
	"sync"
)

// Ensure, that TaskHandlerMock does implement sdk.TaskHandler.
// If this is not the case, regenerate this file with moq.
var _ sdk.TaskHandler = &TaskHandlerMock{}

// TaskHandlerMock is a mock implementation of sdk.TaskHandler.
//
// 	func TestSomethingThatUsesTaskHandler(t *testing.T) {
//
// 		// make and configure a mocked sdk.TaskHandler
// 		mockedTaskHandler := &TaskHandlerMock{
// 			ExecuteFunc: func(keptnHandle sdk.IKeptn, event sdk.KeptnEvent) (interface{}, *sdk.Error) {
// 				panic("mock out the Execute method")
// 			},
// 		}
//
// 		// use mockedTaskHandler in code that requires sdk.TaskHandler
// 		// and then make assertions.
//
// 	}
type TaskHandlerMock struct {
	// ExecuteFunc mocks the Execute method.
	ExecuteFunc func(keptnHandle sdk.IKeptn, event sdk.KeptnEvent) (interface{}, *sdk.Error)

	// calls tracks calls to the methods.
	calls struct {
		// Execute holds details about calls to the Execute method.
		Execute []struct {
			// KeptnHandle is the keptnHandle argument value.
			KeptnHandle sdk.IKeptn
			// Event is the event argument value.
			Event sdk.KeptnEvent
		}
	}
	lockExecute sync.RWMutex
}

// Execute calls ExecuteFunc.
func (mock *TaskHandlerMock) Execute(keptnHandle sdk.IKeptn, event sdk.KeptnEvent) (interface{}, *sdk.Error) {
	if mock.ExecuteFunc == nil {
		panic("TaskHandlerMock.ExecuteFunc: method is nil but TaskHandler.Execute was just called")
	}
	callInfo := struct {
		KeptnHandle sdk.IKeptn
		Event       sdk.KeptnEvent
	}{
		KeptnHandle: keptnHandle,
		Event:       event,
	}
	mock.lockExecute.Lock()
	mock.calls.Execute = append(mock.calls.Execute, callInfo)
	mock.lockExecute.Unlock()
	return mock.ExecuteFunc(keptnHandle, event)
}

// ExecuteCalls gets all the calls that were made to Execute.
// Check the length with:
//     len(mockedTaskHandler.ExecuteCalls())
func (mock *TaskHandlerMock) ExecuteCalls() []struct {
	KeptnHandle sdk.IKeptn
	Event       sdk.KeptnEvent
} {
	var calls []struct {
		KeptnHandle sdk.IKeptn
		Event       sdk.KeptnEvent
	}
	mock.lockExecute.RLock()
	calls = mock.calls.Execute
	mock.lockExecute.RUnlock()
	return calls
}