package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <name>")
		return
	}

	fmt.Printf("Goodbye, %s!\n", os.Args[1])
}
