package helpers_test

import (
	"fmt"
	"testing"

	"github.com/fengdotdev/golibs-future/sandbox/helpers"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestIdentifier(t *testing.T) {

	someString := "some string"
	identifier := helpers.GenerateIdentifier(someString)
	fmt.Println(identifier)

	assert.NotEqual(t, identifier, "")
	assert.Equal(t, len(identifier), 32)
}
