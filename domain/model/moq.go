// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package model

import (
	"context"
	"sync"
)

// Ensure, that InvoiceRepositoryMock does implement InvoiceRepository.
// If this is not the case, regenerate this file with moq.
var _ InvoiceRepository = &InvoiceRepositoryMock{}

// InvoiceRepositoryMock is a mock implementation of InvoiceRepository.
//
//	func TestSomethingThatUsesInvoiceRepository(t *testing.T) {
//
//		// make and configure a mocked InvoiceRepository
//		mockedInvoiceRepository := &InvoiceRepositoryMock{
//			CreateFunc: func(ctx context.Context, invoice *Invoice) error {
//				panic("mock out the Create method")
//			},
//		}
//
//		// use mockedInvoiceRepository in code that requires InvoiceRepository
//		// and then make assertions.
//
//	}
type InvoiceRepositoryMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, invoice *Invoice) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Invoice is the invoice argument value.
			Invoice *Invoice
		}
	}
	lockCreate sync.RWMutex
}

// Create calls CreateFunc.
func (mock *InvoiceRepositoryMock) Create(ctx context.Context, invoice *Invoice) error {
	if mock.CreateFunc == nil {
		panic("InvoiceRepositoryMock.CreateFunc: method is nil but InvoiceRepository.Create was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Invoice *Invoice
	}{
		Ctx:     ctx,
		Invoice: invoice,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, invoice)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedInvoiceRepository.CreateCalls())
func (mock *InvoiceRepositoryMock) CreateCalls() []struct {
	Ctx     context.Context
	Invoice *Invoice
} {
	var calls []struct {
		Ctx     context.Context
		Invoice *Invoice
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}
