package gofutureor

import (

	"github.com/fengdotdev/golibs-traits/trait"
)

func NewDTO[T any](value T, err error) GoFutureOrDTO[T] {
	return GoFutureOrDTO[T]{
		Value: value,
		Error: err,
	}
}

type GoFutureOrDTO[T any] struct {
	Value T
	Error error
}

var _ trait.DataTransferObject[GoFutureOrDTO[any]] = (*GoFutureOr[any])(nil)

// FromDTO implements trait.DataTransferObject.
func (g *GoFutureOr[T]) FromDTO(dto GoFutureOrDTO[T]) error {
	g.value = dto.Value
	g.err = dto.Error
	g.initialized = true
	return nil
}

// ToDTO implements trait.DataTransferObject.
func (g *GoFutureOr[T]) ToDTO() (GoFutureOrDTO[T], error) {

	return GoFutureOrDTO[T]{
		Value: g.value,
		Error: g.err,
	}, nil
}
