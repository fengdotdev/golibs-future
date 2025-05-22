package async

import (

	"github.com/fengdotdev/golibs-traits/trait"
)

// must be a container of some T in the future
type Future[T any] interface {
	FutureOr[T]
	Awaitable[T]
	Completer[T]
	FutureOperations[T]
	FutureState[T]
}

type Awaitable[T any] interface {
	Await() (done chan bool, future Future[T])
}

type Completer[T any] interface {
	Complete(value T)
	CompleteWithError(err error)
	CompleteWith(value T, err error)
}

type FutureOperations[T any] interface {
	Then(onSuccess func(T))
	Catch(onError func(error))
	Finally(onCompletion func(T, error))
}

type FutureState[T any] interface {
	IsDone() bool
}

type FutureOr[T any] = trait.Result[T]
