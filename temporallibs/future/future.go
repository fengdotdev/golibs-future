package future

import (
	"errors"

	"github.com/fengdotdev/golibs-helperfuncs/ident"
	"github.com/fengdotdev/golibs-helperfuncs/web"
	"github.com/fengdotdev/golibs-traits/trait"
)

type identifier = string
type Result[T any] = trait.Result[T]

var _ Result[string] = WebResource[string]{}

var _ Future[string] = WebResource[string]{}

var (
	ErrInvalidAccess = errors.New("invalid access to the future")
)

type Future[T any] interface {
	trait.Result[T]
}

type WebResource[T any] struct {
	id           string
	channel      chan T
	errorChannel chan error
	complete     chan bool
}
type WebResourceAccessor[T any] struct {
	Channel      *chan T
	ErrorChannel *chan error
	Complete     *chan bool
}

// with no scope asure access to the inners of the future
func (wr *WebResource[T]) Unlock(id identifier) (WebResourceAccessor[T], error) {
	if wr.id != id {
		return WebResourceAccessor[T]{}, ErrInvalidAccess
	}
	return WebResourceAccessor[T]{
		Channel:      &wr.channel,
		ErrorChannel: &wr.errorChannel,
		Complete:     &wr.complete,
	}, nil
}

func NewWebResource[T any](url string) (*WebResource[T], identifier) {

	id := ident.DeterministicUUID("webresource", url)

	return &WebResource[T]{
		id:           id,
		channel:      make(chan T, 1),
		errorChannel: make(chan error, 1),
		complete:     make(chan bool, 1),
	}, id
}

func FetchWebResource[T any](url string) *WebResource[[]byte] {
	resource, id := NewWebResource[[]byte](url)
	access, err := resource.Unlock(id)

	// this never should happen, all access should be from the caller
	if err != nil {
		panic(err)
	}

	go func() {

		defer close(*access.Channel)
		defer close(*access.ErrorChannel)
		defer close(*access.Complete)

		data, err := web.GetRemoteResource(url)
		if err != nil {
			access.ErrorChannel <- err
			return
		}

		resource.channel <- data
		resource.complete <- true

	
	}()
	return resource
}
