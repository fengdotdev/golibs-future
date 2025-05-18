package webres

import (
	"context"
	"time"

	"github.com/fengdotdev/golibs-future/helpers"
)

func NewWebResource(url string) *WebResource {

	id := helpers.GenerateIdentifier(url)

	ctxTimeout, cancel := context.WithTimeout(context.Background(), TimeOut)
	return &WebResource{
		id:              id,
		timeout:         TimeOut,
		channel:         make(chan []byte, 1),
		errorChannel:    make(chan error, 1),
		completeChannel: make(chan bool, 1),
		cached:          nil,
		chachedErr:      nil,
		isCompleted:     false,
		ctxTimeout:      ctxTimeout,
		ctxCancel:       cancel,
	}
}

func NewZeroWebResource() *WebResource {
	return &WebResource{
		id:              "",
		timeout:         0,
		channel:         make(chan []byte, 1),
		errorChannel:    make(chan error, 1),
		completeChannel: make(chan bool, 1),
		cached:          nil,
		chachedErr:      nil,
		isCompleted:     false,
		ctxTimeout:      context.Background(),
		ctxCancel:       func() {},
	}
}

func NewWebResourceWithTimeout(url string, timeout time.Duration) *WebResource {
	id := helpers.GenerateIdentifier(url)

	ctxTimeout, cancel := context.WithTimeout(context.Background(), timeout)
	return &WebResource{
		id:              id,
		timeout:         timeout,
		channel:         make(chan []byte, 1),
		errorChannel:    make(chan error, 1),
		completeChannel: make(chan bool, 1),
		cached:          nil,
		chachedErr:      nil,
		isCompleted:     false,
		ctxTimeout:      ctxTimeout,
		ctxCancel:       cancel,
	}
}

func NewWebResourceComplete(url string, data []byte, err error) *WebResource {
	id := helpers.GenerateIdentifier(url)
	ok := make(chan bool, 1)

	defer func() {
		ok <- true
		close(ok)
	}()
	//ctxTimeout, cancel := context.WithTimeout(context.Background(), TimeOut)
	return &WebResource{
		id:              id,
		timeout:         0,
		channel:         nil,
		errorChannel:    nil,
		completeChannel: ok,
		cached:          data,
		chachedErr:      err,
		isCompleted:     true,
		ctxTimeout:      nil,
		ctxCancel:       nil,
	}
}
