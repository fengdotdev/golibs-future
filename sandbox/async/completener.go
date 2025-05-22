package async

//A way to produce Future objects and to complete them later with a value or error.

// Completener is a type that represents a computation that can be completed in the future. and returns the promise or future of T
type Completener[T any] interface {
	Future() Future[T]
	Complete(value T)
	CompleteWithError(err error)
	CompleteWith(value T, err error)
	IsReady() bool
}
