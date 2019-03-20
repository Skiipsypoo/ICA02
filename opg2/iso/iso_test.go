package iso

import (
	"fmt"
	"testing"
	"unicode"
)

func TestGreetingExtendedASCII(t *testing.T) {
	str := GreetingExtendedASCII()
	for _, c := range str {
		if c > '\u00FF' || c < unicode.MaxASCII {
			fmt.Printf("%c", c)
			t.Error("Is not extended ASCII: ", c)
		}
	}
}
