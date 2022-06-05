// Inspired by
// https://pkg.go.dev/github.com/dchest/uniuri
// Slightly modified to better fit this program

package main

import (
	"math/rand"
	"time"
)

var letters = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

func RandomString(length int) string {
	clen := len(letters)
	b := make([]byte, length)
	r := make([]byte, length) // storage for random bytes.

	// if seed is not set it is 1, which makes for not
	// so random numbers
	rand.Seed(time.Now().UnixNano())
	rand.Read(r)

	for i, rb := range r {
		c := int(rb)
		b[i] = letters[c%clen]
	}

	return string(b)
}
