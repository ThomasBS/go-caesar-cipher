package cipher

import "testing"

func TestCan_Shift_Letter_H(t *testing.T) {
	caesarCipher := NewCaesarCipher()
	var word = "H"
	var expected = "L"
	cipherText := caesarCipher.Encrypt(word, 4)
	if cipherText != expected {
		t.Error("Expected "+expected+", got", cipherText)
	}
}

func TestCan_Shift_Characters_With_A_Shift_Key_Of_4(t *testing.T) {
	caesarCipher := NewCaesarCipher()
	var word = "Hello world"
	var expected = "Lipps asvph"
	cipherText := caesarCipher.Encrypt(word, 4)
	if cipherText != expected {
		t.Error("Expected "+expected+", got", cipherText)
	}
}

func TestCan_Shift_A_Uppercase_String_With_A_Shift_key_Of_4(t *testing.T) {
	caesarCipher := NewCaesarCipher()
	var word = "HELLO WORLD"
	var expected = "LIPPS ASVPH"
	cipherText := caesarCipher.Encrypt(word, 4)
	if cipherText != expected {
		t.Error("Expected "+expected+", got", cipherText)
	}
}

func TestCan_Shift_With_A_Custom_Shift_key(t *testing.T) {
	caesarCipher := NewCaesarCipher()
	var word = "HELLO WORLD"
	var expected = "URYYB JBEYQ"
	cipherText := caesarCipher.Encrypt(word, 13)
	if cipherText != expected {
		t.Error("Expected "+expected+", got", cipherText)
	}
}