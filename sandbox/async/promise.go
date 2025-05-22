package async

import "github.com/fengdotdev/golibs-traits/trait"

// Completener is a type that represents a computation that can be completed in the future. and returns the promise or future of T
type Completener[T any] interface {
	Await() (done chan bool, future Future[T])
	Operation()
	FutureOperations[T]
	IsDone() bool
}

type FutureOperations[T any] interface {
	Then(onSuccess func(T))
	Catch(onError func(error))
	Finally(onCompletion func(T, error))
}

// must be a container of some T in the future
type Future[T any] interface {
	trait.Result[T]
	TrueFuture[T]
}

type TrueFuture[T any] interface {
	IsCompleted() bool
	Complete(value T, err error)
}
