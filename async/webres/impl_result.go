package webres

import "github.com/fengdotdev/golibs-traits/trait"

var _ trait.Result[[]byte] = (*WebResource)(nil)

// Error implements async.Future.
func (w *WebResource) Error() error {
	if w.isCompleted {
		return w.chachedErr
	}

	select {
	case err := <-w.errorChannel:
		return err
	}
}

// IsValid implements async.Future.
func (w *WebResource) IsValid() bool {
	if w.isCompleted {
		return w.chachedErr == nil
	}

	select {
	case err := <-w.errorChannel:
		return err == nil
	default:
		return false
	}
}

// String implements async.Future.
func (w *WebResource) String() string {
	if w.isCompleted {
		return string(w.cached)
	}

	return "Future not completed"
}

// Value implements async.Future.
func (w *WebResource) Value() []byte {
	if w.isCompleted {
		return w.cached
	}
	select {
	case data := <-w.channel:
		return data
	}
}

// ValueOr implements async.Future.
func (w *WebResource) ValueOr(or []byte) []byte {
	if w.isCompleted {
		return w.cached
	}

	select {
	case data := <-w.channel:
		return data
	case err := <-w.errorChannel:
		if err != nil {
			return or
		}
	}

	return or
}

// ValueOrErr implements async.Future.
func (w *WebResource) ValueOrErr() ([]byte, error) {
	if w.isCompleted {
		return w.cached, w.chachedErr
	}

	select {
	case data := <-w.channel:
		return data, nil
	case err := <-w.errorChannel:
		return nil, err
	}
}

// ValueOrPanic implements async.Future.
func (w *WebResource) ValueOrPanic() []byte {
	if w.isCompleted {
		return w.cached
	}

	select {
	case data := <-w.channel:
		return data
	case err := <-w.errorChannel:
		panic(err)
	}

}
