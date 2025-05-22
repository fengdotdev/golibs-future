package gofuture

import "github.com/fengdotdev/golibs-future/sandbox/async"

// ensure GoFuture implements async.trueFuture
var _ async.TrueFuture[any] = (*GoFuture[any])(nil)

// IsCompleted implements async.TrueFuture.
func (g *GoFuture[T]) IsCompleted() bool {
	g.mu.Lock()
	defer g.mu.Unlock()

	return g.isCompleted
}
