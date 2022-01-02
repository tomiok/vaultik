package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

//reduce will return a md5 - 128 bits key, or an error if any problem occurs.
func reduce(s string) ([]byte, error) {
	h := md5.New()

	_, err := h.Write([]byte(s))

	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func encrypt(key, value string) (string, error) {
	cipherKey, err := reduce(key)

	if err != nil {
		return "", err
	}

	// cipherKey must be 16, 32 or 48 bytes long
	// 16 bytes (or 128 bits) for this example
	block, err := aes.NewCipher(cipherKey)

	if err != nil {
		return "", err
	}

	plainTextValue := []byte(value)

	// Initialization Vector (IV) needs to be unique but not secure
	cipherText := make([]byte, aes.BlockSize+len(plainTextValue))

	// Put the IV at the beginning of the slice
	iv := cipherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Get the stream for cipher in feedback mode
	stream := cipher.NewCFBEncrypter(block, iv)

	// Add the plain text to the last part from the IV
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainTextValue)

	// Format in base 16
	return fmt.Sprintf("%x", cipherText), nil
}

func decrypt(key, cipherText string) (string, error) {
	// get a 128-bit key
	cipherKey, err := reduce(key)
	if err != nil {
		return "", err
	}
	// decode the cipher text in an byte slice
	text, err := hex.DecodeString(cipherText)

	if err != nil {
		return "", err
	}

	if len(text) == 0 {
		return "", errFileEmpty
	}

	// create a new bock cipher
	block, err := aes.NewCipher(cipherKey)

	if err != nil {
		return "", err
	}

	// same as encrypt
	if len(text) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]

	// Get the stream for cipher in feedback mode

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(text, text)

	// return as readable plain text
	return string(text), nil
}
