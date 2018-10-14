package main

import (
	"crypto/rand"
	"fmt"
)

func uuid() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b[0:4])
}
