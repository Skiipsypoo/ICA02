package slice

//import "fmt"

// AllocateVar har INN-argument b
// b - antall bytes brukeren ønsker å allokere
// Returnerer en slice av type []byte
//
func AllocateVar(b int) []byte {
	// Kode for Oppgave 5a
	var slice []byte
	slice = make([]byte, b)
	return slice
}

// AllocateMake tar lengde og kapasitet som b og lager en ny slice
//
func AllocateMake(b int) []byte {
	// Kode for Oppgave 5a
	slice := make([]byte, b)
	return slice
}

// Reslice takes a slice and reslices it
func Reslice(slc []byte, lidx int, uidx int) []byte {
	slice := AllocateMake(uidx)
	slice = slc[lidx:uidx]
	return slice
}

// CopySlice ???
func CopySlice(slice []byte) []byte {
	// Source https://github.com/go101/go101/wiki/How-to-perfectly-clone-a-slice%3F
	copy := append(slice[:0:0], slice...)
	return copy
}
