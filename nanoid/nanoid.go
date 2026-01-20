package nanoid

import (
	"bytes"
	"math/rand"
	"time"
)

const defaultAlphabet = "_-23456789abcdefghjkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// New generates a random nanoid with length of 4
func New() string {
	return NewWithLength(4)
}

// Generate generates a random nanoid with a custom length; 4 <= length <= 200
// If the size was out of that range, it defaults to 4.
func NewWithLength(size int) string {
	if size < 4 || size > 200 {
		size = 4
	}
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < size; i++ {
		buf.WriteByte(defaultAlphabet[random.Intn(len(defaultAlphabet))])
		random.Seed(time.Now().UnixNano())
	}
	return buf.String()
}
