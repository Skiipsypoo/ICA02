package main

import (
	"fmt"
	"ica02/opg4/unicode"
)

// Variables used to generate hexadecimal representation strings

func main() {
	fmt.Println(unicode.Translate("nord og sr", "jp"))
	fmt.Println(unicode.Translate("nord og sør", "is"))
}
