package gofuture

import "github.com/fengdotdev/golibs-future/sandbox/async"

var _ async.FutureOperations[any] = (*GoFuture[any])(nil)

// Catch implements async.FutureOperations.
func (g *GoFuture[T]) Catch(onError func(error)) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if !g.initialized {
		panic("gofuture: not initialized")
	}

	g.catchs = append(g.catchs, onError)

}

// Finally implements async.FutureOperations.
func (g *GoFuture[T]) Finally(onCompletion func(T, error)) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if !g.initialized {
		panic("gofuture: not initialized")
	}

	g.finallys = append(g.finallys, onCompletion)
}

// Then implements async.FutureOperations.
func (g *GoFuture[T]) Then(onSuccess func(T)) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if !g.initialized {
		panic("gofuture: not initialized")
	}

	g.thens = append(g.thens, onSuccess)
}
