package main

import (
	"fmt"
	"ica02/opg5/slice"
)

func main() {
	byteslice1 := make([]byte, 500)
	fmt.Println("&byteslice1[0]")
	fmt.Println(&byteslice1[0])
	aslice := slice.Reslice(byteslice1, 50, 100)
	fmt.Println("&aslice[0]")
	fmt.Println(&aslice[0])
	fmt.Println("&byteslice1[50]")
	fmt.Println(&byteslice1[50])
}
