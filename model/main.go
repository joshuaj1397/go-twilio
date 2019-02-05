package main

import (
	"encoding/json"
	"net/http"
	"os"
  "strings"
  "net/url"
  "fmt"
)

// Get the TWILIO_API_KEY from env vars
func main() {

  // Set account keys and info
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
  authToken := os.Getenv("TWILIO_AUTH_TOKEN")
  phoneNumber := os.Getenv("TWILIO_PHONE_NUMBER")
  urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
	if (accountSid == "" || authToken == "") {
		panic("TWILIO CREDENTIALS NOT FOUND IN ENV VARS")
	}

  // Produce a message
  msgData := url.Values{}
  msgData.Set("To","+12142125920")
  msgData.Set("From", phoneNumber)
  msgData.Set("Body", "A Test Message")
  msgDataReader := *strings.NewReader(msgData.Encode())

  // POST: /Accounts/[AccountSid]/Messages
  client := &http.Client{}
  req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
  req.SetBasicAuth(accountSid, authToken)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

  // Make the request
  resp, _ := client.Do(req)
  if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
    var data map[string]interface{}
    decoder := json.NewDecoder(resp.Body)
    err := decoder.Decode(&data)
    if (err == nil) {
      fmt.Println(data["sid"])
    }
  } else {
    fmt.Println(resp.Status);
  }
}
