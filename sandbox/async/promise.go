package async

import "github.com/fengdotdev/golibs-traits/trait"

type Promise[T any] interface {
	trait.Result[T]
	IsCompleted() bool
}
