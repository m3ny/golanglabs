package main //Write a program that prints sha256 hash of its standard input by default
// support command line flag to print sha384 or sha512

import (
	"crypto/sha256"
)

func main() {

	c1 := sha256.Sum256([]byte("x"))
}
