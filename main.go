package main

import (
	"github.com/steffen25/go-caesar-cipher/cipher"
	"log"
)

func main()  {
	caesarCipher := cipher.NewCaesarCipher()
	log.Println(caesarCipher.Encrypt("Hello world", 4))
	log.Println(caesarCipher.Decrypt("Lipps asvph", 4))
}
