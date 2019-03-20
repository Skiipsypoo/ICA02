package fileanalysis

import "fmt"
import "unicode"

// FileAnalysis ...
func FileAnalysis(b []byte) {
	for _, c := range b {
		fmt.Printf("%c", c)
	}
	r := []rune(string(b))
	fmt.Println(getLanguage(r))
}

func getLanguage(r []rune) string {
	language := ""
	for _, c := range r {
		if unicode.Is(unicode.Cyrillic, c) {
			language = "The language is cyrillic (russian)"
			break
		} else if unicode.Is(unicode.Latin, c) {
			language = "Latin: Norwegian or Icelandic"
			break
		}
	}
	return language
}
