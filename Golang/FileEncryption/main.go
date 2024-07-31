package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/Syabil76/Golang/FileEncryption/filecrypt"
	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function {
	case "--h":
		printHelp()
	case "-en":
		EncryptFile()
	case "-d":
		DecryptFile()
	default:
		fmt.Println("Run --h for notes, -en to encrypt and -d to decrypt")
		os.Exit(1)
	}
}

func printHelp() {
	filepath := "data.txt"
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("File reading error!", err)
		return
	}
	fmt.Println(string(data))
}

func EncryptFile() {
	if len(os.Args) < 3 {
		fmt.Println("err: missing path to file. Run go run . --h for help.")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")
	}
	password := getPassword()
	fmt.Println("\n encrypting...")
	filecrypt.Encrypt(file, password)
	fmt.Println("\n file successfuly encrypted")
}

func DecryptFile() {
	if len(os.Args) < 3 {
		fmt.Println("err: missing path to file. Run go run . --h for help.")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")
	}
	fmt.Print("\n enter password")
	password, _ := term.ReadPassword(0)
	fmt.Println("decrypting...")
	filecrypt.Decrypt(file, password)
	fmt.Print("\n File successfuly decrypted.")
}

func getPassword() []byte {
	fmt.Print("Enter password: ")
	password, _ := term.ReadPassword(0)
	fmt.Print("\n Confirm password: ")
	password2, _ := term.ReadPassword(0)
	if !validatePassword(password, password2) {
		fmt.Print("\n Passwords do not match, please try again\n")
		return getPassword()
	}
	return password
}

func validatePassword(password []byte, password2 []byte) bool {
	return bytes.Equal(password, password2)

}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
