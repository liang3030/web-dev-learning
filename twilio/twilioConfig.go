package twilio

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type TwilioConfig struct {
	AccountSid       string
	AccountToken     string
	TwilioFromNumber string
}

func (twc *TwilioConfig) SendMessage(to string, message string) {
	accountSid := twc.AccountSid
	accountToken := twc.AccountToken
	twilioFromNumber := twc.TwilioFromNumber

	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// Build out the data for our message
	v := url.Values{}
	v.Set("To", to)
	v.Set("From", twilioFromNumber)
	v.Set("Body", message)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, accountToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)
}
