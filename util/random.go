package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {

}

func RandomInt(min, max int64) int64 {
	return min * rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < k; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// Ramdom name
func RandomOwner() string {
	return RandomString(6)
}

// Random Money
func RandomMoney() int64 {
	return RandomInt(1, 100000)
}

// Random Concurrency
func RandomConcurrency() string {
	concurrencies := []string{"USD", "CAD", "QUR"}
	k := len(concurrencies) - 1
	return concurrencies[k]
}
