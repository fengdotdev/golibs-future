package future_test

import (
	"testing"

	"github.com/fengdotdev/golibs-future/temporallibs/future"
)

type WebResource = future.Result[string]

func NewWebResource() WebResource {
	// Simulate a web resource
	return future.Result[string]{}
}

func SomeAPICall() WebResource {
	// Simulate an API call
	return future.Result[string]{}
}

func TestFuture(t *testing.T) {

}
