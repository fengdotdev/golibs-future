package async

type Async[T any] interface {
	Await() (done chan bool, promise Promise[T])
	Then(then func(T))
	Catch(catch func(error))
	Finally(finally func(T, error))
}
