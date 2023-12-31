package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func getUsername() string {
	fmt.Print("Enter username: ")
	var username string
	fmt.Scanln(&username)
	return username
}

func getPassword() string {
	fmt.Print("Enter password: ")

	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("Error reading password.", err)
		os.Exit(1)
	}

	return string(bytePassword)
}

func returnEncoded(username, password string) string {
	credentials := fmt.Sprintf("%s:%s", username, password)
	encoded := base64.StdEncoding.EncodeToString([]byte(credentials))
	return encoded
}

func main() {
	username := getUsername()
	password := getPassword()
	encoded := returnEncoded(username, password)

	fmt.Println("\nBase64 encoded:")
	fmt.Println(strings.TrimSpace(encoded))
}
