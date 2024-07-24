package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
)

type Password struct {
	Size            int
	Capital_letters bool
	Lowercase       bool
	Numbers         bool
	Symbols         bool
}

func generatePassword(password Password) string {

	characters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()")

	random := rand.Reader

	buffer := make([]byte, password.Size)

	_, err := random.Read(buffer)
	if err != nil {
		panic(err)
	}

	passwordHex := hex.EncodeToString(buffer)

	filteredPassword := ""
	for _, char := range passwordHex {
		if (password.Capital_letters && strings.ToUpper(string(char)) != char) ||
			(!password.Lowercase && strings.ToLower(string(char)) != char) ||
			(!password.Numbers && !strings.Range(string(char), "0", "9")) ||
			(!password.Symbols && !strings.Contains(string(characters), string(char))) {
			continue
		}
		filteredPassword += string(char)
	}

	return filteredPassword[:password.Size]
}

func main() {

	myPassword := Password{
		Size:            16,
		Capital_letters: true,
		Lowercase:       true,
		Numbers:         true,
		Symbols:         true,
	}

	generatedPassword := generatePassword(myPassword)

	fmt.Println("Generated password:", generatedPassword)
}
