package gofutureor

func NewFutureOr[T any](value T) *GoFutureOr[T] {
	return &GoFutureOr[T]{
		value: value,
		err:   nil,
	}
}

func NewFutureOrError[T any](err error) *GoFutureOr[T] {
	var zero T
	return &GoFutureOr[T]{
		value: zero,
		err:   err,
	}
}
