package webres

import (

	"github.com/fengdotdev/golibs-future/async"
)

var _ async.Future[[]byte] = (*WebResource)(nil)

var (
	TimeOut = 1000 // 1 second
)

type WebResource struct {
	timeout      int
	id           string
	channel      chan []byte
	errorChannel chan error
	complete     chan bool
	cached       []byte
	chachedErr   error
	isCompleted  bool
}
