// Package generator provides utilities for generating random data used in network event creation.
// This package contains functions for creating randomized strings, dates, and other primitive types
// required for realistic event data generation at scale.
package generator

import (
	"math/rand"
	"time"
)

// letterRunes defines the character set used for random string generation.
// Uses base64 encoding alphabet (A-Z, a-z, 0-9, +, /) providing 64 possible characters.
// This ensures generated strings are URL-safe and database-friendly.
var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

// RandomString generates a random alphanumeric string with a variable length (0-40 characters).
func RandomString() string {
	strLen := rand.Int31n(40)

	var str string
	for i := 0; i <= int(strLen); i++ {
		str = str + string(letterRunes[int(rand.Int31n(int32(len(letterRunes))))])
	}
	return str
}

// RandomDate generates a random timestamp within a 10-year window (2010-2020).
func RandomDate() *time.Time {
	t := time.Date(rand.Intn(11)+2010, time.Month(rand.Intn(12)+1), rand.Intn(28), rand.Intn(23), rand.Intn(59), rand.Intn(59), rand.Intn(59), time.UTC)
	return &t
}
