package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func init() {
}

// Genereates random number between min and max
func RandomInt(min, max int64) int64 {
	return min + rng.Int63n(max-min+1)
}

// Generates random string of leghth n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rng.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// Generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Generates a random currency
func RandomCurrency() string {
	currencies := []string{EUR, USD, CAD}
	n := len(currencies)
	return currencies[rng.Intn(n)]
}

func RandomEmail() string {
	return RandomString(6) + "@gmail.com"
}

func RandomPassword() string {
	return RandomString(6)
}
