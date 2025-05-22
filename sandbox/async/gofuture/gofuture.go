package gofuture

import (
	"context"
	"sync"
	"time"

	"github.com/fengdotdev/golibs-future/sandbox/async"
)

// ensure GoFuture implements async.Future

var _ async.Future[any] = (*GoFuture[any])(nil)

// GoFuture is a container for a value that may not be available yet, a future value.
type GoFuture[T any] struct {
	mu             sync.Mutex
	initialized    bool
	id             string
	cached         T
	cachedErr      error
	timeout        time.Duration
	ctxTimeout     context.Context
	ctxCancel      context.CancelFunc
	doneRecipients []chan async.FutureOr[T]
	thens          []func(T)
	catchs         []func(error)
	finallys       []func(T, error)
	isCompleted    bool
}
