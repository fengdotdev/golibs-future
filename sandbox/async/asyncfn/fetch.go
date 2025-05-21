package asyncfn

import (
	"github.com/fengdotdev/golibs-future/sandbox/async"
	"github.com/fengdotdev/golibs-future/sandbox/helpers"
)

func Feth(url string) async.Async[[]byte] {

	return AsyncFN(
		func() ([]byte, error) {
			return helpers.FetchURL(url)
		},
	)
}
