package gofuture

import (
	"sync"
	"time"

	"github.com/fengdotdev/golibs-future/sandbox/async"
)

func NewFuture[T any](value T, err error) async.Future[T] {
	return &GoFuture[T]{
		initialized: true,
		cached:      value,
		cachedErr:   nil,
		isCompleted: true,
		mu:          sync.Mutex{},
	}
}

// NewFutureValue retuns a completed future with a value
func NewFutureValue[T any](value T) async.Future[T] {
	return NewFuture(value, nil)
}

// NewFutureError retuns a completed future with an error
func NewFutureError[T any](err error) async.Future[T] {
	var Zero T
	return NewFuture(Zero, err)
}

// NewIncompleteFutureAndCompleter returns a future that is not yet completed and a function to complete it
func NewIncompleteFutureAndCompleter[T any]() (*GoFuture[T], func(T, error)) {
	var Zero T

	p := &GoFuture[T]{
		mu:          sync.Mutex{},
		initialized: true,
		cached:      Zero,
		cachedErr:   ErrPromiseNotResolved,
		isCompleted: false,
	}

	return p, p.update
}

func NewIncompleteFutureWithTimer[T any](value T, err error, timeout time.Duration) (*GoFuture[T], func(T, error)) {
	var Zero T

	p := &GoFuture[T]{
		mu:          sync.Mutex{},
		initialized: true,
		cached:      Zero,
		cachedErr:   ErrPromiseNotResolved,
		isCompleted: false,
	}

	go func() {
		time.Sleep(timeout)
		p.update(value, err)
	}()

	return p, p.update
}
