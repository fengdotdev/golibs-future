package gofuture

import "github.com/fengdotdev/golibs-future/sandbox/async"

var _ async.FutureState[any] = (*GoFuture[any])(nil)



// IsDone implements async.FutureState.
func (g *GoFuture[T]) IsDone() bool {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.initialized {
		panic("gofuture: not initialized")
	}

	return g.isCompleted
}
