package slice

//import "fmt"

// AllocateVar returnerer en ny slice med lengde og kapasitet av parameter int b
func AllocateVar(b int) []byte {
	var slice []byte
	slice = make([]byte, b)
	return slice
}

// AllocateMake returnerer en ny slice med lengde og kapasitet av parameter int b
func AllocateMake(b int) []byte {
	slice := make([]byte, b)
	return slice
}

// Reslice takes a slice and reslices it
func Reslice(slc []byte, lidx int, uidx int) []byte {
	slice := AllocateMake(uidx)
	slice = slc[lidx:uidx]
	return slice
}

// CopySlice copies a slice
// Source https://github.com/go101/go101/wiki/How-to-perfectly-clone-a-slice%3F
func CopySlice(slice []byte) []byte {
	copy := append(slice[:0:0], slice...)
	return copy
}
