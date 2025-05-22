package gofuture

import (
	"context"
	"sync"
	"time"

)

// ensure GoFuture implements async.Future

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
	doneRecipients []chan bool
	thens          []func(T)
	catchs         []func(error)
	finallys       []func(T, error)
	isCompleted    bool
}
