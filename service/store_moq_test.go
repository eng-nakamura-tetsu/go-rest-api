// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package service

import (
	"context"
	"database/sql"
	"github.com/guregu/sqlx"
	"sync"
)

// BeginnerMock is a mock implementation of store.Beginner.
//
// 	func TestSomethingThatUsesBeginner(t *testing.T) {
//
// 		// make and configure a mocked store.Beginner
// 		mockedBeginner := &BeginnerMock{
// 			BeginTxFunc: func(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
// 				panic("mock out the BeginTx method")
// 			},
// 		}
//
// 		// use mockedBeginner in code that requires store.Beginner
// 		// and then make assertions.
//
// 	}
type BeginnerMock struct {
	// BeginTxFunc mocks the BeginTx method.
	BeginTxFunc func(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)

	// calls tracks calls to the methods.
	calls struct {
		// BeginTx holds details about calls to the BeginTx method.
		BeginTx []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Opts is the opts argument value.
			Opts *sql.TxOptions
		}
	}
	lockBeginTx sync.RWMutex
}

// BeginTx calls BeginTxFunc.
func (mock *BeginnerMock) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	if mock.BeginTxFunc == nil {
		panic("BeginnerMock.BeginTxFunc: method is nil but Beginner.BeginTx was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Opts *sql.TxOptions
	}{
		Ctx:  ctx,
		Opts: opts,
	}
	mock.lockBeginTx.Lock()
	mock.calls.BeginTx = append(mock.calls.BeginTx, callInfo)
	mock.lockBeginTx.Unlock()
	return mock.BeginTxFunc(ctx, opts)
}

// BeginTxCalls gets all the calls that were made to BeginTx.
// Check the length with:
//     len(mockedBeginner.BeginTxCalls())
func (mock *BeginnerMock) BeginTxCalls() []struct {
	Ctx  context.Context
	Opts *sql.TxOptions
} {
	var calls []struct {
		Ctx  context.Context
		Opts *sql.TxOptions
	}
	mock.lockBeginTx.RLock()
	calls = mock.calls.BeginTx
	mock.lockBeginTx.RUnlock()
	return calls
}

// PreparerMock is a mock implementation of store.Preparer.
//
// 	func TestSomethingThatUsesPreparer(t *testing.T) {
//
// 		// make and configure a mocked store.Preparer
// 		mockedPreparer := &PreparerMock{
// 			PreparexContextFunc: func(ctx context.Context, query string) (*sqlx.Stmt, error) {
// 				panic("mock out the PreparexContext method")
// 			},
// 		}
//
// 		// use mockedPreparer in code that requires store.Preparer
// 		// and then make assertions.
//
// 	}
type PreparerMock struct {
	// PreparexContextFunc mocks the PreparexContext method.
	PreparexContextFunc func(ctx context.Context, query string) (*sqlx.Stmt, error)

	// calls tracks calls to the methods.
	calls struct {
		// PreparexContext holds details about calls to the PreparexContext method.
		PreparexContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
		}
	}
	lockPreparexContext sync.RWMutex
}

// PreparexContext calls PreparexContextFunc.
func (mock *PreparerMock) PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error) {
	if mock.PreparexContextFunc == nil {
		panic("PreparerMock.PreparexContextFunc: method is nil but Preparer.PreparexContext was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Query string
	}{
		Ctx:   ctx,
		Query: query,
	}
	mock.lockPreparexContext.Lock()
	mock.calls.PreparexContext = append(mock.calls.PreparexContext, callInfo)
	mock.lockPreparexContext.Unlock()
	return mock.PreparexContextFunc(ctx, query)
}

// PreparexContextCalls gets all the calls that were made to PreparexContext.
// Check the length with:
//     len(mockedPreparer.PreparexContextCalls())
func (mock *PreparerMock) PreparexContextCalls() []struct {
	Ctx   context.Context
	Query string
} {
	var calls []struct {
		Ctx   context.Context
		Query string
	}
	mock.lockPreparexContext.RLock()
	calls = mock.calls.PreparexContext
	mock.lockPreparexContext.RUnlock()
	return calls
}

// ExecerMock is a mock implementation of store.Execer.
//
// 	func TestSomethingThatUsesExecer(t *testing.T) {
//
// 		// make and configure a mocked store.Execer
// 		mockedExecer := &ExecerMock{
// 			ExecContextFunc: func(ctx context.Context, query string, args ...any) (sql.Result, error) {
// 				panic("mock out the ExecContext method")
// 			},
// 			NamedExecContextFunc: func(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
// 				panic("mock out the NamedExecContext method")
// 			},
// 		}
//
// 		// use mockedExecer in code that requires store.Execer
// 		// and then make assertions.
//
// 	}
type ExecerMock struct {
	// ExecContextFunc mocks the ExecContext method.
	ExecContextFunc func(ctx context.Context, query string, args ...any) (sql.Result, error)

	// NamedExecContextFunc mocks the NamedExecContext method.
	NamedExecContextFunc func(ctx context.Context, query string, arg interface{}) (sql.Result, error)

	// calls tracks calls to the methods.
	calls struct {
		// ExecContext holds details about calls to the ExecContext method.
		ExecContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []any
		}
		// NamedExecContext holds details about calls to the NamedExecContext method.
		NamedExecContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
			// Arg is the arg argument value.
			Arg interface{}
		}
	}
	lockExecContext      sync.RWMutex
	lockNamedExecContext sync.RWMutex
}

// ExecContext calls ExecContextFunc.
func (mock *ExecerMock) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	if mock.ExecContextFunc == nil {
		panic("ExecerMock.ExecContextFunc: method is nil but Execer.ExecContext was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Query string
		Args  []any
	}{
		Ctx:   ctx,
		Query: query,
		Args:  args,
	}
	mock.lockExecContext.Lock()
	mock.calls.ExecContext = append(mock.calls.ExecContext, callInfo)
	mock.lockExecContext.Unlock()
	return mock.ExecContextFunc(ctx, query, args...)
}

// ExecContextCalls gets all the calls that were made to ExecContext.
// Check the length with:
//     len(mockedExecer.ExecContextCalls())
func (mock *ExecerMock) ExecContextCalls() []struct {
	Ctx   context.Context
	Query string
	Args  []any
} {
	var calls []struct {
		Ctx   context.Context
		Query string
		Args  []any
	}
	mock.lockExecContext.RLock()
	calls = mock.calls.ExecContext
	mock.lockExecContext.RUnlock()
	return calls
}

// NamedExecContext calls NamedExecContextFunc.
func (mock *ExecerMock) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	if mock.NamedExecContextFunc == nil {
		panic("ExecerMock.NamedExecContextFunc: method is nil but Execer.NamedExecContext was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Query string
		Arg   interface{}
	}{
		Ctx:   ctx,
		Query: query,
		Arg:   arg,
	}
	mock.lockNamedExecContext.Lock()
	mock.calls.NamedExecContext = append(mock.calls.NamedExecContext, callInfo)
	mock.lockNamedExecContext.Unlock()
	return mock.NamedExecContextFunc(ctx, query, arg)
}

// NamedExecContextCalls gets all the calls that were made to NamedExecContext.
// Check the length with:
//     len(mockedExecer.NamedExecContextCalls())
func (mock *ExecerMock) NamedExecContextCalls() []struct {
	Ctx   context.Context
	Query string
	Arg   interface{}
} {
	var calls []struct {
		Ctx   context.Context
		Query string
		Arg   interface{}
	}
	mock.lockNamedExecContext.RLock()
	calls = mock.calls.NamedExecContext
	mock.lockNamedExecContext.RUnlock()
	return calls
}

// QueryerMock is a mock implementation of store.Queryer.
//
// 	func TestSomethingThatUsesQueryer(t *testing.T) {
//
// 		// make and configure a mocked store.Queryer
// 		mockedQueryer := &QueryerMock{
// 			GetContextFunc: func(ctx context.Context, dest interface{}, query string, args ...any) error {
// 				panic("mock out the GetContext method")
// 			},
// 			PreparexContextFunc: func(ctx context.Context, query string) (*sqlx.Stmt, error) {
// 				panic("mock out the PreparexContext method")
// 			},
// 			QueryRowxContextFunc: func(ctx context.Context, query string, args ...any) *sqlx.Row {
// 				panic("mock out the QueryRowxContext method")
// 			},
// 			QueryxContextFunc: func(ctx context.Context, query string, args ...any) (*sqlx.Rows, error) {
// 				panic("mock out the QueryxContext method")
// 			},
// 			SelectContextFunc: func(ctx context.Context, dest interface{}, query string, args ...any) error {
// 				panic("mock out the SelectContext method")
// 			},
// 		}
//
// 		// use mockedQueryer in code that requires store.Queryer
// 		// and then make assertions.
//
// 	}
type QueryerMock struct {
	// GetContextFunc mocks the GetContext method.
	GetContextFunc func(ctx context.Context, dest interface{}, query string, args ...any) error

	// PreparexContextFunc mocks the PreparexContext method.
	PreparexContextFunc func(ctx context.Context, query string) (*sqlx.Stmt, error)

	// QueryRowxContextFunc mocks the QueryRowxContext method.
	QueryRowxContextFunc func(ctx context.Context, query string, args ...any) *sqlx.Row

	// QueryxContextFunc mocks the QueryxContext method.
	QueryxContextFunc func(ctx context.Context, query string, args ...any) (*sqlx.Rows, error)

	// SelectContextFunc mocks the SelectContext method.
	SelectContextFunc func(ctx context.Context, dest interface{}, query string, args ...any) error

	// calls tracks calls to the methods.
	calls struct {
		// GetContext holds details about calls to the GetContext method.
		GetContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Dest is the dest argument value.
			Dest interface{}
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []any
		}
		// PreparexContext holds details about calls to the PreparexContext method.
		PreparexContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
		}
		// QueryRowxContext holds details about calls to the QueryRowxContext method.
		QueryRowxContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []any
		}
		// QueryxContext holds details about calls to the QueryxContext method.
		QueryxContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []any
		}
		// SelectContext holds details about calls to the SelectContext method.
		SelectContext []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Dest is the dest argument value.
			Dest interface{}
			// Query is the query argument value.
			Query string
			// Args is the args argument value.
			Args []any
		}
	}
	lockGetContext       sync.RWMutex
	lockPreparexContext  sync.RWMutex
	lockQueryRowxContext sync.RWMutex
	lockQueryxContext    sync.RWMutex
	lockSelectContext    sync.RWMutex
}

// GetContext calls GetContextFunc.
func (mock *QueryerMock) GetContext(ctx context.Context, dest interface{}, query string, args ...any) error {
	if mock.GetContextFunc == nil {
		panic("QueryerMock.GetContextFunc: method is nil but Queryer.GetContext was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Dest  interface{}
		Query string
		Args  []any
	}{
		Ctx:   ctx,
		Dest:  dest,
		Query: query,
		Args:  args,
	}
	mock.lockGetContext.Lock()
	mock.calls.GetContext = append(mock.calls.GetContext, callInfo)
	mock.lockGetContext.Unlock()
	return mock.GetContextFunc(ctx, dest, query, args...)
}

// GetContextCalls gets all the calls that were made to GetContext.
// Check the length with:
//     len(mockedQueryer.GetContextCalls())
func (mock *QueryerMock) GetContextCalls() []struct {
	Ctx   context.Context
	Dest  interface{}
	Query string
	Args  []any
} {
	var calls []struct {
		Ctx   context.Context
		Dest  interface{}
		Query string
		Args  []any
	}
	mock.lockGetContext.RLock()
	calls = mock.calls.GetContext
	mock.lockGetContext.RUnlock()
	return calls
}

// PreparexContext calls PreparexContextFunc.
func (mock *QueryerMock) PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error) {
	if mock.PreparexContextFunc == nil {
		panic("QueryerMock.PreparexContextFunc: method is nil but Queryer.PreparexContext was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Query string
	}{
		Ctx:   ctx,
		Query: query,
	}
	mock.lockPreparexContext.Lock()
	mock.calls.PreparexContext = append(mock.calls.PreparexContext, callInfo)
	mock.lockPreparexContext.Unlock()
	return mock.PreparexContextFunc(ctx, query)
}

// PreparexContextCalls gets all the calls that were made to PreparexContext.
// Check the length with:
//     len(mockedQueryer.PreparexContextCalls())
func (mock *QueryerMock) PreparexContextCalls() []struct {
	Ctx   context.Context
	Query string
} {
	var calls []struct {
		Ctx   context.Context
		Query string
	}
	mock.lockPreparexContext.RLock()
	calls = mock.calls.PreparexContext
	mock.lockPreparexContext.RUnlock()
	return calls
}

// QueryRowxContext calls QueryRowxContextFunc.
func (mock *QueryerMock) QueryRowxContext(ctx context.Context, query string, args ...any) *sqlx.Row {
	if mock.QueryRowxContextFunc == nil {
		panic("QueryerMock.QueryRowxContextFunc: method is nil but Queryer.QueryRowxContext was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Query string
		Args  []any
	}{
		Ctx:   ctx,
		Query: query,
		Args:  args,
	}
	mock.lockQueryRowxContext.Lock()
	mock.calls.QueryRowxContext = append(mock.calls.QueryRowxContext, callInfo)
	mock.lockQueryRowxContext.Unlock()
	return mock.QueryRowxContextFunc(ctx, query, args...)
}

// QueryRowxContextCalls gets all the calls that were made to QueryRowxContext.
// Check the length with:
//     len(mockedQueryer.QueryRowxContextCalls())
func (mock *QueryerMock) QueryRowxContextCalls() []struct {
	Ctx   context.Context
	Query string
	Args  []any
} {
	var calls []struct {
		Ctx   context.Context
		Query string
		Args  []any
	}
	mock.lockQueryRowxContext.RLock()
	calls = mock.calls.QueryRowxContext
	mock.lockQueryRowxContext.RUnlock()
	return calls
}

// QueryxContext calls QueryxContextFunc.
func (mock *QueryerMock) QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error) {
	if mock.QueryxContextFunc == nil {
		panic("QueryerMock.QueryxContextFunc: method is nil but Queryer.QueryxContext was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Query string
		Args  []any
	}{
		Ctx:   ctx,
		Query: query,
		Args:  args,
	}
	mock.lockQueryxContext.Lock()
	mock.calls.QueryxContext = append(mock.calls.QueryxContext, callInfo)
	mock.lockQueryxContext.Unlock()
	return mock.QueryxContextFunc(ctx, query, args...)
}

// QueryxContextCalls gets all the calls that were made to QueryxContext.
// Check the length with:
//     len(mockedQueryer.QueryxContextCalls())
func (mock *QueryerMock) QueryxContextCalls() []struct {
	Ctx   context.Context
	Query string
	Args  []any
} {
	var calls []struct {
		Ctx   context.Context
		Query string
		Args  []any
	}
	mock.lockQueryxContext.RLock()
	calls = mock.calls.QueryxContext
	mock.lockQueryxContext.RUnlock()
	return calls
}

// SelectContext calls SelectContextFunc.
func (mock *QueryerMock) SelectContext(ctx context.Context, dest interface{}, query string, args ...any) error {
	if mock.SelectContextFunc == nil {
		panic("QueryerMock.SelectContextFunc: method is nil but Queryer.SelectContext was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Dest  interface{}
		Query string
		Args  []any
	}{
		Ctx:   ctx,
		Dest:  dest,
		Query: query,
		Args:  args,
	}
	mock.lockSelectContext.Lock()
	mock.calls.SelectContext = append(mock.calls.SelectContext, callInfo)
	mock.lockSelectContext.Unlock()
	return mock.SelectContextFunc(ctx, dest, query, args...)
}

// SelectContextCalls gets all the calls that were made to SelectContext.
// Check the length with:
//     len(mockedQueryer.SelectContextCalls())
func (mock *QueryerMock) SelectContextCalls() []struct {
	Ctx   context.Context
	Dest  interface{}
	Query string
	Args  []any
} {
	var calls []struct {
		Ctx   context.Context
		Dest  interface{}
		Query string
		Args  []any
	}
	mock.lockSelectContext.RLock()
	calls = mock.calls.SelectContext
	mock.lockSelectContext.RUnlock()
	return calls
}