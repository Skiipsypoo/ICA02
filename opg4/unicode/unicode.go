package unicode

import (
	"fmt"
)

// PrintHexString ...
func PrintHexString(byteSlice []byte) {
	for _, c := range byteSlice {
		fmt.Printf("%s%X", "\\x", c)
	}
	fmt.Println()
}

// Translate ...
func Translate(expression string, language string) string {
	str := "\"\x6E\x6F\x72\x64\x20\x6F\x67\x20\x73\xC3\xB8\x72\""
	switch language {
	case "is":
		str += " på islandsk er " + "\"\x6E\x6F\x72\xC3\xB0\x75\x72\x20\x6F\x67\x20\x73\x75\xC3\xB0\x75\x72\""
	case "jp":
		str += " på japansk er " + "\"\xE5\x8C\x97\xE3\x81\xA8\xE5\x8D\x97\""
	}
	return str
}

func getHexString(byteSlice []byte) string {
	str := ""
	for _, c := range byteSlice {
		str += fmt.Sprintf("%s%X", "\\x", c)
	}
	return str
}

// UnicodeCodePointDemo ...
func UnicodeCodePointDemo() {
	// Hva er dette?
	// Er det likt på MS Windows og macOS?
	fmt.Println("\xf0\x9F\x98\x80")
	fmt.Println("\xf0\x9F\x98\x97")
	// Demonstrerer at deler av et tegn representert med flere bytes
	// kan ikke tolkes innenfor koden (unicode)
	fmt.Println("\xf0\x9F\x98")
	fmt.Println("\xf0\x9F")
	fmt.Println("\xf0")
}
