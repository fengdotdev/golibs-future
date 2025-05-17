package async



// AsyncOperation represents an asynchronous operation that can yield a result or an error.
type AsyncOperation[T any] interface {
	Await() (T, error)
}

// Completer is used to control the completion of an AsyncOperation.
type Completer[T any] struct {
	future *future[T]
}

// NewCompleter creates a new Completer and its associated AsyncOperation.
func NewCompleter[T any]() (*Completer[T], AsyncOperation[T]) {
	f := &future[T]{
		result: make(chan result[T]),
	}
	return &Completer[T]{future: f}, f
}

// CompleteWithValue completes the AsyncOperation with a value.
/* func (c *Completer[T]) CompleteWithValue(value T) {
	select {
	case c.future.result <- result[T]{value: value}:
	default:
		// Handle potential multiple completions (optional)
	}
}

// CompleteWithError completes the AsyncOperation with an error.
func (c *Completer[T]) CompleteWithError(err error) {
	select {
	case c.future.result <- result[T]{err: err}:
	default:
		// Handle potential multiple completions (optional)
	}
}
 */
type future[T any] struct {
	result <-chan result[T]
}

type result[T any] struct {
	value T
	err   error
}

// Await blocks until the AsyncOperation completes and returns the result or error.
func (f *future[T]) Await() (T, error) {
	r := <-f.result
	return r.value, r.err
}