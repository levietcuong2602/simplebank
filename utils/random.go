package utils

import (
	"math/rand"
	"strings"
	"time"
)

const letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generate random a integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generate random a string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(letter)

	for i := 0; i < n; i++ {
		ch := letter[RandomInt(0, int64(k))]
		sb.WriteByte(ch)
	}

	return sb.String()
}
