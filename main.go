package main

import (
	"fmt"
	"log"
	"os"
	"password-cracker/md5"
	"password-cracker/passwords"
)

func main() {
	fmt.Println("Password Cracker")

	// var password string
	// fmt.Printf("Enter password: ")
	// fmt.Scanln(&password)
	// md5.Init()

	// hash := md5.Md5("abc123")

	// newHash := md5Encode.Md5Encode("abc123")
	// fmt.Println("Password:", password)
	// fmt.Printf("Hash: %s\nNew Hash: %s\nMatching: %v", hash, newHash, hash == newHash)

	targetHashes := []string{"7a95bf926a0333f57705aeac07a362a2",
		"08054846bbc9933fd0395f8be516a9f9"}

	generatedPasswords := passwords.GeneratePasswords(4)

	for _, password := range generatedPasswords {
		hash := md5.Md5(password)

		for _, targetHash := range targetHashes {

			if hash == targetHash {
				fmt.Printf("Password found: %s\n", password)
				return
			}
		}
	}

}

func readFromWordList(filepath string) ([]byte, error) {

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal("Error opening file")
		return nil, err
	}

	defer file.Close()

	return nil, nil
}
