package gofuture

import "github.com/fengdotdev/golibs-future/sandbox/async"

// ensure GoFuture implements async.completer
var _ async.Completer[any] = (*GoFuture[any])(nil)

// Complete implements async.Completer.
func (g *GoFuture[T]) Complete(value T) {
	g.CompleteWith(value, nil)
}

// CompleteWith implements async.Completer.
func (g *GoFuture[T]) CompleteWith(value T, err error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.cached = value
	g.cachedErr = err
	g.isCompleted = true
}

// CompleteWithError implements async.Completer.
func (g *GoFuture[T]) CompleteWithError(err error) {
	var Zero T
	g.CompleteWith(Zero, err)
}
