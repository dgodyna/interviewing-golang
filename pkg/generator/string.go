package generator

import (
	"math/rand"
	"time"
)

// letterRunes holds standard base64 encoding dictionary for generation of random strings.
var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

// RandomString generates a random string of variable length (up to 40) using alphanumeric and special characters.
func RandomString() string {
	strLen := rand.Int31n(40)

	var str string
	for i := 0; i <= int(strLen); i++ {
		str = str + string(letterRunes[int(rand.Int31n(int32(len(letterRunes))))])
	}
	return str
}

// RandomDate generates a random date between 2010 and 2020 with random month, day, and time values in UTC.
func RandomDate() *time.Time {
	t := time.Date(rand.Intn(11)+2010, time.Month(rand.Intn(12)+1), rand.Intn(28), rand.Intn(23), rand.Intn(59), rand.Intn(59), rand.Intn(59), time.UTC)
	return &t
}
