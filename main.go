package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("************ Welcome to the CLI Password Generator ************")
	menu()

	decisao := readDecision()

	switch decisao {
	case 1:
		var passwordLength int
		fmt.Scan("How long do you want the password to be? ", &passwordLength)

	case 2:
		fmt.Println("Exibindo Logs....")
	case 3:
		fmt.Println("Saindo do programa")
		os.Exit(0)

	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}

func menu() {

	fmt.Scan("Select the desired option: ")

	fmt.Println("1- Generate Password")
	fmt.Println("2- Check Generated Passwords")
	fmt.Println("3- Exit")
	fmt.Println(" ")

}

func readDecision() int {
	var menuDecision int
	fmt.Scan(&menuDecision)
	fmt.Println("O comando escolhido foi", menuDecision)

	return menuDecision
}
