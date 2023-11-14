package main

import (
	"fmt"
	"log"
)

const (
	MAX_SHIFT = 25
	Z_ASCII = 122
	A_ASCII = 97
)
func cryptography(input string, shift int, action string) string {

	if action == "encrypt" {
		return encrypt(input, shift)
	} else {
		return decrypt(input, shift)
	}
}

func encrypt(input string, shift int) string {
	var encrypted string

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

func decrypt(input string, shift int) string {
	var decrypted string

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

func main() {
	input := "zzz"
	shift := 1

	if shift > MAX_SHIFT {
		log.Fatalf("shift is bigger than max: %d", MAX_SHIFT)
	}

	encryptedText := cryptography(input, shift, "encrypt")
	decrypted := cryptography(encryptedText, shift, "decrypt")

	log.Printf("encrypted: %s", encryptedText)
	log.Printf("decrypted: %s", decrypted)
	log.Printf("bool: %t", decrypted == input)
}
