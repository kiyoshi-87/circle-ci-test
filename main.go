package main

import (
	"fmt"
	"strings"
)

func ConvertToUpperCase(input string) string {
	return strings.ToUpper(input)
}

func main() {
	fmt.Print("Enter a string that you want to convert from lowercase to uppercase: ")
	var input string

	fmt.Scanln(&input)

	upperCase := ConvertToUpperCase(input)

	fmt.Println("Your uppercase word is:", upperCase)
}
