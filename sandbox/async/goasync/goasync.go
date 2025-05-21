package goasync

import (
	"context"
	"sync"
	"time"

	"github.com/fengdotdev/golibs-future/sandbox/async"
	"github.com/fengdotdev/golibs-future/sandbox/async/gopromise"
)

var _ async.Async[any] = (*GoAsync[any])(nil)

var timeout = time.Millisecond * 1000 // default 1 second

func NewIncompleteGoAsync[T any]() (goasync *GoAsync[T], completefn func(T, error)) {
	var Zero T

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)

	p := &GoAsync[T]{
		inizialized: true,
		promise:     gopromise.NewGoPromise(Zero, gopromise.ErrPromiseNotResolved),
		mu:          sync.Mutex{},
		timeout:     timeout,
		id:          "",
		ctxTimeout:  ctx,
		ctxCancel:   cancel,
	}

	update := func(value T, err error) {
		p.mu.Lock()
		defer p.mu.Unlock()
		p.promise = gopromise.NewGoPromise(value, err)

		go func() {

			if err == nil {
				p.triggerThen(value)
			} else {
				p.triggerCatch(err)
			}
			p.triggerFinally(value, err)
		}()

		go func() {
			for _, recipient := range p.recipients {
				recipient <- true

				defer close(recipient)
			}
		}()
		p.isCompleted = true
	}

	return p, update
}

// NewGoAsync creates a new GoAsync Complete instance with the given value and error.
// It initializes the GoAsync instance with the provided value and error,
func NewGoAsync[T any](value T, err error) *GoAsync[T] {
	return &GoAsync[T]{
		inizialized: true,
		timeout:     0,
		id:          "",
		promise:     gopromise.NewGoPromise(value, err),
		mu:          sync.Mutex{},
		ctxTimeout:  context.Background(),
		ctxCancel:   func() {},
		recipients:  make([]chan bool, 0),
		thens:       make([]func(T), 0),
		catchs:      make([]func(error), 0),
		finallys:    make([]func(T, error), 0),
		isCompleted: true,
	}
}

type GoAsync[T any] struct {
	inizialized bool
	timeout     time.Duration
	id          string
	promise     async.Promise[T]
	mu          sync.Mutex
	ctxTimeout  context.Context
	ctxCancel   context.CancelFunc
	recipients  []chan bool
	thens       []func(T)
	catchs      []func(error)
	finallys    []func(T, error)
	isCompleted bool
}

func (g *GoAsync[T]) triggerThen(value T) {
	g.mu.Lock()
	defer g.mu.Unlock()

	for _, then := range g.thens {
		then(value)
	}

}

func (g *GoAsync[T]) triggerCatch(err error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	for _, catch := range g.catchs {
		catch(err)
	}
}

func (g *GoAsync[T]) triggerFinally(value T, err error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	for _, finally := range g.finallys {
		finally(value, err)
	}
}
