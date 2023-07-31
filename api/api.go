package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Mauw94/secret_go/utils"
	"github.com/joho/godotenv"
)

type PostData struct {
	EncryptedMessage string `json:"encryptedMessage"`
	Passphrase       string `json:"passphrase"`
}

type PostBody struct {
	EventID   string   `json:"eventID"`
	EventType string   `json:"eventType"`
	Data      PostData `json:"data"`
}

func Call(method string, url string, data PostBody) map[string]string {
	err := godotenv.Load()
	utils.LogErrors(err)

	myEnv, err := godotenv.Read()
	utils.LogErrors(err)

	type formattedData struct {
		Events []PostBody `json:"events" `
	}
	var events []PostBody
	events = append(events, data)

	finalData := &formattedData{Events: events}

	jsonData, err := json.Marshal(finalData)
	utils.LogErrors(err)

	client := &http.Client{}
	req, err := http.NewRequest(
		method,
		url,
		bytes.NewReader(jsonData),
	)
	utils.LogErrors(err)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", " application/json")
	req.Header.Add("Serialized-Access-Key", myEnv["ACCESS_KEY"])
	req.Header.Add("Serialized-Secret-Access-Key", myEnv["SECRET_ACCESS_KEY"])

	resp, err := client.Do(req)
	utils.LogErrors(err)

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	utils.LogErrors(err)

	var responseObject map[string]string
	json.Unmarshal(bodyBytes, &responseObject)

	return responseObject
}
