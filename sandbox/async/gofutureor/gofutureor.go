package gofutureor

//	GoFutureOr is a generic type that implements the FutureOr interface.
//
// this is inmutable and can be used to represent a value that may or may not be present.
// It is designed to be used with the async package and provides methods to handle the value and error state.

// if is not initialized, it will panic if some methods are called. use the constructor for zeros to create the zero value obj.
// this type is not thread safe.
type GoFutureOr[T any] struct {
	initialized bool
	value       T
	err         error
}
