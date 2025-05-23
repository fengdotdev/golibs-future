package gofutureor

// NewFutureOr creates a new GoFutureOr with the provided value.
func NewFutureOr[T any](value T) *GoFutureOr[T] {
	return &GoFutureOr[T]{
		initialized: true,
		value:       value,
		err:         nil,
	}
}

// NewFutureOrErr creates a new GoFutureOr with the provided error.
func NewFutureOrError[T any](err error) *GoFutureOr[T] {
	var zero T
	return &GoFutureOr[T]{
		initialized: true,
		value:       zero,
		err:         err,
	}
}

// NewFutureOrZero creates a new GoFutureOr with the zero value of T. (var zero T)
func NewFutureOrZero[T any]() *GoFutureOr[T] {
	var zero T
	return &GoFutureOr[T]{
		initialized: true,
		value:       zero,
		err:         nil,
	}
}

// / NewFutureOrZeroValue creates a new GoFutureOr with the provided value.
func NewFutureOrZeroValue[T any](zeroValue T) *GoFutureOr[T] {
	return &GoFutureOr[T]{
		initialized: true,
		value:       zeroValue,
		err:         nil,
	}
}

// NewFutureOrZeroFN creates a new GoFutureOr with the provided zero value function.
func NewFutureOrZeroFN[T any](zeroValueFn func() T) *GoFutureOr[T] {
	return &GoFutureOr[T]{
		initialized: true,
		value:       zeroValueFn(),
		err:         nil,
	}
}
