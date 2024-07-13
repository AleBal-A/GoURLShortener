package random

import (
	"math/rand"
	"time"
)

func NewRandomString(length int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	charset := []rune("abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789")

	res := make([]rune, length)

	for i := range res {
		res[i] = charset[rnd.Intn(len(charset))]
	}
	return string(res)
}
