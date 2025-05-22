package gofuture

/* func (g *GoFuture[T]) Complete(value T, err error) {

}
*/

func (g *GoFuture[T]) complete(value T, err error) {
	

	if g.initialized {
		panic("GoFuture already initialized")
	}

	g.value = value
	g.err = err

	g.initialized = true

	for _, recipient := range g.doneRecipients {
		recipient <- &GoFutureOr[T]{value: value, err: err}
		close(recipient)
	}
}
