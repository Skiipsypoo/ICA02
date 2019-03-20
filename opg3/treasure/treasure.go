package treasure

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

// PrintTreasureUTF8 tar en "string literal" som INN-data
// og skriver ut en korrekt text på med UTF-8 koding
// Koden er for Oppgave 3c
// Bruk strengen fra filen treasure.txt som in-data for denne funksjonen
func PrintTreasureUTF8(treasureString string) {
	// Replace UTF-16 'æ' with UTF-8 'æ'
	treasureString = strings.Replace(treasureString, "\\xe6", "\\xc3\\xa6", -1)
	// Replace UTF-16 'ø' with UTF-8 'ø'
	treasureString = strings.Replace(treasureString, "\\xf8", "\\xc3\\xb8", -1)
	// Replace UTF-16 'å' with UTF-8 'å'
	treasureString = strings.Replace(treasureString, "\\xe5", "\\xc3\\xa5", -1)
	// Remove all backslash + x to use hex.DecodeString function to print string
	treasureString = strings.Replace(treasureString, "\\x", "", -1)

	decoded, err := hex.DecodeString(treasureString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", decoded)
}
