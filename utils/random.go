package utils

import (
	"math/rand"
	"strings"
)

const randomStringChars string = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func CreateRandomString(length int) string {
	sb := strings.Builder{}
	for range length {
		idx := rand.Int() % len(randomStringChars)
		sb.WriteByte(randomStringChars[idx])
	}
	return sb.String()
}
