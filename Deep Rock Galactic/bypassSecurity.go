package main

import "fmt"

// Caesar Cipher Encryption Function
func caesarEncrypt(plainText string, shift int) string {
	cipherText := ""

	for _, char := range plainText {
		// Encrypt uppercase letters
		if char >= 'A' && char <= 'Z' {
			cipherText += string((char+rune(shift)-'A')%26 + 'A')
		// Encrypt lowercase letters
		} else if char >= 'a' && char <= 'z' {
			cipherText += string((char+rune(shift)-'a')%26 + 'a')
		// Keep other characters unchanged
		} else {
			cipherText += string(char)
		}
	}

	return cipherText
}

// Caesar Cipher Decryption Function
func caesarDecrypt(cipherText string, shift int) string {
	plainText := ""

	for _, char := range cipherText {
		// Decrypt uppercase letters
		if char >= 'A' && char <= 'Z' {
			plainText += string((char-rune(shift)-'A'+26)%26 + 'A')
		// Decrypt lowercase letters
		} else if char >= 'a' && char <= 'z' {
			plainText += string((char-rune(shift)-'a'+26)%26 + 'a')
		// Keep other characters unchanged
		} else {
			plainText += string(char)
		}
	}

	return plainText
}

func main() {
	plainText := "THIS_IS_AN_ENCRYPTED_PASSWORD"
	shift := 3

	// Encrypt the plaintext
	cipherText := caesarEncrypt(plainText, shift)

	fmt.Printf("Bypassing Security...\n")
	fmt.Printf("Ciphertext: %s\n", cipherText)

	// Decrypt the ciphertext
	plainText = caesarDecrypt(cipherText, shift)
	fmt.Printf("Plaintext: %s\n", plainText)
}
