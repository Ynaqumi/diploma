package main

import (
	"diploma2/status/billing"
	"diploma2/status/email"
	"diploma2/status/incidents"
	"diploma2/status/mms"
	"diploma2/status/sms"
	"diploma2/status/support"
	"diploma2/status/voice"
	"fmt"
)

func main() {
	fmt.Println(sms.Sms(), "\n",
		mms.Mms(), "\n",
		voice.VoiceCall(), "\n",
		email.Email(), "\n",
		billing.Billing(), "\n",
		support.Support(), "\n",
		incidents.Incidents())
}
