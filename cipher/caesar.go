package cipher

import (
	"bytes"
)

// Caesar cipher encryption:
// Replace each letter in the message by its numeric equivalence
// mod 26
type CaesarCipherInterface interface {
	Encrypt(str string, shiftKey int) string
	Decrypt(str string, shiftKey int) string
}

type caesarCipher struct {}

func NewCaesarCipher() CaesarCipherInterface {
	return &caesarCipher{}
}

func (c *caesarCipher) Encrypt(str string, shiftKey int) string {
	var cipherText bytes.Buffer
	// ASCII 97-122 means lowercase
	// ASCII 65-90 means uppercase
	// cipherLetter = (origChar + shiftKey) % 26
	// e.g. "H" cipherLetter = (8+4)%26 : cipherLetter=12 : cipherLetter=l
	for _, char := range str {
		// ignore space
		if char == 32 {
			cipherText.WriteByte(byte(char))
			continue
		}

		var charPos = getCharPos(char)
		var cipherChar = getCharStr((charPos+shiftKey) % 26, char)
		cipherText.WriteString(cipherChar)
	}

	return cipherText.String()
}

func (c *caesarCipher) Decrypt(str string, shiftKey int) string {
	var cipherText bytes.Buffer

	for _, char := range str {
		// ignore space
		if char == 32 {
			cipherText.WriteByte(byte(char))
			continue
		}

		var charPos = getCharPos(char)
		var mod26 = (charPos-shiftKey) % 26
		if mod26 < 0 {
			mod26 += 26
		}
		var origChar = getCharStr(mod26, char)
		cipherText.WriteString(origChar)
	}

	return cipherText.String()
}

// getCharStr will return a character by specifying a position in the alphabet
func getCharStr(alphabetPos int, asciiVal int32) string {
	var isLowercase = asciiVal >= 97
	if isLowercase {
		return string('a' - 1 + alphabetPos)
	}
	return string('A' - 1 + alphabetPos)
}

// getCharPos will return the character's position in the alphabet based on its ascii value.
func getCharPos(asciiVal int32) int {
	var index int
	var isUppercase = asciiVal < 97

	if index = int(asciiVal) - 96; isUppercase {
		index = int(asciiVal) - (64)
	}

	return index
}