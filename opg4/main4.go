package main

import (
	"fmt"
	"ica02/opg4/unicode"
)

// Variables used to generate hexadecimal representation strings
var no = []byte("nord og sør")
var is = []byte("norður og suður")
var jp = []byte("北と南")

func main() {
	unicode.PrintHexString(no)
	unicode.PrintHexString(is)
	unicode.PrintHexString(jp)
	fmt.Println(unicode.Translate("nord og sør", "jp"))
	fmt.Println(unicode.Translate("nord og sør", "is"))
	unicode.UnicodeCodePointDemo()
}
