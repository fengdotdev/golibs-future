package gofutureor

import (
	"fmt"

	"github.com/fengdotdev/golibs-funcs/v0/asserty"
	"github.com/fengdotdev/golibs-future/sandbox/async"
	"github.com/fengdotdev/golibs-traits/trait"
)

var _ trait.Result[any] = (*GoFutureOr[any])(nil)
var _ async.FutureOr[any] = (*GoFutureOr[any])(nil)

// Error implements trait.Result.
// If the value is valid, this method returns nil.
func (g *GoFutureOr[T]) Error() error {
	asserty.TrueWithMessage(g.initialized, "GoFutureOr must be initialized before calling Error()")
	return g.err
}

// IsValid implements trait.Result.
// It returns true if the Result holds a valid value; otherwise, it returns false.
func (g *GoFutureOr[T]) IsValid() bool {
	asserty.TrueWithMessage(g.initialized, "GoFutureOr must be initialized before calling IsValid()")
	return g.err == nil
}

// String implements trait.Result.
// It returns a string representation of the underlying value. retrurns the error message if the value is not valid.
// The format of the string will depend on the type of T.
func (g *GoFutureOr[T]) String() string {

	asserty.TrueWithMessage(g.initialized, "GoFutureOr must be initialized before calling String()")

	if g.err != nil {
		return g.err.Error()
	}
	return fmt.Sprintf("%v", g.value)
}

// Value implements trait.Result.
// It returns the underlying value. if the value is not valid, it will return the zero value of T or nil.
func (g *GoFutureOr[T]) Value() T {

	asserty.TrueWithMessage(g.initialized, "GoFutureOr must be initialized before calling Value()")

	if g.err != nil {
		var zero T
		return zero
	}
	return g.value
}

// ValueOr implements trait.Result.
// It returns the underlying value if it's valid; otherwise, it returns the provided 'or' value.
func (g *GoFutureOr[T]) ValueOr(or T) T {

	asserty.TrueWithMessage(g.initialized, "GoFutureOr must be initialized before calling ValueOr()")

	if g.err != nil {
		return or
	}
	return g.value
}

// ValueOrErr implements trait.Result.
// It returns the underlying value and an error.
func (g *GoFutureOr[T]) ValueOrErr() (T, error) {

	asserty.TrueWithMessage(g.initialized, "GoFutureOr must be initialized before calling ValueOrErr()")

	if g.err != nil {
		var zero T
		return zero, g.err
	}
	return g.value, nil
}

// ValueOrPanic implements trait.Result.
// It returns the underlying value.
// It panics if the value is not considered valid.
func (g *GoFutureOr[T]) ValueOrPanic() T {

	asserty.TrueWithMessage(g.initialized, "GoFutureOr must be initialized before calling ValueOrPanic()")

	if g.err != nil {
		panic(g.err)
	}
	return g.value
}
