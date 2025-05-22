package gofuture

import "github.com/fengdotdev/golibs-traits/trait"

// ensure GoFuture implements trait.Result
var _ trait.Result[any] = (*GoFuture[any])(nil)

// Error implements trait.Result.
func (g *GoFuture[T]) Error() error {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.cachedErr
}

// IsValid implements trait.Result.
func (g *GoFuture[T]) IsValid() bool {
	g.mu.Lock()
	defer g.mu.Unlock()

	if !g.initialized || !g.isCompleted || g.cachedErr != nil {
		return false
	}
	return true
}

// String implements trait.Result.
func (g *GoFuture[T]) String() string {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.isCompleted {
		return "Future completed"
	}
	return "Future pending"
}

// Value implements trait.Result.
func (g *GoFuture[T]) Value() T {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.IsValid() {
		return g.cached
	}

	// return zero value of T
	var Zero T
	return Zero
}

// ValueOr implements trait.Result.
func (g *GoFuture[T]) ValueOr(or T) T {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.IsValid() {
		return g.cached
	}
	return or
}

// ValueOrErr implements trait.Result.
func (g *GoFuture[T]) ValueOrErr() (T, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.IsValid() {
		return g.cached, nil
	}
	var Zero T
	return Zero, g.cachedErr
}

// ValueOrPanic implements trait.Result.
func (g *GoFuture[T]) ValueOrPanic() T {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.IsValid() {
		return g.cached
	}
	panic(g.cachedErr)
}
