package webres

import (
	"errors"

	"github.com/fengdotdev/golibs-future/helpers"
)

func FetchWebResource(url string) *WebResource {
	resource := NewWebResource[[]byte](url)

	go func() {

		defer close(resource.channel)
		defer close(resource.errorChannel)
		defer close(resource.complete)
		defer func() {
			if resource.cached == nil {
				resource.chachedErr = errors.New("something went wrong")
			}
			resource.isCompleted = true
		}()

		data, err := helpers.FetchURL(url)
		if err != nil {
			resource.errorChannel <- err
			resource.chachedErr = err
			return
		}

		// single source of truth
		resource.channel <- data
		resource.cached = data
		resource.complete <- true

	}()
	return resource
}
