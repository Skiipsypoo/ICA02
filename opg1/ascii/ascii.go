package ascii

import (
	"bytes"
	"fmt"
)

// ASCII ...
const ASCII = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f"

// IterateOverASCIIStringLiteral runs a for loop based on stringLiteral length and
// prints out each character based in three formats: ascii hexadecimal, ascii symbol, ascii binary code
func IterateOverASCIIStringLiteral(sl string) {
	for _, c := range sl {
		fmt.Printf("%X %q %b \n", c, c, c)
		// Alternativ for %q er %c
	}
}

// GenerateASCIIStringLiteral generates an ASCII string by looping through the values from HEX 00 to HEX 7F
// or 0 to 127 and writing them to a string using the WriteString method from the bytes package
func GenerateASCIIStringLiteral() string {
	var buffer bytes.Buffer
	for i := '\u0000'; i <= '\u007F'; i++ {
		buffer.WriteString(fmt.Sprintf("%c", i))
	}
	return buffer.String()
}

// GreetingASCII ...
func GreetingASCII() string {
	greeting := "Hello :-)"
	for _, c := range []byte(greeting) {
		fmt.Printf("%c", c)
	}
	return greeting
}
