package handlers

import (
	"math/rand"
)

func RandomString(n int) string {
	var numbers = []rune("0123456789")
	s := make([]rune, n)
	for i := range s {

		s[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(s)
}
