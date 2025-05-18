package webres

import (
	"context"
	"time"

	"github.com/fengdotdev/golibs-future/async"
)

var _ async.Future[[]byte] = (*WebResource)(nil)

var (
	TimeOut = time.Millisecond * 1000 // 1 second
)

type WebResource struct {
	timeout         time.Duration
	id              string
	channel         chan []byte
	errorChannel    chan error
	completeChannel chan bool
	cached          []byte
	chachedErr      error
	isCompleted     bool
	ctxTimeout      context.Context
	ctxCancel       context.CancelFunc
}
