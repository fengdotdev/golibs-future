package async

import "github.com/fengdotdev/golibs-traits/trait"

type Future[T any] interface {
	trait.Result[T]
}
