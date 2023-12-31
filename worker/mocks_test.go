// Code generated by MockGen. DO NOT EDIT.
// Source: ./worker.go

// Package worker is a generated GoMock package.
package worker

import (
	context "context"
	reflect "reflect"

	outbox "github.com/crxfoz/outbox"
	gomock "go.uber.org/mock/gomock"
)

// MockRecord is a mock of Record interface.
type MockRecord[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockRecordMockRecorder[T]
}

// MockRecordMockRecorder is the mock recorder for MockRecord.
type MockRecordMockRecorder[T any] struct {
	mock *MockRecord[T]
}

// NewMockRecord creates a new mock instance.
func NewMockRecord[T any](ctrl *gomock.Controller) *MockRecord[T] {
	mock := &MockRecord[T]{ctrl: ctrl}
	mock.recorder = &MockRecordMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRecord[T]) EXPECT() *MockRecordMockRecorder[T] {
	return m.recorder
}

// GetID mocks base method.
func (m *MockRecord[T]) GetID() T {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(T)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockRecordMockRecorder[T]) GetID() *RecordGetIDCall[T] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockRecord[T])(nil).GetID))
	return &RecordGetIDCall[T]{Call: call}
}

// RecordGetIDCall wrap *gomock.Call
type RecordGetIDCall[T any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RecordGetIDCall[T]) Return(arg0 T) *RecordGetIDCall[T] {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RecordGetIDCall[T]) Do(f func() T) *RecordGetIDCall[T] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RecordGetIDCall[T]) DoAndReturn(f func() T) *RecordGetIDCall[T] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockTransaction is a mock of Transaction interface.
type MockTransaction struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionMockRecorder
}

// MockTransactionMockRecorder is the mock recorder for MockTransaction.
type MockTransactionMockRecorder struct {
	mock *MockTransaction
}

// NewMockTransaction creates a new mock instance.
func NewMockTransaction(ctrl *gomock.Controller) *MockTransaction {
	mock := &MockTransaction{ctrl: ctrl}
	mock.recorder = &MockTransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransaction) EXPECT() *MockTransactionMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockTransaction) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockTransactionMockRecorder) Commit() *TransactionCommitCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTransaction)(nil).Commit))
	return &TransactionCommitCall{Call: call}
}

// TransactionCommitCall wrap *gomock.Call
type TransactionCommitCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *TransactionCommitCall) Return(arg0 error) *TransactionCommitCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *TransactionCommitCall) Do(f func() error) *TransactionCommitCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *TransactionCommitCall) DoAndReturn(f func() error) *TransactionCommitCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Rollback mocks base method.
func (m *MockTransaction) Rollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockTransactionMockRecorder) Rollback() *TransactionRollbackCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockTransaction)(nil).Rollback))
	return &TransactionRollbackCall{Call: call}
}

// TransactionRollbackCall wrap *gomock.Call
type TransactionRollbackCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *TransactionRollbackCall) Return(arg0 error) *TransactionRollbackCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *TransactionRollbackCall) Do(f func() error) *TransactionRollbackCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *TransactionRollbackCall) DoAndReturn(f func() error) *TransactionRollbackCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockRepository is a mock of Repository interface.
type MockRepository[ID any, O Record[ID], TX Transaction] struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder[ID, O, TX]
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder[ID any, O Record[ID], TX Transaction] struct {
	mock *MockRepository[ID, O, TX]
}

// NewMockRepository creates a new mock instance.
func NewMockRepository[ID any, O Record[ID], TX Transaction](ctrl *gomock.Controller) *MockRepository[ID, O, TX] {
	mock := &MockRepository[ID, O, TX]{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder[ID, O, TX]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository[ID, O, TX]) EXPECT() *MockRepositoryMockRecorder[ID, O, TX] {
	return m.recorder
}

// GetTx mocks base method.
func (m *MockRepository[ID, O, TX]) GetTx(ctx context.Context) (TX, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTx", ctx)
	ret0, _ := ret[0].(TX)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTx indicates an expected call of GetTx.
func (mr *MockRepositoryMockRecorder[ID, O, TX]) GetTx(ctx interface{}) *RepositoryGetTxCall[ID, O, TX] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTx", reflect.TypeOf((*MockRepository[ID, O, TX])(nil).GetTx), ctx)
	return &RepositoryGetTxCall[ID, O, TX]{Call: call}
}

// RepositoryGetTxCall wrap *gomock.Call
type RepositoryGetTxCall[ID any, O Record[ID], TX Transaction] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RepositoryGetTxCall[ID, O, TX]) Return(arg0 TX, arg1 error) *RepositoryGetTxCall[ID, O, TX] {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RepositoryGetTxCall[ID, O, TX]) Do(f func(context.Context) (TX, error)) *RepositoryGetTxCall[ID, O, TX] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RepositoryGetTxCall[ID, O, TX]) DoAndReturn(f func(context.Context) (TX, error)) *RepositoryGetTxCall[ID, O, TX] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RecordWithLock mocks base method.
func (m *MockRepository[ID, O, TX]) RecordWithLock(ctx context.Context, tx TX, limit uint, skipLocked bool, status outbox.Status) ([]O, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordWithLock", ctx, tx, limit, skipLocked, status)
	ret0, _ := ret[0].([]O)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordWithLock indicates an expected call of RecordWithLock.
func (mr *MockRepositoryMockRecorder[ID, O, TX]) RecordWithLock(ctx, tx, limit, skipLocked, status interface{}) *RepositoryRecordWithLockCall[ID, O, TX] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordWithLock", reflect.TypeOf((*MockRepository[ID, O, TX])(nil).RecordWithLock), ctx, tx, limit, skipLocked, status)
	return &RepositoryRecordWithLockCall[ID, O, TX]{Call: call}
}

// RepositoryRecordWithLockCall wrap *gomock.Call
type RepositoryRecordWithLockCall[ID any, O Record[ID], TX Transaction] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RepositoryRecordWithLockCall[ID, O, TX]) Return(arg0 []O, arg1 error) *RepositoryRecordWithLockCall[ID, O, TX] {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RepositoryRecordWithLockCall[ID, O, TX]) Do(f func(context.Context, TX, uint, bool, outbox.Status) ([]O, error)) *RepositoryRecordWithLockCall[ID, O, TX] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RepositoryRecordWithLockCall[ID, O, TX]) DoAndReturn(f func(context.Context, TX, uint, bool, outbox.Status) ([]O, error)) *RepositoryRecordWithLockCall[ID, O, TX] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RecordsByStatus mocks base method.
func (m *MockRepository[ID, O, TX]) RecordsByStatus(ctx context.Context, status outbox.Status) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecordsByStatus", ctx, status)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordsByStatus indicates an expected call of RecordsByStatus.
func (mr *MockRepositoryMockRecorder[ID, O, TX]) RecordsByStatus(ctx, status interface{}) *RepositoryRecordsByStatusCall[ID, O, TX] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordsByStatus", reflect.TypeOf((*MockRepository[ID, O, TX])(nil).RecordsByStatus), ctx, status)
	return &RepositoryRecordsByStatusCall[ID, O, TX]{Call: call}
}

// RepositoryRecordsByStatusCall wrap *gomock.Call
type RepositoryRecordsByStatusCall[ID any, O Record[ID], TX Transaction] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RepositoryRecordsByStatusCall[ID, O, TX]) Return(arg0 uint, arg1 error) *RepositoryRecordsByStatusCall[ID, O, TX] {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RepositoryRecordsByStatusCall[ID, O, TX]) Do(f func(context.Context, outbox.Status) (uint, error)) *RepositoryRecordsByStatusCall[ID, O, TX] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RepositoryRecordsByStatusCall[ID, O, TX]) DoAndReturn(f func(context.Context, outbox.Status) (uint, error)) *RepositoryRecordsByStatusCall[ID, O, TX] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateStatus mocks base method.
func (m *MockRepository[ID, O, TX]) UpdateStatus(ctx context.Context, tx TX, id ID, status outbox.Status) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", ctx, tx, id, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockRepositoryMockRecorder[ID, O, TX]) UpdateStatus(ctx, tx, id, status interface{}) *RepositoryUpdateStatusCall[ID, O, TX] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockRepository[ID, O, TX])(nil).UpdateStatus), ctx, tx, id, status)
	return &RepositoryUpdateStatusCall[ID, O, TX]{Call: call}
}

// RepositoryUpdateStatusCall wrap *gomock.Call
type RepositoryUpdateStatusCall[ID any, O Record[ID], TX Transaction] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *RepositoryUpdateStatusCall[ID, O, TX]) Return(arg0 error) *RepositoryUpdateStatusCall[ID, O, TX] {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *RepositoryUpdateStatusCall[ID, O, TX]) Do(f func(context.Context, TX, ID, outbox.Status) error) *RepositoryUpdateStatusCall[ID, O, TX] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *RepositoryUpdateStatusCall[ID, O, TX]) DoAndReturn(f func(context.Context, TX, ID, outbox.Status) error) *RepositoryUpdateStatusCall[ID, O, TX] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockEventMapper is a mock of EventMapper interface.
type MockEventMapper[ID any, O Record[ID], R any] struct {
	ctrl     *gomock.Controller
	recorder *MockEventMapperMockRecorder[ID, O, R]
}

// MockEventMapperMockRecorder is the mock recorder for MockEventMapper.
type MockEventMapperMockRecorder[ID any, O Record[ID], R any] struct {
	mock *MockEventMapper[ID, O, R]
}

// NewMockEventMapper creates a new mock instance.
func NewMockEventMapper[ID any, O Record[ID], R any](ctrl *gomock.Controller) *MockEventMapper[ID, O, R] {
	mock := &MockEventMapper[ID, O, R]{ctrl: ctrl}
	mock.recorder = &MockEventMapperMockRecorder[ID, O, R]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventMapper[ID, O, R]) EXPECT() *MockEventMapperMockRecorder[ID, O, R] {
	return m.recorder
}

// ToEvent mocks base method.
func (m *MockEventMapper[ID, O, R]) ToEvent(arg0 O) (R, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToEvent", arg0)
	ret0, _ := ret[0].(R)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToEvent indicates an expected call of ToEvent.
func (mr *MockEventMapperMockRecorder[ID, O, R]) ToEvent(arg0 interface{}) *EventMapperToEventCall[ID, O, R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToEvent", reflect.TypeOf((*MockEventMapper[ID, O, R])(nil).ToEvent), arg0)
	return &EventMapperToEventCall[ID, O, R]{Call: call}
}

// EventMapperToEventCall wrap *gomock.Call
type EventMapperToEventCall[ID any, O Record[ID], R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *EventMapperToEventCall[ID, O, R]) Return(arg0 R, arg1 error) *EventMapperToEventCall[ID, O, R] {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *EventMapperToEventCall[ID, O, R]) Do(f func(O) (R, error)) *EventMapperToEventCall[ID, O, R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *EventMapperToEventCall[ID, O, R]) DoAndReturn(f func(O) (R, error)) *EventMapperToEventCall[ID, O, R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockHandler is a mock of Handler interface.
type MockHandler[R any] struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder[R]
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder[R any] struct {
	mock *MockHandler[R]
}

// NewMockHandler creates a new mock instance.
func NewMockHandler[R any](ctrl *gomock.Controller) *MockHandler[R] {
	mock := &MockHandler[R]{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder[R]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler[R]) EXPECT() *MockHandlerMockRecorder[R] {
	return m.recorder
}

// Process mocks base method.
func (m *MockHandler[R]) Process(arg0 context.Context, arg1 R) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Process", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Process indicates an expected call of Process.
func (mr *MockHandlerMockRecorder[R]) Process(arg0, arg1 interface{}) *HandlerProcessCall[R] {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockHandler[R])(nil).Process), arg0, arg1)
	return &HandlerProcessCall[R]{Call: call}
}

// HandlerProcessCall wrap *gomock.Call
type HandlerProcessCall[R any] struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *HandlerProcessCall[R]) Return(arg0 error) *HandlerProcessCall[R] {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *HandlerProcessCall[R]) Do(f func(context.Context, R) error) *HandlerProcessCall[R] {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *HandlerProcessCall[R]) DoAndReturn(f func(context.Context, R) error) *HandlerProcessCall[R] {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
