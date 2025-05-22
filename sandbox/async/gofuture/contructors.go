package gofuture

import "time"

func NewFuture[T any](computation func() (T, error)) *GoFuture[T] {
	panic("not implemented")
}

func NewFutureDelayed[T any](computation func() (T, error), delay time.Duration) *GoFuture[T] {
	panic("not implemented")
}

func NewFutureValue[T any](value T) *GoFuture[T] {
	panic("not implemented")
}

func NewFutureError[T any](err error) *GoFuture[T] {
	panic("not implemented")
}

func NewFutureSync[T any](computation func() (T, error)) *GoFuture[T] {
	panic("not implemented")
}
