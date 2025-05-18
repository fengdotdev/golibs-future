package webres

import "github.com/fengdotdev/golibs-traits/trait"

var _ trait.Result[[]byte] = (*WebResource)(nil)

// Error implements async.Future.
func (w *WebResource) Error() error {
	panic("not implemented")
}

// IsValid implements async.Future.
func (w *WebResource) IsValid() bool {
	panic("not implemented")
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
	panic("not implemented")
}

// ValueOr implements async.Future.
func (w *WebResource) ValueOr(or []byte) []byte {
	panic("not implemented")
}

// ValueOrErr implements async.Future.
func (w *WebResource) ValueOrErr() ([]byte, error) {
	panic("not implemented")
}

// ValueOrPanic implements async.Future.
func (w *WebResource) ValueOrPanic() []byte {
	panic("not implemented")

}
