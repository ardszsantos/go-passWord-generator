package main

import (
	"fmt"
)

func main() {
	fmt.Println("************ Welcome to the CLI Password Generator ************")
	menu()
}

func menu() {
	fmt.Println("Select the desired option: ")

	var passwordLength int

	resposta := fmt.Scan("How long do you want the password to be? ", &passwordLength)
}
