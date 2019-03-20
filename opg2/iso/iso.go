package iso

import (
	"bytes"
	"fmt"
)

// GenerateExtendedASCIIStringLiteral ...
func GenerateExtendedASCIIStringLiteral() string {
	var buffer bytes.Buffer
	for i := '\u0080'; i <= '\u00FF'; i++ {
		buffer.WriteString(fmt.Sprintf("%c", i))
	}
	return buffer.String()
}

// IterateOverASCIIStringLiteral ...
func IterateOverExtendedASCIIStringLiteral(sl string) {
	for _, c := range sl {
		fmt.Printf("%X %q %b \n", c, c, c)
	}
}

// GreetingExtendedASCII ...
func GreetingExtendedASCII() string {
	greeting := "Salut, ça va °-) Κοστίζει ​ €50"
	for _, c := range []byte(greeting) {
		fmt.Printf("%c", c)
	}
	return greeting
}
