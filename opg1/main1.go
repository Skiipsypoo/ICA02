package main

import (
	"fmt"
	"ica02/opg1/ascii"
)

func main() {
	ascii.IterateOverASCIIStringLiteral(ascii.ASCII)
	fmt.Println(ascii.GreetingASCII())
}
