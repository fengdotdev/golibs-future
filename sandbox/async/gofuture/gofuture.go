package gofuture

import (
	"sync"

	"github.com/fengdotdev/golibs-future/sandbox/async"
)

// ensure GoFuture implements async.Future
var _ async.Future[any] = (*GoFuture[any])(nil)

//  GoFuture is a container for a value that may not be available yet, a future value. 
type GoFuture[T any] struct {
	mu          sync.Mutex
	initialized bool
	cached      T
	cachedErr   error
	isCompleted bool
}
