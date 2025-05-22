package gofutureor

import (
	"fmt"

	"github.com/fengdotdev/golibs-future/sandbox/async"
	"github.com/fengdotdev/golibs-helperfuncs/asserty"
	"github.com/fengdotdev/golibs-traits/trait"
)

var _ trait.Result[any] = (*GoFutureOr[any])(nil)
var _ async.FutureOr[any] = (*GoFutureOr[any])(nil)

// Error implements trait.Result.
func (g *GoFutureOr[T]) Error() error {
asserty.AssertTrue(g.err != nil, "Error() called on a valid result")
	return g.err
}

// IsValid implements trait.Result.
func (g *GoFutureOr[T]) IsValid() bool {
	return g.err == nil
}

// String implements trait.Result.

func (g *GoFutureOr[T]) String() string {
	if g.err != nil {
		return g.err.Error()
	}
	return fmt.Sprintf("%v", g.value)
}

// Value implements trait.Result.
func (g *GoFutureOr[T]) Value() T {
	if g.err != nil {
		var zero T
		return zero
	}
	return g.value
}

// ValueOr implements trait.Result.
func (g *GoFutureOr[T]) ValueOr(or T) T {
	if g.err != nil {
		return or
	}
	return g.value
}

// ValueOrErr implements trait.Result.
func (g *GoFutureOr[T]) ValueOrErr() (T, error) {
	panic("unimplemented")
}

// ValueOrPanic implements trait.Result.
func (g *GoFutureOr[T]) ValueOrPanic() T {
	if g.err != nil {
		panic(g.err)
	}
	return g.value
}
