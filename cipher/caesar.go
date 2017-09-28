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
	Crack(str string) string
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

func (c *caesarCipher) Crack(str string) string {
	// The frequencies of letters appearing in the English language sorted from a-z
	// http://crypto.interactive-maths.com/frequency-analysis-breaking-the-code.html
	var charFrequencies = []float32{
		.0820, .0150, .0280, .0430, .1270, .0220, .0200, .0610, .0700, .0015, .0080, .0400, .0240,
		.0670, .0750, .0190, .0010, .0600, .0630, .0910, .0280, .0010, .0240, .0015, .0200, .0007,
	}
	var charSums [26]float32

	// Find the frequency of all the letters in the message
	for _, char := range str {
		if char == 32 {
			continue
		}
		var charPos = getCharPos(char)
		// getCharPos is 1 indexed
		charSums[charPos-1] += charFrequencies[charPos-1]
	}

	return "not implemented"
}

// getCharStr will return a character by specifying a position in the alphabet
func getCharStr(alphabetPos int, asciiVal int32) string {
	var isLowercase = asciiVal >= 97
	if isLowercase {
		return string('a' - 1 + alphabetPos)
	}
	return string('A' - 1 + alphabetPos)
}

// getCharPos will return the character's position in the alphabet
func getCharPos(asciiVal int32) int {
	// To convert a uppercase into lowercase and vice versa we simply set the 6'th bit
	// The code below ensures that bit 32 is set using the bitwise OR operation
	// This eliminates a if statement :)
	asciiVal |= 32
	return int(asciiVal) - 96
}