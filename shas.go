package main //Write a program that prints sha256 hash of its standard input by default
// support command line flag to print sha384 or sha512

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	var args string
	if len(os.Args) > 1 {
		args := args + os.Args[1]
	} else {
		args := ""
	}

	fmt.Println("Enter item to hash:")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	i := input.Text()

	switch args {
	case "SHA384":
		fmt.Printf("Result: %x\n", sha512.Sum384([]byte(i)))
	case "SHA512":
		fmt.Printf("Result: %x\n", sha512.Sum512([]byte(i)))
	default:
		fmt.Printf("Result: %x\n", sha256.Sum256([]byte(i)))
	}
	fmt.Println("Exiting")
}
