package cesar_cipher

import "testing"


// Add test cases
func TestCesarEncrypt(t *testing.T) {
	input := "aaa"
	shift := 1
	expected := "bbb"
	encrypted, err := Cesar(input, "encrypt", shift)
	if err != nil {
		t.Fatalf("failed with error: %s", err)
	}

	if encrypted != expected {
		t.Fatalf("received: %s expected: %s", encrypted, expected)
	}
}

func TestCesarDecrypt(t *testing.T) {
	input := "aaa"
	shift := 1
	expected := "bbb"

	encryptedText, err := Cesar(input, "encrypt", shift)
	if err != nil {
		t.Fatalf("failed with error: %s", err)
	}

	if encryptedText != expected {
		t.Fatalf("received: %s expected: %s", encryptedText, expected)
	}

	decrypt, err := Cesar(encryptedText, "decrypt", shift)
	if err != nil {
		t.Fatalf("failed with error: %s", err)
	}

	if decrypt != input {
		t.Fatalf("received: %s expected: %s", decrypt, input)
	}

}

func TestUnknownActionType(t *testing.T) {
	input := "aaa"
	shift := 1
	_, err := Cesar(input, "unknown", shift)

	if err == nil {
		t.Fatal("error should not be nil")
	}
}

