package webres

import "github.com/fengdotdev/golibs-future/helpers"

func NewWebResource[T any](url string) *WebResource {

	id := helpers.GenerateIdentifier(url)
	return &WebResource{
		id:           id,
		channel:      make(chan []byte, 1),
		errorChannel: make(chan error, 1),
		complete:     make(chan bool, 1),
	}
}
