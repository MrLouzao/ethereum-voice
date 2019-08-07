package main

import (
	"unicode"
	"strings"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"fmt"
	"math"
	"strconv"
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


// Cast a number in hexadecimal to ETHER units
func HexAmountWeiToEther(value string) float64 {
	res, _ := strconv.ParseInt(value, 0, 64)
	exponentOfEtherInWeis := math.Pow(10, -18)
	amountInEther := float64(res) * exponentOfEtherInWeis
	return amountInEther
}


func EtherFloatToWeiHex(value float64) string {
	exponentOfEtherInWeis := math.Pow(10, 18)
	amountInWeis := value * exponentOfEtherInWeis
	amountHex := fmt.Sprintf("0x%x", int(amountInWeis))
	return amountHex
}