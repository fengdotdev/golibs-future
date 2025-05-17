package helpers

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/fengdotdev/golibs-future/def"
)

func GenerateIdentifier(input string) def.Identifier {
	hash := md5.Sum([]byte(input))

	hashString := hex.EncodeToString(hash[:])

	return hashString
}
