package identicon

import (
	"testing"
	"crypto/md5"
	"reflect"
	"fmt"
)

var GeneratedPattern = []byte{
	129, 201, 130, 201, 129,
	56, 157, 244, 157, 56,
	232, 187, 241, 187, 232,
	183, 61, 165, 61, 183,
	149, 108, 179, 108, 149,
}

var GeneratedHash = md5.Sum([]byte("Culona"))

func TestItGeneratesAPatternFromAListOfBytes(t *testing.T) {
	pattern := generatePatternFromHash(GeneratedHash)

	if !reflect.DeepEqual(GeneratedPattern, pattern) {
		fmt.Printf("%v", pattern)
		t.Fatal("Failing asserting equality of pattern.")
	}
}
