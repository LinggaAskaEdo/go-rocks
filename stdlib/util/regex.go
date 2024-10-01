package util

import "regexp"

var (
	RegexEncryptedValue = regexp.MustCompile(`^ENC\((.*?)+\)`)
)

func ExtractEncryptedValue(msg string) string {
	return RegexEncryptedValue.ReplaceAllString(msg, "$1")
}
