package main

import (
	"unicode"
	"strings"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)


func CompareStringIgnoringUpperAndAccents(str1, str2 string) bool {
	str1Lower := strings.ToLower(str1);
	str2Lower := strings.ToLower(str2);
	str1LowerNoAccents := RemoveAccents(str1Lower);
	str2LowerNoAccents := RemoveAccents(str2Lower);
	return str1LowerNoAccents == str2LowerNoAccents;
}


func isMn(r rune) bool {
    return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

// Return the same string without accents
func RemoveAccents(str string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
    result, _, _ := transform.String(t, str)
    return result
}