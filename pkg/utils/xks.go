package utils

import (
	"net/http"
	"unicode"
)

func GetExternalKeyId(req *http.Request) string {
	return req.PathValue("externalKeyId")
}

func ReplaceNonPrintable(input string) string {
	result := []rune(input) // Convert string to rune slice to handle Unicode
	for i, r := range result {
		if !unicode.IsPrint(r) { // Check if the rune is printable
			result[i] = '?' // Replace non-printable with '.'
		}
	}
	return string(result) // Convert back to string
}
