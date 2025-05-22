package gofutureor

type GoFutureOr[T any] struct {
	initialized bool
	value       T
	err         error
}
