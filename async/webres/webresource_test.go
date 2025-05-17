package webres_test

import (
	"fmt"
	"testing"

	"github.com/fengdotdev/golibs-future/async/webres"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestWebResource(t *testing.T) {

	t.Run("", func(t *testing.T) {
		wr := webres.NewWebResource[[]byte]("http://example.com")
		assert.NotNil(t, wr)

		fmt.Println(wr)

		select {
		case data := <-wr.GetChannel():
			assert.Nil(t, data)
			fmt.Println("data:", string(data))

		}

	})
}
