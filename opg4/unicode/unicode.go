package unicode

import (
	"encoding/hex"
	"fmt"
	"log"
)

// Variables used to generate hexadecimal representation strings
var no = []byte("nord og sør")
var is = []byte("norður og suður")
var jp = []byte("北と南")

// PrintHexString makes a HEX string from the
func printHexString(byteSlice []byte) {
	fmt.Print("HEX str used: ")
	for _, c := range byteSlice {
		fmt.Printf("%s%X", "\\x", c)
	}
	fmt.Println()
}

// Translate ...
func Translate(expression string, language string) string {
	str := getHexString([]byte(expression))
	if str != string(no) {
		return "unexpected expression"
	}
	switch language {
	case "is":
		str += " på islandsk er " + getHexString(is)
		printHexString(is) // Prints "\x6E\x6F\x72\xC3\xB0\x75\x72\x20\x6F\x67\x20\x73\x75\xC3\xB0\x75\x72"
	case "jp":
		str += " på japansk er " + getHexString(jp)
		printHexString(jp) // Prints \xE5\x8C\x97\xE3\x81\xA8\xE5\x8D\x97
	}
	return str
}

func getHexString(byteSlice []byte) string {
	str := ""
	for _, c := range byteSlice {
		str += fmt.Sprintf("%X", c)
	}
	decoded, err := hex.DecodeString(str)
	if err != nil {
		log.Fatal(err)
	}
	return string(decoded)
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
