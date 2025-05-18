package webres_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/fengdotdev/golibs-future/async/webres"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestWebResource(t *testing.T) {

	t.Run("", func(t *testing.T) {
		wr := webres.NewWebResource("http://example.com")
		assert.NotNil(t, wr)

		fmt.Println(wr)

		wr.Then(func(data []byte) {
			fmt.Println("Then:", string(data))
		})

		// sleep for 3 seconds to allow the async operation to complete
		time.Sleep(3 * time.Second)
	})
}
