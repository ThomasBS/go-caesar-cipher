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

type caesarCipher struct {

}

func NewCaesarCipher() CaesarCipherInterface {
	return &caesarCipher{}
}

func (c *caesarCipher) Encrypt(str string, shiftKey int) string {
	var cipherText bytes.Buffer
	
	for _, char := range str {
		var val = char + int32(shiftKey)
		// ignore space
		if char == 32 {
			cipherText.WriteByte(byte(char))
			continue
		}
		if char < 90 && val > 90 {
			cipherText.WriteByte(byte((val - 26)))
			continue
		}
		if val > 122 {
			cipherText.WriteByte(byte((val - 26)))
			continue
		}

		if val < 65 {
			cipherText.WriteByte(byte((val + 26)))
			continue
		}

		cipherText.WriteByte(byte((val)))
	}

	return cipherText.String()
}

func (c *caesarCipher) Decrypt(str string, shiftKey int) string {
	return "not implemented"
}