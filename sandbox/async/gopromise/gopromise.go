package gopromise

import (
	"github.com/fengdotdev/golibs-future/sandbox/async"
)

var _ async.Promise[any] = (*GoPromise[any])(nil)

func NewGoPromise[T any](value T, err error) *GoPromise[T] {
	return &GoPromise[T]{
		initialized: true,
		cached:      value,
		cachedErr:   err,
		isCompleted: true,
	}
}

func NewIncompleteGoPromise[T any]() (*GoPromise[T], func(T, error)) {
	var Zero T

	p := &GoPromise[T]{
		initialized: true,
		cached:      Zero,
		cachedErr:   ErrPromiseNotResolved,
		isCompleted: false,
	}

	update := func(value T, err error) {
		p.cached = value
		p.cachedErr = err
		p.isCompleted = true
	}

	return p, update
}

// a promise is a the result of an async operation
type GoPromise[T any] struct {
	initialized bool
	cached      T
	cachedErr   error
	isCompleted bool
}

// Catch implements async.Promise.
func (g *GoPromise[T]) Catch(catch func(error)) {
	panic("unimplemented")
}

// Error implements async.Promise.
func (g *GoPromise[T]) Error() error {
	panic("unimplemented")
}

// Finally implements async.Promise.
func (g *GoPromise[T]) Finally(finally func(T, error)) {
	panic("unimplemented")
}

// IsCompleted implements async.Promise.
func (g *GoPromise[T]) IsCompleted() bool {
	panic("unimplemented")
}

// IsValid implements async.Promise.
func (g *GoPromise[T]) IsValid() bool {
	panic("unimplemented")
}

// String implements async.Promise.
func (g *GoPromise[T]) String() string {
	panic("unimplemented")
}

// Then implements async.Promise.
func (g *GoPromise[T]) Then(then func(T)) {
	panic("unimplemented")
}

// Value implements async.Promise.
func (g *GoPromise[T]) Value() T {
	panic("unimplemented")
}

// ValueOr implements async.Promise.
func (g *GoPromise[T]) ValueOr(or T) T {
	panic("unimplemented")
}

// ValueOrErr implements async.Promise.
func (g *GoPromise[T]) ValueOrErr() (T, error) {
	panic("unimplemented")
}

// ValueOrPanic implements async.Promise.
func (g *GoPromise[T]) ValueOrPanic() T {
	panic("unimplemented")
}
