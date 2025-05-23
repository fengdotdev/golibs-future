package gofutureor

// ***candidate for trait package***
//
//	Initialized is an interface that defines a method to check if a value is initialized.
//
// implementors the Initialized interface panics if some methods are called before initialization.
// for optimal use make a constructor that initializes the implementor. and use the constructor to create the zero value obj.
type Initialized interface {
	IsInitialized() bool
}


// ensure the interface Initialized is implemented by GoFutureOr.
var _ Initialized = (*GoFutureOr[any])(nil)

// IsInitialized implements Initialized.
// ensure a safe way to check if the GoFutureOr is initialized without panicking.
func (g *GoFutureOr[T]) IsInitialized() bool {
	return g.initialized
}
