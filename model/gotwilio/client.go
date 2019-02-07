package gotwilio

import (
	"encoding/json"
	"net/http"
	"os"
  "strings"
  "net/url"
  "fmt"
)

// Twilio account credentials are set in environment variables:
// TWILIO_ACCOUNT_SID, TWILIO_AUTH_TOKEN, and TWILIO_PHONE_NUMBER
type Credentials struct {
	accountSid	string
	authToken		string
	phoneNumber string
	baseUrl			string
}

var (
	credentials Credentials

	// Client would instantiate a new http.Client with a 5 sec timeout on requests
	client = http.Client{
		Timeout: time.Second * 5,
	}
)

// Get the Twilio credentials from env vars
func init() {
	// Set account keys and info
	credentials = Credentials {
		accountSid: os.Getenv("TWILIO_ACCOUNT_SID"),
		authToken: os.Getenv("TWILIO_AUTH_TOKEN"),
		phoneNumber: os.Getenv("TWILIO_PHONE_NUMBER"),
		baseUrl: "https://api.twilio.com/2010-04-01/Accounts/"
	}
	if (credentials.accountSid == "" || credentials.authToken == "" || credentials.phoneNumber == "") {
		panic("TWILIO CREDENTIALS NOT FOUND IN ENV VARS")
	}
}

// POST: /Accounts/[AccountSid]/Messages
//
// Sends a message using the Twilio SMS API
func SendMsg(accountCredentials interface{}, recipient, body string) int {
		urlStr := credentials.baseUrl + credentials.accountSid + "/Messages.json"

	  // Produce a message
	  msgData := url.Values{}
	  msgData.Set("To", recipient)
	  msgData.Set("From", credentials.phoneNumber)
	  msgData.Set("Body", body)
	  msgDataReader := *strings.NewReader(msgData.Encode())

	  req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	  req.SetBasicAuth(credentials.accountSid, credentials.authToken)
	  req.Header.Add("Accept", "application/json")
	  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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
