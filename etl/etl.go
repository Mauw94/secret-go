package etl

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"io"
	"os"

	"github.com/Mauw94/secret_go/utils"
)

func ReadData() map[string]string {
	message := make(map[string]string)

	data, err := os.ReadFile("./inputs/message_one.json")
	utils.LogErrors(err)

	json.Unmarshal([]byte(data), &message)

	return message
}

func EncodeData(message string, salt []byte) map[string][]byte {
	text, key := []byte(message), salt

	c, err := aes.NewCipher(key)
	utils.LogErrors(err)

	gcm, err := cipher.NewGCM(c)
	utils.LogErrors(err)

	nonce := make([]byte, gcm.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)
	utils.LogErrors(err)

	finalText := gcm.Seal(nonce, nonce, text, nil)

	return map[string][]byte{
		"passphrase": key,
		"message":    finalText,
	}
}

func DecodeData(encryptedMessage map[string][]byte) string {
	key := encryptedMessage["passphrase"]
	cipherText := encryptedMessage["message"]

	c, err := aes.NewCipher(key)
	utils.LogErrors(err)

	gcm, err := cipher.NewGCM(c)
	utils.LogErrors(err)

	nonceSize := gcm.NonceSize()

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	utils.LogErrors(err)

	return string(plainText)
}
