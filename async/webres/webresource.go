package webres

import (
	"github.com/fengdotdev/golibs-future/helpers"
)

type WebResource struct {
	id           string
	channel      chan []byte
	errorChannel chan error
	complete     chan bool
}
type WebResourceAccessor struct {
	Channel      *chan []byte
	ErrorChannel *chan error
	Complete     *chan bool
}

func NewWebResource[T any](url string) *WebResource {

	id := helpers.GenerateIdentifier(url)
	return &WebResource{
		id:           id,
		channel:      make(chan []byte, 1),
		errorChannel: make(chan error, 1),
		complete:     make(chan bool, 1),
	}
}

func FetchWebResource[T any](url string) *WebResource {
	resource := NewWebResource[[]byte](url)

	go func() {

		defer close(resource.channel)
		defer close(resource.errorChannel)
		defer close(resource.complete)

		data, err := helpers.FetchURL(url)
		if err != nil {
			resource.errorChannel <- err
			return
		}

		resource.channel <- data
		resource.complete <- true

	}()
	return resource
}
