package main

import (
	"fmt"
	"ica02/opg1/ascii"
	"ica02/opg2/iso"
)

func main() {
	iso.IterateOverASCIIStringLiteral(iso.GenerateExtendedASCIIStringLiteral())
	fmt.Println(ascii.GreetingASCII())
}
