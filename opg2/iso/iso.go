package iso

import (
	"bytes"
	"fmt"
)

// GenerateExtendedASCIIStringLiteral ...
func GenerateExtendedASCIIStringLiteral() string {
	var buffer bytes.Buffer
	for i := '\u007F'; i <= '\u00FF'; i++ {
		buffer.WriteString(fmt.Sprintf("%c", i))
	}
	return buffer.String()
}

// IterateOverASCIIStringLiteral ...
func IterateOverASCIIStringLiteral(sl string) {
	for _, c := range sl {
		fmt.Printf("%X %q %b \n", c, c, c)
	}
}

// GreetingExtendedASCII ...
func GreetingExtendedASCII() string {
	return "Salut, ça va °-) Κοστίζει ​ €50"
}
