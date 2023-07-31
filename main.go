package main

import (
	"fmt"

	"github.com/Mauw94/secret_go/api"
	"github.com/Mauw94/secret_go/etl"
	"github.com/Mauw94/secret_go/utils"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("Running!")

	rawMessage := etl.ReadData()

	fmt.Println("Message from filereader is: ")
	fmt.Println(rawMessage)

	passPhrase := utils.MakePassphrase()
	fmt.Println("Random passphrase is: ")
	fmt.Println(string(passPhrase))

	encryptedMessage := etl.EncodeData(rawMessage["body"], passPhrase)
	fmt.Println("Encrypted message and passphrase are: ")
	fmt.Println(encryptedMessage)
	fmt.Println(encryptedMessage["passphrase"])

	decryptedMessage := etl.DecodeData(encryptedMessage)
	fmt.Println("Decrypted message is: ")
	fmt.Println(decryptedMessage)

	newUUID := uuid.New().String()
	url := "https://api.serialized.io/aggregates/new_message/" + newUUID + "/events"

	messageData := api.PostData{
		EncryptedMessage: string(encryptedMessage["message"]),
		Passphrase:       string(encryptedMessage["passphrase"]),
	}

	eventBody := api.PostBody{
		EventID:   newUUID,
		EventType: "new_message",
		Data:      messageData,
	}

	responseObject := api.Call(
		"POST",
		url,
		eventBody,
	)

	fmt.Println("Response object from Serialized is:")
	fmt.Printf("%+v\n", responseObject)
}
