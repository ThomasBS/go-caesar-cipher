package cipher

import "testing"

func TestCan_Shift_4_Characters_In_A_String(t *testing.T) {
	caesarCipher := NewCaesarCipher()
	var word = "Hello world"
	// Each character is shifted 4 times.
	var expected = "Lipps asvph"
	cipherText := caesarCipher.Encrypt(word, 4)
	if cipherText != expected {
		t.Error("Expected "+expected+", got", cipherText)
	}
}

func TestCan_Shift_4_Characters_In_A_Uppercase_String(t *testing.T) {
	caesarCipher := NewCaesarCipher()
	var word = "HELLO WORLD"
	// Each character is shifted 4 times.
	var expected = "LIPPS ASVPH"
	cipherText := caesarCipher.Encrypt(word, 4)
	if cipherText != expected {
		t.Error("Expected "+expected+", got", cipherText)
	}
}

func TestCan_Shift_X_Shifts_In_A_String(t *testing.T) {
	caesarCipher := NewCaesarCipher()
	// The W is the tricky part
	var word = "HELLO WORLD"
	// Each character is shifted 4 times.
	var expected = "URYYB JBEYQ"
	cipherText := caesarCipher.Encrypt(word, 13)
	if cipherText != expected {
		t.Error("Expected "+expected+", got", cipherText)
	}
}