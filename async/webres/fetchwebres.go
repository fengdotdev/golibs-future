package webres

import (
	"errors"

	"github.com/fengdotdev/golibs-future/helpers"
)

func FetchWebResource(url string) *WebResource {
	resource := NewWebResource(url)

	fnlast := func() {
		close(resource.channel)
		close(resource.errorChannel)
		close(resource.completeChannel)

		if resource.cached == nil && resource.chachedErr == nil {
			resource.chachedErr = errors.New("something went wrong")
		}
		resource.isCompleted = true
	}

	fnTimeout := func() {
		// todo
		err := errors.New("timeout")
		resource.channel <- nil
		resource.errorChannel <- err
		resource.chachedErr = err
		resource.completeChannel <- true
		resource.isCompleted = true
	}

	go func() {

		defer fnlast()
		go func() {
			data, err := helpers.FetchURL(url)
			if err != nil {
				resource.errorChannel <- err
				resource.chachedErr = err
				resource.completeChannel <- true
				resource.isCompleted = true
				return
			}

			resource.channel <- data
			resource.cached = data
			resource.completeChannel <- true
			resource.isCompleted = true

		}()

		select {
		case <-resource.ctxTimeout.Done():
			fnTimeout()
			return
		case <-resource.completeChannel:
			return
		}
	}()
	return resource
}
