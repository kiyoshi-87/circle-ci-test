package main

import (
	"fmt"
	"strings"
)

func ConvertToUpperCase(input string) string {
	return strings.ToUpper(input)
}

func main() {
	fmt.Print("Enter a string: ")
	var input string

	fmt.Scanln(&input)

	upperCase := ConvertToUpperCase(input)

	fmt.Println("Uppercase:", upperCase)
}
