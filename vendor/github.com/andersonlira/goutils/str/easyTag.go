package str

import (
	"strings"
)

//EasyTag return a tag array when receive string field
//Example EasyTag("A,'B,C',D",",","'") return "A","B,C","D" values
func EasyTag(s, separator, textMark string) (tags []string) {
	pipe := "|"
	if textMark == "|" {
		pipe = "'"
	}
	aux := s
	first := -1
	limit := 0
	for p := strings.Index(aux, textMark); p > -1 && limit < 10; {
		if first == -1 {
			first = p
			aux = strings.Replace(aux, textMark, pipe, 1)
		} else {
			cut := aux[first : p+1]
			t := strings.Replace(cut, pipe, "", 1)
			t = strings.Replace(t, textMark, "", 1)
			tags = append(tags, t)
			aux = strings.Replace(aux, cut, "", 1)
			first = -1
		}
		p = strings.Index(aux, textMark)
		limit++
	}
	for _, r := range strings.Split(aux, separator) {
		if r != "" {
			r = strings.Replace(r, pipe, textMark, -1)
			tags = append(tags, r)
		}
	}

	return
}
