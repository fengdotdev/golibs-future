package gopromise

import "github.com/fengdotdev/golibs-future/sandbox/async"

// Await implements async.Async.
func (g *GoAsync[T]) Await() (done chan bool, promise async.Future[T]) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.inizialized {
		panic("gofuture: not initialized")
	}

	d := make(chan bool)
	g.recipients = append(g.recipients, d)
	return d, g.promise
}

// Catch implements async.Async.
func (g *GoAsync[T]) Catch(catch func(error)) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.inizialized {
		panic("gofuture: not initialized")
	}

	g.catchs = append(g.catchs, catch)
}

// Finally implements async.Async.
func (g *GoAsync[T]) Finally(finally func(T, error)) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.inizialized {
		panic("gofuture: not initialized")
	}

	g.finallys = append(g.finallys, finally)
}

// Then implements async.Async.
func (g *GoAsync[T]) Then(then func(T)) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.inizialized {
		panic("gofuture: not initialized")
	}

	g.thens = append(g.thens, then)
}
