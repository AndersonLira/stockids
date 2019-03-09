package str

import (
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

//Normalize remove accents from string. Example 'çáŚÉ' returns 'caSE'
func Normalize(s string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
