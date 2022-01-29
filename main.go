package main

import (
	"fmt"
	"os"
	"strings"
)

func wordCount(str string) int {
	return len(strings.Fields(str))
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Err - there is no arguments!")
		os.Exit(-1)
	}

	fmt.Println(wordCount(args[1]))
}
