package str

import (
	"strings"

	"github.com/xrash/smetrics"
)

//Similarity between 2 strings
func Similarity(a string, b string) float64 {
	return smetrics.Jaro(strings.ToLower(a), strings.ToLower(b)) * 100
}

//SimilarityNormalized compare and returns similarity between 2 strings
//ignoring accents
func SimilarityNormalized(a string, b string) float64 {
	return smetrics.Jaro(Normalize(strings.ToLower(a)), Normalize(strings.ToLower(b))) * 100
}
