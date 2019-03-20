package main

import (
	"fmt"
	"ica02/fileutils"
	"ica02/opg3/fileanalysis"
	"ica02/opg3/treasure"
)

var file1, file2, file3 = "../files/lang01.wl",
	"../files/lang02.wl",
	"../files/lang03.wl"

func main() {
	decodeTreasure()
}

func analyseFiles() {
	byteSlice1 := fileutils.FileToByteSlice(file1)
	byteSlice2 := fileutils.FileToByteSlice(file2)
	byteSlice3 := fileutils.FileToByteSlice(file3)
	fileanalysis.FileAnalysis(byteSlice1)
	fileanalysis.FileAnalysis(byteSlice2)
	fileanalysis.FileAnalysis(byteSlice3)
}

func decodeTreasure() {
	treasureByteSlice := fileutils.FileToByteSlice("./treasure/treasure.txt")
	treasureStr := string(treasureByteSlice)
	treasure.PrintTreasureUTF8(treasureStr)
}

func testThings() {
	s := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("Printf with %s:")
	fmt.Printf("%s\n", s)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", s)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", s)

	fmt.Println("Printf with %c:")
	fmt.Printf("%c\n", s)

	fmt.Printf("%s", "\xc2\xbd\x3f\x3d\x3f\xe2\x8c\x98")
}
