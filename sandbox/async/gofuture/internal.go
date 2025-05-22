package gofuture

func (g *GoFuture[T]) Complete(value T, err error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.cached = value
	g.cachedErr = err
	g.isCompleted = true
}
