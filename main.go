package main

import (
	"crypto/rand"
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
		fmt.Println("How long do you want the password to be? ")
		fmt.Scan(&passwordLength)
		fmt.Println("Selected size: ", passwordLength)
		fmt.Println("")
		fmt.Println("Generating...")
		generatePassword(passwordLength)

	case 2:

	case 3:
		fmt.Println("Exiting. . . . . .")
		os.Exit(0)

	default:
		fmt.Println("Unkown Option. . . . . . .")
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
	fmt.Println("Selected option: ", menuDecision)

	return menuDecision
}

func generatePassword(size int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	password := make([]byte, size)

	for i := 0; i < size; i++ {
		randomByte := make([]byte, 1)
		rand.Read(randomByte)
		password[i] = charset[int(randomByte[0])%len(charset)]
	}

	fmt.Println(string(password))

	return string(password)
}
