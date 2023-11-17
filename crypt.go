package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	"golang.org/x/crypto/argon2"
)

const key = "d27b92783f98d466484fc8ace1615e10"

func Encrypt(text string) string {
	block, _ := aes.NewCipher([]byte(key))
	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(text))
	return fmt.Sprintf("%x", ciphertext)
}

func Decrypt(text string) string {
	ciphertext, _ := hex.DecodeString(text)
	block, _ := aes.NewCipher([]byte(key))
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext)
}

func GenerateRandomSalt() ([]byte, error) {
	salt := make([]byte, 16) // Adjust salt length as needed
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func DerivePassword(password string) string {
	const fixedSalt = "00112233445566778899aabbccddeeff"
	salt, err := hex.DecodeString(fixedSalt)
	if err != nil {
		return "Error decoding salt"
	}
	timeCost := 1
	memory := 64 * 1024
	threads := 4
	keyLength := 32

	derivedKey := argon2.IDKey([]byte(password), salt, uint32(timeCost), uint32(memory), uint8(threads), uint32(keyLength))
	derivedKeyHex := hex.EncodeToString(derivedKey)
	return derivedKeyHex
}
