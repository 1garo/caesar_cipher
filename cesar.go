package cesar_cipher

import (
	"fmt"
	"log"
)

type Method string

const (
	MAX_SHIFT = 25
	Z_ASCII = 122
	A_ASCII = 97
	ENCRYPT Method = "encrypt"
	DECRYPT Method = "decrypt"
)

/* Cesar returns encrypted or decrypted string based on `action` param.
	e.g:
		encrpyt -> encryptedText, err := Cesar(input, "encrypt", 5)
		decrypt -> decrypted, err := Cesar(encryptedText, "decrypt", 5)
*/
func Cesar(input string, action Method, shift int) (string, error) {
	if shift > MAX_SHIFT {
		log.Fatalf("shift is bigger than max: %d", MAX_SHIFT)
	}

	if action == "encrypt" {
		return encrypt(input, shift), nil
	} else if action == "decrypt" {
		return decrypt(input, shift), nil
	} else {
		return "", fmt.Errorf("unknown action: '%s'", action)
	}
}

// encrypt returns a encrypted string based on input and shift
func encrypt(input string, shift int) string {
	var encrypted string

	// TODO: make this works for CAPTIAL LETTERS
	for i := 0; i < len(input); i++ {
		ascii := int(input[i]) + shift
		if ascii > Z_ASCII {
			remainer := int(input[i]) - Z_ASCII
			ascii = A_ASCII + remainer
		}

		encrypted += fmt.Sprintf("%c", rune(ascii))
	}

	return encrypted
}

// decrypt returns the string decrypted
func decrypt(input string, shift int) string {
	var decrypted string

	// TODO: make this works for CAPTIAL LETTERS
	for i := 0; i < len(input); i++ {
		ascii := int(input[i]) - shift
		if ascii < A_ASCII {
			remainer := Z_ASCII - int(input[i])
			ascii = A_ASCII + remainer
		}

		decrypted += fmt.Sprintf("%c", rune(ascii))
	}

	return decrypted
}

