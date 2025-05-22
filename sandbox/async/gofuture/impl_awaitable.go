package gofuture

import (
	"github.com/fengdotdev/golibs-future/sandbox/async"
)

var _ async.Awaitable[any] = (*GoFuture[any])(nil)

// Await implements async.Awaitable.
func (g *GoFuture[T]) Await() (result chan async.FutureOr[T]) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.initialized {
		panic("GoFuture not initialized")
	}

	resultchan := make(chan async.FutureOr[T], 1)

	g.doneRecipients = append(g.doneRecipients, resultchan)

	return resultchan
}
