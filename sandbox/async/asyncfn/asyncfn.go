package asyncfn

import (
	"github.com/fengdotdev/golibs-future/sandbox/async"
	 "github.com/fengdotdev/golibs-future/sandbox/async/goasync"
)

func AsyncFN[T any](fn func() (T, error)) async.Async[T] {
	future, complete := goasync.NewIncompleteGoAsync[T]()

	go func() {
		result, err := fn()
		complete(result, err)

	}()

	return future

}
