package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("************ Welcome to the CLI Password Generator ************")
	for {
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
			readPasswords()
		case 3:
			fmt.Println("Exiting. . . . . .")
			os.Exit(0)

		default:
			fmt.Println("Unkown Option. . . . . . .")
			os.Exit(-1)
		}
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

	fmt.Println("Generated password:", string(password), "\n", "Saving it to saved_passwords.txt....")

	var file *os.File
	if _, err := os.Stat("saved_passwords.txt"); os.IsNotExist(err) {
		file, _ = os.Create("saved_passwords.txt")
	} else {
		file, _ = os.OpenFile("saved_passwords.txt", os.O_APPEND|os.O_WRONLY, 0644)
	}
	defer file.Close()

	// Write the password followed by a newline
	file.WriteString(string(password) + "\n")
	return string(password)
}

func readPasswords() {
	data, _ := os.ReadFile("saved_passwords.txt")
	lines := strings.Split(string(data), "\n")
	fmt.Println("Saved passwords: ")
	fmt.Println("")
	for _, line := range lines {
		if line != "" {
			fmt.Println(line)
			fmt.Println("")
		}
	}

}
